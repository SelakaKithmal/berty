package bertyprotocol

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var _ Client = (*client)(nil)

// Client is the main Berty Protocol interface
type Client interface {
	InstanceServer

	Close() error
}

type client struct {
	// variables
	db   *gorm.DB
	opts Opts
}

// Opts contains optional configuration flags for building a new Client
type Opts struct {
	Logger *zap.Logger
}

// New initializes a new Client
func New(db *gorm.DB, opts Opts) (Client, error) {
	if opts.Logger == nil {
		opts.Logger = zap.NewNop()
	}
	return &client{
		db:   db,
		opts: opts,
	}, nil
}

func (c *client) Close() error {
	return nil
}