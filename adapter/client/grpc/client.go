package grpc

import (
	"context"
	"time"

	"github.com/MihasBel/test-transactions-rest/model"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	v1transaction "github.com/MihasBel/test-transactions-service/delivery/grpc/gen/v1/transaction"
)

// Client grpc client
type Client struct {
	log  zerolog.Logger
	conn *grpc.ClientConn

	cli v1transaction.TransactionAPIClient
	cfg Config
}

// New constructor
func New(cfg Config, log zerolog.Logger) *Client {
	return &Client{
		cfg: cfg,
		log: log,
	}
}

// Start client
func (c *Client) Start(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.cfg.DealTimeout)*time.Second)
	defer cancel()

	grpcClientDialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    time.Duration(c.cfg.KeepAliveTime) * time.Second,
			Timeout: time.Duration(c.cfg.KeepAliveTimeout) * time.Second,
		}),
		grpc.WithBlock(),
	}

	clientConn, err := grpc.DialContext(ctx, c.cfg.Endpoint, grpcClientDialOpts...)
	if err != nil {
		return err
	}

	c.cli = v1transaction.NewTransactionAPIClient(clientConn)
	c.conn = clientConn

	return nil
}

// Stop client
func (c *Client) Stop(_ context.Context) error {
	return c.conn.Close()
}

// GetTransaction by ID
func (c *Client) GetTransaction(ctx context.Context, id uuid.UUID) (*model.Transaction, error) {
	respPb, err := c.cli.ByID(ctx, &v1transaction.ByIDRequest{
		Id: id.String(),
	})
	if err != nil {
		return nil, err
	}
	return &model.Transaction{
		ID:          uuid.MustParse(respPb.Id),
		UserID:      uuid.MustParse(respPb.UserId),
		Amount:      int(respPb.Amount),
		CreatedAt:   respPb.CreatedAt.AsTime(),
		Status:      int(respPb.Status),
		Description: respPb.Description,
	}, nil
}
