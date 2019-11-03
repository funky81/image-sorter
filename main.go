package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var outputDir string

func main() {
	//outputDir = string("d:\\temp2")
	outputDir = os.Args[2]
	//err := filepath.Walk("d:\\temp1\\", fileWalk)
	err := filepath.Walk(os.Args[1], fileWalk)
	if err != nil {
		log.Println(err)
	}
}

func fileWalk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() == false {
		//fmt.Println(path, info.Size())
		time := info.ModTime()
		newDirectory := outputDir + string(os.PathSeparator) + time.Format("2006-01-02") + string(os.PathSeparator)
		if _, err := os.Stat(newDirectory); os.IsNotExist(err) {
			os.MkdirAll(newDirectory, os.FileMode(0522))
		}
		newFile := newDirectory + info.Name()
		fmt.Println("Move to ", newFile)
		err := copyFile(path, newFile)
		if err != nil {
			fmt.Errorf("Error %s", err)
		}
	}
	return nil
}

func copyFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	return nil
}

func moveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}
