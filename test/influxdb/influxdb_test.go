package influxdb

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/influxdata/influxdb-client-go/v2"
)

const (
	AuthToken = "my-token"
)

//
//
//
func Test_ConnectDb001(t *testing.T) {
	client := influxdb2.NewClient("http://localhost:8086", AuthToken)
	writeApi := client.WriteAPIBlocking("my-org", "my-bucket")

	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45.0},
		time.Now())

	// write point immediately
	writeApi.WritePoint(context.Background(), p)

	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45.0).
		SetTime(time.Now())

	writeApi.WritePoint(context.Background(), p)

	// 使用行协议
	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	writeApi.WriteRecord(context.Background(), line)

	queryApi := client.QueryAPI("my-org")
	result, err := queryApi.Query(context.Background(), `from(bucket:"my-bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err != nil {
		for result.Next() {
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}

			fmt.Printf("row: %s\n", result.Record().String())
		}

		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	}

	client.Close()
}

func Test_ConnectDb002(t *testing.T) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	// and set batch size to 20
	client := influxdb2.NewClientWithOptions("http://localhost:8086", AuthToken,
		influxdb2.DefaultOptions().SetBatchSize(20))
	// Get non-blocking write client
	writeAPI := client.WriteAPI("my-org", "my-bucket")
	// write some points
	for i := 0; i < 100; i++ {
		// create point
		p := influxdb2.NewPoint(
			"system",
			map[string]string{
				"id":       fmt.Sprintf("rack_%v", i%10),
				"vendor":   "AWS",
				"hostname": fmt.Sprintf("host_%v", i%100),
			},
			map[string]interface{}{
				"temperature": rand.Float64() * 80.0,
				"disk_free":   rand.Float64() * 1000.0,
				"disk_total":  (i/10 + 1) * 1000000,
				"mem_total":   (i/100 + 1) * 10000000,
				"mem_free":    rand.Uint64(),
			},
			time.Now())
		// write asynchronously
		writeAPI.WritePoint(p)
	}
	// Force all unwritten data to be sent
	writeAPI.Flush()
	// Ensures background processes finishes
	client.Close()
}

func Test_ConnectDb003(t *testing.T) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:8086", AuthToken)
	// Get non-blocking write client
	writeAPI := client.WriteAPI("my-org", "my-bucket")
	// Get errors channel
	errorsCh := writeAPI.Errors()
	// Create go proc for reading and logging errors
	go func() {
		for err := range errorsCh {
			fmt.Printf("write error: %s\n", err.Error())
		}
	}()
	// write some points
	for i := 0; i < 100; i++ {
		// create point
		p := influxdb2.NewPointWithMeasurement("stat").
			AddTag("id", fmt.Sprintf("rack_%v", i%10)).
			AddTag("vendor", "AWS").
			AddTag("hostname", fmt.Sprintf("host_%v", i%100)).
			AddField("temperature", rand.Float64()*80.0).
			AddField("disk_free", rand.Float64()*1000.0).
			AddField("disk_total", (i/10+1)*1000000).
			AddField("mem_total", (i/100+1)*10000000).
			AddField("mem_free", rand.Uint64()).
			SetTime(time.Now())
		// write asynchronously
		writeAPI.WritePoint(p)
	}
	// Force all unwritten data to be sent
	writeAPI.Flush()
	// Ensures background processes finishes
	client.Close()
}

func Test_ConnectDb004(t *testing.T) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:8086", AuthToken)
	// Get blocking write client
	writeAPI := client.WriteAPIBlocking("my-org", "my-bucket")
	// write some points
	for i := 0; i < 100; i++ {
		// create data point
		p := influxdb2.NewPoint(
			"system",
			map[string]string{
				"id":       fmt.Sprintf("rack_%v", i%10),
				"vendor":   "AWS",
				"hostname": fmt.Sprintf("host_%v", i%100),
			},
			map[string]interface{}{
				"temperature": rand.Float64() * 80.0,
				"disk_free":   rand.Float64() * 1000.0,
				"disk_total":  (i/10 + 1) * 1000000,
				"mem_total":   (i/100 + 1) * 10000000,
				"mem_free":    rand.Uint64(),
			},
			time.Now())
		// write synchronously
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			panic(err)
		}
	}
	// Ensures background processes finishes
	client.Close()
}

func Test_ConnectDb005(t *testing.T) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:8086", AuthToken)
	// Get query client
	queryAPI := client.QueryAPI("my-org")
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), `from(bucket:"my-bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
	// Ensures background processes finishes
	client.Close()
}

func Test_ConnectDb006(t *testing.T) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:8086", AuthToken)
	// Get query client
	queryAPI := client.QueryAPI("my-org")
	// Query and get complete result as a string
	// Use default dialect
	result, err := queryAPI.QueryRaw(context.Background(), `from(bucket:"my-bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`, influxdb2.DefaultDialect())
	if err == nil {
		fmt.Println("QueryResult:")
		fmt.Println(result)
	} else {
		panic(err)
	}
	// Ensures background processes finishes
	client.Close()
}

func Test_ConnectDb007(t *testing.T) {
	// Create HTTP client
	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(60),
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}
	// Client for server 1
	client1 := influxdb2.NewClientWithOptions("https://server:8086", AuthToken, influxdb2.DefaultOptions().SetHTTPClient(httpClient))
	// Client for server 2
	//client2 := influxdb2.NewClientWithOptions("https://server:9999", "my-token2", influxdb2.DefaultOptions().SetHTTPClient(httpClient))
	// Ensure closing the client
	defer client1.Close()

	// Get write client
	writeApi := client1.WriteAPI("my-org", "my-bucket")

	// Create channel for points feeding
	pointsCh := make(chan *write.Point, 200)

	threads := 5

	var wg sync.WaitGroup
	go func(points int) {
		for i := 0; i < points; i++ {
			p := influxdb2.NewPoint("meas",
				map[string]string{"tag": "tagvalue"},
				map[string]interface{}{"val1": rand.Int63n(1000), "val2": rand.Float64()*100.0 - 50.0},
				time.Now())
			pointsCh <- p
		}
		close(pointsCh)
	}(1000000)

	// Launch write routines
	for t := 0; t < threads; t++ {
		wg.Add(1)
		go func() {
			for p := range pointsCh {
				writeApi.WritePoint(p)
			}
			wg.Done()
		}()
	}
	// Wait for writes complete
	wg.Wait()
}

func Test_ConnectDb008(t *testing.T) {

}
