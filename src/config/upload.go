package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/b3bas/grace"
	gcfg "gopkg.in/gcfg.v1"
)

// ReadConfig is file handler for reading configuration files into variable
// Param: - config pointer of Config, filepath string
// Return: - boolean
func ReadConfig(cfg *Config, path string, module string) bool {
	environ := os.Getenv("B3BAS_ENV")
	if environ == "" {
		environ = ENV_DEVELOPMENT
	}

	environ = strings.ToLower(environ)

	parts := []string{"main"}
	var configString []string

	for _, v := range parts {
		fname := path + "/" + module + "/" + environ + "/" + module + "." + v + ".ini"
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "Reading", fname)

		config, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Println("config/init.go function ReadConfig", err)
			return false
		}

		configString = append(configString, string(config))
	}

	err := gcfg.ReadStringInto(cfg, strings.Join(configString, "\n\n"))

	if err != nil {
		log.Println("config/init.go function ReadConfig", err)
		return false
	}

	return true
}

func GetConfig() *Config {
	CF = &Config{}
	GOPATH := os.Getenv("GOPATH")
	ok := ReadConfig(CF, "/etc", "b3bas-up") || ReadConfig(CF, GOPATH+"/src/github.com/b3bas/b3bas-up/files/etc", "b3bas-up") || ReadConfig(CF, "files/etc", "b3bas-up")
	if !ok {
		log.Fatal("Failed to read config file")
	}

	return CF
}

func (g GraceConfig) ToGraceConfig() grace.Config {
	timeout, err := time.ParseDuration(g.Timeout)
	if err != nil {
		timeout = time.Second * 5
	}

	readTimeout, err := time.ParseDuration(g.HTTPReadTimeout)
	if err != nil {
		timeout = time.Second * 10
	}

	writeTimeout, err := time.ParseDuration(g.HTTPWriteTimeout)
	if err != nil {
		timeout = time.Second * 10
	}

	return grace.Config{
		Timeout:          timeout,
		HTTPReadTimeout:  readTimeout,
		HTTPWriteTimeout: writeTimeout,
	}
}
