rddnet
======

[![Build Status](https://travis-ci.org/reddcoin-project/rddnet.png?branch=master)]
(https://travis-ci.org/reddcoin-project/rddnet) [![Coverage Status]
(https://coveralls.io/repos/reddcoin-project/rddnet/badge.png?branch=master)]
(https://coveralls.io/r/reddcoin-project/rddnet?branch=master)

Package rddnet defines the network parameters for the three standard Reddcoin
networks and provides the ability for callers to define their own custom 
Reddcoin networks.

This package is one of the core packages from rddd, an alternative full-node
implementation of Reddcoin which is under active development by Conformal.
Although it was primarily written for rddd, this package has intentionally been
designed so it can be used as a standalone package for any projects needing to
use parameters for the standard Reddcoin networks or for projects needing to
define their own network.

## Sample Use

```Go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/reddcoin-project/rddnet"
	"github.com/reddcoin-project/rddutil"
)

var testnet = flag.Bool("testnet", false, "operate on the testnet Reddcoin network")

// By default (without -testnet), use mainnet.
var netParams = &rddnet.MainNetParams

func main() {
	flag.Parse()

	// Modify active network parameters if operating on testnet.
	if *testnet {
		netParams = &rddnet.TestNet3Params
	}

	// later...

	// Create and print new payment address, specific to the active network.
	pubKeyHash := make([]byte, 20)
	addr, err := rddutil.NewAddressPubKeyHash(pubKeyHash, netParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
```

## Documentation

[![GoDoc](https://godoc.org/github.com/reddcoin-project/rddnet?status.png)]
(http://godoc.org/github.com/reddcoin-project/rddnet)

Full `go doc` style documentation for the project can be viewed online without
installing this package by using the GoDoc site
[here](http://godoc.org/github.com/reddcoin-project/rddnet).

You can also view the documentation locally once the package is installed with
the `godoc` tool by running `godoc -http=":6060"` and pointing your browser to
http://localhost:6060/pkg/github.com/reddcoin-project/rddnet

## Installation

```bash
$ go get github.com/reddcoin-project/rddnet
```

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from Conformal.  To verify the
signature perform the following:

- Download the public key from the Conformal website at
  https://opensource.conformal.com/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

Package rddnet is licensed under the liberal ISC License.
