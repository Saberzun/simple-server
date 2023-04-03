package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Port          int    `yaml:"port"`
	RestartScript string `yaml:"restartScript"`
}

func main() {
	cfg, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config file: %s", err)
	}

	http.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("start to restart stable diffusion webui...")
		if err := exec.Command("/bin/sh", cfg.RestartScript).Run(); err != nil {
			log.Fatalf("Failed to run restart script: %s", err)
		}
		fmt.Fprintf(w, "Restart success! %s /n", time.Now())
		fmt.Printf("restart stable diffusion webui success")
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}

func loadConfig(path string) (*Config, error) {
	cfg := new(Config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
