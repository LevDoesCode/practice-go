package main

import "fmt"

type vertex struct {
	X int
	Y int
}

type player struct {
	hp       int
	level    int
	location struct {
		x int
		y int
		z int
	}
}

func main() {
	fmt.Println(vertex{1, 2})
	var v2 vertex = vertex{2, 3}
	v3 := vertex{3, 4}
	fmt.Println(v2)
	fmt.Println(v3)
	var player1 player = player{
		hp:    300,
		level: 1,
		location: struct {
			x int
			y int
			z int
		}{6, 7, 9},
	}
	player1.hp = 150
	player1.location.x = 10
	fmt.Println(player1)
	fmt.Println(player1.location)
	fmt.Printf("x is %v, y is %v, z is %v\n", player1.location.x, player1.location.y, player1.location.z)

	var player2 player // dot notation can be used to initialize values inside the struct
	player2.hp = 40
	player2.level = 20
	player2.location.x = 11
	player2.location.y = 22
	player2.location.z = 33
	fmt.Println(player2)
}
