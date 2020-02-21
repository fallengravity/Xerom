// Copyright 2015 The go-ethereum Authors
// Copyright 2019 The Etho.Black Development Team
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

package dnpbridge

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
)

// CheckNodeStatus checks to see if the specified node has been validated
func CheckNodeStatus(nodeType string, nodeId common.Hash, nodeIp common.Hash, blockHash common.Hash, blockNumber uint64) bool {
	blockNumberString := strconv.FormatUint(blockNumber, 10)

	dataId, errId := chainDB.Get([]byte("id" + nodeType + blockNumberString))
	dataIp, errIp := chainDB.Get([]byte("ip" + nodeType + blockNumberString))
	dataHash, errHash := chainDB.Get([]byte("hash" + nodeType + blockNumberString))
	if errId == nil && errIp == nil && errHash == nil {
		if nodeId == common.BytesToHash(dataId) && nodeIp == common.BytesToHash(dataIp) {
			log.Debug("Node ID Found in Node Protocol Data", "Validated", "True", "Type", nodeType, "ID", nodeId, "IP", nodeIp)
			return true
		}
	}
	log.Debug("Node ID Not Found in Node Protocol Data", "Validated", "False", "Type", nodeType, "ID Needing Verification", nodeId, "Hash", blockHash, "IP", nodeIp, "Saved ID", common.BytesToHash([]byte(dataId)), "Saved Hash", dataHash, "Saved IP", common.BytesToHash([]byte(dataIp)))
	return false
}

// CheckUpToDate checks to see if blockHash has been recorded in the mapping
func CheckUpToDate(nodeType string, blockHash common.Hash, blockNumber uint64) bool {
	blockNumberString := strconv.FormatUint(blockNumber, 10)

	dataHash, errHash := chainDB.Get([]byte("hash" + nodeType + blockNumberString))
	if errHash == nil && blockHash == common.BytesToHash(dataHash) {
		return true
	}
	return false
}

// GetNodeProtocolData returns the nodeid at specified blockHash of specific node type
func GetNodeProtocolData(nodeType string, blockHash common.Hash, blockNumber uint64) (string, string, error) {
	blockNumberString := strconv.FormatUint(blockNumber, 10)

	dataId, err := chainDB.Get([]byte("id" + nodeType + blockNumberString))
	if err != nil {
		log.Error("Node Protocol Database Error", "ID", string(dataId), "Error", err)
		return "", "", errors.New("Node Protocol Id Data Not Found")
	}
	dataIp, err := chainDB.Get([]byte("ip" + nodeType + blockNumberString))
	if err != nil {
		log.Error("Node Protocol Database Error", "IP", string(dataIp), "Error", err)
		return "", "", errors.New("Node Protocol Ip Data Not Found")
	}

	return string(dataId), string(dataIp), nil
}

// UpdateNodeProtocolData updates protocol mapping data for verified nodes
func UpdateNodeProtocolData(nodeType string, nodeId string, nodeIp string, peerId string, peerCount int, blockHash common.Hash, blockNumber uint64, syncing bool) {
	blockNumberString := strconv.FormatUint(blockNumber, 10)

	if blockNumberString != HoldBlockNumber {
		chainDB.Put([]byte("id"+nodeType+blockNumberString), []byte(nodeId))
		chainDB.Put([]byte("ip"+nodeType+blockNumberString), []byte(nodeIp))
		chainDB.Put([]byte("hash"+nodeType+blockNumberString), []byte(blockHash.String()))

		log.Debug("Node Protocol Data Updated", "Type", nodeType, "Hash", blockHash)
	}
}

// RemoveNodeProtocolData updates protocol mapping data for verified nodes
func RemoveNodeProtocolData(nodeType string, nodeId string, nodeIp string, blockNumberString string) {
	chainDB.Delete([]byte("id" + nodeType + blockNumberString))
	chainDB.Delete([]byte("ip" + nodeType + blockNumberString))
	chainDB.Delete([]byte("hash" + nodeType + blockNumberString))
	log.Debug("Adjusting Node Protocol Data - Data Removed", "Type", nodeType, "Number", blockNumberString)
}
