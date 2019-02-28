package main

type Tcpd_Rest struct {
	Url 		string 			`json:"url"`
	Id			string 			`json:"id"`
	Client		string 			`json:"client"`
	Context 	string 			`json:"context"`
	Host 		string 			`json:"host"`
	Needheaders []string 		`json:"needheaders"`
	Needbody 	string 			`json:"needbody"`
	Body 		string 			`json:"body"`
}

type Rest_Tcpd struct {
	Ret 		string 			`json:"ret"`
	Id 			string 			`json:"id"`
	Context 	string 			`json:"context"`
	Body 		string 			`json:"body"`
	Headers 	[]string 		`json:"headers"`
}