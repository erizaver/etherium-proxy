package etherium

import (
	"testing"

	"github.com/erizaver/etherium_proxy/pkg/api"
)

func TestEthFacade_getHexblockId(t *testing.T) {
	type fields struct {
		EthService                    EthService
		UnimplementedEthServiceServer api.UnimplementedEthServiceServer
	}
	type args struct {
		rawblockId string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantHexblockId string
	}{
		{name: "latest block", fields: fields{
			EthService: nil,
			UnimplementedEthServiceServer: api.UnimplementedEthServiceServer{},
		}, args: args{rawblockId: "latest"}, wantHexblockId: "latest"},
		{name: "hex blockId", fields: fields{
			EthService: nil,
			UnimplementedEthServiceServer: api.UnimplementedEthServiceServer{},
		}, args: args{rawblockId: "0x2D"}, wantHexblockId: "0x2D"},
		{name: "numerical blockId, cache update", fields: fields{
			EthService: nil,
			UnimplementedEthServiceServer: api.UnimplementedEthServiceServer{},
		}, args: args{rawblockId: "10"}, wantHexblockId: "0xa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EthApi{
				EthService:                    tt.fields.EthService,
				UnimplementedEthServiceServer: tt.fields.UnimplementedEthServiceServer,
			}
			if gotHexblockId := e.getSafeblockId(tt.args.rawblockId); gotHexblockId != tt.wantHexblockId {
				t.Errorf("getSafeblockId() = %v, want %v", gotHexblockId, tt.wantHexblockId)
			}
		})
	}
}