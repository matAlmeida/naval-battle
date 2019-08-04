package field

import (
	"testing"

	"github.com/matalmeida/shipbattle/item"
)

func TestCampo(t *testing.T) {
	campo := Novo(10)
	t.Logf("Cria campo 10x10")

	t.Logf("Coloca item0 em (0, 0)")
	colocou := campo.ColocaItem(0, 0, "item0", item.PortaAviao)
	if !colocou {
		t.Fatalf("Expected True, got False.")
	}

	t.Logf("Coloca item0 em (0, 1)")
	colocou = campo.ColocaItem(0, 1, "item0", item.PortaAviao)
	if !colocou {
		t.Fatalf("Expected True, got False")
	}

	t.Logf("Não consegue colocar item1 em (0, 2)")
	colocou = campo.ColocaItem(0, 2, "item1", item.PortaAviao)
	if colocou {
		t.Fatalf("Expected False, got True")
	}
}
