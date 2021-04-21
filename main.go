package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/imishinist/gend/rule/definition"
)

func main() {
	filename := flag.String("conf", "config.yml", "")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	input, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var conf definition.Config
	if err := yaml.Unmarshal(input, &conf); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(os.Stdout).Encode(conf)
}
