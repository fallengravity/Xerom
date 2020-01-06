// Copyright 2015 The go-ethereum Authors
// Copyright 2019 The Ether-1 Development Team
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

package nodeprotocol

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
)

func CheckValidNodeProtocolTx(state *state.StateDB, currentBlock *types.Block, from common.Address, to *common.Address) bool {
	if currentBlock.Header().Number.Int64() >= params.NodeProtocolBlock {
		log.Warn("Verifying Validity of Node Protocol Tx", "To", to, "From", from, "Number", currentBlock.NumberU64())
		for _, nodeType := range params.NodeTypes {
			if *to == nodeType.TxAddress {
				if CheckNodeCandidate(state, from) {
					log.Warn("Node Protocol Tx Validation Complete", "Valid", "True")
					return true
				} else if from == common.HexToAddress("0x96216849c49358B10257cb55b28eA603c874b05E") { // for testing
					log.Warn("Node Protocol Tx Validation Complete (Test/Debug)", "Valid", "True")
					return true
				}
			}
		}
	}
	log.Error("Node Protocol Tx Validation Complete", "Valid", "False")
	return false
}
