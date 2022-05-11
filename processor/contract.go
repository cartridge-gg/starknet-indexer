package processor

import "github.com/dontpanicdao/caigo/types"

type MatchableContract interface {
	Type() string
	Match() bool
}

// ERC20
type ERC20Contract struct {
	Address string
	Code    *types.Code
	MatchableContract
}

func (c *ERC20Contract) Type() string {
	return "erc20"
}

func (c *ERC20Contract) Match() bool {
	// https://github.com/OpenZeppelin/cairo-contracts/blob/main/src/openzeppelin/token/erc20/interfaces/IERC20.cairo
	// stark_call function / set of functions
	return false
}

// ERC721
type ERC721Contract struct {
	Address string
	Code    *types.Code
	MatchableContract
}

func (c *ERC721Contract) Type() string {
	return "erc721"
}

func (c *ERC721Contract) Match() bool {
	// https://github.com/OpenZeppelin/cairo-contracts/blob/main/src/openzeppelin/token/erc721/ERC721_Mintable_Burnable.cairo
	// supportsInterface
	return false
}

func Match(address string, code *types.Code) MatchableContract {
	var contract MatchableContract = &ERC20Contract{Address: address, Code: code}
	if contract.Match() {
		return contract
	}

	contract = &ERC721Contract{Address: address, Code: code}
	if contract.Match() {
		return contract
	}

	return nil
}
