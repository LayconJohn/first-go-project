package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func sum(x int, y int) (int, bool) {
	if x > 0 {
		return x + y, true
	}
	return x + y, false
}

func testingVariables() {
	//var name string
	name := "Teste"
	name = "Testee"
	result, status := sum(4, 9)

	println(result)
	println(status)
	println(name)

	fmt.Println("Testando pacotes: ", name, result, status)
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func home(res http.ResponseWriter, req *http.Request) {

	course := Course{
		Name:        "Como ficar rico e comprar uma ilha",
		Description: "Fique rico com jogo do tigrinho",
		Price:       100000,
	}
	course.Price = 150000
	fmt.Println(course.GetFullInfo())
	json.NewEncoder(res).Encode(course)
}

// Letra maiúscula: public. Minúscula: private
type Course struct {
	Name        string `json:"course"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

// T1
func main() {
	channel := make(chan int)
	for i := 0; i < 3; i++ {
		go worker(i, channel) //T2, T3, T4
	}

	for i := 0; i < 30; i++ {
		channel <- i
	}

	/*
		go counter() //thread 1 -> Rodar em paralelo (go rountine)
		go counter() //thread 2
		counter()
		testingVariables()

		http.HandleFunc("/", home)
		http.ListenAndServe(":5000", nil)*/
}
