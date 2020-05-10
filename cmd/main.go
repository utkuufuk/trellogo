package main

import (
	"log"
	"os"

	"github.com/utkuufuk/trellogo/internal/commands"
	"github.com/utkuufuk/trellogo/internal/config"
	"github.com/utkuufuk/trellogo/pkg/api"
)

func main() {
	cfg, err := config.ReadConfig("config.yml")
	if err != nil {
		log.Fatalf("[-] could not read config variables: %v", err)
	}

	client := api.NewClient(cfg.ApiKey, cfg.ApiToken)
	lists := mapLists(cfg.Lists)

	command := commands.Parse(os.Args[1:])
	if err := command.Run(client, lists); err != nil {
		log.Fatalf("[-] could not run command: %v", err)
	}
}

// mapLists creates a name->ID mapping for the lists specified in config
func mapLists(lists []config.List) map[string]string {
	m := make(map[string]string)
	for _, l := range lists {
		m[l.Name] = l.Id
	}
	return m
}
