package processor

import (
	"context"
	"errors"
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

func (c *ERC1155Contract) Process(ctx context.Context, rpc *jsonrpc.Client, b *types.Block, txn *types.Transaction, evt *Event) (func(tx *ent.Tx) error, error) {
	if len(evt.Keys) == 0 {
		return nil, nil
	}

	handleTransfer := func(tx *ent.Tx) error {
		sender, receiver, tokenId, value := evt.Data[1], evt.Data[2], evt.Data[3], evt.Data[4]
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

	// https://eips.ethereum.org/EIPS/eip-1155#specification
	// event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _value)
	// event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _values)
	// NOTE: balance entity id is in this structure: account:contract:tokenId
	if evt.Keys[0].Cmp(caigo.GetSelectorFromName("TransferSingle")) != 0 {
		return handleTransfer, nil
	} else if evt.Keys[0].Cmp(caigo.GetSelectorFromName("TransferBatch")) != 0 {
		// TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _values)
		return nil, &ProcessorError{
			Scope: "erc1155:TransferBatch",
			Err:   errors.New("not implemented"),
		}
	}

	return nil, nil
}
