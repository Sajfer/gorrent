package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to gorrent!")
	fmt.Println("Endpoint hit: homepage")
}

func returnTorrents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listing all torrents!")
	fmt.Println("Endpoint hit: all")
}

func addTorrent(w http.ResponseWriter, r *http.Request) {

	err := os.Mkdir("tmp", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := ioutil.TempFile("tmp", "*.torrent:")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
	if err != nil {
		return
	}

}

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/all", returnTorrents)
	router.HandleFunc("/add", addTorrent).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", router))
}
