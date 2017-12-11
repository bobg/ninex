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
const NinexABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"mPayments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mRevealedTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mCommitmentSetTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMinGuessWei\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mGuessWindowSecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setRevealTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\"},{\"name\":\"digits\",\"type\":\"bytes\"}],\"name\":\"guess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mRevealTimeoutSecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mAfterLastGuessDelaySecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mCommitment\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mFirstGuessTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mBank\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setGuessWindow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAfterRevealDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAfterCommitmentDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"setCommitment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mMinGuessWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mAfterCommitmentDelaySecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAfterLastGuessDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numGuesses\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mAfterRevealDelaySecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mLastGuessTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mGuesses\",\"outputs\":[{\"name\":\"guesser\",\"type\":\"address\"},{\"name\":\"digits\",\"type\":\"bytes\"},{\"name\":\"escrow\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guesser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"digits\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"evGuess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"evNewCommitment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guesser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"digits\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"evWin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guesser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"digits\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"evLose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"evRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guesser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"evWinByDefault\",\"type\":\"event\"}]"

// NinexBin is the compiled bytecode used for deploying new contracts.
const NinexBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161781556102586008819055610384600955600a819055600b5562015180600c55620f4240600d5560028190556004819055600581905560068190556007556112a2806100766000396000f3006060604052600436106101745763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630342a929811461017957806309ecabf5146101aa5780630f6f2503146101bd578063196161fa146101d057806328453d6f146101e85780632e1a7d4d146101fb57806338b64085146102115780633d58f99b146102275780633ef314aa1461023e5780634ae800f0146102515780634d4f604a14610264578063525b6b7c14610277578063592dac2a1461028a578063663ae0281461029d5780636a2403b8146102b357806370dea79a146102c957806372f12a5d146102dc5780639594c3dc146102fa578063abe2003214610310578063b60d428814610326578063b8df37eb1461032e578063c571598f14610341578063cb18c64314610354578063d659d9a51461036a578063d7b5412d1461037d578063dcda4bf314610390578063e5225381146103bf578063e66c3791146103d2578063f33fffd8146103e5575b600080fd5b341561018457600080fd5b610198600160a060020a036004351661049b565b60405190815260200160405180910390f35b34156101b557600080fd5b6101986104ad565b34156101c857600080fd5b6101986104b3565b34156101db57600080fd5b6101e66004356104b9565b005b34156101f357600080fd5b6101986104e6565b341561020657600080fd5b6101e66004356104ec565b341561021c57600080fd5b6101e6600435610555565b6101e6600480359060248035908101910135610582565b341561024957600080fd5b6101986107f3565b341561025c57600080fd5b6101986107f9565b341561026f57600080fd5b6101986107ff565b341561028257600080fd5b610198610805565b341561029557600080fd5b61019861080b565b34156102a857600080fd5b6101e6600435610811565b34156102be57600080fd5b6101e660043561083e565b34156102d457600080fd5b6101e661086b565b34156102e757600080fd5b6101e660048035602481019101356109d8565b341561030557600080fd5b6101e6600435610f3d565b341561031b57600080fd5b6101e6600435610f6a565b6101e661101c565b341561033957600080fd5b610198611026565b341561034c57600080fd5b61019861102c565b341561035f57600080fd5b6101e6600435611032565b341561037557600080fd5b61019861105f565b341561038857600080fd5b610198611066565b341561039b57600080fd5b6103a361106c565b604051600160a060020a03909116815260200160405180910390f35b34156103ca57600080fd5b6101e661107b565b34156103dd57600080fd5b6101986110e0565b34156103f057600080fd5b6103fb6004356110e6565b604051600160a060020a0384168152604081018290526060602082018181528454600260001961010060018416150201909116049183018290529060808301908590801561048a5780601f1061045f5761010080835404028352916020019161048a565b820191906000526020600020905b81548152906001019060200180831161046d57829003601f168201915b505094505050505060405180910390f35b600e6020526000908152604090205481565b60075481565b60045481565b60005433600160a060020a039081169116146104d457600080fd5b600354156104e157600080fd5b600d55565b60095481565b60005433600160a060020a0390811691161461050757600080fd5b60025481111561051657600080fd5b600280548290039055600054600160a060020a031681156108fc0282604051600060405180830381858888f19350505050151561055257600080fd5b50565b60005433600160a060020a0390811691161461057057600080fd5b6003541561057d57600080fd5b600c55565b60008060008060045411151561059757600080fd5b600754156105a457600080fd5b600854600454014210156105b757600080fd5b60055415806105cb57506009546005540142105b15156105d657600080fd5b60018410156105e457600080fd5b604e8411156105f257600080fd5b600d5434101561060157600080fd5b600254346000198601600a0a810260090203935083111561062157600080fd5b506001905060005b602081101561068857600154816020811061064057fe5b1a60f860020a02600160f860020a031916868260208110151561065f57fe5b1a60f860020a02600160f860020a0319161415156106805760009150610688565b600101610629565b81151561069457600080fd5b60038054600181016106a68382611121565b9160005260206000209060030201600060606040519081016040528033600160a060020a0316815260200189898080601f01602080910402602001604051908101604052818152929190602084018383808284375050509284525050503488016020909101529190508151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015181600101908051610753929160200190611152565b506040820151600291820155805486900390555050600554151561077657426005555b426006557f543732d2dc4e018310316a0d52a4331a936ee078134d9767f05d605d82394b6a3387878734604051600160a060020a0386168152602081018590526060810182905260806040820181815290820184905260a082018585808284378201915050965050505050505060405180910390a1505050505050565b600c5481565b600a5481565b60015481565b60055481565b60025481565b60005433600160a060020a0390811691161461082c57600080fd5b6003541561083957600080fd5b600955565b60005433600160a060020a0390811691161461085957600080fd5b6003541561086657600080fd5b600b55565b60008060045411151561087d57600080fd5b6005546000901161088d57600080fd5b6007541561089a57600080fd5b600c54600554014210156108ad57600080fd5b5060005b6003548110156109b35760038054829081106108c957fe5b906000526020600020906003020160020154600e60006003848154811015156108ee57fe5b6000918252602080832060039283020154600160a060020a0316845283019390935260409091019020805490920190915580547fd9c1674c23feb7f9dbb60dc6e16dc4e33d1ca6e6b9c1b714d5012c3b1ed138a691908390811061094e57fe5b60009182526020909120600391820201548154600160a060020a0390911691908490811061097857fe5b906000526020600020906003020160020154604051600160a060020a03909216825260208201526040908101905180910390a16001016108b1565b60006109c0600382611121565b50506000600481905560058190556006819055600755565b6000805481908190819033600160a060020a039081169116146109fa57600080fd5b60045460009011610a0a57600080fd5b60075415610a1757600080fd5b604e8514610a2457600080fd5b6006541580610a395750600a54600654014210155b1515610a4457600080fd5b6005541580610a585750600c546005540142105b1515610a6357600080fd5b600154600287876000604051602001526040518083838082843782019150509250505060206040518083038160008661646e5a03f11515610aa357600080fd5b50506040518051919091149050610ab957600080fd5b600093505b600354841015610eca57600380546001945085908110610ada57fe5b9060005260206000209060030201600101805460018160011615610100020316600290049050604e0391508190505b604e811015610be557858582818110610b1e57fe5b905090013560f860020a900460f860020a02600160f860020a031916600160f860020a031916600385815481101515610b5357fe5b90600052602060002090600302016001018383038154600181600116156101000203166002900481101515610b8457fe5b815460011615610ba35790600052602060002090602091828204019190065b9054901a60f860020a027fff000000000000000000000000000000000000000000000000000000000000001614610bdd5760009250610be5565b600101610b09565b8215610d89576003805485908110610bf957fe5b906000526020600020906003020160020154600e6000600387815481101515610c1e57fe5b6000918252602080832060039283020154600160a060020a0316845283019390935260409091019020805490920190915580547f678c162f1e47778e3e037fb6846912e380d12b271b4c20ad03d80a2a82af87f6919086908110610c7e57fe5b60009182526020909120600391820201548154600160a060020a03909116919087908110610ca857fe5b906000526020600020906003020160020154600387815481101515610cc957fe5b9060005260206000209060030201600101600154604051600160a060020a038516815260208101849052606081018290526080604082018181528454600261010060018316150260001901909116049183018290529060a083019085908015610d735780601f10610d4857610100808354040283529160200191610d73565b820191906000526020600020905b815481529060010190602001808311610d5657829003601f168201915b50509550505050505060405180910390a1610ebf565b6003805485908110610d9757fe5b60009182526020909120600391820201600290810154815401905580547faca1f720e44074903b69648d849804d3d9b05c83a2e07736fb897bb1e14506c7919086908110610de157fe5b60009182526020909120600391820201548154600160a060020a03909116919087908110610e0b57fe5b9060005260206000209060030201600101600154604051600160a060020a03841681526040810182905260606020820181815284546002610100600183161502600019019091160491830182905290608083019085908015610eae5780601f10610e8357610100808354040283529160200191610eae565b820191906000526020600020905b815481529060010190602001808311610e9157829003601f168201915b505094505050505060405180910390a15b600190930192610abe565b6000610ed7600382611121565b50426007556001547f6fb34c0e0d06882f1dfe593dfc5da87edc1d80db75915c125d360c270b3510e390878760405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a1505050505050565b60005433600160a060020a03908116911614610f5857600080fd5b60035415610f6557600080fd5b600855565b60005433600160a060020a03908116911614610f8557600080fd5b60035415610f9257600080fd5b6004541580610fa357506000600754115b1515610fae57600080fd5b6007541580610fc35750600b54600754014210155b1515610fce57600080fd5b6001819055426004556000600781905560058190556006557f4a25b99420a56be9a349969b39d040302441e1576e1e92c453ccd9d85a67ed0d8160405190815260200160405180910390a150565b6002805434019055565b600d5481565b60085481565b60005433600160a060020a0390811691161461104d57600080fd5b6003541561105a57600080fd5b600a55565b6003545b90565b600b5481565b600054600160a060020a031681565b600160a060020a0333166000908152600e60205260408120549081111561055257600160a060020a0333166000818152600e60205260408082209190915582156108fc0290839051600060405180830381858888f19350505050151561055257600080fd5b60065481565b60038054829081106110f457fe5b6000918252602090912060039091020180546002820154600160a060020a03909116925060019091019083565b81548183558181151161114d5760030281600302836000526020600020918201910161114d91906111d0565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061119357805160ff19168380011785556111c0565b828001600101855582156111c0579182015b828111156111c05782518255916020019190600101906111a5565b506111cc929150611218565b5090565b61106391905b808211156111cc57805473ffffffffffffffffffffffffffffffffffffffff1916815560006112086001830182611232565b50600060028201556003016111d6565b61106391905b808211156111cc576000815560010161121e565b50805460018160011615610100020316600290046000825580601f106112585750610552565b601f01602090049060005260206000209081019061055291906112185600a165627a7a72305820585c11872edfd82fa3f307273abaeaf352500dde2ff0501720f1676d9e13a2410029`

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

// MAfterLastGuessDelaySecs is a free data retrieval call binding the contract method 0x4ae800f0.
//
// Solidity: function mAfterLastGuessDelaySecs() constant returns(uint256)
func (_Ninex *NinexCaller) MAfterLastGuessDelaySecs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mAfterLastGuessDelaySecs")
	return *ret0, err
}

// MAfterLastGuessDelaySecs is a free data retrieval call binding the contract method 0x4ae800f0.
//
// Solidity: function mAfterLastGuessDelaySecs() constant returns(uint256)
func (_Ninex *NinexSession) MAfterLastGuessDelaySecs() (*big.Int, error) {
	return _Ninex.Contract.MAfterLastGuessDelaySecs(&_Ninex.CallOpts)
}

// MAfterLastGuessDelaySecs is a free data retrieval call binding the contract method 0x4ae800f0.
//
// Solidity: function mAfterLastGuessDelaySecs() constant returns(uint256)
func (_Ninex *NinexCallerSession) MAfterLastGuessDelaySecs() (*big.Int, error) {
	return _Ninex.Contract.MAfterLastGuessDelaySecs(&_Ninex.CallOpts)
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

// MFirstGuessTime is a free data retrieval call binding the contract method 0x525b6b7c.
//
// Solidity: function mFirstGuessTime() constant returns(uint256)
func (_Ninex *NinexCaller) MFirstGuessTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mFirstGuessTime")
	return *ret0, err
}

// MFirstGuessTime is a free data retrieval call binding the contract method 0x525b6b7c.
//
// Solidity: function mFirstGuessTime() constant returns(uint256)
func (_Ninex *NinexSession) MFirstGuessTime() (*big.Int, error) {
	return _Ninex.Contract.MFirstGuessTime(&_Ninex.CallOpts)
}

// MFirstGuessTime is a free data retrieval call binding the contract method 0x525b6b7c.
//
// Solidity: function mFirstGuessTime() constant returns(uint256)
func (_Ninex *NinexCallerSession) MFirstGuessTime() (*big.Int, error) {
	return _Ninex.Contract.MFirstGuessTime(&_Ninex.CallOpts)
}

// MGuessWindowSecs is a free data retrieval call binding the contract method 0x28453d6f.
//
// Solidity: function mGuessWindowSecs() constant returns(uint256)
func (_Ninex *NinexCaller) MGuessWindowSecs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mGuessWindowSecs")
	return *ret0, err
}

// MGuessWindowSecs is a free data retrieval call binding the contract method 0x28453d6f.
//
// Solidity: function mGuessWindowSecs() constant returns(uint256)
func (_Ninex *NinexSession) MGuessWindowSecs() (*big.Int, error) {
	return _Ninex.Contract.MGuessWindowSecs(&_Ninex.CallOpts)
}

// MGuessWindowSecs is a free data retrieval call binding the contract method 0x28453d6f.
//
// Solidity: function mGuessWindowSecs() constant returns(uint256)
func (_Ninex *NinexCallerSession) MGuessWindowSecs() (*big.Int, error) {
	return _Ninex.Contract.MGuessWindowSecs(&_Ninex.CallOpts)
}

// MGuesses is a free data retrieval call binding the contract method 0xf33fffd8.
//
// Solidity: function mGuesses( uint256) constant returns(guesser address, digits bytes, escrow uint256)
func (_Ninex *NinexCaller) MGuesses(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Guesser common.Address
	Digits  []byte
	Escrow  *big.Int
}, error) {
	ret := new(struct {
		Guesser common.Address
		Digits  []byte
		Escrow  *big.Int
	})
	out := ret
	err := _Ninex.contract.Call(opts, out, "mGuesses", arg0)
	return *ret, err
}

// MGuesses is a free data retrieval call binding the contract method 0xf33fffd8.
//
// Solidity: function mGuesses( uint256) constant returns(guesser address, digits bytes, escrow uint256)
func (_Ninex *NinexSession) MGuesses(arg0 *big.Int) (struct {
	Guesser common.Address
	Digits  []byte
	Escrow  *big.Int
}, error) {
	return _Ninex.Contract.MGuesses(&_Ninex.CallOpts, arg0)
}

// MGuesses is a free data retrieval call binding the contract method 0xf33fffd8.
//
// Solidity: function mGuesses( uint256) constant returns(guesser address, digits bytes, escrow uint256)
func (_Ninex *NinexCallerSession) MGuesses(arg0 *big.Int) (struct {
	Guesser common.Address
	Digits  []byte
	Escrow  *big.Int
}, error) {
	return _Ninex.Contract.MGuesses(&_Ninex.CallOpts, arg0)
}

// MLastGuessTime is a free data retrieval call binding the contract method 0xe66c3791.
//
// Solidity: function mLastGuessTime() constant returns(uint256)
func (_Ninex *NinexCaller) MLastGuessTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mLastGuessTime")
	return *ret0, err
}

// MLastGuessTime is a free data retrieval call binding the contract method 0xe66c3791.
//
// Solidity: function mLastGuessTime() constant returns(uint256)
func (_Ninex *NinexSession) MLastGuessTime() (*big.Int, error) {
	return _Ninex.Contract.MLastGuessTime(&_Ninex.CallOpts)
}

// MLastGuessTime is a free data retrieval call binding the contract method 0xe66c3791.
//
// Solidity: function mLastGuessTime() constant returns(uint256)
func (_Ninex *NinexCallerSession) MLastGuessTime() (*big.Int, error) {
	return _Ninex.Contract.MLastGuessTime(&_Ninex.CallOpts)
}

// MMinGuessWei is a free data retrieval call binding the contract method 0xb8df37eb.
//
// Solidity: function mMinGuessWei() constant returns(uint256)
func (_Ninex *NinexCaller) MMinGuessWei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mMinGuessWei")
	return *ret0, err
}

// MMinGuessWei is a free data retrieval call binding the contract method 0xb8df37eb.
//
// Solidity: function mMinGuessWei() constant returns(uint256)
func (_Ninex *NinexSession) MMinGuessWei() (*big.Int, error) {
	return _Ninex.Contract.MMinGuessWei(&_Ninex.CallOpts)
}

// MMinGuessWei is a free data retrieval call binding the contract method 0xb8df37eb.
//
// Solidity: function mMinGuessWei() constant returns(uint256)
func (_Ninex *NinexCallerSession) MMinGuessWei() (*big.Int, error) {
	return _Ninex.Contract.MMinGuessWei(&_Ninex.CallOpts)
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

// NumGuesses is a free data retrieval call binding the contract method 0xd659d9a5.
//
// Solidity: function numGuesses() constant returns(uint256)
func (_Ninex *NinexCaller) NumGuesses(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "numGuesses")
	return *ret0, err
}

// NumGuesses is a free data retrieval call binding the contract method 0xd659d9a5.
//
// Solidity: function numGuesses() constant returns(uint256)
func (_Ninex *NinexSession) NumGuesses() (*big.Int, error) {
	return _Ninex.Contract.NumGuesses(&_Ninex.CallOpts)
}

// NumGuesses is a free data retrieval call binding the contract method 0xd659d9a5.
//
// Solidity: function numGuesses() constant returns(uint256)
func (_Ninex *NinexCallerSession) NumGuesses() (*big.Int, error) {
	return _Ninex.Contract.NumGuesses(&_Ninex.CallOpts)
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

// SetAfterLastGuessDelay is a paid mutator transaction binding the contract method 0xcb18c643.
//
// Solidity: function setAfterLastGuessDelay(val uint256) returns()
func (_Ninex *NinexTransactor) SetAfterLastGuessDelay(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "setAfterLastGuessDelay", val)
}

// SetAfterLastGuessDelay is a paid mutator transaction binding the contract method 0xcb18c643.
//
// Solidity: function setAfterLastGuessDelay(val uint256) returns()
func (_Ninex *NinexSession) SetAfterLastGuessDelay(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetAfterLastGuessDelay(&_Ninex.TransactOpts, val)
}

// SetAfterLastGuessDelay is a paid mutator transaction binding the contract method 0xcb18c643.
//
// Solidity: function setAfterLastGuessDelay(val uint256) returns()
func (_Ninex *NinexTransactorSession) SetAfterLastGuessDelay(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetAfterLastGuessDelay(&_Ninex.TransactOpts, val)
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

// SetGuessWindow is a paid mutator transaction binding the contract method 0x663ae028.
//
// Solidity: function setGuessWindow(val uint256) returns()
func (_Ninex *NinexTransactor) SetGuessWindow(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "setGuessWindow", val)
}

// SetGuessWindow is a paid mutator transaction binding the contract method 0x663ae028.
//
// Solidity: function setGuessWindow(val uint256) returns()
func (_Ninex *NinexSession) SetGuessWindow(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetGuessWindow(&_Ninex.TransactOpts, val)
}

// SetGuessWindow is a paid mutator transaction binding the contract method 0x663ae028.
//
// Solidity: function setGuessWindow(val uint256) returns()
func (_Ninex *NinexTransactorSession) SetGuessWindow(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetGuessWindow(&_Ninex.TransactOpts, val)
}

// SetMinGuessWei is a paid mutator transaction binding the contract method 0x196161fa.
//
// Solidity: function setMinGuessWei(val uint256) returns()
func (_Ninex *NinexTransactor) SetMinGuessWei(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Ninex.contract.Transact(opts, "setMinGuessWei", val)
}

// SetMinGuessWei is a paid mutator transaction binding the contract method 0x196161fa.
//
// Solidity: function setMinGuessWei(val uint256) returns()
func (_Ninex *NinexSession) SetMinGuessWei(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetMinGuessWei(&_Ninex.TransactOpts, val)
}

// SetMinGuessWei is a paid mutator transaction binding the contract method 0x196161fa.
//
// Solidity: function setMinGuessWei(val uint256) returns()
func (_Ninex *NinexTransactorSession) SetMinGuessWei(val *big.Int) (*types.Transaction, error) {
	return _Ninex.Contract.SetMinGuessWei(&_Ninex.TransactOpts, val)
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
