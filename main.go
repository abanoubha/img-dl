package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	version bool
	get     string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "img-dl",
		Short: "Download all images from a webpage",
		Long:  "Download all images from a webpage",
		Example: `img-dl -v # show version
img-dl -g x.com/abanoubha # download all images shown in the webpage`,
	}

	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "show the release version of img-dl")

	rootCmd.Flags().StringVarP(&get, "get", "g", "", "download all images shown in the specified webpage")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if get != "" {
			getFileFromUrl(get)
		} else if version {
			fmt.Println(`
img-dl v0.2.0

Software Developer  : Abanoub Hanna
Source code         : https://github.com/abanoubha/gobrew
X Platform          : https://x.com/@AbanoubHA
Developer's Website : https://AbanoubHanna.com`)
		} else {
			fmt.Println(`
You need to specify add a flag.

Example:
  img-dl -v # show version
  img-dl -g x.com/abanoubha # download all images shown in the webpage

img-dl v0.2.0

Software Developer  : Abanoub Hanna
Source code         : https://github.com/abanoubha/gobrew
X Platform          : https://x.com/@AbanoubHA
Developer's Website : https://AbanoubHanna.com`)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getFileFromUrl(webpage string) {
	// download the webpage itself
	resp, err := http.Get(webpage)
	if err != nil {
		fmt.Println("No Internet Connection!")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("no webpage body!!")
	}

	r, _ := regexp.Compile("[a-zA-Z0-9/_.:-]+.(jpg|jpeg|png|webp|avif)")
	// cmd.Println(r.FindAllString(string(body), -1))
	arr := r.FindAllString(string(body), -1)
	for i := 0; i < len(arr); i++ {
		if strings.HasPrefix(arr[i], "//") {
			arr[i] = "https:" + arr[i]
		}
		if strings.HasPrefix(arr[i], "/") {
			arr[i] = webpage + arr[i]
		}
		fileName := filepath.Base(arr[i])
		if err := downloadFile(fileName, arr[i]); err != nil {
			panic(err)
		}
		fmt.Println("downloaded image : ", arr[i])
	}
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create("saved_images/" + filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
