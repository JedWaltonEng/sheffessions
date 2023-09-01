package services

import (
	"sheffessions/api/store"
	"testing"
)

type ConfessionsStoreStub struct {
	saveCalled       bool
	randomCalled     bool
	deleteByIDCalled bool
	forcedError      error
}

type PostedConfessionsStoreStub struct {
	MarkConfessionAsPublishedCalled bool
	IsConfessionPublishedCalled     bool
	forcedError                     error
}

func (s *ConfessionsStoreStub) SaveConfession(content, source string) (int64, error) {
	s.saveCalled = true
	return 1, s.forcedError // return dummy data
}

func (s *ConfessionsStoreStub) RandomConfession() (*store.Confession, error) {
	s.randomCalled = true
	return &store.Confession{}, s.forcedError // return dummy data
}

func (s *ConfessionsStoreStub) DeleteConfessionByID(id int64) error {
	s.deleteByIDCalled = true
	return s.forcedError
}

func (s *PostedConfessionsStoreStub) MarkConfessionAsPublished(id int64) error {
	s.MarkConfessionAsPublishedCalled = true
	return s.forcedError
}

func (s *PostedConfessionsStoreStub) IsConfessionPublished(id int64) (bool, error) {
	s.IsConfessionPublishedCalled = true
	return true, s.forcedError
}

func TestGeneratePost(t *testing.T) {
	confessionsStoreStub := &ConfessionsStoreStub{}
	postedConfessionsStoreStub := &PostedConfessionsStoreStub{}
	service := NewPostGenerationService(confessionsStoreStub, postedConfessionsStoreStub)

	service.GeneratePost()

	if !confessionsStoreStub.randomCalled {
		t.Error("Expected SaveConfession to be called on the store")
	}
}
