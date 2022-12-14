package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"post-data/models"
	"strconv"
)

func main() {
	num := 1

	for i := 1; i <= 4; i++ {
		str := ".xml"
		name := strconv.Itoa(num) + str

		xmlFile, err := os.Open(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Successfully Opened ")

		defer xmlFile.Close()

		XMLdata, _ := ioutil.ReadAll(xmlFile)

		var data models.TrademarkApplication
		xml.Unmarshal([]byte(XMLdata), &data)

		jsonData, _ := json.MarshalIndent(data, "", "\t")

		fmt.Println(string(jsonData))
		num++
	}
}
