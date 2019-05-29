// Copyright 2017 The go-ethereum Authors
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

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xf8D\u1511\x89\x91\xf9\x0f\xcdZ\x8d\x1e.\x81_\x8a\xe3\xb8\x1e\x8c\x1a(\x92\x8b\aZK\xa1\xd8\xe4\xb2\b\xe3\x8e8\u153b\xe7D#m\xc5\xfb\xa5\xcb\\c2\x9e\x96B\xc8y\xe2n\\\x8b\bE\x95\x16\x14\x01HI\xff\xff\xff"
const testnetAllocData = ""
const rinkebyAllocData = ""
const goerliAllocData = ""
