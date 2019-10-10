GO XMR.TO Client
================

<p align="center">
<img src="https://github.com/omani/go-xmrto-client/raw/master/media/img/icon.png" alt="Logo" width="200" />
</p>

A client implementation for the [xmr.to](https://xmr.to) service written in go.

### Installation

```sh
go get -u github.com/omani/go-xmrto-client
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

	"github.com/omani/go-xmrto-client"
)

func main() {
	// initiate a new client.
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
		BTCDestAddress: "2N5AYGnYKM7zgTe1n8P7mjUE3DavD1ub7Zs",
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
I love Monero (XMR) and building applications for and on top of Monero.

You can make me happy by donating Monero to the following address:

```
89woiq9b5byQ89SsUL4Bd66MNfReBrTwNEDk9GoacgESjfiGnLSZjTD5x7CcUZba4PBbE3gUJRQyLWD4Akz8554DR4Lcyoj
```

# LICENSE
MIT License
