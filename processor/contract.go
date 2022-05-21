package processor

import (
	"context"

	"github.com/dontpanicdao/caigo/jsonrpc"
)

type MatchableContract interface {
	Address() string
	Type() string
	Match(ctx context.Context, provider *jsonrpc.Client) bool
}

// ERC20
type ERC20Contract struct {
	MatchableContract
	address string
}

func NewERC20Contract(address string) *ERC20Contract {
	return &ERC20Contract{address: address}
}

func (c *ERC20Contract) Address() string {
	return c.address
}

func (c *ERC20Contract) Type() string {
	return "ERC20"
}

func (c *ERC20Contract) Match(ctx context.Context, provider *jsonrpc.Client) bool {
	// https://github.com/OpenZeppelin/cairo-contracts/blob/main/src/openzeppelin/token/erc20/interfaces/IERC20.cairo
	// stark_call function / set of functions

	// check symbol, decimals and balanceOf functions
	// if _, err := provider.Call(ctx, jsonrpc.FunctionCall{
	// 	ContractAddress:    c.address,
	// 	EntryPointSelector: "symbol",
	// }, "latest"); err != nil {
	// 	return false
	// }

	// if _, err := provider.Call(ctx, jsonrpc.FunctionCall{
	// 	ContractAddress:    c.address,
	// 	EntryPointSelector: "decimals",
	// }, "latest"); err != nil {
	// 	return false
	// }

	if _, err := provider.Call(ctx, jsonrpc.FunctionCall{
		ContractAddress:    c.address,
		EntryPointSelector: "balanceOf",
		Calldata:           []string{"0x050c47150563ff7cf60dd60f7d0bd4d62a9cc5331441916e5099e905bdd8c4bc"},
	}, "latest"); err != nil {
		return false
	}

	return true
}

// ERC721
type ERC721Contract struct {
	MatchableContract
	address string
}

func NewERC721Contract(address string) *ERC721Contract {
	return &ERC721Contract{address: address}
}

func (c *ERC721Contract) Address() string {
	return c.address
}

func (c *ERC721Contract) Type() string {
	return "ERC721"
}

func (c *ERC721Contract) Match(ctx context.Context, provider *jsonrpc.Client) bool {
	// https://github.com/OpenZeppelin/cairo-contracts/blob/main/src/openzeppelin/token/erc721/ERC721_Mintable_Burnable.cairo
	// supportsInterface
	if res, err := provider.Call(ctx, jsonrpc.FunctionCall{
		ContractAddress:    c.address,
		EntryPointSelector: "supportsInterface",
		Calldata:           []string{"0x80ac58cd"},
	}, "latest"); err != nil || res[0] != "0x01" {
		return false
	}

	return true
}

func Match(ctx context.Context, address string, provider *jsonrpc.Client) MatchableContract {
	if c := NewERC20Contract(address); c.Match(ctx, provider) {
		return c
	}

	if c := NewERC721Contract(address); c.Match(ctx, provider) {
		return c
	}

	return nil
}
