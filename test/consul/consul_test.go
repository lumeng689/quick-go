package consul

import (
	"fmt"
	"testing"

	"github.com/hashicorp/consul/api"
)

func Test_ConsulConnect001(t *testing.T) {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)

	if err != nil {
		fmt.Println("create client failed", err)
	}

	client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    "my-service",
		Address: "127.0.0.1",
		Port:    8888,
	})

	mp, err := client.Agent().Services()

	if err != nil {
		fmt.Println("get services failed", err)
	}

	for k, v := range mp {
		fmt.Printf("%v: %v", k, v)
	}
}
