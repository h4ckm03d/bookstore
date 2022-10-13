package handlers

import "fmt"

func GetHeaderBlabla() {
	b := B{
		Name: "Lutfi",
		A: A{
			ID: 1,
		},
	}
	fmt.Println(b.A.ID, b.ID)
}

type A struct {
	ID int
}

type B struct {
	A
	Name string
}
