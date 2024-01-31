package tests

import (
	"test/domain"
	"test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {

	testProduct := domain.Product{
		ProductName: "Shirt",
		Price:       1250,
		Catergory:   "men",
	}

	err := repository.GetProductByName(testProduct)
	assert.NoError(t, err, "Expected no error")

	err = repository.CreateProduct(testProduct)
	assert.NoError(t, err, "Expected no error")

	err = repository.DeleteProduct(testProduct)
	assert.NoError(t, err, "Expected no error")

}

func TestGetAllProduct(t *testing.T) {

	expectedProducts := []domain.Product{
		{ProductName: "T-shirt", Price: 1250.00, Catergory: "men"},
	}

	list, err := repository.GetAllProduct()
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedProducts, list, "Returned products matched expected products")

}
