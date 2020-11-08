package data

// ResponseObjects holds an Array with the results of the test
type ResponseObjects struct {
	ResponseObject []ResponseObject `json:"ResponseObject"`
}

// ResponseObject detail information about 1 testcase
type ResponseObject struct {
	Name             string `json:"Name"`
	URL              string `json:"Url"`
	HTTPStatus       int16  `json:"HTTP_status"`
	TotalTime        int16  `json:"Total_time"`
	DNSLookup        int16  `json:"Dns_lookup"`
	TCPConnection    int16  `json:"Tcp_connection"`
	TLSHandshake     int16  `json:"Tls_handshake"`
	ServerProcessing int16  `json:"Server_processing"`
	ContentTransfer  int16  `json:"Content_transfare"`
}
