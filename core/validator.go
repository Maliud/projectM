package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct{
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("zincir zaten hash ile blok (%d) içeriyor (%s)", b.Height, b.Hash(BlockHasher{}))
	}

	if b.Height != v.bc.Height() + 1 {
		return fmt.Errorf("bloktan (%s) yüksek", b.Hash(BlockHasher{}))
	}

	prevHeader, err := v.bc.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}

	hash := BlockHasher{}.Hash(prevHeader)
	if hash != b.PrevBlockHash {
		return fmt.Errorf("önceki bloğun hash'i (%s) geçersizdir.", b.PrevBlockHash)
	}

	if err := b.Verify(); err != nil{
		return err
	}

	return nil
}