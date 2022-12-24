package main

import (
	cgf "mailService/Config"
	rout "mailService/Router"
)

func main() {
	cgf.Init()
	rout.HandleFunc()
}
