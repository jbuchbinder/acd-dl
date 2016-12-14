package main

import (
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"log"
	"net/http"
	"os"
)

func retrieve(shareId string, r Resource, basedir string) {
	os.MkdirAll(basedir, 0755)
	for _, item := range r.Data {
		if item.Kind == "FILE" {
			if *Debug {
				log.Printf("[%s] retrieve %s, kind %s, tempLink %s", basedir, item.Id, item.Kind, item.TempLink)
			} else {
				log.Printf("FILE: %s", basedir+string(os.PathSeparator)+item.Name)
			}
			get(shareId, item, basedir)
		} else if item.Kind == "FOLDER" {
			if *Debug {
				log.Printf("[%s] retrieve %s, kind %s", basedir, item.Id, item.Kind, item.TempLink)
			} else {
				log.Printf("FOLDER: %s", basedir+string(os.PathSeparator)+item.Name)
			}
			f, err := getFolder(shareId, item.Id)
			if err != nil {
				log.Printf("ERR: %s", err.Error())
				continue
			}
			retrieve(shareId, f, basedir+string(os.PathSeparator)+item.Name)
		}
	}
}

func get(shareId string, data ResourceData, basedir string) error {
	bar := pb.New64(data.Properties.Size).SetUnits(pb.U_BYTES)
	bar.Start()
	defer bar.Finish()

	dest := basedir + string(os.PathSeparator) + data.Name

	// Check to see if we have already retrieved the file
	var start int64
	if fileExists(dest) {
		if fileSize(dest) >= data.Properties.Size {
			log.Printf("%s fully retrieved", dest)
			return nil
		}
		start = fileSize(dest)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", data.TempLink, nil)
	if start > 0 {
		bar.Set64(start)
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", start))
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERR: %s", err.Error())
		return err
	}
	defer resp.Body.Close()
	var out *os.File
	if start == 0 {
		out, err = os.Create(dest)
		if err != nil {
			log.Printf("ERR: %s", err.Error())
			return err
		}
	} else {
		out, err = os.OpenFile(dest, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("ERR: %s", err.Error())
			return err
		}
	}
	defer out.Close()
	//io.Copy(out, resp.Body)
	io.Copy(out, bar.NewProxyReader(resp.Body))
	return nil
}
