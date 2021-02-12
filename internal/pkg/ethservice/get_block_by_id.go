package ethservice

import (
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/erizaver/etherium_proxy/internal/pkg/model"
)

const uncachableBlocksThreshold = 25

// GetBlockByNumber will return block by its number, hex number, or "latest" block. Also if its older than 25 block, it will be cached.
// we can`t cache younger blocks since they are not trustworthy
func (s *EthService) GetBlockByNumber(hexBlockId string) (model.Block, error) {
	if hexBlockId == "" {
		return model.Block{}, errors.New("block ID can`t be empty")
	}

	blockFromCache, ok := s.getBlockFromCache(hexBlockId)
	if ok {
		return blockFromCache, nil
	}

	block, err := s.getBlockAndUpdateCache(hexBlockId)
	if err != nil {
		return model.Block{}, err
	}

	return block, nil
}

//getBlockAndUpdateCache will get block from external API and update cache if cacheUpdate is true
func (s *EthService) getBlockAndUpdateCache(blockId string) (model.Block, error) {
	block, err := s.EthClient.GetBlockByNumber(blockId)
	if err != nil {
		return model.Block{}, errors.Wrap(err, "can`t get block from client")
	}

	block.FastTransactions = make(map[string]model.Transaction, len(block.Transactions))
	for _, v := range block.Transactions {
		block.FastTransactions[v.Hash] = v
	}

	numBlockId, err := strconv.ParseInt(strings.Replace(block.Number, "0x", "", -1), 16, 64)
	if err != nil {
		return model.Block{}, errors.New("unable to parse hex block id to int")
	}

	if s.BlockCounter-numBlockId > uncachableBlocksThreshold {
		s.Cache.Add(block.Number, block)
	}

	blockNumber, err := strconv.ParseInt(strings.Replace(block.Number, "0x", "", -1), 16, 64)
	if err != nil {
		return model.Block{}, errors.Wrap(err, "can`t parse block ID")
	}

	if s.BlockCounter < blockNumber {
		s.BlockCounter = blockNumber
	}

	return block, nil
}

func (s *EthService) getBlockFromCache(blockID string) (model.Block, bool) {
	block, ok := s.Cache.Get(blockID)
	if !ok {
		return model.Block{}, false
	}

	res, ok := block.(model.Block)
	if !ok {
		glog.Error("unable to cast block from cache")
	}

	return res, true
}
