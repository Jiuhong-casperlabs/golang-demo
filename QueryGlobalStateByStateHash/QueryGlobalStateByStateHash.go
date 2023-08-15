package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/make-software/casper-go-sdk/casper"
)

func main() {
	handler := casper.NewRPCHandler("http://94.130.10.55:7777/rpc", http.DefaultClient)
	client := casper.NewRPCClient(handler)
	key := "uref-d98fedf17d6de2ac8d09952e14373fb9a0d25a7dbb1c63e34111132dd2dc3734-007"
	res, err := client.QueryGlobalStateByStateHash(context.Background(), nil, key, nil)

	if err != nil {
		fmt.Println("error is ", err)
		return
	}
	fmt.Println(res.StoredValue.CLValue.Value())
}
