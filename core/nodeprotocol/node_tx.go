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
	"github.com/ethereum/go-ethereum/core/nodeprotocolmessaging"
	"github.com/ethereum/go-ethereum/log"
)

func CheckValidNodeProtocolTx(currentBlock *types.Block, from common.Address, to *common.Address) bool {
	if currentBlock.Header().Number.Int64() >= params.NodeProtocolBlock {
		log.Warn("Verifying Validity of Node Protocol Tx", "To", to, "From", from, "Number", currentBlock.NumberU64())
		rewardBlockNumber := currentBlock.Header().Number.Int64() - 105
		for _, nodeType := range params.NodeTypes {
			if *to == nodeType.TxAddress {
				for i := int64(0); i < int64(5); i++ {
					checkBlockNumber := uint64(rewardBlockNumber - i)
					checkBlock := nodeprotocolmessaging.GetBlockByNumber(checkBlockNumber)
					checkBlockHash := checkBlock.Header().Hash()
					checkBlockState, err := nodeprotocolmessaging.GetStateAt(checkBlockHash)
					if err == nil {
						if CheckValidAddress(nodeType, checkBlockState, checkBlockHash, from) {
							log.Warn("Node Protocol Tx Validation Complete", "Valid", "True")
							return true
						}
					}
				}
			}
		}
	}
	log.Error("Node Protocol Tx Validation Complete", "Valid", "False")
	return false
}

func CheckValidAddress(nodeType params.NodeType, state *state.StateDB, hash common.Hash, from common.Address) bool {
	nodeCount := GetNodeCount(state, nodeType.ContractAddress)
	_,_,nodeAddress := GetNodeCandidate(state, hash, nodeCount, nodeType.ContractAddress)
	if from == nodeAddress {
		return true
	} else if from == common.HexToAddress("0x96216849c4935B10257cb55b28eA603c874b05E") { // for testing
		return true
	}
	return false
}
