// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nodeprotocol

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NodeValidationsABI is the input ABI used to generate the binding from.
const NodeValidationsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"nodeCheckIn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastNodeActivity\",\"outputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"},{\"name\":\"signatureCount\",\"type\":\"uint256\"},{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NodeValidationsBin is the compiled bytecode used for deploying new contracts.
const NodeValidationsBin = `0x608060405234801561001057600080fd5b5061088c806100206000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c511613b811461005b578063c7cfc9ea1461007d578063ed5b5015146100b6575b600080fd5b34801561006757600080fd5b5061007b610076366004610606565b6100e3565b005b34801561008957600080fd5b5061009d6100983660046105e0565b6101ab565b6040516100ad9493929190610723565b60405180910390f35b3480156100c257600080fd5b506100d66100d13660046105e0565b610272565b6040516100ad9190610768565b6100eb61036e565b506040805160a081018252328082526020808301868152865184860152606084018690524360808501526000928352828252939091208251815473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff909116178155925180519293849390926101729260018501929101906103b4565b506040820151600282015560608201518051610198916003840191602090910190610411565b5060808201518160040155905050505050565b60006020818152918152604090819020805460028083015460038401805486516101006001831615026000190190911693909304601f810188900488028401880190965285835273ffffffffffffffffffffffffffffffffffffffff909316959094919291908301828280156102625780601f1061023757610100808354040283529160200191610262565b820191906000526020600020905b81548152906001019060200180831161024557829003601f168201915b5050505050908060040154905084565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260208181526040808320600101805482518185028101850190935280835260609492939192909184015b828210156103635760008481526020908190208301805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561034f5780601f106103245761010080835404028352916020019161034f565b820191906000526020600020905b81548152906001019060200180831161033257829003601f168201915b5050505050815260200190600101906102b8565b505050509050919050565b60a060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016000815260200160608152602001600081525090565b828054828255906000526020600020908101928215610401579160200282015b8281111561040157825180516103f1918491602090910190610411565b50916020019190600101906103d4565b5061040d92915061048b565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061045257805160ff191683800117855561047f565b8280016001018555821561047f579182015b8281111561047f578251825591602001919060010190610464565b5061040d9291506104b1565b6104ae91905b8082111561040d5760006104a582826104cb565b50600101610491565b90565b6104ae91905b8082111561040d57600081556001016104b7565b50805460018160011615610100020316600290046000825580601f106104f1575061050f565b601f01602090049060005260206000209081019061050f91906104b1565b50565b600061051e82356107f3565b9392505050565b6000601f8201831361053657600080fd5b8135610549610544826107a0565b610779565b81815260209384019390925082018360005b8381101561058757813586016105718882610591565b845250602092830192919091019060010161055b565b5050505092915050565b6000601f820183136105a257600080fd5b81356105b0610544826107c1565b915080825260208301602083018583830111156105cc57600080fd5b6105d783828461080c565b50505092915050565b6000602082840312156105f257600080fd5b60006105fe8484610512565b949350505050565b6000806040838503121561061957600080fd5b823567ffffffffffffffff81111561063057600080fd5b61063c85828601610525565b925050602083013567ffffffffffffffff81111561065957600080fd5b61066585828601610591565b9150509250929050565b610678816107f3565b82525050565b6000610689826107ef565b808452602084019350836020820285016106a2856107e9565b60005b848110156106d95783830388526106bd8383516106e5565b92506106c8826107e9565b6020989098019791506001016106a5565b50909695505050505050565b60006106f0826107ef565b808452610704816020860160208601610818565b61070d81610848565b9093016020019392505050565b610678816104ae565b60808101610731828761066f565b61073e602083018661071a565b818103604083015261075081856106e5565b905061075f606083018461071a565b95945050505050565b6020808252810161051e818461067e565b60405181810167ffffffffffffffff8111828210171561079857600080fd5b604052919050565b600067ffffffffffffffff8211156107b757600080fd5b5060209081020190565b600067ffffffffffffffff8211156107d857600080fd5b506020601f91909101601f19160190565b60200190565b5190565b73ffffffffffffffffffffffffffffffffffffffff1690565b82818337506000910152565b60005b8381101561083357818101518382015260200161081b565b83811115610842576000848401525b50505050565b601f01601f1916905600a265627a7a723058200c5fb3ae10421352e8eb8779739d8d5de18a160b7d991b61cb66d2fc0bae57b76c6578706572696d656e74616cf50037`

// DeployNodeValidations deploys a new Ethereum contract, binding an instance of NodeValidations to it.
func DeployNodeValidations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeValidations, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeValidationsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeValidationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeValidations{NodeValidationsCaller: NodeValidationsCaller{contract: contract}, NodeValidationsTransactor: NodeValidationsTransactor{contract: contract}, NodeValidationsFilterer: NodeValidationsFilterer{contract: contract}}, nil
}

// NodeValidations is an auto generated Go binding around an Ethereum contract.
type NodeValidations struct {
	NodeValidationsCaller     // Read-only binding to the contract
	NodeValidationsTransactor // Write-only binding to the contract
	NodeValidationsFilterer   // Log filterer for contract events
}

// NodeValidationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeValidationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeValidationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeValidationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeValidationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeValidationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeValidationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeValidationsSession struct {
	Contract     *NodeValidations  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeValidationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeValidationsCallerSession struct {
	Contract *NodeValidationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NodeValidationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeValidationsTransactorSession struct {
	Contract     *NodeValidationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NodeValidationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeValidationsRaw struct {
	Contract *NodeValidations // Generic contract binding to access the raw methods on
}

// NodeValidationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeValidationsCallerRaw struct {
	Contract *NodeValidationsCaller // Generic read-only contract binding to access the raw methods on
}

// NodeValidationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeValidationsTransactorRaw struct {
	Contract *NodeValidationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeValidations creates a new instance of NodeValidations, bound to a specific deployed contract.
func NewNodeValidations(address common.Address, backend bind.ContractBackend) (*NodeValidations, error) {
	contract, err := bindNodeValidations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeValidations{NodeValidationsCaller: NodeValidationsCaller{contract: contract}, NodeValidationsTransactor: NodeValidationsTransactor{contract: contract}, NodeValidationsFilterer: NodeValidationsFilterer{contract: contract}}, nil
}

// NewNodeValidationsCaller creates a new read-only instance of NodeValidations, bound to a specific deployed contract.
func NewNodeValidationsCaller(address common.Address, caller bind.ContractCaller) (*NodeValidationsCaller, error) {
	contract, err := bindNodeValidations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeValidationsCaller{contract: contract}, nil
}

// NewNodeValidationsTransactor creates a new write-only instance of NodeValidations, bound to a specific deployed contract.
func NewNodeValidationsTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeValidationsTransactor, error) {
	contract, err := bindNodeValidations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeValidationsTransactor{contract: contract}, nil
}

// NewNodeValidationsFilterer creates a new log filterer instance of NodeValidations, bound to a specific deployed contract.
func NewNodeValidationsFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeValidationsFilterer, error) {
	contract, err := bindNodeValidations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeValidationsFilterer{contract: contract}, nil
}

// bindNodeValidations binds a generic wrapper to an already deployed contract.
func bindNodeValidations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeValidationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeValidations *NodeValidationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeValidations.Contract.NodeValidationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeValidations *NodeValidationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeValidations.Contract.NodeValidationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeValidations *NodeValidationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeValidations.Contract.NodeValidationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeValidations *NodeValidationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeValidations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeValidations *NodeValidationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeValidations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeValidations *NodeValidationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeValidations.Contract.contract.Transact(opts, method, params...)
}

// GetSignatures is a free data retrieval call binding the contract method 0xed5b5015.
//
// Solidity: function getSignatures(address nodeAddress) constant returns(bytes[])
func (_NodeValidations *NodeValidationsCaller) GetSignatures(opts *bind.CallOpts, nodeAddress common.Address) ([][]byte, error) {
	var (
		ret0 = new([][]byte)
	)
	out := ret0
	err := _NodeValidations.contract.Call(opts, out, "getSignatures", nodeAddress)
	return *ret0, err
}

// GetSignatures is a free data retrieval call binding the contract method 0xed5b5015.
//
// Solidity: function getSignatures(address nodeAddress) constant returns(bytes[])
func (_NodeValidations *NodeValidationsSession) GetSignatures(nodeAddress common.Address) ([][]byte, error) {
	return _NodeValidations.Contract.GetSignatures(&_NodeValidations.CallOpts, nodeAddress)
}

// GetSignatures is a free data retrieval call binding the contract method 0xed5b5015.
//
// Solidity: function getSignatures(address nodeAddress) constant returns(bytes[])
func (_NodeValidations *NodeValidationsCallerSession) GetSignatures(nodeAddress common.Address) ([][]byte, error) {
	return _NodeValidations.Contract.GetSignatures(&_NodeValidations.CallOpts, nodeAddress)
}

// LastNodeActivity is a free data retrieval call binding the contract method 0xc7cfc9ea.
//
// Solidity: function lastNodeActivity(address ) constant returns(address nodeAddress, uint256 signatureCount, bytes publicKey, uint256 blockHeight)
func (_NodeValidations *NodeValidationsCaller) LastNodeActivity(opts *bind.CallOpts, arg0 common.Address) (struct {
	NodeAddress    common.Address
	SignatureCount *big.Int
	PublicKey      []byte
	BlockHeight    *big.Int
}, error) {
	ret := new(struct {
		NodeAddress    common.Address
		SignatureCount *big.Int
		PublicKey      []byte
		BlockHeight    *big.Int
	})
	out := ret
	err := _NodeValidations.contract.Call(opts, out, "lastNodeActivity", arg0)
	return *ret, err
}

// LastNodeActivity is a free data retrieval call binding the contract method 0xc7cfc9ea.
//
// Solidity: function lastNodeActivity(address ) constant returns(address nodeAddress, uint256 signatureCount, bytes publicKey, uint256 blockHeight)
func (_NodeValidations *NodeValidationsSession) LastNodeActivity(arg0 common.Address) (struct {
	NodeAddress    common.Address
	SignatureCount *big.Int
	PublicKey      []byte
	BlockHeight    *big.Int
}, error) {
	return _NodeValidations.Contract.LastNodeActivity(&_NodeValidations.CallOpts, arg0)
}

// LastNodeActivity is a free data retrieval call binding the contract method 0xc7cfc9ea.
//
// Solidity: function lastNodeActivity(address ) constant returns(address nodeAddress, uint256 signatureCount, bytes publicKey, uint256 blockHeight)
func (_NodeValidations *NodeValidationsCallerSession) LastNodeActivity(arg0 common.Address) (struct {
	NodeAddress    common.Address
	SignatureCount *big.Int
	PublicKey      []byte
	BlockHeight    *big.Int
}, error) {
	return _NodeValidations.Contract.LastNodeActivity(&_NodeValidations.CallOpts, arg0)
}

// NodeCheckIn is a paid mutator transaction binding the contract method 0xc511613b.
//
// Solidity: function nodeCheckIn(bytes[] signatures, bytes publicKey) returns()
func (_NodeValidations *NodeValidationsTransactor) NodeCheckIn(opts *bind.TransactOpts, signatures [][]byte, publicKey []byte) (*types.Transaction, error) {
	return _NodeValidations.contract.Transact(opts, "nodeCheckIn", signatures, publicKey)
}

// NodeCheckIn is a paid mutator transaction binding the contract method 0xc511613b.
//
// Solidity: function nodeCheckIn(bytes[] signatures, bytes publicKey) returns()
func (_NodeValidations *NodeValidationsSession) NodeCheckIn(signatures [][]byte, publicKey []byte) (*types.Transaction, error) {
	return _NodeValidations.Contract.NodeCheckIn(&_NodeValidations.TransactOpts, signatures, publicKey)
}

// NodeCheckIn is a paid mutator transaction binding the contract method 0xc511613b.
//
// Solidity: function nodeCheckIn(bytes[] signatures, bytes publicKey) returns()
func (_NodeValidations *NodeValidationsTransactorSession) NodeCheckIn(signatures [][]byte, publicKey []byte) (*types.Transaction, error) {
	return _NodeValidations.Contract.NodeCheckIn(&_NodeValidations.TransactOpts, signatures, publicKey)
}
