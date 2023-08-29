package services

import (
	"errors"
	"sheffessions/api/store"
)

type postGenerationServiceImpl struct {
	db store.StorerConfessions
}

func NewPostGenerationService(db store.StorerConfessions) PostGenerationService {
	return &postGenerationServiceImpl{db: db}
}

func (s *postGenerationServiceImpl) GeneratePost() (*store.Confession, error) {
	confession, err := s.db.RandomConfession()
	if err != nil {
		return nil, errors.New("Error getting random confession: " + err.Error())
	}

	// You can add the rest of your business logic here:
	// - Delete from db if so
	// - Check if confession is a duplicate
	// - Check if the confession is empty
	// - Check if confession goes against Instagram TOS
	// - Check if confession is too long
	// - Store posted confession in db.
	// - Generate post image
	// - Make post to Instagram

	return confession, nil
}

func isPostDuplicate() bool {
	return false
}
