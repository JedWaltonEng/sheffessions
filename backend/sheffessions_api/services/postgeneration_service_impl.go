package services

import (
	"errors"
	"sheffessions/api/store"
)

type postService struct {
	confessionStore store.StorerConfessions
	postedStore     store.StorerPostedConfessions
}

func NewPostGenerationService(confessionStore store.StorerConfessions, postedConfessionStore store.StorerPostedConfessions) *postService {
	return &postService{confessionStore, postedConfessionStore}
}

func (p *postService) GeneratePost() (*store.Confession, error) {
	confession, err := p.confessionStore.RandomConfession()
	if err != nil {
		return nil, errors.New("Error getting random confession: " + err.Error())
	}
	// Utilize methods from the store here if required, like checking if posted.
	return confession, nil
}

func (p *postService) IsPostDuplicate() bool {
	// If you want, you can use methods from p.store here
	// to check if the post is a duplicate.
	return false
}

// You can add the rest of your business logic here:
// - Delete from db if so
// - Check if confession is a duplicate
// - Check if the confession is empty
// - Check if confession goes against Instagram TOS
// - Check if confession is too long
// - Store posted confession in posted confessions db.
// - Generate post image
// - Make post to Instagram
