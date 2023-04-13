package service

import (
	"assigment10/models"
	"assigment10/repository"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductFound(t *testing.T) {
	data := models.Product{
		GormModel: models.GormModel{ID: 1},
	}
	productRepository.Mock.On("FindById", "1").Return(data)

	result, err := productService.GetOneProduct("1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	// assert.Equal(t, "product not found", err.Error(), "error has be 'product not found'")
	assert.Equal(t, data.GormModel.ID, result.GormModel.ID, "result has be 1")
}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "2").Return(nil)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error should be 'product not found'")
}

func TestProductServiceGetAllProductFound(t *testing.T) {
	// Persiapan data
	products := []models.Product{
		{GormModel: models.GormModel{ID: 1}},
		{GormModel: models.GormModel{ID: 2}},
		{GormModel: models.GormModel{ID: 3}},
	}
	productRepository.Mock.On("FindAll").Return(products)

	// Pemanggilan fungsi yang diuji
	result, err := productService.GetAllProduct()

	// Asset bahwa tidak ada error
	assert.Nil(t, err)

	// Asset bahwa hasil tidak nil
	assert.NotNil(t, result)

	// Asset bahwa jumlah produk sesuai
	assert.Equal(t, len(products), len(result))

	// Asset bahwa setiap ID produk ada di hasil
	for _, product := range products {
		found := false
		for _, r := range result {
			if product.GormModel.ID == r.GormModel.ID {
				found = true
				break
			}
		}
		assert.True(t, found, fmt.Sprintf("Product with ID %d not found in result", product.GormModel.ID))
	}
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	// Persiapan data
	products := []models.Product{}
	productRepository.Mock.On("FindAll").Return(products)

	// Pemanggilan fungsi yang diuji
	result, err := productService.GetAllProduct()

	// Asset bahwa harus ada error
	assert.NotNil(t, err)

	// Asset bahwa hasil harus nil
	assert.Nil(t, result)

	// Asset bahwa error sesuai dengan yang diharapkan
	assert.Equal(t, "no products found", err.Error(), "error should be 'no products found'")
}
