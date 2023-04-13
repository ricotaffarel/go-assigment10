package repository

import (
	"assigment10/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *models.Product {
	argument := repository.Mock.Called(id)
	if argument.Get(0) == nil {
		return nil
	}
	product := argument.Get(0).(models.Product)

	return &product
}

func (repository *ProductRepositoryMock) FindAll() []models.Product {
	argument := repository.Mock.Called()
	if argument.Get(0) == nil {
		return []models.Product{}
	}
	products := argument.Get(0).([]models.Product)

	return products
}
