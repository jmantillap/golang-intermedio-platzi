package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var x int
	x = 8
	y := 1

	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("7ss", 0, 64)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	} else {
		fmt.Println(myValue)
	}
	// Mapa: estructura de clave valor, con Make especificar un map que mapee llaves de tipo string a valores de tipo entero
	m := make(map[string]int)
	//Key= string, 6= int
	m["key"] = 6
	fmt.Println(m["key"])

	// Slice: Estructura como un array
	s := []int{1, 2, 3}
	// Con el for recorremos el slice
	for index, value := range s {
		//index: valor en memoria que estamos accediendo
		fmt.Println(index)
		//value: valor almacenado en el slice
		fmt.Println(value)
	}
	// Agregar un valor nuevo al final del slice
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25
	fmt.Println(g) // imprime el valor entero 25
	h := &g
	fmt.Println(h) // imprimer la direccion de memoria.
	i := *h
	fmt.Println(i) // Imprime el valor por de g
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	c <- 1
}
