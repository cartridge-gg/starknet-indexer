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

// ERC721
type ERC721Contract struct {
	EventProcessor
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

func (c *ERC721Contract) Process(ctx context.Context, rpc *jsonrpc.Client, b *types.Block, txn *types.Transaction, evt *Event) (func(tx *ent.Tx) error, error) {
	if len(evt.Keys) == 0 || evt.Keys[0].Cmp(caigo.GetSelectorFromName("Transfer")) != 0 {
		return nil, nil
	}

	return func(tx *ent.Tx) error {
		if err := tx.Balance.Create().
			SetID(fmt.Sprintf("%s:%s", evt.Data[0].Hex(), evt.FromAddress)).
			SetAccountID(evt.Data[0].Hex()).
			SetContractID(evt.FromAddress).
			SetTokenId(big.FromBase(evt.Data[2].Int)).
			SetBalance(big.NewInt(0)).
			OnConflictColumns("id").
			SetBalance(big.NewInt(0)).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("balance:%s:%s", evt.Data[0].Hex(), evt.FromAddress),
				Err:   err,
			}
		}

		if err := tx.Balance.Create().
			SetID(fmt.Sprintf("%s:%s", evt.Data[1].Hex(), evt.FromAddress)).
			SetAccountID(evt.Data[0].Hex()).
			SetContractID(evt.FromAddress).
			SetTokenId(big.FromBase(evt.Data[2].Int)).
			SetBalance(big.NewInt(1)).
			OnConflictColumns("id").
			SetBalance(big.NewInt(1)).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("balance:%s:%s", evt.Data[0].Hex(), evt.FromAddress),
				Err:   err,
			}
		}
		return nil
	}, nil
}
