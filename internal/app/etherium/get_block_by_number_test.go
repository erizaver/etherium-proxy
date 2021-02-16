package etherium

import (
	"testing"

	"github.com/erizaver/etherium_proxy/pkg/api"
)

func TestEthFacade_getHexBlockId(t *testing.T) {
	type fields struct {
		EthService                    EthService
		UnimplementedEthServiceServer api.UnimplementedEthServiceServer
	}
	type args struct {
		rawBlockId string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantHexBlockId string
	}{
		{name: "latest block", fields: fields{
			EthService: nil,
			UnimplementedEthServiceServer: api.UnimplementedEthServiceServer{},
		}, args: args{rawBlockId: "latest"}, wantHexBlockId: "latest"},
		{name: "hex blockId", fields: fields{
			EthService: nil,
			UnimplementedEthServiceServer: api.UnimplementedEthServiceServer{},
		}, args: args{rawBlockId: "0x2D"}, wantHexBlockId: "0x2D"},
		{name: "numerical blockId", fields: fields{
			EthService: nil,
			UnimplementedEthServiceServer: api.UnimplementedEthServiceServer{},
		}, args: args{rawBlockId: "10"}, wantHexBlockId: "0xa"},
		{name: "bad block ID", fields: fields{
			EthService: nil,
			UnimplementedEthServiceServer: api.UnimplementedEthServiceServer{},
		}, args: args{rawBlockId: "testBlockId"}, wantHexBlockId: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EthApi{
				EthService:                    tt.fields.EthService,
				UnimplementedEthServiceServer: tt.fields.UnimplementedEthServiceServer,
			}
			if gotHexblockId := e.getSafeBlockId(tt.args.rawBlockId); gotHexblockId != tt.wantHexBlockId {
				t.Errorf("getSafeBlockId() = %v, want %v", gotHexblockId, tt.wantHexBlockId)
			}
		})
	}
}