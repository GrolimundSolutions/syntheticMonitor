package main

import (
	"fmt"
	"github.com/GrolimundSolutions/syntheticMonitor/data"
	"github.com/GrolimundSolutions/syntheticMonitor/reader"
	"github.com/GrolimundSolutions/syntheticMonitor/util"
	"log"
)

func main() {

	myChannel := make(chan data.ResponseObject)
	var resObj = data.ResponseObjects{}
	jsonData, err := reader.ReadFromJSON("urlList.json")
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
}
