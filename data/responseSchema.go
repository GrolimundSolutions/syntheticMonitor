package data

type ResponseObjects struct {
	ResponseObject []ResponseObject `json:"ResponseObject"`
}

type ResponseObject struct {
	Name             string `json:"Name"`
	URL              string `json:"Url"`
	HttpStatus       int16  `json:"HTTP_status"`
	TotalTime        int16  `json:"Total_time"`
	DNSLookup        int16  `json:"Dns_lookup"`
	TCPConnection    int16  `json:"Tcp_connection"`
	TLSHandshake     int16  `json:"Tls_handshake"`
	ServerProcessing int16  `json:"Server_processing"`
	ContentTransfer  int16  `json:"Content_transfare"`
}
