package field

import (
	"testing"

	"github.com/matalmeida/shipbattle/item"
)

func TestCampo(t *testing.T) {
	campo := Novo(10)
	t.Logf("Cria campo 10x10")

	t.Logf("Coloca PortaAvião0 em (0, 0)")
	colocou := campo.ColocaItem(0, 0, "PortaAvião0", item.PortaAviao)
	if !colocou {
		t.Logf("\n%s", campo.String())
		t.Fatalf("Expected True, got False.")
	}

	t.Logf("Coloca Hidroavião0 em (2, 1) para baixo")
	colocou = campo.ColocaItem(2, 1, "Hidroavião0", item.Hidroaviao, Baixo)
	if !colocou {
		t.Logf("\n%s", campo.String())
		t.Fatalf("Expected True, got False.")
	}

	t.Logf("\n%s", campo.String())

}
