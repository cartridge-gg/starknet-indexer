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
type ERC1155Contract struct {
	EventProcessor
	MatchableContract
	address string
}

func NewERC1155Contract(address string) *ERC721Contract {
	return &ERC721Contract{address: address}
}

func (c *ERC1155Contract) Address() string {
	return c.address
}

func (c *ERC1155Contract) Type() string {
	return "ERC1155"
}

func (c *ERC1155Contract) Match(ctx context.Context, provider *jsonrpc.Client) bool {
	// https://github.com/OpenZeppelin/cairo-contracts/blob/main/src/openzeppelin/token/erc721/ERC721_Mintable_Burnable.cairo
	// supportsInterface
	if res, err := provider.Call(ctx, jsonrpc.FunctionCall{
		ContractAddress:    c.address,
		EntryPointSelector: "supportsInterface",
		Calldata:           []string{"0xd9b67a26"},
	}, "latest"); err != nil || res[0] != "0x01" {
		return false
	}

	return true
}

func (c *ERC1155Contract) handleTransfer(ctx context.Context, evt *Event, transferIdx uint64) func(tx *ent.Tx) error {
	return func(tx *ent.Tx) error {
		sender, receiver, tokenId, value := evt.Data[1], evt.Data[2], evt.Data[3+transferIdx], evt.Data[4+transferIdx]
		// sender
		if err := tx.Balance.Create().
			SetID(fmt.Sprintf("%s:%s:%s", sender.Hex(), evt.FromAddress, tokenId.String())).
			SetAccountID(sender.Hex()).
			SetContractID(evt.FromAddress).
			SetTokenId(big.FromBase(tokenId.Int)).
			OnConflictColumns("id").
			AddBalance(big.FromBase(value.Int).Neg()).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("balance:%s:%s", sender.Hex(), evt.FromAddress),
				Err:   err,
			}
		}

		// receiver
		if err := tx.Balance.Create().
			SetID(fmt.Sprintf("%s:%s:%s", receiver.Hex(), evt.FromAddress, tokenId.String())).
			SetAccountID(evt.Data[2].Hex()).
			SetContractID(evt.FromAddress).
			SetTokenId(big.FromBase(tokenId.Int)).
			SetBalance(big.FromBase(value.Int)).
			OnConflictColumns("id").
			AddBalance(big.FromBase(value.Int)).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("balance:%s:%s", receiver.Hex(), evt.FromAddress),
				Err:   err,
			}
		}

		return nil
	}
}

func (c *ERC1155Contract) Process(ctx context.Context, rpc *jsonrpc.Client, b *types.Block, txn *types.Transaction, evt *Event) (func(tx *ent.Tx) error, error) {
	if len(evt.Keys) == 0 {
		return nil, nil
	}

	// https://eips.ethereum.org/EIPS/eip-1155#specification
	// event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _value)
	// event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _values)
	// NOTE: balance entity id is in this structure: account:contract:tokenId
	if evt.Keys[0].Cmp(caigo.GetSelectorFromName("TransferSingle")) == 0 {
		return c.handleTransfer(ctx, evt, 0), nil
	} else if evt.Keys[0].Cmp(caigo.GetSelectorFromName("TransferBatch")) == 0 {
		// get length of "_ids" and "_values" in event data
		// and divide by 2 to get the number of transfers
		// NOTE: we assume that _ids and _values have the same length
		transfers := uint64((len(evt.Data) - 3) / 2)
		return func(tx *ent.Tx) error {
			for i := uint64(0); i < transfers; i++ {
				if err := c.handleTransfer(ctx, evt, i)(tx); err != nil {
					return err
				}
			}
			return nil
		}, nil
	}

	return nil, nil
}
