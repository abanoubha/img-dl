# img-dl
Images Download Tool written in Go

# Download All Images From Spicific webpage
- cURL the webpage
- get the links of all Images
- download all images into a known directory (locally)

LET'S TEST THE IDEA MANUALLY

regex for all jpg and png files: [a-zA-Z0-9/_.:-]+.(jpg|png)

get all image files into array, then loop for download 'em all
_note_: if the file starts with // then replace it with https://


LET'S IMPLEMET THIS IDEA IN GO
! D O N E !

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
