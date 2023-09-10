package store_test

import (
	"testing"

	"github.com/sadensmol/test_gymshark/internal/domain"
	packstore "github.com/sadensmol/test_gymshark/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestValidPackStore(t *testing.T) {
	store, err := packstore.NewPackStore()
	assert.NoError(t, err)

	packs := store.AvailablePacks()
	assert.ElementsMatch(t, domain.Packs{
		{Size: 250},
		{Size: 500},
		{Size: 1000},
		{Size: 2000},
		{Size: 5000},
	}, packs)
}
func TestImmutability(t *testing.T) {
	store, err := packstore.NewPackStore()
	assert.NoError(t, err)

	packs := store.AvailablePacks()
	packs[0] = domain.Pack{Size: 1}
	assert.ElementsMatch(t, domain.Packs{
		{Size: 250},
		{Size: 500},
		{Size: 1000},
		{Size: 2000},
		{Size: 5000},
	}, store.AvailablePacks())
}

func TestAddNewPack(t *testing.T) {
	ps, err := packstore.NewPackStore()
	ps.AddPack(10000)
	assert.NoError(t, err)
	assert.ElementsMatch(t, domain.Packs{
		{Size: 250},
		{Size: 500},
		{Size: 1000},
		{Size: 2000},
		{Size: 5000},
		{Size: 10000},
	}, ps.AvailablePacks())
}
