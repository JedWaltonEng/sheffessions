package services

import "sheffessions/api/store"

type service struct {
	store store.StorerConfessions
}

func NewConfessionService(s store.StorerConfessions) ConfessionService {
	return &service{s}
}

func (s *service) SaveConfession(content, source string) (int64, error) {
	return s.store.SaveConfession(content, source)
}

func (s *service) RandomConfession() (*store.Confession, error) {
	return s.store.RandomConfession()
}

func (s *service) DeleteConfessionByID(id int64) error {
	return s.store.DeleteConfessionByID(id)
}
