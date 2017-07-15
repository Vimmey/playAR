package main

import (
	"fmt"
	"log"
	"net/http"
	"playAR/apis/common/connections"
	"playAR/apis/model/restaurants"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

var mysqlconn *connections.MysqlConn

func main() {

	mysqlconn = connections.GetNewMysqlConn()
	mysqlconn.InitMysqlClient()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/restaurants", restaurantsApi)
	router.HandleFunc("/menu", menu)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	spew.Dump(restaurants.GetById(mysqlconn.Dbm, 69859898))
	spew.Dump(restaurants.GetByLatLong(mysqlconn.Dbm, "17.934378", "75.613404"))
	fmt.Fprintln(w, "Team playAR rocks!!!")
}

func restaurantsApi(w http.ResponseWriter, r *http.Request) {

	t1 := time.Now()
	log.Println("/restaurant API hit at time", t1)
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func menu(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	log.Println("/menu API hit at time", t1)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintln(w, "{\x22menu_details\x22 : [{\x22item_name\x22: \x22Pepperoni Supremo Pizza \x22,\x22item_price\x22: \x221350.00\x22},{\x22item_name\x22: \x22Spicy Medi Pizza \x22,\x22item_price\x22: \x22490.00\x22},{\x22item_name\x22: \x22Indiano Rustica Pizza \x22,\x22item_price\x22: \x221070.00\x22},{\x22item_name\x22: \x22Frijole Pizza \x22,\x22item_price\x22: \x22690.00\x22},{\x22item_name\x22: \x22Chicken Spicy Cilantro Pizza \x22,\x22item_price\x22: \x221010.00\x22},{\x22item_name\x22: \x22Chicken Fajita Pizza \x22,\x22item_price\x22: \x221230.00\x22},{\x22item_name\x22: \x22Meat Cravers Pizza \x22,\x22item_price\x22: \x221450.00\x22},{\x22item_name\x22: \x22Chipotle Chicken Pizza \x22,\x22item_price\x22: \x221190.00\x22},{\x22item_name\x22: \x22Roasted Garlic Chicken Pizza \x22,\x22item_price\x22: \x221190.00\x22},{\x22item_name\x22: \x22Fanta\x22,\x22item_price\x22: \x2270.00\x22},{\x22item_name\x22: \x22California Kung Pao Spaghetti Pasta\x22,\x22item_price\x22: \x22850.00\x22},{\x22item_name\x22: \x22Original BBQ Chicken Sandwich\x22,\x22item_price\x22: \x22850.00\x22},{\x22item_name\x22: \x22Chocolate Trinity\x22,\x22item_price\x22: \x22590.00\x22},{\x22item_name\x22: \x22Tomato Basil Spaghetti Pasta\x22,\x22item_price\x22: \x22950.00\x22},{\x22item_name\x22: \x22Red Velvet\x22,\x22item_price\x22: \x22630.00\x22}]}")
}
