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
func (s *EthService) GetBlockByNumber(blockId string) (*model.Block, error) {
	if blockId == "" {
		return nil, errors.New("block Id can`t be empty")
	}

	if blockId != "latest"{
		blockFromCache, ok := s.getBlockFromCache(blockId)
		if ok {
			return blockFromCache, nil
		}
	}

	block, err := s.getBlockAndUpdateCache(blockId)
	if err != nil {
		return nil, err
	}

	return block, nil
}

//getBlockAndUpdateCache will get block from external API and update cache if cacheUpdate is true
func (s *EthService) getBlockAndUpdateCache(blockId string) (*model.Block, error) {
	block, err := s.EthClient.GetBlockByNumber(strings.ToLower(blockId))
	if err != nil {
		return nil, errors.Wrap(err, "can`t get block from client")
	}

	block.FastTransactions = make(map[string]model.Transaction, len(block.Transactions))
	for _, v := range block.Transactions {
		block.FastTransactions[v.Hash] = v
	}

	numblockId, err := strconv.ParseInt(strings.Replace(block.Number, "0x", "", -1), 16, 64)
	if err != nil {
		return nil, errors.New("unable to parse hex block id to int")
	}

	if s.BlockCounter-numblockId > uncachableBlocksThreshold {
		s.Cache.Add(block.Number, *block)
	}

	blockNumber, err := strconv.ParseInt(strings.Replace(block.Number, "0x", "", -1), 16, 64)
	if err != nil {
		return nil, errors.Wrap(err, "can`t parse block ID")
	}

	if s.BlockCounter < blockNumber {
		s.BlockCounter = blockNumber
	}

	return block, nil
}

func (s *EthService) getBlockFromCache(blockId string) (*model.Block, bool) {
	block, ok := s.Cache.Get(strings.ToLower(blockId))
	if !ok {
		return nil, false
	}

	res, ok := block.(model.Block)
	if !ok {
		glog.Error("unable to cast block from cache")
		return nil, false
	}

	return &res, true
}
