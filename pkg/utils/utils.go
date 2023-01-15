package utils

import (
	"Sample-APIs/pkg/config"
	"Sample-APIs/pkg/models"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"io"
	"os"
	"strconv"
	"strings"
)

var DB *gorm.DB

func init() {
	config.Connect()
	DB = config.GetDB()
}

func Parsing() {
	DB.AutoMigrate(models.TrademarkTable{})
	num := 1057613
	for i := 1; i < 19529; i++ {
		str := "-00.xml"
		name := strconv.Itoa(num) + str
		file, err := os.Open("data/" + name)
		if err != nil {
			fmt.Println("open error")
			return
		}
		defer file.Close()
		Xmldata, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("xml data error")
			return
		}
		var data models.TrademarkApplication

		xml.Unmarshal([]byte(Xmldata), &data)
		id := data.TrademarkBag.Trademark.ApplicationNumber.ST13ApplicationNumber
		trademarkName := data.TrademarkBag.Trademark.ApplicantBag.Applicant.LegalEntityName
		trademarkName = strings.Replace(trademarkName, ".", "", -1)
		trademarkName = strings.Replace(trademarkName, ",", "", -1)
		trademarkName = strings.Replace(trademarkName, "-", " ", -1)
		country := data.TrademarkBag.Trademark.ApplicantBag.Applicant.NationalLegalEntityCode

		var jsondata datatypes.JSON

		jsondata, err = json.Marshal(data)
		if err != nil {
			return
		}

		var test models.TrademarkTable
		test.ApplicationNumber = int(id)
		test.TrademarkName = trademarkName
		test.Country = country
		test.Attributes = jsondata

		DB.Create(&test)
		fmt.Println(num)
		num++
	}

}
