package client

type Client struct {
	url          string
	token        string
	organization string
	debug        bool
}

const (
	defaultBaseURL = "https://api.telry.io/v1"
	libraryName    = "telry-client-go"
	libraryVersion = "0.0.1"
)

type Option func(*Client)

func Debug() Option {
	return func(c *Client) {
		c.debug = true
	}
}

func WithURL(url string) Option {
	return func(c *Client) {
		c.url = url
	}
}

func New(token, organization string, o ...Option) *Client {
	c := &Client{
		url:          defaultBaseURL,
		token:        token,
		organization: organization,
	}

	for _, fn := range o {
		fn(c)
	}

	return c
}
