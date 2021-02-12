package etherium

import (
	"github.com/erizaver/etherium_proxy/internal/pkg/model"
	"github.com/erizaver/etherium_proxy/pkg/api"
)

func castModelBlockToPb(block model.Block) *api.Block {
	transactions := make([]*api.Transaction, len(block.Transactions))
	for k,v := range block.Transactions{
		transactions[k] = castModelTransactionToPb(&v)
	}

	uncles := make([]string, len(block.Uncles))
	for k,v := range block.Uncles{
		uncles[k] = v
	}

	return &api.Block{
		Difficulty:       block.Difficulty,
		ExtraData:        block.ExtraDate,
		GasLimit:         block.GasLimit,
		GasUsed:          block.GasUsed,
		Hash:             block.Hash,
		LogsBloom:        block.LogsBloom,
		Miner:            block.Miner,
		MixHash:          block.MixHash,
		Nonce:            block.Nonce,
		Number:           block.Number,
		ParentHash:       block.ParentHash,
		ReceiptsRoot:     block.ReceiptsRoot,
		Sha3Uncles:       block.Sha3Uncles,
		Size:             block.Size,
		StateRoot:        block.StateRoot,
		Timestamp:        block.Timestamp,
		TotalDifficulty:  block.TotalDifficulty,
		Transactions:     transactions,
		TransactionsRoot: block.TransactionsRoot,
		Uncles:           uncles,
	}
}

func castModelTransactionToPb(transaction *model.Transaction) *api.Transaction {
	return &api.Transaction{
		BlockHash:        transaction.BlockHash,
		BlockNumber:      transaction.BlockNumber,
		From:             transaction.From,
		Gas:              transaction.Gas,
		GasPrice:         transaction.GasPrice,
		Hash:             transaction.Hash,
		Input:            transaction.Input,
		Nonce:            transaction.Nonce,
		To:               transaction.To,
		TransactionIndex: transaction.TransactionIndex,
		Value:            transaction.Value,
		V:                transaction.V,
		R:                transaction.R,
		S:                transaction.S,
	}
}
