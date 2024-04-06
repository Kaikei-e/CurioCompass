package gateway

type RegisterFeedGateway interface {
	RegisterFeed()
}

type RegisterFeedGatewayImpl struct{}

func NewRegisterFeedGateway() RegisterFeedGateway {
	return &RegisterFeedGatewayImpl{}
}

func (r *RegisterFeedGatewayImpl) RegisterFeed() {
	// Register the feed
}
