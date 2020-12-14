package main

import (
	"github.com/manh737/bitking-return-golang/db"
	"github.com/manh737/bitking-return-golang/router"
)

func main() {

	db, err := db.New("mongodb://localhost:27017", "bitkingreturns")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := router.Create(db)

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
