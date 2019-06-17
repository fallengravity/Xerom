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
        "sync"
        "errors"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
        "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/core/nodeprotocolmessaging"
)

var activeNode *node.Node
var nodeProtocolData map[string]protocolData
var mux = sync.RWMutex{}
type protocolData struct {
        nodeConsensusMap map[common.Hash]map[string]map[string]string
}
var protocolInitiationFlag bool
var protocolSyncFlag bool
var targetSyncNumber uint64
var protocolSyncStatus map[string]bool

func ActiveNode() *node.Node {
	return activeNode
}

func SetActiveNode(stack *node.Node) {
        nodeprotocolmessaging.SyncWg.Add(1)
	activeNode = stack
        SetupNodeProtocolMapping()
        protocolSyncFlag = false
        protocolSyncStatus = make(map[string]bool)
        initializeProtocolHeadTracker()
}

func SetSyncFlag(synced bool) {
        protocolSyncFlag = synced
}

func SetTargetSyncNumber(number uint64) {
        targetSyncNumber = number
}

func initializeProtocolHeadTracker() {
        protocolHeadTracker := nodeprotocolmessaging.GetChan()

        done := make(chan uint64)

        go nodeprotocolmessaging.Background(protocolHeadTracker, done)
}

func CheckSyncStatus() bool {
        if IsSynced() {
                return true
        }
        synced := true
        for _, nodeType := range params.NodeTypes {
                if !protocolSyncStatus[nodeType.Name] {
                        headNumber, _ := GetHeadNodeProtocolDataEntry(nodeType.Name)
                        if headNumber >= targetSyncNumber {
                                 protocolSyncStatus[nodeType.Name] = true
                        } else {
                                 synced = false
                        }
                }
        }
        SetSyncFlag(synced)
        return synced
}

// IsSynced lets us know if the state has specific data saved
func IsSynced() bool {
        return protocolSyncFlag
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

                log.Info("Initializing Node Protocol Data Mapping")
                nodeProtocolData = make(map[string]protocolData)
                for _, nodeType := range params.NodeTypes {
                        nodeConsensus := make(map[common.Hash]map[string]map[string]string)
                        data := protocolData{nodeConsensusMap: nodeConsensus}
                        nodeProtocolData[nodeType.Name] = data
                }
        }
}

// CheckNodeStatus checks to see if specified node has been validated
//func CheckNodeStatus(blockHeight uint64, currentHash common.Hash, parentHash common.Hash, grandParentHash common.Hash, nodeType string, nodeId string, blockHash common.Hash, blockNumber uint64) bool {
func CheckNodeStatus(nodeType string, nodeId common.Hash, nodeIp common.Hash, blockHash common.Hash, blockNumber uint64) bool {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        data, err := GetNodeDataState(nodeType, blockNumber)
        if err == nil {
                if nodeId == common.BytesToHash([]byte(data.Id)) && nodeIp == common.BytesToHash([]byte(data.Ip)) {
                        log.Info("Node ID Found in Node Protocol Data", "Validated", "True", "Type", nodeType, "ID", nodeId, "IP", nodeIp)
                        return true
                }
        }
        log.Warn("Node ID Not Found in Node Protocol Data", "Validated", "False", "Type", nodeType, "ID Needing Verification", nodeId, "Hash", blockHash, "IP", nodeIp, "Saved ID", common.BytesToHash([]byte(data.Id)), "Saved Hash", data.Hash, "Saved IP", common.BytesToHash([]byte(data.Ip)))

        return false
}

// CheckUpToDate checks to see if blockHash has been recorded in mapping
func CheckUpToDate(nodeType string, blockHash common.Hash, blockNumber uint64) bool {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        _, err := GetNodeDataState(nodeType, blockNumber)
        if err == nil {
                return true
        }
        return false
}

func GetHeadNodeProtocolDataEntry(nodeType string) (uint64, common.Hash) {
        number, data, err := GetNodeDataStateLatest(nodeType)
        if err == nil {
                return number, data.Hash
        }
        return 0, common.HexToHash("")
}

// GetNodeProtocolData returns the nodeid at specified blockHash of specific node type
func GetNodeProtocolData(nodeType string, blockHash common.Hash, blockNumber uint64) (string, string, error) {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        data, err := GetNodeDataState(nodeType, blockNumber)
        if err == nil {
                return data.Id, data.Ip, nil
        }
        return "", "", errors.New("Node Protocol Data Not Found")
}

// RollBackNodeProtocolData removes a specific amount of data starting with head
// Used for rolling back when node goes out of sync - returns new head block number
func RollBackNodeProtocolData(nodeType string, count uint64) {
        mux.Lock()
        defer mux.Unlock()
        currentHeadBlockNumber, data, _ := GetNodeDataStateLatest(nodeType)
        for i := currentHeadBlockNumber; i > (currentHeadBlockNumber + count); i-- {
                _, err := GetNodeDataState(nodeType, i)
                if err == nil {
                        DeleteNodeDataState(nodeType, data.Hash, data.Number)
                        if _, ok := nodeProtocolData[nodeType].nodeConsensusMap[data.Hash]; ok {
                                delete(nodeProtocolData[nodeType].nodeConsensusMap, data.Hash)
                        }
                        //go RemoveNodeProtocolData(nodeType, data.Hash, data.Number)
                }
        }
}

// RemoveNodeProtocolData removes protocal data
func RemoveNodeProtocolData(nodeType string, blockHash common.Hash, blockNumber uint64) {
        DeleteNodeDataState(nodeType, blockHash, blockNumber)
        mux.Lock()
        defer mux.Unlock()
        for {
                if _, ok := nodeProtocolData[nodeType].nodeConsensusMap[blockHash]; ok {
                        delete(nodeProtocolData[nodeType].nodeConsensusMap, blockHash)
                        return
                }
        }
}

// UpdateNodeProtocolData updates protocol mapping data for verified nodes
func UpdateNodeProtocolData(nodeType string, nodeId string, nodeIp string, peerId string, peerCount int, blockHash common.Hash, blockNumber uint64, syncing bool) {
        if !syncing {
               nodeprotocolmessaging.SyncWg.Wait()
        }
        mux.Lock()
        defer mux.Unlock()
        log.Trace("Attempting To Update Node Protocol Data", "Type", nodeType, "ID", nodeId, "IP", nodeIp, "Number", blockNumber, "Hash", blockHash)
        consensusData := nodeProtocolData[nodeType].nodeConsensusMap
        if peerMap, ok := consensusData[blockHash][nodeId]; ok  && !syncing {

                localData, err := GetNodeDataState(nodeType, blockNumber)
                if err == nil && localData.Id == "" && nodeId != "" {
                        consensusData[blockHash] = make(map[string]map[string]string)
                        consensusAddition := make(map[string]string)
                        consensusAddition[peerId] = peerId
                        consensusData[blockHash][nodeId] = consensusAddition
                        data := protocolData{nodeConsensusMap: consensusData}
                        nodeProtocolData[nodeType] = data
                        ReplaceNodeDataState(nodeType, blockHash, nodeId, blockNumber, nodeIp)
                        return
                } else if  _, ok := peerMap[peerId]; ok && err == nil {
                        return
                } else {
                        peerMap[peerId] = peerId
                        if len(peerMap) > (peerCount / 2) && localData.Id != nodeId {
                                ReplaceNodeDataState(nodeType, blockHash, nodeId, blockNumber, nodeIp)
                                log.Trace("Node Protocol Data Updated - Node Consensus Achieved", "Type", nodeType, "ID", nodeId, "IP", nodeIp, "Hash", blockHash)
                        }
                        consensusData[blockHash][nodeId] = peerMap
                        data := protocolData{nodeConsensusMap: consensusData}
                        nodeProtocolData[nodeType] = data
                        return
                }
        } else {

                // Initiate consensus mapping
                consensusData[blockHash] = make(map[string]map[string]string)
                consensusAddition := make(map[string]string)

                consensusAddition[peerId] = peerId
                consensusData[blockHash][nodeId] = consensusAddition

                data := protocolData{nodeConsensusMap: consensusData}
                nodeProtocolData[nodeType] = data
                log.Trace("Node Protocol Data Updated - Initial Data Added", "Type", nodeType, "ID", nodeId, "IP", nodeIp, "Hash", blockHash)

                // With enough nodes this should go away in order to only add via consensus unless syncing
                ReplaceNodeDataState(nodeType, blockHash, nodeId, blockNumber, nodeIp)
                return
        }
}

// SyncNodeProtocolDataGroup adds a slice of NodeData to state is consenus is reached
func SyncNodeProtocolDataGroup(nodeType string, nodeData map[uint64]NodeData, peerId string, peerCount int) {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        largestBlockNumber := uint64(0)
        smallestBlockNumber := uint64(99999999)
        for blockNumber, data := range nodeData {
                if blockNumber > largestBlockNumber {
                        largestBlockNumber = blockNumber
                } else if blockNumber < smallestBlockNumber {
                        smallestBlockNumber = blockNumber
                }
                go UpdateNodeProtocolData(nodeType, data.Id, data.Ip, peerId, peerCount, data.Hash, blockNumber, true)
        }

        if len(nodeData) > 0 {
                log.Info("Importing Node Protocol Data", "Entries", len(nodeData), "Blocks", strconv.FormatUint(smallestBlockNumber, 10) + "->" + strconv.FormatUint(largestBlockNumber, 10))
        }
        nodeprotocolmessaging.Set(largestBlockNumber)
}

//func GetNodeProtocolDataGroup(nodeType string, startBlock uint64, endBlock uint64) (map[uint64]NodeData, error) {
func GetNodeProtocolDataGroup(nodeType string, startBlock uint64, endBlock uint64) ([]string, []string, []string, []string, error) {
        var hashes []string
        var nodes  []string
        var ips  []string
        var numbers []string
        for i := startBlock; i <= endBlock; i++ {
                data, err := GetNodeDataState(nodeType, i)
                if err == nil {
                        hashes  = append(hashes, data.Hash.String())
                        nodes   = append(nodes, data.Id)
                        ips     = append(ips, data.Ip)
                        numbers = append(numbers, strconv.FormatUint(i, 10))
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
