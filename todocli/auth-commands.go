package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"todos/api"
)

var (
	authFlags = flag.NewFlagSet("auth", flag.ExitOnError)
	email     = authFlags.String("email", "", "email for authentication")
	password  = authFlags.String("password", "", "password for authentication")
)

func authCommands(args []string) {
	if len(args) < 1 {
		log.Println("An auth subcommand is required")
		authFlags.Usage()
		return
	}

	authFlags.Parse(args[1:])
	switch args[0] {
	case "signup":
		checkAuthFlags()
		client := api.NewClient(*host, "")
		token, err := client.Signup(*email, *password)
		if err != nil {
			log.Fatalf("Error during %s: %v", args[0], err)
		}
		storeToken(token)
		fmt.Println("Successfully logged in")

	case "signin":
		checkAuthFlags()
		client := api.NewClient(*host, "")
		token, err := client.Signin(*email, *password)
		if err != nil {
			log.Fatalf("Error during %s: %v", args[0], err)
		}
		storeToken(token)
		fmt.Println("Successfully logged in")

	case "identity":
		token, err := readToken()
		if err != nil {
			log.Fatal("Error reading token", err)
		}
		client := api.NewClient(*host, token)
		user, err := client.Identity()
		if err != nil {
			log.Fatal("Error fetching identity", err)
		}
		b, _ := json.MarshalIndent(&user, "", "  ")
		fmt.Print(string(b))

	default:
		log.Println("Unknown auth command:", args[0])
		authFlags.Usage()
	}
}

func checkAuthFlags() {
	if *email == "" || *password == "" {
		log.Fatal("Email and password are required")
	}
}
