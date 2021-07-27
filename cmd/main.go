package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	"goNorthwindApiDemo/cfg"
	"goNorthwindApiDemo/routing"
	"net/http"
)

func main()  {
	cfg.DbExport = cfg.DB{
		Server:   "127.0.0.1",
		Port:     1433,
		User:     "sa",
		Password: "Password12*",
		Database: "Northwind",
	}
	router := mux.NewRouter()
	routing.Router(router)

	err := http.ListenAndServe(":52",
			handlers.CORS(handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		fmt.Println(err.Error())
	}
}