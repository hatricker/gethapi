package main

import (
	"flag"
	"fmt"

	"github.com/hatricker/gethapi/client"
)

func main() {
	var (
		url = flag.String("url", "", "json-rpc endpoint")
	)

	flag.Parse()
	if *url == "" {
		fmt.Printf("JSON-RPC endpoint address not provided")
		return
	}

	if len(flag.Args()) < 1 {
		fmt.Printf("Usage: gethapi --url <rpc endpoint> rpc_method [args...]")
		return
	}

	agent, err := client.New(*url)
	if err != nil {
		fmt.Printf("cannot set up RPC client, %v", err)
	}
	defer agent.Close()

	var result interface{}
	args := flag.Args()
	if err = agent.CallMethod(&result, args[0], args[1:]); err != nil {
		fmt.Printf("RPC call failed, %v", err)
	}
	fmt.Println(result)
}
