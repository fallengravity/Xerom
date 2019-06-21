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

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

// Get next node reward candidate based on current state and nodeCount
func GetNodeCandidate(state *state.StateDB, blockHash common.Hash, nodeCount int64, contractAddress common.Address) (string, string, common.Address) {
         nodeIndex := new(big.Int).Mod(blockHash.Big(), big.NewInt(nodeCount)).Int64()
         return getNodeData(state, getNodeKey(state, nodeIndex, contractAddress), contractAddress)
}

func GetNodeCount(state *state.StateDB, contractAddress common.Address) int64 {
	// Get storage state form db using index
	nodeCount := state.GetState(contractAddress, common.HexToHash("2")).Big().Int64()

	return nodeCount
}

func getNodeKey(state *state.StateDB, nodeIndex int64, contractAddress common.Address) string {
	solcIndex := int64(1)

	hash := sha3.NewKeccak256()
	var buf []byte

	// Index is the contract variable index based on solc storage state standards
	index := abi.U256(big.NewInt(solcIndex))

	// Key is the mapping key to lookup
	key := abi.U256(big.NewInt(int64(nodeIndex)))

	// Prepare the Keccak256 seed
	location := append(key, index...)

	hash.Write(location)
	buf = hash.Sum(nil)
	storageLocation := common.BytesToHash(buf)

	// Get storage state form db using the hashed data
	response := state.GetState(contractAddress, storageLocation)

	// Assemble the strings
	nodeAddressString := response

	return nodeAddressString.String()
}

func getNodeData(state *state.StateDB, nodeAddress string, contractAddress common.Address) (string, string, common.Address) {
	solcIndex := int64(0)

	hash := sha3.NewKeccak256()
	var buf []byte

	// Left-fill with zeros to meet abi packing standards
	prepend := make([]byte, 12)

	// Index is the contract variable index based on solc storage state standards
	index := abi.U256(big.NewInt(solcIndex))

	// Key is the mapping key to lookup
	key := common.HexToAddress(nodeAddress).Bytes()[:]

	// Prepare the Keccak256 seed
	location := append(prepend, key...)
	location = append(location, index...)

	hash.Write(location)
	buf = hash.Sum(nil)
	storageLocation := common.BytesToHash(buf)

	nodeIdLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(2))).Bytes()

	// Get offsets for long enodeid string
	hash = sha3.NewKeccak256()
	var buf1 []byte
	hash.Write(nodeIdLocation)
	buf1 = hash.Sum(nil)
	finalNodeIdLocation1 := common.BytesToHash(buf1)
	finalNodeIdLocation2 := common.BigToHash(new(big.Int).Add(finalNodeIdLocation1.Big(), big.NewInt(1)))
	finalNodeIdLocation3 := common.BigToHash(new(big.Int).Add(finalNodeIdLocation1.Big(), big.NewInt(2)))
	finalNodeIdLocation4 := common.BigToHash(new(big.Int).Add(finalNodeIdLocation1.Big(), big.NewInt(3)))

	nodeAddressLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(1)))
        nodeIpLocation := common.BigToHash(new(big.Int).Add(storageLocation.Big(), big.NewInt(3)))

	// Get storage state form db using the hashed data
	responseNodeId1 := state.GetState(contractAddress, finalNodeIdLocation1)
	responseNodeId2 := state.GetState(contractAddress, finalNodeIdLocation2)
	responseNodeId3 := state.GetState(contractAddress, finalNodeIdLocation3)
	responseNodeId4 := state.GetState(contractAddress, finalNodeIdLocation4)
	responseNodeAddress := state.GetState(contractAddress, nodeAddressLocation)
        responseNodeIp := state.GetState(contractAddress, nodeIpLocation)

	// Assemble the strings
        contractNodeId := stripCtlAndExtFromBytes(string(responseNodeId1.Bytes())) + stripCtlAndExtFromBytes(string(responseNodeId2.Bytes())) + stripCtlAndExtFromBytes(string(responseNodeId3.Bytes())) + stripCtlAndExtFromBytes(string(responseNodeId4.Bytes()))
	contractNodeAddress := common.BytesToAddress(responseNodeAddress.Bytes())
        contractNodeIp := stripCtlAndExtFromBytes(string(responseNodeIp.Bytes()))

	return contractNodeId, contractNodeIp, contractNodeAddress
}

func UpdateNodeCount(state *state.StateDB, currentNodeCount int64, countAddresses []common.Address) uint64 {

        var nodeCount = common.BytesToAddress(state.GetCode(countAddresses[len(countAddresses)-1])).Big().Uint64()

        // Rotate addresses for caching behavior
        for i := len(countAddresses)-1; i > 0; i-- {
                state.SetCode(countAddresses[i], state.GetCode(countAddresses[i-1]))
        }
        currentNodeCountBytes := big.NewInt(currentNodeCount).Bytes()
        state.SetCode(countAddresses[0], currentNodeCountBytes)

        log.Trace("Updating Node Counts", "Count", nodeCount)

        return nodeCount
}

func UpdateNodeCandidate(state *state.StateDB, currentNodeId string, currentNodeIp string, currentNodeAddress common.Address, nodeIds []common.Address, nodeIps []common.Address, nodeAddresses []common.Address) (common.Hash, common.Hash, common.Address) {

        var nodeId = common.BytesToHash(state.GetCode(nodeIds[len(nodeIds)-1]))
        var nodeIp = common.BytesToHash(state.GetCode(nodeIps[len(nodeIps)-1]))
        var rewardAddress = common.BytesToAddress(state.GetCode(nodeAddresses[len(nodeAddresses)-1]))

        // Rotate addresses for caching behavior
        for i := len(nodeIds)-1; i > 0; i-- {
                state.SetCode(nodeIds[i], state.GetCode(nodeIds[i-1]))
                state.SetCode(nodeIps[i], state.GetCode(nodeIps[i-1]))
                state.SetCode(nodeAddresses[i], state.GetCode(nodeAddresses[i-1]))
        }
        state.SetCode(nodeIds[0], []byte(currentNodeId))
        state.SetCode(nodeIps[0], []byte(currentNodeIp))
        state.SetCode(nodeAddresses[0], currentNodeAddress.Bytes())

        log.Trace("Updating Node Reward Candidates", "ID", nodeId, "IP", nodeIp, "Address", rewardAddress)

        return nodeId, nodeIp, rewardAddress
}
