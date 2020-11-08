# SyntheticMonitor
<p align="center">
  <a href="https://github.com/GrolimundSolutions/syntheticMonitor/issues"><img alt="GitHub issues" src="https://img.shields.io/github/issues/GrolimundSolutions/syntheticMonitor"></a>
  <a href="https://github.com/GrolimundSolutions/syntheticMonitor"><img alt="GitHub license" src="https://img.shields.io/github/license/GrolimundSolutions/syntheticMonitor"></a>
  <a href="https://github.com/GrolimundSolutions/syntheticMonitor/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/release/GrolimundSolutions/syntheticMonitor?logo=github&style=flat-square"></a>
  <a href="https://codecov.io/gh/GrolimundSolutions/syntheticMonitor">
    <img src="https://codecov.io/gh/GrolimundSolutions/syntheticMonitor/branch/master/graph/badge.svg?token=WNTU19RFJX"/>
  </a>
  <a href="https://goreportcard.com/report/github.com/GrolimundSolutions/syntheticMonitor"><img alt="Go Report" src="https://goreportcard.com/badge/github.com/GrolimundSolutions/syntheticMonitor"></a>
  <a href="#"> <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="Go"></a>
  <a href="https://github.com/GrolimundSolutions/syntheticMonitor/blob/master/LICENSE"> <img src="https://img.shields.io/github/license/GrolimundSolutions/SyntheticMonitor.svg" alt="Go"></a>
  <a href=""> <img src="https://img.shields.io/github/commits-since/GrolimundSolutions/SyntheticMonitor/latest" alt="Go"></a>
</p>

<hr>

Mit dem SyntethicMonitor kann man eine/mehrere url's Prüfen auf deren Geschwindigkeit und via Regex nach einem Inhalt suchen um zu Prüfen das die Webseite auch wirklich Online ist.
Die abarbeitung erfolgt via "concurrency" sprich, es wird Parallel abarbeitet.
Weiter erhaltet man Detailliertere Informationen zu den Request-Zeiten:

* Name             
* URL             
* HTTPStatus       
* TotalTime        
* DNSLookup        
* TCPConnection    
* TLSHandshake     
* ServerProcessing 
* *ContentTransfer

Json:
```json
[
	{
		"Name": "Google Swiss",
		"URL": "https://www.google.ch/",
		"HTTPStatus": "200",
		"TotalTime": "134",
		"DNSLookup": "16",
		"TCPConnection": "10",
		"TLSHandshake": "48",
		"ServerProcessing": "58",
		"ContentTransfer": "23"
	},
	{
		"Name": "shodan",
		"URL": "https://www.shodan.io/",
		"HTTPStatus": "200",
		"TotalTime": "319",
		"DNSLookup": "6",
		"TCPConnection": "13",
		"TLSHandshake": "43",
		"ServerProcessing": "257",
		"ContentTransfer": "23"
	}
]
```

Settings:
```json
{
	"Location": "Swiss",
	"SyntheticUrls": [
		{
			"URL": "https://google.ch",
			"Name": "Google Swiss",
			"Expect": "In Progress"
		},
		{
			"URL": "https://www.shodan.io/",
			"Name": "shodan",
			"Expect": "In Progress"
		}
	]
}
```


