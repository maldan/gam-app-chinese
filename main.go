package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-chinese/internal/app/chinese"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
