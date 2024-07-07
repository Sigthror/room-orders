package handlers

import (
	"fmt"
	"net/http"
)

type Order struct{}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented") // TODO: Implement
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented") // TODO: Implement
}

func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented") // TODO: Implement
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented") // TODO: Implement
}
