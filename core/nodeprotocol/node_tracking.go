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
//var EnodeId string
//var blockChainService *core.BlockChain
//var awaitingDataRequests map[string][]string

func ActiveNode() *node.Node {
	return activeNode
}

func SetActiveNode(stack *node.Node) {
	activeNode = stack
        SetupNodeProtocolMapping()
        //setNodeId()
        //Enodeid = enode.ID()
        //log.Info("Node Protocol Node ID Setup Confirmed", "ID", enode.ID())
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
/*
func UpdateBlockChainService(bc *core.BlockChain) {
        blockChainService = bc
}

func GetBlockChainService() *core.BlockChain {
        return blockChainService
}
*/
// CheckNodeProtocolStatus determines if node protocol data has been initiated
func CheckNodeProtocolStatus() bool {
        if len(nodeProtocolData) > 0 {
                return true
        }
        return false

}

/*// AddDataRequest adds requests for node activity data to queue
func AddDataRequest(nodeType string, blockHash string) {
        if len(awaitingDataRequests) == 0 {
                awaitingDataRequests = make(map[string][]string)
        }
        awaitingDataRequests[nodeType + blockHash] = []string{nodeType, blockHash}
}

// GetDataRequests returns queued data requests and resets list
func GetDataRequests() [][]string {
        var dataRequests [][]string
        for _, requests := range awaitingDataRequests {
                dataRequests = append(dataRequests, requests)
        }
        awaitingDataRequests = make(map[string][]string)
        return dataRequests
}*/

// SetupNodeProtocolMap initiates node protocol mapping
func SetupNodeProtocolMapping() {
        if !protocolInitiationFlag {
                protocolInitiationFlag = true
                // Lock protocol mapping on initiation to prevent bad data during sync
                Lock()

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
        //log.Warn("Block Height", "Number", blockHeight, "Current Hash", currentHash, "Parent Hash", parentHash, "Grand Parent Hash", grandParentHash, "Great Grand Parent Hash", blockHash)
        //log.Warn("Block Height", "Number", blockHeight, "Current Node", nodeProtocolData[nodeType].dataMap[currentHash], "Parent Node", nodeProtocolData[nodeType].dataMap[parentHash], "Grand Parent Node", nodeProtocolData[nodeType].dataMap[grandParentHash], "Great Grand Parent Node", nodeProtocolData[nodeType].dataMap[blockHash])

        // log.Warn("Newest Saved Block Hash", "Hash", nodeProtocolData[nodeType].hashData[len(nodeProtocolData[nodeType].hashData) - 1])
        //log.Warn("Newest Saved Verified Node ID", "ID", nodeProtocolData[nodeType].nodeData[len(nodeProtocolData[nodeType].nodeData) - 1])

        log.Warn("Node ID Not Found in Node Protocol Data", "Validated", "False", "Type", nodeType, "ID Needing Verification", nodeId, "Hash", blockHash, "Saved ID", nodeProtocolData[nodeType].dataMap[blockHash])

        /*for i := 0; i < len(nodeProtocolData[nodeType].hashData); i++ {
                log.Warn("Outputting Current Node Protocol Data", "I", i, "Hash", nodeProtocolData[nodeType].hashData[i], "ID", nodeProtocolData[nodeType].nodeData[i])
        }*/


        return false
}
/*
// CheckNodeStateStatus checks to see if specified node has been validated and saved to statedb
func CheckNodeStateStatus(state *state.StateDB, nodeType string, nodeId string, blockHash common.Hash) bool {
        stateData := common.BytesToAddress(state.GetCode(common.BytesToAddress(append([]byte(nodeType), blockHash.Bytes()...))))

        if stateData == common.BytesToAddress([]byte(nodeId)) {
                log.Info("Node ID Found in Node Protocol State Data", "Validated", "True", "Type", nodeType, "ID", nodeId, "State Data", stateData, "Compare Data", common.BytesToAddress([]byte(nodeId)))
                return true
        }
        log.Warn("Node ID Not Found in Node Protocol State Data", "Validated", "False", "Type", nodeType, "ID", nodeId)
        return false
}
*/
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
                log.Info("Node ID Found in Node Protocol Data", "Type", nodeType, "ID", nodeId, "Hash", blockHash)
                return nodeId
        }

        log.Warn("Node ID Not Found in Node Protocol Data", "Type", nodeType, "Hash", blockHash)
        return ""
}

// UpdateNodeProtocolData updates protocol mapping data for verified nodes
func UpdateNodeProtocolData(nodeType string, nodeId string, peerId string, peerCount int, blockHash common.Hash) {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        // Check to see if protocol mapping is locked prior to modifying/adding data
        if !lockedFlag {
                if id, ok := nodeProtocolData[nodeType].dataMap[blockHash]; ok {
                        log.Warn("Duplicate Node ID Data Found in Node Protocol Data - Checking Peer Consensus", "Type", nodeType, "ID", id, "Hash", blockHash)

                        consensusData := nodeProtocolData[nodeType].nodeConsensusMap
                        if peerMap, ok := consensusData[blockHash][nodeId]; ok {
                                 if _, ok := peerMap[peerId]; ok {
                                 } else {
                                        peerMap[peerId] = peerId
                                        hashes := nodeProtocolData[nodeType].hashData
                                        nodes := nodeProtocolData[nodeType].nodeData
                                        dataMapping := nodeProtocolData[nodeType].dataMap
                                        if len(peerMap) > (peerCount / 2) {
                                                dataMapping[blockHash] = nodeId

                                                for i := 0; i < len(hashes); i++ {
                                                        if hashes[i] == blockHash {
                                                                nodes[i] = nodeId
                                                                break
                                                        }
                                                }
                                                log.Warn("Node ID Updated in Node Protocol Data - Consensus Override", "Type", nodeType, "ID", nodeId, "Hash", blockHash)
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
                                hashes = hashes[1:(hashDataLength-1)]
                        }
                        // Trim nodeDataLength - keep history to around 2 hours
                        if nodeDataLength > 550 {
                                nodes = nodes[1:(nodeDataLength-1)]
                        }

                        // One final check to be sure hash was not previously added
                        if nodeId, ok := nodeProtocolData[nodeType].dataMap[blockHash]; ok {
                                data := protocolData{nodeConsensusMap: nodeConsensus, dataMap: dataMapping, nodeData: nodes, hashData: hashes}
                                nodeProtocolData[nodeType] = data

                                log.Info("Node ID Saved To Node Protocol Data", "Type", nodeType, "ID", nodeId, "Hash", blockHash)
                        }
                }
        }
}

// SyncNodeProtocolData initially syncs validated node data from peerset
func SyncNodeProtocolData(nodeType string, nodes []string, hashes []string) {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }

        if lockedFlag {
                nodeDataLength := len(nodeProtocolData[nodeType].nodeData)
                hashDataLength := len(nodeProtocolData[nodeType].hashData)

                log.Warn("Incoming Node Data Length", "Length", len(nodes))
                log.Warn("Incoming Hash Data Length", "Length", len(hashes))

                localDataMap := nodeProtocolData[nodeType].dataMap
                localNodeConsensus := nodeProtocolData[nodeType].nodeConsensusMap

                if len(nodes) >= nodeDataLength && len(hashes) >= hashDataLength && len(nodes) == len(hashes) {
                        var updatedHashes []common.Hash
                        //for _, data := range hashes {
                        for i := 0; i < len(hashes); i++ {
                                //log.Info("Syncing Node Hash Data", "Type", nodeType, "Hash", common.HexToHash(data))
                                updatedHash := common.HexToHash(hashes[i])
                                updatedHashes = append(updatedHashes, updatedHash)
                                localDataMap[updatedHash] = nodes[i]
                                localNodeConsensus[updatedHash] = make(map[string]map[string]string)
                        }

                        /*for _, data := range nodes {
                                log.Info("Syncing Node ID Data", "Type", nodeType, "ID", data)
                        }*/
                        /*for _, data := range updatedHashes {
                                log.Info("Syncing Node ID Data", "Type", nodeType, "Hash", data)
                        }*/

                        updatedData := protocolData{nodeConsensusMap: localNodeConsensus, dataMap: localDataMap, nodeData: nodes, hashData: updatedHashes}
                        nodeProtocolData[nodeType] = updatedData

                       // Unlock data mapping now that sync is complete
                       Unlock()
                } else {
                       log.Error("Invalid Node Protocol Sync Data Receieved - Rolling Back Data Mapping")
                }
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
/*
// UpdateNodeProtocolStateData updates protocol mapping data for verified nodes
func UpdateNodeProtocolStateData(state *state.StateDB, nodeType string, nodeId string, blockHash common.Hash) {
        log.Info("Saving Node Protocol Data To State", "Type", nodeType, "ID", nodeId, "Hash", blockHash, "State Address", common.BytesToAddress(append([]byte(nodeType), blockHash.Bytes()...)))
        state.SetCode(common.BytesToAddress(append([]byte(nodeType), blockHash.Bytes()...)), []byte(nodeId))
}
*/

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
/*
// Get user home directory from env
func getHomeDirectory() string {
    usr, err := user.Current()
    if err != nil {
        log.Error("Unable To Find Home Directory" )
    }
    return usr.HomeDir
}

// Retrieve nodekey and calculate enodeid
func setNodeId() {
    b, err := ioutil.ReadFile(getHomeDirectory() + "/.xerom/geth/nodekey")
    if err != nil {
        fmt.Print(err)
    }
    enodeId, err := crypto.HexToECDSA(string(b))
    if err != nil {
        fmt.Print(err)
    }
    pubkeyBytes := crypto.FromECDSAPub(&enodeId.PublicKey)[1:]
    EnodeId = hex.EncodeToString(pubkeyBytes)
}
*/
