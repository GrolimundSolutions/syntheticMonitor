package writer

import (
	"encoding/json"
	"grolimundSolutions.ch/syntheticMonitor/data"
	"log"
	"os"
)

func WriteToJson(filePath *string, objectSchema *data.SyntheticSettings) error {
	// Write struct to JSON file
	file, err := os.OpenFile(*filePath, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(objectSchema)
	if err != nil {
		return err
	}
	return nil
}