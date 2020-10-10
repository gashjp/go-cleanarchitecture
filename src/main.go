package main

import (
	"github.com/gashjp1994/go-ca/db/mysql"
	"github.com/gashjp1994/go-ca/handler"
)

func main() {
	defer mysql.CloseConnect()
	handler.Setup()
}
