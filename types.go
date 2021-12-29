package main

type content struct {
	Files       []file      `json:"files" xml:"files"`
	Directories []directory `json:"directories" xml:"directories"`
}

type file struct {
	Name string `json:"name" xml:"name"`
	Size int    `json:"size" xml:"size"`
	Mode string `json:"mode" xml:"mode"`
	Link string `json:"link"  xml:"link"`
}

type directory struct {
	Name string `json:"name" xml:"name"`
	Mode string `json:"mode" xml:"mode"`
	Link string `json:"link" xml:"link"`
}

type httpError struct {
	Status  int    `json:"status" xml:"status"`
	Message string `json:"message" xml:"message"`
}
