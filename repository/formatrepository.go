package repository

import "agora/core/models"

type FormatRepository interface {
	Reverse(model *models.Transport) (*models.Transport, error)
	Upper(model *models.Transport) (*models.Transport, error)
	Lower(model *models.Transport) (*models.Transport, error)
}
