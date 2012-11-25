package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	PREFIX="/webshare/"
)

func currentDir() string {
	dir,err:=os.Getwd()
	if err!=nil {
		log.Fatal("Can't get working directory!")
	}
	return dir
}

func main() {
	// Load command line config
	var prefix,dir string
	var port int 
	flag.StringVar(&prefix,"prefix", PREFIX, "Web prefix before accesing the shared files")
	flag.StringVar(&dir,"dir",currentDir(),"Local directory to be shared on the web")
	flag.IntVar(&port,"port",80,"Port to share the files from HTTP")
	flag.Parse()
	
	// Advertise configuration loaded
	fmt.Printf("WebShare directory %s from :%v%s...\n",dir,port,prefix)

	// Serve files...
	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(dir))))
	err := http.ListenAndServe(fmt.Sprintf(":%v",port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}