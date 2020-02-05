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

   The mkdnpalloc tool creates the RPL encoded dnp sync data for future use with dnp node syncing

       go run mkdnpalloc.go dnp.json

*/
package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type allocItem struct{ Hash, Data *big.Int }

type allocList []allocItem

var dnpData map[common.Hash]int64

func (a allocList) Len() int           { return len(a) }
func (a allocList) Less(i, j int) bool { return a[i].Hash.Cmp(a[j].Hash) < 0 }
func (a allocList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func makelist(d map[common.Hash]int64) allocList {
	a := make(allocList, 0, len(d))
	for key, data  := range d {
                bigHash := new(big.Int).SetBytes(key.Bytes())
                bigData := big.NewInt(data)
		a = append(a, allocItem{bigHash, bigData})
	}
	sort.Sort(a)
	return a
}

func makealloc(d map[common.Hash]int64) string {
	a := makelist(d)
	data, err := rlp.EncodeToBytes(a)
	if err != nil {
		panic(err)
	}
	return strconv.QuoteToASCII(string(data))
}

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
	fmt.Println("const allocData =", makealloc(dnpData))
}
