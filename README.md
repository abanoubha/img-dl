# img-dl
Images Download Tool written in Go

# Download All Images From Spicific webpage

steps to test the idea :
1. cURL the webpage
2. get the links of all Images
3. download all images into a known directory (locally)

Notes :
- regex for all jpg and png files: `[a-zA-Z0-9/_.:-]+.(jpg|png)`
- get all image files into array, then loop for download them all
- if the file starts with `//` then replace it with `https://` or add `http:` before `//`.

# Download All Images From a Specific website
I will develop it further to download all images from a website url by crawling its webpages.

# Installation
build the CLI program using this command
```Go
go build main.go
```
run the tool using this command
```Go
./main get https://twitter.com/devabanoub
```

## Links
Video: [coding a tool to download all images from a webpage url in Go](https://youtu.be/qJ5RlAFk5QI)
