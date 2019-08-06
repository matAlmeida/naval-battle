package game

import (
	"github.com/matalmeida/shipbattle/field"
)

const GAME_SIZE = 10

type Jogo struct {
	Campo *field.Campo
}

func Novo() *Jogo {
	g := &Jogo{Campo: field.Novo(GAME_SIZE)}

	return g
}
