package main

import (
	"bytes"
	"context"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/imishinist/gend/rule/definition"
	"github.com/imishinist/gend/rule/generator"
	targetdef "github.com/imishinist/gend/target/definition"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func do(ctx context.Context, gtx *generator.Context, config *definition.Config) {
	items := make([]string, 0, len(config.Rules))
	for _, rule := range config.Rules {
		target := targetdef.NewTarget(rule.Key)

		length, err := generator.Length(rule.Length)
		if err != nil {
			log.Println(rule.Key, err)
			continue
		}
		for i := 0; i < length; i++ {
			res, err := generator.Value(ctx, gtx, rule)
			if err != nil {
				log.Println(rule.Key, err)
				continue
			}
			target.Add(res)
		}
		buf := new(bytes.Buffer)
		if err := generator.KVGenerator(ctx, gtx, target, rule, buf); err != nil {
			log.Println(rule.Key, err)
			continue
		}
		items = append(items, buf.String())
	}
	if err := generator.Generator(ctx, gtx, items, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

// iff rule has same key, simply override it
func mergeConfigRule(conf ...*definition.Config) *definition.Config {
	rules := make([]definition.Rule, 0, 10)
	var gen definition.Generator
	for _, c := range conf {
		if c.Generator != nil {
			gen = *c.Generator
		}
		rules = append(rules, c.Rules...)
	}
	return &definition.Config{Rules: rules, Generator: &gen}
}

func readConfig(filename string) (*definition.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	input, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var conf definition.Config
	if err := yaml.Unmarshal(input, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func parseVars(varsStr string) []generator.Variable {
	ret := make([]generator.Variable, 0)
	vs := strings.Split(varsStr, ",")
	for _, v := range vs {
		trimed := strings.TrimSpace(v)
		tmp := strings.SplitN(trimed, "=", 2)
		if len(tmp) < 2 {
			continue
		}
		name := tmp[0]
		value := tmp[1]
		ret = append(ret, generator.Variable{
			Name:  name,
			Value: value,
		})
	}
	return ret
}

func main() {
	num := flag.Int("n", 1, "")
	varsStr := flag.String("vars", "", "")
	flag.Parse()

	confs := make([]*definition.Config, 0, flag.NArg())
	for _, filename := range flag.Args() {
		conf, err := readConfig(filename)
		if err != nil {
			log.Fatal(err)
		}
		confs = append(confs, conf)
	}

	conf := mergeConfigRule(confs...)
	vars := parseVars(*varsStr)

	ctx := context.Background()
	gtx, err := generator.Build(*conf, vars)
	if err != nil {
		log.Fatal(err)
	}
	defer gtx.Close()

	for i := 0; i < *num; i++ {
		do(ctx, gtx, conf)
	}
}
