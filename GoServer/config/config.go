package config



import (

	"os"
	"encoding/json"

)

//config entries

type Configuration struct{

	Database DatabaseConfig

}

//database config entries

type DatabaseConfig struct{

	Url string
	Name string
	Tables DatabaseTables
}

type DatabaseTables struct{
	Items string
}

var MainConfig = GetConfig("config.json")


//gets the config from json file

func GetConfig(filePath string) *Configuration{

	file, _ := os.Open(filePath)

	decoder := json.NewDecoder(file)

	config := Configuration{}



	err := decoder.Decode(&config)



	if err != nil{

		panic(err)

	}



	return &config

}