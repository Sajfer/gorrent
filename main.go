package gorrent

import (
	"io"
	"log"
	"os"

	"./torrentfile"

	"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `benvode:"name"`
}

type bencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func Open(r io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)
	if err != nil {
		return nil, err
	}
	return &bto, nil
}

func main() {
	torrentPath := os.Args[1]
	downloadPath := os.Args[2]

	torrent, err := torrentfile.Open(torrentPath)
	if err != nil {
		log.Fatal(err)
	}

	err = torrent.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
