package main

import (
	"context"
	"fmt"
	"os"

	"github.com/paultyng/go-unifi/unifi"
)

func main() {
	url := os.Getenv("UNIFI_URL")
	user := os.Getenv("UNIFI_USER")
	password := os.Getenv("UNIFI_PASSWORD")
	client := new(unifi.Client)
	if err := client.SetBaseURL(url); err != nil {
		panic(err)
	}

	fmt.Println("Login")
	if err := client.Login(context.Background(), user, password); err != nil {
		panic(err)
	}
	fmt.Println("List networks")
	networks, err := client.ListNetwork(context.Background(), "default")
	if err != nil {
		panic(err)
	}
	for _, n := range networks {
		fmt.Printf("Name: %s, ID: %s\n", n.Name, n.ID)
	}

	// fmt.Println("Create Network")
	// network := unifi.Network{
	// 	Name:        "Test Network",
	// 	Enabled:     true,
	// 	Purpose:     "vlan-only",
	// 	VLAN:        100,
	// 	VLANEnabled: true,
	// }
	// createdNetwork, err := client.CreateNetwork(context.Background(), "default", &network)
	// if err != nil {
	// 	panic(err)
	// }
	// b, err := createdNetwork.MarshalJSON()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b))
}
