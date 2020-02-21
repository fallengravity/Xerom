// Copyright 2017 The go-ethereum Authors
// Copyright 2020 by The etho.black Development Team
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// +build none

/*

   The mkdnpupdate tool creates the RPL encoded dnp sync data for future use with dnp node syncing

       go run mkdnpalloc.go dnp.json

*/
package main

import (
	"encoding/json"
	"fmt"
        "io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

var dnpData map[common.Hash]int64

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: mkalloc genesis.json")
		os.Exit(1)
	}

        dnpData := make(map[common.Hash]int64)
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(file).Decode(&dnpData); err != nil {
		panic(err)
	}

        fmt.Println(dnpData)

        dnpData[common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000005511")] = 0100

        updatedDnpData, _ := json.Marshal(dnpData)
        err = ioutil.WriteFile("updatedDnpData.json", updatedDnpData, 0644)
        if err != nil {
            panic(err)
        }
}
