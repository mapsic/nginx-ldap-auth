package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var (
	configFile = flag.String("config", "/etc/nginx-ldap-auth/config.yaml", "Configuration file")
	config     Config
)

func main() {
	flag.Parse()

	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Could not read file \"%s\": %v\n", *configFile, err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error on parse config: %v\n", err)
	}

	fmt.Printf("Config: %+v\n", config)
}