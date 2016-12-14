package main

import (
	"flag"
	"log"
	"os"
)

var (
	Debug = flag.Bool("debug", false, "Enable debugging")
)

func main() {
	flag.Parse()

	for _, id := range flag.Args() {
		sr, err := getLink(id)
		if err != nil {
			log.Printf("Unable to retrieve %s, skipping remainder", id)
			continue
		}
		//log.Printf("%#v\n", sr)

		f, err := getFolder(id, sr.NodeInfo.Id)
		if err != nil {
			log.Printf("Unable to retrieve %s, skipping remainder", id)
			continue
		}
		//log.Printf("%#v\n", f)
		retrieve(id, f, "."+string(os.PathSeparator)+sr.NodeInfo.Name)
	}
}
