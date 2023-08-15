package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/make-software/casper-go-sdk/casper"
)

func main() {
	handler := casper.NewRPCHandler("http://94.130.10.55:7777/rpc", http.DefaultClient)
	client := casper.NewRPCClient(handler)
	deployHash := "62972eddc6fdc03b7ec53e52f7da7e24f01add9a74d68e3e21d924051c43f126"
	deploy, err := client.GetDeploy(context.Background(), deployHash)
	if err != nil {
		return
	}

	resultStr, _ := json.Marshal(deploy)
	log.Print("Parsed rpc response: \n" + string(resultStr))
}
