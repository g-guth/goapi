package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	viper.AddConfigPath("../app/config")
	viper.AddConfigPath("./config")
	viper.SetConfigName("environment")
	viper.ReadInConfig()
	fmt.Fprintf(w, "Welcome home! env: %s", viper.Get("env"))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
