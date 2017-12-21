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
const NinexABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"mPayments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mRevealedTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mCommitmentSetTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMinGuessWei\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mGuessWindowSecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mPreimage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setRevealTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\"},{\"name\":\"digits\",\"type\":\"bytes\"}],\"name\":\"guess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mRevealTimeoutSecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mAfterLastGuessDelaySecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mCommitment\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mFirstGuessTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mBank\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setGuessWindow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAfterRevealDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAfterCommitmentDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"setCommitment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mMinGuessWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mAfterCommitmentDelaySecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAfterLastGuessDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numGuesses\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mAfterRevealDelaySecs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"collect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mLastGuessTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mGuesses\",\"outputs\":[{\"name\":\"guesser\",\"type\":\"address\"},{\"name\":\"digits\",\"type\":\"bytes\"},{\"name\":\"escrow\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guesser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"digits\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"evGuess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"evNewCommitment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guesser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"digits\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"evWin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guesser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"digits\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"evLose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"evRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"guesser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"evWinByDefault\",\"type\":\"event\"}]"

// NinexBin is the compiled bytecode used for deploying new contracts.
const NinexBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161781556102586009819055610384600a55600b819055600c5562015180600d55620f4240600e556003819055600581905560068190556007819055600855611479806100766000396000f30060606040526004361061017f5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630342a929811461018457806309ecabf5146101b55780630f6f2503146101c8578063196161fa146101db57806328453d6f146101f35780632b98ea9e146102065780632e1a7d4d1461029057806338b64085146102a65780633d58f99b146102bc5780633ef314aa146102d35780634ae800f0146102e65780634d4f604a146102f9578063525b6b7c1461030c578063592dac2a1461031f578063663ae028146103325780636a2403b81461034857806370dea79a1461035e57806372f12a5d146103715780639594c3dc1461038f578063abe20032146103a5578063b60d4288146103bb578063b8df37eb146103c3578063c571598f146103d6578063cb18c643146103e9578063d659d9a5146103ff578063d7b5412d14610412578063dcda4bf314610425578063e522538114610454578063e66c379114610467578063f33fffd81461047a575b600080fd5b341561018f57600080fd5b6101a3600160a060020a0360043516610530565b60405190815260200160405180910390f35b34156101c057600080fd5b6101a3610542565b34156101d357600080fd5b6101a3610548565b34156101e657600080fd5b6101f160043561054e565b005b34156101fe57600080fd5b6101a361057b565b341561021157600080fd5b610219610581565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561025557808201518382015260200161023d565b50505050905090810190601f1680156102825780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561029b57600080fd5b6101f160043561061f565b34156102b157600080fd5b6101f1600435610688565b6101f16004803590602480359081019101356106b5565b34156102de57600080fd5b6101a3610928565b34156102f157600080fd5b6101a361092e565b341561030457600080fd5b6101a3610934565b341561031757600080fd5b6101a361093a565b341561032a57600080fd5b6101a3610940565b341561033d57600080fd5b6101f1600435610946565b341561035357600080fd5b6101f1600435610973565b341561036957600080fd5b6101f16109a0565b341561037c57600080fd5b6101f16004803560248101910135610b0f565b341561039a57600080fd5b6101f1600435611083565b34156103b057600080fd5b6101f16004356110b0565b6101f1611185565b34156103ce57600080fd5b6101a361118f565b34156103e157600080fd5b6101a3611195565b34156103f457600080fd5b6101f160043561119b565b341561040a57600080fd5b6101a36111c8565b341561041d57600080fd5b6101a36111cf565b341561043057600080fd5b6104386111d5565b604051600160a060020a03909116815260200160405180910390f35b341561045f57600080fd5b6101f16111e4565b341561047257600080fd5b6101a3611249565b341561048557600080fd5b61049060043561124f565b604051600160a060020a0384168152604081018290526060602082018181528454600260001961010060018416150201909116049183018290529060808301908590801561051f5780601f106104f45761010080835404028352916020019161051f565b820191906000526020600020905b81548152906001019060200180831161050257829003601f168201915b505094505050505060405180910390f35b600f6020526000908152604090205481565b60085481565b60055481565b60005433600160a060020a0390811691161461056957600080fd5b6004541561057657600080fd5b600e55565b600a5481565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106175780601f106105ec57610100808354040283529160200191610617565b820191906000526020600020905b8154815290600101906020018083116105fa57829003601f168201915b505050505081565b60005433600160a060020a0390811691161461063a57600080fd5b60035481111561064957600080fd5b600380548290039055600054600160a060020a031681156108fc0282604051600060405180830381858888f19350505050151561068557600080fd5b50565b60005433600160a060020a039081169116146106a357600080fd5b600454156106b057600080fd5b600d55565b6000806000806005541115156106ca57600080fd5b600854156106d757600080fd5b600954600554014210156106ea57600080fd5b60065415806106fe5750600a546006540142105b151561070957600080fd5b600184101561071757600080fd5b604e84111561072557600080fd5b600e5434101561073457600080fd5b600354346000198601600a0a810260090203935083111561075457600080fd5b506001905060005b60208110156107bb57600154816020811061077357fe5b1a60f860020a02600160f860020a031916868260208110151561079257fe5b1a60f860020a02600160f860020a0319161415156107b357600091506107bb565b60010161075c565b8115156107c757600080fd5b60048054600181016107d9838261128a565b9160005260206000209060030201600060606040519081016040528033600160a060020a0316815260200189898080601f01602080910402602001604051908101604052818152929190602084018383808284375050509284525050503488016020909101529190508151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151816001019080516108869291602001906112bb565b506040820151600290910155505060038054849003905560065415156108ab57426006555b426007557f543732d2dc4e018310316a0d52a4331a936ee078134d9767f05d605d82394b6a3387878734604051600160a060020a0386168152602081018590526060810182905260806040820181815290820184905260a082018585808284378201915050965050505050505060405180910390a1505050505050565b600d5481565b600b5481565b60015481565b60065481565b60035481565b60005433600160a060020a0390811691161461096157600080fd5b6004541561096e57600080fd5b600a55565b60005433600160a060020a0390811691161461098e57600080fd5b6004541561099b57600080fd5b600c55565b6000806005541115156109b257600080fd5b600654600090116109c257600080fd5b600854156109cf57600080fd5b600d54600654014210156109e257600080fd5b5060005b600454811015610aea5760048054829081106109fe57fe5b906000526020600020906003020160020154600f6000600484815481101515610a2357fe5b60009182526020808320600390920290910154600160a060020a03168352820192909252604001902080549091019055600480547fd9c1674c23feb7f9dbb60dc6e16dc4e33d1ca6e6b9c1b714d5012c3b1ed138a6919083908110610a8457fe5b600091825260209091206003909102015460048054600160a060020a039092169184908110610aaf57fe5b906000526020600020906003020160020154604051600160a060020a03909216825260208201526040908101905180910390a16001016109e6565b6000610af760048261128a565b50506000600581905560068190556007819055600855565b6000805481908190819033600160a060020a03908116911614610b3157600080fd5b60055460009011610b4157600080fd5b60085415610b4e57600080fd5b604e8514610b5b57600080fd5b6007541580610b705750600b54600754014210155b1515610b7b57600080fd5b6006541580610b8f5750600d546006540142105b1515610b9a57600080fd5b600154600287876000604051602001526040518083838082843782019150509250505060206040518083038160008661646e5a03f11515610bda57600080fd5b50506040518051919091149050610bf057600080fd5b600093505b60045484101561100457600480546001945085908110610c1157fe5b9060005260206000209060030201600101805460018160011615610100020316600290049050604e0391508190505b604e811015610d1c57858582818110610c5557fe5b905090013560f860020a900460f860020a02600160f860020a031916600160f860020a031916600485815481101515610c8a57fe5b90600052602060002090600302016001018383038154600181600116156101000203166002900481101515610cbb57fe5b815460011615610cda5790600052602060002090602091828204019190065b9054901a60f860020a027fff000000000000000000000000000000000000000000000000000000000000001614610d145760009250610d1c565b600101610c40565b8215610ec2576004805485908110610d3057fe5b906000526020600020906003020160020154600f6000600487815481101515610d5557fe5b60009182526020808320600390920290910154600160a060020a03168352820192909252604001902080549091019055600480547f678c162f1e47778e3e037fb6846912e380d12b271b4c20ad03d80a2a82af87f6919086908110610db657fe5b600091825260209091206003909102015460048054600160a060020a039092169187908110610de157fe5b906000526020600020906003020160020154600487815481101515610e0257fe5b9060005260206000209060030201600101600154604051600160a060020a038516815260208101849052606081018290526080604082018181528454600261010060018316150260001901909116049183018290529060a083019085908015610eac5780601f10610e8157610100808354040283529160200191610eac565b820191906000526020600020905b815481529060010190602001808311610e8f57829003601f168201915b50509550505050505060405180910390a1610ff9565b6004805485908110610ed057fe5b60009182526020909120600391820201600201548154019055600480547faca1f720e44074903b69648d849804d3d9b05c83a2e07736fb897bb1e14506c7919086908110610f1a57fe5b600091825260209091206003909102015460048054600160a060020a039092169187908110610f4557fe5b9060005260206000209060030201600101600154604051600160a060020a03841681526040810182905260606020820181815284546002610100600183161502600019019091160491830182905290608083019085908015610fe85780601f10610fbd57610100808354040283529160200191610fe8565b820191906000526020600020905b815481529060010190602001808311610fcb57829003601f168201915b505094505050505060405180910390a15b600190930192610bf5565b600061101160048261128a565b504260085561102260028787611339565b507f6fb34c0e0d06882f1dfe593dfc5da87edc1d80db75915c125d360c270b3510e3600154878760405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a1505050505050565b60005433600160a060020a0390811691161461109e57600080fd5b600454156110ab57600080fd5b600955565b60005433600160a060020a039081169116146110cb57600080fd5b600454156110d857600080fd5b60055415806110e957506000600854115b15156110f457600080fd5b60085415806111095750600c54600854014210155b151561111457600080fd5b6001819055602060405190810160405260008152600290805161113b9291602001906112bb565b50426005556000600881905560068190556007557f4a25b99420a56be9a349969b39d040302441e1576e1e92c453ccd9d85a67ed0d8160405190815260200160405180910390a150565b6003805434019055565b600e5481565b60095481565b60005433600160a060020a039081169116146111b657600080fd5b600454156111c357600080fd5b600b55565b6004545b90565b600c5481565b600054600160a060020a031681565b600160a060020a0333166000908152600f60205260408120549081111561068557600160a060020a0333166000818152600f60205260408082209190915582156108fc0290839051600060405180830381858888f19350505050151561068557600080fd5b60075481565b600480548290811061125d57fe5b6000918252602090912060039091020180546002820154600160a060020a03909116925060019091019083565b8154818355818115116112b6576003028160030283600052602060002091820191016112b691906113a7565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112fc57805160ff1916838001178555611329565b82800160010185558215611329579182015b8281111561132957825182559160200191906001019061130e565b506113359291506113ef565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061137a5782800160ff19823516178555611329565b82800160010185558215611329579182015b8281111561132957823582559160200191906001019061138c565b6111cc91905b8082111561133557805473ffffffffffffffffffffffffffffffffffffffff1916815560006113df6001830182611409565b50600060028201556003016113ad565b6111cc91905b8082111561133557600081556001016113f5565b50805460018160011615610100020316600290046000825580601f1061142f5750610685565b601f01602090049060005260206000209081019061068591906113ef5600a165627a7a723058201cb53282cbcbbb5582ed7cc0d1e596a3f5fc3637e35f9135c3146f6a217f307d0029`

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

// MPreimage is a free data retrieval call binding the contract method 0x2b98ea9e.
//
// Solidity: function mPreimage() constant returns(bytes)
func (_Ninex *NinexCaller) MPreimage(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Ninex.contract.Call(opts, out, "mPreimage")
	return *ret0, err
}

// MPreimage is a free data retrieval call binding the contract method 0x2b98ea9e.
//
// Solidity: function mPreimage() constant returns(bytes)
func (_Ninex *NinexSession) MPreimage() ([]byte, error) {
	return _Ninex.Contract.MPreimage(&_Ninex.CallOpts)
}

// MPreimage is a free data retrieval call binding the contract method 0x2b98ea9e.
//
// Solidity: function mPreimage() constant returns(bytes)
func (_Ninex *NinexCallerSession) MPreimage() ([]byte, error) {
	return _Ninex.Contract.MPreimage(&_Ninex.CallOpts)
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
