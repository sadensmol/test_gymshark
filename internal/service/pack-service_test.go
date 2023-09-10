package service_test

import (
	"testing"

	"github.com/sadensmol/test_gymshark/internal/domain"
	packservice "github.com/sadensmol/test_gymshark/internal/service"
	packstore "github.com/sadensmol/test_gymshark/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestPackService_PackItems(t *testing.T) {
	store, err := packstore.NewPackStore()
	assert.NoError(t, err)

	service, err := packservice.NewPackService(store)

	tests := []struct {
		name          string
		numberOfItems int
		expectedPacks []domain.Pack
	}{
		{
			name:          "1 item packs into one 250-size pack",
			numberOfItems: 1,
			expectedPacks: []domain.Pack{
				{Size: 250},
			},
		},
		{
			name:          "250 items packs into one 250-size pack",
			numberOfItems: 250,
			expectedPacks: []domain.Pack{
				{Size: 250},
			},
		},
		{
			name:          "251 items packs into one 500-size pack",
			numberOfItems: 251,
			expectedPacks: []domain.Pack{
				{Size: 500},
			},
		},
		{
			name:          "501 items packs into two 500-size and 250-size packs",
			numberOfItems: 501,
			expectedPacks: []domain.Pack{
				{Size: 500},
				{Size: 250},
			},
		},
		{
			name:          "12001 items packs into three 5000-size, 2000-size and 250-size packs",
			numberOfItems: 12001,
			expectedPacks: []domain.Pack{
				{Size: 5000},
				{Size: 5000},
				{Size: 2000},
				{Size: 250},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packs, err := service.PackItems(tt.numberOfItems)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedPacks, packs)
		})
	}
}
