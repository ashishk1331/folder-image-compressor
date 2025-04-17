package main

import (
	"fmt"
	"github.com/chai2010/webp"
	"image/jpeg"
	"os"
	"strings"
)

func getImageSize(filename string) int64 {
	info, err := os.Stat(filename)

	if err != nil {
		fmt.Println("Unable to read file.")
	}

	return info.Size()
}

func ConvertImage(filename string) int64 {
	outFileName := fmt.Sprintf("%s.webp", filename)

	inFile, _ := os.Open(fmt.Sprintf("%s.jpg", filename))
	defer inFile.Close()

	img, _ := jpeg.Decode(inFile)

	outFile, _ := os.Create(outFileName)
	defer outFile.Close()

	err := webp.Encode(outFile, img, &webp.Options{Quality: 80})
	if err != nil {
		panic(err)
	}

	return getImageSize(outFileName)
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
			oldSize := float64(getImageSize(fmt.Sprintf("%s.jpg", file)))
			newSize := float64(ConvertImage(file))
			percentReduction := -((oldSize - newSize) / oldSize * 100)
			fmt.Printf("✅ %.2f%% %s converted.\n", percentReduction, file)
		}
	} else {
		fmt.Println("No images found.")
	}
}
