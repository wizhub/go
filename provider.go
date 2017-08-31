package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type ShippingPrice struct {
	FromAddr string `json:"From Address"`
	ToAddr   string `json:"To Address"`
	Price    string `json:"Price"`
}

type ShippingPrices []ShippingPrice

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to SHIBCO Home Page!")
	fmt.Printf("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/all", returnAllPrices)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func returnAllPrices(w http.ResponseWriter, r *http.Request) {
	prices := ShippingPrices{
		ShippingPrice{FromAddr: "UAE, Dubai, JLT, Cluster P, AstroLab", ToAddr: "UAE, Sharjar, Airport, Terminal 1",
			Price: "100 AED"},
		ShippingPrice{FromAddr: "UAE, Sharjar, Airport, Terminal 1", ToAddr: "UAE, Dubai, JLT, Cluster P, AstroLab",
			Price: "100 AED"},
	}
	fmt.Println("Endpoint Hit: returnAllPrice")

	json.NewEncoder(w).Encode(prices)
}

func main() {
	handleRequests()
}
