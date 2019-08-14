package game

import (
	"github.com/matalmeida/shipbattle/field"
	"github.com/matalmeida/shipbattle/item"
)

const GAME_SIZE = 10

type Coordenadas struct {
	x int
	y int
}

type Jogo struct {
	Campo        *field.Campo
	CampoInimigo *field.Campo
	PilhaAtaques []Coordenadas
}

func Novo() *Jogo {
	j := &Jogo{Campo: field.Novo(GAME_SIZE), CampoInimigo: field.Novo(GAME_SIZE)}

	j.criaAtaqueSePilhaVazia()

	return j
}

func (j *Jogo) criaAtaqueSePilhaVazia() {
	j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: 0, y: 0})
}

func (j *Jogo) atk(x int, y int, tipo item.Nave) {
	i, iE := j.CampoInimigo.GetItem(x, y)
	if iE == nil && i.Id == "enemy" {
		return
	}
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

func (j *Jogo) sugereAtaque(x int, y int, tipo item.Nave) {
	switch tipo {
	case item.Destroyer:
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y + 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y - 1})
		break
	case item.Cruzador:
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y + 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y - 1})
		break
	case item.PortaAviao:
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y + 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y - 1})
		break
	case item.Hidroaviao:
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y + 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y - 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y + 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y - 1})
		break
	}
}

func (j *Jogo) checaSeGanhou() bool {
	return false
}

func (j *Jogo) RetornoDeAtaque(x int, y int, tipo item.Nave) bool {
	j.atk(x, y, tipo)
	j.sugereAtaque(x, y, tipo)

	switch tipo {
	case item.Submarino:
		j.atkVazioDiagonal(x, y)
		j.atkVazioCruz(x, y)
		break
	case item.Destroyer:
		j.atkVazioDiagonal(x, y)
		d, dE := j.CampoInimigo.GetItem(x+1, y)
		if dE == nil && d.Tipo != item.Vazio {
			j.atk(x-1, y, item.Vazio)
			j.atk(x+2, y, item.Vazio)
			break
		}
		e, eE := j.CampoInimigo.GetItem(x-1, y)
		if eE == nil && e.Tipo != item.Vazio {
			j.atk(x+1, y, item.Vazio)
			j.atk(x-2, y, item.Vazio)
			break
		}
		c, cE := j.CampoInimigo.GetItem(x, y+1)
		if cE == nil && c.Tipo != item.Vazio {
			j.atk(x, y-1, item.Vazio)
			j.atk(x, y+2, item.Vazio)
			break
		}
		b, bE := j.CampoInimigo.GetItem(x, y-1)
		if bE == nil && b.Tipo != item.Vazio {
			j.atk(x, y+1, item.Vazio)
			j.atk(x, y-2, item.Vazio)
			break
		}
		break
	case item.Cruzador:
		j.atkVazioDiagonal(x, y)
		// DIREITA
		d, dE := j.CampoInimigo.GetItem(x+1, y)
		d2, d2E := j.CampoInimigo.GetItem(x+2, y)
		d3, d3E := j.CampoInimigo.GetItem(x+3, y)
		// ESQUERDA
		e, eE := j.CampoInimigo.GetItem(x-1, y)
		e2, e2E := j.CampoInimigo.GetItem(x-2, y)
		e3, e3E := j.CampoInimigo.GetItem(x-3, y)

		if dE == nil && d.Tipo != item.Vazio {
			if d2E == nil && d2.Tipo != item.Vazio {
				if d3E == nil && d3.Tipo != item.Vazio {
					j.atk(x-1, y, item.Vazio)
					j.atk(x+4, y, item.Vazio)
					break
				}
				if eE == nil && e.Tipo != item.Vazio {
					j.atk(x-2, y, item.Vazio)
					j.atk(x+3, y, item.Vazio)
					break
				}
			}
			if eE == nil && e.Tipo != item.Vazio {
				if e2E == nil && e2.Tipo != item.Vazio {
					j.atk(x-3, y, item.Vazio)
					j.atk(x+2, y, item.Vazio)
					break
				}
			}
		}

		if eE == nil && e.Tipo != item.Vazio {
			if e2E == nil && e2.Tipo != item.Vazio {
				if e3E == nil && e3.Tipo != item.Vazio {
					j.atk(x-4, y, item.Vazio)
					j.atk(x+1, y, item.Vazio)
					break
				}
			}
		}

		// BAIXO
		b, bE := j.CampoInimigo.GetItem(x, y+1)
		b2, b2E := j.CampoInimigo.GetItem(x, y+2)
		b3, b3E := j.CampoInimigo.GetItem(x, y+3)
		// CIMA
		c, cE := j.CampoInimigo.GetItem(x, y-1)
		c2, c2E := j.CampoInimigo.GetItem(x, y-2)
		c3, c3E := j.CampoInimigo.GetItem(x, y-3)

		if bE == nil && b.Tipo != item.Vazio {
			if b2E == nil && b2.Tipo != item.Vazio {
				if b3E == nil && b3.Tipo != item.Vazio {
					j.atk(x, y-1, item.Vazio)
					j.atk(x, y+4, item.Vazio)
					break
				}
				if cE == nil && c.Tipo != item.Vazio {
					j.atk(x, y-2, item.Vazio)
					j.atk(x, y+3, item.Vazio)
					break
				}
			}
			if cE == nil && c.Tipo != item.Vazio {
				if c2E == nil && c2.Tipo != item.Vazio {
					j.atk(x, y-3, item.Vazio)
					j.atk(x, y+2, item.Vazio)
					break
				}
			}
		}

		if cE == nil && c.Tipo != item.Vazio {
			if c2E == nil && c2.Tipo != item.Vazio {
				if c3E == nil && c3.Tipo != item.Vazio {
					j.atk(x, y-4, item.Vazio)
					j.atk(x, y+1, item.Vazio)
					break
				}
			}
		}
		break
	case item.PortaAviao:
		j.atkVazioDiagonal(x, y)
		// DIREITA
		d, dE := j.CampoInimigo.GetItem(x+1, y)
		d2, d2E := j.CampoInimigo.GetItem(x+2, y)
		d3, d3E := j.CampoInimigo.GetItem(x+3, y)
		d4, d4E := j.CampoInimigo.GetItem(x+4, y)
		// ESQUERDA
		e, eE := j.CampoInimigo.GetItem(x-1, y)
		e2, e2E := j.CampoInimigo.GetItem(x-2, y)
		e3, e3E := j.CampoInimigo.GetItem(x-3, y)
		e4, e4E := j.CampoInimigo.GetItem(x-4, y)

		if dE == nil && d.Tipo != item.Vazio {
			if d2E == nil && d2.Tipo != item.Vazio {
				if d3E == nil && d3.Tipo != item.Vazio {
					if d4E == nil && d4.Tipo != item.Vazio {
						j.atk(x-1, y, item.Vazio)
						j.atk(x+5, y, item.Vazio)
						break
					}
					if eE == nil && e.Tipo != item.Vazio {
						j.atk(x-2, y, item.Vazio)
						j.atk(x+4, y, item.Vazio)
						break
					}
				}
				if eE == nil && e.Tipo != item.Vazio {
					if e2E == nil && e2.Tipo != item.Vazio {
						j.atk(x-3, y, item.Vazio)
						j.atk(x+3, y, item.Vazio)
						break
					}
				}
			}
			if eE == nil && e.Tipo != item.Vazio {
				if e2E == nil && e2.Tipo != item.Vazio {
					if e3E == nil && e3.Tipo != item.Vazio {
						j.atk(x-4, y, item.Vazio)
						j.atk(x+2, y, item.Vazio)
						break
					}
				}
			}
		}

		if eE == nil && e.Tipo != item.Vazio {
			if e2E == nil && e2.Tipo != item.Vazio {
				if e3E == nil && e3.Tipo != item.Vazio {
					if e4E == nil && e4.Tipo != item.Vazio {
						j.atk(x-5, y, item.Vazio)
						j.atk(x+1, y, item.Vazio)
						break
					}
				}
			}
		}

		// BAIXO
		b, bE := j.CampoInimigo.GetItem(x, y+1)
		b2, b2E := j.CampoInimigo.GetItem(x, y+2)
		b3, b3E := j.CampoInimigo.GetItem(x, y+3)
		b4, b4E := j.CampoInimigo.GetItem(x, y+4)
		// CIMA
		c, cE := j.CampoInimigo.GetItem(x, y-1)
		c2, c2E := j.CampoInimigo.GetItem(x, y-2)
		c3, c3E := j.CampoInimigo.GetItem(x, y-3)
		c4, c4E := j.CampoInimigo.GetItem(x, y-4)

		if bE == nil && b.Tipo != item.Vazio {
			if b2E == nil && b2.Tipo != item.Vazio {
				if b3E == nil && b3.Tipo != item.Vazio {
					if b4E == nil && b4.Tipo != item.Vazio {
						j.atk(x, y-1, item.Vazio)
						j.atk(x, y+5, item.Vazio)
						break
					}
					if cE == nil && c.Tipo != item.Vazio {
						j.atk(x, y-2, item.Vazio)
						j.atk(x, y+4, item.Vazio)
						break
					}
				}
				if cE == nil && c.Tipo != item.Vazio {
					if c2E == nil && c2.Tipo != item.Vazio {
						j.atk(x, y-3, item.Vazio)
						j.atk(x, y+3, item.Vazio)
						break
					}
				}
			}
			if cE == nil && c.Tipo != item.Vazio {
				if c2E == nil && c2.Tipo != item.Vazio {
					if c2E == nil && c2.Tipo != item.Vazio {
						j.atk(x, y-4, item.Vazio)
						j.atk(x, y+2, item.Vazio)
						break
					}
				}
			}
		}

		if cE == nil && c.Tipo != item.Vazio {
			if c2E == nil && c2.Tipo != item.Vazio {
				if c3E == nil && c3.Tipo != item.Vazio {
					if c4E == nil && c4.Tipo != item.Vazio {
						j.atk(x, y-5, item.Vazio)
						j.atk(x, y+1, item.Vazio)
						break
					}
				}
			}
		}

		break
	case item.Hidroaviao:
		j.atkVazioCruz(x, y)
		// CIMA
		c, cE := j.CampoInimigo.GetItem(x, y-2)
		cEsq, cEsqE := j.CampoInimigo.GetItem(x-1, y-1)
		cDir, cDirE := j.CampoInimigo.GetItem(x+1, y-1)
		// BAIXO
		b, bE := j.CampoInimigo.GetItem(x, y+2)
		bEsq, bEsqE := j.CampoInimigo.GetItem(x-1, y+1)
		bDir, bDirE := j.CampoInimigo.GetItem(x+1, y+1)
		// LATERAL
		d, dE := j.CampoInimigo.GetItem(x+2, y)
		e, eE := j.CampoInimigo.GetItem(x-2, y)

		if cE == nil && c.Tipo != item.Vazio {
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x, y-2)
				j.atkVazioDiagonal(x-1, y-1)
				break
			}
			if cDirE == nil && cDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x, y-2)
				j.atkVazioDiagonal(x+1, y-1)
				break
			}
		}

		if bE == nil && b.Tipo != item.Vazio {
			if bEsqE == nil && bEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x, y+2)
				j.atkVazioDiagonal(x-1, y+1)
				break
			}
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x, y+2)
				j.atkVazioDiagonal(x+1, y+1)
				break
			}
		}

		if dE == nil && d.Tipo != item.Vazio {
			if cDirE == nil && cDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x+2, y)
				j.atkVazioDiagonal(x+1, y-1)
				break
			}
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x+2, y)
				j.atkVazioDiagonal(x+1, y+1)
				break
			}
		}

		if eE == nil && e.Tipo != item.Vazio {
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x-2, y)
				j.atkVazioDiagonal(x-1, y-1)
				break
			}
			if bEsqE == nil && bEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x-2, y)
				j.atkVazioDiagonal(x-1, y+1)
				break
			}
		}

		if cDirE == nil && cDir.Tipo != item.Vazio {
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x+1, y-1)
				j.atkVazioDiagonal(x-1, y-1)
			}
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x+1, y-1)
				j.atkVazioDiagonal(x+1, y+1)
			}
		}

		if bEsqE == nil && bEsq.Tipo != item.Vazio {
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x-1, y+1)
				j.atkVazioDiagonal(x+1, y+1)
			}
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x-1, y+1)
				j.atkVazioDiagonal(x-1, y-1)
			}
		}

		break
	}

	if j.checaSeGanhou() {
		return true
	} else {
		return false
	}

}
