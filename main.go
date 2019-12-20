package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PrzemyslawMorski/backing-fulfillment/service"

	"github.com/hudl/fargo"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}

	c := fargo.NewConn("http://localhost:18080/eureka/v2")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Port is not an int")
		return
	}

	i := fargo.Instance{
		HostName:         "backing-fulfillment",
		Port:             portInt,
		App:              "backing-fulfillment",
		IPAddr:           "127.0.0.1",
		VipAddress:       "127.0.0.1",
		SecureVipAddress: "127.0.0.1",
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
		Status:           fargo.UP,
	}

	_ = c.RegisterInstance(&i)

	// Ordinarily we'd use a CF environment here, but we don't need it for the fake data we're returning.
	server := service.NewServer()
	server.Run(":" + port)
}
