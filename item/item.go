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

func (i *Item) String() string {
	switch i.Tipo {
	case Vazio:
		return "s0"
	case PortaAviao:
		return "s1"
	case Cruzador:
		return "s2"
	case Destroyer:
		return "s3"
	case Submarino:
		return "s4"
	case Hidroaviao:
		return "s5"
	default:
		return "unk"
	}
}
