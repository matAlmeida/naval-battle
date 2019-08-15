package item

type Nave int32

const (
	Vazio Nave = 0

	Submarino  Nave = 1
	Destroyer  Nave = 2
	Cruzador   Nave = 4
	PortaAviao Nave = 5

	Hidroaviao Nave = 3
)

type Item struct {
	Id          string
	Bombardeado bool
	Tipo        Nave
}

func Novo(id string, tipo Nave) *Item {
	i := &Item{Id: id, Tipo: tipo, Bombardeado: false}

	return i
}

func (i *Item) Afundar() {
	if !i.Bombardeado {
		i.Bombardeado = true
	}
}

func (i *Item) String() string {
	switch i.Tipo {
	case Vazio:
		if i.Bombardeado {
			return "A"
		}
		return "B"
	case PortaAviao:
		return "P"
	case Cruzador:
		return "C"
	case Destroyer:
		return "D"
	case Submarino:
		return "S"
	case Hidroaviao:
		return "H"
	default:
		return "unk"
	}
}
