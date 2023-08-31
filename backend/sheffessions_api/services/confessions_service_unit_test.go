package services

import (
	"sheffessions/api/store"
	"testing"
)

type StoreStub struct {
	saveCalled       bool
	randomCalled     bool
	deleteByIDCalled bool
	forcedError      error
}

func (s *StoreStub) SaveConfession(content, source string) (int64, error) {
	s.saveCalled = true
	return 1, s.forcedError // return dummy data
}

func (s *StoreStub) RandomConfession() (*store.Confession, error) {
	s.randomCalled = true
	return &store.Confession{}, s.forcedError // return dummy data
}

func (s *StoreStub) DeleteConfessionByID(id int64) error {
	s.deleteByIDCalled = true
	return s.forcedError
}

func TestSaveConfession(t *testing.T) {
	stub := &StoreStub{}
	service := NewConfessionService(stub)

	service.SaveConfession("testContent", "testSource")

	if !stub.saveCalled {
		t.Error("Expected SaveConfession to be called on the store")
	}
}

func TestRandomConfession(t *testing.T) {
	stub := &StoreStub{}
	service := NewConfessionService(stub)

	service.RandomConfession()

	if !stub.randomCalled {
		t.Error("Expected RandomConfession to be called on the store")
	}
}

func TestDeleteConfessionByID(t *testing.T) {
	stub := &StoreStub{}
	service := NewConfessionService(stub)

	service.DeleteConfessionByID(1)

	if !stub.deleteByIDCalled {
		t.Error("Expected DeleteConfessionByID to be called on the store")
	}
}
