package main

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
)

func checkFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func writeItemToZipWriter(w *zip.Writer, item string) {
	filepath.Walk(item, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fDest, err := w.Create(path)
			if err != nil {
				log.Printf("Create failed: %s.\n", err.Error())
				return nil
			}

			fSrc, err := os.OpenFile(path, os.O_RDONLY, 0644)
			if err != nil {
				log.Printf("Open failed: %s.\n", err.Error())
				return nil
			}
			defer fSrc.Close()

			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				log.Printf("Copy failed: %s.\n", err.Error())
				return nil
			}
		}
		return nil
	})
}

func generateZip(itemList []string, zipName string) {
	// 创建zip.Writer
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	for _, item := range itemList {
		writeItemToZipWriter(w, item)
	}

	err := w.Close()
	if err != nil {
		info.Fatalln(err)
	}

	if checkFileExist(zipName) {
		err := os.Remove(zipName)
		if err != nil {
			info.Println(err)
			log.Fatal(err)
		}
	}

	f, err := os.OpenFile(zipName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)
	info.Println("finished.")
}
