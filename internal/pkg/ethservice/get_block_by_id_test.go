package ethservice

import (
	"testing"

	"github.com/golang/glog"
	lru "github.com/hashicorp/golang-lru"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/erizaver/etherium_proxy/internal/pkg/ethservice/mocks"
	"github.com/erizaver/etherium_proxy/internal/pkg/model"
)

const (
	cacheSize = 5
)

func TestEthService_GetBlockByNumber(t *testing.T) {
	cache, err := lru.New(cacheSize)
	if err != nil {
		glog.Fatal("unable to start user server", err)
	}
	mockEthCli := new(mocks.EthClient)
	mockEthCli.On("GetBlockByNumber", mock.AnythingOfType("string")).Once().Return(model.Block{
		Hash:   "testHash",
		Number: "0x12",
	}, nil)

	service := NewEthService(mockEthCli, cache)
	service.BlockCounter = 50

	block, err := service.GetBlockByNumber("0x12")
	assert.NoError(t, err)
	assert.Equal(t, "0x12", block.Number)
	assert.Equal(t, "testHash", block.Hash)

	// check if it will get same info from cache
	block2, err := service.GetBlockByNumber("0x12")
	assert.NoError(t, err)
	assert.Equal(t, "0x12", block2.Number)
	assert.Equal(t, "testHash", block2.Hash)
}