package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/imishinist/gend/rule/definition"
	"github.com/imishinist/gend/rule/generator"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
	ctx := context.Background()

	for _, rule := range conf.Rules {
		length, err := generator.Length(rule.Length)
		if err != nil {
			continue
		}
		value, err := generator.Value(ctx, rule.Value)
		if err != nil {
			continue
		}
		fmt.Println(length, value)
	}
}
