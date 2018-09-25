package generator

import "github.com/satori/go.uuid"

type Generator interface {
	GenerateUUID() string
}

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (g *UUIDGenerator) GenerateUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
