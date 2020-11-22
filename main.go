package main

import (
	"log"
	"os"

	"github.com/sajfer/gorrent/torrentfile"
)

func main() {
	torrentPath := os.Args[1]
	downloadPath := os.Args[2]

	torrent, err := torrentfile.Open(torrentPath)
	if err != nil {
		log.Fatal(err)
	}

	err = torrent.DownloadToFile(downloadPath)
	if err != nil {
		log.Fatal(err)
	}
}
