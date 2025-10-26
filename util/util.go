package util

import (
	"github.com/GrolimundSolutions/syntheticMonitor/data"
	"github.com/tcnksm/go-httpstat"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

// Count returns the count of url object
func Count(data *data.SyntheticSettings) int {
	return len(data.SyntheticUrls)
}

// SyntheticCall calls the website and generate the detail information
func SyntheticCall(urlObject data.SyntheticUrls, channel chan data.ResponseObject) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", urlObject.URL, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Create a httpstat powered context
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)
	// Send request by default HTTP client
	client := http.DefaultClient
	start := time.Now()
	res, err := client.Do(req)
	//defer res.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err := io.Copy(io.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var timeTotal = time.Since(start) / time.Millisecond
	channel <- data.ResponseObject{
		Name:             urlObject.Name,
		URL:              urlObject.URL,
		HTTPStatus:       int16(res.StatusCode),
		TotalTime:        int16(timeTotal),
		DNSLookup:        int16(result.DNSLookup / time.Millisecond),
		TCPConnection:    int16(result.TCPConnection / time.Millisecond),
		TLSHandshake:     int16(result.TLSHandshake / time.Millisecond),
		ServerProcessing: int16(result.ServerProcessing / time.Millisecond),
		ContentTransfer:  int16(result.ContentTransfer(time.Now()) / time.Millisecond),
	}

}

// windowsOS const for GOOS windows
const windowsOS = "windows"

// GetDefaultConfigPath returns the default path for the actual OS
func GetDefaultConfigPath() string {
	if runtime.GOOS == windowsOS {
		return "\\ProgramData\\SyntheticMonitor\\urlList.json"
	}
	return "/etc/syntheticmonitor/urlList.json"
}
