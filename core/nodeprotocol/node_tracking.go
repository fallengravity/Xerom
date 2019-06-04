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
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/p2p/enode"
        "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

var activeNode *node.Node
var nodeConsensusMap map[string]uint64
var localNodeConsensusMap map[string]nodeConsensus

type nodeConsensus struct {
        blockHeight uint64
	peerIdMap map[string]uint64
}

func ActiveNode() *node.Node {
	return activeNode
}

func SetActiveNode(stack *node.Node) {
	activeNode = stack
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

// CheckNodeStatus checks the node protocol map and verifies the node is
// active and not stale or invalid
func CheckNodeStatus(state *state.StateDB, blockHeight uint64, nodeId string) bool {
        nodeStatus := common.BytesToAddress(state.GetCode(common.HexToAddress(nodeId)))
        compareStatus := common.BytesToAddress([]byte("TRUE"))
        if nodeStatus == compareStatus {
                log.Info("Retrieving Node Status From State", "Status", nodeStatus, "Confirmed", "True")
                return true
        } else {
                log.Info("Retrieving Node Status From State", "Status", nodeStatus, "Confirmed", "False")
        }
        return false
}

// SetNodeStatus records node active status in statedb
func SetNodeStatus(state *state.StateDB, blockHeight uint64, nodeId string) {
        if len(nodeConsensusMap) == 0 {
                SetupNodeProtocolMapping()
        }
        if nodeBlockHeight, ok := nodeConsensusMap[nodeId]; ok {
                // Only allow a state update is nodeBlockHeight has been confirmed by peers to be active
                if nodeBlockHeight > 0 {
                        // Verifify that node is active within last 276 block (roughly an hour)
                        if blockHeight < uint64(276) || nodeBlockHeight > (blockHeight - uint64(276)) {
                                state.SetCode(common.HexToAddress(nodeId), []byte("TRUE"))
                        } else {
                                state.SetCode(common.HexToAddress(nodeId), []byte("FALSE"))
                        }
                }
        }
}

// SetupNodeProtocolMap initiates node protocol mapping
func SetupNodeProtocolMapping() {
        log.Info("Initializing Node Protocol Data Mapping")
        nodeConsensusMap = make(map[string]uint64)
        localNodeConsensusMap = make(map[string]nodeConsensus)
}

// UpdateRemoteNodeProtocolData updates mapping data of the message sender
// to latest block height
func UpdateRemoteNodeProtocolData(peerId string, newBlockHeight uint64) {
        if len(nodeConsensusMap) == 0 {
                SetupNodeProtocolMapping()
        }
        // Check to see if node protocl mapping has been initiated
        // then update known block height
        if nodeConsensusStruct, ok := localNodeConsensusMap[peerId]; ok {
                if newBlockHeight > nodeConsensusStruct.blockHeight {
                        nodeConsensusStruct.blockHeight = newBlockHeight
                        nodeConsensusStruct.peerIdMap = make(map[string]uint64)
                        nodeConsensusMap[peerId] = newBlockHeight
                }
        } else {
               var nodeConsensusData nodeConsensus
               nodeConsensusData.blockHeight = newBlockHeight
               nodeConsensusData.peerIdMap = make(map[string]uint64)
               localNodeConsensusMap[peerId] = nodeConsensusData
               nodeConsensusMap[peerId] = newBlockHeight
        }
}

// UpdateNodeProtocolMap updates mapping data to latest block heights
// based on node messaging that has been received on node protocol layer
func UpdateNodeProtocolData(nodeId string, newBlockHeight uint64, peerId string, peerCount int) {
        if len(nodeConsensusMap) == 0 {
                SetupNodeProtocolMapping()
        }
        if nodeConsensusStruct, ok := localNodeConsensusMap[nodeId]; ok {
                if newBlockHeight > nodeConsensusStruct.blockHeight {

                        nodeConsensusStruct.peerIdMap[peerId] = newBlockHeight

                        // Set peerCount comparison number
                        peerCountComparison := 1
                        if peerCount > 20 {
                                peerCountComparison = peerCount / 2
                        }

                        // Update consensus mapping blockheight is enough peers confirm node as active
                        // and clear peer mapping to start over
                        if len(nodeConsensusStruct.peerIdMap) > peerCountComparison {
                                updatedBlockHeight := GetMinMappingValue(nodeConsensusStruct.peerIdMap)
                                nodeConsensusStruct.blockHeight = updatedBlockHeight
                                nodeConsensusStruct.peerIdMap = make(map[string]uint64)
                                nodeConsensusMap[nodeId] = updatedBlockHeight
                        } else {
                              nodeConsensusStruct.peerIdMap[peerId] = newBlockHeight
                        }
                        localNodeConsensusMap[nodeId] = nodeConsensusStruct
                }
        } else {
                var nodeConsensusData nodeConsensus
                nodeConsensusData.blockHeight = 0
                nodeConsensusData.peerIdMap = make(map[string]uint64)
                nodeConsensusData.peerIdMap[peerId] = newBlockHeight
                localNodeConsensusMap[nodeId] = nodeConsensusData
                nodeConsensusMap[nodeId] = 0
        }
}

// GetMinMappingValue min value from node protocol mapping data
func GetMinMappingValue(peerIdMap map[string]uint64) uint64 {
        minMappingValue := uint64(0)
        for _, value := range peerIdMap {
                if minMappingValue == 0 {
                        minMappingValue = value
                } else if value < minMappingValue {
                        minMappingValue = value
                }
        }
        return minMappingValue
}

// GetNodeProtocolMap returns node protocol mapping data
func GetNodeProtocolData() [][]string {
        var nodeConsensusData [][]string
        for nodeId, blockHeight := range nodeConsensusMap {
                nodeConsensusData = append(nodeConsensusData, []string{nodeId, strconv.FormatUint(blockHeight, 10)})
        }
        return nodeConsensusData
}
