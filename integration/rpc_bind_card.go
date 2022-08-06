package integration

var BindCardRpc IBindCardRpc

//go:generate mockgen -source rpc_bind_card.go -destination mock/rpc_bind_card_mock.go -package mock
type IBindCardRpc interface {
	BindCardRpc(bindCardNo string) error
}

type BindCardRpcImpl struct {
}

func (b BindCardRpcImpl) BindCardRpc(bindCardNo string) error {
	return nil
}

func InitBindCardRpc() {
	BindCardRpc = &BindCardRpcImpl{}
}
