package main

import (
	"bytes"
	"io"
	"os"

	"github.com/dimiro1/banner"
	_ "github.com/dimiro1/banner/autoload"
	colorable "github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

const B3BAS_LOG = "/var/log/b3bas-up/b3bas-up.log"

func InitLogo() {
	isEnabled := true
	isColorEnabled := true
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString(" B3BAS Golang Framework {{ .AnsiColor.Green }}(Running){{ .AnsiColor.Default }}\n"))
}

func SaveLog() {
	// create the logger
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	fName, err := os.OpenFile(B3BAS_LOG, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}
	defer fName.Close()
	// multiwriter simultaneously
	logger.SetOutput(io.MultiWriter(os.Stdout, fName))
	println(" Log file saved to:", B3BAS_LOG)
}

func main() {
	InitLogo()
	SaveLog()
}
