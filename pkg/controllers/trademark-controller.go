package controllers

import (
	"Sample-APIs/pkg/config"
	"Sample-APIs/pkg/models"
	"encoding/json"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	_ "github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
)

var DB *gorm.DB

func init() {
	config.Connect()
	DB = config.GetDB()
	DB.AutoMigrate(&models.TrademarkTable{})

}

func HandleError(c *fiber.Ctx) error {
	c.Status(400)
	var error datatypes.JSON
	error = datatypes.JSON(`{
		"error":"wrong end point"	
		}`)
	return c.JSON(error)
}

func GetDataByName(c *fiber.Ctx) error {
	n := c.Params("name")
	name := strings.Replace(n, "%20", " ", -1)
	country := c.Query("country")

	var res []models.TrademarkTable

	var check *gorm.DB

	if country != "" {
		check = DB.Where("trademark_name ILIKE ? and country ILIKE ?", "%"+name+"%", country).Find(&res)
	} else {
		check = DB.Where("trademark_name ILIKE ?", "%"+name+"%").Find(&res)
	}

	if check.RowsAffected == 0 {
		c.Status(404)
		var error datatypes.JSON
		error = datatypes.JSON(`{
		"error":"trademark not found"	
		}`)
		return c.JSON(error)
	}

	var data models.TrademarkApplication

	DataMap := make([]interface{}, len(res))
	for i := 0; i < len(res); i++ {
		err := json.Unmarshal(res[i].Attributes, &data)
		if err != nil {
			return nil
		}

		dmap := GetDto(data)
		DataMap[i] = dmap
	}

	return c.JSON(DataMap)

}

func GetDataById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var res models.TrademarkTable

	check := DB.Select("attributes").Where("application_number= ?", id).Find(&res)
	if check.RowsAffected == 0 {
		c.Status(404)
		var error datatypes.JSON
		error = datatypes.JSON(`{
		"error":"trademark not found"	
		}`)
		return c.JSON(error)
	}

	var data models.TrademarkApplication
	err := json.Unmarshal(res.Attributes, &data)
	if err != nil {
		return nil
	}

	//var DataMap map[string]interface{}
	var DataMap interface{}
	DataMap = GetDto(data)

	return c.JSON(DataMap)
}

func GetDto(data models.TrademarkApplication) *linkedhashmap.Map {

	DataMap := linkedhashmap.New()

	DataMap.Put("trademark_name", data.TrademarkBag.Trademark.ApplicantBag.Applicant.LegalEntityName)
	DataMap.Put("application_number", data.TrademarkBag.Trademark.ApplicationNumber.ST13ApplicationNumber)

	status := data.TrademarkBag.Trademark.NationalTrademarkInformation.MarkCurrentStatusInternalDescriptionText
	DataMap.Put("status", status)

	class := data.TrademarkBag.Trademark.GoodsServicesBag.GoodsServices.GoodsServicesClassificationBag.GoodsServicesClassification[0].ClassTitleText
	DataMap.Put("class", class)

	description := data.TrademarkBag.Trademark.GoodsServicesBag.GoodsServices.ClassDescriptionBag.ClassDescription[0].GoodsServicesDescriptionText
	DataMap.Put("class_description", description)

	registrationNum := data.TrademarkBag.Trademark.RegistrationNumber
	DataMap.Put("registration_number", registrationNum)

	registrationDate := data.TrademarkBag.Trademark.RegistrationDate
	DataMap.Put("registration_date", registrationDate)

	renewalDate := data.TrademarkBag.Trademark.NationalTrademarkInformation.RenewalDate
	DataMap.Put("renewal_date", renewalDate)

	abandonedDate := data.TrademarkBag.Trademark.NationalTrademarkInformation.ApplicationAbandonedDate
	DataMap.Put("abandoned_date", abandonedDate)

	return DataMap
}
