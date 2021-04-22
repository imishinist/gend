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

	for _, rule := range conf.Rules {
		target := targetdef.NewTarget(rule.Key)

		length, err := generator.Length(rule.Length)
		if err != nil {
			log.Println(rule.Key, err)
			continue
		}
		for i := 0; i < length; i++ {
			value, err := generator.Value(ctx, rule)
			if err != nil {
				log.Println(rule.Key, err)
				continue
			}
			target.Add(value)
		}
		res, err := generator.Generator(ctx, target, rule)
		if err != nil {
			log.Println(rule.Key, err)
			continue
		}
		fmt.Print(res)
	}
}
