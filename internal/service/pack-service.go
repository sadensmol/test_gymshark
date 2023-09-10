package service

import (
	"sort"

	"github.com/sadensmol/test_gymshark/internal/domain"
)

type IPackStore interface {
	AvailablePacks() []domain.Pack
	RemoveBySize(size int) error
	AddPack(size int) error
}

type PackService struct {
	store IPackStore
}

func NewPackService(store IPackStore) (*PackService, error) {
	return &PackService{store: store}, nil
}

func (ps *PackService) AvailablePacks() domain.Packs {
	return ps.store.AvailablePacks()
}

func (ps *PackService) Remove(size int) error {
	return ps.store.RemoveBySize(size)
}

func (ps *PackService) Add(size int) error {
	return ps.store.AddPack(size)
}

// PackItems takes a number of items and returns a slice of packs.
// The following rules apply:
// 1. Only whole packs can be sent. Packs cannot be broken open.
// 2. Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order.
// 3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.
func (ps *PackService) PackItems(numberOfItems int) (domain.Packs, error) {
	availablePacks := ps.store.AvailablePacks()

	sort.Slice(availablePacks, func(i, j int) bool {
		return availablePacks[i].Size > availablePacks[j].Size
	})

	var result domain.Packs
	minDiff := -1

	var recSearch func(int, int, domain.Packs)
	recSearch = func(index int, totalSize int, packs domain.Packs) {
		diff := totalSize - numberOfItems
		if diff >= 0 && (result == nil || diff < minDiff) {
			result = packs
			minDiff = diff
		}

		if index >= len(availablePacks) || totalSize >= numberOfItems {
			return
		}

		recSearch(index, totalSize+availablePacks[index].Size, append(packs, availablePacks[index]))
		recSearch(index+1, totalSize, packs)
	}

	recSearch(0, 0, domain.Packs{})
	return result, nil
}
