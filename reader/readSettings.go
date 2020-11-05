package reader

import (
	"encoding/json"
	"github.com/GrolimundSolutions/syntheticMonitor/data"
	"io/ioutil"
	"log"
	"os"
)

func ReadFromJson(filePath string) (data.SyntheticSettings, error) {
	// Read JSON file to struct
	var synset data.SyntheticSettings

	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
		return synset, err
	}
	log.Printf("Successfully Opened %s", filePath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer func() {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
		return synset, err
	}
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &synset)
	if err != nil {
		log.Fatal(err)
		return synset, err
	}
	return synset, nil
}
