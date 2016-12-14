package main

import (
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

	resp, err := http.Get(data.TempLink)
	if err != nil {
		log.Printf("ERR: %s", err.Error())
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(basedir + string(os.PathSeparator) + data.Name)
	if err != nil {
		log.Printf("ERR: %s", err.Error())
		return err
	}
	defer out.Close()
	//io.Copy(out, resp.Body)
	io.Copy(out, bar.NewProxyReader(resp.Body))
	return nil
}
