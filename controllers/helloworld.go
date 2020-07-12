package controllers

import (
	"net/http"
)

type Hello interface {
	SayHello() http.HandlerFunc
}

type name struct {
	name string
}

func NewName() *name{
	return &name{name: "hello1"}
}

//SayHello
func (n *name) SayHello() http.HandlerFunc{

	return func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Println("Say hello", n.name)
		writer.Write([]byte(n.name))
	}

	//return nil
}