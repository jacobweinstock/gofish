//
// SPDX-License-Identifier: BSD-3-Clause
//
package main

import (
	"context"
	"time"

	"github.com/stmcginnis/gofish"
)

func changeLogin() {
	// Create a new instance of gofish client, ignoring self-signed certs
	username := "my-username"
	config := gofish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: username,
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

	// Query the AccountService using the session token
	accountService, err := service.AccountService(ctx)
	if err != nil {
		panic(err)
	}
	// Get list of accounts
	accounts, err := accountService.Accounts(ctx)
	if err != nil {
		panic(err)
	}
	// Iterate over accounts to find the current user
	for _, account := range accounts {
		if account.UserName == username {
			account.UserName = "new-username"
			// New password must follow the rules set in AccountService :
			// MinPasswordLength and MaxPasswordLength
			account.Password = "new-password"
			err := account.Update(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}
