package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/make-software/casper-go-sdk/casper"
	"github.com/make-software/casper-go-sdk/types"
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

	accountPublicKey, err := casper.NewPublicKey("0152836c51eac04205bb7febe9d92da50758178b0bf388bd03e1da13147b99e2c5")
	if err != nil {
		fmt.Println("error for accountPublicKey: ", err)
		return
	}
	targetPublicKey, err := casper.NewPublicKey("0193b3800386aefe11648150f6779158f2c7e1233c8e9b423338eb71b93ae6c5a9")
	if err != nil {
		fmt.Println("error for targetPublicKey: ", err)
		return
	}
	amount := big.NewInt(100000000)
	sessionArgs := types.Args{}
	sessionArgs.AddArgument("amount", *clvalue.NewCLUInt512(big.NewInt(2500000000)))
	sessionArgs.AddArgument("target", clvalue.NewCLByteArray(targetPublicKey.AccountHash().Bytes()))
	sessionArgs.AddArgument("id", clvalue.NewCLOption(*clvalue.NewCLUInt64(123)))

	session := types.ExecutableDeployItem{
		Transfer: &types.TransferDeployItem{
			Args: sessionArgs,
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
