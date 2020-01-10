// Copyright 2015 The go-ethereum Authors
// Copyright 2019 The Ether-1 Development Team
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
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

var id = "ba7e67b457e0c746bed770300a3fef89c24f77cf467acda9131f271186b9b81d"
var keys = []string{"fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19", "8dbe9fe18ddce807baf17ac41a9246c023a63df6a4fb481182fa0b4958605665", "a3ccdced2c63c721202145402c09d63a3b747dd62e5f47bdc573f5d00bfc1663", "fca0bb13ac202e001d9a11d5bfb62b4e91564774b7bd6a5f69c3cb6a05392050", "19815fea11e6c188f90ca15cb87b85f2a919c641e447f9636830f734f73ee45d"}

func txTest() {
	var validations [][]byte

	privateKey, err := crypto.HexToECDSA(id)
	if err != nil {
		log.Error("Error", "Error", err)
	}

	publicKey := privateKey.Public()
	publicId, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("Error", "Error", "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	idBytes := crypto.FromECDSAPub(publicId)

	for _,key := range keys {
		// Respond to validation requests
		nodePrivateKey, err := crypto.HexToECDSA(key)
		if err != nil {
			log.Error("Error", "Error", err)
		}

		nodePublicKey := nodePrivateKey.Public()
		nodeId, ok := nodePublicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Error("Error", "Error", "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		nodePublicAddr := crypto.PubkeyToAddress(*nodeId)
		nodeIdString := fmt.Sprintf("%x", crypto.FromECDSAPub(nodeId)[1:])

		log.Info("Initial Address", "Address", nodePublicAddr)
		fmt.Println("Initial ID: " + nodeIdString)
		fmt.Println("Initial Address: " + nodePublicAddr.String())

		// After Validations have been revceived
		validations = append(validations, SignNodeProtocolValidation(nodePrivateKey, idBytes))
	}

	//nodeValidations := NodeValidation{Id: idBytes, Validations:validations}
	/*tx := SendSignedNodeProtocolTx(privateKey, nodeValidations)


	// POS - Validation on DNP tx receipt
	var v NodeValidations
	err = json.Unmarshal(tx.Data(), &v)
	if err != nil {
		log.Error("Error", "Error", err)
	}

	for _, validation := range v.Validations {
		ValidateNodeProtocolSignature(idBytes, validation, idBytes)
	}*/
}
