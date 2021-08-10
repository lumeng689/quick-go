package etcd

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func Test_EtcdConnect001(t *testing.T) {
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}

	var client *clientv3.Client
	var err error

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("connect success")
	defer client.Close()
}

func GetEtcdClient() *clientv3.Client {
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println("connect success")
	}
	return client
}

/*
  测试设置值
*/
func Test_Put001(t *testing.T) {
	client := GetEtcdClient()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	rsp, err := client.Put(ctx, "/demo/demo1_key", "demo1_value")
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	fmt.Println("response::", rsp.Header.String())
}

func Test_Put002(t *testing.T) {
	client := GetEtcdClient()
	defer client.Close()

	wg := sync.WaitGroup{}

	client.Put(context.Background(), "/demo/demo2_key", "demo2_value")

	wg.Add(1)

	go func() {
		//watch 返回值是个channel
		watchKey := client.Watch(context.Background(), "/demo/demo2_key")
		for resp := range watchKey {
			for _, item := range resp.Events {
				fmt.Printf("events:::%s %q : %q \n", item.Type, item.Kv.Key, item.Kv.Value)
			}
		}
		wg.Done()
		fmt.Println("done....")
	}()

	time.Sleep(1 * time.Second)
	if resp, err := client.Put(context.TODO(), "/demo/demo2_key", "demo2_watch"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("resp", resp)
	}
	fmt.Println("waiting...")
	wg.Wait()
	fmt.Println("bye byte...")
}

func Test_Put003(t *testing.T) {
	client := GetEtcdClient()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	putOp := clientv3.OpPut("/demo/demo1_key", "demo1_value")

	opResp, err := client.Do(ctx, putOp)

	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	fmt.Println("response::", opResp)
}

/*
  获取值
*/
func Test_Get001(t *testing.T) {
	client := GetEtcdClient()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	resp, err := client.Get(ctx, "/demo/demo1_key")
	cancel()
	if err != nil {
		fmt.Println("get failed err:", err)
		return
	}

	for _, item := range resp.Kvs {
		fmt.Println("item key:", string(item.Key))
		fmt.Println("item value:", string(item.Value))
	}
}

func Test_Get002(t *testing.T) {
	client := GetEtcdClient()
	defer client.Close()
}

func Test_GetAll001(t *testing.T) {
	client := GetEtcdClient()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	resp, err := client.Get(ctx, "/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println("get failed err:", err)
		return
	}

	for _, item := range resp.Kvs {
		fmt.Println("item key:", string(item.Key))
		fmt.Println("item value:", string(item.Value))
	}
}

func Test_GetAll002(t *testing.T) {
	client := GetEtcdClient()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	resp, err := client.Get(ctx, "/")
	cancel()
	if err != nil {
		fmt.Println("get failed err:", err)
		return
	}

	for _, item := range resp.Kvs {
		fmt.Println("item key:", string(item.Key))
		fmt.Println("item value:", string(item.Value))
	}
}

/*
  删除
*/
func Test_Delete002(t *testing.T) {
	client := GetEtcdClient()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	resp, err := client.Delete(ctx, "/demo/demo1_key")
	cancel()
	if err != nil {
		fmt.Println("get failed err:", err)
		return
	}

	fmt.Println("deleted items:", resp.Deleted)
}

/*
  事务
*/
func Test_Transaction001(t *testing.T) {

}

/*
  租约
*/
func Test_Lease001(t *testing.T) {
	client := GetEtcdClient()
	defer client.Close()

	// 申请一个租约
	lease := clientv3.NewLease(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	leaseGrantResp, err := lease.Grant(ctx, 10)
	if err != nil {
		fmt.Println("create lease error::", err)
		return
	}

	leaseId := leaseGrantResp.ID

	// 获得kv API子集
	kv := clientv3.NewKV(client)

	if _, err = kv.Put(context.TODO(), "/school/class/students", "h", clientv3.WithLease(clientv3.LeaseID(leaseId))); err != nil {
		fmt.Println("put kv error", err)
		return
	}

	for {
		getResp, err := kv.Get(context.TODO(), "/school/class/students")
		if err != nil {
			fmt.Println("get error", err)
			return
		}
		if len(getResp.Kvs) == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期:", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}
