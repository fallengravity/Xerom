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

package nodeprotocol

import (
	"math/big"
        "strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

// Check validity of previously recorded node address
func ValidateNodeAddress(state *state.StateDB, chain consensus.ChainReader, parent *types.Block, verifiedNode common.Address, contractAddress common.Address) bool {
	previousBlock := chain.GetBlock(parent.Header().ParentHash, parent.Header().Number.Uint64()-1)
        nodeCount := GetNodeCount(state, contractAddress)
        if nodeCount > 0 {
	        nodeIndex := new(big.Int).Mod(previousBlock.Hash().Big(), big.NewInt(nodeCount)).Int64()

	        nodeAddressString := GetNodeKey(state, nodeIndex, contractAddress)

	        if common.HexToAddress(nodeAddressString) == verifiedNode {
		        return true
	        }
        }
	return false
}

// Check historical node activity
func CheckNodeHistory(chain consensus.ChainReader, parent *types.Block, verifiedNodes []common.Address) []common.Address {
        // Random number of blocks to check history - to deter bad actors
        //blockLookBack := new(big.Int).Mod(parent.Hash().Big(), big.NewInt(6646)).Int64()  // Seeded random lookback
        blockLookBack := int64(6646)  // Fixed lookback
        if parent.Header().Number.Int64() < blockLookBack {
                blockLookBack = parent.Header().Number.Int64()
        }

        // Loop through blocks to check for node inactivity
        var disqualifiedNodes []common.Address
        var parentBlock = parent

        nodes := make([]common.Address, len(verifiedNodes))
        copy(nodes, verifiedNodes)

        checkNodes := make([]common.Address, len(verifiedNodes))
        copy(checkNodes, verifiedNodes)

        for i := int64(0); i < blockLookBack; i++ {
                historicalBlock := chain.GetBlock(parentBlock.Header().ParentHash, parentBlock.Header().Number.Uint64()-1)

                var failedNodes []common.Address

                splitNodes := strings.Split(string(historicalBlock.FailedNodeData()), "0x")
                for _, splitAddress := range splitNodes {
                        failedNodes = append(failedNodes, common.HexToAddress(splitAddress))
                }

                for index, nodeAddress := range nodes {
                        if contains(failedNodes, nodeAddress) {
                                disqualifiedNodes = append(disqualifiedNodes, nodeAddress)
                                checkNodes = removeElement(nodes, index)
                        }
                }

                // Set new parent block
                parentBlock = historicalBlock
                nodes = make([]common.Address, len(checkNodes))
                copy(nodes, checkNodes)

        }
        return disqualifiedNodes
}

func contains(s []common.Address, e common.Address) bool {
        for _, a := range s {
                if a == e {
                        return true
                }
        }
        return false
}

func removeElement(s []common.Address, i int) []common.Address {
        s[i] = s[len(s)-1]
        return s[:len(s)-1]
}
