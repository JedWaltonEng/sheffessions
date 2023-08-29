package services

import "sheffessions/api/store"

type service struct {
	store store.StorerConfessions
}

func NewConfessionService(s store.StorerConfessions) ConfessionService {
	return &service{s}
}

func (s *service) CreateConfession(content, source string) (int64, error) {
	return s.store.SaveConfession(content, source)
}

func (s *service) GetRandomConfession() (*store.Confession, error) {
	return s.store.RandomConfession()
}

func (s *service) RemoveConfession(id int64) error {
	return s.store.DeleteConfessionByID(id)
}
