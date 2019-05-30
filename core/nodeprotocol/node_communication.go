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
        "io/ioutil"
        "os/user"
        "strconv"
        "encoding/hex"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
        "github.com/ethereum/go-ethereum/crypto"
)

var activeNode *node.Node
var nodeConsensusMap map[string]uint64

func ActiveNode() *node.Node {
	return activeNode
}

func SetActiveNode(stack *node.Node) {
	activeNode = stack
}

// Get user home directory from env
func getHomeDirectory() string {
    usr, err := user.Current()
    if err != nil {
        fmt.Print(err)
    }
    return usr.HomeDir
}

// Retrieve nodekey and calculate enodeid
func getNodeId() []byte {
    b, err := ioutil.ReadFile(getHomeDirectory() + "/.xerom/geth/nodekey")
    if err != nil {
        fmt.Print(err)
        return []byte{}
    }
    enodeId, err := crypto.HexToECDSA(string(b))
    if err != nil {
        fmt.Print(err)
        return []byte{}
    }
    pubkeyBytes := crypto.FromECDSAPub(&enodeId.PublicKey)[1:]
    return pubkeyBytes
}

func ConfirmNodeActivity(nodeID string) (bool, error) {
	log.Info("Attempting to Contact Node", "NodeID", nodeID)
	destinationNode, err := enode.ParseV4(nodeID)
	if err != nil {
		return true, fmt.Errorf("Node ID Format Error: %v", err)
	}
	// Try to contact node
	response := p2p.Resolve(activeNode.Server(), destinationNode)
	if response {
		log.Info("Node Contacted Successfully", "Node", "Verified")
		return true, nil
	}
	return false, fmt.Errorf("Node Communication Error: %v", err)
}

// SetupNodeProtocolMap initiates node protocol mapping
func SetupNodeProtocolMap() {
        nodeConsensusMap = make(map[string]uint64)
}

// UpdateNodeProtocolMap updates mapping data to latest block height
func UpdateNodeProtocolLocalBlock(newBlockHeight uint64) {

        nodeId := hex.EncodeToString(getNodeId())
        // Check to see if node protocl mapping has been initiated
        if oldBlockHeight, ok := nodeConsensusMap[nodeId]; ok {
                if newBlockHeight > oldBlockHeight {
                        nodeConsensusMap[nodeId] = newBlockHeight
                }
        } else {
                // Initiate node protocol mapping
                SetupNodeProtocolMap()
                nodeConsensusMap[nodeId] = newBlockHeight
        }
}

// UpdateNodeProtocolMap updates mapping data to latest block heights
// based on node messaging that has been received on node protocol layer
func UpdateNodeProtocolData(nodeId string, newBlockHeight uint64) {
        if oldBlockHeight, ok := nodeConsensusMap[nodeId]; ok {
                if newBlockHeight > oldBlockHeight {
                        nodeConsensusMap[nodeId] = newBlockHeight
                }
        } else {
                nodeConsensusMap[nodeId] = newBlockHeight
        }
}

// GetNodeProtocolMap returns node protocol mapping data
func GetNodeProtocolData() [][]string {
        var nodeConsensusData [][]string
        for nodeId, blockHeight := range nodeConsensusMap {
                nodeConsensusData = append(nodeConsensusData, []string{nodeId, strconv.FormatUint(blockHeight, 10)})
        }
        return nodeConsensusData
}
