// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iFAllocationSaleV8

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

// IFVestableCliff is an auto generated low-level Go binding around an user-defined struct.
type IFVestableCliff struct {
	ClaimTime *big.Int
	Pct       uint8
}

// IFAllocationSaleV8MetaData contains all meta data concerning the IFAllocationSaleV8 contract.
var IFAllocationSaleV8MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_funder\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"_saleToken\",\"type\":\"address\"},{\"internalType\":\"contractIIFRetrievableStakeWeight\",\"name\":\"_allocationMaster\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_trackId\",\"type\":\"uint24\"},{\"internalType\":\"uint80\",\"name\":\"_allocSnapshotTimestamp\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalPayment\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paymentTokenBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"saleTokenBalance\",\"type\":\"uint256\"}],\"name\":\"Cash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyTokenRetrieve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Fund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"OptInBuyback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paymentAmount\",\"type\":\"uint256\"}],\"name\":\"Purchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"casher\",\"type\":\"address\"}],\"name\":\"SetCasher\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"claimTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pct\",\"type\":\"uint8\"}],\"indexed\":true,\"internalType\":\"structIFVestable.Cliff[]\",\"name\":\"cliffPeriod\",\"type\":\"tuple[]\"}],\"name\":\"SetCliffVestingPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"}],\"name\":\"SetFunder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linearVestingEndTime\",\"type\":\"uint256\"}],\"name\":\"SetLinearVestingEndTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_maxTotalPurchasable\",\"type\":\"uint256\"}],\"name\":\"SetMaxTotalPurchasable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"minTotalPayment\",\"type\":\"uint256\"}],\"name\":\"SetMinTotalPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"whitelistRootHash\",\"type\":\"bytes32\"}],\"name\":\"SetWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"whitelistSetter\",\"type\":\"address\"}],\"name\":\"SetWhitelistSetter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"withdrawDelay\",\"type\":\"uint24\"}],\"name\":\"SetWithdrawDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allocSnapshotTimestamp\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocationMaster\",\"outputs\":[{\"internalType\":\"contractIIFRetrievableStakeWeight\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buybackClaimableNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"casher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"checkWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cliffPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pct\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"emergencyTokenRetrieve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"funder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCliffPeriod\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"claimTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"pct\",\"type\":\"uint8\"}],\"internalType\":\"structIFVestable.Cliff[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getCurrentClaimableToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getMaxPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getTotalPaymentAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalPurchased\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimable\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUnlockedToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStakeValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasCashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasOptInBuyback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasWithdrawn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestClaimTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linearVestingEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalPurchasable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTotalPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optInBuyback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paymentAmount\",\"type\":\"uint256\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"purchaserCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleTokenPurchased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_buybackClaimableNumber\",\"type\":\"uint256\"}],\"name\":\"setBuybackClaimableNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_casher\",\"type\":\"address\"}],\"name\":\"setCasher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"claimTimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"pct\",\"type\":\"uint8[]\"}],\"name\":\"setCliffPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_funder\",\"type\":\"address\"}],\"name\":\"setFunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vestingEndTime\",\"type\":\"uint256\"}],\"name\":\"setLinearVestingEndTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTotalPurchasable\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalPurchasable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minTotalPayment\",\"type\":\"uint256\"}],\"name\":\"setMinTotalPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_vestingEditableOverride\",\"type\":\"bool\"}],\"name\":\"setVestingEditable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_whitelistRootHash\",\"type\":\"bytes32\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_whitelistSetter\",\"type\":\"address\"}],\"name\":\"setWhitelistSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_withdrawDelay\",\"type\":\"uint24\"}],\"name\":\"setWithdrawDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPaymentReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalPurchased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trackId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vestingEditableOverride\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paymentAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"whitelistedPurchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDelay\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"withdrawGiveaway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawerCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200752d3803806200752d8339818101604052810190620000389190620008fd565b888a8989868686848484848a868a8c88620000686200005c6200067b60201b60201c565b6200068360201b60201c565b600180819055506000821480620000c7575060008214158015620000b95750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614155b8015620000c65750818110155b5b62000109576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001009062000aa0565b60405180910390fd5b816002819055508273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508060038190555050505080600b81905550508473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415620001c8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001bf9062000b12565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156200023b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002329062000b84565b60405180910390fd5b82421062000280576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002779062000bf6565b60405180910390fd5b426301e1853e67ffffffffffffffff16846200029d919062000c47565b10620002e0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002d79062000cf8565b60405180910390fd5b81831062000325576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200031c9062000d90565b60405180910390fd5b826312cc030067ffffffffffffffff168362000342919062000c47565b1062000385576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200037c9062000e28565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620003f8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003ef9062000e9a565b60405180910390fd5b80600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508473ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff16815250508373ffffffffffffffffffffffffffffffffffffffff166101008173ffffffffffffffffffffffffffffffffffffffff16815250508260a081815250508160c08181525050505050505050505050505050428469ffffffffffffffffffff1611806200059c5750428469ffffffffffffffffffff16111580156200059b575060008673ffffffffffffffffffffffffffffffffffffffff16635301f3ad87876040518363ffffffff1660e01b81526004016200052b92919062000ede565b60206040518083038186803b1580156200054457600080fd5b505afa15801562000559573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200057f919062000f60565b77ffffffffffffffffffffffffffffffffffffffffffffffff16115b5b620005de576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620005d59062001008565b60405180910390fd5b84601a60006101000a81548162ffffff021916908362ffffff16021790555085601960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083601960146101000a81548169ffffffffffffffffffff021916908369ffffffffffffffffffff160217905550505050505050505050506200102a565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b6000819050919050565b62000761816200074c565b81146200076d57600080fd5b50565b600081519050620007818162000756565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620007b48262000787565b9050919050565b620007c681620007a7565b8114620007d257600080fd5b50565b600081519050620007e681620007bb565b92915050565b6000620007f982620007a7565b9050919050565b6200080b81620007ec565b81146200081757600080fd5b50565b6000815190506200082b8162000800565b92915050565b60006200083e82620007a7565b9050919050565b620008508162000831565b81146200085c57600080fd5b50565b600081519050620008708162000845565b92915050565b600062ffffff82169050919050565b620008908162000876565b81146200089c57600080fd5b50565b600081519050620008b08162000885565b92915050565b600069ffffffffffffffffffff82169050919050565b620008d781620008b6565b8114620008e357600080fd5b50565b600081519050620008f781620008cc565b92915050565b6000806000806000806000806000806101408b8d03121562000924576200092362000747565b5b6000620009348d828e0162000770565b9a50506020620009478d828e01620007d5565b99505060406200095a8d828e016200081a565b98505060606200096d8d828e016200081a565b9750506080620009808d828e016200085f565b96505060a0620009938d828e016200089f565b95505060c0620009a68d828e01620008e6565b94505060e0620009b98d828e0162000770565b935050610100620009cd8d828e0162000770565b925050610120620009e18d828e0162000770565b9150509295989b9194979a5092959850565b600082825260208201905092915050565b7f7061796d656e74546f6b656e206f72206d6178546f74616c5061796d656e742060008201527f73686f756c64206e6f742062652030207768656e2073616c655072696365206960208201527f73206e6f74203000000000000000000000000000000000000000000000000000604082015250565b600062000a88604783620009f3565b915062000a958262000a04565b606082019050919050565b6000602082019050818103600083015262000abb8162000a79565b9050919050565b7f73616c65546f6b656e203d207061796d656e74546f6b656e0000000000000000600082015250565b600062000afa601883620009f3565b915062000b078262000ac2565b602082019050919050565b6000602082019050818103600083015262000b2d8162000aeb565b9050919050565b7f3078302073616c65546f6b656e00000000000000000000000000000000000000600082015250565b600062000b6c600d83620009f3565b915062000b798262000b34565b602082019050919050565b6000602082019050818103600083015262000b9f8162000b5d565b9050919050565b7f73746172742074696d657374616d7020746f6f206561726c7900000000000000600082015250565b600062000bde601983620009f3565b915062000beb8262000ba6565b602082019050919050565b6000602082019050818103600083015262000c118162000bcf565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600062000c54826200074c565b915062000c61836200074c565b92508282101562000c775762000c7662000c18565b5b828203905092915050565b7f73746172742074696d652068617320746f2062652077697468696e203120796560008201527f6172000000000000000000000000000000000000000000000000000000000000602082015250565b600062000ce0602283620009f3565b915062000ced8262000c82565b604082019050919050565b6000602082019050818103600083015262000d138162000cd1565b9050919050565b7f656e642074696d657374616d70206d757374206265206166746572207374617260008201527f742074696d657374616d70000000000000000000000000000000000000000000602082015250565b600062000d78602b83620009f3565b915062000d858262000d1a565b604082019050919050565b6000602082019050818103600083015262000dab8162000d69565b9050919050565b7f656e642074696d652068617320746f2062652077697468696e2031302079656160008201527f7273000000000000000000000000000000000000000000000000000000000000602082015250565b600062000e10602283620009f3565b915062000e1d8262000db2565b604082019050919050565b6000602082019050818103600083015262000e438162000e01565b9050919050565b7f3078302066756e64657200000000000000000000000000000000000000000000600082015250565b600062000e82600a83620009f3565b915062000e8f8262000e4a565b602082019050919050565b6000602082019050818103600083015262000eb58162000e73565b9050919050565b62000ec78162000876565b82525050565b62000ed881620008b6565b82525050565b600060408201905062000ef5600083018562000ebc565b62000f04602083018462000ecd565b9392505050565b600077ffffffffffffffffffffffffffffffffffffffffffffffff82169050919050565b62000f3a8162000f0b565b811462000f4657600080fd5b50565b60008151905062000f5a8162000f2f565b92915050565b60006020828403121562000f795762000f7862000747565b5b600062000f898482850162000f49565b91505092915050565b7f746f74616c207765696768742069732030206f6e207768696c65207573696e6760008201527f206f6c6465722074696d657374616d7000000000000000000000000000000000602082015250565b600062000ff0603083620009f3565b915062000ffd8262000f92565b604082019050919050565b60006020820190508181036000830152620010238162000fe1565b9050919050565b60805160a05160c05160e05161010051616432620010fb600039600081816113b501528181612602015281816126ec01528181612ac50152613e6201526000818161250301526125ba015260008181610bbe01528181610dea015281816113490152818161154a015281816115850152818161242b015281816128980152612de8015260008181610b5c01528181612006015281816122950152818161282d01528181612a5601528181612bab01528181612d8601526137590152600081816115260152613b8601526164326000f3fe608060405234801561001057600080fd5b50600436106103af5760003560e01c8063642d0f70116101f4578063c6632d551161011a578063d98b1f9d116100ad578063efef39a11161007c578063efef39a114610adf578063f2fde38b14610afb578063f51f96dd14610b17578063fbf4062414610b35576103af565b8063d98b1f9d14610a45578063da943cee14610a63578063dbc4b4d314610a7f578063e3af552714610aaf576103af565b8063cddfb5fd116100e9578063cddfb5fd146109cf578063d03df6dd146109eb578063d1a12d6a14610a09578063d6ca214d14610a27576103af565b8063c6632d5514610949578063c9a2ff4314610965578063ca1d209d14610995578063cc89629b146109b1576103af565b80637dd2094611610192578063961be39111610161578063961be391146108e9578063a590c84e146108f3578063a89e85221461090f578063aacc557a1461092d576103af565b80637dd209461461087357806389910cac146108915780638aae995a146108ad5780638da5cb5b146108cb576103af565b8063715018a6116101ce578063715018a6146107fd57806376ebbbc21461080757806378e97925146108255780637a11cf5b14610843576103af565b8063642d0f7014610791578063665e8acb146107c15780636713cea2146107df576103af565b80633197cbb6116102d957806347ae8109116102775780635559638f116102465780635559638f146106d05780635b2ee404146107005780635dbb4dbd146107305780635e2c19db14610761576103af565b806347ae81091461063657806348faade414610666578063497aef2814610684578063547a5eee146106b4576103af565b8063402914f5116102b3578063402914f5146105ae578063440bc7f3146105de57806345cb3dde146105fa5780634797a54814610618576103af565b80633197cbb6146105685780633ccfd60b146105865780633d9205c114610590576103af565b80631649a8a9116103515780632270e82d116103205780632270e82d146104f6578063252b99c51461051257806326c654821461052e5780633013ce291461054a576103af565b80631649a8a9146104805780631d6a4581146104b05780631efddc2f146104ce57806322524c9d146104ec576103af565b8063041ae8801161038d578063041ae8801461040c5780630acc8cd11461042a5780631101eaaf146104465780631637cde014610464576103af565b8063017287ae146103b457806301fc191c146103d25780630288a39c146103ee575b600080fd5b6103bc610b53565b6040516103c99190614283565b60405180910390f35b6103ec60048036038101906103e79190614339565b610b59565b005b6103f6610c81565b60405161040391906143b7565b60405180910390f35b610414610c96565b6040516104219190614413565b60405180910390f35b610444600480360381019061043f919061445a565b610cbc565b005b61044e610dbb565b60405161045b91906144a6565b60405180910390f35b61047e600480360381019061047991906144c1565b610dd1565b005b61049a6004803603810190610495919061450e565b61109b565b6040516104a79190614589565b60405180910390f35b6104b861111f565b6040516104c59190614283565b60405180910390f35b6104d6611125565b6040516104e391906145c9565b60405180910390f35b6104f4611141565b005b610510600480360381019061050b9190614610565b6112c4565b005b61052c6004803603810190610527919061463d565b6112e9565b005b6105486004803603810190610543919061445a565b611328565b005b610552611524565b60405161055f91906146c9565b60405180910390f35b610570611548565b60405161057d9190614283565b60405180910390f35b61058e61156c565b005b6105986116f4565b6040516105a591906147ed565b60405180910390f35b6105c860048036038101906105c3919061445a565b61177a565b6040516105d59190614283565b60405180910390f35b6105f860048036038101906105f39190614845565b611792565b005b6106026118a4565b60405161060f9190614283565b60405180910390f35b6106206118aa565b60405161062d9190614283565b60405180910390f35b610650600480360381019061064b919061445a565b6118b0565b60405161065d9190614283565b60405180910390f35b61066e611b59565b60405161067b9190614283565b60405180910390f35b61069e6004803603810190610699919061445a565b611b5f565b6040516106ab9190614283565b60405180910390f35b6106ce60048036038101906106c9919061445a565b611b77565b005b6106ea60048036038101906106e5919061445a565b611c06565b6040516106f79190614283565b60405180910390f35b61071a6004803603810190610715919061445a565b611c1e565b6040516107279190614283565b60405180910390f35b61074a6004803603810190610745919061463d565b611c36565b604051610758929190614881565b60405180910390f35b61077b6004803603810190610776919061445a565b611c77565b6040516107889190614589565b60405180910390f35b6107ab60048036038101906107a691906148aa565b611c97565b6040516107b89190614283565b60405180910390f35b6107c9611fb1565b6040516107d6919061491e565b60405180910390f35b6107e7611fd7565b6040516107f49190614283565b60405180910390f35b610805611fdd565b005b61080f611ff1565b60405161081c9190614589565b60405180910390f35b61082d612004565b60405161083a9190614283565b60405180910390f35b61085d6004803603810190610858919061445a565b612028565b60405161086a9190614283565b60405180910390f35b61087b612278565b6040516108889190614589565b60405180910390f35b6108ab60048036038101906108a691906149e5565b61228b565b005b6108b5612307565b6040516108c29190614283565b60405180910390f35b6108d361230d565b6040516108e09190614413565b60405180910390f35b6108f1612336565b005b61090d6004803603810190610908919061445a565b61278e565b005b61091761281d565b6040516109249190614283565b60405180910390f35b61094760048036038101906109429190614a92565b612823565b005b610963600480360381019061095e919061463d565b6128d2565b005b61097f600480360381019061097a919061445a565b61292b565b60405161098c9190614283565b60405180910390f35b6109af60048036038101906109aa919061463d565b6129bd565b005b6109b9612b7b565b6040516109c69190614413565b60405180910390f35b6109e960048036038101906109e4919061463d565b612ba1565b005b6109f3612c17565b604051610a009190614413565b60405180910390f35b610a11612c3d565b604051610a1e9190614ace565b60405180910390f35b610a2f612c43565b604051610a3c91906144a6565b60405180910390f35b610a4d612c59565b604051610a5a9190614283565b60405180910390f35b610a7d6004803603810190610a78919061463d565b612c5f565b005b610a996004803603810190610a94919061445a565b612cf2565b604051610aa69190614589565b60405180910390f35b610ac96004803603810190610ac4919061445a565b612d12565b604051610ad69190614283565b60405180910390f35b610af96004803603810190610af4919061463d565b612d83565b005b610b156004803603810190610b10919061445a565b612eb2565b005b610b1f612f36565b604051610b2c9190614283565b60405180910390f35b610b3d612f3c565b604051610b4a91906143b7565b60405180910390f35b600a5481565b427f00000000000000000000000000000000000000000000000000000000000000001115610bbc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb390614b46565b60405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000421115610c1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1690614bb2565b60405180910390fd5b610c31610c2a612f51565b838361109b565b610c70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6790614c1e565b60405180910390fd5b610c7c83600354612f59565b505050565b601060149054906101000a900462ffffff1681565b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610cc461308c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2b90614c8a565b60405180910390fd5b80600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f73e0366d7ceb5a21fb27f8b2aa2720e2a9dee755bd85bfb5dbccf2830753b02260405160405180910390a250565b600860009054906101000a900463ffffffff1681565b601060149054906101000a900462ffffff1662ffffff167f0000000000000000000000000000000000000000000000000000000000000000610e139190614cd9565b4211610e54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4b90614da1565b60405180910390fd5b60026001541415610e9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9190614e0d565b60405180910390fd5b60026001819055506000610eac612f51565b9050600060025414610ef3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eea90614e79565b60405180910390fd5b6000801b6016541480610f0d5750610f0c81848461109b565b5b610f4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4390614c1e565b60405180910390fd5b601160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611034576000610fa882612028565b905080601760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080601860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b600061103f8261292b565b905061104a8161310a565b600081141561108e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108590614ee5565b60405180910390fd5b5050600180819055505050565b600080846040516020016110af9190614f4d565b604051602081830303815290604052805190602001209050611115848480806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050601654836131be565b9150509392505050565b60045481565b601960149054906101000a900469ffffffffffffffffffff1681565b600061114b612f51565b905060001515600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515146111e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111d790614fb4565b60405180910390fd5b6000600a541415611226576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121d90615020565b60405180910390fd5b6001600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f91b0dcebf54773cfc86cca32698ae9cd725b3f4e407d4b83e7e19f73c54eb95460405160405180910390a250565b6112cc61308c565b80600860046101000a81548160ff02191690831515021790555050565b6112f161308c565b80600581905550807f4d893dbff365afe590ba8bf0f5d258598f5e63fde5b989dcd4902cf89d403c2f60405160405180910390a250565b61133061308c565b601060149054906101000a900462ffffff1662ffffff167f00000000000000000000000000000000000000000000000000000000000000006113729190614cd9565b42116113b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113aa90614da1565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561140c57600080fd5b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016114479190614413565b60206040518083038186803b15801561145f57600080fd5b505afa158015611473573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114979190615055565b90506114cb6114a4612f51565b828473ffffffffffffffffffffffffffffffffffffffff166131d59092919063ffffffff16565b6114d3612f51565b73ffffffffffffffffffffffffffffffffffffffff167fb1d34d1c064a5cb36c65797ef779e448eefccf2b978edbcfd206cc687c8cea5a826040516115189190614283565b60405180910390a25050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b601060149054906101000a900462ffffff1662ffffff167f00000000000000000000000000000000000000000000000000000000000000006115ae9190614cd9565b42116115ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115e690614da1565b60405180910390fd5b60026001541415611635576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161162c90614e0d565b60405180910390fd5b60026001819055506000611647612f51565b90506000600254141561168f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611686906150ce565b60405180910390fd5b600061169a8261292b565b90506116a58161310a565b60008114156116e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116e09061513a565b60405180910390fd5b505060018081905550565b6060600e805480602002602001604051908101604052809291908181526020016000905b828210156117715783829060005260206000209060020201604051806040016040529081600082015481526020016001820160009054906101000a900460ff1660ff1660ff168152505081526020019060010190611718565b50505050905090565b60176020528060005260406000206000915090505481565b601560049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166117d3612f51565b73ffffffffffffffffffffffffffffffffffffffff16148061182e57506117f861230d565b73ffffffffffffffffffffffffffffffffffffffff16611816612f51565b73ffffffffffffffffffffffffffffffffffffffff16145b61186d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611864906151cc565b60405180910390fd5b80601681905550807fc504a95022b9d01b3024e95b0a85c200d0c538f417160776512b95fcbf7b2daa60405160405180910390a250565b600b5481565b60055481565b600080601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637db85246601a60009054906101000a900462ffffff1685601960149054906101000a900469ffffffffffffffffffff166040518463ffffffff1660e01b815260040161193b939291906151ec565b60206040518083038186803b15801561195357600080fd5b505afa158015611967573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061198b9190615273565b77ffffffffffffffffffffffffffffffffffffffffffffffff1690506000601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635301f3ad601a60009054906101000a900462ffffff16601960149054906101000a900469ffffffffffffffffffff166040518363ffffffff1660e01b8152600401611a2f9291906152a0565b60206040518083038186803b158015611a4757600080fd5b505afa158015611a5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7f9190615273565b77ffffffffffffffffffffffffffffffffffffffffffffffff16905060008111611ade576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ad590615315565b60405180910390fd5b60008082670de0b6b3a764000085611af69190615335565b611b0091906153be565b9050600081601254611b129190615335565b9050670de0b6b3a76400008067ffffffffffffffff1660025483611b369190615335565b611b4091906153be565b611b4a91906153be565b92508295505050505050919050565b60145481565b60186020528060005260406000206000915090505481565b611b7f61308c565b80601560046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f7aba5fc71e3607d34203924738fbacc00b5782879615f86db108794b4bdcc95b60405160405180910390a250565b60076020528060005260406000206000915090505481565b600c6020528060005260406000206000915090505481565b600e8181548110611c4657600080fd5b90600052602060002090600202016000915090508060000154908060010160009054906101000a900460ff16905082565b60116020528060005260406000206000915054906101000a900460ff1681565b600042600d541115611d2557600b54600d54611cb391906153ef565b611cfe600c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600b5461325b565b42611d0991906153ef565b85611d149190615335565b611d1e91906153be565b9050611faa565b6000600e80549050905060008114158015611dcb575042600e600183611d4b91906153ef565b81548110611d5c57611d5b615423565b5b9060005260206000209060020201600001541180611dca575060011515600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515145b5b15611fa5576000805b828160ff161015611f695760011515600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515148015611e4457508060ff16600a5411155b15611e6e5760648260ff1688611e5a9190615335565b611e6491906153be565b9350505050611faa565b42600e8260ff1681548110611e8657611e85615423565b5b9060005260206000209060020201600001541115611ea357611f69565b600e8160ff1681548110611eba57611eb9615423565b5b906000526020600020906002020160000154600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611f5657600e8160ff1681548110611f2957611f28615423565b5b906000526020600020906002020160010160009054906101000a900460ff1682611f539190615452565b91505b8080611f6190615489565b915050611dd4565b5060008160ff161415611f8157600092505050611faa565b60648160ff1687611f929190615335565b611f9c91906153be565b92505050611faa565b839150505b9392505050565b601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60065481565b611fe561308c565b611fef6000613275565b565b600860049054906101000a900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b600080601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637db85246601a60009054906101000a900462ffffff1685601960149054906101000a900469ffffffffffffffffffff166040518463ffffffff1660e01b81526004016120b3939291906151ec565b60206040518083038186803b1580156120cb57600080fd5b505afa1580156120df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121039190615273565b77ffffffffffffffffffffffffffffffffffffffffffffffff1690506000601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635301f3ad601a60009054906101000a900462ffffff16601960149054906101000a900469ffffffffffffffffffff166040518363ffffffff1660e01b81526004016121a79291906152a0565b60206040518083038186803b1580156121bf57600080fd5b505afa1580156121d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121f79190615273565b77ffffffffffffffffffffffffffffffffffffffffffffffff16905060008111612256576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161224d90615315565b60405180910390fd5b80826012546122659190615335565b61226f91906153be565b92505050919050565b601360009054906101000a900460ff1681565b61229361308c565b7f000000000000000000000000000000000000000000000000000000000000000042106122f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122ec906154ff565b60405180910390fd5b61230184848484613339565b50505050565b60125481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16612377612f51565b73ffffffffffffffffffffffffffffffffffffffff1614806123d2575061239c61230d565b73ffffffffffffffffffffffffffffffffffffffff166123ba612f51565b73ffffffffffffffffffffffffffffffffffffffff16145b612411576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124089061556b565b60405180910390fd5b42601060149054906101000a900462ffffff1662ffffff167f00000000000000000000000000000000000000000000000000000000000000006124549190614cd9565b10612494576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161248b906155d7565b60405180910390fd5b601360009054906101000a900460ff16156124e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124db90615643565b60405180910390fd5b6001601360006101000a81548160ff02191690831515021790555060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161255a9190614413565b60206040518083038186803b15801561257257600080fd5b505afa158015612586573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125aa9190615055565b90506125fe6125b7612f51565b827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166131d59092919063ffffffff16565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016126599190614413565b60206040518083038186803b15801561267157600080fd5b505afa158015612685573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126a99190615055565b905060006126b5613711565b9050600082601254106126ca576012546126cc565b825b9050600082826126dc91906153ef565b90506127306126e9612f51565b827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166131d59092919063ffffffff16565b612738612f51565b73ffffffffffffffffffffffffffffffffffffffff167f83205c70ca31ffcb57664adecfd9894647d48665aefae1fb38bbc7ca4c1b86fb868360405161277f929190615663565b60405180910390a25050505050565b61279661308c565b80601060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167fde112653552cba8a4f696cac12b4478ce2b9b8c0e04429455a2052ec7c0412ce60405160405180910390a250565b60035481565b61282b61308c565b7f0000000000000000000000000000000000000000000000000000000000000000421061288d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612884906154ff565b60405180910390fd5b6128c68162ffffff167f00000000000000000000000000000000000000000000000000000000000000006128c19190614cd9565b613745565b6128cf8161374f565b50565b6128da61308c565b600e805490508110612921576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612918906156fe565b60405180910390fd5b80600a8190555050565b60006129b6601860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054601760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205484611c97565b9050919050565b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166129fe612f51565b73ffffffffffffffffffffffffffffffffffffffff1614612a54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a4b9061576a565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000004210612ab6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612aad906154ff565b60405180910390fd5b612b0a612ac1612f51565b30837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16613862909392919063ffffffff16565b8060126000828254612b1c9190614cd9565b92505081905550612b2b612f51565b73ffffffffffffffffffffffffffffffffffffffff167fda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd82604051612b709190614283565b60405180910390a250565b601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b612ba961308c565b7f00000000000000000000000000000000000000000000000000000000000000004210612c0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c02906154ff565b60405180910390fd5b612c14816138eb565b50565b601560049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60165481565b601560009054906101000a900463ffffffff1681565b600d5481565b612c6761308c565b60025481612c759190615335565b6006819055506004546006541015612cc2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cb990615822565b60405180910390fd5b807f851466103668359b383470efccd1760dc8caf09bc6d5a74acc7d78ac751c21e960405160405180910390a250565b60096020528060005260406000206000915054906101000a900460ff1681565b600080612d1e836118b0565b9050806003541015612d305760035490505b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481612d7b91906153ef565b915050919050565b427f00000000000000000000000000000000000000000000000000000000000000001115612de6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ddd90614b46565b60405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000421115612e49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e4090614bb2565b60405180910390fd5b6000801b60165414612e90576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e879061588e565b60405180910390fd5b6000612ea2612e9d612f51565b612d12565b9050612eae8282612f59565b5050565b612eba61308c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612f2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f2190615920565b60405180910390fd5b612f3381613275565b50565b60025481565b601a60009054906101000a900462ffffff1681565b600033905090565b8160146000828254612f6b9190614cd9565b92505081905550612f7c8282613a34565b6000600254670de0b6b3a764000067ffffffffffffffff1660076000612fa0612f51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612fe59190615335565b612fef91906153be565b90508060186000612ffe612f51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508060176000613049612f51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050565b613094612f51565b73ffffffffffffffffffffffffffffffffffffffff166130b261230d565b73ffffffffffffffffffffffffffffffffffffffff1614613108576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130ff9061598c565b60405180910390fd5b565b61311381613d18565b42600c6000613120612f51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806017600061316b612f51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546131b491906153ef565b9250508190555050565b6000826131cb8584613efe565b1490509392505050565b6132568363a9059cbb60e01b84846040516024016131f49291906159ac565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613f54565b505050565b60008183101561326b578161326d565b825b905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61334161308c565b600860049054906101000a900460ff168061335d5750600b5442105b61339c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161339390615a21565b60405180910390fd5b8181905084849050146133e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133db90615a8d565b60405180910390fd5b6000848490501161342a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161342190615af9565b60405180910390fd5b6064848490501115613471576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161346890615b65565b60405180910390fd5b600e600061347f9190614213565b600080600b548686600081811061349957613498615423565b5b90506020020135116134e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134d790615bf7565b60405180910390fd5b60005b8686905081101561365d5786868281811061350157613500615423565b5b905060200201358310613549576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161354090615c63565b60405180910390fd5b86868281811061355c5761355b615423565b5b90506020020135925084848281811061357857613577615423565b5b905060200201602081019061358d9190615caf565b826135989190615452565b9150600e60405180604001604052808989858181106135ba576135b9615423565b5b9050602002013581526020018787858181106135d9576135d8615423565b5b90506020020160208101906135ee9190615caf565b60ff1681525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010160006101000a81548160ff021916908360ff1602179055505050808061365590615cdc565b9150506134e3565b506312cc030067ffffffffffffffff168261367891906153ef565b600b54116136bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136b290615d97565b60405180910390fd5b60648160ff1614613701576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136f890615e29565b60405180910390fd5b6000600d81905550505050505050565b6000600254670de0b6b3a764000067ffffffffffffffff166014546137369190615335565b61374091906153be565b905090565b80600b8190555050565b61375761308c565b7f000000000000000000000000000000000000000000000000000000000000000042106137b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137b0906154ff565b60405180910390fd5b6309679a3667ffffffffffffffff168162ffffff161061380e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161380590615ebb565b60405180910390fd5b80601060146101000a81548162ffffff021916908362ffffff1602179055508062ffffff167f0dac7b34fe917a51c24c51f379615699f5c23ce9168eaed469bc161c9002860860405160405180910390a250565b6138e5846323b872dd60e01b85858560405160240161388393929190615edb565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613f54565b50505050565b6138f361308c565b600860049054906101000a900460ff168061390f5750600b5442105b61394e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161394590615a21565b60405180910390fd5b600b548111613992576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161398990615f84565b60405180910390fd5b6312cc030067ffffffffffffffff16816139ac91906153ef565b600b54116139ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139e690615d97565b60405180910390fd5b80600d81905550600e6000613a049190614213565b807f83d190eb78c1206ea1abb0222d475a5f70b7f63bcd534ab65e9404b39016c7a460405160405180910390a250565b60026001541415613a7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a7190614e0d565b60405180910390fd5b6002600181905550600554821015613ac7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613abe90615ff0565b60405180910390fd5b80821115613b0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b019061605c565b60405180910390fd5b8160046000828254613b1c9190614cd9565b9250508190555060006006541480613b38575060045460065410155b613b77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b6e906160c8565b60405180910390fd5b613bcb613b82612f51565b30847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16613862909392919063ffffffff16565b600060076000613bd9612f51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415613c5b576001600860008282829054906101000a900463ffffffff16613c3c91906160e8565b92506101000a81548163ffffffff021916908363ffffffff1602179055505b8160076000613c68612f51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254613cb19190614cd9565b92505081905550613cc0612f51565b73ffffffffffffffffffffffffffffffffffffffff167f2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f63283604051613d059190614283565b60405180910390a2600180819055505050565b6000811415613d5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d539061513a565b60405180910390fd5b60116000613d68612f51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16613e54576001601560008282829054906101000a900463ffffffff16613dd691906160e8565b92506101000a81548163ffffffff021916908363ffffffff160217905550600160116000613e02612f51565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b613ea6613e5f612f51565b827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166131d59092919063ffffffff16565b613eae612f51565b73ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436482604051613ef39190614283565b60405180910390a250565b60008082905060005b8451811015613f4957613f3482868381518110613f2757613f26615423565b5b602002602001015161401b565b91508080613f4190615cdc565b915050613f07565b508091505092915050565b6000613fb6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166140469092919063ffffffff16565b90506000815111156140165780806020019051810190613fd69190616137565b614015576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161400c906161d6565b60405180910390fd5b5b505050565b60008183106140335761402e828461405e565b61403e565b61403d838361405e565b5b905092915050565b60606140558484600085614075565b90509392505050565b600082600052816020526040600020905092915050565b6060824710156140ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016140b190616268565b60405180910390fd5b6140c385614189565b614102576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016140f9906162d4565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161412b919061636e565b60006040518083038185875af1925050503d8060008114614168576040519150601f19603f3d011682016040523d82523d6000602084013e61416d565b606091505b509150915061417d8282866141ac565b92505050949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b606083156141bc5782905061420c565b6000835111156141cf5782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161420391906163da565b60405180910390fd5b9392505050565b50805460008255600202906000526020600020908101906142349190614237565b50565b5b80821115614266576000808201600090556001820160006101000a81549060ff021916905550600201614238565b5090565b6000819050919050565b61427d8161426a565b82525050565b60006020820190506142986000830184614274565b92915050565b600080fd5b600080fd5b6142b18161426a565b81146142bc57600080fd5b50565b6000813590506142ce816142a8565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126142f9576142f86142d4565b5b8235905067ffffffffffffffff811115614316576143156142d9565b5b602083019150836020820283011115614332576143316142de565b5b9250929050565b6000806000604084860312156143525761435161429e565b5b6000614360868287016142bf565b935050602084013567ffffffffffffffff811115614381576143806142a3565b5b61438d868287016142e3565b92509250509250925092565b600062ffffff82169050919050565b6143b181614399565b82525050565b60006020820190506143cc60008301846143a8565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006143fd826143d2565b9050919050565b61440d816143f2565b82525050565b60006020820190506144286000830184614404565b92915050565b614437816143f2565b811461444257600080fd5b50565b6000813590506144548161442e565b92915050565b6000602082840312156144705761446f61429e565b5b600061447e84828501614445565b91505092915050565b600063ffffffff82169050919050565b6144a081614487565b82525050565b60006020820190506144bb6000830184614497565b92915050565b600080602083850312156144d8576144d761429e565b5b600083013567ffffffffffffffff8111156144f6576144f56142a3565b5b614502858286016142e3565b92509250509250929050565b6000806000604084860312156145275761452661429e565b5b600061453586828701614445565b935050602084013567ffffffffffffffff811115614556576145556142a3565b5b614562868287016142e3565b92509250509250925092565b60008115159050919050565b6145838161456e565b82525050565b600060208201905061459e600083018461457a565b92915050565b600069ffffffffffffffffffff82169050919050565b6145c3816145a4565b82525050565b60006020820190506145de60008301846145ba565b92915050565b6145ed8161456e565b81146145f857600080fd5b50565b60008135905061460a816145e4565b92915050565b6000602082840312156146265761462561429e565b5b6000614634848285016145fb565b91505092915050565b6000602082840312156146535761465261429e565b5b6000614661848285016142bf565b91505092915050565b6000819050919050565b600061468f61468a614685846143d2565b61466a565b6143d2565b9050919050565b60006146a182614674565b9050919050565b60006146b382614696565b9050919050565b6146c3816146a8565b82525050565b60006020820190506146de60008301846146ba565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6147198161426a565b82525050565b600060ff82169050919050565b6147358161471f565b82525050565b6040820160008201516147516000850182614710565b506020820151614764602085018261472c565b50505050565b6000614776838361473b565b60408301905092915050565b6000602082019050919050565b600061479a826146e4565b6147a481856146ef565b93506147af83614700565b8060005b838110156147e05781516147c7888261476a565b97506147d283614782565b9250506001810190506147b3565b5085935050505092915050565b60006020820190508181036000830152614807818461478f565b905092915050565b6000819050919050565b6148228161480f565b811461482d57600080fd5b50565b60008135905061483f81614819565b92915050565b60006020828403121561485b5761485a61429e565b5b600061486984828501614830565b91505092915050565b61487b8161471f565b82525050565b60006040820190506148966000830185614274565b6148a36020830184614872565b9392505050565b6000806000606084860312156148c3576148c261429e565b5b60006148d1868287016142bf565b93505060206148e2868287016142bf565b92505060406148f386828701614445565b9150509250925092565b600061490882614696565b9050919050565b614918816148fd565b82525050565b6000602082019050614933600083018461490f565b92915050565b60008083601f84011261494f5761494e6142d4565b5b8235905067ffffffffffffffff81111561496c5761496b6142d9565b5b602083019150836020820283011115614988576149876142de565b5b9250929050565b60008083601f8401126149a5576149a46142d4565b5b8235905067ffffffffffffffff8111156149c2576149c16142d9565b5b6020830191508360208202830111156149de576149dd6142de565b5b9250929050565b600080600080604085870312156149ff576149fe61429e565b5b600085013567ffffffffffffffff811115614a1d57614a1c6142a3565b5b614a2987828801614939565b9450945050602085013567ffffffffffffffff811115614a4c57614a4b6142a3565b5b614a588782880161498f565b925092505092959194509250565b614a6f81614399565b8114614a7a57600080fd5b50565b600081359050614a8c81614a66565b92915050565b600060208284031215614aa857614aa761429e565b5b6000614ab684828501614a7d565b91505092915050565b614ac88161480f565b82525050565b6000602082019050614ae36000830184614abf565b92915050565b600082825260208201905092915050565b7f73616c6520686173206e6f7420626567756e0000000000000000000000000000600082015250565b6000614b30601283614ae9565b9150614b3b82614afa565b602082019050919050565b60006020820190508181036000830152614b5f81614b23565b9050919050565b7f73616c65206f7665720000000000000000000000000000000000000000000000600082015250565b6000614b9c600983614ae9565b9150614ba782614b66565b602082019050919050565b60006020820190508181036000830152614bcb81614b8f565b9050919050565b7f70726f6f6620696e76616c696400000000000000000000000000000000000000600082015250565b6000614c08600d83614ae9565b9150614c1382614bd2565b602082019050919050565b60006020820190508181036000830152614c3781614bfb565b9050919050565b7f3078302066756e64657200000000000000000000000000000000000000000000600082015250565b6000614c74600a83614ae9565b9150614c7f82614c3e565b602082019050919050565b60006020820190508181036000830152614ca381614c67565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000614ce48261426a565b9150614cef8361426a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115614d2457614d23614caa565b5b828201905092915050565b7f63616e2774207769746864726177206265666f726520636c61696d206973207360008201527f7461727465640000000000000000000000000000000000000000000000000000602082015250565b6000614d8b602683614ae9565b9150614d9682614d2f565b604082019050919050565b60006020820190508181036000830152614dba81614d7e565b9050919050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6000614df7601f83614ae9565b9150614e0282614dc1565b602082019050919050565b60006020820190508181036000830152614e2681614dea565b9050919050565b7f6e6f742061206769766561776179000000000000000000000000000000000000600082015250565b6000614e63600e83614ae9565b9150614e6e82614e2d565b602082019050919050565b60006020820190508181036000830152614e9281614e56565b9050919050565b7f776974686472617720676976656177617920616d6f756e742030000000000000600082015250565b6000614ecf601a83614ae9565b9150614eda82614e99565b602082019050919050565b60006020820190508181036000830152614efe81614ec2565b9050919050565b60008160601b9050919050565b6000614f1d82614f05565b9050919050565b6000614f2f82614f12565b9050919050565b614f47614f42826143f2565b614f24565b82525050565b6000614f598284614f36565b60148201915081905092915050565b7f757365722068617320616c7265616479206f7074656420696e00000000000000600082015250565b6000614f9e601983614ae9565b9150614fa982614f68565b602082019050919050565b60006020820190508181036000830152614fcd81614f91565b9050919050565b7f6275796261636b206973206e6f7420656e61626c656400000000000000000000600082015250565b600061500a601683614ae9565b915061501582614fd4565b602082019050919050565b6000602082019050818103600083015261503981614ffd565b9050919050565b60008151905061504f816142a8565b92915050565b60006020828403121561506b5761506a61429e565b5b600061507984828501615040565b91505092915050565b7f7573652077697468647261774769766561776179000000000000000000000000600082015250565b60006150b8601483614ae9565b91506150c382615082565b602082019050919050565b600060208201905081810360008301526150e7816150ab565b9050919050565b7f6e6f20746f6b656e20746f2062652077697468647261776e0000000000000000600082015250565b6000615124601883614ae9565b915061512f826150ee565b602082019050919050565b6000602082019050818103600083015261515381615117565b9050919050565b7f63616c6c6572206e6f742077686974656c69737420736574746572206f72206f60008201527f776e657200000000000000000000000000000000000000000000000000000000602082015250565b60006151b6602483614ae9565b91506151c18261515a565b604082019050919050565b600060208201905081810360008301526151e5816151a9565b9050919050565b600060608201905061520160008301866143a8565b61520e6020830185614404565b61521b60408301846145ba565b949350505050565b600077ffffffffffffffffffffffffffffffffffffffffffffffff82169050919050565b61525081615223565b811461525b57600080fd5b50565b60008151905061526d81615247565b92915050565b6000602082840312156152895761528861429e565b5b60006152978482850161525e565b91505092915050565b60006040820190506152b560008301856143a8565b6152c260208301846145ba565b9392505050565b7f746f74616c207765696768742069732030000000000000000000000000000000600082015250565b60006152ff601183614ae9565b915061530a826152c9565b602082019050919050565b6000602082019050818103600083015261532e816152f2565b9050919050565b60006153408261426a565b915061534b8361426a565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561538457615383614caa565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006153c98261426a565b91506153d48361426a565b9250826153e4576153e361538f565b5b828204905092915050565b60006153fa8261426a565b91506154058361426a565b92508282101561541857615417614caa565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061545d8261471f565b91506154688361471f565b92508260ff0382111561547e5761547d614caa565b5b828201905092915050565b60006154948261471f565b915060ff8214156154a8576154a7614caa565b5b600182019050919050565b7f73616c6520616c72656164792073746172746564000000000000000000000000600082015250565b60006154e9601483614ae9565b91506154f4826154b3565b602082019050919050565b60006020820190508181036000830152615518816154dc565b9050919050565b7f63616c6c6572206e6f7420636173686572206f72206f776e6572000000000000600082015250565b6000615555601a83614ae9565b91506155608261551f565b602082019050919050565b6000602082019050818103600083015261558481615548565b9050919050565b7f63616e6e6f742077697468647261772079657400000000000000000000000000600082015250565b60006155c1601383614ae9565b91506155cc8261558b565b602082019050919050565b600060208201905081810360008301526155f0816155b4565b9050919050565b7f616c726561647920636173686564000000000000000000000000000000000000600082015250565b600061562d600e83614ae9565b9150615638826155f7565b602082019050919050565b6000602082019050818103600083015261565c81615620565b9050919050565b60006040820190506156786000830185614274565b6156856020830184614274565b9392505050565b7f6275796261636b20636c61696d61626c65206e756d6265722063616e6e6f742060008201527f657863656564206e756d626572206f6620636c69666620706572696f64000000602082015250565b60006156e8603d83614ae9565b91506156f38261568c565b604082019050919050565b60006020820190508181036000830152615717816156db565b9050919050565b7f63616c6c6572206e6f742066756e646572000000000000000000000000000000600082015250565b6000615754601183614ae9565b915061575f8261571e565b602082019050919050565b6000602082019050818103600083015261578381615747565b9050919050565b7f4d6178207075726368617361626c652073686f756c64206e6f74206265206c6f60008201527f776572207468616e2074686520616d6f756e74206f6620746f6b656e2070757260208201527f6368617365640000000000000000000000000000000000000000000000000000604082015250565b600061580c604683614ae9565b91506158178261578a565b606082019050919050565b6000602082019050818103600083015261583b816157ff565b9050919050565b7f7573652077686974656c69737465645075726368617365000000000000000000600082015250565b6000615878601783614ae9565b915061588382615842565b602082019050919050565b600060208201905081810360008301526158a78161586b565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061590a602683614ae9565b9150615915826158ae565b604082019050919050565b60006020820190508181036000830152615939816158fd565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000615976602083614ae9565b915061598182615940565b602082019050919050565b600060208201905081810360008301526159a581615969565b9050919050565b60006040820190506159c16000830185614404565b6159ce6020830184614274565b9392505050565b7f43616e277420656469742076657374696e672061667465722073616c65000000600082015250565b6000615a0b601d83614ae9565b9150615a16826159d5565b602082019050919050565b60006020820190508181036000830152615a3a816159fe565b9050919050565b7f646174657320616e642070637420646f65736e2774206d617463680000000000600082015250565b6000615a77601b83614ae9565b9150615a8282615a41565b602082019050919050565b60006020820190508181036000830152615aa681615a6a565b9050919050565b7f696e70757420697320656d707479000000000000000000000000000000000000600082015250565b6000615ae3600e83614ae9565b9150615aee82615aad565b602082019050919050565b60006020820190508181036000830152615b1281615ad6565b9050919050565b7f696e707574206c656e6774682063616e6e6f7420657863656564203130300000600082015250565b6000615b4f601e83614ae9565b9150615b5a82615b19565b602082019050919050565b60006020820190508181036000830152615b7e81615b42565b9050919050565b7f666972737420636c61696d2074696d65206973206265666f726520656e64207460008201527f696d65202b2077697468647261772064656c6179000000000000000000000000602082015250565b6000615be1603483614ae9565b9150615bec82615b85565b604082019050919050565b60006020820190508181036000830152615c1081615bd4565b9050919050565b7f6461746573206e6f7420696e20617363656e64696e67206f7264657200000000600082015250565b6000615c4d601c83614ae9565b9150615c5882615c17565b602082019050919050565b60006020820190508181036000830152615c7c81615c40565b9050919050565b615c8c8161471f565b8114615c9757600080fd5b50565b600081359050615ca981615c83565b92915050565b600060208284031215615cc557615cc461429e565b5b6000615cd384828501615c9a565b91505092915050565b6000615ce78261426a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415615d1a57615d19614caa565b5b600182019050919050565b7f76657374696e6720656e642074696d652068617320746f20626520776974686960008201527f6e20313020796561727300000000000000000000000000000000000000000000602082015250565b6000615d81602a83614ae9565b9150615d8c82615d25565b604082019050919050565b60006020820190508181036000830152615db081615d74565b9050919050565b7f746f74616c20696e7075742070657263656e7461676520646f65736e2774206560008201527f7175616c20746f20313030000000000000000000000000000000000000000000602082015250565b6000615e13602b83614ae9565b9150615e1e82615db7565b604082019050919050565b60006020820190508181036000830152615e4281615e06565b9050919050565b7f776974686472617744656c61792068617320746f2062652077697468696e203560008201527f2079656172730000000000000000000000000000000000000000000000000000602082015250565b6000615ea5602683614ae9565b9150615eb082615e49565b604082019050919050565b60006020820190508181036000830152615ed481615e98565b9050919050565b6000606082019050615ef06000830186614404565b615efd6020830185614404565b615f0a6040830184614274565b949350505050565b7f76657374696e6720656e642074696d652068617320746f20626520616674657260008201527f207769746864726177616c2073746172742074696d6500000000000000000000602082015250565b6000615f6e603683614ae9565b9150615f7982615f12565b604082019050919050565b60006020820190508181036000830152615f9d81615f61565b9050919050565b7f616d6f756e742062656c6f77206d696e00000000000000000000000000000000600082015250565b6000615fda601083614ae9565b9150615fe582615fa4565b602082019050919050565b6000602082019050818103600083015261600981615fcd565b9050919050565b7f65786365656473206d6178207061796d656e7400000000000000000000000000600082015250565b6000616046601383614ae9565b915061605182616010565b602082019050919050565b6000602082019050818103600083015261607581616039565b9050919050565b7f657863656564206d6178207075726368617361626c6500000000000000000000600082015250565b60006160b2601683614ae9565b91506160bd8261607c565b602082019050919050565b600060208201905081810360008301526160e1816160a5565b9050919050565b60006160f382614487565b91506160fe83614487565b92508263ffffffff0382111561611757616116614caa565b5b828201905092915050565b600081519050616131816145e4565b92915050565b60006020828403121561614d5761614c61429e565b5b600061615b84828501616122565b91505092915050565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b60006161c0602a83614ae9565b91506161cb82616164565b604082019050919050565b600060208201905081810360008301526161ef816161b3565b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b6000616252602683614ae9565b915061625d826161f6565b604082019050919050565b6000602082019050818103600083015261628181616245565b9050919050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b60006162be601d83614ae9565b91506162c982616288565b602082019050919050565b600060208201905081810360008301526162ed816162b1565b9050919050565b600081519050919050565b600081905092915050565b60005b8381101561632857808201518184015260208101905061630d565b83811115616337576000848401525b50505050565b6000616348826162f4565b61635281856162ff565b935061636281856020860161630a565b80840191505092915050565b600061637a828461633d565b915081905092915050565b600081519050919050565b6000601f19601f8301169050919050565b60006163ac82616385565b6163b68185614ae9565b93506163c681856020860161630a565b6163cf81616390565b840191505092915050565b600060208201905081810360008301526163f481846163a1565b90509291505056fea2646970667358221220233193ee41a1893da015b934d737a7bab220ac4f64b50990b45015ed71d7fb5a64736f6c63430008090033",
}

// IFAllocationSaleV8ABI is the input ABI used to generate the binding from.
// Deprecated: Use IFAllocationSaleV8MetaData.ABI instead.
var IFAllocationSaleV8ABI = IFAllocationSaleV8MetaData.ABI

// IFAllocationSaleV8Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IFAllocationSaleV8MetaData.Bin instead.
var IFAllocationSaleV8Bin = IFAllocationSaleV8MetaData.Bin

// DeployIFAllocationSaleV8 deploys a new Ethereum contract, binding an instance of IFAllocationSaleV8 to it.
func DeployIFAllocationSaleV8(auth *bind.TransactOpts, backend bind.ContractBackend, _salePrice *big.Int, _funder common.Address, _paymentToken common.Address, _saleToken common.Address, _allocationMaster common.Address, _trackId *big.Int, _allocSnapshotTimestamp *big.Int, _startTime *big.Int, _endTime *big.Int, _maxTotalPayment *big.Int) (common.Address, *types.Transaction, *IFAllocationSaleV8, error) {
	parsed, err := IFAllocationSaleV8MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IFAllocationSaleV8Bin), backend, _salePrice, _funder, _paymentToken, _saleToken, _allocationMaster, _trackId, _allocSnapshotTimestamp, _startTime, _endTime, _maxTotalPayment)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IFAllocationSaleV8{IFAllocationSaleV8Caller: IFAllocationSaleV8Caller{contract: contract}, IFAllocationSaleV8Transactor: IFAllocationSaleV8Transactor{contract: contract}, IFAllocationSaleV8Filterer: IFAllocationSaleV8Filterer{contract: contract}}, nil
}

// IFAllocationSaleV8 is an auto generated Go binding around an Ethereum contract.
type IFAllocationSaleV8 struct {
	IFAllocationSaleV8Caller     // Read-only binding to the contract
	IFAllocationSaleV8Transactor // Write-only binding to the contract
	IFAllocationSaleV8Filterer   // Log filterer for contract events
}

// IFAllocationSaleV8Caller is an auto generated read-only Go binding around an Ethereum contract.
type IFAllocationSaleV8Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFAllocationSaleV8Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IFAllocationSaleV8Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFAllocationSaleV8Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFAllocationSaleV8Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFAllocationSaleV8Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFAllocationSaleV8Session struct {
	Contract     *IFAllocationSaleV8 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IFAllocationSaleV8CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFAllocationSaleV8CallerSession struct {
	Contract *IFAllocationSaleV8Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IFAllocationSaleV8TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFAllocationSaleV8TransactorSession struct {
	Contract     *IFAllocationSaleV8Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IFAllocationSaleV8Raw is an auto generated low-level Go binding around an Ethereum contract.
type IFAllocationSaleV8Raw struct {
	Contract *IFAllocationSaleV8 // Generic contract binding to access the raw methods on
}

// IFAllocationSaleV8CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFAllocationSaleV8CallerRaw struct {
	Contract *IFAllocationSaleV8Caller // Generic read-only contract binding to access the raw methods on
}

// IFAllocationSaleV8TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFAllocationSaleV8TransactorRaw struct {
	Contract *IFAllocationSaleV8Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIFAllocationSaleV8 creates a new instance of IFAllocationSaleV8, bound to a specific deployed contract.
func NewIFAllocationSaleV8(address common.Address, backend bind.ContractBackend) (*IFAllocationSaleV8, error) {
	contract, err := bindIFAllocationSaleV8(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8{IFAllocationSaleV8Caller: IFAllocationSaleV8Caller{contract: contract}, IFAllocationSaleV8Transactor: IFAllocationSaleV8Transactor{contract: contract}, IFAllocationSaleV8Filterer: IFAllocationSaleV8Filterer{contract: contract}}, nil
}

// NewIFAllocationSaleV8Caller creates a new read-only instance of IFAllocationSaleV8, bound to a specific deployed contract.
func NewIFAllocationSaleV8Caller(address common.Address, caller bind.ContractCaller) (*IFAllocationSaleV8Caller, error) {
	contract, err := bindIFAllocationSaleV8(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8Caller{contract: contract}, nil
}

// NewIFAllocationSaleV8Transactor creates a new write-only instance of IFAllocationSaleV8, bound to a specific deployed contract.
func NewIFAllocationSaleV8Transactor(address common.Address, transactor bind.ContractTransactor) (*IFAllocationSaleV8Transactor, error) {
	contract, err := bindIFAllocationSaleV8(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8Transactor{contract: contract}, nil
}

// NewIFAllocationSaleV8Filterer creates a new log filterer instance of IFAllocationSaleV8, bound to a specific deployed contract.
func NewIFAllocationSaleV8Filterer(address common.Address, filterer bind.ContractFilterer) (*IFAllocationSaleV8Filterer, error) {
	contract, err := bindIFAllocationSaleV8(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8Filterer{contract: contract}, nil
}

// bindIFAllocationSaleV8 binds a generic wrapper to an already deployed contract.
func bindIFAllocationSaleV8(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFAllocationSaleV8MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFAllocationSaleV8 *IFAllocationSaleV8Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFAllocationSaleV8.Contract.IFAllocationSaleV8Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFAllocationSaleV8 *IFAllocationSaleV8Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.IFAllocationSaleV8Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFAllocationSaleV8 *IFAllocationSaleV8Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.IFAllocationSaleV8Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFAllocationSaleV8.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.contract.Transact(opts, method, params...)
}

// AllocSnapshotTimestamp is a free data retrieval call binding the contract method 0x1efddc2f.
//
// Solidity: function allocSnapshotTimestamp() view returns(uint80)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) AllocSnapshotTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "allocSnapshotTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocSnapshotTimestamp is a free data retrieval call binding the contract method 0x1efddc2f.
//
// Solidity: function allocSnapshotTimestamp() view returns(uint80)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) AllocSnapshotTimestamp() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.AllocSnapshotTimestamp(&_IFAllocationSaleV8.CallOpts)
}

// AllocSnapshotTimestamp is a free data retrieval call binding the contract method 0x1efddc2f.
//
// Solidity: function allocSnapshotTimestamp() view returns(uint80)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) AllocSnapshotTimestamp() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.AllocSnapshotTimestamp(&_IFAllocationSaleV8.CallOpts)
}

// AllocationMaster is a free data retrieval call binding the contract method 0x665e8acb.
//
// Solidity: function allocationMaster() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) AllocationMaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "allocationMaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationMaster is a free data retrieval call binding the contract method 0x665e8acb.
//
// Solidity: function allocationMaster() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) AllocationMaster() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.AllocationMaster(&_IFAllocationSaleV8.CallOpts)
}

// AllocationMaster is a free data retrieval call binding the contract method 0x665e8acb.
//
// Solidity: function allocationMaster() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) AllocationMaster() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.AllocationMaster(&_IFAllocationSaleV8.CallOpts)
}

// BuybackClaimableNumber is a free data retrieval call binding the contract method 0x017287ae.
//
// Solidity: function buybackClaimableNumber() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) BuybackClaimableNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "buybackClaimableNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuybackClaimableNumber is a free data retrieval call binding the contract method 0x017287ae.
//
// Solidity: function buybackClaimableNumber() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) BuybackClaimableNumber() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.BuybackClaimableNumber(&_IFAllocationSaleV8.CallOpts)
}

// BuybackClaimableNumber is a free data retrieval call binding the contract method 0x017287ae.
//
// Solidity: function buybackClaimableNumber() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) BuybackClaimableNumber() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.BuybackClaimableNumber(&_IFAllocationSaleV8.CallOpts)
}

// Casher is a free data retrieval call binding the contract method 0xcc89629b.
//
// Solidity: function casher() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) Casher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "casher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Casher is a free data retrieval call binding the contract method 0xcc89629b.
//
// Solidity: function casher() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) Casher() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.Casher(&_IFAllocationSaleV8.CallOpts)
}

// Casher is a free data retrieval call binding the contract method 0xcc89629b.
//
// Solidity: function casher() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) Casher() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.Casher(&_IFAllocationSaleV8.CallOpts)
}

// CheckWhitelist is a free data retrieval call binding the contract method 0x1649a8a9.
//
// Solidity: function checkWhitelist(address user, bytes32[] merkleProof) view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) CheckWhitelist(opts *bind.CallOpts, user common.Address, merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "checkWhitelist", user, merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckWhitelist is a free data retrieval call binding the contract method 0x1649a8a9.
//
// Solidity: function checkWhitelist(address user, bytes32[] merkleProof) view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) CheckWhitelist(user common.Address, merkleProof [][32]byte) (bool, error) {
	return _IFAllocationSaleV8.Contract.CheckWhitelist(&_IFAllocationSaleV8.CallOpts, user, merkleProof)
}

// CheckWhitelist is a free data retrieval call binding the contract method 0x1649a8a9.
//
// Solidity: function checkWhitelist(address user, bytes32[] merkleProof) view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) CheckWhitelist(user common.Address, merkleProof [][32]byte) (bool, error) {
	return _IFAllocationSaleV8.Contract.CheckWhitelist(&_IFAllocationSaleV8.CallOpts, user, merkleProof)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) Claimable(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "claimable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) Claimable(arg0 common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.Claimable(&_IFAllocationSaleV8.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.Claimable(&_IFAllocationSaleV8.CallOpts, arg0)
}

// CliffPeriod is a free data retrieval call binding the contract method 0x5dbb4dbd.
//
// Solidity: function cliffPeriod(uint256 ) view returns(uint256 claimTime, uint8 pct)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) CliffPeriod(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ClaimTime *big.Int
	Pct       uint8
}, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "cliffPeriod", arg0)

	outstruct := new(struct {
		ClaimTime *big.Int
		Pct       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ClaimTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pct = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// CliffPeriod is a free data retrieval call binding the contract method 0x5dbb4dbd.
//
// Solidity: function cliffPeriod(uint256 ) view returns(uint256 claimTime, uint8 pct)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) CliffPeriod(arg0 *big.Int) (struct {
	ClaimTime *big.Int
	Pct       uint8
}, error) {
	return _IFAllocationSaleV8.Contract.CliffPeriod(&_IFAllocationSaleV8.CallOpts, arg0)
}

// CliffPeriod is a free data retrieval call binding the contract method 0x5dbb4dbd.
//
// Solidity: function cliffPeriod(uint256 ) view returns(uint256 claimTime, uint8 pct)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) CliffPeriod(arg0 *big.Int) (struct {
	ClaimTime *big.Int
	Pct       uint8
}, error) {
	return _IFAllocationSaleV8.Contract.CliffPeriod(&_IFAllocationSaleV8.CallOpts, arg0)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) EndTime() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.EndTime(&_IFAllocationSaleV8.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) EndTime() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.EndTime(&_IFAllocationSaleV8.CallOpts)
}

// Funder is a free data retrieval call binding the contract method 0x041ae880.
//
// Solidity: function funder() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) Funder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "funder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Funder is a free data retrieval call binding the contract method 0x041ae880.
//
// Solidity: function funder() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) Funder() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.Funder(&_IFAllocationSaleV8.CallOpts)
}

// Funder is a free data retrieval call binding the contract method 0x041ae880.
//
// Solidity: function funder() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) Funder() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.Funder(&_IFAllocationSaleV8.CallOpts)
}

// GetCliffPeriod is a free data retrieval call binding the contract method 0x3d9205c1.
//
// Solidity: function getCliffPeriod() view returns((uint256,uint8)[])
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) GetCliffPeriod(opts *bind.CallOpts) ([]IFVestableCliff, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "getCliffPeriod")

	if err != nil {
		return *new([]IFVestableCliff), err
	}

	out0 := *abi.ConvertType(out[0], new([]IFVestableCliff)).(*[]IFVestableCliff)

	return out0, err

}

// GetCliffPeriod is a free data retrieval call binding the contract method 0x3d9205c1.
//
// Solidity: function getCliffPeriod() view returns((uint256,uint8)[])
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) GetCliffPeriod() ([]IFVestableCliff, error) {
	return _IFAllocationSaleV8.Contract.GetCliffPeriod(&_IFAllocationSaleV8.CallOpts)
}

// GetCliffPeriod is a free data retrieval call binding the contract method 0x3d9205c1.
//
// Solidity: function getCliffPeriod() view returns((uint256,uint8)[])
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) GetCliffPeriod() ([]IFVestableCliff, error) {
	return _IFAllocationSaleV8.Contract.GetCliffPeriod(&_IFAllocationSaleV8.CallOpts)
}

// GetCurrentClaimableToken is a free data retrieval call binding the contract method 0xc9a2ff43.
//
// Solidity: function getCurrentClaimableToken(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) GetCurrentClaimableToken(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "getCurrentClaimableToken", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentClaimableToken is a free data retrieval call binding the contract method 0xc9a2ff43.
//
// Solidity: function getCurrentClaimableToken(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) GetCurrentClaimableToken(user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetCurrentClaimableToken(&_IFAllocationSaleV8.CallOpts, user)
}

// GetCurrentClaimableToken is a free data retrieval call binding the contract method 0xc9a2ff43.
//
// Solidity: function getCurrentClaimableToken(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) GetCurrentClaimableToken(user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetCurrentClaimableToken(&_IFAllocationSaleV8.CallOpts, user)
}

// GetMaxPayment is a free data retrieval call binding the contract method 0xe3af5527.
//
// Solidity: function getMaxPayment(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) GetMaxPayment(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "getMaxPayment", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPayment is a free data retrieval call binding the contract method 0xe3af5527.
//
// Solidity: function getMaxPayment(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) GetMaxPayment(user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetMaxPayment(&_IFAllocationSaleV8.CallOpts, user)
}

// GetMaxPayment is a free data retrieval call binding the contract method 0xe3af5527.
//
// Solidity: function getMaxPayment(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) GetMaxPayment(user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetMaxPayment(&_IFAllocationSaleV8.CallOpts, user)
}

// GetTotalPaymentAllocation is a free data retrieval call binding the contract method 0x47ae8109.
//
// Solidity: function getTotalPaymentAllocation(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) GetTotalPaymentAllocation(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "getTotalPaymentAllocation", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPaymentAllocation is a free data retrieval call binding the contract method 0x47ae8109.
//
// Solidity: function getTotalPaymentAllocation(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) GetTotalPaymentAllocation(user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetTotalPaymentAllocation(&_IFAllocationSaleV8.CallOpts, user)
}

// GetTotalPaymentAllocation is a free data retrieval call binding the contract method 0x47ae8109.
//
// Solidity: function getTotalPaymentAllocation(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) GetTotalPaymentAllocation(user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetTotalPaymentAllocation(&_IFAllocationSaleV8.CallOpts, user)
}

// GetUnlockedToken is a free data retrieval call binding the contract method 0x642d0f70.
//
// Solidity: function getUnlockedToken(uint256 totalPurchased, uint256 claimable, address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) GetUnlockedToken(opts *bind.CallOpts, totalPurchased *big.Int, claimable *big.Int, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "getUnlockedToken", totalPurchased, claimable, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnlockedToken is a free data retrieval call binding the contract method 0x642d0f70.
//
// Solidity: function getUnlockedToken(uint256 totalPurchased, uint256 claimable, address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) GetUnlockedToken(totalPurchased *big.Int, claimable *big.Int, user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetUnlockedToken(&_IFAllocationSaleV8.CallOpts, totalPurchased, claimable, user)
}

// GetUnlockedToken is a free data retrieval call binding the contract method 0x642d0f70.
//
// Solidity: function getUnlockedToken(uint256 totalPurchased, uint256 claimable, address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) GetUnlockedToken(totalPurchased *big.Int, claimable *big.Int, user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetUnlockedToken(&_IFAllocationSaleV8.CallOpts, totalPurchased, claimable, user)
}

// GetUserStakeValue is a free data retrieval call binding the contract method 0x7a11cf5b.
//
// Solidity: function getUserStakeValue(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) GetUserStakeValue(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "getUserStakeValue", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserStakeValue is a free data retrieval call binding the contract method 0x7a11cf5b.
//
// Solidity: function getUserStakeValue(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) GetUserStakeValue(user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetUserStakeValue(&_IFAllocationSaleV8.CallOpts, user)
}

// GetUserStakeValue is a free data retrieval call binding the contract method 0x7a11cf5b.
//
// Solidity: function getUserStakeValue(address user) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) GetUserStakeValue(user common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.GetUserStakeValue(&_IFAllocationSaleV8.CallOpts, user)
}

// HasCashed is a free data retrieval call binding the contract method 0x7dd20946.
//
// Solidity: function hasCashed() view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) HasCashed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "hasCashed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCashed is a free data retrieval call binding the contract method 0x7dd20946.
//
// Solidity: function hasCashed() view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) HasCashed() (bool, error) {
	return _IFAllocationSaleV8.Contract.HasCashed(&_IFAllocationSaleV8.CallOpts)
}

// HasCashed is a free data retrieval call binding the contract method 0x7dd20946.
//
// Solidity: function hasCashed() view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) HasCashed() (bool, error) {
	return _IFAllocationSaleV8.Contract.HasCashed(&_IFAllocationSaleV8.CallOpts)
}

// HasOptInBuyback is a free data retrieval call binding the contract method 0xdbc4b4d3.
//
// Solidity: function hasOptInBuyback(address ) view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) HasOptInBuyback(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "hasOptInBuyback", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOptInBuyback is a free data retrieval call binding the contract method 0xdbc4b4d3.
//
// Solidity: function hasOptInBuyback(address ) view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) HasOptInBuyback(arg0 common.Address) (bool, error) {
	return _IFAllocationSaleV8.Contract.HasOptInBuyback(&_IFAllocationSaleV8.CallOpts, arg0)
}

// HasOptInBuyback is a free data retrieval call binding the contract method 0xdbc4b4d3.
//
// Solidity: function hasOptInBuyback(address ) view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) HasOptInBuyback(arg0 common.Address) (bool, error) {
	return _IFAllocationSaleV8.Contract.HasOptInBuyback(&_IFAllocationSaleV8.CallOpts, arg0)
}

// HasWithdrawn is a free data retrieval call binding the contract method 0x5e2c19db.
//
// Solidity: function hasWithdrawn(address ) view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) HasWithdrawn(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "hasWithdrawn", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasWithdrawn is a free data retrieval call binding the contract method 0x5e2c19db.
//
// Solidity: function hasWithdrawn(address ) view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) HasWithdrawn(arg0 common.Address) (bool, error) {
	return _IFAllocationSaleV8.Contract.HasWithdrawn(&_IFAllocationSaleV8.CallOpts, arg0)
}

// HasWithdrawn is a free data retrieval call binding the contract method 0x5e2c19db.
//
// Solidity: function hasWithdrawn(address ) view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) HasWithdrawn(arg0 common.Address) (bool, error) {
	return _IFAllocationSaleV8.Contract.HasWithdrawn(&_IFAllocationSaleV8.CallOpts, arg0)
}

// LatestClaimTime is a free data retrieval call binding the contract method 0x5b2ee404.
//
// Solidity: function latestClaimTime(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) LatestClaimTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "latestClaimTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestClaimTime is a free data retrieval call binding the contract method 0x5b2ee404.
//
// Solidity: function latestClaimTime(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) LatestClaimTime(arg0 common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.LatestClaimTime(&_IFAllocationSaleV8.CallOpts, arg0)
}

// LatestClaimTime is a free data retrieval call binding the contract method 0x5b2ee404.
//
// Solidity: function latestClaimTime(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) LatestClaimTime(arg0 common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.LatestClaimTime(&_IFAllocationSaleV8.CallOpts, arg0)
}

// LinearVestingEndTime is a free data retrieval call binding the contract method 0xd98b1f9d.
//
// Solidity: function linearVestingEndTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) LinearVestingEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "linearVestingEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LinearVestingEndTime is a free data retrieval call binding the contract method 0xd98b1f9d.
//
// Solidity: function linearVestingEndTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) LinearVestingEndTime() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.LinearVestingEndTime(&_IFAllocationSaleV8.CallOpts)
}

// LinearVestingEndTime is a free data retrieval call binding the contract method 0xd98b1f9d.
//
// Solidity: function linearVestingEndTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) LinearVestingEndTime() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.LinearVestingEndTime(&_IFAllocationSaleV8.CallOpts)
}

// MaxTotalPayment is a free data retrieval call binding the contract method 0xa89e8522.
//
// Solidity: function maxTotalPayment() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) MaxTotalPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "maxTotalPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalPayment is a free data retrieval call binding the contract method 0xa89e8522.
//
// Solidity: function maxTotalPayment() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) MaxTotalPayment() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.MaxTotalPayment(&_IFAllocationSaleV8.CallOpts)
}

// MaxTotalPayment is a free data retrieval call binding the contract method 0xa89e8522.
//
// Solidity: function maxTotalPayment() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) MaxTotalPayment() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.MaxTotalPayment(&_IFAllocationSaleV8.CallOpts)
}

// MaxTotalPurchasable is a free data retrieval call binding the contract method 0x6713cea2.
//
// Solidity: function maxTotalPurchasable() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) MaxTotalPurchasable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "maxTotalPurchasable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalPurchasable is a free data retrieval call binding the contract method 0x6713cea2.
//
// Solidity: function maxTotalPurchasable() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) MaxTotalPurchasable() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.MaxTotalPurchasable(&_IFAllocationSaleV8.CallOpts)
}

// MaxTotalPurchasable is a free data retrieval call binding the contract method 0x6713cea2.
//
// Solidity: function maxTotalPurchasable() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) MaxTotalPurchasable() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.MaxTotalPurchasable(&_IFAllocationSaleV8.CallOpts)
}

// MinTotalPayment is a free data retrieval call binding the contract method 0x4797a548.
//
// Solidity: function minTotalPayment() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) MinTotalPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "minTotalPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTotalPayment is a free data retrieval call binding the contract method 0x4797a548.
//
// Solidity: function minTotalPayment() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) MinTotalPayment() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.MinTotalPayment(&_IFAllocationSaleV8.CallOpts)
}

// MinTotalPayment is a free data retrieval call binding the contract method 0x4797a548.
//
// Solidity: function minTotalPayment() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) MinTotalPayment() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.MinTotalPayment(&_IFAllocationSaleV8.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) Owner() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.Owner(&_IFAllocationSaleV8.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) Owner() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.Owner(&_IFAllocationSaleV8.CallOpts)
}

// PaymentReceived is a free data retrieval call binding the contract method 0x5559638f.
//
// Solidity: function paymentReceived(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) PaymentReceived(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "paymentReceived", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaymentReceived is a free data retrieval call binding the contract method 0x5559638f.
//
// Solidity: function paymentReceived(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) PaymentReceived(arg0 common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.PaymentReceived(&_IFAllocationSaleV8.CallOpts, arg0)
}

// PaymentReceived is a free data retrieval call binding the contract method 0x5559638f.
//
// Solidity: function paymentReceived(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) PaymentReceived(arg0 common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.PaymentReceived(&_IFAllocationSaleV8.CallOpts, arg0)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) PaymentToken() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.PaymentToken(&_IFAllocationSaleV8.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) PaymentToken() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.PaymentToken(&_IFAllocationSaleV8.CallOpts)
}

// PurchaserCount is a free data retrieval call binding the contract method 0x1101eaaf.
//
// Solidity: function purchaserCount() view returns(uint32)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) PurchaserCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "purchaserCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PurchaserCount is a free data retrieval call binding the contract method 0x1101eaaf.
//
// Solidity: function purchaserCount() view returns(uint32)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) PurchaserCount() (uint32, error) {
	return _IFAllocationSaleV8.Contract.PurchaserCount(&_IFAllocationSaleV8.CallOpts)
}

// PurchaserCount is a free data retrieval call binding the contract method 0x1101eaaf.
//
// Solidity: function purchaserCount() view returns(uint32)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) PurchaserCount() (uint32, error) {
	return _IFAllocationSaleV8.Contract.PurchaserCount(&_IFAllocationSaleV8.CallOpts)
}

// SaleAmount is a free data retrieval call binding the contract method 0x8aae995a.
//
// Solidity: function saleAmount() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) SaleAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "saleAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SaleAmount is a free data retrieval call binding the contract method 0x8aae995a.
//
// Solidity: function saleAmount() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SaleAmount() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.SaleAmount(&_IFAllocationSaleV8.CallOpts)
}

// SaleAmount is a free data retrieval call binding the contract method 0x8aae995a.
//
// Solidity: function saleAmount() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) SaleAmount() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.SaleAmount(&_IFAllocationSaleV8.CallOpts)
}

// SalePrice is a free data retrieval call binding the contract method 0xf51f96dd.
//
// Solidity: function salePrice() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) SalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "salePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SalePrice is a free data retrieval call binding the contract method 0xf51f96dd.
//
// Solidity: function salePrice() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SalePrice() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.SalePrice(&_IFAllocationSaleV8.CallOpts)
}

// SalePrice is a free data retrieval call binding the contract method 0xf51f96dd.
//
// Solidity: function salePrice() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) SalePrice() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.SalePrice(&_IFAllocationSaleV8.CallOpts)
}

// SaleTokenPurchased is a free data retrieval call binding the contract method 0x1d6a4581.
//
// Solidity: function saleTokenPurchased() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) SaleTokenPurchased(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "saleTokenPurchased")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SaleTokenPurchased is a free data retrieval call binding the contract method 0x1d6a4581.
//
// Solidity: function saleTokenPurchased() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SaleTokenPurchased() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.SaleTokenPurchased(&_IFAllocationSaleV8.CallOpts)
}

// SaleTokenPurchased is a free data retrieval call binding the contract method 0x1d6a4581.
//
// Solidity: function saleTokenPurchased() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) SaleTokenPurchased() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.SaleTokenPurchased(&_IFAllocationSaleV8.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) StartTime() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.StartTime(&_IFAllocationSaleV8.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) StartTime() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.StartTime(&_IFAllocationSaleV8.CallOpts)
}

// TotalPaymentReceived is a free data retrieval call binding the contract method 0x48faade4.
//
// Solidity: function totalPaymentReceived() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) TotalPaymentReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "totalPaymentReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPaymentReceived is a free data retrieval call binding the contract method 0x48faade4.
//
// Solidity: function totalPaymentReceived() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) TotalPaymentReceived() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.TotalPaymentReceived(&_IFAllocationSaleV8.CallOpts)
}

// TotalPaymentReceived is a free data retrieval call binding the contract method 0x48faade4.
//
// Solidity: function totalPaymentReceived() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) TotalPaymentReceived() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.TotalPaymentReceived(&_IFAllocationSaleV8.CallOpts)
}

// TotalPurchased is a free data retrieval call binding the contract method 0x497aef28.
//
// Solidity: function totalPurchased(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) TotalPurchased(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "totalPurchased", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPurchased is a free data retrieval call binding the contract method 0x497aef28.
//
// Solidity: function totalPurchased(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) TotalPurchased(arg0 common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.TotalPurchased(&_IFAllocationSaleV8.CallOpts, arg0)
}

// TotalPurchased is a free data retrieval call binding the contract method 0x497aef28.
//
// Solidity: function totalPurchased(address ) view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) TotalPurchased(arg0 common.Address) (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.TotalPurchased(&_IFAllocationSaleV8.CallOpts, arg0)
}

// TrackId is a free data retrieval call binding the contract method 0xfbf40624.
//
// Solidity: function trackId() view returns(uint24)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) TrackId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "trackId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrackId is a free data retrieval call binding the contract method 0xfbf40624.
//
// Solidity: function trackId() view returns(uint24)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) TrackId() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.TrackId(&_IFAllocationSaleV8.CallOpts)
}

// TrackId is a free data retrieval call binding the contract method 0xfbf40624.
//
// Solidity: function trackId() view returns(uint24)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) TrackId() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.TrackId(&_IFAllocationSaleV8.CallOpts)
}

// VestingEditableOverride is a free data retrieval call binding the contract method 0x76ebbbc2.
//
// Solidity: function vestingEditableOverride() view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) VestingEditableOverride(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "vestingEditableOverride")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VestingEditableOverride is a free data retrieval call binding the contract method 0x76ebbbc2.
//
// Solidity: function vestingEditableOverride() view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) VestingEditableOverride() (bool, error) {
	return _IFAllocationSaleV8.Contract.VestingEditableOverride(&_IFAllocationSaleV8.CallOpts)
}

// VestingEditableOverride is a free data retrieval call binding the contract method 0x76ebbbc2.
//
// Solidity: function vestingEditableOverride() view returns(bool)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) VestingEditableOverride() (bool, error) {
	return _IFAllocationSaleV8.Contract.VestingEditableOverride(&_IFAllocationSaleV8.CallOpts)
}

// WhitelistRootHash is a free data retrieval call binding the contract method 0xd1a12d6a.
//
// Solidity: function whitelistRootHash() view returns(bytes32)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) WhitelistRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "whitelistRootHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WhitelistRootHash is a free data retrieval call binding the contract method 0xd1a12d6a.
//
// Solidity: function whitelistRootHash() view returns(bytes32)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) WhitelistRootHash() ([32]byte, error) {
	return _IFAllocationSaleV8.Contract.WhitelistRootHash(&_IFAllocationSaleV8.CallOpts)
}

// WhitelistRootHash is a free data retrieval call binding the contract method 0xd1a12d6a.
//
// Solidity: function whitelistRootHash() view returns(bytes32)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) WhitelistRootHash() ([32]byte, error) {
	return _IFAllocationSaleV8.Contract.WhitelistRootHash(&_IFAllocationSaleV8.CallOpts)
}

// WhitelistSetter is a free data retrieval call binding the contract method 0xd03df6dd.
//
// Solidity: function whitelistSetter() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) WhitelistSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "whitelistSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistSetter is a free data retrieval call binding the contract method 0xd03df6dd.
//
// Solidity: function whitelistSetter() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) WhitelistSetter() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.WhitelistSetter(&_IFAllocationSaleV8.CallOpts)
}

// WhitelistSetter is a free data retrieval call binding the contract method 0xd03df6dd.
//
// Solidity: function whitelistSetter() view returns(address)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) WhitelistSetter() (common.Address, error) {
	return _IFAllocationSaleV8.Contract.WhitelistSetter(&_IFAllocationSaleV8.CallOpts)
}

// WithdrawDelay is a free data retrieval call binding the contract method 0x0288a39c.
//
// Solidity: function withdrawDelay() view returns(uint24)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) WithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "withdrawDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawDelay is a free data retrieval call binding the contract method 0x0288a39c.
//
// Solidity: function withdrawDelay() view returns(uint24)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) WithdrawDelay() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.WithdrawDelay(&_IFAllocationSaleV8.CallOpts)
}

// WithdrawDelay is a free data retrieval call binding the contract method 0x0288a39c.
//
// Solidity: function withdrawDelay() view returns(uint24)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) WithdrawDelay() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.WithdrawDelay(&_IFAllocationSaleV8.CallOpts)
}

// WithdrawTime is a free data retrieval call binding the contract method 0x45cb3dde.
//
// Solidity: function withdrawTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) WithdrawTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "withdrawTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawTime is a free data retrieval call binding the contract method 0x45cb3dde.
//
// Solidity: function withdrawTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) WithdrawTime() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.WithdrawTime(&_IFAllocationSaleV8.CallOpts)
}

// WithdrawTime is a free data retrieval call binding the contract method 0x45cb3dde.
//
// Solidity: function withdrawTime() view returns(uint256)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) WithdrawTime() (*big.Int, error) {
	return _IFAllocationSaleV8.Contract.WithdrawTime(&_IFAllocationSaleV8.CallOpts)
}

// WithdrawerCount is a free data retrieval call binding the contract method 0xd6ca214d.
//
// Solidity: function withdrawerCount() view returns(uint32)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Caller) WithdrawerCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IFAllocationSaleV8.contract.Call(opts, &out, "withdrawerCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// WithdrawerCount is a free data retrieval call binding the contract method 0xd6ca214d.
//
// Solidity: function withdrawerCount() view returns(uint32)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) WithdrawerCount() (uint32, error) {
	return _IFAllocationSaleV8.Contract.WithdrawerCount(&_IFAllocationSaleV8.CallOpts)
}

// WithdrawerCount is a free data retrieval call binding the contract method 0xd6ca214d.
//
// Solidity: function withdrawerCount() view returns(uint32)
func (_IFAllocationSaleV8 *IFAllocationSaleV8CallerSession) WithdrawerCount() (uint32, error) {
	return _IFAllocationSaleV8.Contract.WithdrawerCount(&_IFAllocationSaleV8.CallOpts)
}

// Cash is a paid mutator transaction binding the contract method 0x961be391.
//
// Solidity: function cash() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) Cash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "cash")
}

// Cash is a paid mutator transaction binding the contract method 0x961be391.
//
// Solidity: function cash() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) Cash() (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.Cash(&_IFAllocationSaleV8.TransactOpts)
}

// Cash is a paid mutator transaction binding the contract method 0x961be391.
//
// Solidity: function cash() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) Cash() (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.Cash(&_IFAllocationSaleV8.TransactOpts)
}

// EmergencyTokenRetrieve is a paid mutator transaction binding the contract method 0x26c65482.
//
// Solidity: function emergencyTokenRetrieve(address token) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) EmergencyTokenRetrieve(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "emergencyTokenRetrieve", token)
}

// EmergencyTokenRetrieve is a paid mutator transaction binding the contract method 0x26c65482.
//
// Solidity: function emergencyTokenRetrieve(address token) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) EmergencyTokenRetrieve(token common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.EmergencyTokenRetrieve(&_IFAllocationSaleV8.TransactOpts, token)
}

// EmergencyTokenRetrieve is a paid mutator transaction binding the contract method 0x26c65482.
//
// Solidity: function emergencyTokenRetrieve(address token) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) EmergencyTokenRetrieve(token common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.EmergencyTokenRetrieve(&_IFAllocationSaleV8.TransactOpts, token)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 amount) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "fund", amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 amount) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) Fund(amount *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.Fund(&_IFAllocationSaleV8.TransactOpts, amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 amount) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) Fund(amount *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.Fund(&_IFAllocationSaleV8.TransactOpts, amount)
}

// OptInBuyback is a paid mutator transaction binding the contract method 0x22524c9d.
//
// Solidity: function optInBuyback() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) OptInBuyback(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "optInBuyback")
}

// OptInBuyback is a paid mutator transaction binding the contract method 0x22524c9d.
//
// Solidity: function optInBuyback() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) OptInBuyback() (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.OptInBuyback(&_IFAllocationSaleV8.TransactOpts)
}

// OptInBuyback is a paid mutator transaction binding the contract method 0x22524c9d.
//
// Solidity: function optInBuyback() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) OptInBuyback() (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.OptInBuyback(&_IFAllocationSaleV8.TransactOpts)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 paymentAmount) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) Purchase(opts *bind.TransactOpts, paymentAmount *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "purchase", paymentAmount)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 paymentAmount) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) Purchase(paymentAmount *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.Purchase(&_IFAllocationSaleV8.TransactOpts, paymentAmount)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 paymentAmount) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) Purchase(paymentAmount *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.Purchase(&_IFAllocationSaleV8.TransactOpts, paymentAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) RenounceOwnership() (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.RenounceOwnership(&_IFAllocationSaleV8.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.RenounceOwnership(&_IFAllocationSaleV8.TransactOpts)
}

// SetBuybackClaimableNumber is a paid mutator transaction binding the contract method 0xc6632d55.
//
// Solidity: function setBuybackClaimableNumber(uint256 _buybackClaimableNumber) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetBuybackClaimableNumber(opts *bind.TransactOpts, _buybackClaimableNumber *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setBuybackClaimableNumber", _buybackClaimableNumber)
}

// SetBuybackClaimableNumber is a paid mutator transaction binding the contract method 0xc6632d55.
//
// Solidity: function setBuybackClaimableNumber(uint256 _buybackClaimableNumber) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetBuybackClaimableNumber(_buybackClaimableNumber *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetBuybackClaimableNumber(&_IFAllocationSaleV8.TransactOpts, _buybackClaimableNumber)
}

// SetBuybackClaimableNumber is a paid mutator transaction binding the contract method 0xc6632d55.
//
// Solidity: function setBuybackClaimableNumber(uint256 _buybackClaimableNumber) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetBuybackClaimableNumber(_buybackClaimableNumber *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetBuybackClaimableNumber(&_IFAllocationSaleV8.TransactOpts, _buybackClaimableNumber)
}

// SetCasher is a paid mutator transaction binding the contract method 0xa590c84e.
//
// Solidity: function setCasher(address _casher) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetCasher(opts *bind.TransactOpts, _casher common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setCasher", _casher)
}

// SetCasher is a paid mutator transaction binding the contract method 0xa590c84e.
//
// Solidity: function setCasher(address _casher) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetCasher(_casher common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetCasher(&_IFAllocationSaleV8.TransactOpts, _casher)
}

// SetCasher is a paid mutator transaction binding the contract method 0xa590c84e.
//
// Solidity: function setCasher(address _casher) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetCasher(_casher common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetCasher(&_IFAllocationSaleV8.TransactOpts, _casher)
}

// SetCliffPeriod is a paid mutator transaction binding the contract method 0x89910cac.
//
// Solidity: function setCliffPeriod(uint256[] claimTimes, uint8[] pct) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetCliffPeriod(opts *bind.TransactOpts, claimTimes []*big.Int, pct []uint8) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setCliffPeriod", claimTimes, pct)
}

// SetCliffPeriod is a paid mutator transaction binding the contract method 0x89910cac.
//
// Solidity: function setCliffPeriod(uint256[] claimTimes, uint8[] pct) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetCliffPeriod(claimTimes []*big.Int, pct []uint8) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetCliffPeriod(&_IFAllocationSaleV8.TransactOpts, claimTimes, pct)
}

// SetCliffPeriod is a paid mutator transaction binding the contract method 0x89910cac.
//
// Solidity: function setCliffPeriod(uint256[] claimTimes, uint8[] pct) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetCliffPeriod(claimTimes []*big.Int, pct []uint8) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetCliffPeriod(&_IFAllocationSaleV8.TransactOpts, claimTimes, pct)
}

// SetFunder is a paid mutator transaction binding the contract method 0x0acc8cd1.
//
// Solidity: function setFunder(address _funder) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetFunder(opts *bind.TransactOpts, _funder common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setFunder", _funder)
}

// SetFunder is a paid mutator transaction binding the contract method 0x0acc8cd1.
//
// Solidity: function setFunder(address _funder) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetFunder(_funder common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetFunder(&_IFAllocationSaleV8.TransactOpts, _funder)
}

// SetFunder is a paid mutator transaction binding the contract method 0x0acc8cd1.
//
// Solidity: function setFunder(address _funder) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetFunder(_funder common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetFunder(&_IFAllocationSaleV8.TransactOpts, _funder)
}

// SetLinearVestingEndTime is a paid mutator transaction binding the contract method 0xcddfb5fd.
//
// Solidity: function setLinearVestingEndTime(uint256 _vestingEndTime) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetLinearVestingEndTime(opts *bind.TransactOpts, _vestingEndTime *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setLinearVestingEndTime", _vestingEndTime)
}

// SetLinearVestingEndTime is a paid mutator transaction binding the contract method 0xcddfb5fd.
//
// Solidity: function setLinearVestingEndTime(uint256 _vestingEndTime) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetLinearVestingEndTime(_vestingEndTime *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetLinearVestingEndTime(&_IFAllocationSaleV8.TransactOpts, _vestingEndTime)
}

// SetLinearVestingEndTime is a paid mutator transaction binding the contract method 0xcddfb5fd.
//
// Solidity: function setLinearVestingEndTime(uint256 _vestingEndTime) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetLinearVestingEndTime(_vestingEndTime *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetLinearVestingEndTime(&_IFAllocationSaleV8.TransactOpts, _vestingEndTime)
}

// SetMaxTotalPurchasable is a paid mutator transaction binding the contract method 0xda943cee.
//
// Solidity: function setMaxTotalPurchasable(uint256 _maxTotalPurchasable) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetMaxTotalPurchasable(opts *bind.TransactOpts, _maxTotalPurchasable *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setMaxTotalPurchasable", _maxTotalPurchasable)
}

// SetMaxTotalPurchasable is a paid mutator transaction binding the contract method 0xda943cee.
//
// Solidity: function setMaxTotalPurchasable(uint256 _maxTotalPurchasable) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetMaxTotalPurchasable(_maxTotalPurchasable *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetMaxTotalPurchasable(&_IFAllocationSaleV8.TransactOpts, _maxTotalPurchasable)
}

// SetMaxTotalPurchasable is a paid mutator transaction binding the contract method 0xda943cee.
//
// Solidity: function setMaxTotalPurchasable(uint256 _maxTotalPurchasable) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetMaxTotalPurchasable(_maxTotalPurchasable *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetMaxTotalPurchasable(&_IFAllocationSaleV8.TransactOpts, _maxTotalPurchasable)
}

// SetMinTotalPayment is a paid mutator transaction binding the contract method 0x252b99c5.
//
// Solidity: function setMinTotalPayment(uint256 _minTotalPayment) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetMinTotalPayment(opts *bind.TransactOpts, _minTotalPayment *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setMinTotalPayment", _minTotalPayment)
}

// SetMinTotalPayment is a paid mutator transaction binding the contract method 0x252b99c5.
//
// Solidity: function setMinTotalPayment(uint256 _minTotalPayment) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetMinTotalPayment(_minTotalPayment *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetMinTotalPayment(&_IFAllocationSaleV8.TransactOpts, _minTotalPayment)
}

// SetMinTotalPayment is a paid mutator transaction binding the contract method 0x252b99c5.
//
// Solidity: function setMinTotalPayment(uint256 _minTotalPayment) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetMinTotalPayment(_minTotalPayment *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetMinTotalPayment(&_IFAllocationSaleV8.TransactOpts, _minTotalPayment)
}

// SetVestingEditable is a paid mutator transaction binding the contract method 0x2270e82d.
//
// Solidity: function setVestingEditable(bool _vestingEditableOverride) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetVestingEditable(opts *bind.TransactOpts, _vestingEditableOverride bool) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setVestingEditable", _vestingEditableOverride)
}

// SetVestingEditable is a paid mutator transaction binding the contract method 0x2270e82d.
//
// Solidity: function setVestingEditable(bool _vestingEditableOverride) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetVestingEditable(_vestingEditableOverride bool) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetVestingEditable(&_IFAllocationSaleV8.TransactOpts, _vestingEditableOverride)
}

// SetVestingEditable is a paid mutator transaction binding the contract method 0x2270e82d.
//
// Solidity: function setVestingEditable(bool _vestingEditableOverride) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetVestingEditable(_vestingEditableOverride bool) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetVestingEditable(&_IFAllocationSaleV8.TransactOpts, _vestingEditableOverride)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x440bc7f3.
//
// Solidity: function setWhitelist(bytes32 _whitelistRootHash) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetWhitelist(opts *bind.TransactOpts, _whitelistRootHash [32]byte) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setWhitelist", _whitelistRootHash)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x440bc7f3.
//
// Solidity: function setWhitelist(bytes32 _whitelistRootHash) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetWhitelist(_whitelistRootHash [32]byte) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetWhitelist(&_IFAllocationSaleV8.TransactOpts, _whitelistRootHash)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x440bc7f3.
//
// Solidity: function setWhitelist(bytes32 _whitelistRootHash) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetWhitelist(_whitelistRootHash [32]byte) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetWhitelist(&_IFAllocationSaleV8.TransactOpts, _whitelistRootHash)
}

// SetWhitelistSetter is a paid mutator transaction binding the contract method 0x547a5eee.
//
// Solidity: function setWhitelistSetter(address _whitelistSetter) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetWhitelistSetter(opts *bind.TransactOpts, _whitelistSetter common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setWhitelistSetter", _whitelistSetter)
}

// SetWhitelistSetter is a paid mutator transaction binding the contract method 0x547a5eee.
//
// Solidity: function setWhitelistSetter(address _whitelistSetter) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetWhitelistSetter(_whitelistSetter common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetWhitelistSetter(&_IFAllocationSaleV8.TransactOpts, _whitelistSetter)
}

// SetWhitelistSetter is a paid mutator transaction binding the contract method 0x547a5eee.
//
// Solidity: function setWhitelistSetter(address _whitelistSetter) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetWhitelistSetter(_whitelistSetter common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetWhitelistSetter(&_IFAllocationSaleV8.TransactOpts, _whitelistSetter)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0xaacc557a.
//
// Solidity: function setWithdrawDelay(uint24 _withdrawDelay) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) SetWithdrawDelay(opts *bind.TransactOpts, _withdrawDelay *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "setWithdrawDelay", _withdrawDelay)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0xaacc557a.
//
// Solidity: function setWithdrawDelay(uint24 _withdrawDelay) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) SetWithdrawDelay(_withdrawDelay *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetWithdrawDelay(&_IFAllocationSaleV8.TransactOpts, _withdrawDelay)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0xaacc557a.
//
// Solidity: function setWithdrawDelay(uint24 _withdrawDelay) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) SetWithdrawDelay(_withdrawDelay *big.Int) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.SetWithdrawDelay(&_IFAllocationSaleV8.TransactOpts, _withdrawDelay)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.TransferOwnership(&_IFAllocationSaleV8.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.TransferOwnership(&_IFAllocationSaleV8.TransactOpts, newOwner)
}

// WhitelistedPurchase is a paid mutator transaction binding the contract method 0x01fc191c.
//
// Solidity: function whitelistedPurchase(uint256 paymentAmount, bytes32[] merkleProof) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) WhitelistedPurchase(opts *bind.TransactOpts, paymentAmount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "whitelistedPurchase", paymentAmount, merkleProof)
}

// WhitelistedPurchase is a paid mutator transaction binding the contract method 0x01fc191c.
//
// Solidity: function whitelistedPurchase(uint256 paymentAmount, bytes32[] merkleProof) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) WhitelistedPurchase(paymentAmount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.WhitelistedPurchase(&_IFAllocationSaleV8.TransactOpts, paymentAmount, merkleProof)
}

// WhitelistedPurchase is a paid mutator transaction binding the contract method 0x01fc191c.
//
// Solidity: function whitelistedPurchase(uint256 paymentAmount, bytes32[] merkleProof) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) WhitelistedPurchase(paymentAmount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.WhitelistedPurchase(&_IFAllocationSaleV8.TransactOpts, paymentAmount, merkleProof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) Withdraw() (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.Withdraw(&_IFAllocationSaleV8.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) Withdraw() (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.Withdraw(&_IFAllocationSaleV8.TransactOpts)
}

// WithdrawGiveaway is a paid mutator transaction binding the contract method 0x1637cde0.
//
// Solidity: function withdrawGiveaway(bytes32[] merkleProof) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Transactor) WithdrawGiveaway(opts *bind.TransactOpts, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IFAllocationSaleV8.contract.Transact(opts, "withdrawGiveaway", merkleProof)
}

// WithdrawGiveaway is a paid mutator transaction binding the contract method 0x1637cde0.
//
// Solidity: function withdrawGiveaway(bytes32[] merkleProof) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8Session) WithdrawGiveaway(merkleProof [][32]byte) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.WithdrawGiveaway(&_IFAllocationSaleV8.TransactOpts, merkleProof)
}

// WithdrawGiveaway is a paid mutator transaction binding the contract method 0x1637cde0.
//
// Solidity: function withdrawGiveaway(bytes32[] merkleProof) returns()
func (_IFAllocationSaleV8 *IFAllocationSaleV8TransactorSession) WithdrawGiveaway(merkleProof [][32]byte) (*types.Transaction, error) {
	return _IFAllocationSaleV8.Contract.WithdrawGiveaway(&_IFAllocationSaleV8.TransactOpts, merkleProof)
}

// IFAllocationSaleV8CashIterator is returned from FilterCash and is used to iterate over the raw logs and unpacked data for Cash events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8CashIterator struct {
	Event *IFAllocationSaleV8Cash // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8CashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8Cash)
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
		it.Event = new(IFAllocationSaleV8Cash)
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
func (it *IFAllocationSaleV8CashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8CashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8Cash represents a Cash event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8Cash struct {
	Sender              common.Address
	PaymentTokenBalance *big.Int
	SaleTokenBalance    *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCash is a free log retrieval operation binding the contract event 0x83205c70ca31ffcb57664adecfd9894647d48665aefae1fb38bbc7ca4c1b86fb.
//
// Solidity: event Cash(address indexed sender, uint256 paymentTokenBalance, uint256 saleTokenBalance)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterCash(opts *bind.FilterOpts, sender []common.Address) (*IFAllocationSaleV8CashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "Cash", senderRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8CashIterator{contract: _IFAllocationSaleV8.contract, event: "Cash", logs: logs, sub: sub}, nil
}

// WatchCash is a free log subscription operation binding the contract event 0x83205c70ca31ffcb57664adecfd9894647d48665aefae1fb38bbc7ca4c1b86fb.
//
// Solidity: event Cash(address indexed sender, uint256 paymentTokenBalance, uint256 saleTokenBalance)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchCash(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8Cash, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "Cash", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8Cash)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "Cash", log); err != nil {
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

// ParseCash is a log parse operation binding the contract event 0x83205c70ca31ffcb57664adecfd9894647d48665aefae1fb38bbc7ca4c1b86fb.
//
// Solidity: event Cash(address indexed sender, uint256 paymentTokenBalance, uint256 saleTokenBalance)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseCash(log types.Log) (*IFAllocationSaleV8Cash, error) {
	event := new(IFAllocationSaleV8Cash)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "Cash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8EmergencyTokenRetrieveIterator is returned from FilterEmergencyTokenRetrieve and is used to iterate over the raw logs and unpacked data for EmergencyTokenRetrieve events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8EmergencyTokenRetrieveIterator struct {
	Event *IFAllocationSaleV8EmergencyTokenRetrieve // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8EmergencyTokenRetrieveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8EmergencyTokenRetrieve)
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
		it.Event = new(IFAllocationSaleV8EmergencyTokenRetrieve)
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
func (it *IFAllocationSaleV8EmergencyTokenRetrieveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8EmergencyTokenRetrieveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8EmergencyTokenRetrieve represents a EmergencyTokenRetrieve event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8EmergencyTokenRetrieve struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyTokenRetrieve is a free log retrieval operation binding the contract event 0xb1d34d1c064a5cb36c65797ef779e448eefccf2b978edbcfd206cc687c8cea5a.
//
// Solidity: event EmergencyTokenRetrieve(address indexed sender, uint256 amount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterEmergencyTokenRetrieve(opts *bind.FilterOpts, sender []common.Address) (*IFAllocationSaleV8EmergencyTokenRetrieveIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "EmergencyTokenRetrieve", senderRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8EmergencyTokenRetrieveIterator{contract: _IFAllocationSaleV8.contract, event: "EmergencyTokenRetrieve", logs: logs, sub: sub}, nil
}

// WatchEmergencyTokenRetrieve is a free log subscription operation binding the contract event 0xb1d34d1c064a5cb36c65797ef779e448eefccf2b978edbcfd206cc687c8cea5a.
//
// Solidity: event EmergencyTokenRetrieve(address indexed sender, uint256 amount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchEmergencyTokenRetrieve(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8EmergencyTokenRetrieve, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "EmergencyTokenRetrieve", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8EmergencyTokenRetrieve)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "EmergencyTokenRetrieve", log); err != nil {
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

// ParseEmergencyTokenRetrieve is a log parse operation binding the contract event 0xb1d34d1c064a5cb36c65797ef779e448eefccf2b978edbcfd206cc687c8cea5a.
//
// Solidity: event EmergencyTokenRetrieve(address indexed sender, uint256 amount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseEmergencyTokenRetrieve(log types.Log) (*IFAllocationSaleV8EmergencyTokenRetrieve, error) {
	event := new(IFAllocationSaleV8EmergencyTokenRetrieve)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "EmergencyTokenRetrieve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8FundIterator is returned from FilterFund and is used to iterate over the raw logs and unpacked data for Fund events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8FundIterator struct {
	Event *IFAllocationSaleV8Fund // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8FundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8Fund)
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
		it.Event = new(IFAllocationSaleV8Fund)
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
func (it *IFAllocationSaleV8FundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8FundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8Fund represents a Fund event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8Fund struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFund is a free log retrieval operation binding the contract event 0xda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd.
//
// Solidity: event Fund(address indexed sender, uint256 amount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterFund(opts *bind.FilterOpts, sender []common.Address) (*IFAllocationSaleV8FundIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "Fund", senderRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8FundIterator{contract: _IFAllocationSaleV8.contract, event: "Fund", logs: logs, sub: sub}, nil
}

// WatchFund is a free log subscription operation binding the contract event 0xda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd.
//
// Solidity: event Fund(address indexed sender, uint256 amount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchFund(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8Fund, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "Fund", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8Fund)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "Fund", log); err != nil {
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

// ParseFund is a log parse operation binding the contract event 0xda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd.
//
// Solidity: event Fund(address indexed sender, uint256 amount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseFund(log types.Log) (*IFAllocationSaleV8Fund, error) {
	event := new(IFAllocationSaleV8Fund)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "Fund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8OptInBuybackIterator is returned from FilterOptInBuyback and is used to iterate over the raw logs and unpacked data for OptInBuyback events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8OptInBuybackIterator struct {
	Event *IFAllocationSaleV8OptInBuyback // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8OptInBuybackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8OptInBuyback)
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
		it.Event = new(IFAllocationSaleV8OptInBuyback)
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
func (it *IFAllocationSaleV8OptInBuybackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8OptInBuybackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8OptInBuyback represents a OptInBuyback event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8OptInBuyback struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOptInBuyback is a free log retrieval operation binding the contract event 0x91b0dcebf54773cfc86cca32698ae9cd725b3f4e407d4b83e7e19f73c54eb954.
//
// Solidity: event OptInBuyback(address indexed user)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterOptInBuyback(opts *bind.FilterOpts, user []common.Address) (*IFAllocationSaleV8OptInBuybackIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "OptInBuyback", userRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8OptInBuybackIterator{contract: _IFAllocationSaleV8.contract, event: "OptInBuyback", logs: logs, sub: sub}, nil
}

// WatchOptInBuyback is a free log subscription operation binding the contract event 0x91b0dcebf54773cfc86cca32698ae9cd725b3f4e407d4b83e7e19f73c54eb954.
//
// Solidity: event OptInBuyback(address indexed user)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchOptInBuyback(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8OptInBuyback, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "OptInBuyback", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8OptInBuyback)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "OptInBuyback", log); err != nil {
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

// ParseOptInBuyback is a log parse operation binding the contract event 0x91b0dcebf54773cfc86cca32698ae9cd725b3f4e407d4b83e7e19f73c54eb954.
//
// Solidity: event OptInBuyback(address indexed user)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseOptInBuyback(log types.Log) (*IFAllocationSaleV8OptInBuyback, error) {
	event := new(IFAllocationSaleV8OptInBuyback)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "OptInBuyback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8OwnershipTransferredIterator struct {
	Event *IFAllocationSaleV8OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8OwnershipTransferred)
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
		it.Event = new(IFAllocationSaleV8OwnershipTransferred)
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
func (it *IFAllocationSaleV8OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8OwnershipTransferred represents a OwnershipTransferred event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IFAllocationSaleV8OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8OwnershipTransferredIterator{contract: _IFAllocationSaleV8.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8OwnershipTransferred)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseOwnershipTransferred(log types.Log) (*IFAllocationSaleV8OwnershipTransferred, error) {
	event := new(IFAllocationSaleV8OwnershipTransferred)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8PurchaseIterator is returned from FilterPurchase and is used to iterate over the raw logs and unpacked data for Purchase events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8PurchaseIterator struct {
	Event *IFAllocationSaleV8Purchase // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8PurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8Purchase)
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
		it.Event = new(IFAllocationSaleV8Purchase)
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
func (it *IFAllocationSaleV8PurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8PurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8Purchase represents a Purchase event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8Purchase struct {
	Sender        common.Address
	PaymentAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPurchase is a free log retrieval operation binding the contract event 0x2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f632.
//
// Solidity: event Purchase(address indexed sender, uint256 paymentAmount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterPurchase(opts *bind.FilterOpts, sender []common.Address) (*IFAllocationSaleV8PurchaseIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "Purchase", senderRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8PurchaseIterator{contract: _IFAllocationSaleV8.contract, event: "Purchase", logs: logs, sub: sub}, nil
}

// WatchPurchase is a free log subscription operation binding the contract event 0x2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f632.
//
// Solidity: event Purchase(address indexed sender, uint256 paymentAmount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchPurchase(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8Purchase, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "Purchase", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8Purchase)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "Purchase", log); err != nil {
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

// ParsePurchase is a log parse operation binding the contract event 0x2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f632.
//
// Solidity: event Purchase(address indexed sender, uint256 paymentAmount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParsePurchase(log types.Log) (*IFAllocationSaleV8Purchase, error) {
	event := new(IFAllocationSaleV8Purchase)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "Purchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8SetCasherIterator is returned from FilterSetCasher and is used to iterate over the raw logs and unpacked data for SetCasher events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetCasherIterator struct {
	Event *IFAllocationSaleV8SetCasher // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8SetCasherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8SetCasher)
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
		it.Event = new(IFAllocationSaleV8SetCasher)
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
func (it *IFAllocationSaleV8SetCasherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8SetCasherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8SetCasher represents a SetCasher event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetCasher struct {
	Casher common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetCasher is a free log retrieval operation binding the contract event 0xde112653552cba8a4f696cac12b4478ce2b9b8c0e04429455a2052ec7c0412ce.
//
// Solidity: event SetCasher(address indexed casher)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterSetCasher(opts *bind.FilterOpts, casher []common.Address) (*IFAllocationSaleV8SetCasherIterator, error) {

	var casherRule []interface{}
	for _, casherItem := range casher {
		casherRule = append(casherRule, casherItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "SetCasher", casherRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8SetCasherIterator{contract: _IFAllocationSaleV8.contract, event: "SetCasher", logs: logs, sub: sub}, nil
}

// WatchSetCasher is a free log subscription operation binding the contract event 0xde112653552cba8a4f696cac12b4478ce2b9b8c0e04429455a2052ec7c0412ce.
//
// Solidity: event SetCasher(address indexed casher)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchSetCasher(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8SetCasher, casher []common.Address) (event.Subscription, error) {

	var casherRule []interface{}
	for _, casherItem := range casher {
		casherRule = append(casherRule, casherItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "SetCasher", casherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8SetCasher)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetCasher", log); err != nil {
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

// ParseSetCasher is a log parse operation binding the contract event 0xde112653552cba8a4f696cac12b4478ce2b9b8c0e04429455a2052ec7c0412ce.
//
// Solidity: event SetCasher(address indexed casher)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseSetCasher(log types.Log) (*IFAllocationSaleV8SetCasher, error) {
	event := new(IFAllocationSaleV8SetCasher)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetCasher", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8SetCliffVestingPeriodIterator is returned from FilterSetCliffVestingPeriod and is used to iterate over the raw logs and unpacked data for SetCliffVestingPeriod events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetCliffVestingPeriodIterator struct {
	Event *IFAllocationSaleV8SetCliffVestingPeriod // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8SetCliffVestingPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8SetCliffVestingPeriod)
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
		it.Event = new(IFAllocationSaleV8SetCliffVestingPeriod)
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
func (it *IFAllocationSaleV8SetCliffVestingPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8SetCliffVestingPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8SetCliffVestingPeriod represents a SetCliffVestingPeriod event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetCliffVestingPeriod struct {
	CliffPeriod []IFVestableCliff
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetCliffVestingPeriod is a free log retrieval operation binding the contract event 0x3593839147bb7832f5ec3cbd35ab371fb0283d503cfc9460ed0bd7b473f8efd4.
//
// Solidity: event SetCliffVestingPeriod((uint256,uint8)[] indexed cliffPeriod)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterSetCliffVestingPeriod(opts *bind.FilterOpts, cliffPeriod [][]IFVestableCliff) (*IFAllocationSaleV8SetCliffVestingPeriodIterator, error) {

	var cliffPeriodRule []interface{}
	for _, cliffPeriodItem := range cliffPeriod {
		cliffPeriodRule = append(cliffPeriodRule, cliffPeriodItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "SetCliffVestingPeriod", cliffPeriodRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8SetCliffVestingPeriodIterator{contract: _IFAllocationSaleV8.contract, event: "SetCliffVestingPeriod", logs: logs, sub: sub}, nil
}

// WatchSetCliffVestingPeriod is a free log subscription operation binding the contract event 0x3593839147bb7832f5ec3cbd35ab371fb0283d503cfc9460ed0bd7b473f8efd4.
//
// Solidity: event SetCliffVestingPeriod((uint256,uint8)[] indexed cliffPeriod)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchSetCliffVestingPeriod(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8SetCliffVestingPeriod, cliffPeriod [][]IFVestableCliff) (event.Subscription, error) {

	var cliffPeriodRule []interface{}
	for _, cliffPeriodItem := range cliffPeriod {
		cliffPeriodRule = append(cliffPeriodRule, cliffPeriodItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "SetCliffVestingPeriod", cliffPeriodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8SetCliffVestingPeriod)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetCliffVestingPeriod", log); err != nil {
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

// ParseSetCliffVestingPeriod is a log parse operation binding the contract event 0x3593839147bb7832f5ec3cbd35ab371fb0283d503cfc9460ed0bd7b473f8efd4.
//
// Solidity: event SetCliffVestingPeriod((uint256,uint8)[] indexed cliffPeriod)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseSetCliffVestingPeriod(log types.Log) (*IFAllocationSaleV8SetCliffVestingPeriod, error) {
	event := new(IFAllocationSaleV8SetCliffVestingPeriod)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetCliffVestingPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8SetFunderIterator is returned from FilterSetFunder and is used to iterate over the raw logs and unpacked data for SetFunder events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetFunderIterator struct {
	Event *IFAllocationSaleV8SetFunder // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8SetFunderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8SetFunder)
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
		it.Event = new(IFAllocationSaleV8SetFunder)
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
func (it *IFAllocationSaleV8SetFunderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8SetFunderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8SetFunder represents a SetFunder event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetFunder struct {
	Funder common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFunder is a free log retrieval operation binding the contract event 0x73e0366d7ceb5a21fb27f8b2aa2720e2a9dee755bd85bfb5dbccf2830753b022.
//
// Solidity: event SetFunder(address indexed funder)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterSetFunder(opts *bind.FilterOpts, funder []common.Address) (*IFAllocationSaleV8SetFunderIterator, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "SetFunder", funderRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8SetFunderIterator{contract: _IFAllocationSaleV8.contract, event: "SetFunder", logs: logs, sub: sub}, nil
}

// WatchSetFunder is a free log subscription operation binding the contract event 0x73e0366d7ceb5a21fb27f8b2aa2720e2a9dee755bd85bfb5dbccf2830753b022.
//
// Solidity: event SetFunder(address indexed funder)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchSetFunder(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8SetFunder, funder []common.Address) (event.Subscription, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "SetFunder", funderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8SetFunder)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetFunder", log); err != nil {
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

// ParseSetFunder is a log parse operation binding the contract event 0x73e0366d7ceb5a21fb27f8b2aa2720e2a9dee755bd85bfb5dbccf2830753b022.
//
// Solidity: event SetFunder(address indexed funder)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseSetFunder(log types.Log) (*IFAllocationSaleV8SetFunder, error) {
	event := new(IFAllocationSaleV8SetFunder)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetFunder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8SetLinearVestingEndTimeIterator is returned from FilterSetLinearVestingEndTime and is used to iterate over the raw logs and unpacked data for SetLinearVestingEndTime events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetLinearVestingEndTimeIterator struct {
	Event *IFAllocationSaleV8SetLinearVestingEndTime // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8SetLinearVestingEndTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8SetLinearVestingEndTime)
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
		it.Event = new(IFAllocationSaleV8SetLinearVestingEndTime)
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
func (it *IFAllocationSaleV8SetLinearVestingEndTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8SetLinearVestingEndTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8SetLinearVestingEndTime represents a SetLinearVestingEndTime event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetLinearVestingEndTime struct {
	LinearVestingEndTime *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetLinearVestingEndTime is a free log retrieval operation binding the contract event 0x83d190eb78c1206ea1abb0222d475a5f70b7f63bcd534ab65e9404b39016c7a4.
//
// Solidity: event SetLinearVestingEndTime(uint256 indexed linearVestingEndTime)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterSetLinearVestingEndTime(opts *bind.FilterOpts, linearVestingEndTime []*big.Int) (*IFAllocationSaleV8SetLinearVestingEndTimeIterator, error) {

	var linearVestingEndTimeRule []interface{}
	for _, linearVestingEndTimeItem := range linearVestingEndTime {
		linearVestingEndTimeRule = append(linearVestingEndTimeRule, linearVestingEndTimeItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "SetLinearVestingEndTime", linearVestingEndTimeRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8SetLinearVestingEndTimeIterator{contract: _IFAllocationSaleV8.contract, event: "SetLinearVestingEndTime", logs: logs, sub: sub}, nil
}

// WatchSetLinearVestingEndTime is a free log subscription operation binding the contract event 0x83d190eb78c1206ea1abb0222d475a5f70b7f63bcd534ab65e9404b39016c7a4.
//
// Solidity: event SetLinearVestingEndTime(uint256 indexed linearVestingEndTime)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchSetLinearVestingEndTime(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8SetLinearVestingEndTime, linearVestingEndTime []*big.Int) (event.Subscription, error) {

	var linearVestingEndTimeRule []interface{}
	for _, linearVestingEndTimeItem := range linearVestingEndTime {
		linearVestingEndTimeRule = append(linearVestingEndTimeRule, linearVestingEndTimeItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "SetLinearVestingEndTime", linearVestingEndTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8SetLinearVestingEndTime)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetLinearVestingEndTime", log); err != nil {
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

// ParseSetLinearVestingEndTime is a log parse operation binding the contract event 0x83d190eb78c1206ea1abb0222d475a5f70b7f63bcd534ab65e9404b39016c7a4.
//
// Solidity: event SetLinearVestingEndTime(uint256 indexed linearVestingEndTime)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseSetLinearVestingEndTime(log types.Log) (*IFAllocationSaleV8SetLinearVestingEndTime, error) {
	event := new(IFAllocationSaleV8SetLinearVestingEndTime)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetLinearVestingEndTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8SetMaxTotalPurchasableIterator is returned from FilterSetMaxTotalPurchasable and is used to iterate over the raw logs and unpacked data for SetMaxTotalPurchasable events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetMaxTotalPurchasableIterator struct {
	Event *IFAllocationSaleV8SetMaxTotalPurchasable // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8SetMaxTotalPurchasableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8SetMaxTotalPurchasable)
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
		it.Event = new(IFAllocationSaleV8SetMaxTotalPurchasable)
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
func (it *IFAllocationSaleV8SetMaxTotalPurchasableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8SetMaxTotalPurchasableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8SetMaxTotalPurchasable represents a SetMaxTotalPurchasable event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetMaxTotalPurchasable struct {
	MaxTotalPurchasable *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetMaxTotalPurchasable is a free log retrieval operation binding the contract event 0x851466103668359b383470efccd1760dc8caf09bc6d5a74acc7d78ac751c21e9.
//
// Solidity: event SetMaxTotalPurchasable(uint256 indexed _maxTotalPurchasable)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterSetMaxTotalPurchasable(opts *bind.FilterOpts, _maxTotalPurchasable []*big.Int) (*IFAllocationSaleV8SetMaxTotalPurchasableIterator, error) {

	var _maxTotalPurchasableRule []interface{}
	for _, _maxTotalPurchasableItem := range _maxTotalPurchasable {
		_maxTotalPurchasableRule = append(_maxTotalPurchasableRule, _maxTotalPurchasableItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "SetMaxTotalPurchasable", _maxTotalPurchasableRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8SetMaxTotalPurchasableIterator{contract: _IFAllocationSaleV8.contract, event: "SetMaxTotalPurchasable", logs: logs, sub: sub}, nil
}

// WatchSetMaxTotalPurchasable is a free log subscription operation binding the contract event 0x851466103668359b383470efccd1760dc8caf09bc6d5a74acc7d78ac751c21e9.
//
// Solidity: event SetMaxTotalPurchasable(uint256 indexed _maxTotalPurchasable)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchSetMaxTotalPurchasable(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8SetMaxTotalPurchasable, _maxTotalPurchasable []*big.Int) (event.Subscription, error) {

	var _maxTotalPurchasableRule []interface{}
	for _, _maxTotalPurchasableItem := range _maxTotalPurchasable {
		_maxTotalPurchasableRule = append(_maxTotalPurchasableRule, _maxTotalPurchasableItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "SetMaxTotalPurchasable", _maxTotalPurchasableRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8SetMaxTotalPurchasable)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetMaxTotalPurchasable", log); err != nil {
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

// ParseSetMaxTotalPurchasable is a log parse operation binding the contract event 0x851466103668359b383470efccd1760dc8caf09bc6d5a74acc7d78ac751c21e9.
//
// Solidity: event SetMaxTotalPurchasable(uint256 indexed _maxTotalPurchasable)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseSetMaxTotalPurchasable(log types.Log) (*IFAllocationSaleV8SetMaxTotalPurchasable, error) {
	event := new(IFAllocationSaleV8SetMaxTotalPurchasable)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetMaxTotalPurchasable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8SetMinTotalPaymentIterator is returned from FilterSetMinTotalPayment and is used to iterate over the raw logs and unpacked data for SetMinTotalPayment events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetMinTotalPaymentIterator struct {
	Event *IFAllocationSaleV8SetMinTotalPayment // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8SetMinTotalPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8SetMinTotalPayment)
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
		it.Event = new(IFAllocationSaleV8SetMinTotalPayment)
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
func (it *IFAllocationSaleV8SetMinTotalPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8SetMinTotalPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8SetMinTotalPayment represents a SetMinTotalPayment event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetMinTotalPayment struct {
	MinTotalPayment *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetMinTotalPayment is a free log retrieval operation binding the contract event 0x4d893dbff365afe590ba8bf0f5d258598f5e63fde5b989dcd4902cf89d403c2f.
//
// Solidity: event SetMinTotalPayment(uint256 indexed minTotalPayment)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterSetMinTotalPayment(opts *bind.FilterOpts, minTotalPayment []*big.Int) (*IFAllocationSaleV8SetMinTotalPaymentIterator, error) {

	var minTotalPaymentRule []interface{}
	for _, minTotalPaymentItem := range minTotalPayment {
		minTotalPaymentRule = append(minTotalPaymentRule, minTotalPaymentItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "SetMinTotalPayment", minTotalPaymentRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8SetMinTotalPaymentIterator{contract: _IFAllocationSaleV8.contract, event: "SetMinTotalPayment", logs: logs, sub: sub}, nil
}

// WatchSetMinTotalPayment is a free log subscription operation binding the contract event 0x4d893dbff365afe590ba8bf0f5d258598f5e63fde5b989dcd4902cf89d403c2f.
//
// Solidity: event SetMinTotalPayment(uint256 indexed minTotalPayment)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchSetMinTotalPayment(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8SetMinTotalPayment, minTotalPayment []*big.Int) (event.Subscription, error) {

	var minTotalPaymentRule []interface{}
	for _, minTotalPaymentItem := range minTotalPayment {
		minTotalPaymentRule = append(minTotalPaymentRule, minTotalPaymentItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "SetMinTotalPayment", minTotalPaymentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8SetMinTotalPayment)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetMinTotalPayment", log); err != nil {
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

// ParseSetMinTotalPayment is a log parse operation binding the contract event 0x4d893dbff365afe590ba8bf0f5d258598f5e63fde5b989dcd4902cf89d403c2f.
//
// Solidity: event SetMinTotalPayment(uint256 indexed minTotalPayment)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseSetMinTotalPayment(log types.Log) (*IFAllocationSaleV8SetMinTotalPayment, error) {
	event := new(IFAllocationSaleV8SetMinTotalPayment)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetMinTotalPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8SetWhitelistIterator is returned from FilterSetWhitelist and is used to iterate over the raw logs and unpacked data for SetWhitelist events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetWhitelistIterator struct {
	Event *IFAllocationSaleV8SetWhitelist // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8SetWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8SetWhitelist)
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
		it.Event = new(IFAllocationSaleV8SetWhitelist)
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
func (it *IFAllocationSaleV8SetWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8SetWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8SetWhitelist represents a SetWhitelist event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetWhitelist struct {
	WhitelistRootHash [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetWhitelist is a free log retrieval operation binding the contract event 0xc504a95022b9d01b3024e95b0a85c200d0c538f417160776512b95fcbf7b2daa.
//
// Solidity: event SetWhitelist(bytes32 indexed whitelistRootHash)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterSetWhitelist(opts *bind.FilterOpts, whitelistRootHash [][32]byte) (*IFAllocationSaleV8SetWhitelistIterator, error) {

	var whitelistRootHashRule []interface{}
	for _, whitelistRootHashItem := range whitelistRootHash {
		whitelistRootHashRule = append(whitelistRootHashRule, whitelistRootHashItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "SetWhitelist", whitelistRootHashRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8SetWhitelistIterator{contract: _IFAllocationSaleV8.contract, event: "SetWhitelist", logs: logs, sub: sub}, nil
}

// WatchSetWhitelist is a free log subscription operation binding the contract event 0xc504a95022b9d01b3024e95b0a85c200d0c538f417160776512b95fcbf7b2daa.
//
// Solidity: event SetWhitelist(bytes32 indexed whitelistRootHash)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchSetWhitelist(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8SetWhitelist, whitelistRootHash [][32]byte) (event.Subscription, error) {

	var whitelistRootHashRule []interface{}
	for _, whitelistRootHashItem := range whitelistRootHash {
		whitelistRootHashRule = append(whitelistRootHashRule, whitelistRootHashItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "SetWhitelist", whitelistRootHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8SetWhitelist)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetWhitelist", log); err != nil {
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

// ParseSetWhitelist is a log parse operation binding the contract event 0xc504a95022b9d01b3024e95b0a85c200d0c538f417160776512b95fcbf7b2daa.
//
// Solidity: event SetWhitelist(bytes32 indexed whitelistRootHash)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseSetWhitelist(log types.Log) (*IFAllocationSaleV8SetWhitelist, error) {
	event := new(IFAllocationSaleV8SetWhitelist)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8SetWhitelistSetterIterator is returned from FilterSetWhitelistSetter and is used to iterate over the raw logs and unpacked data for SetWhitelistSetter events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetWhitelistSetterIterator struct {
	Event *IFAllocationSaleV8SetWhitelistSetter // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8SetWhitelistSetterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8SetWhitelistSetter)
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
		it.Event = new(IFAllocationSaleV8SetWhitelistSetter)
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
func (it *IFAllocationSaleV8SetWhitelistSetterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8SetWhitelistSetterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8SetWhitelistSetter represents a SetWhitelistSetter event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetWhitelistSetter struct {
	WhitelistSetter common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetWhitelistSetter is a free log retrieval operation binding the contract event 0x7aba5fc71e3607d34203924738fbacc00b5782879615f86db108794b4bdcc95b.
//
// Solidity: event SetWhitelistSetter(address indexed whitelistSetter)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterSetWhitelistSetter(opts *bind.FilterOpts, whitelistSetter []common.Address) (*IFAllocationSaleV8SetWhitelistSetterIterator, error) {

	var whitelistSetterRule []interface{}
	for _, whitelistSetterItem := range whitelistSetter {
		whitelistSetterRule = append(whitelistSetterRule, whitelistSetterItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "SetWhitelistSetter", whitelistSetterRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8SetWhitelistSetterIterator{contract: _IFAllocationSaleV8.contract, event: "SetWhitelistSetter", logs: logs, sub: sub}, nil
}

// WatchSetWhitelistSetter is a free log subscription operation binding the contract event 0x7aba5fc71e3607d34203924738fbacc00b5782879615f86db108794b4bdcc95b.
//
// Solidity: event SetWhitelistSetter(address indexed whitelistSetter)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchSetWhitelistSetter(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8SetWhitelistSetter, whitelistSetter []common.Address) (event.Subscription, error) {

	var whitelistSetterRule []interface{}
	for _, whitelistSetterItem := range whitelistSetter {
		whitelistSetterRule = append(whitelistSetterRule, whitelistSetterItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "SetWhitelistSetter", whitelistSetterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8SetWhitelistSetter)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetWhitelistSetter", log); err != nil {
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

// ParseSetWhitelistSetter is a log parse operation binding the contract event 0x7aba5fc71e3607d34203924738fbacc00b5782879615f86db108794b4bdcc95b.
//
// Solidity: event SetWhitelistSetter(address indexed whitelistSetter)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseSetWhitelistSetter(log types.Log) (*IFAllocationSaleV8SetWhitelistSetter, error) {
	event := new(IFAllocationSaleV8SetWhitelistSetter)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetWhitelistSetter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8SetWithdrawDelayIterator is returned from FilterSetWithdrawDelay and is used to iterate over the raw logs and unpacked data for SetWithdrawDelay events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetWithdrawDelayIterator struct {
	Event *IFAllocationSaleV8SetWithdrawDelay // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8SetWithdrawDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8SetWithdrawDelay)
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
		it.Event = new(IFAllocationSaleV8SetWithdrawDelay)
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
func (it *IFAllocationSaleV8SetWithdrawDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8SetWithdrawDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8SetWithdrawDelay represents a SetWithdrawDelay event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8SetWithdrawDelay struct {
	WithdrawDelay *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawDelay is a free log retrieval operation binding the contract event 0x0dac7b34fe917a51c24c51f379615699f5c23ce9168eaed469bc161c90028608.
//
// Solidity: event SetWithdrawDelay(uint24 indexed withdrawDelay)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterSetWithdrawDelay(opts *bind.FilterOpts, withdrawDelay []*big.Int) (*IFAllocationSaleV8SetWithdrawDelayIterator, error) {

	var withdrawDelayRule []interface{}
	for _, withdrawDelayItem := range withdrawDelay {
		withdrawDelayRule = append(withdrawDelayRule, withdrawDelayItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "SetWithdrawDelay", withdrawDelayRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8SetWithdrawDelayIterator{contract: _IFAllocationSaleV8.contract, event: "SetWithdrawDelay", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawDelay is a free log subscription operation binding the contract event 0x0dac7b34fe917a51c24c51f379615699f5c23ce9168eaed469bc161c90028608.
//
// Solidity: event SetWithdrawDelay(uint24 indexed withdrawDelay)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchSetWithdrawDelay(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8SetWithdrawDelay, withdrawDelay []*big.Int) (event.Subscription, error) {

	var withdrawDelayRule []interface{}
	for _, withdrawDelayItem := range withdrawDelay {
		withdrawDelayRule = append(withdrawDelayRule, withdrawDelayItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "SetWithdrawDelay", withdrawDelayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8SetWithdrawDelay)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetWithdrawDelay", log); err != nil {
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

// ParseSetWithdrawDelay is a log parse operation binding the contract event 0x0dac7b34fe917a51c24c51f379615699f5c23ce9168eaed469bc161c90028608.
//
// Solidity: event SetWithdrawDelay(uint24 indexed withdrawDelay)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseSetWithdrawDelay(log types.Log) (*IFAllocationSaleV8SetWithdrawDelay, error) {
	event := new(IFAllocationSaleV8SetWithdrawDelay)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "SetWithdrawDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFAllocationSaleV8WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8WithdrawIterator struct {
	Event *IFAllocationSaleV8Withdraw // Event containing the contract specifics and raw log

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
func (it *IFAllocationSaleV8WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFAllocationSaleV8Withdraw)
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
		it.Event = new(IFAllocationSaleV8Withdraw)
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
func (it *IFAllocationSaleV8WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFAllocationSaleV8WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFAllocationSaleV8Withdraw represents a Withdraw event raised by the IFAllocationSaleV8 contract.
type IFAllocationSaleV8Withdraw struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address) (*IFAllocationSaleV8WithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.FilterLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return &IFAllocationSaleV8WithdrawIterator{contract: _IFAllocationSaleV8.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IFAllocationSaleV8Withdraw, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFAllocationSaleV8.contract.WatchLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFAllocationSaleV8Withdraw)
				if err := _IFAllocationSaleV8.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount)
func (_IFAllocationSaleV8 *IFAllocationSaleV8Filterer) ParseWithdraw(log types.Log) (*IFAllocationSaleV8Withdraw, error) {
	event := new(IFAllocationSaleV8Withdraw)
	if err := _IFAllocationSaleV8.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
