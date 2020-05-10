package commands

import (
	"fmt"

	"github.com/utkuufuk/trellogo/pkg/api"
)

type Command interface {
	Run(client api.Client, lists map[string]string) error
}

// UndefinedCommand returns an error indicating an undefined command.
type UndefinedCommand struct{ Err error }

// FetchAllListsCommand fetches cards in all lists.
type FetchAllListsCommand struct{}

// FetchListsCommand fetches cards in selected lists.
type FetchListsCommand struct {
	Lists []string
}

func (cmd UndefinedCommand) Run(_ api.Client, _ map[string]string) error {
	return cmd.Err
}

// TODO: fetch lists in parallel
func (cmd FetchAllListsCommand) Run(client api.Client, lists map[string]string) error {
	for name, id := range lists {
		if err := fetchAndPrintList(client, name, id); err != nil {
			return err
		}
	}
	return nil
}

// TODO: fetch lists in parallel
func (cmd FetchListsCommand) Run(client api.Client, lists map[string]string) error {
	for _, list := range cmd.Lists {
		id, ok := lists[list]
		if !ok {
			return fmt.Errorf("undefined list: '%s'", list)
		}

		if err := fetchAndPrintList(client, list, id); err != nil {
			return err
		}
	}
	return nil
}

// fetchAndPrintList fetches & prints all cards in the given list
func fetchAndPrintList(client api.Client, name, id string) error {
	data, err := client.FetchListCards(id)
	if err != nil {
		return fmt.Errorf("could not fecth cards in '%s' list: %w", name, err)
	}

	fmt.Printf("\n%s List:\n=====================\n", name)
	for _, card := range data {
		fmt.Println(card.Name)
	}
	return nil
}
