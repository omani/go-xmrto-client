GO XMR.TO Client
================

[![GoDoc](https://godoc.org/github.com/monero-ecosystem/go-xmrto-client?status.svg)](https://godoc.org/github.com/monero-ecosystem/go-xmrto-client)


<p align="center">
<img src="https://github.com/monero-ecosystem/go-xmrto-client/raw/master/media/img/icon.png" alt="Logo" width="300" />
</p>

A client implementation for the [xmr.to](https://xmr.to) service written in go.

# NOTICE
The xmr.to service shut down long time ago. This repo has been archived.

### Installation

```sh
go get -u github.com/monero-ecosystem/go-xmrto-client
```

#### Example code:

```Go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/monero-ecosystem/go-xmrto-client"
)

func main() {
	// initiate a new client for the testnet
	client := xmrto.New(&xmrto.Config{Testnet: true})

	// here we check if we run the programm with a check parameter.
	// in case we already got a secret-key from xmr.to we can track the order.
	if len(os.Args) > 2 {
		if os.Args[1] == "check" {
			checkorder, err := client.GetOrderStatus(&xmrto.RequestGetOrderStatus{UUID: os.Args[2]})
			checkerr(err)

			log.Println(string(prettyPrintJSON(checkorder)))
			fmt.Printf("Your order has state: %s\n", checkorder.State)

			return
		}
	}
	// otherwise

	// get order parameters:
	// - how much bitcoin can we send?
	// - what is the upper limit? lower limit?
	// - the current exchange rate (price) of the btc amount we want to send
	// - etc.
	getorder, err := client.GetOrderParameters()
	checkerr(err)

	// pretty print.
	log.Println(string(prettyPrintJSON(getorder)))

	// let's create an order with 0.001 btc.
	createorder, err := client.CreateOrder(&xmrto.RequestCreateOrder{
		BTCAmount:      .001,
		BTCDestAddress: "2N5AYGnYKM7zgTe1n8P7mjUE3DavD1ub7Zs", // this is the testnet btc address of xmr.to itself.
	})
	checkerr(err)

	log.Println(string(prettyPrintJSON(createorder)))

	// we got a secret-key from xmr.to (for later).
	fmt.Printf("UUID (secret key) is: %s\n", createorder.UUID)

	// important: give it time to let "CreateOrder" to settle down
	// before we query our order.
	time.Sleep(time.Second * 1)

	// now check the order state with the secret-key
	// we received from xmr.to for this particular order.
	orderstatus, err := client.GetOrderStatus(&xmrto.RequestGetOrderStatus{UUID: createorder.UUID})
	checkerr(err)

	log.Println(string(prettyPrintJSON(orderstatus)))

	// print a nice message to the user
	// how much xmr to deposit to which xmr address.
	fmt.Printf("Please deposit %f to Monero address: %s\n", orderstatus.XMRAmountTotal, orderstatus.XMRReceivingSubAddress)

	return
}

/* helper functions */
func prettyPrintJSON(data interface{}) (pretty []byte) {
	pretty, _ = json.MarshalIndent(data, "", "    ")
	return
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
```

# Run the code
```Go
go run main.go
```
to create an order, or:
```Go
go run main.go check <your secret-key>
```
to check an order about its state

# Contribution
* You can fork this, extend it and contribute back.
* You can contribute with pull requests.

# Donations
You can make me happy by donating Bitcoin to the following address:
```
bc1qgezvfp4s0xme8pdv6aaqu9ayfgnv4mejdlv3tx
```

# LICENSE
MIT License
