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
	"enode://079763eff99b63261a1804ac5a28401d8d69137106e9022e588041a63c66594af910ffb28984eff4a009553246d5589d4075a61beb0c184a97d4bbdda793ec4a@157.230.144.98:30456", // Cat
	"enode://cbfe746bcc1793286470393f9d83d045943e71b9f35f81af27e7cbf0d91e3340f915f3c47ee99b5f2c91438ffb0b7267d9bb48da19c04d3786a26fd78cd53953@173.212.202.40:50505", // Monkey See
	"enode://f40a0e383a9096acdcd919af3b99fecf15c154f51d349725be703d0e68a285eadf91b9254374ab3a7c6d410f2d5eb25de6947dd16406b8e018e73df4efe8fcb8@178.238.229.61:30305", // XERO
	"enode://514c56ab4f0735bb92de18b04e91c2e7994302555e0ce31ec01dbcb8463ecb8a2cd692a7cea7f47faa01f9915018fbaf759d62c3ac629e7ddc6f684d571a31e2@45.63.22.232:60606",   // Explorer
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
