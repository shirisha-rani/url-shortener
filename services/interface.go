package services

import "url-shortener/models"

type URL interface {
	GetShortURL(models.URL) (string, error)
}
