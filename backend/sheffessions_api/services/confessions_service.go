package services

import "sheffessions/api/store"

type ConfessionService interface {
	SaveConfession(content, source string) (int64, error)
	RandomConfession() (*store.Confession, error)
	DeleteConfessionByID(id int64) error
}
