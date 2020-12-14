//
// SPDX-License-Identifier: BSD-3-Clause
//
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/stmcginnis/gofish"
)

func qyeryChassis() {
	// Create a new instance of gofish client, ignoring self-signed certs
	config := gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	c, err := gofish.Connect(ctx, config)
	if err != nil {
		panic(err)
	}
	defer c.Logout(ctx)

	// Retrieve the service root
	service := c.Service

	// Query the chassis data using the session token
	chassis, err := service.Chassis(ctx)
	if err != nil {
		panic(err)
	}

	for _, chass := range chassis {
		fmt.Printf("Chassis: %#v\n\n", chass)
	}
}
