package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"telry.io/client/types"
)

func (c *Client) CreateThread(ctx context.Context, recipient string, recipientName string) error {
	body := map[string]any{
		"to": recipient,
		"message": map[string]any{
			"template": "hello_app",
			"template_data": types.TemplateData{
				"body": []types.TemplateDataEntry{
					{
						Name:  "name",
						Value: recipientName,
					}, {
						Name:  "app",
						Value: "Telry Prod",
					}, {
						Name:  "message",
						Value: "Welcome to Telry",
					},
				},
			},
		},
	}

	b, _ := json.Marshal(body)
	url, err := c.prepareUrl("/threads", types.Query{})
	if err != nil {
		return err
	}

	thread, err := post[any](ctx, url, bytes.NewReader(b), newAuthorizationHeader(c.token))
	if err != nil {
		return err
	}

	log.Printf("Thread: %v", thread)
	return nil
}

func (c *Client) GetThreads(ctx context.Context, query types.Query) (types.ThreadsResponse, error) {
	url, err := c.prepareUrl("/threads", query)
	if err != nil {
		return types.ThreadsResponse{}, err
	}
	return get[types.ThreadsResponse](ctx, url, newAuthorizationHeader(c.token))
}

func (c *Client) NewTextMessage(ctx context.Context, threadID string, text string) error {
	url, err := c.prepareUrl("/threads/"+threadID, types.Query{})
	if err != nil {
		return err
	}

	body := map[string]string{
		"text": text,
	}

	b, _ := json.Marshal(body)

	zero, err := post[any](ctx, url, bytes.NewReader(b), newAuthorizationHeader(c.token))
	if err != nil {
		return err
	}

	log.Printf("TextMessageResponse: %v", zero)
	return nil
}

func (c *Client) NewTemplateMessage(ctx context.Context, threadID string, templateName string, data types.TemplateData) error {
	return nil
}

func (c *Client) GetMessages(ctx context.Context, threadID string, query types.Query) (types.MessagesResponse, error) {
	threadPath := fmt.Sprintf("/threads/%s/messages", threadID)
	url, err := c.prepareUrl(threadPath, query)
	if err != nil {
		return types.MessagesResponse{}, err
	}

	return get[types.MessagesResponse](ctx, url, newAuthorizationHeader(c.token))
}

func (c *Client) Welcome(ctx context.Context, name string, phone string) (types.WelcomeMessage, error) {
	url, err := c.prepareUrl("/welcome", types.Query{})
	if err != nil {
		panic(err)
	}
	body := map[string]string{
		"name":  name,
		"phone": phone,
	}

	b, _ := json.Marshal(body)

	return post[types.WelcomeMessage](ctx, url, bytes.NewReader(b), newAuthorizationHeader(c.token))
}

func (c *Client) MarkRead(ctx context.Context, threadID string) error {
	threadPath := fmt.Sprintf("/threads/%s/mark-read", threadID)
	url, err := c.prepareUrl(threadPath, types.Query{})
	if err != nil {
		return nil
	}

	if _, err := do[any](ctx, http.MethodPost, url, nil, newAuthorizationHeader(c.token)); err != nil {
		return err
	}

	return nil
}
