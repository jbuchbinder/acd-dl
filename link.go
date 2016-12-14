package main

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
)

func getLink(id string) (SharedCollection, error) {
	url := fmt.Sprintf("https://www.amazon.com/drive/v1/shares/%s/?resourceVersion=V2&ContentType=JSON&asset=ALL", id)
	if *Debug {
		log.Printf("GET %s", url)
	}
	_, body, errs := gorequest.New().Get(url).EndBytes()
	if len(errs) > 0 {
		return SharedCollection{}, errs[0]
	}
	if *Debug {
		log.Printf("LINK %s : %d bytes read", id, len(body))
	}

	var data SharedCollection
	err := json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
