package main

import (
	"fmt"
	"os"
	"strconv"
)

var field [][]int //créé un tableau à 2 dimensions

func draw() {
	for _, row := range field {
		fmt.Println(row) //écrit chaque ligne du tableau
	}
	fmt.Println()
}

func canPut(x int, y int, value int) bool { //appelle les 3 fonctions de vérifications pour chaque position x y et chaque valeur
	return !alreadyInVertical(x, y, value) &&
		!alreadyInHorizontal(x, y, value) &&
		!alreadyInSquare(x, y, value)
}

func alreadyInVertical(x int, y int, value int) bool { //vérifie si value est déjà écrit dans la ligne horizontal du tableau field
	for i := range [9]int{} {
		if field[i][x] == value {
			return true
		}
	}
	return false
}

func alreadyInHorizontal(x int, y int, value int) bool { //vérifie si value est déjà écrit dans la ligne vertical du tableau field
	for i := range [9]int{} {
		if field[y][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(x int, y int, value int) bool { //vérifie si value est déjà écrit le carré 3*3 du tableau field
	sx, sy := int(x/3)*3, int(y/3)*3 // sx = (x/3)*3 && sy = (y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if field[sy+dy][sx+dx] == value {
				return true //renvoie true si la valeur est déjà présente
			}
		}
	}
	return false // sinon renvoie false
}

func next(x int, y int) (int, int) { //avance de 1 dans le tableau
	nextX, nextY := (x+1)%9, y // nextX = (x+1)%9 et nextY = y
	if nextX == 0 {            // si nextX == 0 nextY = y+1
		nextY = y + 1
	}
	return nextX, nextY //renvoi nextX et nextY
}

func solve(x int, y int) bool { // résout le sudoku
	if y == 9 { // si on a fini d'append le tableau return true
		return true
	}
	if field[y][x] != 0 { // si la valeur de field[y][x] != 0 on lance une recursive
		return solve(next(x, y))
	} else { // sinon
		for i := range [9]int{} {
			var v = i + 1
			if canPut(x, y, v) { // verifie si la case peut être écrit avec la valeur de v
				field[y][x] = v        // ajoute la valeur de v au tableau field
				if solve(next(x, y)) { // si solve next est vrai (vérifie si la place est libre) return true
					return true
				}
				field[y][x] = 0 // sinon field = 0
			}
		}
		return false //si ce n'est pas faisable return false
	}
}

func min() { // initialise les variables et appelle la fonction recursive solve et vérifie la légalité du os.Args[1:]
	args := os.Args[1:]

	if len(args) != 9 { //si il y a plus de 9 paramètres écrit Error et return
		fmt.Println("Error")
		return
	}

	for a := range args { // si on vérifie la taille de chaque string et on return error si la taille n'est pas bonne
		if len(args[a]) < 9 || len(args[a]) > 9 {
			fmt.Println("Error")
			fmt.Println()
			return
		}
	}

	for a := 0; a < 8; a++ {
		for b := a + 1; b < 9; b++ {
			for c := 0; c < 9; c++ {
				if args[a][c] == args[b][c] && args[a][c] != '.' && args[b][c] != '.' { // affiche une erreur d'argument si deux chiffres donnés sont
					fmt.Println("Erreur") //identiques dans la même colonne (corrige une erreur de la fonction alreadyInVertical qui vérifie mal la dernière ligne)
					fmt.Println()
					return
				}
			}
		}
	}

	compteur := 0
	for a := range args {
		for _, b := range args[a] {
			if b != '.' {
				compteur++
			}
		}
	}
	if compteur < 17 {
		fmt.Println("Error")
		fmt.Println()
		return
	}

	field = make([][]int, 9) // ajoute 9 tableaux vide au premier tableau de field
	for i, arg := range args {
		row := make([]int, 9) // créé un tableau row de 9 entrer
		for j, char := range arg {
			if char == '.' {
				row[j] = 0 //remplace tous les '.' par des 0
			} else {
				num, err := strconv.Atoi(string(char)) //convertie char de string en int et si erreur écrit "Erreur" et return
				if err != nil || num < 1 || num > 9 {
					fmt.Println("Error")
					fmt.Println()
					return
				}
				row[j] = num //ajoute num au tableau row
			}
		}
		field[i] = row // ajoute row dans le tableau field
	}

	if solve(0, 0) { //si solve réussi dessine sinon print "Error"
		draw()
	} else {
		fmt.Println("Error")
		fmt.Println()
	}
}
