package consul

import (
	"fmt"
	"testing"

	"github.com/hashicorp/consul/api"
)

func GetConsulClient() *api.Client {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)

	if err != nil {
		fmt.Println("create client failed", err)
	}
	return client
}

func Test_ConsulConnect001(t *testing.T) {

	client := GetConsulClient()

	client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    "my-service",
		Address: "127.0.0.1",
		Port:    8888,
	})

	mp, err := client.Agent().Services()

	if err != nil {
		fmt.Println("register services failed", err)
	}

	for k, v := range mp {
		fmt.Printf("%v: %v", k, v)
	}
}

func Test_GetService001(t *testing.T) {

	client := GetConsulClient()

	services, metaData, err := client.Catalog().Service("my-service", "", nil)

	if err != nil {
		fmt.Println("get services failed", err)
	}

	fmt.Printf("meta data: %v \n", metaData)

	for _,s := range services {
		fmt.Printf("Address: %v \n", s.Address)
	}
}