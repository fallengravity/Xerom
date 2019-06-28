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
	"enode://22b31e7fd3cf44072387d7e3765bfac545b9e623746e3ca19233b010997d6074a36e0851d8fe6fda98964315c81d445e5a2ba077960d9950446c514d56e61224@149.28.49.244:30305",   // Checked
	"enode://3480c503c27de9220b9440eb1f4f92d7e16e396f223ace88661bf319d8871f179b0214d8121bbb229985090c1d2eaf7b046329dd27f605851b2860b7661922ca@45.32.170.196:30305",   // Checked
	"enode://2c2fc4d4eb62081096aad210daaaba59f6aa84dc1408351f05d35143c6b08a8bae8e5283fc59756e0dca99b33e541d4307ea41f295ad5b486942e4b7ee74cbf9@45.32.152.162:30305",   // Checked
	"enode://9e3db882078176028076cd4510c299e112069f969c7fa5c662517eb6cf119f54e74ab0348b110d7f3d82b40fe3a9c855f5e5c11673088378d3e4f56813dbca4b@140.82.50.249:30305",   // Checked
	"enode://2b58a5e8fd16050d3de67edc081ccc3a95f41c502c34679064514cbe87ef71afea21a694f8e0a2cb4fca3a6b76c6702881e601a909e9d085f9050b17699dafa5@45.32.221.141:30305",   // Checked
	"enode://d3b89d32e89febbd0c28f58c1b5c72fd770343739cec9b9c525b0cc7f2383d68720fe752f06259d157d3b4489ca6214085b0e337512b4c0a53f93e5020d89a21@139.180.160.239:30305", // Checked
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
