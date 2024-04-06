package registerFeed

import "feedFetcher/gateway"

func RegisterSingleFeed() error {
	registerFeedGateway := gateway.NewRegisterFeedGateway()
	registerFeedGateway.RegisterFeed()

	return nil
}
