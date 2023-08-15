package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/make-software/casper-go-sdk/casper"
	"github.com/make-software/casper-go-sdk/types/clvalue"
	"github.com/make-software/casper-go-sdk/types/keypair"
)

func main() {
	privateKeyPath := "/home/jh/keys/test1/secret_key.pem"
	privateKey, err := keypair.NewPrivateKeyED25518(privateKeyPath)
	if err != nil {
		fmt.Println("error for privateKey: ", err)
		return
	}

	modulePath := "/home/jh/casper-node/target/wasm32-unknown-unknown/release/delegate.wasm"
	accountPublicKey, err := casper.NewPublicKey("0152836c51eac04205bb7febe9d92da50758178b0bf388bd03e1da13147b99e2c5")
	if err != nil {
		fmt.Println("error for accountPublicKey: ", err)
		return
	}
	amount := big.NewInt(100000000)
	module, _ := os.ReadFile(modulePath)
	session := casper.ExecutableDeployItem{
		ModuleBytes: &casper.ModuleBytes{
			ModuleBytes: hex.EncodeToString([]byte(module)),
			Args: (&casper.Args{}).
				AddArgument("target", clvalue.NewCLByteArray(accountPublicKey.AccountHash().Bytes())).
				AddArgument("amount", *clvalue.NewCLUInt512(amount)),
		},
	}

	payment := casper.StandardPayment(amount)

	deployHeader := casper.DefaultHeader()
	deployHeader.Account = accountPublicKey
	deployHeader.ChainName = "casper-test"

	newDeploy, err := casper.MakeDeploy(deployHeader, payment, session)
	if err != nil {
		fmt.Println("error for MakeDeploy: ", err)
		return
	}

	err = newDeploy.SignDeploy(privateKey)
	if err != nil {
		fmt.Println("error for SignDeploy: ", err)
		return
	}

	// print out deploy json
	resultStr, _ := json.Marshal(newDeploy)
	log.Print("Parsed rpc response: \n" + string(resultStr))
	//

	handler := casper.NewRPCHandler("http://142.132.208.215:7777/rpc", http.DefaultClient)
	client := casper.NewRPCClient(handler)

	result, err := client.PutDeploy(context.Background(), *newDeploy)

	if err != nil {
		fmt.Println("error for PutDeploy: ", err)
		return
	}

	log.Println(result.DeployHash)
}
