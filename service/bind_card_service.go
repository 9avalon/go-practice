package service

import "go-practice/integration"

var BindCardService IBindCardService

type BindCardReq struct {
	BankCardName string
	BankCardNo   string
}

type BindCardRsp struct {
	Code int
	Msg  string
}

type IBindCardService interface {
	BindCard(req *BindCardReq) (*BindCardRsp, error)
}

type BindCardServiceImpl struct {
	BindCardRpc integration.IBindCardRpc
}

func InitBindCardService() {
	BindCardService = &BindCardServiceImpl{
		BindCardRpc: integration.BindCardRpc,
	}
}

//go:generate mockgen -source bind_card_service.go -destination mock/bind_card_service_mock.go -package service
func (b BindCardServiceImpl) BindCard(req *BindCardReq) (*BindCardRsp, error) {
	// validate params
	err := validateParams(req)
	if err != nil {
		return nil, err
	}

	// do bind card
	err = doBindCard(req)
	if err != nil {
		return nil, err
	}

	// bind card rpc
	err = b.BindCardRpc.BindCardRpc(req.BankCardName)
	if err != nil {
		return nil, err
	}

	rsp := &BindCardRsp{
		Code: 200,
		Msg:  "",
	}
	return rsp, nil
}

func validateParams(req *BindCardReq) error {
	return nil
}

func doBindCard(req *BindCardReq) error {

	return nil
}
