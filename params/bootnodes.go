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

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Xerom Maine network
var MainnetBootnodes = []string{
	"enode://6a4d405ba53336f0d12b2d37250bd8058b0b32c4d1bc040fc82c0566e7be105962f4464984c54c6754bf21d4829f281aac75126dc5ae56f0b8fa1322115ea72f@149.28.49.244:30305",
	"enode://80c3fb881c2924bfd0fc2a4b9fb618332560683966acebcc65f24d17839e6d5b4071bc53f19a016d9633a1cf9f39540813cea8973d11807d5ee0d0e7a1c64290@45.32.170.196:30305",
	"enode://f5e8a147396ca5947ffb505ef2363f2e6163cb5b13dba26770c1a090e88df3ac438f9d7149bdc4ad68cef2c1f7ffd11095ca3bfcfbd0057ecf202a17424eb120@173.212.202.40:50505",
	"enode://eaf87d3bf2717886578f49a7be340c465bb631d9d7bd07f2f410ca0e66bee5c5a7ab2bf5ecb0f8a57df4e11b095e4936cf253f51e866bbdba1e90235bd7f9b5a@45.63.22.232:60606",
	"enode://fd69edb2b17b185b9b9591e774204208abafea803196ffc823cc03ef2168f786e13e210fd5e7507ed60b430d5f1f454d388789be294216b6e88d3f42cd6e81f4@178.238.229.61:30305",
	"enode://c2fcba64f2dc285e4dd4578632dbf96ce3057547f08877e6f11d41042de8ee76efc34dd71fc9642c992149c36948b2db4e657825daef4e4ce34c54cb82b72319@144.91.86.189:30307",
	// Crypto_Saiyan's Bootnode
	"enode://bbd7965ea3b8d8b22c608327cfa63c540087cbe6f8764d2e9e0dfe64faf0df737743e01b03717d186ef9cf2e4e29c31e3c0a421f4ece42c4bcb352f8fe0997ca@144.91.117.137:30307",
	// Pistol's Bootnode
	"enode://db7c09733ae4ad7c26f98e1a112d5de2c6de915dcb1cfa91f3c24ec328fb3af02ff62dea922989b43026948a4decf866f998b410a4e04b0c9a3042246756c13f@207.180.197.210:30307",
	// Dylie's Bootnode
	"enode://8affe6d20222e6b28f1bae64ed3c73cf4779c18bcc5efcd0fbea1ee6c4d2a4d588b9d7e9be9f9fa59af28f169feaa9181e7ed523d0beb1cdc3850bdab6180c5b@5.189.157.250:30307",
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
