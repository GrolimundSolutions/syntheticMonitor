package pkg

import (
	"bytes"
	"encoding/json"
	"github.com/GrolimundSolutions/syntheticMonitor/data"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Results declare a Interface
type Results interface {
	WriteToYAML(string) error
	WriteToJSON(string) error
	SendToHTTP(string) error
}

// Result type declare for the Interface
type Result data.ResponseObjects

// WriteToYAML save the results to a yaml file
func (r Result) WriteToYAML(path string) error {

	filePath := path
	// Write struct to JSON file
	file, err := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	y, err := yaml.Marshal(r)
	if err != nil {
		log.Printf("err: %v\n", err)
		return err
	}
	_, err = file.Write(y)
	if err != nil {
		log.Printf("err: %v\n", err)
		return err
	}

	log.Println(string(y))
	return nil
}

// WriteToJSON save the results to a json file
func (r Result) WriteToJSON(path string) error {
	filePath := path
	// Write struct to JSON file
	file, err := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
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
	err = encoder.Encode(r)
	if err != nil {
		return err
	}
	return nil
}

// SendToHTTP send the results to a url
func (r Result) SendToHTTP(url string) error {
	requestBody, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err.Error())
	}

	resp, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(string(body))
	return nil
}
