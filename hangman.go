package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Inicializar el juego
	// Inicializar un jugador
	// Generar una palabra
	// Si el juego esta activo
	// Imput del jugador
	// Comprobar si la palabra es correcta o no
	// Si el jugador tiene vidas
	// Si es su turno

	displayMainMenu()
}

func displayMainMenu() {

	fmt.Println(`
Main menu:
(N) - New Game
(P) - Two Player Mode
(S) - Settings
(Q) - Quit
		`)

	isValidOptionMenu := false
	for !isValidOptionMenu {

		menuReader := bufio.NewReader(os.Stdin)
		fmt.Print("Select an option from the menu above: ")
		option, _ := menuReader.ReadString('\n')
		option = strings.Replace(option, "\n", "", -1)

		switch option {
		case "N":
			fmt.Println("You've selected start a new game.")
			isValidOptionMenu = true
		case "P":
			fmt.Println("You've selected a two player mode game.")
			isValidOptionMenu = true
		case "S":
			fmt.Println("You've selected change settings.")
			isValidOptionMenu = true
		case "Q":
			fmt.Println("Bye.")
			isValidOptionMenu = true
		default:
			fmt.Println("Sorry unknown option, please try again.")
		}
	}
}
