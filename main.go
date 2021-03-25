package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/n3k0fi5t/ShortTheURL/app/route"
	"github.com/n3k0fi5t/ShortTheURL/common/cache"
	"github.com/n3k0fi5t/ShortTheURL/common/config"
	"github.com/n3k0fi5t/ShortTheURL/common/dbx"
)

var flagConfig = flag.String("config", "./config/test.yml", "path to the config file")

func main() {
	flag.Parse()

	cfg, err := config.Load(*flagConfig)
	if err != nil {
		fmt.Printf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	db, err := dbx.Init(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer db.Close()

	rdb, err := cache.Init(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := route.BuildRouter(cfg, db, rdb)
	r.Run()
}
