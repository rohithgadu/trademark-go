package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-demo/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

	//num := 1
	//
	//for i := 1; i <= 1; i++ {
	//	str := ".xml"
	//	name := strconv.Itoa(num) + str
	//
	//	xmlFile, err := os.Open(name)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	fmt.Println("Successfully Opened ")
	//
	//	defer xmlFile.Close()
	//
	//	XMLdata, _ := ioutil.ReadAll(xmlFile)
	//
	//	var data models.TrademarkApplication
	//	xml.Unmarshal([]byte(XMLdata), &data)
	//
	//	jsonData, _ := json.MarshalIndent(data, "", "\t")
	//
	//	fmt.Println(string(jsonData))
	//	num++
	//}
}
