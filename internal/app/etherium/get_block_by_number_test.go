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
		rawBlockID string
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
		}, args: args{rawBlockID: "latest"}, wantHexBlockId: "latest"},
		{name: "hex blockID", fields: fields{
			EthService: nil,
			UnimplementedEthServiceServer: api.UnimplementedEthServiceServer{},
		}, args: args{rawBlockID: "0x2D"}, wantHexBlockId: "0x2D"},
		{name: "numerical blockID, cache update", fields: fields{
			EthService: nil,
			UnimplementedEthServiceServer: api.UnimplementedEthServiceServer{},
		}, args: args{rawBlockID: "10"}, wantHexBlockId: "0xa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EthFacade{
				EthService:                    tt.fields.EthService,
				UnimplementedEthServiceServer: tt.fields.UnimplementedEthServiceServer,
			}
			if gotHexBlockId := e.getHexBlockId(tt.args.rawBlockID); gotHexBlockId != tt.wantHexBlockId {
				t.Errorf("getHexBlockId() = %v, want %v", gotHexBlockId, tt.wantHexBlockId)
			}
		})
	}
}