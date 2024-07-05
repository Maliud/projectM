package core

import (
	"fmt"
	"sync"

	"github.com/go-kit/log"
)

// 10 tx [1 -4] 5 'de hata
// s => `S
type Blockchain struct {
	logger    log.Logger
	store     Storage
	lock      sync.RWMutex
	headers   []*Header
	validator Validator
}

func NewBlockchain(l log.Logger, genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemorystore(),
		logger: l,
	}
	bc.validator = NewBlockValidator(bc)
	err := bc.addBlockWithoutValidation(genesis)
	return bc, err
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}
	return bc.addBlockWithoutValidation(b)

}

// [0,1]
// len(arr) 2

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("verilen Yükseklik(%d) çok yüksek", height)
	}
	bc.lock.Lock()
	defer bc.lock.Unlock()

	return bc.headers[height], nil

}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

// [0, 1, 2, 3] => 4
// [0, 1, 2, 3] => 3 height
func (bc *Blockchain) Height() uint32 {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.lock.Lock()
	bc.headers = append(bc.headers, b.Header)
	bc.lock.Unlock()


	bc.logger.Log(
		"msg", "Yeni Blok",
		"hash", b.Hash(BlockHasher{}),
		"Yükseklik", b.Height,
		"İŞLEMLER", len(b.Transactions),
	)

	return bc.store.Put(b)

}
