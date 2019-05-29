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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
)

var NodeFlag bool

func SetProtocolFlag(active bool) {
        NodeFlag = active
}

func CheckActiveNode() {
        if NodeFlag {
                log.Info("Node Protocol is Active", "Active", "True")
        }
}
// Clean up data
func stripCtlAndExtFromBytes(str string) string {
	b := make([]byte, len(str))
	var bl int
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 32 && c < 127 {
			b[bl] = c
			bl++
		}
	}
	return string(b[:bl])
}

// Calculate and return node remainder balance payment share
func GetNodeRemainder(state *state.StateDB, nodeCount uint64, remainderAddress common.Address) *big.Int {

	remainderBalance := state.GetBalance(remainderAddress)

	if remainderBalance.Cmp(big.NewInt(0)) > 0 && nodeCount > 0 {

                // Disburse remainder funds over extended period using a full days block count as divisor
		var remainderPayment *big.Int
		remainderPayment = new(big.Int).Div(remainderBalance, big.NewInt(6646))

		return remainderPayment
	}
	return big.NewInt(0)
}
