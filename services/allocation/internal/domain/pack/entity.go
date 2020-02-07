package pack

type Entity struct {
	size int
}

func NewPackEntity(size int) Entity {
	return Entity{size: size}
}
