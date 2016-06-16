package gaufre;

import (
	"github.com/soywod/file64"
	"log"
	"github.com/soywod/archive"
	"os"
)

var router *Router
var Config struct {
	ArchiveCode string
	ArchivePath string
}

func init() {
	Config.ArchiveCode = ""
	Config.ArchivePath = ""

	router = &Router{Routes: make(map[string]Route)}
}

func Bootstrap() {
	if Config.ArchiveCode != "" {
		defer os.Remove("archive.zip")

		if err := file64.Decode(Config.ArchiveCode, "archive.zip"); err != nil {
			log.Fatal(err)
		}

		if err := archive.Decompress("archive.zip", Config.ArchivePath); err != nil {
			log.Fatal(err)
		}
	}
}
