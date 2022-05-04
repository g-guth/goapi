package main

import (
	"fmt"
	"io/ioutil"
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

	fmt.Fprintf(w, "Welcome home! With env: %s", viper.Get("env"))
}

func request(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://goapi02:8081/")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Fprintf(w, sb)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/app", request)
	log.Fatal(http.ListenAndServe(":8080", router))
}
