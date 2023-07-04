package main

import (
	"errors"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/rs/zerolog/log"
)

type Args struct {
	A, B int
}

type Calculator struct{}

func (c *Calculator) Add(args *Args, result *int) error {
	log.Info().Msg("Operation: Add")
	*result = args.A + args.B
	return nil
}

func (c *Calculator) Subtract(args *Args, result *int) error {
	log.Info().Msg("Operation: Subtract")
	*result = args.A - args.B
	return nil
}

func (c *Calculator) Multiply(args *Args, result *int) error {
	log.Info().Msg("Operation: Multiply")
	*result = args.A * args.B
	return nil
}

func (c *Calculator) Divide(args *Args, result *int) error {
	log.Info().Msg("Operation: Divide")

	if args.B == 0 {
		return errors.New("can't divide by zero")
	}

	*result = args.A / args.B
	return nil
}

func main() {
	log.Info().Msg("Starting RPC Server")

	var calculator Calculator

	rpcServer := rpc.NewServer()

	rpcServer.Register(&calculator)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Can't create tcp listener")
	}

	log.Info().Msg("Started RPC Server")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		log.Info().Msg("Accepted new connection")

		go rpcServer.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
