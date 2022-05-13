package processor

import (
	"context"

	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
)

type MatchableContract interface {
	Type() string
	Match(ctx context.Context, provider jsonrpc.Client) bool
}

// ERC20
type ERC20Contract struct {
	Address string
	Code    *types.Code
	MatchableContract
}

func (c *ERC20Contract) Type() string {
	return "ERC20"
}

func (c *ERC20Contract) Match(ctx context.Context, provider jsonrpc.Client) bool {
	// https://github.com/OpenZeppelin/cairo-contracts/blob/main/src/openzeppelin/token/erc20/interfaces/IERC20.cairo
	// stark_call function / set of functions

	// check symbol, decimals and balanceOf functions
	if _, err := provider.Call(ctx, jsonrpc.FunctionCall{
		ContractAddress:    c.Address,
		EntryPointSelector: "symbol",
	}, "latest"); err != nil {
		return false
	}

	if _, err := provider.Call(ctx, jsonrpc.FunctionCall{
		ContractAddress:    c.Address,
		EntryPointSelector: "decimals",
	}, "latest"); err != nil {
		return false
	}

	if _, err := provider.Call(ctx, jsonrpc.FunctionCall{
		ContractAddress:    c.Address,
		EntryPointSelector: "balanceOf",
		Calldata:           []string{"0x050c47150563ff7cf60dd60f7d0bd4d62a9cc5331441916e5099e905bdd8c4bc"},
	}, "latest"); err != nil {
		return false
	}

	return true
}

// ERC721
type ERC721Contract struct {
	Address string
	Code    *types.Code
	MatchableContract
}

func (c *ERC721Contract) Type() string {
	return "ERC721"
}

func (c *ERC721Contract) Match(ctx context.Context, provider jsonrpc.Client) bool {
	// https://github.com/OpenZeppelin/cairo-contracts/blob/main/src/openzeppelin/token/erc721/ERC721_Mintable_Burnable.cairo
	// supportsInterface
	if res, err := provider.Call(ctx, jsonrpc.FunctionCall{
		ContractAddress:    c.Address,
		EntryPointSelector: "supportsInterface",
		Calldata:           []string{"0x01ffc9a7"},
	}, "latest"); err != nil || res[0] != "0x01" {
		return false
	}

	return true
}

func Match(ctx context.Context, address string, code *types.Code, provider jsonrpc.Client) MatchableContract {
	var contract MatchableContract = &ERC20Contract{Address: address, Code: code}
	if contract.Match(ctx, provider) {
		return contract
	}

	contract = &ERC721Contract{Address: address, Code: code}
	if contract.Match(ctx, provider) {
		return contract
	}

	return nil
}
