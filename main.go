package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var (
	images = []string{".jpeg", ".jpg", ".png", ".svg", ".gif"}
	reads  = []string{".docx", ".pdf", ".md", ".txt", ".pptx"}
	sheets = []string{".xlsx", ".csv"}
	media  = []string{".mp3", ".mp4"}
)

func main() {
	t := time.Now()
	d, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, file := range d {
		fileName := file.Name()
		lowerCaseFileName := strings.ToLower(file.Name())
		switch {
		case containsAny(lowerCaseFileName, sheets):
			if err := os.MkdirAll("SHEETS", os.ModeAppend); err != nil {
				fmt.Println(err)
			}
			dst := fmt.Sprintf("./SHEETS/%s", fileName)
			if err := moveFile(fileName, dst); err != nil {
				fmt.Println(err)
			}
		case containsAny(lowerCaseFileName, reads):
			if err := os.MkdirAll("READINGS", os.ModeAppend); err != nil {
				fmt.Println(err)
			}
			dst := fmt.Sprintf("./READINGS/%s", fileName)
			if err := moveFile(fileName, dst); err != nil {
				fmt.Println(err)
			}
		case containsAny(lowerCaseFileName, images):
			if err := os.MkdirAll("IMAGES", os.ModeAppend); err != nil {
				fmt.Println(err)
			}
			dst := fmt.Sprintf("./IMAGES/%s", file.Name())
			if err := moveFile(file.Name(), dst); err != nil {
				fmt.Println(err)
			}
		case containsAny(lowerCaseFileName, media):
			if err := os.MkdirAll("MEDIA", os.ModeAppend); err != nil {
				fmt.Println(err)
			}
			dst := fmt.Sprintf("./MEDIA/%s", file.Name())
			if err := moveFile(file.Name(), dst); err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Undefined File")
		}
	}

	duration := time.Since(t)
	fmt.Println(duration)
	time.Sleep(3 * time.Second)
}

func moveFile(source, dst string) error {
	inputFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("couldn't open dest file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("couldn't copy to dest from source: %v", err)
	}

	inputFile.Close() // for Windows, close before trying to remove

	err = os.Remove(source)
	if err != nil {
		return fmt.Errorf("couldn't remove source file: %v", err)
	}
	return nil
}

func containsAny(str string, substr []string) bool {
	for _, v := range substr {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}
