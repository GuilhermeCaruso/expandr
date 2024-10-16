package server

type config struct {
	port int
}

type Option func(*config)

func WithPort(port int) Option {
	return func(c *config) {
		c.port = port
	}
}

func (c *config) defaults() {
	c.port = 3001
}

func newConfig(opts ...Option) *config {
	c := new(config)

	c.defaults()

	for _, opt := range opts {
		opt(c)
	}

	return c
}
