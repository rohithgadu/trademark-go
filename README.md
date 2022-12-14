# Trademark-GoLang

- Read XML files and convert XML to JSON. 
- And then storing the parsed data in the Database
- An API for searching the Trademark Details using application number


## Tech Stack

- `Backend Framework:` `GoLang`
- `Databases:` `MySQL`
- `Libraries` `Gorilla Mux, GORM`


## Message to recruitments@trademarkia.com

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


