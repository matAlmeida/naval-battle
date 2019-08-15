package field

import (
	"errors"

	"github.com/matalmeida/shipbattle/item"
)

type Direcao int32

const (
	Cima     Direcao = 0
	Direita  Direcao = 1
	Baixo    Direcao = 2
	Esquerda Direcao = 3
)

type Campo struct {
	Tamanho int
	itens   [][]*item.Item
}

func Novo(tamanho int) *Campo {
	c := &Campo{Tamanho: tamanho, itens: make([][]*item.Item, tamanho)}

	for x := range c.itens {
		c.itens[x] = make([]*item.Item, tamanho)

		for y := range c.itens[x] {
			i := item.Novo("nil", item.Vazio)
			c.itens[x][y] = i
		}
	}

	return c
}

func (c *Campo) String() string {
	ret := ""

	for i := 0; i < c.Tamanho; i++ {
		for j := 0; j < c.Tamanho; j++ {
			ret += c.itens[i][j].String() + " "
		}
		ret += "\n"
	}

	return ret
}

func (c *Campo) GetItem(x int, y int) (*item.Item, error) {
	if (x >= 0 && x < c.Tamanho) && (y >= 0 && y < c.Tamanho) {
		return c.itens[x][y], nil
	}
	// _, file, no, ok := runtime.Caller(1)
	// if ok {
	// 	fmt.Printf("called from %s#%d\n", file, no)
	// }

	// fmt.Printf("\nCoord fora da matriz\n")
	return item.Novo("nil", item.Vazio), errors.New("coordenadas fora da matriz")
}

func (c *Campo) LimpaPosicao(x int, y int) {
	_, err := c.GetItem(x, y)
	if err != nil {
		return
	}

	c.itens[x][y] = item.Novo("nil", item.Vazio)
	return
}

func (c *Campo) Atacar(x int, y int) (bool, *item.Item) {
	i, err := c.GetItem(x, y)
	if err != nil {
		return false, nil
	}

	if i.Bombardeado {
		return false, i
	}

	i.Afundar()
	if i.Tipo == item.Vazio {
		return false, i
	} else {
		return true, i
	}

}

func (c *Campo) ColocaItem(x int, y int, id string, tipo item.Nave, direcao ...Direcao) bool {
	if id == "enemy" {
		_, eItemErro := c.GetItem(x, y)
		if eItemErro == nil {
			c.itens[x][y] = item.Novo(id, tipo)
		}
		return true
	}

	dir := Direita

	if len(direcao) > 0 {
		dir = direcao[0]
	}

	if tipo == item.Hidroaviao {
		c.itens[x][y] = item.Novo(id, tipo)
		switch dir {
		case Cima:
			c.itens[x-1][y-1] = item.Novo(id, tipo)
			c.itens[x-1][y+1] = item.Novo(id, tipo)
			break
		case Baixo:
			c.itens[x+1][y-1] = item.Novo(id, tipo)
			c.itens[x+1][y+1] = item.Novo(id, tipo)
			break
		case Esquerda:
			c.itens[x-1][y-1] = item.Novo(id, tipo)
			c.itens[x+1][y-1] = item.Novo(id, tipo)
			break
		default:
			c.itens[x-1][y+1] = item.Novo(id, tipo)
			c.itens[x+1][y+1] = item.Novo(id, tipo)
			break
		}
	} else if tipo != item.Vazio {
		for j := 0; j < int(tipo); j++ {
			switch dir {
			case Cima:
				c.itens[x-j][y] = item.Novo(id, tipo)
				break
			case Baixo:
				c.itens[x+j][y] = item.Novo(id, tipo)
				break
			case Esquerda:
				c.itens[x][y-j] = item.Novo(id, tipo)
				break
			default:
				c.itens[x][y+j] = item.Novo(id, tipo)
				break
			}
		}
	}

	return true
}
