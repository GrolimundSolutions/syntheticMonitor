package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GrolimundSolutions/syntheticMonitor/data"
	"github.com/GrolimundSolutions/syntheticMonitor/util"
	"github.com/ghodss/yaml"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Results declare a Interface
type Results interface {
	WriteToYAML(string) error
	WriteToJSON(string) error
	SendToHTTP(settings *data.SyntheticSettings) error
}

// Result type declare for the Interface
type Result data.ResponseObjects

// WriteToYAML save the results to a yaml file
func (r Result) WriteToYAML(path string) error {

	filePath := path
	_time := time.Now()
	filename := fmt.Sprintf("%s.yaml", _time.Format("2006-01-02_15-04-05"))
	// Write struct to JSON file
	file, err := os.OpenFile(filePath+filename, os.O_CREATE, os.ModePerm)
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
	return nil
}

// WriteToJSON save the results to a json file
func (r Result) WriteToJSON(path string) error {
	filePath := path
	_time := time.Now()
	filename := fmt.Sprintf("%s.json", _time.Format("2006-01-02_15-04-05"))
	// Write struct to JSON file
	file, err := os.OpenFile(filePath+filename, os.O_CREATE, os.ModePerm)
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
func (r Result) SendToHTTP(settings *data.SyntheticSettings) error {
	client := &http.Client{}

	url, err := util.URLBuilder(settings.EndpointURL, settings.EndpointPort, settings.EndpointPath)
	requestBody, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err.Error())
	}
	req.Header.Add(settings.EndpointTokenKey, settings.EndpointTokenValue)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(string(body))
	return nil
}
