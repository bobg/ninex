// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ninex

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// NinexABI is the input ABI used to generate the binding from.
const NinexABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"mPayments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mRevealedTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mCommitmentSetTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mGuessedDigits\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setRevealTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\"},{\"name\":\"digits\",\"type\":\"bytes\"}],\"name\":\"guess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mRevealTimeoutSecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mCommitment\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAfterGuessDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mGuessedTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mBank\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mMinGuess\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAfterRevealDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMinGuess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAfterCommitmentDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mAfterGuessDelaySecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"setCommitment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mGuessedBy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mAfterCommitmentDelaySecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mAfterRevealDelaySecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mEscrow\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NinexBin is the compiled bytecode used for deploying new contracts.
const NinexBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161781556102586009819055600a819055600b5562015180600c55620f4240600d556002819055600381905560078190556008819055600655610baa806100706000396000f30060606040526004361061015e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630342a929811461016357806309ecabf5146101945780630f6f2503146101a75780632c46c8f9146101ba5780632e1a7d4d1461024457806338b640851461025c5780633d58f99b146102725780633ef314aa146102895780634d4f604a1461029c578063503e114b146102af57806351091b33146102c5578063592dac2a146102d857806361a14c90146102eb5780636a2403b8146102fe5780636d5cfa581461031457806370dea79a1461032a57806372f12a5d1461033d5780639594c3dc1461035b578063a61ddc1714610371578063abe2003214610384578063b0b2d5d91461039a578063b60d4288146103c9578063c571598f146103d1578063d7b5412d146103e4578063dcda4bf3146103f7578063dfbec5901461040a578063e52253811461041d575b600080fd5b341561016e57600080fd5b610182600160a060020a0360043516610430565b60405190815260200160405180910390f35b341561019f57600080fd5b610182610442565b34156101b257600080fd5b610182610448565b34156101c557600080fd5b6101cd61044e565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102095780820151838201526020016101f1565b50505050905090810190601f1680156102365780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561024f57600080fd5b61025a6004356104ec565b005b341561026757600080fd5b61025a600435610555565b61025a600480359060248035908101910135610582565b341561029457600080fd5b6101826106d7565b34156102a757600080fd5b6101826106dd565b34156102ba57600080fd5b61025a6004356106e3565b34156102d057600080fd5b610182610710565b34156102e357600080fd5b610182610716565b34156102f657600080fd5b61018261071c565b341561030957600080fd5b61025a600435610722565b341561031f57600080fd5b61025a60043561074f565b341561033557600080fd5b61025a61077c565b341561034857600080fd5b61025a60048035602481019101356107ce565b341561036657600080fd5b61025a6004356109b6565b341561037c57600080fd5b6101826109e3565b341561038f57600080fd5b61025a6004356109e9565b34156103a557600080fd5b6103ad610a44565b604051600160a060020a03909116815260200160405180910390f35b61025a610a53565b34156103dc57600080fd5b610182610a5d565b34156103ef57600080fd5b610182610a63565b341561040257600080fd5b6103ad610a69565b341561041557600080fd5b610182610a78565b341561042857600080fd5b61025a610a7e565b600e6020526000908152604090205481565b60085481565b60065481565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104e45780601f106104b9576101008083540402835291602001916104e4565b820191906000526020600020905b8154815290600101906020018083116104c757829003601f168201915b505050505081565b60005433600160a060020a0390811691161461050757600080fd5b60025481111561051657600080fd5b600280548290039055600054600160a060020a031681156108fc0282604051600060405180830381858888f19350505050151561055257600080fd5b50565b60005433600160a060020a0390811691161461057057600080fd5b6003541561057d57600080fd5b600c55565b6000806000600354600014151561059857600080fd5b600654600090116105a857600080fd5b60018410156105b657600080fd5b604e8411156105c457600080fd5b600d543410156105d357600080fd5b600254346000198601600a0a81026009020393508311156105f357600080fd5b6009546006540142101561060657600080fd5b506001905060005b602081101561066d57600154816020811061062557fe5b1a60f860020a02600160f860020a031916868260208110151561064457fe5b1a60f860020a02600160f860020a031916141515610665576000915061066d565b60010161060e565b81151561067957600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03161790556106ad60058686610ae3565b50506002805483900390555060038054909101340190555050426007555060006006819055600855565b600c5481565b60015481565b60005433600160a060020a039081169116146106fe57600080fd5b6003541561070b57600080fd5b600a55565b60075481565b60025481565b600d5481565b60005433600160a060020a0390811691161461073d57600080fd5b6003541561074a57600080fd5b600b55565b60005433600160a060020a0390811691161461076a57600080fd5b6003541561077757600080fd5b600d55565b6003546000901161078c57600080fd5b600c546007540142101561079f57600080fd5b60038054600454600160a060020a03166000908152600e60205260408120805490920190915590819055600755565b600080548190819033600160a060020a039081169116146107ee57600080fd5b604e84146107fb57600080fd5b600154600286866000604051602001526040518083838082843782019150509250505060206040518083038160008661646e5a03f1151561083b57600080fd5b5050604051805191909114905061085157600080fd5b600060035411156109a157600a546007540142101561086f57600080fd5b600c5460075401421061088157600080fd5b505060055460019150600260001961010083851615020190911604604e03805b604e81101561095f578484828181106108b657fe5b905090013560f860020a900460f860020a02600160f860020a031916600160f860020a031916600583830381546001816001161561010002031660029004811015156108fe57fe5b81546001161561091d5790600052602060002090602091828204019190065b9054901a60f860020a027fff000000000000000000000000000000000000000000000000000000000000001614610957576000925061095f565b6001016108a1565b821561098e57600354600454600160a060020a03166000908152600e602052604090208054909101905561099b565b6003546002805490910190555b60006003555b50504260085550506000600681905560075550565b60005433600160a060020a039081169116146109d157600080fd5b600354156109de57600080fd5b600955565b600a5481565b60005433600160a060020a03908116911614610a0457600080fd5b60035415610a1157600080fd5b60065415610a1e57600080fd5b600b5460085401421015610a3157600080fd5b6001554260065560006008819055600755565b600454600160a060020a031681565b6002805434019055565b60095481565b600b5481565b600054600160a060020a031681565b60035481565b600160a060020a0333166000908152600e60205260408120549081111561055257600160a060020a0333166000818152600e60205260408082209190915582156108fc0290839051600060405180830381858888f19350505050151561055257600080fd5b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b245782800160ff19823516178555610b51565b82800160010185558215610b51579182015b82811115610b51578235825591602001919060010190610b36565b50610b5d929150610b61565b5090565b610b7b91905b80821115610b5d5760008155600101610b67565b905600a165627a7a72305820afbc354b109d7faef1842c0732745d512ad52bd75b00a901fdf0969829595c610029`

// DeployNinex deploys a new Ethereum contract, binding an instance of Ninex to it.
func DeployNinex(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ninex, error) {
	parsed, err := abi.JSON(strings.NewReader(NinexABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NinexBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ninex{NinexCaller: NinexCaller{contract: contract}, NinexTransactor: NinexTransactor{contract: contract}}, nil
}

// Ninex is an auto generated Go binding around an Ethereum contract.
type Ninex struct {
	NinexCaller     // Read-only binding to the contract
	NinexTransactor // Write-only binding to the contract
}

// NinexCaller is an auto generated read-only Go binding around an Ethereum contract.
type NinexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NinexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NinexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NinexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NinexSession struct {
	Contract     *Ninex            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NinexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NinexCallerSession struct {
	Contract *NinexCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NinexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NinexTransactorSession struct {
	Contract     *NinexTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NinexRaw is an auto generated low-level Go binding around an Ethereum contract.
type NinexRaw struct {
	Contract *Ninex // Generic contract binding to access the raw methods on
}

// NinexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NinexCallerRaw struct {
	Contract *NinexCaller // Generic read-only contract binding to access the raw methods on
}

// NinexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NinexTransactorRaw struct {
	Contract *NinexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNinex creates a new instance of Ninex, bound to a specific deployed contract.
func NewNinex(address common.Address, backend bind.ContractBackend) (*Ninex, error) {
	contract, err := bindNinex(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ninex{NinexCaller: NinexCaller{contract: contract}, NinexTransactor: NinexTransactor{contract: contract}}, nil
}

// NewNinexCaller creates a new read-only instance of Ninex, bound to a specific deployed contract.
func NewNinexCaller(address common.Address, caller bind.ContractCaller) (*NinexCaller, error) {
	contract, err := bindNinex(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &NinexCaller{contract: contract}, nil
}

// NewNinexTransactor creates a new write-only instance of Ninex, bound to a specific deployed contract.
func NewNinexTransactor(address common.Address, transactor bind.ContractTransactor) (*NinexTransactor, error) {
	contract, err := bindNinex(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &NinexTransactor{contract: contract}, nil
}

// bindNinex binds a generic wrapper to an already deployed contract.
func bindNinex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NinexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ninex *NinexRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ninex.Contract.NinexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ninex *NinexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ninex.Contract.NinexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ninex *NinexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ninex.Contract.NinexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ninex *NinexCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ninex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ninex *NinexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ninex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ninex *NinexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ninex.Contract.contract.Transact(opts, method, params...)
}

// MAfterCommitmentDelaySecs is a free data retrieval call binding the contract method 0xc571598f.
//
// Solidity: function mAfterCommitmentDelaySecs() constant returns(uint256)
func (_Ninex *NinexCaller) MAfterCommitmentDelaySecs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mAfterCommitmentDelaySecs")
	return *ret0, err
}

// MAfterCommitmentDelaySecs is a free data retrieval call binding the contract method 0xc571598f.
//
// Solidity: function mAfterCommitmentDelaySecs() constant returns(uint256)
func (_Ninex *NinexSession) MAfterCommitmentDelaySecs() (*big.Int, error) {
	return _Ninex.Contract.MAfterCommitmentDelaySecs(&_Ninex.CallOpts)
}

// MAfterCommitmentDelaySecs is a free data retrieval call binding the contract method 0xc571598f.
//
// Solidity: function mAfterCommitmentDelaySecs() constant returns(uint256)
func (_Ninex *NinexCallerSession) MAfterCommitmentDelaySecs() (*big.Int, error) {
	return _Ninex.Contract.MAfterCommitmentDelaySecs(&_Ninex.CallOpts)
}

// MAfterGuessDelaySecs is a free data retrieval call binding the contract method 0xa61ddc17.
//
// Solidity: function mAfterGuessDelaySecs() constant returns(uint256)
func (_Ninex *NinexCaller) MAfterGuessDelaySecs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mAfterGuessDelaySecs")
	return *ret0, err
}

// MAfterGuessDelaySecs is a free data retrieval call binding the contract method 0xa61ddc17.
//
// Solidity: function mAfterGuessDelaySecs() constant returns(uint256)
func (_Ninex *NinexSession) MAfterGuessDelaySecs() (*big.Int, error) {
	return _Ninex.Contract.MAfterGuessDelaySecs(&_Ninex.CallOpts)
}

// MAfterGuessDelaySecs is a free data retrieval call binding the contract method 0xa61ddc17.
//
// Solidity: function mAfterGuessDelaySecs() constant returns(uint256)
func (_Ninex *NinexCallerSession) MAfterGuessDelaySecs() (*big.Int, error) {
	return _Ninex.Contract.MAfterGuessDelaySecs(&_Ninex.CallOpts)
}

// MAfterRevealDelaySecs is a free data retrieval call binding the contract method 0xd7b5412d.
//
// Solidity: function mAfterRevealDelaySecs() constant returns(uint256)
func (_Ninex *NinexCaller) MAfterRevealDelaySecs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mAfterRevealDelaySecs")
	return *ret0, err
}

// MAfterRevealDelaySecs is a free data retrieval call binding the contract method 0xd7b5412d.
//
// Solidity: function mAfterRevealDelaySecs() constant returns(uint256)
func (_Ninex *NinexSession) MAfterRevealDelaySecs() (*big.Int, error) {
	return _Ninex.Contract.MAfterRevealDelaySecs(&_Ninex.CallOpts)
}

// MAfterRevealDelaySecs is a free data retrieval call binding the contract method 0xd7b5412d.
//
// Solidity: function mAfterRevealDelaySecs() constant returns(uint256)
func (_Ninex *NinexCallerSession) MAfterRevealDelaySecs() (*big.Int, error) {
	return _Ninex.Contract.MAfterRevealDelaySecs(&_Ninex.CallOpts)
}

// MBank is a free data retrieval call binding the contract method 0x592dac2a.
//
// Solidity: function mBank() constant returns(uint256)
func (_Ninex *NinexCaller) MBank(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mBank")
	return *ret0, err
}

// MBank is a free data retrieval call binding the contract method 0x592dac2a.
//
// Solidity: function mBank() constant returns(uint256)
func (_Ninex *NinexSession) MBank() (*big.Int, error) {
	return _Ninex.Contract.MBank(&_Ninex.CallOpts)
}

// MBank is a free data retrieval call binding the contract method 0x592dac2a.
//
// Solidity: function mBank() constant returns(uint256)
func (_Ninex *NinexCallerSession) MBank() (*big.Int, error) {
	return _Ninex.Contract.MBank(&_Ninex.CallOpts)
}

// MCommitment is a free data retrieval call binding the contract method 0x4d4f604a.
//
// Solidity: function mCommitment() constant returns(bytes32)
func (_Ninex *NinexCaller) MCommitment(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mCommitment")
	return *ret0, err
}

// MCommitment is a free data retrieval call binding the contract method 0x4d4f604a.
//
// Solidity: function mCommitment() constant returns(bytes32)
func (_Ninex *NinexSession) MCommitment() ([32]byte, error) {
	return _Ninex.Contract.MCommitment(&_Ninex.CallOpts)
}

// MCommitment is a free data retrieval call binding the contract method 0x4d4f604a.
//
// Solidity: function mCommitment() constant returns(bytes32)
func (_Ninex *NinexCallerSession) MCommitment() ([32]byte, error) {
	return _Ninex.Contract.MCommitment(&_Ninex.CallOpts)
}

// MCommitmentSetTime is a free data retrieval call binding the contract method 0x0f6f2503.
//
// Solidity: function mCommitmentSetTime() constant returns(uint256)
func (_Ninex *NinexCaller) MCommitmentSetTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mCommitmentSetTime")
	return *ret0, err
}

// MCommitmentSetTime is a free data retrieval call binding the contract method 0x0f6f2503.
//
// Solidity: function mCommitmentSetTime() constant returns(uint256)
func (_Ninex *NinexSession) MCommitmentSetTime() (*big.Int, error) {
	return _Ninex.Contract.MCommitmentSetTime(&_Ninex.CallOpts)
}

// MCommitmentSetTime is a free data retrieval call binding the contract method 0x0f6f2503.
//
// Solidity: function mCommitmentSetTime() constant returns(uint256)
func (_Ninex *NinexCallerSession) MCommitmentSetTime() (*big.Int, error) {
	return _Ninex.Contract.MCommitmentSetTime(&_Ninex.CallOpts)
}

// MEscrow is a free data retrieval call binding the contract method 0xdfbec590.
//
// Solidity: function mEscrow() constant returns(uint256)
func (_Ninex *NinexCaller) MEscrow(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mEscrow")
	return *ret0, err
}

// MEscrow is a free data retrieval call binding the contract method 0xdfbec590.
//
// Solidity: function mEscrow() constant returns(uint256)
func (_Ninex *NinexSession) MEscrow() (*big.Int, error) {
	return _Ninex.Contract.MEscrow(&_Ninex.CallOpts)
}

// MEscrow is a free data retrieval call binding the contract method 0xdfbec590.
//
// Solidity: function mEscrow() constant returns(uint256)
func (_Ninex *NinexCallerSession) MEscrow() (*big.Int, error) {
	return _Ninex.Contract.MEscrow(&_Ninex.CallOpts)
}

// MGuessedBy is a free data retrieval call binding the contract method 0xb0b2d5d9.
//
// Solidity: function mGuessedBy() constant returns(address)
func (_Ninex *NinexCaller) MGuessedBy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mGuessedBy")
	return *ret0, err
}

// MGuessedBy is a free data retrieval call binding the contract method 0xb0b2d5d9.
//
// Solidity: function mGuessedBy() constant returns(address)
func (_Ninex *NinexSession) MGuessedBy() (common.Address, error) {
	return _Ninex.Contract.MGuessedBy(&_Ninex.CallOpts)
}

// MGuessedBy is a free data retrieval call binding the contract method 0xb0b2d5d9.
//
// Solidity: function mGuessedBy() constant returns(address)
func (_Ninex *NinexCallerSession) MGuessedBy() (common.Address, error) {
	return _Ninex.Contract.MGuessedBy(&_Ninex.CallOpts)
}

// MGuessedDigits is a free data retrieval call binding the contract method 0x2c46c8f9.
//
// Solidity: function mGuessedDigits() constant returns(bytes)
func (_Ninex *NinexCaller) MGuessedDigits(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mGuessedDigits")
	return *ret0, err
}

// MGuessedDigits is a free data retrieval call binding the contract method 0x2c46c8f9.
//
// Solidity: function mGuessedDigits() constant returns(bytes)
func (_Ninex *NinexSession) MGuessedDigits() ([]byte, error) {
	return _Ninex.Contract.MGuessedDigits(&_Ninex.CallOpts)
}

// MGuessedDigits is a free data retrieval call binding the contract method 0x2c46c8f9.
//
// Solidity: function mGuessedDigits() constant returns(bytes)
func (_Ninex *NinexCallerSession) MGuessedDigits() ([]byte, error) {
	return _Ninex.Contract.MGuessedDigits(&_Ninex.CallOpts)
}

// MGuessedTime is a free data retrieval call binding the contract method 0x51091b33.
//
// Solidity: function mGuessedTime() constant returns(uint256)
func (_Ninex *NinexCaller) MGuessedTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mGuessedTime")
	return *ret0, err
}

// MGuessedTime is a free data retrieval call binding the contract method 0x51091b33.
//
// Solidity: function mGuessedTime() constant returns(uint256)
func (_Ninex *NinexSession) MGuessedTime() (*big.Int, error) {
	return _Ninex.Contract.MGuessedTime(&_Ninex.CallOpts)
}

// MGuessedTime is a free data retrieval call binding the contract method 0x51091b33.
//
// Solidity: function mGuessedTime() constant returns(uint256)
func (_Ninex *NinexCallerSession) MGuessedTime() (*big.Int, error) {
	return _Ninex.Contract.MGuessedTime(&_Ninex.CallOpts)
}

// MMinGuess is a free data retrieval call binding the contract method 0x61a14c90.
//
// Solidity: function mMinGuess() constant returns(uint256)
func (_Ninex *NinexCaller) MMinGuess(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mMinGuess")
	return *ret0, err
}

// MMinGuess is a free data retrieval call binding the contract method 0x61a14c90.
//
// Solidity: function mMinGuess() constant returns(uint256)
func (_Ninex *NinexSession) MMinGuess() (*big.Int, error) {
	return _Ninex.Contract.MMinGuess(&_Ninex.CallOpts)
}

// MMinGuess is a free data retrieval call binding the contract method 0x61a14c90.
//
// Solidity: function mMinGuess() constant returns(uint256)
func (_Ninex *NinexCallerSession) MMinGuess() (*big.Int, error) {
	return _Ninex.Contract.MMinGuess(&_Ninex.CallOpts)
}

// MOwner is a free data retrieval call binding the contract method 0xdcda4bf3.
//
// Solidity: function mOwner() constant returns(address)
func (_Ninex *NinexCaller) MOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mOwner")
	return *ret0, err
}

// MOwner is a free data retrieval call binding the contract method 0xdcda4bf3.
//
// Solidity: function mOwner() constant returns(address)
func (_Ninex *NinexSession) MOwner() (common.Address, error) {
	return _Ninex.Contract.MOwner(&_Ninex.CallOpts)
}

// MOwner is a free data retrieval call binding the contract method 0xdcda4bf3.
//
// Solidity: function mOwner() constant returns(address)
func (_Ninex *NinexCallerSession) MOwner() (common.Address, error) {
	return _Ninex.Contract.MOwner(&_Ninex.CallOpts)
}

// MPayments is a free data retrieval call binding the contract method 0x0342a929.
//
// Solidity: function mPayments( address) constant returns(uint256)
func (_Ninex *NinexCaller) MPayments(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mPayments", arg0)
	return *ret0, err
}

// MPayments is a free data retrieval call binding the contract method 0x0342a929.
//
// Solidity: function mPayments( address) constant returns(uint256)
func (_Ninex *NinexSession) MPayments(arg0 common.Address) (*big.Int, error) {
	return _Ninex.Contract.MPayments(&_Ninex.CallOpts, arg0)
}

// MPayments is a free data retrieval call binding the contract method 0x0342a929.
//
// Solidity: function mPayments( address) constant returns(uint256)
func (_Ninex *NinexCallerSession) MPayments(arg0 common.Address) (*big.Int, error) {
	return _Ninex.Contract.MPayments(&_Ninex.CallOpts, arg0)
}

// MRevealTimeoutSecs is a free data retrieval call binding the contract method 0x3ef314aa.
//
// Solidity: function mRevealTimeoutSecs() constant returns(uint256)
func (_Ninex *NinexCaller) MRevealTimeoutSecs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mRevealTimeoutSecs")
	return *ret0, err
}

// MRevealTimeoutSecs is a free data retrieval call binding the contract method 0x3ef314aa.
//
// Solidity: function mRevealTimeoutSecs() constant returns(uint256)
func (_Ninex *NinexSession) MRevealTimeoutSecs() (*big.Int, error) {
	return _Ninex.Contract.MRevealTimeoutSecs(&_Ninex.CallOpts)
}

// MRevealTimeoutSecs is a free data retrieval call binding the contract method 0x3ef314aa.
//
// Solidity: function mRevealTimeoutSecs() constant returns(uint256)
func (_Ninex *NinexCallerSession) MRevealTimeoutSecs() (*big.Int, error) {
	return _Ninex.Contract.MRevealTimeoutSecs(&_Ninex.CallOpts)
}

// MRevealedTime is a free data retrieval call binding the contract method 0x09ecabf5.
//
// Solidity: function mRevealedTime() constant returns(uint256)
func (_Ninex *NinexCaller) MRevealedTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mRevealedTime")
	return *ret0, err
}

// MRevealedTime is a free data retrieval call binding the contract method 0x09ecabf5.
//
// Solidity: function mRevealedTime() constant returns(uint256)
func (_Ninex *NinexSession) MRevealedTime() (*big.Int, error) {
	return _Ninex.Contract.MRevealedTime(&_Ninex.CallOpts)
}

// MRevealedTime is a free data retrieval call binding the contract method 0x09ecabf5.
//
// Solidity: function mRevealedTime() constant returns(uint256)
func (_Ninex *NinexCallerSession) MRevealedTime() (*big.Int, error) {
	return _Ninex.Contract.MRevealedTime(&_Ninex.CallOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_Ninex *NinexTransactor) Collect(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "collect")
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_Ninex *NinexSession) Collect() (*types.Transaction, error) {
	return _Ninex.Contract.Collect(&_Ninex.TransactOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_Ninex *NinexTransactorSession) Collect() (*types.Transaction, error) {
	return _Ninex.Contract.Collect(&_Ninex.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_Ninex *NinexTransactor) Fund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "fund")
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_Ninex *NinexSession) Fund() (*types.Transaction, error) {
	return _Ninex.Contract.Fund(&_Ninex.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_Ninex *NinexTransactorSession) Fund() (*types.Transaction, error) {
	return _Ninex.Contract.Fund(&_Ninex.TransactOpts)
}

// Guess is a paid mutator transaction binding the contract method 0x3d58f99b.
//
// Solidity: function guess(commitment bytes32, digits bytes) returns()
func (_Ninex *NinexTransactor) Guess(opts *bind.TransactOpts, commitment [32]byte, digits []byte) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "guess", commitment, digits)
}

// Guess is a paid mutator transaction binding the contract method 0x3d58f99b.
//
// Solidity: function guess(commitment bytes32, digits bytes) returns()
func (_Ninex *NinexSession) Guess(commitment [32]byte, digits []byte) (*types.Transaction, error) {
	return _Ninex.Contract.Guess(&_Ninex.TransactOpts, commitment, digits)
}

// Guess is a paid mutator transaction binding the contract method 0x3d58f99b.
//
// Solidity: function guess(commitment bytes32, digits bytes) returns()
func (_Ninex *NinexTransactorSession) Guess(commitment [32]byte, digits []byte) (*types.Transaction, error) {
	return _Ninex.Contract.Guess(&_Ninex.TransactOpts, commitment, digits)
}

// Reveal is a paid mutator transaction binding the contract method 0x72f12a5d.
//
// Solidity: function reveal(preimage bytes) returns()
func (_Ninex *NinexTransactor) Reveal(opts *bind.TransactOpts, preimage []byte) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "reveal", preimage)
}

// Reveal is a paid mutator transaction binding the contract method 0x72f12a5d.
//
// Solidity: function reveal(preimage bytes) returns()
func (_Ninex *NinexSession) Reveal(preimage []byte) (*types.Transaction, error) {
	return _Ninex.Contract.Reveal(&_Ninex.TransactOpts, preimage)
}

// Reveal is a paid mutator transaction binding the contract method 0x72f12a5d.
//
// Solidity: function reveal(preimage bytes) returns()
func (_Ninex *NinexTransactorSession) Reveal(preimage []byte) (*types.Transaction, error) {
	return _Ninex.Contract.Reveal(&_Ninex.TransactOpts, preimage)
}

// SetAfterCommitmentDelay is a paid mutator transaction binding the contract method 0x9594c3dc.
//
// Solidity: function setAfterCommitmentDelay(val uint256) returns()
func (_Ninex *NinexTransactor) SetAfterCommitmentDelay(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "setAfterCommitmentDelay", val)
}

// SetAfterCommitmentDelay is a paid mutator transaction binding the contract method 0x9594c3dc.
//
// Solidity: function setAfterCommitmentDelay(val uint256) returns()
func (_Ninex *NinexSession) SetAfterCommitmentDelay(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetAfterCommitmentDelay(&_Ninex.TransactOpts, val)
}

// SetAfterCommitmentDelay is a paid mutator transaction binding the contract method 0x9594c3dc.
//
// Solidity: function setAfterCommitmentDelay(val uint256) returns()
func (_Ninex *NinexTransactorSession) SetAfterCommitmentDelay(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetAfterCommitmentDelay(&_Ninex.TransactOpts, val)
}

// SetAfterGuessDelay is a paid mutator transaction binding the contract method 0x503e114b.
//
// Solidity: function setAfterGuessDelay(val uint256) returns()
func (_Ninex *NinexTransactor) SetAfterGuessDelay(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "setAfterGuessDelay", val)
}

// SetAfterGuessDelay is a paid mutator transaction binding the contract method 0x503e114b.
//
// Solidity: function setAfterGuessDelay(val uint256) returns()
func (_Ninex *NinexSession) SetAfterGuessDelay(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetAfterGuessDelay(&_Ninex.TransactOpts, val)
}

// SetAfterGuessDelay is a paid mutator transaction binding the contract method 0x503e114b.
//
// Solidity: function setAfterGuessDelay(val uint256) returns()
func (_Ninex *NinexTransactorSession) SetAfterGuessDelay(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetAfterGuessDelay(&_Ninex.TransactOpts, val)
}

// SetAfterRevealDelay is a paid mutator transaction binding the contract method 0x6a2403b8.
//
// Solidity: function setAfterRevealDelay(val uint256) returns()
func (_Ninex *NinexTransactor) SetAfterRevealDelay(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "setAfterRevealDelay", val)
}

// SetAfterRevealDelay is a paid mutator transaction binding the contract method 0x6a2403b8.
//
// Solidity: function setAfterRevealDelay(val uint256) returns()
func (_Ninex *NinexSession) SetAfterRevealDelay(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetAfterRevealDelay(&_Ninex.TransactOpts, val)
}

// SetAfterRevealDelay is a paid mutator transaction binding the contract method 0x6a2403b8.
//
// Solidity: function setAfterRevealDelay(val uint256) returns()
func (_Ninex *NinexTransactorSession) SetAfterRevealDelay(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetAfterRevealDelay(&_Ninex.TransactOpts, val)
}

// SetCommitment is a paid mutator transaction binding the contract method 0xabe20032.
//
// Solidity: function setCommitment(commitment bytes32) returns()
func (_Ninex *NinexTransactor) SetCommitment(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "setCommitment", commitment)
}

// SetCommitment is a paid mutator transaction binding the contract method 0xabe20032.
//
// Solidity: function setCommitment(commitment bytes32) returns()
func (_Ninex *NinexSession) SetCommitment(commitment [32]byte) (*types.Transaction, error) {
	return _Ninex.Contract.SetCommitment(&_Ninex.TransactOpts, commitment)
}

// SetCommitment is a paid mutator transaction binding the contract method 0xabe20032.
//
// Solidity: function setCommitment(commitment bytes32) returns()
func (_Ninex *NinexTransactorSession) SetCommitment(commitment [32]byte) (*types.Transaction, error) {
	return _Ninex.Contract.SetCommitment(&_Ninex.TransactOpts, commitment)
}

// SetMinGuess is a paid mutator transaction binding the contract method 0x6d5cfa58.
//
// Solidity: function setMinGuess(val uint256) returns()
func (_Ninex *NinexTransactor) SetMinGuess(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "setMinGuess", val)
}

// SetMinGuess is a paid mutator transaction binding the contract method 0x6d5cfa58.
//
// Solidity: function setMinGuess(val uint256) returns()
func (_Ninex *NinexSession) SetMinGuess(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetMinGuess(&_Ninex.TransactOpts, val)
}

// SetMinGuess is a paid mutator transaction binding the contract method 0x6d5cfa58.
//
// Solidity: function setMinGuess(val uint256) returns()
func (_Ninex *NinexTransactorSession) SetMinGuess(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetMinGuess(&_Ninex.TransactOpts, val)
}

// SetRevealTimeout is a paid mutator transaction binding the contract method 0x38b64085.
//
// Solidity: function setRevealTimeout(val uint256) returns()
func (_Ninex *NinexTransactor) SetRevealTimeout(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "setRevealTimeout", val)
}

// SetRevealTimeout is a paid mutator transaction binding the contract method 0x38b64085.
//
// Solidity: function setRevealTimeout(val uint256) returns()
func (_Ninex *NinexSession) SetRevealTimeout(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetRevealTimeout(&_Ninex.TransactOpts, val)
}

// SetRevealTimeout is a paid mutator transaction binding the contract method 0x38b64085.
//
// Solidity: function setRevealTimeout(val uint256) returns()
func (_Ninex *NinexTransactorSession) SetRevealTimeout(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetRevealTimeout(&_Ninex.TransactOpts, val)
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_Ninex *NinexTransactor) Timeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "timeout")
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_Ninex *NinexSession) Timeout() (*types.Transaction, error) {
	return _Ninex.Contract.Timeout(&_Ninex.TransactOpts)
}

// Timeout is a paid mutator transaction binding the contract method 0x70dea79a.
//
// Solidity: function timeout() returns()
func (_Ninex *NinexTransactorSession) Timeout() (*types.Transaction, error) {
	return _Ninex.Contract.Timeout(&_Ninex.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_Ninex *NinexTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_Ninex *NinexSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.Withdraw(&_Ninex.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_Ninex *NinexTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.Withdraw(&_Ninex.TransactOpts, amount)
}
