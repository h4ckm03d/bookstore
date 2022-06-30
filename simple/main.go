package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovering from panic:", err)
		}
	}()

	r := Rectangle{
		Width:  10,
		Height: 10,
		AreaFn: func() float64 {
			return 10 * 19
		},
	}
	var r2 *Rectangle

	fmt.Println(r2.Width)
	r.MutatedWidth(4)
	fmt.Println(r.Width)
	update(&r)
	fmt.Println(r.Width)
	PrintService(&r)

	fmt.Println(r.AreaFn())
	fmt.Println([]interface{}{1, "satu", 3})
	sampleMap := map[string]any{"satu": 1, "dua": true, "tiga": "333333"}
	keys := []string{"satu", "dua"}

	for _, v := range []interface{}{1, "satu", 3} {
		fmt.Println(v)
	}

	for _, key := range keys {
		fmt.Println(sampleMap[key])
	}

	for key, value := range sampleMap {
		fmt.Println(key, value)
	}

	if v, ok := sampleMap["tiga"]; ok {
		fmt.Println("tiga is in the map", v)
	}

	delete(sampleMap, "dua")

	fmt.Println(sampleMap)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = append(slice[:2], slice[3:4]...)
	fmt.Println(slice)

}

func getSomething() *Rectangle {
	return &Rectangle{}
}

func update(r *Rectangle) {
	r.Width = 10
}

func update2(r Rectangle) {
	r.Width = 10
}

// abstract class
// inherit

// contract - interface

// struct = class

type AreaCalc interface {
	Area() float64
}

type ServiceProvider interface {
	GetService() string
}

type AreaProvider interface {
	ServiceProvider
	AreaCalc
}

var _ AreaProvider = &Rectangle{}

type Rectangle struct {
	Width  float64        `json:"width,omitempty"`
	Height float64        `json:"height,omitempty"`
	AreaFn func() float64 `json:"area_fn,omitempty"`
}

func (r *Rectangle) MutatedWidth(w float64) {
	r.Width = w
}

func (r *Rectangle) GetService() string {
	return "you got a service"
}

func (r Rectangle) NonMutated(w float64) {
	r.Width = w
}

func (r Rectangle) Area() float64 {
	return 0
}

func (r Rectangle) Copy() (*Rectangle, error) {
	return nil, nil
}

func PrintArea(r Rectangle) {
	fmt.Println(r.Area())
}

func PrintService(s ServiceProvider) {
	fmt.Println(s.GetService())
}
