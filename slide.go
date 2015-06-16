package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
    "os/exec"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

const (
	// permissions for the serving directory
	perm os.FileMode = os.FileMode(int(0777))
)

var noTmp bool        // whether to use a base directory other than /tmp
var tmp string        // temp directory created to serve static files from
var subDirs []string  // dirs to be made for copying data into www/ dir
var dir string        // user directory where the slide images reside
var reducedDir string // tmp minus /www on the end
var port int64        // port to be used

func init() {
	// seed rand
	rand.Seed(time.Now().UTC().UnixNano())

	// set current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Could not get current working directory\n\tGot Error %v\n", err))
	}

	// see if we shouldn't use the /tmp/ directory
	baseDir := os.TempDir()
	flag.BoolVar(&noTmp, "no-tmp", false, "Sets if you don't want to use the /tmp/ directory to host the static files, but you want to have the directory be set . Could be used if you want to host this persistantly on a server, although this isn't the most practical thing because the files are recopied every time. Defaults to false. When true, baseDir set to '/.slide/'")
	if noTmp {
		baseDir = "/.slide/"
	}

	// get port
	flag.Int64Var(&port, "port", 3000, "Sets the port to run the slide application from. Defaults to 3000")

	// get images directory
	flag.StringVar(&dir, "img", cwd, "Set the directory containing either slides in pdf form (as slides.pdf), images – .jpg, .jpeg, or .png – separated as (1.png, 2.png, 3.png, etc.) Defaults to current directory.")

	// set tmp to a tmp/slide directory to serve the static files from
	flag.StringVar(&reducedDir, "serve", baseDir+".slide/"+strconv.FormatInt(int64(rand.Uint32()), 10), "Assigns the serving directory for the static files. Files will be copies into this directory. Defaults to {operating system temp dir}/slides{+ some random 32 bit unsigned integer}")
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

	pdfArr, png, jpg := Match(files)
    pdfLen := len(pdfArr)
	pngLen := len(png)
	jpgLen := len(jpg)

	if pdfLen != 0 {
		// migrate pdf to a group of pngs
        //
        // image-magick Command:
        // $ convert file.pdf image.png
        pdf := pdfArr[0]
        
        pdfFile := filepath.Join(dir, pdf.Name())
        output := filepath.Join(dir + "/slide.png")
        
        command := exec.Command("convert", pdfFile, output)

        execErr := command.Run()
        if execErr != nil {
                panic(fmt.Sprintf("Could not create png's from pdf %v\n\tGot Error <%v>\n\t with output %v\n", pdfFile, execErr, output))
        }
        
        // copy over images again
        files, err = ioutil.ReadDir("/" + dir)
        if err != nil {
            panic(fmt.Sprintf("Could not open directory %v\n\tGot Error %v\n", dir, err))
        }
        fmt.Printf("Files returned: %v\n\n", files)

        _, png, _ = Match(files)
        pngLen = len(png)
	}
    
    if pngLen != 0 {
		// copy png files over to the tmp directory

        for i, file := range png {
            path := filepath.Join(dir, file.Name())
            cpPath := filepath.Join(tmp, "img", strconv.FormatInt(int64(i+1), 10) + ".png")
            command := exec.Command("cp", path, cpPath)

            execErr := command.Run()
            if execErr != nil {
                    panic(fmt.Sprintf("Could not copy image %v\n\tGot Error %v\n", dir, execErr))
            }
            
            fmt.Printf("copied %v\n\t -> %v\n", path, cpPath)
		}
	} else if jpgLen != 0 {
		// copy png files over to the tmp directory

        for i, file := range jpg {
            path := filepath.Join(dir, file.Name())
            cpPath := filepath.Join(tmp, "img", strconv.FormatInt(int64(i), 10) + ".jpg")
          
            command := exec.Command("cp", path, cpPath)

            execErr := command.Run()
            if execErr != nil {
                    panic(fmt.Sprintf("Could not copy image %v\n\tGot Error %v\n", dir, execErr))
            }
            
            fmt.Printf("copied %v into %v\n", path, cpPath)
		}
    } else {
        panic("Slide files not in correct format")
    }

	// handle file server
	fs := http.FileServer(http.Dir(tmp))
	http.Handle("/", fs)

	// listen and serve file server
	portString := ":" + strconv.FormatInt(port, 10)

	fmt.Printf("\nServing %v %v ...\n", tmp, portString)
	log.Fatal(http.ListenAndServe(portString, nil))
}

// Match takes in a slice of os.FileInfo's and reduces it to only those files
// who end with PNG, JPG/JPEG, or PDF
//
// Any file paths not matched will return a nil array or file pointer
//
// Preference order for return of filtered files is 'slides.pdf', and then
// the other JPEG/JPG and PNG formats as n.{jpg,jpeg,png}
//
// Returns: a *os.FileInfo for the PDF file, an []os.FileInfo for the PNG files,
// and an []os.FileInfo for the JPG/JPEG files.
func Match(files []os.FileInfo) ([]os.FileInfo, []os.FileInfo, []os.FileInfo) {
    pdf := []os.FileInfo{}
	png := []os.FileInfo{}
	jpg := []os.FileInfo{}

	for _, file := range files {
		if !file.Mode().IsRegular() {
			continue
		}

		// split name at '.'
		ext := filepath.Ext(file.Name())

		if ext == ".pdf" {
            // if file matches *.pdf
            
            pdf = append(pdf, file)

		} else if ext == ".png" {
			// if the file matches *.png

			png = append(png, file)

		} else if ext == ".jpeg" || ext == ".jpg" {
			// if the file matches *.jpg or *.jpeg

			jpg = append(jpg, file)
		}
	}

	return pdf, png, jpg
}
