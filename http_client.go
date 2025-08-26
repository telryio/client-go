package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"telry.io/client/types"
)

func do[T any](ctx context.Context, method string, url string, body io.Reader, headers map[string]string) (T, error) {
	var zero T
	req, _ := http.NewRequestWithContext(ctx, method, url, body)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return zero, errors.New(string(bodyBytes))
	}

	if err := json.NewDecoder(resp.Body).Decode(&zero); err != nil {
		return zero, err
	}

	return zero, nil
}

func post[T any](ctx context.Context, url string, body io.Reader, headers map[string]string) (T, error) {
	val, err := do[T](ctx, http.MethodPost, url, body, headers)
	return val, err
}

func get[T any](ctx context.Context, url string, headers map[string]string) (T, error) {
	return do[T](ctx, http.MethodGet, url, nil, headers)
}

func newAuthorizationHeader(token string) map[string]string {
	return map[string]string{
		"Authorization": "Bearer " + token,
	}
}

func (c *Client) prepareUrl(path string, query types.Query) (string, error) {
	turl := fmt.Sprintf("%s/organizations/%s/%s", c.url, c.organization, strings.Trim(path, "/"))
	u, err := url.Parse(turl)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	if query.Direction != "" {
		q.Add("direction", query.Direction)
	}

	if query.Limit != 0 {
		q.Add("limit", fmt.Sprintf("%d", query.Limit))
	}

	if query.Offset != 0 {
		q.Add("offset", fmt.Sprintf("%d", query.Offset))
	}

	if query.OrderBy != "" {
		q.Add("order_by", query.OrderBy)
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}
