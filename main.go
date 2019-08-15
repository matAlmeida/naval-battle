package main

import (
	"bufio"
	"fmt"
	"os"
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
			fmt.Printf("Atacando (%d, %d)\n", x, y)
			fmt.Printf("\nQual foi o resultado do ataque? \n")
			fmt.Printf("%s\n", atkMenu)
			fmt.Printf(PROMPT)
			atkResponse := readString()
			tipoNave, _ := strconv.ParseInt(atkResponse, 10, 32)
			ganhou := g.RetornoDeAtaque(1, 1, item.Nave(tipoNave))
			if ganhou {
				fmt.Printf("%s\n", v("Ganhei o Jogo!"))
				os.Exit(0)
			} else {
				fmt.Printf("%s\n", r("Mais alguns ataques para ganhar!"))
			}
			break
		case "2":
			fmt.Printf("Digite as coordenadas: ")
			atkResponse := readString()
			xS := strings.Split(atkResponse, ",")[0]
			x, _ := strconv.ParseInt(xS, 10, 32)
			yS := strings.Split(atkResponse, ",")[1]
			y, _ := strconv.ParseInt(yS, 10, 32)
			naveAtacada := g.SerAtacado(int(x), int(y))

			fmt.Printf("Atingiu %s.", naveAtacada.String())
			break
		default:
			break
		}
	}
}
