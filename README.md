# img-dl

Images Download Tool written in Go

## Download All Images From Spicific webpage

steps to test the idea :

1. cURL the webpage
2. get the links of all Images
3. download all images into a known directory (locally)

Notes :

- regex for all jpg and png files: `[a-zA-Z0-9/_.:-]+.(jpg|png)`
- get all image files into array, then loop for download them all
- if the file starts with `//` then replace it with `https://` or add `http:` before `//`.

## Download All Images From a Specific website

I will develop it further to download all images from a website url by crawling its webpages.

## Build from source

build the CLI program using this command

```sh
# download dependencies
go mod tidy

# build the project to get the executable app
go build -o imgdl .
```

run the tool using this command

```sh
# run the app to download all images from the URL (webpage)
./imgdl get https://x.com/abanoubha
```

## Installation

```sh
go install github.com/abanoubha/img-dl
```

## Links

Video: [coding a tool to download all images from a webpage url in Go](https://youtu.be/qJ5RlAFk5QI)
