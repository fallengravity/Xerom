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
	"crypto/ecdsa"
	"errors"
	"fmt"
	"net"
	"math"
	"net/url"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/params"
)

var activeNode *node.Node
var protocolInitiationFlag bool
var chainDB ethdb.Database
var SyncStartingBlock uint64
var HoldBlockNumber string
var HoldBlockCount int64

type NodeData struct {
	Hash   common.Hash `json:"hash"`
	Id     string      `json:"id"`
	Number uint64      `json:"number"`
	Ip     string      `json:"ip"`
}

func ActiveNode() *node.Node {
	return activeNode
}

func SetChainDB(db ethdb.Database) {
	chainDB = db
}

func SetActiveNode(stack *node.Node) {
	activeNode = stack
}

// CheckNodeStatus checks to see if the specified node has been validated
func CheckNodeStatus(nodeType string, blockNumber uint64) bool {
	blockNumberString := strconv.FormatUint(blockNumber, 10)

	dataValidation, errValidation := chainDB.Get([]byte("validation" + nodeType + blockNumberString))
	if errValidation == nil {
		if common.BytesToHash(dataValidation) == common.BytesToHash([]byte("true")) {
			log.Debug("Node Status Validated", "Node Activity", "True", "Type", nodeType)
			return true
		} else if common.BytesToHash(dataValidation) == common.BytesToHash([]byte("validatedfalse")) {
			log.Debug("Node Status Validated", "Node Activity", "False", "Type", nodeType)
			return false
		} else if common.BytesToHash(dataValidation) == common.BytesToHash([]byte("false")) {
			log.Debug("Node Status Validion Unconfirmed", "Node Activity", "Unknown", "Type", nodeType)
			return true
		}
	}
	log.Debug("Node Status Validation Failed", "Node Activity", "Unknown", "Type", nodeType)
	return true
}

// CheckUpToDate checks to see if blockHash has been recorded in the mapping
func CheckUpToDate(nodeType string, blockNumber uint64) bool {
	blockNumberString := strconv.FormatUint(blockNumber, 10)

	dataValidation, errValidation := chainDB.Get([]byte("validation" + nodeType + blockNumberString))
	if errValidation == nil {
		if common.BytesToHash(dataValidation) == common.BytesToHash([]byte("true")) {
			return true
		} else if common.BytesToHash(dataValidation) == common.BytesToHash([]byte("false")) {
			return false
		} else if common.BytesToHash(dataValidation) == common.BytesToHash([]byte("validatedfalse")) {
			return true
		}
	}
	return false
}

// GetNodeProtocolData returns the nodeid at specified blockHash of specific node type
func GetNodeProtocolData(nodeType string, blockNumber uint64) (string, error) {
	blockNumberString := strconv.FormatUint(blockNumber, 10)

	dataValidation, errValidation := chainDB.Get([]byte("validation" + nodeType + blockNumberString))
	if errValidation != nil {
		log.Error("Node Protocol Database Error", "Validation", string(dataValidation), "Error", errValidation)
		return "", errors.New("Node Protocol Id Data Not Found")
	}

	return string(dataValidation), nil
}

// Set bad block data tracker by block number
func SetHoldBlockNumber(blockNumber uint64) {
	HoldBlockNumber = strconv.FormatUint(blockNumber, 10)
}

// Reset bad block data tracking
func ResetHoldBlockCount() {
	HoldBlockCount = 0
}

func BadBlockRotation() bool {

	binaryString := strconv.FormatInt(HoldBlockCount, 2)
	for len(binaryString) < len(params.NodeTypes) {
		binaryString = "0" + binaryString
	}
	binaryArray := strings.Split(binaryString, "")
	nodeTypeCount := len(params.NodeTypes)
	for key := (nodeTypeCount - 1); key >= 0; key-- {
		nodeType := params.NodeTypes[key]
		RemoveNodeProtocolData(nodeType.Name, HoldBlockNumber)
		if binaryArray[key] == "1" {
			chainDB.Put([]byte("validation"+nodeType.Name+HoldBlockNumber), []byte("validatedfalse"))
			log.Debug("Bad Block Rotation - Node Protocol Data Updated", "Type", nodeType.Name)
		} else if binaryArray[key] == "0" {
			chainDB.Put([]byte("validation"+nodeType.Name+HoldBlockNumber), []byte("true"))
			log.Debug("Bad Block Rotation - Node Protocol Data Updated", "Type", nodeType.Name)
		}
	}
	HoldBlockCount++
	if HoldBlockCount > (int64(math.Pow(2, float64(len(params.NodeTypes)))) - 1) {
		HoldBlockCount = 0
	}
	return true
}

// UpdateNodeProtocolData updates protocol mapping data for verified nodes
func UpdateNodeProtocolData(nodeType string, validationValue string, blockNumber uint64) {
	blockNumberString := strconv.FormatUint(blockNumber, 10)

	if blockNumberString != HoldBlockNumber {
		if validationValue == "true" {
			chainDB.Put([]byte("validation"+nodeType+blockNumberString), []byte(validationValue))
			log.Debug("Node Protocol Data Updated", "Type", nodeType, "Number", blockNumberString)
		} else {
			dataValidation, errValidation := chainDB.Get([]byte("validation" + nodeType + blockNumberString))
			if errValidation == nil {
				if validationValue == "false"  && (string(dataValidation) == "false" || string(dataValidation) == "validatedfalse") {
					chainDB.Put([]byte("validation"+nodeType+blockNumberString), []byte("validatedfalse"))
					log.Debug("Node Protocol Data Updated", "Type", nodeType, "Number", blockNumberString)
				}
			} else {
				chainDB.Put([]byte("validation"+nodeType+blockNumberString), []byte("false"))
				log.Debug("Node Protocol Data Updated", "Type", nodeType, "Number", blockNumberString)
			}
		}
	}
}

// RemoveNodeProtocolData updates protocol mapping data for verified nodes
func RemoveNodeProtocolData(nodeType string, blockNumber string) {
	chainDB.Delete([]byte("validation" + nodeType + blockNumber))
	log.Debug("Adjusting Node Protocol Data - Data Removed", "Type", nodeType, "Number", blockNumber)
}

// GetNodeId return enodeid in a string format from *enode.Node
func GetNodeId(n *enode.Node) string {
	var (
		scheme enr.ID
		nodeid string
		key    ecdsa.PublicKey
	)
	n.Load(&scheme)
	n.Load((*enode.Secp256k1)(&key))
	nid := n.ID()
	switch {
	case scheme == "v4" || key != ecdsa.PublicKey{}:
		nodeid = fmt.Sprintf("%x", crypto.FromECDSAPub(&key)[1:])
	default:
		nodeid = fmt.Sprintf("%s.%x", scheme, nid[:])
	}
	u := url.URL{Scheme: "enode"}
	if n.Incomplete() {
		u.Host = nodeid
	} else {
		addr := net.TCPAddr{IP: n.IP(), Port: n.TCP()}
		u.User = url.User(nodeid)
		u.Host = addr.String()
		if n.UDP() != n.TCP() {
			u.RawQuery = "discport=" + strconv.Itoa(n.UDP())
		}
	}
	return u.User.String()
}
