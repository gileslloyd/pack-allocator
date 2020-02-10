package microrepo

import "github.com/gileslloyd/gs-allocation-service/internal/domain/pack"

type Pack struct {
}

func NewMicroPackRepo() pack.Repository {
	return Pack{}
}

func (p Pack) GetAll() []pack.Entity {
	return []pack.Entity{
		pack.NewPackEntity(250),
		pack.NewPackEntity(500),
	}
}
