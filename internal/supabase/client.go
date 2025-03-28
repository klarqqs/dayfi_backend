package supabase

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
)

type Client struct {
	*supabase.Client
}

func NewClient(url, key string) (*Client, error) {
	client, err := supabase.InitClient(url, key, nil)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (c *Client) AddToWaitlist(name, email string) error {
	data := map[string]interface{}{
		"name":  name,
		"email": email,
	}
	_, count, err := c.From("waitlist").Insert(data, false, "", "", "").Execute()
	fmt.Println("Rows affected:", count)
	return err
}
