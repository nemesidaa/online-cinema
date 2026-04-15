package main

import "flag"

var confFile string

func InitFlags() {
	flag.StringVar(&confFile, "config", "./etc/auth.yaml", "path to config file")
	flag.Parse()
}
