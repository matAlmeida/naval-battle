package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

const PROMPT = ">> "
const TENTATIVAS = 3

func main() {
	// g := game.Novo()
	reader := bufio.NewReader(os.Stdin)
	p := color.New(color.FgBlack).Add(color.BgGreen).SprintfFunc()
	menu := p(" 0 ") + " Sair\n" +
		p(" 1 ") + " Atacar\n" +
		p(" 2 ") + " Receber Ataque\n"

	fmt.Printf("%s\n\n", p("Bem vindo a Batalha Naval!"))

	for {
		fmt.Printf(menu)
		fmt.Printf(PROMPT)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", 1)

		switch userInput {
		case "0":
			fmt.Printf("\nAté a próxima!\n")
			os.Exit(0)
			break
		default:
			break
		}

		fmt.Printf(userInput)
	}
}
