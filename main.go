package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type MemConfig struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type JarConfig struct {
	Provider string    `json:"provider"`
	Version  string    `json:"version"`
	Memory   MemConfig `json:"memory"`
}

type PluginConfig string

type Settings struct {
	JVM        map[string]string `json:"jvm"`
	Properties map[string]string `json:"properties"`
}

type Config struct {
	Name     string         `json:"name"`
	Version  string         `json:"version"`
	Jar      JarConfig      `json:"jar"`
	Plugins  []PluginConfig `json:"plugins"`
	Settings Settings       `json:"settings"`
}

func main() {
	configFile := flag.String("config", "mcli.yml", "Config file path")

	flag.Parse()

	f, err := os.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config

	if err := json.Unmarshal(f, &cfg); err != nil {
		log.Fatal(err)
	}

	// TODO: download server jar, paper eg:
	// https://api.papermc.io/v2/projects/paper/versions/{minecraft_version}/builds/{build_version}/downloads/paper-{minecraft_version}-{build_version}.jar

	// NOTE: Need to evaluate different ways of caching locally.

	fmt.Printf("%+v\n", cfg)
}
