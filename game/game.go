package game

import (
	"github.com/matalmeida/shipbattle/field"
	"github.com/matalmeida/shipbattle/item"
)

const GAME_SIZE = 10

type Jogo struct {
	Campo        *field.Campo
	CampoInimigo *field.Campo
}

func Novo() *Jogo {
	g := &Jogo{Campo: field.Novo(GAME_SIZE), CampoInimigo: field.Novo(GAME_SIZE)}

	return g
}

func (j *Jogo) atk(x int, y int, tipo item.Nave) {
	j.CampoInimigo.ColocaItem(x, y, "enemy", tipo)
	j.CampoInimigo.Atacar(x, y)
}

func (j *Jogo) atkVazioCruz(x int, y int) {
	j.atk(x+1, y, item.Vazio)
	j.atk(x-1, y, item.Vazio)
	j.atk(x, y+1, item.Vazio)
	j.atk(x, y-1, item.Vazio)
}

func (j *Jogo) atkVazioDiagonal(x int, y int) {
	j.atk(x+1, y+1, item.Vazio)
	j.atk(x+1, y-1, item.Vazio)
	j.atk(x-1, y+1, item.Vazio)
	j.atk(x-1, y-1, item.Vazio)
}

func (j *Jogo) checaSeGanhou() bool {
	return false
}

func (j *Jogo) RetornoDeAtaque(x int, y int, tipo item.Nave) bool {
	j.atk(x, y, tipo)

	switch tipo {
	case item.Submarino:
		j.atkVazioDiagonal(x, y)
		j.atkVazioCruz(x, y)
		break
	case item.Destroyer:
		j.atkVazioDiagonal(x, y)
		d, dE := j.CampoInimigo.GetItem(x+1, y)
		if dE == nil && d.Tipo == item.Destroyer {
			j.atk(x-1, y, item.Vazio)
			j.atk(x+2, y, item.Vazio)
			break
		}
		e, eE := j.CampoInimigo.GetItem(x-1, y)
		if eE == nil && e.Tipo == item.Destroyer {
			j.atk(x+1, y, item.Vazio)
			j.atk(x-2, y, item.Vazio)
			break
		}
		c, cE := j.CampoInimigo.GetItem(x, y+1)
		if cE == nil && c.Tipo == item.Destroyer {
			j.atk(x, y-1, item.Vazio)
			j.atk(x, y+2, item.Vazio)
			break
		}
		b, bE := j.CampoInimigo.GetItem(x, y-1)
		if bE == nil && b.Tipo == item.Destroyer {
			j.atk(x, y+1, item.Vazio)
			j.atk(x, y-2, item.Vazio)
			break
		}
		break
	case item.Cruzador:
		j.atkVazioDiagonal(x, y)
		break
	case item.PortaAviao:
		j.atkVazioDiagonal(x, y)
		break
	case item.Hidroaviao:
		j.atkVazioCruz(x, y)
		break
	}

	if j.checaSeGanhou() {
		return true
	} else {
		return false
	}

}
