# slide

### present slides created in other programs nicely in a browser

`present` (a Golang tool) is great for creating clean, easy to make slides. But when you want to create well designed slides in another program (because let's be real, sometimes plain text and raw images don't cut it) it can be a hassle to get them running easily in a web browser and looking clean, and even harder to host them on a web server. Enter slide. Always spelled with a lowercase 's', especially at the beginning of sentences, slide helps people present things to the best of their ability. 

#### slide in action

![screenshot](screenshot.png "slide in action")

# installation

#### with `go get`

from your `$GOPATH`, with `$GOPATH/bin` in your `$PATH`:

```bash
$ go get github.com/cdipaolo/slide
$ go install ./...
$ cp bin/slide /usr/bin/slide

# now just use as normal!
```

# usage

place a pdf of your slideshow (created in whatever) as 'slides.pdf' or (with extensions JPG/JPEG/PNG) images of your slides (named *.png, *.jpg, or *.jpeg as long as the files will sort into the correct order) into a directory (lets call it `slide-dir` for now.) Run the following and open `localhost:3000` (it'll prompt you anyways.) 

That's it.

```bash
$ cd slide-dir
$ slide 
```

# license: MIT

see [LICENSE](LICENSE) for details.