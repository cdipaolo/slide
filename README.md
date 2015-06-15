# slide

### present slides created in other programs nicely in a browser

`present` (a Golang tool) is great for creating clean, easy to make slides. But when you want to create well designed slides in another program (because let's be real, sometimes plain text and raw images don't cut it) it can be a hassle to get them running easily in a web browser and looking clean, and even harder to host them on a web server. Enter slide. Always spelled with a lowercase 's', especially at the beginning of sentences, slide helps people present things to the best of their ability. 

![screenshot](screenshot.png "slide in action")

# usage

place a pdf of your slideshow (created in whatever) as 'slides.pdf' or (with extensions JPG/JPEG/PNG) images of your slides (named 1.png, 2.png, 3.png, etc.) into a directory (lets call it `slide-dir` for now.) Run the following and open `localhost:3000` (it'll prompt you anyways.) That's it.

```bash
$ cd slide-dir
$ slide 
```

# license: MIT

see [LICENSE](LICENSE) for details.