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
        "enode://9e3db882078176028076cd4510c299e112069f969c7fa5c662517eb6cf119f54e74ab0348b110d7f3d82b40fe3a9c855f5e5c11673088378d3e4f56813dbca4b@140.82.50.249:30305",
        "enode://22b31e7fd3cf44072387d7e3765bfac545b9e623746e3ca19233b010997d6074a36e0851d8fe6fda98964315c81d445e5a2ba077960d9950446c514d56e61224@149.28.49.244:30305",
        "enode://3480c503c27de9220b9440eb1f4f92d7e16e396f223ace88661bf319d8871f179b0214d8121bbb229985090c1d2eaf7b046329dd27f605851b2860b7661922ca@45.32.170.196:30305",
        "enode://2b58a5e8fd16050d3de67edc081ccc3a95f41c502c34679064514cbe87ef71afea21a694f8e0a2cb4fca3a6b76c6702881e601a909e9d085f9050b17699dafa5@45.32.221.141:30305",
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
