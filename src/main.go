package main

import (
	"github.com/gashjp1994/go-ca/db/mysql"
	"github.com/gashjp1994/go-ca/router"
)

func main() {
	defer mysql.CloseConnect()
	router.Setup()
}
