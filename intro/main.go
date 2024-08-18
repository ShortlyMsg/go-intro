package main

import (
	"fmt"
)

// Gin, Middleware, JWT, Redis, PostgreSQL, Ent Alternatifi bul
type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "test2", Age: 2}
	person2 := Person{"deneme3", 2}

	fmt.Println(person)
	fmt.Println(person2)
}
