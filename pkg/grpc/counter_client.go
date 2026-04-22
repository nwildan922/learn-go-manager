package grpc

import (
	"context"
	"time"

	"github.com/nwildan922/learn-go-manager/proto/counterpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type CounterClient struct {
	conn   *grpc.ClientConn
	client counterpb.CounterServiceClient
}

func NewCounterClient(address string) (*CounterClient, error) {

	conf := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(conf),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	)
	if err != nil {
		return nil, err
	}
	client := counterpb.NewCounterServiceClient(conn)
	return &CounterClient{
		conn:   conn,
		client: client,
	}, nil
}

func (c *CounterClient) SendCounter(counter int32, timestamp string) (*counterpb.CounterResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.client.SendCounter(ctx, &counterpb.CounterRequest{
		Counter: counter,
	})
}

func (c *CounterClient) Close() error {
	return c.conn.Close()
}
