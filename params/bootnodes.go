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

package params

// Ether-1 Bootnodes
var MainnetBootnodes = []string{
	// Xerom Bootnodes
        "enode://079763eff99b63261a1804ac5a28401d8d69137106e9022e588041a63c66594af910ffb28984eff4a009553246d5589d4075a61beb0c184a97d4bbdda793ec4a@157.230.144.98:30456",  // Cat -- Checked
	"enode://b2267f97ffedb56626de2911442daf70fa9904f34b1106a17a93086898f5d6ce5ba9f88f38b34c7f154aa3e96ae6956dbfc002d1dbbaae288850b4dca1c92e74@178.238.229.61:30305",  // RPC -- Checked
	"enode://cbfe746bcc1793286470393f9d83d045943e71b9f35f81af27e7cbf0d91e3340f915f3c47ee99b5f2c91438ffb0b7267d9bb48da19c04d3786a26fd78cd53953@173.212.202.40:50505",  // Monkey See -- Checked
	"enode://514c56ab4f0735bb92de18b04e91c2e7994302555e0ce31ec01dbcb8463ecb8a2cd692a7cea7f47faa01f9915018fbaf759d62c3ac629e7ddc6f684d571a31e2@45.63.22.232:60606",    // Explorer -- Checked
	"enode://fa73c40a0eda74b67314165099c6bbafb3eb02747bd5cc0c928e10addb1a6277195e758a6c1752f74ed29ec69c3bf05a6a5388654dc9e17f314bc9c32cabdd3b@54.38.158.16:30307",    // Checked
	"enode://7bbf02732b5dc875e3925aa68817059b81a1a9059d4b1132e9d379e9b4225cc18e30197ffc3ece8c38ed3d4ff113de0cda7e6628e501df087d85d218784999fc@54.38.158.16:30309",    // Checked
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
