package main

import "github.com/manh737/bitking-return-golang/router"

func main() {
	r := router.New()

	v1 := r.Group("/api")

	h := handler.NewHandler(us, as)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
