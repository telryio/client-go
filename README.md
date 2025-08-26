# Telry Go Client (Experimental)

Go SDK for the [Telry API](https:/pwd/api.telry.io).

Status: early scaffold (0.0.0-dev) – surface & APIs may change.

## Install

```
go get telry.io/client@latest
```

## Quick Start

```go
import (
	"context"
	telry "telry.io/client"
)

func example() error {
	c, err := telry.New(
		telry.WithAPIKey("your_api_key"),
		// telry.WithToken("bearer-token"),
	)
	if err != nil { return err }
	user, err := c.Users.Get(context.Background(), "usr_123")
	if err != nil { return err }
	_ = user
	return nil
}
```

## Features (current)

- Configurable client with functional options
- Automatic API key / bearer auth headers
- Middleware chain (auth, retry placeholder, custom)
- Basic Users service example

## Roadmap (short-term)

- Unified error handling & typed errors
- Pagination helpers
- Retry, rate limiting, logging, instrumentation middlewares
- Additional resource services

## Contributing

Issues / PRs welcome. Keep public API small & stable; prefer internal packages for helpers.

## License

Apache 2.0
