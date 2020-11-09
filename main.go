package main

import (
	"fmt"
	"github.com/GrolimundSolutions/syntheticMonitor/data"
	"github.com/GrolimundSolutions/syntheticMonitor/pkg"
	"github.com/GrolimundSolutions/syntheticMonitor/reader"
	"github.com/GrolimundSolutions/syntheticMonitor/util"
	"github.com/GrolimundSolutions/syntheticMonitor/writer"
	"log"
	"os"
	"strings"
)

func init() {
	defaultData := data.SyntheticSettings{
		Location: "Swiss",
		SyntheticUrls: []data.SyntheticUrls{
			data.SyntheticUrls{
				URL:    "https://google.ch",
				Name:   "Google Swiss",
				Expect: "In Progress",
			},
			data.SyntheticUrls{
				URL:    "https://google.de",
				Name:   "Google Germany",
				Expect: "In Progress",
			},
		},
	}

	defaultPath := strings.Trim(util.GetDefaultConfigPath(), "urlList.json")

	// Check if the defaultlocation exists
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		log.Println("Info default settings folder not exists")
		err := os.MkdirAll(defaultPath, 0700)
		log.Println("Info default settings folder created")
		if err != nil {
			panic(err.Error())
		}
		// Create a default file
		err = writer.WriteToJSON(&defaultData)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Info default settings file created")
	}
	// Check if the file exists
	if _, err := os.Stat(util.GetDefaultConfigPath()); os.IsNotExist(err) {
		log.Println("Info default settings file not found")
		// Create a default file
		err = writer.WriteToJSON(&defaultData)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Info default settings file created")
	}
}

func main() {

	myChannel := make(chan data.ResponseObject)
	var resObj = data.ResponseObjects{}
	jsonData, err := reader.ReadFromJSON()
	if err != nil {
		log.Fatalln(err.Error())
	}

	urlObjectCount := util.Count(&jsonData)
	log.Println("urlobjects: ", urlObjectCount)

	for i := 0; i <= urlObjectCount-1; i++ {
		go util.SyntheticCall(jsonData.SyntheticUrls[i], myChannel)
	}

	for j := 0; j <= urlObjectCount-1; j++ {
		retItem := <-myChannel
		log.Println(retItem)
		resObj.ResponseObject = append(resObj.ResponseObject, retItem)
	}
	fmt.Println(resObj)

	res := pkg.Result(resObj)

	switch jsonData.EndpointType {
	case "URL":
		err = res.SendToHTTP(&jsonData)
		if err != nil {
			log.Fatalln(err.Error())
		}
	case "File":
		log.Println("File")
		log.Println(jsonData.FileLocation)
		err = res.WriteToJSON(jsonData.FileLocation)
	default:
		fmt.Println("EndpointType::NIL")
	}

}
