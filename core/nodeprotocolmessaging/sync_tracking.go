// Copyright 2015 The go-ethereum Authors
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

package nodeprotocolmessaging

import (

//        "github.com/ethereum/go-ethereum/log"
)
/*
var nodeProtocolHead uint64

var nodeProtocolHeadChannel chan uint64

func GetNodeProtocolHead() uint64 {
        return nodeProtocolHead
}

func GetNodeProtocolHeadChannel() chan uint64 {
        return nodeProtocolHeadChannel
}

func Set(newHead uint64) {
	nodeProtocolHead = newHead
	ch := nodeProtocolHeadChannel
	ch <- nodeProtocolHead
}

func GetChan() chan uint64 {
	listener := make(chan uint64, 5)
	nodeProtocolHeadChannel = listener
	return listener
}

func Background(ch chan uint64, done chan uint64) {
        for v := range ch {
                log.Trace("Node Protocol Data Head State Sync Update", "Updated Head Block Number", v)
	}
	done <- 0
}

*/
