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

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://573b6607cd59f241e30e4c4943fd50e99e2b6f42f9bd5ca111659d309c06741247f4f1e93843ad3e8c8c18b6e2d94c161b7ef67479b3938780a97134b618b5ce@52.56.136.200:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
