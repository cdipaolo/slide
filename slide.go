package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
    perm os.FileMode = os.FileMode(int(0777))
)

var noTmp bool
var noCopy bool
var tmp string
var subDirs []string
var dir string
var reducedDir string // tmp minus /www on the end
var port int64

func init() {
	// seed rand
	rand.Seed(time.Now().UTC().UnixNano())

    // set current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Could not get current working directory\n\tGot Error %v\n", err))
	}
    
    // see if we shouldn't use the /tmp/ directory
    baseDir := "/tmp/"
    flag.BoolVar(&noTmp, "no-tmp", false, "Sets if you don't want to use the /tmp/ directory to host the static files, but you want to have the directory be set . Could be used if you want to host this persistantly on a server, although the directory is recopied over every time unless you set '-no-copy'. Defaults to false. When true, baseDir set to '/.slide/'")
    if noTmp {
        baseDir = "/.slide/"
    }

	// get port
	flag.Int64Var(&port, "port", 3000, "Sets the port to run the slide application from. Defaults to 3000")

	// get images directory
	flag.StringVar(&dir, "img", cwd, "Set the directory containing either slides in pdf form (as slides.pdf), images – .jpg, .jpeg, or .png – separated as (1.png, 2.png, 3.png, etc.) Defaults to current directory.")

	// set tmp to a tmp/slide directory to serve the static files from
	flag.StringVar(&reducedDir, "serve", baseDir+".slide/"+strconv.FormatInt(int64(rand.Uint32()), 10), "Assigns the serving directory for the static files. Files will be copies into this directory. Defaults to /tmp/slides{+ some random 32 bit unsigned integer}")
    tmp = reducedDir + "/www"

	// parse flags
	flag.Parse()

	// expand dir if applicable
	if !path.IsAbs(dir) {
		dir = filepath.Join(cwd, dir)
	}
    
    // set all necessary sub directories to be copied into
    // later when making tmp directory structure
    sub := []string{"img"}
    
    for _, subDir := range sub {
        subDirs = append(subDirs, filepath.Join(tmp, subDir))
    }
}

func main() {
    // copy files over to the new serving directory
    fmt.Printf("Creating Temp Directories –\n")
    for _, directory := range subDirs {
        err := os.MkdirAll(directory, perm)
        if err != nil {
            panic(fmt.Sprintf("Could not create directory %v\n\tGot Error %v\n", directory, err))
        }
        
        fmt.Printf("created %v\n", directory)        
    }
    
    err := RestoreAssets(reducedDir, "www")
    if err != nil {
        panic(fmt.Sprintf("Could not restore assets in 'www'\n\tGot Error %v\n", err))
    }

    fmt.Printf("copied 'www/*' into %v\n", reducedDir)   

    
    // copy over images 
    files, err := ioutil.ReadDir("/" + dir)
	if err != nil {
		panic(fmt.Sprintf("Could not open directory %v\n\tGot Error %v\n", dir, err))
	}
	fmt.Printf("Files returned: %v\n\n", files)
    

	// handle file server
	fs := http.FileServer(http.Dir(tmp))
	http.Handle("/", fs)

	// listen and serve file server
	portString := ":" + strconv.FormatInt(port, 10)

    fmt.Printf("\nServing %v %v ...", tmp, portString)
	log.Fatal(http.ListenAndServe(portString, nil))
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
