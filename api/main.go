package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"playAR/apis/common/connections"
	"playAR/apis/model/menu"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

var mysqlconn *connections.MysqlConn
var cred int

func main() {

	mysqlconn = connections.GetNewMysqlConn()
	mysqlconn.InitMysqlClient()

	cred = 70

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/credits", getCredit).Methods("GET")
	router.HandleFunc("/credits", updateCredit).Methods("PUT")
	router.HandleFunc("/restaurants", restaurantsApi)
	router.HandleFunc("/v2/restaurants", restaurantsApiV2)
	router.HandleFunc("/menu/{id}", menuApi)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	// spew.Dump(restaurants.GetById(mysqlconn.Dbm, 69859898))
	// spew.Dump(menu.GetById(mysqlconn.Dbm, 1939034))
	// spew.Dump(restaurants.GetByLatLong(mysqlconn.Dbm, "17.934378", "75.613404"))
	fmt.Fprintln(w, "Team playAR rocks!!!")
}

func restaurantsApi(w http.ResponseWriter, r *http.Request) {

	t1 := time.Now()
	log.Println("/restaurant API hit at time", t1)
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "[{\x22id\x22:\x22363222332\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_363222332\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x224200\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22363915580\x22,\x22latitude\x22:\x2215.459199\x22,\x22longitude\x22:\x2283.385102\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_363915580\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x223340\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22364256036\x22,\x22latitude\x22:\x2217.934672\x22,\x22longitude\x22:\x2275.616922\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_364256036\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x221100\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22364408085\x22,\x22latitude\x22:\x2217.936529\x22,\x22longitude\x22:\x2275.614497\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_364408085\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22364646334\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_364646334\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222344\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22365657581\x22,\x22latitude\x22:\x2215.459199\x22,\x22longitude\x22:\x2283.385102\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_365657581\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x223300\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22366087320\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_366087320\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22366681035\x22,\x22latitude\x22:\x2217.925656\x22,\x22longitude\x22:\x2275.625058\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_366681035\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222200\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22367575582\x22,\x22latitude\x22:\x2215.459199\x22,\x22longitude\x22:\x2283.385102\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_367575582\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x224400\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22368261078\x22,\x22latitude\x22:\x2217.933699\x22,\x22longitude\x22:\x2275.618562\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368261078\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22368347579\x22,\x22latitude\x22:\x2215.459199\x22,\x22longitude\x22:\x2283.385102\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368347579\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x221340\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22368708333\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368708333\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22368721133\x22,\x22latitude\x22:\x2217.934463\x22,\x22longitude\x22:\x2275.616111\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368721133\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22368919079\x22,\x22latitude\x22:\x2217.933699\x22,\x22longitude\x22:\x2275.618562\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368919079\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x224000\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22369037432\x22,\x22latitude\x22:\x2217.936811\x22,\x22longitude\x22:\x2275.620287\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_369037432\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22369879319\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_369879319\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222390\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22374223094\x22,\x22latitude\x22:\x2217.934248\x22,\x22longitude\x22:\x2275.616868\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_374223094\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x224100\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22},{\x22id\x22:\x22376463040\x22,\x22latitude\x22:\x2215.450625\x22,\x22longitude\x22:\x2283.390954\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_376463040\x22,\x22description\x22:\x22\x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x22\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22\x22}]")
}

func restaurantsApiV2(w http.ResponseWriter, r *http.Request) {

	t1 := time.Now()
	log.Println("/v2/restaurant API hit at time", t1)
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "{\x22restaurants\x22:[{\x22id\x22:\x22363222332\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_363222332\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x224200\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22363915580\x22,\x22latitude\x22:\x2215.459199\x22,\x22longitude\x22:\x2283.385102\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_363915580\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x223340\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22364256036\x22,\x22latitude\x22:\x2217.934672\x22,\x22longitude\x22:\x2275.616922\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_364256036\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x221100\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22364408085\x22,\x22latitude\x22:\x2217.936529\x22,\x22longitude\x22:\x2275.614497\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_364408085\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22364646334\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_364646334\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222344\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22365657581\x22,\x22latitude\x22:\x2215.459199\x22,\x22longitude\x22:\x2283.385102\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_365657581\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x223300\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22366087320\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_366087320\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22366681035\x22,\x22latitude\x22:\x2217.925656\x22,\x22longitude\x22:\x2275.625058\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_366681035\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222200\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22367575582\x22,\x22latitude\x22:\x2215.459199\x22,\x22longitude\x22:\x2283.385102\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_367575582\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x224400\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22368261078\x22,\x22latitude\x22:\x2217.933699\x22,\x22longitude\x22:\x2275.618562\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368261078\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22368347579\x22,\x22latitude\x22:\x2215.459199\x22,\x22longitude\x22:\x2283.385102\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368347579\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x221340\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22368708333\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368708333\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22368721133\x22,\x22latitude\x22:\x2217.934463\x22,\x22longitude\x22:\x2275.616111\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368721133\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22368919079\x22,\x22latitude\x22:\x2217.933699\x22,\x22longitude\x22:\x2275.618562\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_368919079\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x224000\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22369037432\x22,\x22latitude\x22:\x2217.936811\x22,\x22longitude\x22:\x2275.620287\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_369037432\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22369879319\x22,\x22latitude\x22:\x2217.934847\x22,\x22longitude\x22:\x2275.616123\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_369879319\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222390\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22374223094\x22,\x22latitude\x22:\x2217.934248\x22,\x22longitude\x22:\x2275.616868\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_374223094\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x224100\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22},{\x22id\x22:\x22376463040\x22,\x22latitude\x22:\x2215.450625\x22,\x22longitude\x22:\x2283.390954\x22,\x22altitude\x22:\x22\x22,\x22title\x22:\x22rest_376463040\x22,\x22description\x22:\x22Famous for its late night servings, dosa with chicken and shawarma served till endless hours in the night. \x22,\x22distance\x22:\x222340\x22,\x22rating\x22:\x224.8\x22,\x22pickupTime\x22:\x22\x22,\x22costForTwo\x22:\x22851\x22}]}")
}

func menuApi(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	log.Println("/menu API hit at time", t1)
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	m, err := menu.GetById(mysqlconn.Dbm, id)
	if err != nil {
		log.Println("/menu API ERROR = ", err)
	}

	finalArr := []menu.MenuRes{}

	for _, element := range m {
		if len(*element.Item_name) == 0 || *element.Item_name == "" {
			continue
		}
		m := menu.NewMenuRes(element.Item_name, fmt.Sprint(*element.Item_price))
		// jsonStr, err := json.Marshal(m)
		// spew.Dump(m)
		// spew.Dump(jsonStr)

		jsonStr, err := json.Marshal(*m)
		spew.Dump(*m)
		spew.Dump(jsonStr)

		log.Println("\n\n\n\n")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		finalArr = append(finalArr, *m)
	}

	// spew.Dump(finalArr)

	js, err := json.Marshal(finalArr)
	// spew.Dump(js)

	jsStr := string(js[:])
	jsStr = "{\"menu_details\" : " + jsStr + "}"
	// strings.Replace(jsStr,  "\", "\x22", -1)

	// w.Write(js)
	// spew.Dump(m)

	fmt.Fprintln(w, jsStr)

	// fmt.Fprintln(w, m)
	// fmt.Printf("%v", projects)

}

func getCredit(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	log.Println("/getCredit API hit at time", t1)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintln(w, creditResponse())

}

func updateCredit(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	log.Println("/updateCredit API hit at time", t1)
	w.Header().Set("Content-Type", "application/json")
	cred += 10
	fmt.Fprintln(w, creditResponse())
}

func creditResponse() string {
	credits := cred
	res := "{\x22id\x22:44354,\x22credits\x22:" + strconv.Itoa(credits) + "}"
	return res
}
