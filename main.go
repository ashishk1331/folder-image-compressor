package main

import (
	"fmt"
	"github.com/chai2010/webp"
	"image/jpeg"
	"os"
	"strings"
)

func ConvertImage(filename string) {
	inFile, _ := os.Open(fmt.Sprintf("%s.jpg", filename))
	defer inFile.Close()

	img, _ := jpeg.Decode(inFile)

	outFile, _ := os.Create(fmt.Sprintf("%s.webp", filename))
	defer outFile.Close()

	err := webp.Encode(outFile, img, &webp.Options{Quality: 80})
	if err != nil {
		panic(err)
	}
}

func ListAllJPGs() []string {
	list := []string{}
	files, err := os.ReadDir(".")

	if err != nil {
		fmt.Println("Unable to read the directory.")
	}

	for _, file := range files {
		name := file.Name()

		if strings.HasSuffix(name, ".jpg") {
			before, found := strings.CutSuffix(name, ".jpg")
			if found {
				list = append(list, before)
			}
		}
	}

	return list
}

func PrintFiles(files []string) {
	for _, file := range files {
		fmt.Println(file)
	}
}

func main() {
	list := ListAllJPGs()
	if len(list) > 0 {
		for _, file := range list {
			ConvertImage(file)
			fmt.Printf("✅ %s converted.\n", file)
		}
	} else {
		fmt.Println("No images found.")
	}
}
