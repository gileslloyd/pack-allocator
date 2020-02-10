package inmemoryrepo

import "github.com/gileslloyd/gs-allocation-service/internal/domain/pack"

type Pack struct {
}

func NewInMemoryPackRepo() pack.Repository {
	return Pack{}
}

func (p Pack) GetAll() []pack.Entity {
	return []pack.Entity{
		pack.NewPackEntity(250),
		pack.NewPackEntity(500),
	}
}
