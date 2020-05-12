package commands

import (
	"fmt"

	"github.com/adlio/trello"
)

type Command interface {
	Run(client *trello.Client, lists map[string]string) error
}

// UndefinedCommand returns an error indicating an undefined command.
type UndefinedCommand struct{ Err error }

// FetchAllListsCommand fetches cards in all lists.
type FetchAllListsCommand struct{}

// FetchListsCommand fetches cards in selected lists.
type FetchListsCommand struct {
	Lists []string
}

func (cmd UndefinedCommand) Run(_ *trello.Client, _ map[string]string) error {
	return cmd.Err
}

// TODO: fetch lists in parallel
func (cmd FetchAllListsCommand) Run(client *trello.Client, lists map[string]string) error {
	for name, id := range lists {
		if err := fetchAndPrintList(client, name, id); err != nil {
			return err
		}
	}
	return nil
}

// TODO: fetch lists in parallel
func (cmd FetchListsCommand) Run(client *trello.Client, lists map[string]string) error {
	for _, name := range cmd.Lists {
		id, ok := lists[name]
		if !ok {
			return fmt.Errorf("undefined list: '%s'", name)
		}

		if err := fetchAndPrintList(client, name, id); err != nil {
			return err
		}
	}
	return nil
}

// fetchAndPrintList fetches & prints all cards in the given list
func fetchAndPrintList(client *trello.Client, name, id string) error {
	list, err := client.GetList(id, trello.Defaults())
	if err != nil {
		return fmt.Errorf("could not fetch fetch list '%s': %w", name, err)
	}

	cards, err := list.GetCards(trello.Defaults())
	if err != nil {
		return fmt.Errorf("could not fetch cards from list '%s': %w", name, err)
	}

	fmt.Printf("\n%s List:\n=====================\n", name)
	for _, card := range cards {
		fmt.Println(card.Name)
	}
	return nil
}
