package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"go-practice/integration"
	"go-practice/integration/mock"
	"reflect"
	"testing"
)

func TestBindCardServiceImpl_BindCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()

	// mock绑卡远程服务
	mockBindCardRpc := mock.NewMockIBindCardRpc(ctrl)
	mockBindCardRpc.EXPECT().BindCardRpc(gomock.Eq("Miguel")).Return(nil)
	mockBindCardRpc.EXPECT().BindCardRpc(gomock.Eq("Tom")).Return(errors.New("rpc fail"))

	type fields struct {
		BindCardRpc integration.IBindCardRpc
	}
	type args struct {
		req *BindCardReq
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BindCardRsp
		wantErr bool
	}{
		{
			name:   "success bind card",
			fields: fields{BindCardRpc: mockBindCardRpc},
			args: args{req: &BindCardReq{
				BankCardName: "Miguel",
				BankCardNo:   "123",
			}},
			want: &BindCardRsp{
				Code: 200,
				Msg:  "",
			},
			wantErr: false,
		},
		{
			name:   "fail bind card",
			fields: fields{BindCardRpc: mockBindCardRpc},
			args: args{req: &BindCardReq{
				BankCardName: "Tom",
				BankCardNo:   "123",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BindCardServiceImpl{
				BindCardRpc: tt.fields.BindCardRpc,
			}
			got, err := b.BindCard(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("BindCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BindCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}
