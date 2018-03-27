package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var version = "undefined"

func main() {
	var config = &Config{}
	flag.IntVar(&config.Port, "port", 9999, "HTTP server port")
	flag.StringVar(&config.MongoUri, "uri", "mongodb://mwe_mongo1:27017,mwe_mongo2:27017,mwe_mongo3:27017", "Mongos URI")
	appVersion := flag.Bool("v", false, "prints version")
	flag.Parse()

	if *appVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	repo, err := NewRepository(config.MongoUri)
	if err != nil {
		logrus.Fatal(err)
	}

	server := &HttpServer{
		Port:       config.Port,
		Repository: repo,
	}

	logrus.Infof("Starting HTTP server on port %v", config.Port)
	server.Start()
}
