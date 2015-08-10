package flowpush

import (
	"log"
	"sync"
)

type Client struct {
	apiKey         string
	organizations  []Organization
	availableFlows []Flow // TODO: Change to map[ID]Flow
	users          map[string]User
}

func (c *Client) DetailsForUser(id string) User {
	return c.users[id]
}

func (c *Client) DetailsForFlow(id string) *Flow {
	for _, flow := range c.availableFlows {
		if flow.ID == id {
			return &flow
		}
	}

	return nil
}

func NewClient(apiKey string) *Client {
	client := &Client{apiKey: apiKey}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		var err error
		client.users, err = getUsers(apiKey)
		if err != nil {
			log.Printf("Failed to get users: %v", err)
		}
		wg.Done()
	}()

	go func() {
		var err error
		client.availableFlows, err = getFlows(apiKey)
		if err != nil {
			log.Printf("Failed to get flows: %v", err)
		}
		wg.Done()
	}()

	go func() {
		var err error
		client.organizations, err = getOrganizations(apiKey)
		if err != nil {
			log.Printf("Failed to get organizations: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()

	return client
}
