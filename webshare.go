/*
Copyright (c) 2012, <Jose Luis Vázquez González> josvazg@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and 
associated documentation files (the "Software"), to deal in the Software without restriction, 
including without limitation the rights to use, copy, modify, merge, publish, distribute, 
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software 
is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or 
substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING 
BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND 
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, 
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	PREFIX="/webshare/"
)

// currentDir returns the current directory or makes the program fail
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
	flag.StringVar(&prefix,"prefix", "", "Web prefix before accesing the shared files")
	flag.StringVar(&dir,"dir",currentDir(),"Local directory to be shared on the web")
	flag.IntVar(&port,"port",80,"Port to share the files from HTTP")
	flag.Parse()

	// Check directory
	dir,err:=filepath.Abs(dir)
	if err != nil {
		log.Fatal("Can't open: ", err)
	}

	// prefix defaults to the last name of the shared directory
	if prefix=="" {
		if dir=="/" {
			prefix=PREFIX
		} else {
			prefix="/"+filepath.Base(dir)+"/"
		}
	}
	
	// Advertise configuration loaded
	fmt.Printf("WebShare directory %s from :%v%s...\n",dir,port,prefix)

	// Serve files...
	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(dir))))
	err = http.ListenAndServe(fmt.Sprintf(":%v",port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}