package geth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tenderly/tenderly-cli/ethereum/schema"
	"github.com/tenderly/tenderly-cli/ethereum/types"
	"github.com/tenderly/tenderly-cli/jsonrpc2"
)

var DefaultSchema = Schema{
	ValueEth:    ethSchema{},
	ValueNet:    netSchema{},
	ValueTrace:  traceSchema{},
	ValuePubSub: pubSubSchema{},
}

type Schema struct {
	ValueEth    schema.EthSchema
	ValueNet    schema.NetSchema
	ValueTrace  schema.TraceSchema
	ValuePubSub schema.PubSubSchema
}

func (s *Schema) Eth() schema.EthSchema {
	return s.ValueEth
}

func (s *Schema) Net() schema.NetSchema {
	return s.ValueNet
}

func (s *Schema) Trace() schema.TraceSchema {
	return s.ValueTrace
}

func (s *Schema) PubSub() schema.PubSubSchema {
	return s.ValuePubSub
}

// Eth

type ethSchema struct {
}

func (ethSchema) BlockNumber() (*jsonrpc2.Request, *types.Number) {
	var num types.Number

	return jsonrpc2.NewRequest("eth_blockNumber"), &num
}

func (ethSchema) GetBlockByNumber(num types.Number) (*jsonrpc2.Request, types.Block) {
	var block Block

	return jsonrpc2.NewRequest("eth_getBlockByNumber", num.Hex(), true), &block
}

func (ethSchema) GetBlockByHash(hash string) (*jsonrpc2.Request, types.BlockHeader) {
	var block BlockHeader

	return jsonrpc2.NewRequest("eth_getBlockByHash", hash, false), &block
}

func (ethSchema) GetTransaction(hash string) (*jsonrpc2.Request, types.Transaction) {
	var t Transaction

	return jsonrpc2.NewRequest("eth_getTransactionByHash", hash), &t
}

func (ethSchema) GetTransactionReceipt(hash string) (*jsonrpc2.Request, types.TransactionReceipt) {
	var receipt TransactionReceipt

	return jsonrpc2.NewRequest("eth_getTransactionReceipt", hash), &receipt
}

func (ethSchema) GetBalance(address string, block *types.Number) (*jsonrpc2.Request, *hexutil.Big) {
	var balance hexutil.Big

	param := "latest"
	if block != nil {
		param = fmt.Sprintf("0x%x", *block)
	}

	return jsonrpc2.NewRequest("eth_getBalance", address, param), &balance // "latest"
}

func (ethSchema) GetCode(address string, block *types.Number) (*jsonrpc2.Request, *string) {
	var code string

	param := "latest"
	if block != nil {
		param = fmt.Sprintf("0x%x", *block)
	}

	return jsonrpc2.NewRequest("eth_getCode", address, param), &code
}

func (ethSchema) GetNonce(address string, block *types.Number) (*jsonrpc2.Request, *hexutil.Uint64) {
	var nonce hexutil.Uint64

	param := "latest"
	if block != nil {
		param = fmt.Sprintf("0x%x", *block)
	}

	return jsonrpc2.NewRequest("eth_getTransactionCount", address, param), &nonce
}

func (ethSchema) GetStorage(address string, offset common.Hash, block *types.Number) (*jsonrpc2.Request, *string) {
	var data string

	param := "latest"
	if block != nil {
		param = fmt.Sprintf("0x%x", *block)
	}

	return jsonrpc2.NewRequest("eth_getStorageAt", address, offset, param), &data
}

// Net

type netSchema struct {
}

func (netSchema) Version() (*jsonrpc2.Request, *string) {
	var v string

	return jsonrpc2.NewRequest("net_version"), &v
}

// States

type traceSchema struct {
}

func (traceSchema) VMTrace(hash string) (*jsonrpc2.Request, types.TransactionStates) {
	var trace TraceResult

	return jsonrpc2.NewRequest("debug_traceTransaction", hash, struct{}{}), &trace
}
func (traceSchema) CallTrace(hash string) (*jsonrpc2.Request, types.CallTraces) {
	var trace CallTrace

	return jsonrpc2.NewRequest("debug_traceTransaction", hash, map[string]string{"tracer": "callTracer"}), &trace
}

// PubSub

type PubSubSchema interface {
	Subscribe() (*jsonrpc2.Request, *types.SubscriptionID)
	Unsubscribe(id types.SubscriptionID) (*jsonrpc2.Request, *types.UnsubscribeSuccess)
}

type pubSubSchema struct {
}

func (pubSubSchema) Subscribe() (*jsonrpc2.Request, *types.SubscriptionID) {
	id := types.NewNilSubscriptionID()

	return jsonrpc2.NewRequest("eth_subscribe", "newHeads"), &id
}

func (pubSubSchema) Unsubscribe(id types.SubscriptionID) (*jsonrpc2.Request, *types.UnsubscribeSuccess) {
	var success types.UnsubscribeSuccess

	return jsonrpc2.NewRequest("eth_unsubscribe", id.String()), &success
}
