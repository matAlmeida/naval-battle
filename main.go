package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"github.com/matalmeida/shipbattle/field"
	"github.com/matalmeida/shipbattle/game"
	"github.com/matalmeida/shipbattle/item"
)

const PROMPT = ">> "
const TENTATIVAS = 3
const GAME_SIZE = 10

type Jogo struct {
	Campo        *field.Campo
	CampoInimigo *field.Campo
}

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", 1)

	return userInput
}

func main() {
	g := game.Novo()
	g = nossoCampo(g)
	v := color.New(color.FgBlack).Add(color.BgGreen).SprintfFunc()
	r := color.New(color.FgBlack).Add(color.BgRed).SprintfFunc()

	menu := v(" 0 ") + " Sair\n" +
		v(" 1 ") + " Atacar\n" +
		v(" 2 ") + " Receber Ataque\n"

	atkMenu := r(" 0 ") + " Água\n" +
		r(" 1 ") + " Submarino\n" +
		r(" 2 ") + " Destroyer\n" +
		r(" 3 ") + " Hidro-Avião\n" +
		r(" 4 ") + " Cruzador\n" +
		r(" 5 ") + " Porta-Avião\n"

	fmt.Printf("%s\n", v("Bem vindo a Batalha Naval!"))

	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		fmt.Printf("Nosso Campo\n")
		fmt.Printf("%s\n", g.Campo.String())
		fmt.Printf("Campo Inimigo\n")
		fmt.Printf("%s\n", g.CampoInimigo.String())
		fmt.Printf("\n%s", menu)
		fmt.Printf("\n%s", PROMPT)
		userInput := readString()

		switch userInput {
		case "0":
			fmt.Printf("\nAté a próxima!\n")
			os.Exit(0)
			break
		case "1":
			x, y := g.Atacar()

			fmt.Printf("Atacando (%c, %d)\n", 'a'+x, y+1)

			fmt.Printf("\nQual foi o resultado do ataque? \n")
			fmt.Printf("%s\n", atkMenu)
			fmt.Printf(PROMPT)

			atkResponse := readString()
			tipoNave, _ := strconv.ParseInt(atkResponse, 10, 32)
			ganhou := g.RetornoDeAtaque(x, y, item.Nave(tipoNave))
			if ganhou {
				fmt.Printf("%s\n", v("Ganhei o Jogo!"))
				os.Exit(0)
			}
			break
		case "2":
			fmt.Printf("Digite as coordenadas([a-z],[1-10]): ")
			atkResponse := readString()
			xS := strings.Split(atkResponse, ",")[0]
			x := xS[0] - 'a'
			yS := strings.Split(atkResponse, ",")[1]
			y, _ := strconv.ParseInt(yS, 10, 32)
			naveAtacada := g.SerAtacado(int(x), int(y-1))

			fmt.Printf("Atingiu %s.\n", naveAtacada.String())
			break
		}
	}
}

func nossoCampo(j *game.Jogo) *game.Jogo {
	// PORTA AVIÃO
	j.Campo.ColocaItem(1, 9, "p1", item.PortaAviao, field.Baixo)
	// SUBMARINO
	j.Campo.ColocaItem(3, 0, "s1", item.Submarino, field.Baixo)
	j.Campo.ColocaItem(7, 7, "s2", item.Submarino, field.Baixo)
	j.Campo.ColocaItem(7, 9, "s3", item.Submarino, field.Baixo)
	j.Campo.ColocaItem(9, 9, "s4", item.Submarino, field.Baixo)
	// DESTROYER
	j.Campo.ColocaItem(3, 2, "d1", item.Destroyer, field.Baixo)
	j.Campo.ColocaItem(3, 4, "d2", item.Destroyer, field.Baixo)
	j.Campo.ColocaItem(9, 6, "d3", item.Destroyer, field.Direita)
	// CRUZADOR
	j.Campo.ColocaItem(5, 0, "c1", item.Cruzador, field.Baixo)
	j.Campo.ColocaItem(6, 2, "c2", item.Cruzador, field.Direita)
	// HIDRO AVIÃO
	j.Campo.ColocaItem(0, 2, "h1", item.Hidroaviao, field.Baixo)
	j.Campo.ColocaItem(0, 6, "h2", item.Hidroaviao, field.Baixo)
	j.Campo.ColocaItem(4, 6, "h3", item.Hidroaviao, field.Direita)
	j.Campo.ColocaItem(8, 3, "h4", item.Hidroaviao, field.Baixo)

	return j
}
