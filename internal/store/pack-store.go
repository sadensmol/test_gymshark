package packstore

import (
	"sync"

	"github.com/sadensmol/test_gymshark/internal/domain"
)

// PackStore is a simple in-memory store.
// It provides list of available pack sizes.
type PackStore struct {
	packs []domain.Pack
	mutex sync.RWMutex
}

func NewPackStore() (*PackStore, error) {
	store := &PackStore{}
	// default packs
	store.AddPack(250)
	store.AddPack(500)
	store.AddPack(1000)
	store.AddPack(2000)
	store.AddPack(5000)
	return store, nil
}

func (ps *PackStore) AddPack(size int) error {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.packs = append(ps.packs, domain.Pack{Size: size})
	return nil
}

func (ps *PackStore) AvailablePacks() []domain.Pack {
	r := make([]domain.Pack, len(ps.packs))
	ps.mutex.RLock()
	defer ps.mutex.RUnlock()
	copy(r, ps.packs)
	return r
}
