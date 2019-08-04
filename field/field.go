package field

import (
	"errors"

	"github.com/matalmeida/shipbattle/item"
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
			ret += c.itens[i][j].String() + "\t"
		}
		ret += "\n"
	}

	return ret
}

func (c *Campo) GetItem(x int, y int) (*item.Item, error) {
	if (x >= 0 && x < c.Tamanho) && (y >= 0 && y < c.Tamanho) {
		return c.itens[x][y], nil
	}

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

func (c *Campo) checaAdjacentes(x int, y int, i *item.Item) (bool, error) {
	_, err := c.GetItem(x, y)
	if err != nil {
		return false, errors.New("item não existe")
	}

	upleft, _ := c.GetItem(x-1, y-1)
	up, _ := c.GetItem(x, y-1)
	upright, _ := c.GetItem(x+1, y-1)
	left, _ := c.GetItem(x-1, y)
	right, _ := c.GetItem(x+1, y)
	downleft, _ := c.GetItem(x-1, y+1)
	down, _ := c.GetItem(x, y+1)
	downright, _ := c.GetItem(x+1, y+1)

	if upleft.Tipo == item.Vazio && up.Tipo == item.Vazio && upright.Tipo == item.Vazio && left.Tipo == item.Vazio && right.Tipo == item.Vazio && downleft.Tipo == item.Vazio && down.Tipo == item.Vazio && downright.Tipo == item.Vazio {
		return true, nil
	} else {
		if upleft.Id != i.Id && up.Id != i.Id && upright.Id != i.Id && left.Id != i.Id && right.Id != i.Id && downleft.Id != i.Id && down.Id != i.Id && downright.Id != i.Id {
			return false, errors.New("existe alguma nave adjacente")
		}

		return true, nil
	}
}

func (c *Campo) ColocaItem(x int, y int, id string, tipo item.Nave) bool {
	i := item.Novo(id, tipo)
	_, err := c.checaAdjacentes(x, y, i)
	if err != nil {
		return false
	}

	c.itens[x][y] = i
	return true
}
