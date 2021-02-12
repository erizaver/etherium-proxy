package ethservice

import (
	lru "github.com/hashicorp/golang-lru"

	"github.com/erizaver/etherium_proxy/internal/pkg/model"
)

type EthService struct {
	// We have to keep im mind, that last~20 block can`t be trusted, do we will not add them to cache, to get most trusted info we can
	BlockCounter int64
	EthClient    EthClient
	Cache        *lru.Cache
}

func NewEthService(ec EthClient, c *lru.Cache) *EthService {
	return &EthService{
		EthClient: ec,
		Cache: c,
	}
}

type EthClient interface {
	GetBlockByNumber(blockID string) (model.Block, error)
}