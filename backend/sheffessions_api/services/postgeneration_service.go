package services

import "sheffessions/api/store"

type PostGenerationService interface {
	GeneratePost() (*store.Confession, error)
}
