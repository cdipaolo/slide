package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
    Port = ":3000"
)

var tmp string
var dir string

func init() {
	// seed rand
	rand.Seed(time.Now().UTC().UnixNano())

	// get images directory
	flag.StringVar(&dir, "img", "."+string(filepath.Separator), "Set the directory containing either slides in pdf form (as slides.pdf), images – .jpg, .jpeg, or .png – separated as (1.png, 2.png, 3.png, etc.) Defaults to current directory.")

	// set tmp to a tmp/slide directory to serve the static files from
	flag.StringVar(&tmp, "serve", "/tmp/slides"+string(rand.Uint32()), "Assigns the serving directory for the static files. Files will be copies into this directory. Defaults to /tmp/slides{+ some random 32 bit unsigned integer}")

	// parse flags
	flag.Parse()
}

func main() {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(fmt.Sprintf("Could not open directory %v\n\tGot Error %v\n", dir, err))
	}

	fmt.Printf("Files returned: %v", files)

    // handle file server
	fs := http.FileServer(http.Dir(tmp))
    http.Handle("/", fs)
    
    // listen and serve file server
    fmt.Printf("\nListening on port %v...", Port)
}

// Match takes in a slice of os.FileInfo's and reduces it to only those files
// who end with PNG and have a number as the actual file name, JPG/JPEG and
// have a number as the file name, or PDF and have 'slides' as the file name
//
// Any file paths not matched will return a nil array or file pointer
//
// Preference order for return of filtered files is 'slides.pdf', and then
// the other JPEG/JPG and PNG formats as n.{jpg,jpeg,png}
//
// Returns: a *os.FileInfo for the PDF file, an []os.FileInfo for the PNG files,
// and an []os.FileInfo for the JPG/JPEG files.
func Match(files []os.FileInfo) (os.FileInfo, []os.FileInfo, []os.FileInfo) {
	png := []os.FileInfo{}
	jpg := []os.FileInfo{}

	for i, file := range files {
		if !file.Mode().IsRegular() {
			continue
		}

		// split name at '.'
		name := strings.Split(file.Name(), ".")
		if len(name) != 2 {
			// if there is more than 2 'segments' of the file, continue
			// eg. if the file is slide.1.png –> it should be 1.png

			continue
		}

		if file.Name() == "slides.pdf" {
			return file, nil, nil

		} else if name[0] != string(i+1) {
			// if the name doesn't fit n.ext where n is one more than
			// the index, continue

			continue

		} else if name[1] == "png" {
			// if the file matches n.png, where n is one more than the
			// index, then add it to the array of png files

			png = append(png, file)

		} else if name[1] == "jpeg" || name[1] == "jpg" {
			// if the file matches n.jpg or n.jpeg, where n is one
			// more than the index, then add it to the array of png files

			jpg = append(jpg, file)
		}
	}

	return nil, png, jpg
}
