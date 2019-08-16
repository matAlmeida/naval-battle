package game

import (
	"fmt"

	"github.com/fatih/color"
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
	CopiaCampo   *field.Campo
	CampoInimigo *field.Campo
	PilhaAtaques []Coordenadas
	NavesInimigo map[item.Nave]int
	NossasNaves  map[item.Nave]int
}

func Novo() *Jogo {
	j := &Jogo{Campo: field.Novo(GAME_SIZE), CopiaCampo: field.Novo(GAME_SIZE), CampoInimigo: field.Novo(GAME_SIZE)}

	j.NavesInimigo = make(map[item.Nave]int)
	j.NossasNaves = make(map[item.Nave]int)

	j.NavesInimigo[item.Hidroaviao] = 4
	j.NavesInimigo[item.Submarino] = 4
	j.NavesInimigo[item.Destroyer] = 3
	j.NavesInimigo[item.Cruzador] = 2
	j.NavesInimigo[item.PortaAviao] = 1

	j.NossasNaves[item.Hidroaviao] = 4
	j.NossasNaves[item.Submarino] = 4
	j.NossasNaves[item.Destroyer] = 3
	j.NossasNaves[item.Cruzador] = 2
	j.NossasNaves[item.PortaAviao] = 1

	j.iniciaPilhaDeAtaque()

	return j
}

func (j *Jogo) Atacar() (int, int) {
	nextAtk := j.PilhaAtaques[len(j.PilhaAtaques)-1]
	j.PilhaAtaques = j.PilhaAtaques[:len(j.PilhaAtaques)-1]
	for {
		alvo, alvoE := j.CampoInimigo.GetItem(nextAtk.x, nextAtk.y)
		// fmt.Printf("cood: %#v - alvo: %#v\n", nextAtk, alvo)
		if alvoE == nil && !alvo.Bombardeado {
			return nextAtk.x, nextAtk.y
		}
		nextAtk = j.PilhaAtaques[len(j.PilhaAtaques)-1]
		j.PilhaAtaques = j.PilhaAtaques[:len(j.PilhaAtaques)-1]
	}
}

func (j *Jogo) SerAtacado(x int, y int) *item.Item {
	successo, alvo := j.Campo.Atacar(x, y)
	if successo {
		return alvo
	}

	return alvo
}

func (j *Jogo) iniciaPilhaDeAtaque() {
	for x := 9; x >= 0; x-- {
		for y := 9; y >= 0; y-- {
			j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y})
		}
	}
}

func (j *Jogo) atk(x int, y int, tipo item.Nave, o ...string) {
	i, iE := j.CampoInimigo.GetItem(x, y)
	if len(o) != 0 {
		i, iE = j.CopiaCampo.GetItem(x, y)
		if iE == nil && i.Id == "enemy" {
			return
		}
		j.CopiaCampo.ColocaItem(x, y, "enemy", tipo)
		j.CopiaCampo.Atacar(x, y)
		return
	}

	if iE == nil && i.Id == "enemy" {
		return
	}
	j.CampoInimigo.ColocaItem(x, y, "enemy", tipo)
	j.CampoInimigo.Atacar(x, y)
}

func (j *Jogo) atkVazioCruz(x int, y int, o ...string) {
	if len(o) != 0 {
		j.atk(x+1, y, item.Vazio, o[0])
		j.atk(x-1, y, item.Vazio, o[0])
		j.atk(x, y+1, item.Vazio, o[0])
		j.atk(x, y-1, item.Vazio, o[0])
		return
	}
	j.atk(x+1, y, item.Vazio)
	j.atk(x-1, y, item.Vazio)
	j.atk(x, y+1, item.Vazio)
	j.atk(x, y-1, item.Vazio)
}

func (j *Jogo) atkVazioDiagonal(x int, y int, o ...string) {
	if len(o) != 0 {
		j.atk(x+1, y+1, item.Vazio, o[0])
		j.atk(x+1, y-1, item.Vazio, o[0])
		j.atk(x-1, y+1, item.Vazio, o[0])
		j.atk(x-1, y-1, item.Vazio, o[0])
		return
	}
	j.atk(x+1, y+1, item.Vazio)
	j.atk(x+1, y-1, item.Vazio)
	j.atk(x-1, y+1, item.Vazio)
	j.atk(x-1, y-1, item.Vazio)
}

func (j *Jogo) sugereAtaque(x int, y int, tipo item.Nave) {
	switch tipo {
	case item.Destroyer:
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y - 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y + 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y})
		break
	case item.Cruzador:
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y - 1})
		e, eE := j.CampoInimigo.GetItem(x, y-1)
		if y < 7 || (eE == nil && e.Tipo != item.Vazio) {
			j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y + 1})
		}
		c, cE := j.CampoInimigo.GetItem(x-1, y)
		if x < 7 || (cE == nil && c.Tipo != item.Vazio) {
			j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y})
		}
		break
	case item.PortaAviao:
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y - 1})
		e, eE := j.CampoInimigo.GetItem(x, y-1)
		if y < 6 || (eE == nil && e.Tipo != item.Vazio) {
			j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y + 1})
		}
		c, cE := j.CampoInimigo.GetItem(x-1, y)
		if x < 6 || (cE == nil && c.Tipo != item.Vazio) {
			j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y})
		}
		break
	case item.Hidroaviao:
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 2, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 2, y: y})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y - 2})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x, y: y + 2})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y - 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y - 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x + 1, y: y + 1})
		j.PilhaAtaques = append(j.PilhaAtaques, Coordenadas{x: x - 1, y: y + 1})
		break
	}
}

func (j *Jogo) checaSeGanhou() bool {
	ganhou := true
	for _, qntd := range j.NavesInimigo {
		if qntd != 0 {
			ganhou = ganhou && false
		}
	}
	return ganhou
}

func (j *Jogo) checaSePerdeu() bool {
	perdeu := true
	for _, qntd := range j.NossasNaves {
		if qntd != 0 {
			perdeu = perdeu && false
		}
	}
	return perdeu
}

func (j *Jogo) RetornoDeAtaque(x int, y int, tipo item.Nave) bool {
	j.atk(x, y, tipo)
	if tipo == item.Hidroaviao {
		bD, bDE := j.CampoInimigo.GetItem(x+1, y+1)
		cD, cDE := j.CampoInimigo.GetItem(x+1, y-1)
		bE, bEE := j.CampoInimigo.GetItem(x-1, y+1)
		cE, cEE := j.CampoInimigo.GetItem(x-1, y-1)

		sugere := true
		if bDE == nil && bD.Tipo == item.Hidroaviao {
			sugere = false
		} else if cDE == nil && cD.Tipo == item.Hidroaviao {
			sugere = false
		} else if bEE == nil && bE.Tipo == item.Hidroaviao {
			sugere = false
		} else if cEE == nil && cE.Tipo == item.Hidroaviao {
			sugere = false
		}

		if sugere {
			j.sugereAtaque(x, y, tipo)
		}
	} else {
		j.sugereAtaque(x, y, tipo)
	}

	switch tipo {
	case item.Submarino:
		j.atkVazioDiagonal(x, y)
		j.atkVazioCruz(x, y)
		j.NavesInimigo[item.Submarino]--

		break
	case item.Destroyer:
		j.atkVazioDiagonal(x, y)
		d, dE := j.CampoInimigo.GetItem(x+1, y)
		if dE == nil && d.Tipo != item.Vazio {
			j.atk(x-1, y, item.Vazio)
			j.atk(x+2, y, item.Vazio)
			j.NavesInimigo[item.Destroyer]--

			break
		}
		e, eE := j.CampoInimigo.GetItem(x-1, y)
		if eE == nil && e.Tipo != item.Vazio {
			j.atk(x+1, y, item.Vazio)
			j.atk(x-2, y, item.Vazio)
			j.NavesInimigo[item.Destroyer]--

			break
		}
		c, cE := j.CampoInimigo.GetItem(x, y+1)
		if cE == nil && c.Tipo != item.Vazio {
			j.atk(x, y-1, item.Vazio)
			j.atk(x, y+2, item.Vazio)
			j.NavesInimigo[item.Destroyer]--

			break
		}
		b, bE := j.CampoInimigo.GetItem(x, y-1)
		if bE == nil && b.Tipo != item.Vazio {
			j.atk(x, y+1, item.Vazio)
			j.atk(x, y-2, item.Vazio)
			j.NavesInimigo[item.Destroyer]--

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
					j.NavesInimigo[item.Cruzador]--

					break
				}
				if eE == nil && e.Tipo != item.Vazio {
					j.atk(x-2, y, item.Vazio)
					j.atk(x+3, y, item.Vazio)
					j.NavesInimigo[item.Cruzador]--

					break
				}
			}
			if eE == nil && e.Tipo != item.Vazio {
				if e2E == nil && e2.Tipo != item.Vazio {
					j.atk(x-3, y, item.Vazio)
					j.atk(x+2, y, item.Vazio)
					j.NavesInimigo[item.Cruzador]--

					break
				}
			}
		}

		if eE == nil && e.Tipo != item.Vazio {
			if e2E == nil && e2.Tipo != item.Vazio {
				if e3E == nil && e3.Tipo != item.Vazio {
					j.atk(x-4, y, item.Vazio)
					j.atk(x+1, y, item.Vazio)
					j.NavesInimigo[item.Cruzador]--

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
					j.NavesInimigo[item.Cruzador]--

					break
				}
				if cE == nil && c.Tipo != item.Vazio {
					j.atk(x, y-2, item.Vazio)
					j.atk(x, y+3, item.Vazio)
					j.NavesInimigo[item.Cruzador]--

					break
				}
			}
			if cE == nil && c.Tipo != item.Vazio {
				if c2E == nil && c2.Tipo != item.Vazio {
					j.atk(x, y-3, item.Vazio)
					j.atk(x, y+2, item.Vazio)
					j.NavesInimigo[item.Cruzador]--

					break
				}
			}
		}

		if cE == nil && c.Tipo != item.Vazio {
			if c2E == nil && c2.Tipo != item.Vazio {
				if c3E == nil && c3.Tipo != item.Vazio {
					j.atk(x, y-4, item.Vazio)
					j.atk(x, y+1, item.Vazio)
					j.NavesInimigo[item.Cruzador]--

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
						j.NavesInimigo[item.PortaAviao]--

						break
					}
					if eE == nil && e.Tipo != item.Vazio {
						j.atk(x-2, y, item.Vazio)
						j.atk(x+4, y, item.Vazio)
						j.NavesInimigo[item.PortaAviao]--

						break
					}
				}
				if eE == nil && e.Tipo != item.Vazio {
					if e2E == nil && e2.Tipo != item.Vazio {
						j.atk(x-3, y, item.Vazio)
						j.atk(x+3, y, item.Vazio)
						j.NavesInimigo[item.PortaAviao]--

						break
					}
				}
			}
			if eE == nil && e.Tipo != item.Vazio {
				if e2E == nil && e2.Tipo != item.Vazio {
					if e3E == nil && e3.Tipo != item.Vazio {
						j.atk(x-4, y, item.Vazio)
						j.atk(x+2, y, item.Vazio)
						j.NavesInimigo[item.PortaAviao]--

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
						j.NavesInimigo[item.PortaAviao]--

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
						j.NavesInimigo[item.PortaAviao]--

						break
					}
					if cE == nil && c.Tipo != item.Vazio {
						j.atk(x, y-2, item.Vazio)
						j.atk(x, y+4, item.Vazio)
						j.NavesInimigo[item.PortaAviao]--

						break
					}
				}
				if cE == nil && c.Tipo != item.Vazio {
					if c2E == nil && c2.Tipo != item.Vazio {
						j.atk(x, y-3, item.Vazio)
						j.atk(x, y+3, item.Vazio)
						j.NavesInimigo[item.PortaAviao]--

						break
					}
				}
			}
			if cE == nil && c.Tipo != item.Vazio {
				if c2E == nil && c2.Tipo != item.Vazio {
					if c2E == nil && c2.Tipo != item.Vazio {
						j.atk(x, y-4, item.Vazio)
						j.atk(x, y+2, item.Vazio)
						j.NavesInimigo[item.PortaAviao]--

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
						j.NavesInimigo[item.PortaAviao]--

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
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
			if cDirE == nil && cDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x, y-2)
				j.atkVazioDiagonal(x+1, y-1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
		}

		if bE == nil && b.Tipo != item.Vazio {
			if bEsqE == nil && bEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x, y+2)
				j.atkVazioDiagonal(x-1, y+1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x, y+2)
				j.atkVazioDiagonal(x+1, y+1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
		}

		if dE == nil && d.Tipo != item.Vazio {
			if cDirE == nil && cDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x+2, y)
				j.atkVazioDiagonal(x+1, y-1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x+2, y)
				j.atkVazioDiagonal(x+1, y+1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
		}

		if eE == nil && e.Tipo != item.Vazio {
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x-2, y)
				j.atkVazioDiagonal(x-1, y-1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
			if bEsqE == nil && bEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x-2, y)
				j.atkVazioDiagonal(x-1, y+1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
		}

		if cDirE == nil && cDir.Tipo != item.Vazio {
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x+1, y-1)
				j.atkVazioDiagonal(x-1, y-1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x+1, y-1)
				j.atkVazioDiagonal(x+1, y+1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
		}

		if bEsqE == nil && bEsq.Tipo != item.Vazio {
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x-1, y+1)
				j.atkVazioDiagonal(x+1, y+1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y)
				j.atkVazioDiagonal(x-1, y+1)
				j.atkVazioDiagonal(x-1, y-1)
				j.NavesInimigo[item.Hidroaviao]--

				break
			}
		}

		break
	}

	return j.checaSeGanhou()
}

func (j *Jogo) RetornoDeReceberAtaque(x int, y int, tipo item.Nave) bool {
	j.atk(x, y, tipo, "sup")
	r := color.New(color.FgBlack).Add(color.BgRed).SprintfFunc()

	switch tipo {
	case item.Submarino:
		j.atkVazioDiagonal(x, y, "sup")
		j.atkVazioCruz(x, y, "sup")
		j.NossasNaves[item.Submarino]--
		fmt.Printf("\nDestruiram um %s.\n", r("Submarino"))
		break
	case item.Destroyer:
		j.atkVazioDiagonal(x, y, "sup")
		d, dE := j.CopiaCampo.GetItem(x+1, y)
		if dE == nil && d.Tipo != item.Vazio {
			j.atk(x-1, y, item.Vazio, "sup")
			j.atk(x+2, y, item.Vazio, "sup")
			j.NossasNaves[item.Destroyer]--
			fmt.Printf("\nDestruiram um %s.\n", r("Destroyer"))
			break
		}
		e, eE := j.CopiaCampo.GetItem(x-1, y)
		if eE == nil && e.Tipo != item.Vazio {
			j.atk(x+1, y, item.Vazio, "sup")
			j.atk(x-2, y, item.Vazio, "sup")
			j.NossasNaves[item.Destroyer]--
			fmt.Printf("\nDestruiram um %s.\n", r("Destroyer"))
			break
		}
		c, cE := j.CopiaCampo.GetItem(x, y+1)
		if cE == nil && c.Tipo != item.Vazio {
			j.atk(x, y-1, item.Vazio, "sup")
			j.atk(x, y+2, item.Vazio, "sup")
			j.NossasNaves[item.Destroyer]--
			fmt.Printf("\nDestruiram um %s.\n", r("Destroyer"))
			break
		}
		b, bE := j.CopiaCampo.GetItem(x, y-1)
		if bE == nil && b.Tipo != item.Vazio {
			j.atk(x, y+1, item.Vazio, "sup")
			j.atk(x, y-2, item.Vazio, "sup")
			j.NossasNaves[item.Destroyer]--
			fmt.Printf("\nDestruiram um %s.\n", r("Destroyer"))
			break
		}
		break
	case item.Cruzador:
		j.atkVazioDiagonal(x, y, "sup")
		// DIREITA
		d, dE := j.CopiaCampo.GetItem(x+1, y)
		d2, d2E := j.CopiaCampo.GetItem(x+2, y)
		d3, d3E := j.CopiaCampo.GetItem(x+3, y)
		// ESQUERDA
		e, eE := j.CopiaCampo.GetItem(x-1, y)
		e2, e2E := j.CopiaCampo.GetItem(x-2, y)
		e3, e3E := j.CopiaCampo.GetItem(x-3, y)

		if dE == nil && d.Tipo != item.Vazio {
			if d2E == nil && d2.Tipo != item.Vazio {
				if d3E == nil && d3.Tipo != item.Vazio {
					j.atk(x-1, y, item.Vazio)
					j.atk(x+4, y, item.Vazio)
					j.NossasNaves[item.Cruzador]--
					fmt.Printf("\nDestruiram um %s.\n", r("Cruzador"))
					break
				}
				if eE == nil && e.Tipo != item.Vazio {
					j.atk(x-2, y, item.Vazio)
					j.atk(x+3, y, item.Vazio)
					j.NossasNaves[item.Cruzador]--
					fmt.Printf("\nDestruiram um %s.\n", r("Cruzador"))
					break
				}
			}
			if eE == nil && e.Tipo != item.Vazio {
				if e2E == nil && e2.Tipo != item.Vazio {
					j.atk(x-3, y, item.Vazio)
					j.atk(x+2, y, item.Vazio)
					j.NossasNaves[item.Cruzador]--
					fmt.Printf("\nDestruiram um %s.\n", r("Cruzador"))
					break
				}
			}
		}

		if eE == nil && e.Tipo != item.Vazio {
			if e2E == nil && e2.Tipo != item.Vazio {
				if e3E == nil && e3.Tipo != item.Vazio {
					j.atk(x-4, y, item.Vazio)
					j.atk(x+1, y, item.Vazio)
					j.NossasNaves[item.Cruzador]--
					fmt.Printf("\nDestruiram um %s.\n", r("Cruzador"))
					break
				}
			}
		}

		// BAIXO
		b, bE := j.CopiaCampo.GetItem(x, y+1)
		b2, b2E := j.CopiaCampo.GetItem(x, y+2)
		b3, b3E := j.CopiaCampo.GetItem(x, y+3)
		// CIMA
		c, cE := j.CopiaCampo.GetItem(x, y-1)
		c2, c2E := j.CopiaCampo.GetItem(x, y-2)
		c3, c3E := j.CopiaCampo.GetItem(x, y-3)

		if bE == nil && b.Tipo != item.Vazio {
			if b2E == nil && b2.Tipo != item.Vazio {
				if b3E == nil && b3.Tipo != item.Vazio {
					j.atk(x, y-1, item.Vazio)
					j.atk(x, y+4, item.Vazio)
					j.NossasNaves[item.Cruzador]--
					fmt.Printf("\nDestruiram um %s.\n", r("Cruzador"))
					break
				}
				if cE == nil && c.Tipo != item.Vazio {
					j.atk(x, y-2, item.Vazio)
					j.atk(x, y+3, item.Vazio)
					j.NossasNaves[item.Cruzador]--
					fmt.Printf("\nDestruiram um %s.\n", r("Cruzador"))
					break
				}
			}
			if cE == nil && c.Tipo != item.Vazio {
				if c2E == nil && c2.Tipo != item.Vazio {
					j.atk(x, y-3, item.Vazio)
					j.atk(x, y+2, item.Vazio)
					j.NossasNaves[item.Cruzador]--
					fmt.Printf("\nDestruiram um %s.\n", r("Cruzador"))
					break
				}
			}
		}

		if cE == nil && c.Tipo != item.Vazio {
			if c2E == nil && c2.Tipo != item.Vazio {
				if c3E == nil && c3.Tipo != item.Vazio {
					j.atk(x, y-4, item.Vazio)
					j.atk(x, y+1, item.Vazio)
					j.NossasNaves[item.Cruzador]--
					fmt.Printf("\nDestruiram um %s.\n", r("Cruzador"))
					break
				}
			}
		}
		break
	case item.PortaAviao:
		j.atkVazioDiagonal(x, y, "sup")
		// DIREITA
		d, dE := j.CopiaCampo.GetItem(x+1, y)
		d2, d2E := j.CopiaCampo.GetItem(x+2, y)
		d3, d3E := j.CopiaCampo.GetItem(x+3, y)
		d4, d4E := j.CopiaCampo.GetItem(x+4, y)
		// ESQUERDA
		e, eE := j.CopiaCampo.GetItem(x-1, y)
		e2, e2E := j.CopiaCampo.GetItem(x-2, y)
		e3, e3E := j.CopiaCampo.GetItem(x-3, y)
		e4, e4E := j.CopiaCampo.GetItem(x-4, y)

		if dE == nil && d.Tipo != item.Vazio {
			if d2E == nil && d2.Tipo != item.Vazio {
				if d3E == nil && d3.Tipo != item.Vazio {
					if d4E == nil && d4.Tipo != item.Vazio {
						j.atk(x-1, y, item.Vazio)
						j.atk(x+5, y, item.Vazio)
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
						break
					}
					if eE == nil && e.Tipo != item.Vazio {
						j.atk(x-2, y, item.Vazio)
						j.atk(x+4, y, item.Vazio)
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
						break
					}
				}
				if eE == nil && e.Tipo != item.Vazio {
					if e2E == nil && e2.Tipo != item.Vazio {
						j.atk(x-3, y, item.Vazio)
						j.atk(x+3, y, item.Vazio)
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
						break
					}
				}
			}
			if eE == nil && e.Tipo != item.Vazio {
				if e2E == nil && e2.Tipo != item.Vazio {
					if e3E == nil && e3.Tipo != item.Vazio {
						j.atk(x-4, y, item.Vazio)
						j.atk(x+2, y, item.Vazio)
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
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
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
						break
					}
				}
			}
		}

		// BAIXO
		b, bE := j.CopiaCampo.GetItem(x, y+1)
		b2, b2E := j.CopiaCampo.GetItem(x, y+2)
		b3, b3E := j.CopiaCampo.GetItem(x, y+3)
		b4, b4E := j.CopiaCampo.GetItem(x, y+4)
		// CIMA
		c, cE := j.CopiaCampo.GetItem(x, y-1)
		c2, c2E := j.CopiaCampo.GetItem(x, y-2)
		c3, c3E := j.CopiaCampo.GetItem(x, y-3)
		c4, c4E := j.CopiaCampo.GetItem(x, y-4)

		if bE == nil && b.Tipo != item.Vazio {
			if b2E == nil && b2.Tipo != item.Vazio {
				if b3E == nil && b3.Tipo != item.Vazio {
					if b4E == nil && b4.Tipo != item.Vazio {
						j.atk(x, y-1, item.Vazio)
						j.atk(x, y+5, item.Vazio)
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
						break
					}
					if cE == nil && c.Tipo != item.Vazio {
						j.atk(x, y-2, item.Vazio)
						j.atk(x, y+4, item.Vazio)
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
						break
					}
				}
				if cE == nil && c.Tipo != item.Vazio {
					if c2E == nil && c2.Tipo != item.Vazio {
						j.atk(x, y-3, item.Vazio)
						j.atk(x, y+3, item.Vazio)
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
						break
					}
				}
			}
			if cE == nil && c.Tipo != item.Vazio {
				if c2E == nil && c2.Tipo != item.Vazio {
					if c2E == nil && c2.Tipo != item.Vazio {
						j.atk(x, y-4, item.Vazio)
						j.atk(x, y+2, item.Vazio)
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
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
						j.NossasNaves[item.PortaAviao]--
						fmt.Printf("\nDestruiram um %s.\n", r("PortaAviao"))
						break
					}
				}
			}
		}

		break
	case item.Hidroaviao:
		j.atkVazioCruz(x, y)
		// CIMA
		c, cE := j.CopiaCampo.GetItem(x, y-2)
		cEsq, cEsqE := j.CopiaCampo.GetItem(x-1, y-1)
		cDir, cDirE := j.CopiaCampo.GetItem(x+1, y-1)
		// BAIXO
		b, bE := j.CopiaCampo.GetItem(x, y+2)
		bEsq, bEsqE := j.CopiaCampo.GetItem(x-1, y+1)
		bDir, bDirE := j.CopiaCampo.GetItem(x+1, y+1)
		// LATERAL
		d, dE := j.CopiaCampo.GetItem(x+2, y)
		e, eE := j.CopiaCampo.GetItem(x-2, y)

		if cE == nil && c.Tipo != item.Vazio {
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x, y-2, "sup")
				j.atkVazioDiagonal(x-1, y-1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
			if cDirE == nil && cDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x, y-2, "sup")
				j.atkVazioDiagonal(x+1, y-1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
		}

		if bE == nil && b.Tipo != item.Vazio {
			if bEsqE == nil && bEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x, y+2, "sup")
				j.atkVazioDiagonal(x-1, y+1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x, y+2, "sup")
				j.atkVazioDiagonal(x+1, y+1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
		}

		if dE == nil && d.Tipo != item.Vazio {
			if cDirE == nil && cDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x+2, y, "sup")
				j.atkVazioDiagonal(x+1, y-1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x+2, y, "sup")
				j.atkVazioDiagonal(x+1, y+1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
		}

		if eE == nil && e.Tipo != item.Vazio {
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x-2, y, "sup")
				j.atkVazioDiagonal(x-1, y-1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
			if bEsqE == nil && bEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x-2, y, "sup")
				j.atkVazioDiagonal(x-1, y+1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
		}

		if cDirE == nil && cDir.Tipo != item.Vazio {
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x+1, y-1, "sup")
				j.atkVazioDiagonal(x-1, y-1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x+1, y-1, "sup")
				j.atkVazioDiagonal(x+1, y+1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
		}

		if bEsqE == nil && bEsq.Tipo != item.Vazio {
			if bDirE == nil && bDir.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x-1, y+1, "sup")
				j.atkVazioDiagonal(x+1, y+1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
			if cEsqE == nil && cEsq.Tipo != item.Vazio {
				j.atkVazioDiagonal(x, y, "sup")
				j.atkVazioDiagonal(x-1, y+1, "sup")
				j.atkVazioDiagonal(x-1, y-1, "sup")
				j.NossasNaves[item.Hidroaviao]--
				fmt.Printf("\nDestruiram um %s.\n", r("Hidroaviao"))
				break
			}
		}

		break
	}

	return j.checaSePerdeu()
}
