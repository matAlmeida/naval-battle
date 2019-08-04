package field

import (
	"testing"

	"github.com/matalmeida/shipbattle/item"
)

func TestCampo(t *testing.T) {
	campo := Novo(10)
	t.Logf("Cria campo 10x10")
	t.Logf("\n%s", campo.String())

	t.Logf("Coloca item0 em (0, 0)")
	colocou := campo.ColocaItem(0, 0, "item0", item.PortaAviao)
	if !colocou {
		t.Logf("\n%s", campo.String())
		t.Fatalf("Expected True, got False.")
	}

	t.Logf("Coloca item0 em (0, 1)")
	colocou = campo.ColocaItem(0, 1, "item0", item.PortaAviao)
	if colocou {
		t.Logf("\n%s", campo.String())
		t.Fatalf("Expected False, got True")
	}

	t.Logf("Não consegue colocar item1 em (0, 2)")
	colocou = campo.ColocaItem(0, 2, "item1", item.PortaAviao)
	if colocou {
		t.Logf("\n%s", campo.String())
		t.Fatalf("Expected False, got True")
	}

	t.Logf("Coloca item1 em (1, 0) para baixo")
	colocou = campo.ColocaItem(0, 0, "item0", item.PortaAviao, Baixo)
	if colocou {
		t.Logf("\n%s", campo.String())
		t.Fatalf("Expected True, got False.")
	}

}
