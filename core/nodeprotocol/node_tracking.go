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
 	"fmt"
        "strconv"
        "crypto/ecdsa"
        "net/url"
        "net"
        "errors"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
        "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

var activeNode *node.Node
var protocolInitiationFlag bool
var chainDB ethdb.Database
var SyncStartingBlock uint64

type NodeData struct {
        Hash   common.Hash  `json:"hash"`
        Id     string       `json:"id"`
        Number uint64       `json:"number"`
        Ip     string       `json:"ip"`
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

// CheckNodeStatus checks to see if specified node has been validated
func CheckNodeStatus(nodeType string, nodeId common.Hash, nodeIp common.Hash, blockHash common.Hash, blockNumber uint64) bool {
        blockNumberString := strconv.FormatUint(blockNumber, 10)

        dataId, errId := chainDB.Get([]byte("id" + nodeType + blockNumberString))
        dataIp, errIp := chainDB.Get([]byte("ip" + nodeType + blockNumberString))
        dataHash, errHash := chainDB.Get([]byte("hash" + nodeType + blockNumberString))
        if errId == nil && errIp == nil && errHash == nil{
                if nodeId == common.BytesToHash(dataId) && nodeIp == common.BytesToHash(dataIp) {
                        log.Debug("Node ID Found in Node Protocol Data", "Validated", "True", "Type", nodeType, "ID", nodeId, "IP", nodeIp)
                        return true
                }
        }
        log.Debug("Node ID Not Found in Node Protocol Data", "Validated", "False", "Type", nodeType, "ID Needing Verification", nodeId, "Hash", blockHash, "IP", nodeIp, "Saved ID", common.BytesToHash([]byte(dataId)), "Saved Hash", dataHash, "Saved IP", common.BytesToHash([]byte(dataIp)))
        return false
}

// CheckUpToDate checks to see if blockHash has been recorded in mapping
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

        chainDB.Put([]byte("id" + nodeType + blockNumberString), []byte(nodeId))
        chainDB.Put([]byte("ip" + nodeType + blockNumberString), []byte(nodeIp))
        chainDB.Put([]byte("hash" + nodeType + blockNumberString), []byte(blockHash.String()))

        log.Debug("Node Protocol Data Updated", "Type", nodeType, "ID", nodeId, "IP", nodeIp, "Hash", blockHash)
}

// SyncNodeProtocolDataGroup adds a slice of NodeData to state is consenus is reached
func SyncNodeProtocolDataGroup(nodeType string, nodeData map[uint64]NodeData, peerId string, peerCount int) {

        largestBlockNumber := uint64(0)
        smallestBlockNumber := uint64(99999999)
        for blockNumber, _ := range nodeData {
                if blockNumber > largestBlockNumber {
                        largestBlockNumber = blockNumber
                } else if blockNumber < smallestBlockNumber {
                        smallestBlockNumber = blockNumber
                }
        }
        if smallestBlockNumber == uint64(99999999) {
                smallestBlockNumber = uint64(0)
        }

        for number, data := range nodeData {
                blockNumberString := strconv.FormatUint(number, 10)
                chainDB.Put([]byte("id" + nodeType + blockNumberString), []byte(data.Id))
                chainDB.Put([]byte("ip" + nodeType + blockNumberString), []byte(data.Ip))
                chainDB.Put([]byte("hash" + nodeType + blockNumberString), []byte(data.Hash.String()))
        }


        if len(nodeData) > 0 {
                log.Info("Imported Node-Protocol Data", "entries", len(nodeData), "blocks", strconv.FormatUint(smallestBlockNumber, 10) + "->" + strconv.FormatUint(largestBlockNumber, 10))
        }
}

//func GetNodeProtocolDataGroup(nodeType string, startBlock uint64, endBlock uint64) (map[uint64]NodeData, error) {
func GetNodeProtocolDataGroup(nodeType string, startBlock uint64, endBlock uint64) ([]string, []string, []string, []string, error) {
        var hashes []string
        var nodes  []string
        var ips  []string
        var numbers []string

        for i := startBlock; i <= endBlock; i++ {
                blockNumberString := strconv.FormatUint(i, 10)

                dataId, errId := chainDB.Get([]byte("id" + nodeType + blockNumberString))
                dataIp, errIp := chainDB.Get([]byte("ip" + nodeType + blockNumberString))
                dataHash, errHash := chainDB.Get([]byte("hash" + nodeType + blockNumberString))
                if errId == nil && errIp == nil && errHash == nil {
                        hashes  = append(hashes, string(dataHash))
                        nodes   = append(nodes, string(dataId))
                        ips     = append(ips, string(dataIp))
                        numbers = append(numbers, blockNumberString)
                }
        }
        return hashes, nodes, numbers, ips, nil
}

// GetNodeId return enodeid in string format from *enode.Node
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
