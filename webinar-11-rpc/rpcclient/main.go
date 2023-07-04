package main

import (
	"net/rpc/jsonrpc"

	"github.com/rs/zerolog/log"
)

type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Can't create jsonrpc client")
	}

	args := &Args{1, 2}
	var result int

	if err := client.Call("Calculator.Add", args, &result); err != nil {
		log.Fatal().Err(err).Msg("Can't call calculator method")
	}

	log.Info().Msgf("%v + %v = %v", args.A, args.B, result)

	args = &Args{5, 1}

	if err := client.Call("Calculator.Divide", args, &result); err != nil {
		log.Fatal().Err(err).Msg("Can't call calculator method")
	}

	log.Info().Msgf("%v / %v = %v", args.A, args.B, result)
}
