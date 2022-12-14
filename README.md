# Trademark-GoLang

- Read XML files and convert XML to JSON. 
- And then storing the parsed data in the Database
- An API for searching the Trademark Details using application number


## Tech Stack

- `Backend Framework:` `GoLang`
- `Databases:` `MySQL`
- `Libraries` `Gorilla Mux, GORM`


## To recruitments@trademarkia.com

### Approach to Parse XML file to JSON
 - First I stored the given files in post-data folder (I did not commit those to githud as there were more than 20,000 files).
 - Wrote a script to open each file and unmarshlled the XML file and then marshled it into JSON data.
 - Using a small cmd script I renamed all the files in sequencal order:
	
	- Example

 		```
		 1.xml
		 2.xml
		 3.xml
		 4.xml 
		 . 
		 .
		 19531.xml
		```
- post-data.models.TrademarkApplication

	I used this stuct to store the XML data from the file. 

 - Below is my main function in post-data directory which parses all the XML files stored in the directory:

	```go
	num := 1

	for i := 1; i <= 19531; i++ {
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
	```


### Approach to POST the parsed JSON to a database

- My idea was to have two GoLang applications, one which handles all the requests and interactions with the database, and the other application to parse the XML files and POST the data to the main GoLang application so that the parsed data can be stored into a database.

- So I created a POST endpoint in the main GoLang application 

- My main GoLang application was receiving the data(checked it by printing the request in the console), but it was not able to Marshal the request into the struct which I declared. 

- I used GORM, it  did not create my entire struct, instead it  created a table with fields only for a few attributes.
- I checked my json tags in the struct it was matching with the JSON but still I could not find the error.
- Mostly I build rest services and Microservices using Java spring boot, recently I started doing it in GoLang and I'm loving it, I have not able to complete the assignment that is creating an API and dockerizing it , but that does not mean I do not know how to do it, I did some small projects using GoLang and building REST APIs, if you have time please see my repositories. Thank you                  
- Output of the parsed 1.xml

	```json
	{
   	"xml_name": {
      "Space": "http://www.wipo.int/standards/XMLSchema/ST96/Trademark",
      "Local": "TrademarkApplication"
   },
   "catmk": "http://www.cipo.ic.gc.ca/standards/XMLSchema/ST96/Trademark",
   "ns_4": "http://www.w3.org/1998/Math/MathML",
   "ns_3": "http://www.oasis-open.org/tables/exchange/1.0",
   "com": "http://www.wipo.int/standards/XMLSchema/ST96/Common",
   "cacom": "http://www.cipo.ic.gc.ca/standards/XMLSchema/ST96/Common",
   "xs": "http://www.w3.org/2001/XMLSchema",
   "tmk": "http://www.wipo.int/standards/XMLSchema/ST96/Trademark",
   "xsi": "http://www.w3.org/2001/XMLSchema-instance",
   "st_96_version": "V3_0",
   "ipo_version": "V1_4",
   "request_search": {
      "ID": 0,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "DeletedAt": null,
      "request_search_category": "Normal"
   },
   "request_examination": {
      "ID": 0,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "DeletedAt": null,
      "request_examination_category": "Normal"
   },
   "trademark_bag": {
      "ID": 0,
      "CreatedAt": "0001-01-01T00:00:00Z",
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "DeletedAt": null,
      "trademark": {
         "ID": 0,
         "CreatedAt": "0001-01-01T00:00:00Z",
         "UpdatedAt": "0001-01-01T00:00:00Z",
         "DeletedAt": null,
         "operation_category": "Update",
         "registration_office_code": "CA",
         "receiving_office_code": "CA",
         "receiving_office_date": "2001-12-06",
         "application_number": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "ip_office_code": "CA",
            "st_13_application_number": "300000105761300"
         },
         "registration_number": "TMA580662",
         "application_date": "2000-05-03",
         "registration_date": "2003-05-06",
         "filing_place": "Canada",
         "application_language_code": "en",
         "expiry_date": "2033-05-06",
         "mark_current_status_code": "Registration published",
         "mark_current_status_date": "2003-05-06",
         "mark_category": "Trademark",
         "mark_representation": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "mark_feature_category": "Figurative",
            "mark_reproduction": {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "mark_image_bag": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_image": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "file_name": "1057613.png",
                     "image_format_category": "PNG",
                     "mark_image_classification": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "figurative_element_classification_bag": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "vienna_classification_bag": {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "vienna_classification": [
                                 {
                                    "ID": 0,
                                    "CreatedAt": "0001-01-01T00:00:00Z",
                                    "UpdatedAt": "0001-01-01T00:00:00Z",
                                    "DeletedAt": null,
                                    "type": "cacom:ViennaClassificationType",
                                    "vienna_category": "4",
                                    "vienna_division": "5",
                                    "vienna_section": "15",
                                    "vienna_description_bag": {
                                       "ID": 0,
                                       "CreatedAt": "0001-01-01T00:00:00Z",
                                       "UpdatedAt": "0001-01-01T00:00:00Z",
                                       "DeletedAt": null,
                                       "vienna_description": [
                                          {
                                             "ID": 0,
                                             "CreatedAt": "0001-01-01T00:00:00Z",
                                             "UpdatedAt": "0001-01-01T00:00:00Z",
                                             "DeletedAt": null,
                                             "language_code": "en"
                                          },
                                          {
                                             "ID": 0,
                                             "CreatedAt": "0001-01-01T00:00:00Z",
                                             "UpdatedAt": "0001-01-01T00:00:00Z",
                                             "DeletedAt": null,
                                             "language_code": "fr"
                                          }
                                       ]
                                    }
                                 },
                                 {
                                    "ID": 0,
                                    "CreatedAt": "0001-01-01T00:00:00Z",
                                    "UpdatedAt": "0001-01-01T00:00:00Z",
                                    "DeletedAt": null,
                                    "type": "cacom:ViennaClassificationType",
                                    "vienna_category": "16",
                                    "vienna_division": "3",
                                    "vienna_section": "13",
                                    "vienna_description_bag": {
                                       "ID": 0,
                                       "CreatedAt": "0001-01-01T00:00:00Z",
                                       "UpdatedAt": "0001-01-01T00:00:00Z",
                                       "DeletedAt": null,
                                       "vienna_description": [
                                          {
                                             "ID": 0,
                                             "CreatedAt": "0001-01-01T00:00:00Z",
                                             "UpdatedAt": "0001-01-01T00:00:00Z",
                                             "DeletedAt": null,
                                             "language_code": "en"
                                          },
                                          {
                                             "ID": 0,
                                             "CreatedAt": "0001-01-01T00:00:00Z",
                                             "UpdatedAt": "0001-01-01T00:00:00Z",
                                             "DeletedAt": null,
                                             "language_code": "fr"
                                          }
                                       ]
                                    }
                                 }
                              ]
                           }
                        }
                     }
                  }
               }
            }
         },
         "non_use_cancelled_indicator": "false",
         "trade_distinctiveness_indicator": "false",
         "use_right": [
            {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "use_right_indicator": "true",
               "use_right_text": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "language_code": "en"
               }
            },
            {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "use_right_indicator": "true",
               "use_right_text": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "language_code": "en"
               }
            }
         ],
         "opposition_period_start_date": "2002-03-27",
         "opposition_period_end_date": "2002-05-27",
         "goods_services_bag": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "goods_services": {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "goods_services_classification_bag": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "goods_services_classification": [
                     {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "classification_kind_code": "Nice",
                        "comment_text": "CIPO Classification Code",
                        "class_number": "9",
                        "class_title_text": "Electrical, scientific and teaching apparatus and software"
                     },
                     {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "classification_kind_code": "Nice",
                        "comment_text": "CIPO Classification Code",
                        "class_number": "35",
                        "class_title_text": "Advertising, marketing, promotional and business"
                     },
                     {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "classification_kind_code": "Nice",
                        "comment_text": "CIPO Classification Code",
                        "class_number": "41",
                        "class_title_text": "Education and entertainment"
                     }
                  ]
               },
               "national_goods_services": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "national_class_total_quantity": "3"
               },
               "national_filing_basis": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "current_basis": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "basis_foreign_application_indicator": "false",
                     "basis_foreign_registration_indicator": "true",
                     "basis_use_indicator": "true",
                     "basis_intent_to_use_indicator": "false",
                     "no_basis_indicator": "false"
                  }
               },
               "classification_kind_code": "Nice",
               "class_description_bag": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "class_description": [
                     {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "goods_services_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "sequence_number": "Goods1",
                           "language_code": "en"
                        }
                     },
                     {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "goods_services_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "sequence_number": "Services1",
                           "language_code": "en"
                        }
                     }
                  ]
               }
            }
         },
         "priority_bag": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "priority": [
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "priority_country_code": "US",
                  "application_number": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "application_number_text": "75/847,216"
                  },
                  "priority_application_filing_date": "1999-11-12",
                  "priority_partial_goods_services": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "class_description_bag": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "class_description": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "goods_services_description_text": {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "sequence_number": "Goods1",
                              "language_code": "en"
                           }
                        }
                     }
                  }
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "priority_country_code": "US",
                  "application_number": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "application_number_text": "75/863,750"
                  },
                  "priority_application_filing_date": "1999-11-12",
                  "priority_partial_goods_services": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "class_description_bag": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "class_description": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "goods_services_description_text": {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "sequence_number": "Services1",
                              "language_code": "en"
                           }
                        }
                     }
                  }
               }
            ]
         },
         "publication_bag": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "publication": {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "publication_identifier": "Vol.49 Issue 2474",
               "national_publication": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "publication_status_category": "Advertised",
                  "publication_action_date": "2002-03-27"
               }
            }
         },
         "applicant_bag": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "applicant": {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "legal_entity_name": "ALEKS CORPORATION,",
               "contact": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "language_code": "en",
                  "name": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "entity_name": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en"
                     }
                  },
                  "postal_address_bag": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "postal_address": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "postal_structured_address": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "en",
                           "address_line_text": [
                              {
                                 "ID": 0,
                                 "CreatedAt": "0001-01-01T00:00:00Z",
                                 "UpdatedAt": "0001-01-01T00:00:00Z",
                                 "DeletedAt": null,
                                 "sequence_number": "1"
                              },
                              {
                                 "ID": 0,
                                 "CreatedAt": "0001-01-01T00:00:00Z",
                                 "UpdatedAt": "0001-01-01T00:00:00Z",
                                 "DeletedAt": null,
                                 "sequence_number": "2"
                              }
                           ],
                           "country_code": "US"
                        }
                     }
                  }
               },
               "national_legal_entity_code": "US"
            }
         },
         "national_representative": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "comment_text": "6208",
            "contact": {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "language_code": "en",
               "name": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "entity_name": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "language_code": "en"
                  }
               },
               "postal_address_bag": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "postal_address": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "postal_structured_address": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en",
                        "address_line_text": [
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "sequence_number": "1"
                           },
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "sequence_number": "2"
                           },
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "sequence_number": "3"
                           }
                        ],
                        "geographic_region_name": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "geographic_region_category": "Province"
                        },
                        "country_code": "CA",
                        "postal_code": "K1Y4S1"
                     }
                  }
               }
            }
         },
         "mark_event_bag": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "mark_event": [
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "Application filed",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "30",
                     "mark_event_description_text": "Filed",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     }
                  },
                  "mark_event_date": "2000-05-03"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "National prosecution history entry",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "1",
                     "mark_event_description_text": "Created",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     }
                  },
                  "mark_event_date": "2000-05-10"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "Filing date accorded",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "31",
                     "mark_event_description_text": "Formalized",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     }
                  },
                  "mark_event_date": "2000-06-08"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "National prosecution history entry",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "22",
                     "mark_event_description_text": "Search Recorded",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     }
                  },
                  "mark_event_date": "2001-09-28"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "National prosecution history entry",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "20",
                     "mark_event_description_text": "Examiner's First Report",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     }
                  },
                  "mark_event_date": "2001-10-09",
                  "mark_event_response_date": "2002-02-09"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "Application accepted",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "26",
                     "mark_event_description_text": "Approved",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     }
                  },
                  "mark_event_date": "2002-02-22"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "Application published",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "42",
                     "mark_event_description_text": "Advertised",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     },
                     "mark_event_additional_text": "Vol.49 Issue 2474"
                  },
                  "mark_event_date": "2002-03-27"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "National prosecution history entry",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "50",
                     "mark_event_description_text": "Allowed",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     }
                  },
                  "mark_event_date": "2002-06-13"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "National prosecution history entry",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "51",
                     "mark_event_description_text": "Allowance Notice Sent",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     }
                  },
                  "mark_event_date": "2002-06-13",
                  "mark_event_response_date": "2003-05-03"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "Registration published",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "55",
                     "mark_event_description_text": "Registered",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     }
                  },
                  "mark_event_date": "2003-05-06",
                  "mark_event_response_date": "2018-05-06"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "National prosecution history entry",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "87",
                     "mark_event_description_text": "Rep for Service Changed",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     },
                     "mark_event_additional_text": "From: 12   To: 6208  /  Voir Preuve au dossier/See evidence on File No. 1056844"
                  },
                  "mark_event_date": "2014-02-28"
               },
               {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "mark_event_category": "National prosecution history entry",
                  "national_mark_event": {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "type": "catmk:NationalMarkEventType",
                     "mark_event_code": "61",
                     "mark_event_description_text": "Renewed",
                     "mark_event_other_language_description_text_bag": {
                        "mark_event_description_text": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "language_code": "fr"
                        }
                     },
                     "mark_event_additional_text": "DP:2018/05/04  RD:2018/05/04  RR:(6208)  MOFFAT & CO."
                  },
                  "mark_event_date": "2018-05-06",
                  "mark_event_response_date": "2033-05-06"
               }
            ]
         },
         "mark_feature_description": "FANCIFUL DOG DESIGN",
         "national_trademark_information": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "type": "catmk:NationalTrademarkInformationType",
            "register_category": "Primary",
            "mark_current_status_internal_description_text": "Registered",
            "renewal_date": "2018-05-06",
            "allowed_date": "2002-06-13",
            "trademark_class": {
               "trademark_class_code": "1",
               "trademark_class_description_bag": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "trademark_class_description": [
                     {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en"
                     },
                     {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "fr"
                     }
                  ]
               }
            },
            "legislation": {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "legislation_code": "1",
               "legislation_description_bag": {
                  "ID": 0,
                  "CreatedAt": "0001-01-01T00:00:00Z",
                  "UpdatedAt": "0001-01-01T00:00:00Z",
                  "DeletedAt": null,
                  "legislation_description": [
                     {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en"
                     },
                     {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "fr"
                     }
                  ]
               }
            },
            "interested_party_bag": {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "interested_party": [
                  {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "interested_party_category": "Applicant",
                     "contact": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en",
                        "name": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "entity_name": {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "language_code": "en"
                           }
                        },
                        "postal_address_bag": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "postal_address": {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "postal_structured_address": {
                                 "ID": 0,
                                 "CreatedAt": "0001-01-01T00:00:00Z",
                                 "UpdatedAt": "0001-01-01T00:00:00Z",
                                 "DeletedAt": null,
                                 "language_code": "en",
                                 "address_line_text": [
                                    {
                                       "ID": 0,
                                       "CreatedAt": "0001-01-01T00:00:00Z",
                                       "UpdatedAt": "0001-01-01T00:00:00Z",
                                       "DeletedAt": null,
                                       "sequence_number": "1"
                                    },
                                    {
                                       "ID": 0,
                                       "CreatedAt": "0001-01-01T00:00:00Z",
                                       "UpdatedAt": "0001-01-01T00:00:00Z",
                                       "DeletedAt": null,
                                       "sequence_number": "2"
                                    }
                                 ],
                                 "country_code": "US"
                              }
                           }
                        }
                     }
                  },
                  {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "interested_party_category": "Registrant",
                     "contact": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en",
                        "name": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "entity_name": {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "language_code": "en"
                           }
                        },
                        "postal_address_bag": {
                           "ID": 0,
                           "CreatedAt": "0001-01-01T00:00:00Z",
                           "UpdatedAt": "0001-01-01T00:00:00Z",
                           "DeletedAt": null,
                           "postal_address": {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "postal_structured_address": {
                                 "ID": 0,
                                 "CreatedAt": "0001-01-01T00:00:00Z",
                                 "UpdatedAt": "0001-01-01T00:00:00Z",
                                 "DeletedAt": null,
                                 "language_code": "en",
                                 "address_line_text": [
                                    {
                                       "ID": 0,
                                       "CreatedAt": "0001-01-01T00:00:00Z",
                                       "UpdatedAt": "0001-01-01T00:00:00Z",
                                       "DeletedAt": null,
                                       "sequence_number": "1"
                                    },
                                    {
                                       "ID": 0,
                                       "CreatedAt": "0001-01-01T00:00:00Z",
                                       "UpdatedAt": "0001-01-01T00:00:00Z",
                                       "DeletedAt": null,
                                       "sequence_number": "2"
                                    }
                                 ],
                                 "country_code": "US"
                              }
                           }
                        }
                     }
                  }
               ]
            },
            "claim_bag": {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "claim": [
                  {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "claim_category_type": "12",
                     "claim_type_description": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en"
                     },
                     "claim_number": "1",
                     "claim_date": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "structured_claim_date": "1999-11-12"
                     },
                     "claim_country_code": "US",
                     "claim_foreign_registration_nbr": "75/847,216",
                     "claim_text": "Priority Filing Date: November 12, 1999, Country: UNITED STATES OF AMERICA, Application No. 75/847,216 in association with the same kind of goods",
                     "goods_services_reference_bag": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "goods_services_reference": [
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "goods_services_reference_identifier": "Goods1"
                           }
                        ]
                     }
                  },
                  {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "claim_category_type": "12",
                     "claim_type_description": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en"
                     },
                     "claim_number": "2",
                     "claim_date": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "structured_claim_date": "1999-11-12"
                     },
                     "claim_country_code": "US",
                     "claim_foreign_registration_nbr": "75/863,750",
                     "claim_text": "Priority Filing Date: November 12, 1999, Country: UNITED STATES OF AMERICA, Application No. 75/863,750 in association with the same kind of services",
                     "goods_services_reference_bag": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "goods_services_reference": [
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "goods_services_reference_identifier": "Services1"
                           }
                        ]
                     }
                  },
                  {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "claim_category_type": "13",
                     "claim_type_description": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en"
                     },
                     "claim_number": "1",
                     "claim_date": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null
                     },
                     "claim_country_code": "US",
                     "claim_text": "Used in UNITED STATES OF AMERICA on goods",
                     "goods_services_reference_bag": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "goods_services_reference": [
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "goods_services_reference_identifier": "Goods1"
                           }
                        ]
                     }
                  },
                  {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "claim_category_type": "14",
                     "claim_type_description": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en"
                     },
                     "claim_number": "1",
                     "claim_date": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "structured_claim_date": "2001-05-22"
                     },
                     "claim_country_code": "US",
                     "claim_foreign_registration_nbr": "2,452,773",
                     "claim_text": "Registered in or for UNITED STATES OF AMERICA on May 22, 2001, under No. 2,452,773 on services",
                     "goods_services_reference_bag": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "goods_services_reference": [
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "goods_services_reference_identifier": "Services1"
                           }
                        ]
                     }
                  },
                  {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "claim_category_type": "14",
                     "claim_type_description": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en"
                     },
                     "claim_number": "2",
                     "claim_date": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "structured_claim_date": "2000-10-03"
                     },
                     "claim_country_code": "US",
                     "claim_foreign_registration_nbr": "2,391,684",
                     "claim_text": "Registered in or for UNITED STATES OF AMERICA on October 03, 2000, under No. 2,391,684 on goods",
                     "goods_services_reference_bag": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "goods_services_reference": [
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "goods_services_reference_identifier": "Goods1"
                           }
                        ]
                     }
                  },
                  {
                     "ID": 0,
                     "CreatedAt": "0001-01-01T00:00:00Z",
                     "UpdatedAt": "0001-01-01T00:00:00Z",
                     "DeletedAt": null,
                     "claim_category_type": "16",
                     "claim_type_description": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "language_code": "en"
                     },
                     "claim_number": "1",
                     "claim_date": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "structured_claim_date": "2003-04-15"
                     },
                     "claim_text": "Declaration of Use filed April 15, 2003",
                     "goods_services_reference_bag": {
                        "ID": 0,
                        "CreatedAt": "0001-01-01T00:00:00Z",
                        "UpdatedAt": "0001-01-01T00:00:00Z",
                        "DeletedAt": null,
                        "goods_services_reference": [
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "goods_services_reference_identifier": "Goods1"
                           },
                           {
                              "ID": 0,
                              "CreatedAt": "0001-01-01T00:00:00Z",
                              "UpdatedAt": "0001-01-01T00:00:00Z",
                              "DeletedAt": null,
                              "goods_services_reference_identifier": "Services1"
                           }
                        ]
                     }
                  }
               ]
            },
            "mark_feature_category_bag": {
               "ID": 0,
               "CreatedAt": "0001-01-01T00:00:00Z",
               "UpdatedAt": "0001-01-01T00:00:00Z",
               "DeletedAt": null,
               "mark_feature_category": "Figurative"
            }
         }
      }
   	}
	}

