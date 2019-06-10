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

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
        "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/params"
)

var activeNode *node.Node
var nodeProtocolData map[string]protocolData
type protocolData struct {
        nodeConsensusMap map[common.Hash]map[string]map[string]string
        dataMap map[common.Hash]string
        nodeData []string
        hashData []common.Hash
}
var lockedFlag bool
var protocolInitiationFlag bool

func ActiveNode() *node.Node {
	return activeNode
}

func SetActiveNode(stack *node.Node) {
	activeNode = stack
        SetupNodeProtocolMapping()
        Lock()
}

func Unlock() {
        lockedFlag = false
}

func Lock() {
        lockedFlag = true
}

func IsLocked() bool {
        return lockedFlag
}

func IsSynced(nodeType string) bool {
        if  len(nodeProtocolData[nodeType].nodeData) > 100 && len(nodeProtocolData[nodeType].hashData) > 100 {
                log.Info("Node Protocol Mapping Sync Status", "Status", "Synced")
                return true
        }
        log.Warn("Node Protocol Mapping Sync Status", "Status", "Not Synced", "Node Entries", len(nodeProtocolData[nodeType].nodeData), "Hash Entries", len(nodeProtocolData[nodeType].hashData))
        return false
}

// CheckNodeProtocolStatus determines if node protocol data has been initiated
func CheckNodeProtocolStatus() bool {
        if len(nodeProtocolData) > 0 {
                return true
        }
        return false

}

// SetupNodeProtocolMap initiates node protocol mapping
func SetupNodeProtocolMapping() {
        if !protocolInitiationFlag {
                protocolInitiationFlag = true
                // Lock protocol mapping on initiation to prevent bad data during sync
                //Lock()

                log.Info("Initializing Node Protocol Data Mapping")
                nodeProtocolData = make(map[string]protocolData)
                for _, nodeType := range params.NodeTypes {
                        nodeConsensus := make(map[common.Hash]map[string]map[string]string)
                        hashes := make(map[common.Hash]string)
                        var data1 []string
                        var data2 []common.Hash
                        data := protocolData{nodeConsensusMap: nodeConsensus, dataMap: hashes, nodeData: data1, hashData: data2}
                        nodeProtocolData[nodeType.Name] = data
                }
        }
}

func ResetNodeProtocolData() {
        protocolInitiationFlag = false
        SetupNodeProtocolMapping()
}

// CheckNodeStatus checks to see if specified node has been validated
func CheckNodeStatus(blockHeight uint64, currentHash common.Hash, parentHash common.Hash, grandParentHash common.Hash, nodeType string, nodeId string, blockHash common.Hash) bool {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        if value, ok := nodeProtocolData[nodeType].dataMap[blockHash]; ok {
                if nodeId == value {
                        log.Info("Node ID Found in Node Protocol Data", "Validated", "True", "Type", nodeType, "ID", nodeId)
                        return true
                }
        }
        log.Warn("Node ID Not Found in Node Protocol Data", "Validated", "False", "Type", nodeType, "ID Needing Verification", nodeId, "Hash", blockHash, "Saved ID", nodeProtocolData[nodeType].dataMap[blockHash])

        return false
}

// CheckUpToDate checks to see if blockHash has been recorderd in mapping
func CheckUpToDate(nodeType string, blockHash common.Hash) bool {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        if _, ok := nodeProtocolData[nodeType].dataMap[blockHash]; ok {
                return true
        }
        return false
}

// GetNodeProtocolData returns the nodeid at specified blockHash of specific node type
func GetNodeProtocolData(nodeType string, blockHash common.Hash) string {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        if nodeId, ok := nodeProtocolData[nodeType].dataMap[blockHash]; ok {
                return nodeId
        }

        return ""
}

// UpdateNodeProtocolData updates protocol mapping data for verified nodes
func UpdateNodeProtocolData(nodeType string, nodeId string, peerId string, peerCount int, blockHash common.Hash) {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        // Check to see if protocol mapping is locked prior to modifying/adding data
        if !lockedFlag {
                Lock()  // Lock to eliminate data errors
                if _, ok := nodeProtocolData[nodeType].dataMap[blockHash]; ok {

                        consensusData := nodeProtocolData[nodeType].nodeConsensusMap
                        if peerMap, ok := consensusData[blockHash][nodeId]; ok {
                                 if _, ok := peerMap[peerId]; ok {
                                 } else {
                                        peerMap[peerId] = peerId
                                        hashes := nodeProtocolData[nodeType].hashData
                                        nodes := nodeProtocolData[nodeType].nodeData
                                        dataMapping := nodeProtocolData[nodeType].dataMap
                                        if len(peerMap) > (peerCount / 2) && dataMapping[blockHash] != nodeId {
                                                dataMapping[blockHash] = nodeId

                                                for i := 0; i < len(hashes); i++ {
                                                        if hashes[i] == blockHash {
                                                                nodes[i] = nodeId
                                                                break
                                                        }
                                                }
                                                log.Warn("Node Protocol Data Updated - Node ID Consensus Override", "Type", nodeType, "ID", nodeId, "Hash", blockHash)
                                         }
                                         consensusData[blockHash][nodeId] = peerMap
                                         data := protocolData{nodeConsensusMap: consensusData, dataMap: dataMapping, nodeData: nodes, hashData: hashes}
                                         nodeProtocolData[nodeType] = data
                                 }
                        }
                } else {

                        nodeConsensus := nodeProtocolData[nodeType].nodeConsensusMap
                        dataMapping := nodeProtocolData[nodeType].dataMap

                        nodeConsensus[blockHash] = make(map[string]map[string]string)
                        consensusAddition := make(map[string]string)

                        consensusAddition[peerId] = peerId
                        nodeConsensus[blockHash][nodeId] = consensusAddition

                        dataMapping[blockHash] = nodeId

                        nodes := append(nodeProtocolData[nodeType].nodeData, nodeId)
                        hashes := append(nodeProtocolData[nodeType].hashData, blockHash)

                        hashDataLength := len(hashes)
                        nodeDataLength := len(nodes)

                        // Trim hashDataLength - keep history to around 2 hours
                        if hashDataLength > 550 {
                                lastBlockHash := hashes[0]
                                delete(dataMapping, lastBlockHash)
                                delete(nodeConsensus, lastBlockHash)
                                hashes = hashes[1:(hashDataLength)]
                        }
                        // Trim nodeDataLength - keep history to around 2 hours
                        if nodeDataLength > 550 {
                                nodes = nodes[1:(nodeDataLength)]
                        }

                        data := protocolData{nodeConsensusMap: nodeConsensus, dataMap: dataMapping, nodeData: nodes, hashData: hashes}
                        nodeProtocolData[nodeType] = data
                        log.Info("Node Protocol Data Updated - Node ID Saved To Node Protocol Data", "Type", nodeType, "ID", nodeId, "Hash", blockHash)
                }
                Unlock()
        }
}

// SyncNodeProtocolData initially syncs validated node data from peerset
func SyncNodeProtocolData(nodeType string, nodes []string, hashes []string) {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        Lock() // Lock so no data is added during sync

        if lockedFlag {

                nodeDataLength := len(nodeProtocolData[nodeType].nodeData)
                hashDataLength := len(nodeProtocolData[nodeType].hashData)

                log.Warn("Incoming Node Data Length", "Length", len(nodes))
                log.Warn("Incoming Hash Data Length", "Length", len(hashes))

                localDataMap := nodeProtocolData[nodeType].dataMap
                localNodeConsensus := nodeProtocolData[nodeType].nodeConsensusMap

                if len(nodes) >= nodeDataLength && len(hashes) >= hashDataLength && len(nodes) == len(hashes) {

                        log.Info("Saving Received Node Protocol Data Into Mapping")

                        var updatedHashes []common.Hash
                        //for _, data := range hashes {
                        for i := 0; i < len(hashes); i++ {
                                //log.Info("Syncing Node Hash Data", "Type", nodeType, "Hash", common.HexToHash(data))
                                updatedHash := common.HexToHash(hashes[i])
                                updatedHashes = append(updatedHashes, updatedHash)
                                localDataMap[updatedHash] = nodes[i]
                                localNodeConsensus[updatedHash] = make(map[string]map[string]string)
                        }

                        updatedData := protocolData{nodeConsensusMap: localNodeConsensus, dataMap: localDataMap, nodeData: nodes, hashData: updatedHashes}
                        nodeProtocolData[nodeType] = updatedData

                       // Unlock data mapping now that sync is complete
                       Unlock()
                } else {
                       log.Error("Invalid Node Protocol Sync Data Receieved - Rolling Back Data Mapping")
                }
                // Unlock data mapping now that sync attempt is complete
                Unlock()
        }
}

// GetNodeProtocolSyncNodeData returns all node data
func GetNodeProtocolSyncNodeData(nodeType string) []string {
        localNodeData := nodeProtocolData[nodeType].nodeData
        return localNodeData
}

// GetNodeProtocolSyncHashData returns all hash data
func GetNodeProtocolSyncHashData(nodeType string) []string {
        var hashes []string
        localHashData := nodeProtocolData[nodeType].hashData
        for _, data := range localHashData {
                hashes = append(hashes, data.String())
        }
        return hashes
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
