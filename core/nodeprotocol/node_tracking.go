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
        "os/user"
        "io/ioutil"
        "encoding/hex"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
        "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/params"
)

var activeNode *node.Node
var nodeProtocolData map[string]map[common.Hash]string
var EnodeId string
var awaitingDataRequests map[string][]string

func ActiveNode() *node.Node {
	return activeNode
}

func SetActiveNode(stack *node.Node) {
	activeNode = stack
        setNodeId()
        log.Info("Node Protocol Node ID Setup Confirmed", "ID", EnodeId)
}

// CheckNodeProtocolStatus determines if node protocol data has been initiated
func CheckNodeProtocolStatus() bool {
        if len(nodeProtocolData) > 0 {
                return true
        }
        return false

}

// AddDataRequest adds requests for node activity data to queue
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
}

// SetupNodeProtocolMap initiates node protocol mapping
func SetupNodeProtocolMapping() {
        log.Info("Initializing Node Protocol Data Mapping")
        nodeProtocolData = make(map[string]map[common.Hash]string)
        for _, nodeType := range params.NodeTypes {
                nodeProtocolData[nodeType.Name] = make(map[common.Hash]string)
        }
}

// CheckNodeStatus checks to see if specified node has been validated
func CheckNodeStatus(nodeType string, nodeId string, blockHash common.Hash) bool {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }
        if value, ok := nodeProtocolData[nodeType][blockHash]; ok {
                if nodeId == value {
                        log.Info("Node ID Found in Node Protocol Data", "Validated", "True", "Type", nodeType, "ID", nodeId)
                        return true
                }
        }
        log.Warn("Node ID Not Found in Node Protocol Data", "Validated", "False", "Type", nodeType, "ID", nodeId)
        return false
}

// CheckUpToDate checks to see if blockHash has been recorderd in mapping
func CheckUpToDate(nodeType string, blockHash common.Hash) bool {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }
        if _, ok := nodeProtocolData[nodeType][blockHash]; ok {
                return true
        }
        return false
}

// GetNodeProtocolData returns the nodeid at specified blockHash of specific node type
func GetNodeProtocolData(nodeType string, blockHash common.Hash) string {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }
        if nodeId, ok := nodeProtocolData[nodeType][blockHash]; ok {
                log.Info("Node ID Found in Node Protocol Data", "Type", nodeType, "ID", nodeId, "Hash", blockHash)
                return nodeId
        }

        log.Warn("Node ID Not Found in Node Protocol Data", "Type", nodeType, "Hash", blockHash)
        return ""
}

// UpdateNodeProtocolDate updates protocol mapping data for verified nodes
func UpdateNodeProtocolData(nodeType string, nodeId string, blockHash common.Hash) {
        if len(nodeProtocolData) == 0 {
                SetupNodeProtocolMapping()
        }
        if id, ok := nodeProtocolData[nodeType][blockHash]; ok {
                log.Warn("Duplicate Node ID Data Found in Node Protocol Data", "Type", nodeType, "ID", id, "Hash", blockHash)
        } else {
                nodeProtocolData[nodeType][blockHash] = nodeId
                log.Info("Node ID Saved To Node Protocol Data", "Type", nodeType, "ID", nodeId, "Hash", blockHash)
        }
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
