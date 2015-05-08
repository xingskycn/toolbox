package main

import (
	"log"
    "flag"
	"net/http"
	"fmt"
)

func Usage(){
	var usage string = `A Simple HTTP Server
Usage:
	-h		Help.
	-l		Listen address. Default: 0.0.0.0:80.
	-d		Absolute path of Root directory. Default: ".".
`
	fmt.Print(usage)
}

var h bool
var l string
var d string

func main() {
    flag.BoolVar(&h, "h", false, "Usage.")
	flag.StringVar(&l, "l", "0.0.0.0:80", "Listen address.")
	flag.StringVar(&d, "d", ".", "Root directory.")
    flag.Parse()
    if h {
        Usage()
        return
    }
	log.Fatal(http.ListenAndServe(l, http.FileServer(http.Dir(d))))
}
