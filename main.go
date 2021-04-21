package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/imishinist/gend/rule/definition"
	"github.com/imishinist/gend/rule/generator"
	targetdef "github.com/imishinist/gend/target/definition"
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

	target := targetdef.New()
	for _, rule := range conf.Rules {
		length, err := generator.Length(rule.Length)
		if err != nil {
			continue
		}
		for i := 0; i < length; i++ {
			value, err := generator.Value(ctx, rule)
			if err != nil {
				continue
			}
			target.Add(rule.Key, value)
		}
	}
	json.NewEncoder(os.Stdout).Encode(target)
}
