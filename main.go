package main

import (
	"github.com/kerlexov/back_core/db"
	"github.com/kerlexov/back_core/handler"
	"github.com/kerlexov/back_core/router"
	"github.com/kerlexov/back_core/store"
	"net/http"
	"time"

)

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	d.Close()

	h := handler.NewHandler(us)
	h.Register(v1)

	s := &http.Server{
		Addr:         ":420",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	r.Logger.Fatal(r.StartServer(s))
}
