package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"telry.io/client/types"
)

func (c *Client) CreateOTP(ctx context.Context, phone string) (types.OtpResponse, error) {
	url := fmt.Sprintf("%s/organizations/%s/otp", c.url, c.organization)
	b, _ := json.Marshal(types.OtpRequest{
		Phone: phone,
	})

	return post[types.OtpResponse](ctx, url, bytes.NewReader(b), newAuthorizationHeader(c.token))
}

func (c *Client) VerifyOTP(ctx context.Context, phone string, code string) (types.OtpResponse, error) {
	url := fmt.Sprintf("%s/organizations/%s/otp/verify", c.url, c.organization)
	b, _ := json.Marshal(types.OtpRequest{
		Phone: phone,
		Code:  code,
	})

	return post[types.OtpResponse](ctx, url, bytes.NewReader(b), newAuthorizationHeader(c.token))
}

func (c *Client) GetOrganization(ctx context.Context) (types.Organization, error) {
	url := fmt.Sprintf("%s/organizations/%s", c.url, c.organization)
	return get[types.Organization](ctx, url, newAuthorizationHeader(c.token))
}
