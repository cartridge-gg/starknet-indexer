package processor

import (
	"context"
	"fmt"

	"github.com/cartridge-gg/starknet-indexer/ent"
	"github.com/cartridge-gg/starknet-indexer/ent/schema/big"
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
)

// ERC20
type ERC20Contract struct {
	EventProcessor
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

func (c *ERC20Contract) Process(ctx context.Context, rpc *jsonrpc.Client, b *types.Block, txn *types.Transaction, evt *Event) (func(tx *ent.Tx) error, error) {
	if len(evt.Keys) == 0 || evt.Keys[0].Cmp(caigo.GetSelectorFromName("Transfer")) != 0 {
		return nil, nil
	}

	return func(tx *ent.Tx) error {
		if err := tx.Balance.Create().
			SetID(fmt.Sprintf("%s:%s", evt.Data[0].Hex(), evt.FromAddress)).
			SetAccountID(evt.Data[0].Hex()).
			SetContractID(evt.FromAddress).
			OnConflictColumns("id").
			AddBalance(big.FromBase(evt.Data[2].Int).Neg()).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("balance:%s:%s", evt.Data[0].Hex(), evt.FromAddress),
				Err:   err,
			}
		}

		if err := tx.Balance.Create().
			SetID(fmt.Sprintf("%s:%s", evt.Data[0].Hex(), evt.FromAddress)).
			SetAccountID(evt.Data[0].Hex()).
			SetContractID(evt.FromAddress).
			SetBalance(big.FromBase(evt.Data[2].Int)).
			OnConflictColumns("id").
			AddBalance(big.FromBase(evt.Data[2].Int)).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("balance:%s:%s", evt.Data[0].Hex(), evt.FromAddress),
				Err:   err,
			}
		}
		return nil
	}, nil
}
