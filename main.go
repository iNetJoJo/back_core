package main

import (
	"back_core/db"
	"back_core/handler"
	"back_core/router"
	"back_core/store"
	"net/http"
	"time"

)

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	//as := store.NewArticleStore(d)
	h := handler.NewHandler(us)
	h.Register(v1)

	s := &http.Server{
		Addr:         ":420",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	r.Logger.Fatal(r.StartServer(s))
}
