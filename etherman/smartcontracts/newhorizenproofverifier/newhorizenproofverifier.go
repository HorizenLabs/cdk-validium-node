// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package newhorizenproofverifier

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NewhorizenproofverifierMetaData contains all meta data concerning the Newhorizenproofverifier contract.
var NewhorizenproofverifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBatchCounts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCannotRenounce\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_attestationId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_proofsAttestation\",\"type\":\"bytes32\"}],\"name\":\"AttestationPosted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipIsEnforcingSequentialAttestations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnforcingSequentialAttestations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAttestationId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proofsAttestations\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_attestationId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_proofsAttestation\",\"type\":\"bytes32\"}],\"name\":\"submitAttestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_attestationIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proofsAttestation\",\"type\":\"bytes32[]\"}],\"name\":\"submitAttestationBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_attestationId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merklePath\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_leafCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"verifyProofAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051620013d9380380620013d98339810160408190526100319161010b565b61003c60003361006c565b6100667f523a704056dcd17bcf83bed8b68c59416dac1119be77755efe3bde0a64e46e0c8261006c565b5061013b565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16610107576000828152602081815260408083206001600160a01b03851684529091529020805460ff191660011790556100c63390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b60006020828403121561011d57600080fd5b81516001600160a01b038116811461013457600080fd5b9392505050565b61128e806200014b6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806391d1485411610097578063bd93187c11610066578063bd93187c1461023d578063d547741f1461024a578063ddea78861461025d578063e0fec29a1461027057600080fd5b806391d14854146101aa578063983d2737146101ee578063a217fddf14610215578063b99047df1461021d57600080fd5b806336568abe116100d357806336568abe14610168578063560f3ba41461017b57806357ae920d146101845780636736ec951461019757600080fd5b806301ffc9a7146100fa578063248a9ca3146101225780632f2ff15d14610153575b600080fd5b61010d610108366004610df1565b610278565b60405190151581526020015b60405180910390f35b610145610130366004610e33565b60009081526020819052604090206001015490565b604051908152602001610119565b610166610161366004610e4c565b610311565b005b610166610176366004610e4c565b61033b565b61014560015481565b610166610192366004610ee1565b610380565b6101666101a5366004610f4d565b610559565b61010d6101b8366004610e4c565b60009182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101457f523a704056dcd17bcf83bed8b68c59416dac1119be77755efe3bde0a64e46e0c81565b610145600081565b61014561022b366004610e33565b60026020526000908152604090205481565b60035461010d9060ff1681565b610166610258366004610e4c565b61061d565b61010d61026b366004610f6f565b610642565b610166610681565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061030b57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b60008281526020819052604090206001015461032c816106de565b61033683836106eb565b505050565b81610372576040517f1d88406500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61037c82826107db565b5050565b7f523a704056dcd17bcf83bed8b68c59416dac1119be77755efe3bde0a64e46e0c6103aa816106de565b8382146103e3576040517f73d99d3700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015460035485919060ff161561046e5760005b8281101561046c5761040a816001611003565b6104149083611003565b88888381811061042657610426611016565b9050602002013514610464576040517fbd8ba84d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001016103f7565b505b60005b828110156105275785858281811061048b5761048b611016565b90506020020135600260008a8a858181106104a8576104a8611016565b905060200201358152602001908152602001600020819055508585828181106104d3576104d3611016565b905060200201358888838181106104ec576104ec611016565b905060200201357fe64e12e2dfd684ae91bf1a6c52fbf2af6986db5927d23f5b643e0d50f7de487860405160405180910390a3600101610471565b508686610535600182611045565b81811061054457610544611016565b60200291909101356001555050505050505050565b7f523a704056dcd17bcf83bed8b68c59416dac1119be77755efe3bde0a64e46e0c610583816106de565b60035460ff1680156105a157506001805461059d91611003565b8314155b156105d8576040517fbd8ba84d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600183905560008381526002602052604080822084905551839185917fe64e12e2dfd684ae91bf1a6c52fbf2af6986db5927d23f5b643e0d50f7de48789190a3505050565b600082815260208190526040902060010154610638816106de565b610336838361088b565b600060015487111561065657506000610677565b60008781526002602052604090205461067381878787878c610942565b9150505b9695505050505050565b7f523a704056dcd17bcf83bed8b68c59416dac1119be77755efe3bde0a64e46e0c6106ab816106de565b50600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00811660ff90911615179055565b6106e88133610ad0565b50565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661037c5760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff85168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561077d3390565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b73ffffffffffffffffffffffffffffffffffffffff81163314610885576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084015b60405180910390fd5b61037c82825b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff161561037c5760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b600083831061097d576040517f4e23d03500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260405160200161099291815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209050838560005b88811015610ac05760008a8a838181106109eb576109eb611016565b905060200201359050600284610a019190611087565b60011480610a18575082610a16856001611003565b145b15610a4e576040805160208101839052908101869052606001604051602081830303815290604052805190602001209450610a7b565b60408051602081018790529081018290526060016040516020818303038152906040528051906020012094505b610a8660028561109b565b93506002610a95600185611045565b610a9f919061109b565b610aaa906001611003565b9250508080610ab8906110af565b9150506109cf565b5050509096149695505050505050565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661037c57610b0e81610b88565b610b19836020610ba7565b604051602001610b2a92919061110b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261087c9160040161118c565b606061030b73ffffffffffffffffffffffffffffffffffffffff831660145b60606000610bb68360026111dd565b610bc1906002611003565b67ffffffffffffffff811115610bd957610bd96111f4565b6040519080825280601f01601f191660200182016040528015610c03576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610c3a57610c3a611016565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610c9d57610c9d611016565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000610cd98460026111dd565b610ce4906001611003565b90505b6001811115610d81577f303132333435363738396162636465660000000000000000000000000000000085600f1660108110610d2557610d25611016565b1a60f81b828281518110610d3b57610d3b611016565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c93610d7a81611223565b9050610ce7565b508315610dea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161087c565b9392505050565b600060208284031215610e0357600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610dea57600080fd5b600060208284031215610e4557600080fd5b5035919050565b60008060408385031215610e5f57600080fd5b82359150602083013573ffffffffffffffffffffffffffffffffffffffff81168114610e8a57600080fd5b809150509250929050565b60008083601f840112610ea757600080fd5b50813567ffffffffffffffff811115610ebf57600080fd5b6020830191508360208260051b8501011115610eda57600080fd5b9250929050565b60008060008060408587031215610ef757600080fd5b843567ffffffffffffffff80821115610f0f57600080fd5b610f1b88838901610e95565b90965094506020870135915080821115610f3457600080fd5b50610f4187828801610e95565b95989497509550505050565b60008060408385031215610f6057600080fd5b50508035926020909101359150565b60008060008060008060a08789031215610f8857600080fd5b8635955060208701359450604087013567ffffffffffffffff811115610fad57600080fd5b610fb989828a01610e95565b979a9699509760608101359660809091013595509350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561030b5761030b610fd4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8181038181111561030b5761030b610fd4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261109657611096611058565b500690565b6000826110aa576110aa611058565b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110e0576110e0610fd4565b5060010190565b60005b838110156111025781810151838201526020016110ea565b50506000910152565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008152600083516111438160178501602088016110e7565b7f206973206d697373696e6720726f6c652000000000000000000000000000000060179184019182015283516111808160288401602088016110e7565b01602801949350505050565b60208152600082518060208401526111ab8160408501602087016110e7565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b808202811582820484141761030b5761030b610fd4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008161123257611232610fd4565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fea2646970667358221220adcee907f87793daf50397db38deca3628000a6693a39cfbc0703c2646a3530b64736f6c63430008140033",
}

// NewhorizenproofverifierABI is the input ABI used to generate the binding from.
// Deprecated: Use NewhorizenproofverifierMetaData.ABI instead.
var NewhorizenproofverifierABI = NewhorizenproofverifierMetaData.ABI

// NewhorizenproofverifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NewhorizenproofverifierMetaData.Bin instead.
var NewhorizenproofverifierBin = NewhorizenproofverifierMetaData.Bin

// DeployNewhorizenproofverifier deploys a new Ethereum contract, binding an instance of Newhorizenproofverifier to it.
func DeployNewhorizenproofverifier(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address) (common.Address, *types.Transaction, *Newhorizenproofverifier, error) {
	parsed, err := NewhorizenproofverifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NewhorizenproofverifierBin), backend, _operator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Newhorizenproofverifier{NewhorizenproofverifierCaller: NewhorizenproofverifierCaller{contract: contract}, NewhorizenproofverifierTransactor: NewhorizenproofverifierTransactor{contract: contract}, NewhorizenproofverifierFilterer: NewhorizenproofverifierFilterer{contract: contract}}, nil
}

// Newhorizenproofverifier is an auto generated Go binding around an Ethereum contract.
type Newhorizenproofverifier struct {
	NewhorizenproofverifierCaller     // Read-only binding to the contract
	NewhorizenproofverifierTransactor // Write-only binding to the contract
	NewhorizenproofverifierFilterer   // Log filterer for contract events
}

// NewhorizenproofverifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type NewhorizenproofverifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewhorizenproofverifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NewhorizenproofverifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewhorizenproofverifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NewhorizenproofverifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewhorizenproofverifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NewhorizenproofverifierSession struct {
	Contract     *Newhorizenproofverifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NewhorizenproofverifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NewhorizenproofverifierCallerSession struct {
	Contract *NewhorizenproofverifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// NewhorizenproofverifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NewhorizenproofverifierTransactorSession struct {
	Contract     *NewhorizenproofverifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// NewhorizenproofverifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type NewhorizenproofverifierRaw struct {
	Contract *Newhorizenproofverifier // Generic contract binding to access the raw methods on
}

// NewhorizenproofverifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NewhorizenproofverifierCallerRaw struct {
	Contract *NewhorizenproofverifierCaller // Generic read-only contract binding to access the raw methods on
}

// NewhorizenproofverifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NewhorizenproofverifierTransactorRaw struct {
	Contract *NewhorizenproofverifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNewhorizenproofverifier creates a new instance of Newhorizenproofverifier, bound to a specific deployed contract.
func NewNewhorizenproofverifier(address common.Address, backend bind.ContractBackend) (*Newhorizenproofverifier, error) {
	contract, err := bindNewhorizenproofverifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Newhorizenproofverifier{NewhorizenproofverifierCaller: NewhorizenproofverifierCaller{contract: contract}, NewhorizenproofverifierTransactor: NewhorizenproofverifierTransactor{contract: contract}, NewhorizenproofverifierFilterer: NewhorizenproofverifierFilterer{contract: contract}}, nil
}

// NewNewhorizenproofverifierCaller creates a new read-only instance of Newhorizenproofverifier, bound to a specific deployed contract.
func NewNewhorizenproofverifierCaller(address common.Address, caller bind.ContractCaller) (*NewhorizenproofverifierCaller, error) {
	contract, err := bindNewhorizenproofverifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NewhorizenproofverifierCaller{contract: contract}, nil
}

// NewNewhorizenproofverifierTransactor creates a new write-only instance of Newhorizenproofverifier, bound to a specific deployed contract.
func NewNewhorizenproofverifierTransactor(address common.Address, transactor bind.ContractTransactor) (*NewhorizenproofverifierTransactor, error) {
	contract, err := bindNewhorizenproofverifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NewhorizenproofverifierTransactor{contract: contract}, nil
}

// NewNewhorizenproofverifierFilterer creates a new log filterer instance of Newhorizenproofverifier, bound to a specific deployed contract.
func NewNewhorizenproofverifierFilterer(address common.Address, filterer bind.ContractFilterer) (*NewhorizenproofverifierFilterer, error) {
	contract, err := bindNewhorizenproofverifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NewhorizenproofverifierFilterer{contract: contract}, nil
}

// bindNewhorizenproofverifier binds a generic wrapper to an already deployed contract.
func bindNewhorizenproofverifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NewhorizenproofverifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Newhorizenproofverifier *NewhorizenproofverifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Newhorizenproofverifier.Contract.NewhorizenproofverifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Newhorizenproofverifier *NewhorizenproofverifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.NewhorizenproofverifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Newhorizenproofverifier *NewhorizenproofverifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.NewhorizenproofverifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Newhorizenproofverifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Newhorizenproofverifier.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Newhorizenproofverifier.Contract.DEFAULTADMINROLE(&_Newhorizenproofverifier.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Newhorizenproofverifier.Contract.DEFAULTADMINROLE(&_Newhorizenproofverifier.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierCaller) OPERATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Newhorizenproofverifier.contract.Call(opts, &out, "OPERATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) OPERATOR() ([32]byte, error) {
	return _Newhorizenproofverifier.Contract.OPERATOR(&_Newhorizenproofverifier.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerSession) OPERATOR() ([32]byte, error) {
	return _Newhorizenproofverifier.Contract.OPERATOR(&_Newhorizenproofverifier.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Newhorizenproofverifier.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Newhorizenproofverifier.Contract.GetRoleAdmin(&_Newhorizenproofverifier.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Newhorizenproofverifier.Contract.GetRoleAdmin(&_Newhorizenproofverifier.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Newhorizenproofverifier.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Newhorizenproofverifier.Contract.HasRole(&_Newhorizenproofverifier.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Newhorizenproofverifier.Contract.HasRole(&_Newhorizenproofverifier.CallOpts, role, account)
}

// IsEnforcingSequentialAttestations is a free data retrieval call binding the contract method 0xbd93187c.
//
// Solidity: function isEnforcingSequentialAttestations() view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierCaller) IsEnforcingSequentialAttestations(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Newhorizenproofverifier.contract.Call(opts, &out, "isEnforcingSequentialAttestations")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnforcingSequentialAttestations is a free data retrieval call binding the contract method 0xbd93187c.
//
// Solidity: function isEnforcingSequentialAttestations() view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) IsEnforcingSequentialAttestations() (bool, error) {
	return _Newhorizenproofverifier.Contract.IsEnforcingSequentialAttestations(&_Newhorizenproofverifier.CallOpts)
}

// IsEnforcingSequentialAttestations is a free data retrieval call binding the contract method 0xbd93187c.
//
// Solidity: function isEnforcingSequentialAttestations() view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerSession) IsEnforcingSequentialAttestations() (bool, error) {
	return _Newhorizenproofverifier.Contract.IsEnforcingSequentialAttestations(&_Newhorizenproofverifier.CallOpts)
}

// LatestAttestationId is a free data retrieval call binding the contract method 0x560f3ba4.
//
// Solidity: function latestAttestationId() view returns(uint256)
func (_Newhorizenproofverifier *NewhorizenproofverifierCaller) LatestAttestationId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Newhorizenproofverifier.contract.Call(opts, &out, "latestAttestationId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAttestationId is a free data retrieval call binding the contract method 0x560f3ba4.
//
// Solidity: function latestAttestationId() view returns(uint256)
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) LatestAttestationId() (*big.Int, error) {
	return _Newhorizenproofverifier.Contract.LatestAttestationId(&_Newhorizenproofverifier.CallOpts)
}

// LatestAttestationId is a free data retrieval call binding the contract method 0x560f3ba4.
//
// Solidity: function latestAttestationId() view returns(uint256)
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerSession) LatestAttestationId() (*big.Int, error) {
	return _Newhorizenproofverifier.Contract.LatestAttestationId(&_Newhorizenproofverifier.CallOpts)
}

// ProofsAttestations is a free data retrieval call binding the contract method 0xb99047df.
//
// Solidity: function proofsAttestations(uint256 ) view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierCaller) ProofsAttestations(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Newhorizenproofverifier.contract.Call(opts, &out, "proofsAttestations", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProofsAttestations is a free data retrieval call binding the contract method 0xb99047df.
//
// Solidity: function proofsAttestations(uint256 ) view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) ProofsAttestations(arg0 *big.Int) ([32]byte, error) {
	return _Newhorizenproofverifier.Contract.ProofsAttestations(&_Newhorizenproofverifier.CallOpts, arg0)
}

// ProofsAttestations is a free data retrieval call binding the contract method 0xb99047df.
//
// Solidity: function proofsAttestations(uint256 ) view returns(bytes32)
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerSession) ProofsAttestations(arg0 *big.Int) ([32]byte, error) {
	return _Newhorizenproofverifier.Contract.ProofsAttestations(&_Newhorizenproofverifier.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Newhorizenproofverifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Newhorizenproofverifier.Contract.SupportsInterface(&_Newhorizenproofverifier.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Newhorizenproofverifier.Contract.SupportsInterface(&_Newhorizenproofverifier.CallOpts, interfaceId)
}

// VerifyProofAttestation is a free data retrieval call binding the contract method 0xddea7886.
//
// Solidity: function verifyProofAttestation(uint256 _attestationId, bytes32 _leaf, bytes32[] _merklePath, uint256 _leafCount, uint256 _index) view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierCaller) VerifyProofAttestation(opts *bind.CallOpts, _attestationId *big.Int, _leaf [32]byte, _merklePath [][32]byte, _leafCount *big.Int, _index *big.Int) (bool, error) {
	var out []interface{}
	err := _Newhorizenproofverifier.contract.Call(opts, &out, "verifyProofAttestation", _attestationId, _leaf, _merklePath, _leafCount, _index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProofAttestation is a free data retrieval call binding the contract method 0xddea7886.
//
// Solidity: function verifyProofAttestation(uint256 _attestationId, bytes32 _leaf, bytes32[] _merklePath, uint256 _leafCount, uint256 _index) view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) VerifyProofAttestation(_attestationId *big.Int, _leaf [32]byte, _merklePath [][32]byte, _leafCount *big.Int, _index *big.Int) (bool, error) {
	return _Newhorizenproofverifier.Contract.VerifyProofAttestation(&_Newhorizenproofverifier.CallOpts, _attestationId, _leaf, _merklePath, _leafCount, _index)
}

// VerifyProofAttestation is a free data retrieval call binding the contract method 0xddea7886.
//
// Solidity: function verifyProofAttestation(uint256 _attestationId, bytes32 _leaf, bytes32[] _merklePath, uint256 _leafCount, uint256 _index) view returns(bool)
func (_Newhorizenproofverifier *NewhorizenproofverifierCallerSession) VerifyProofAttestation(_attestationId *big.Int, _leaf [32]byte, _merklePath [][32]byte, _leafCount *big.Int, _index *big.Int) (bool, error) {
	return _Newhorizenproofverifier.Contract.VerifyProofAttestation(&_Newhorizenproofverifier.CallOpts, _attestationId, _leaf, _merklePath, _leafCount, _index)
}

// FlipIsEnforcingSequentialAttestations is a paid mutator transaction binding the contract method 0xe0fec29a.
//
// Solidity: function flipIsEnforcingSequentialAttestations() returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactor) FlipIsEnforcingSequentialAttestations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Newhorizenproofverifier.contract.Transact(opts, "flipIsEnforcingSequentialAttestations")
}

// FlipIsEnforcingSequentialAttestations is a paid mutator transaction binding the contract method 0xe0fec29a.
//
// Solidity: function flipIsEnforcingSequentialAttestations() returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) FlipIsEnforcingSequentialAttestations() (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.FlipIsEnforcingSequentialAttestations(&_Newhorizenproofverifier.TransactOpts)
}

// FlipIsEnforcingSequentialAttestations is a paid mutator transaction binding the contract method 0xe0fec29a.
//
// Solidity: function flipIsEnforcingSequentialAttestations() returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactorSession) FlipIsEnforcingSequentialAttestations() (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.FlipIsEnforcingSequentialAttestations(&_Newhorizenproofverifier.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Newhorizenproofverifier.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.GrantRole(&_Newhorizenproofverifier.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.GrantRole(&_Newhorizenproofverifier.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Newhorizenproofverifier.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.RenounceRole(&_Newhorizenproofverifier.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.RenounceRole(&_Newhorizenproofverifier.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Newhorizenproofverifier.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.RevokeRole(&_Newhorizenproofverifier.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.RevokeRole(&_Newhorizenproofverifier.TransactOpts, role, account)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6736ec95.
//
// Solidity: function submitAttestation(uint256 _attestationId, bytes32 _proofsAttestation) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactor) SubmitAttestation(opts *bind.TransactOpts, _attestationId *big.Int, _proofsAttestation [32]byte) (*types.Transaction, error) {
	return _Newhorizenproofverifier.contract.Transact(opts, "submitAttestation", _attestationId, _proofsAttestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6736ec95.
//
// Solidity: function submitAttestation(uint256 _attestationId, bytes32 _proofsAttestation) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) SubmitAttestation(_attestationId *big.Int, _proofsAttestation [32]byte) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.SubmitAttestation(&_Newhorizenproofverifier.TransactOpts, _attestationId, _proofsAttestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6736ec95.
//
// Solidity: function submitAttestation(uint256 _attestationId, bytes32 _proofsAttestation) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactorSession) SubmitAttestation(_attestationId *big.Int, _proofsAttestation [32]byte) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.SubmitAttestation(&_Newhorizenproofverifier.TransactOpts, _attestationId, _proofsAttestation)
}

// SubmitAttestationBatch is a paid mutator transaction binding the contract method 0x57ae920d.
//
// Solidity: function submitAttestationBatch(uint256[] _attestationIds, bytes32[] _proofsAttestation) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactor) SubmitAttestationBatch(opts *bind.TransactOpts, _attestationIds []*big.Int, _proofsAttestation [][32]byte) (*types.Transaction, error) {
	return _Newhorizenproofverifier.contract.Transact(opts, "submitAttestationBatch", _attestationIds, _proofsAttestation)
}

// SubmitAttestationBatch is a paid mutator transaction binding the contract method 0x57ae920d.
//
// Solidity: function submitAttestationBatch(uint256[] _attestationIds, bytes32[] _proofsAttestation) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierSession) SubmitAttestationBatch(_attestationIds []*big.Int, _proofsAttestation [][32]byte) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.SubmitAttestationBatch(&_Newhorizenproofverifier.TransactOpts, _attestationIds, _proofsAttestation)
}

// SubmitAttestationBatch is a paid mutator transaction binding the contract method 0x57ae920d.
//
// Solidity: function submitAttestationBatch(uint256[] _attestationIds, bytes32[] _proofsAttestation) returns()
func (_Newhorizenproofverifier *NewhorizenproofverifierTransactorSession) SubmitAttestationBatch(_attestationIds []*big.Int, _proofsAttestation [][32]byte) (*types.Transaction, error) {
	return _Newhorizenproofverifier.Contract.SubmitAttestationBatch(&_Newhorizenproofverifier.TransactOpts, _attestationIds, _proofsAttestation)
}

// NewhorizenproofverifierAttestationPostedIterator is returned from FilterAttestationPosted and is used to iterate over the raw logs and unpacked data for AttestationPosted events raised by the Newhorizenproofverifier contract.
type NewhorizenproofverifierAttestationPostedIterator struct {
	Event *NewhorizenproofverifierAttestationPosted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NewhorizenproofverifierAttestationPostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewhorizenproofverifierAttestationPosted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NewhorizenproofverifierAttestationPosted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NewhorizenproofverifierAttestationPostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewhorizenproofverifierAttestationPostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewhorizenproofverifierAttestationPosted represents a AttestationPosted event raised by the Newhorizenproofverifier contract.
type NewhorizenproofverifierAttestationPosted struct {
	AttestationId     *big.Int
	ProofsAttestation [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAttestationPosted is a free log retrieval operation binding the contract event 0xe64e12e2dfd684ae91bf1a6c52fbf2af6986db5927d23f5b643e0d50f7de4878.
//
// Solidity: event AttestationPosted(uint256 indexed _attestationId, bytes32 indexed _proofsAttestation)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) FilterAttestationPosted(opts *bind.FilterOpts, _attestationId []*big.Int, _proofsAttestation [][32]byte) (*NewhorizenproofverifierAttestationPostedIterator, error) {

	var _attestationIdRule []interface{}
	for _, _attestationIdItem := range _attestationId {
		_attestationIdRule = append(_attestationIdRule, _attestationIdItem)
	}
	var _proofsAttestationRule []interface{}
	for _, _proofsAttestationItem := range _proofsAttestation {
		_proofsAttestationRule = append(_proofsAttestationRule, _proofsAttestationItem)
	}

	logs, sub, err := _Newhorizenproofverifier.contract.FilterLogs(opts, "AttestationPosted", _attestationIdRule, _proofsAttestationRule)
	if err != nil {
		return nil, err
	}
	return &NewhorizenproofverifierAttestationPostedIterator{contract: _Newhorizenproofverifier.contract, event: "AttestationPosted", logs: logs, sub: sub}, nil
}

// WatchAttestationPosted is a free log subscription operation binding the contract event 0xe64e12e2dfd684ae91bf1a6c52fbf2af6986db5927d23f5b643e0d50f7de4878.
//
// Solidity: event AttestationPosted(uint256 indexed _attestationId, bytes32 indexed _proofsAttestation)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) WatchAttestationPosted(opts *bind.WatchOpts, sink chan<- *NewhorizenproofverifierAttestationPosted, _attestationId []*big.Int, _proofsAttestation [][32]byte) (event.Subscription, error) {

	var _attestationIdRule []interface{}
	for _, _attestationIdItem := range _attestationId {
		_attestationIdRule = append(_attestationIdRule, _attestationIdItem)
	}
	var _proofsAttestationRule []interface{}
	for _, _proofsAttestationItem := range _proofsAttestation {
		_proofsAttestationRule = append(_proofsAttestationRule, _proofsAttestationItem)
	}

	logs, sub, err := _Newhorizenproofverifier.contract.WatchLogs(opts, "AttestationPosted", _attestationIdRule, _proofsAttestationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewhorizenproofverifierAttestationPosted)
				if err := _Newhorizenproofverifier.contract.UnpackLog(event, "AttestationPosted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAttestationPosted is a log parse operation binding the contract event 0xe64e12e2dfd684ae91bf1a6c52fbf2af6986db5927d23f5b643e0d50f7de4878.
//
// Solidity: event AttestationPosted(uint256 indexed _attestationId, bytes32 indexed _proofsAttestation)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) ParseAttestationPosted(log types.Log) (*NewhorizenproofverifierAttestationPosted, error) {
	event := new(NewhorizenproofverifierAttestationPosted)
	if err := _Newhorizenproofverifier.contract.UnpackLog(event, "AttestationPosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NewhorizenproofverifierRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Newhorizenproofverifier contract.
type NewhorizenproofverifierRoleAdminChangedIterator struct {
	Event *NewhorizenproofverifierRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NewhorizenproofverifierRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewhorizenproofverifierRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NewhorizenproofverifierRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NewhorizenproofverifierRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewhorizenproofverifierRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewhorizenproofverifierRoleAdminChanged represents a RoleAdminChanged event raised by the Newhorizenproofverifier contract.
type NewhorizenproofverifierRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NewhorizenproofverifierRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Newhorizenproofverifier.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NewhorizenproofverifierRoleAdminChangedIterator{contract: _Newhorizenproofverifier.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NewhorizenproofverifierRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Newhorizenproofverifier.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewhorizenproofverifierRoleAdminChanged)
				if err := _Newhorizenproofverifier.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) ParseRoleAdminChanged(log types.Log) (*NewhorizenproofverifierRoleAdminChanged, error) {
	event := new(NewhorizenproofverifierRoleAdminChanged)
	if err := _Newhorizenproofverifier.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NewhorizenproofverifierRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Newhorizenproofverifier contract.
type NewhorizenproofverifierRoleGrantedIterator struct {
	Event *NewhorizenproofverifierRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NewhorizenproofverifierRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewhorizenproofverifierRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NewhorizenproofverifierRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NewhorizenproofverifierRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewhorizenproofverifierRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewhorizenproofverifierRoleGranted represents a RoleGranted event raised by the Newhorizenproofverifier contract.
type NewhorizenproofverifierRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NewhorizenproofverifierRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Newhorizenproofverifier.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NewhorizenproofverifierRoleGrantedIterator{contract: _Newhorizenproofverifier.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NewhorizenproofverifierRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Newhorizenproofverifier.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewhorizenproofverifierRoleGranted)
				if err := _Newhorizenproofverifier.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) ParseRoleGranted(log types.Log) (*NewhorizenproofverifierRoleGranted, error) {
	event := new(NewhorizenproofverifierRoleGranted)
	if err := _Newhorizenproofverifier.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NewhorizenproofverifierRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Newhorizenproofverifier contract.
type NewhorizenproofverifierRoleRevokedIterator struct {
	Event *NewhorizenproofverifierRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NewhorizenproofverifierRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewhorizenproofverifierRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NewhorizenproofverifierRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NewhorizenproofverifierRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewhorizenproofverifierRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewhorizenproofverifierRoleRevoked represents a RoleRevoked event raised by the Newhorizenproofverifier contract.
type NewhorizenproofverifierRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NewhorizenproofverifierRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Newhorizenproofverifier.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NewhorizenproofverifierRoleRevokedIterator{contract: _Newhorizenproofverifier.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NewhorizenproofverifierRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Newhorizenproofverifier.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewhorizenproofverifierRoleRevoked)
				if err := _Newhorizenproofverifier.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Newhorizenproofverifier *NewhorizenproofverifierFilterer) ParseRoleRevoked(log types.Log) (*NewhorizenproofverifierRoleRevoked, error) {
	event := new(NewhorizenproofverifierRoleRevoked)
	if err := _Newhorizenproofverifier.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
