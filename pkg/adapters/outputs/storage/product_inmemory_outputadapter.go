package storage

import (
	"context"
	"errors"

	"github.com/landistas/enterprise-go-examples/pkg/entities"
	"github.com/landistas/enterprise-go-examples/pkg/usecases"
)

var _ usecases.ProductStorageAdapter = (*ProductStorageAdapter)(nil)

type ProductStorageAdapter struct {
	infraStorage InfraStorage
}

func NewProductStorageAdapter(infraStorage InfraStorage) *ProductStorageAdapter {
	return &ProductStorageAdapter{
		infraStorage: infraStorage,
	}
}

func (adapter *ProductStorageAdapter) Add(ctx context.Context, product entities.Product) (*entities.Product, error) {
	return &product, adapter.infraStorage.Save(ctx, product.ID, product)
}

func (adapter ProductStorageAdapter) Get(ctx context.Context, productID string) (*entities.Product, error) {
	product, err := adapter.infraStorage.Read(ctx, productID)
	if err != nil {
		return nil, err
	}
	
	if product, ok := product.(entities.Product); ok {
		return &product, nil
	}
	return nil, errors.New("casting: can not casting product from storage")
}

func (adapter ProductStorageAdapter) List(ctx context.Context, filterOptions entities.CatalogFilterOptions) (*entities.Catalog, error) {
	// TODO implement

	return nil, nil
}
