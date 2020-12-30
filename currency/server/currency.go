package server

import (
	"context"

	// alias for protoc generated code
	protogc "github.com/fahmi1597/microservices-go/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

// Currency is a gRPC server, it implements the method defined in CurrencyServer interface
type Currency struct {
	log hclog.Logger
	protogc.UnimplementedCurrencyServer
}

// NewCurrencyServer creates a new instance of Currency with log.
func NewCurrencyServer(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

// GetRate implements CurrencyServer GetRate method from the protoc generated code
func (cs *Currency) GetRate(ctx context.Context, rr *protogc.RateRequest) (*protogc.RateResponse, error) {
	cs.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protogc.RateResponse{Rate: 0.5}, nil
}