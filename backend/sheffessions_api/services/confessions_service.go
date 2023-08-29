package services

import "sheffessions/api/store"

type ConfessionService interface {
	CreateConfession(content, source string) (int64, error)
	GetRandomConfession() (*store.Confession, error)
	RemoveConfession(id int64) error
}
