package main

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
)

func getFolder(shareId, id string) (Resource, error) {
	url := fmt.Sprintf("https://www.amazon.com/drive/v1/nodes/%s/children?resourceVersion=V2&ContentType=JSON&limit=200&sort=%%5B%%22kind+DESC%%22%%2C+%%22modifiedDate+DESC%%22%%5D&asset=ALL&tempLink=true&shareId=%s", id, shareId)
	if *Debug {
		log.Printf("GET %s", url)
	}
	_, body, errs := gorequest.New().Get(url).EndBytes()
	if len(errs) > 0 {
		return Resource{}, errs[0]
	}
	if *Debug {
		log.Printf("FOLDER %s : %d bytes read", id, len(body))
	}

	var data Resource
	err := json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
