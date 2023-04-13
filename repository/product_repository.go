package repository

import "assigment10/models"

type ProductRepository interface {
	FindById(id string) *models.Product
	FindAll() []models.Product
}
