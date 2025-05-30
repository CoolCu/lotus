// Code generated by github.com/filecoin-project/lotus/gen/api. DO NOT EDIT.

package v2api

import (
	"context"

	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
)

var ErrNotSupported = xerrors.New("method not supported")

type FullNodeStruct struct {
	Internal FullNodeMethods
}

type FullNodeMethods struct {
	ChainGetTipSet func(p0 context.Context, p1 types.TipSetSelector) (*types.TipSet, error) `perm:"read"`

	EthAccounts func(p0 context.Context) ([]ethtypes.EthAddress, error) `perm:"read"`

	EthAddressToFilecoinAddress func(p0 context.Context, p1 ethtypes.EthAddress) (address.Address, error) `perm:"read"`

	EthBlockNumber func(p0 context.Context) (ethtypes.EthUint64, error) `perm:"read"`

	EthCall func(p0 context.Context, p1 ethtypes.EthCall, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error) `perm:"read"`

	EthChainId func(p0 context.Context) (ethtypes.EthUint64, error) `perm:"read"`

	EthEstimateGas func(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthUint64, error) `perm:"read"`

	EthFeeHistory func(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthFeeHistory, error) `perm:"read"`

	EthGasPrice func(p0 context.Context) (ethtypes.EthBigInt, error) `perm:"read"`

	EthGetBalance func(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBigInt, error) `perm:"read"`

	EthGetBlockByHash func(p0 context.Context, p1 ethtypes.EthHash, p2 bool) (ethtypes.EthBlock, error) `perm:"read"`

	EthGetBlockByNumber func(p0 context.Context, p1 string, p2 bool) (ethtypes.EthBlock, error) `perm:"read"`

	EthGetBlockReceipts func(p0 context.Context, p1 ethtypes.EthBlockNumberOrHash) ([]*ethtypes.EthTxReceipt, error) `perm:"read"`

	EthGetBlockReceiptsLimited func(p0 context.Context, p1 ethtypes.EthBlockNumberOrHash, p2 abi.ChainEpoch) ([]*ethtypes.EthTxReceipt, error) `perm:"read"`

	EthGetBlockTransactionCountByHash func(p0 context.Context, p1 ethtypes.EthHash) (ethtypes.EthUint64, error) `perm:"read"`

	EthGetBlockTransactionCountByNumber func(p0 context.Context, p1 string) (ethtypes.EthUint64, error) `perm:"read"`

	EthGetCode func(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error) `perm:"read"`

	EthGetFilterChanges func(p0 context.Context, p1 ethtypes.EthFilterID) (*ethtypes.EthFilterResult, error) `perm:"read"`

	EthGetFilterLogs func(p0 context.Context, p1 ethtypes.EthFilterID) (*ethtypes.EthFilterResult, error) `perm:"read"`

	EthGetLogs func(p0 context.Context, p1 *ethtypes.EthFilterSpec) (*ethtypes.EthFilterResult, error) `perm:"read"`

	EthGetMessageCidByTransactionHash func(p0 context.Context, p1 *ethtypes.EthHash) (*cid.Cid, error) `perm:"read"`

	EthGetStorageAt func(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBytes, p3 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error) `perm:"read"`

	EthGetTransactionByBlockHashAndIndex func(p0 context.Context, p1 ethtypes.EthHash, p2 ethtypes.EthUint64) (*ethtypes.EthTx, error) `perm:"read"`

	EthGetTransactionByBlockNumberAndIndex func(p0 context.Context, p1 string, p2 ethtypes.EthUint64) (*ethtypes.EthTx, error) `perm:"read"`

	EthGetTransactionByHash func(p0 context.Context, p1 *ethtypes.EthHash) (*ethtypes.EthTx, error) `perm:"read"`

	EthGetTransactionByHashLimited func(p0 context.Context, p1 *ethtypes.EthHash, p2 abi.ChainEpoch) (*ethtypes.EthTx, error) `perm:"read"`

	EthGetTransactionCount func(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthUint64, error) `perm:"read"`

	EthGetTransactionHashByCid func(p0 context.Context, p1 cid.Cid) (*ethtypes.EthHash, error) `perm:"read"`

	EthGetTransactionReceipt func(p0 context.Context, p1 ethtypes.EthHash) (*ethtypes.EthTxReceipt, error) `perm:"read"`

	EthGetTransactionReceiptLimited func(p0 context.Context, p1 ethtypes.EthHash, p2 abi.ChainEpoch) (*ethtypes.EthTxReceipt, error) `perm:"read"`

	EthMaxPriorityFeePerGas func(p0 context.Context) (ethtypes.EthBigInt, error) `perm:"read"`

	EthNewBlockFilter func(p0 context.Context) (ethtypes.EthFilterID, error) `perm:"read"`

	EthNewFilter func(p0 context.Context, p1 *ethtypes.EthFilterSpec) (ethtypes.EthFilterID, error) `perm:"read"`

	EthNewPendingTransactionFilter func(p0 context.Context) (ethtypes.EthFilterID, error) `perm:"read"`

	EthProtocolVersion func(p0 context.Context) (ethtypes.EthUint64, error) `perm:"read"`

	EthSendRawTransaction func(p0 context.Context, p1 ethtypes.EthBytes) (ethtypes.EthHash, error) `perm:"read"`

	EthSendRawTransactionUntrusted func(p0 context.Context, p1 ethtypes.EthBytes) (ethtypes.EthHash, error) `perm:"read"`

	EthSubscribe func(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthSubscriptionID, error) `perm:"read"`

	EthSyncing func(p0 context.Context) (ethtypes.EthSyncingResult, error) `perm:"read"`

	EthTraceBlock func(p0 context.Context, p1 string) ([]*ethtypes.EthTraceBlock, error) `perm:"read"`

	EthTraceFilter func(p0 context.Context, p1 ethtypes.EthTraceFilterCriteria) ([]*ethtypes.EthTraceFilterResult, error) `perm:"read"`

	EthTraceReplayBlockTransactions func(p0 context.Context, p1 string, p2 []string) ([]*ethtypes.EthTraceReplayBlockTransaction, error) `perm:"read"`

	EthTraceTransaction func(p0 context.Context, p1 string) ([]*ethtypes.EthTraceTransaction, error) `perm:"read"`

	EthUninstallFilter func(p0 context.Context, p1 ethtypes.EthFilterID) (bool, error) `perm:"read"`

	EthUnsubscribe func(p0 context.Context, p1 ethtypes.EthSubscriptionID) (bool, error) `perm:"read"`

	FilecoinAddressToEthAddress func(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthAddress, error) `perm:"read"`

	NetListening func(p0 context.Context) (bool, error) `perm:"read"`

	NetVersion func(p0 context.Context) (string, error) `perm:"read"`

	StateGetActor func(p0 context.Context, p1 address.Address, p2 types.TipSetSelector) (*types.Actor, error) `perm:"read"`

	StateGetID func(p0 context.Context, p1 address.Address, p2 types.TipSetSelector) (*address.Address, error) `perm:"read"`

	Web3ClientVersion func(p0 context.Context) (string, error) `perm:"read"`
}

type FullNodeStub struct {
}

func (s *FullNodeStruct) ChainGetTipSet(p0 context.Context, p1 types.TipSetSelector) (*types.TipSet, error) {
	if s.Internal.ChainGetTipSet == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.ChainGetTipSet(p0, p1)
}

func (s *FullNodeStub) ChainGetTipSet(p0 context.Context, p1 types.TipSetSelector) (*types.TipSet, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthAccounts(p0 context.Context) ([]ethtypes.EthAddress, error) {
	if s.Internal.EthAccounts == nil {
		return *new([]ethtypes.EthAddress), ErrNotSupported
	}
	return s.Internal.EthAccounts(p0)
}

func (s *FullNodeStub) EthAccounts(p0 context.Context) ([]ethtypes.EthAddress, error) {
	return *new([]ethtypes.EthAddress), ErrNotSupported
}

func (s *FullNodeStruct) EthAddressToFilecoinAddress(p0 context.Context, p1 ethtypes.EthAddress) (address.Address, error) {
	if s.Internal.EthAddressToFilecoinAddress == nil {
		return *new(address.Address), ErrNotSupported
	}
	return s.Internal.EthAddressToFilecoinAddress(p0, p1)
}

func (s *FullNodeStub) EthAddressToFilecoinAddress(p0 context.Context, p1 ethtypes.EthAddress) (address.Address, error) {
	return *new(address.Address), ErrNotSupported
}

func (s *FullNodeStruct) EthBlockNumber(p0 context.Context) (ethtypes.EthUint64, error) {
	if s.Internal.EthBlockNumber == nil {
		return *new(ethtypes.EthUint64), ErrNotSupported
	}
	return s.Internal.EthBlockNumber(p0)
}

func (s *FullNodeStub) EthBlockNumber(p0 context.Context) (ethtypes.EthUint64, error) {
	return *new(ethtypes.EthUint64), ErrNotSupported
}

func (s *FullNodeStruct) EthCall(p0 context.Context, p1 ethtypes.EthCall, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error) {
	if s.Internal.EthCall == nil {
		return *new(ethtypes.EthBytes), ErrNotSupported
	}
	return s.Internal.EthCall(p0, p1, p2)
}

func (s *FullNodeStub) EthCall(p0 context.Context, p1 ethtypes.EthCall, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error) {
	return *new(ethtypes.EthBytes), ErrNotSupported
}

func (s *FullNodeStruct) EthChainId(p0 context.Context) (ethtypes.EthUint64, error) {
	if s.Internal.EthChainId == nil {
		return *new(ethtypes.EthUint64), ErrNotSupported
	}
	return s.Internal.EthChainId(p0)
}

func (s *FullNodeStub) EthChainId(p0 context.Context) (ethtypes.EthUint64, error) {
	return *new(ethtypes.EthUint64), ErrNotSupported
}

func (s *FullNodeStruct) EthEstimateGas(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthUint64, error) {
	if s.Internal.EthEstimateGas == nil {
		return *new(ethtypes.EthUint64), ErrNotSupported
	}
	return s.Internal.EthEstimateGas(p0, p1)
}

func (s *FullNodeStub) EthEstimateGas(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthUint64, error) {
	return *new(ethtypes.EthUint64), ErrNotSupported
}

func (s *FullNodeStruct) EthFeeHistory(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthFeeHistory, error) {
	if s.Internal.EthFeeHistory == nil {
		return *new(ethtypes.EthFeeHistory), ErrNotSupported
	}
	return s.Internal.EthFeeHistory(p0, p1)
}

func (s *FullNodeStub) EthFeeHistory(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthFeeHistory, error) {
	return *new(ethtypes.EthFeeHistory), ErrNotSupported
}

func (s *FullNodeStruct) EthGasPrice(p0 context.Context) (ethtypes.EthBigInt, error) {
	if s.Internal.EthGasPrice == nil {
		return *new(ethtypes.EthBigInt), ErrNotSupported
	}
	return s.Internal.EthGasPrice(p0)
}

func (s *FullNodeStub) EthGasPrice(p0 context.Context) (ethtypes.EthBigInt, error) {
	return *new(ethtypes.EthBigInt), ErrNotSupported
}

func (s *FullNodeStruct) EthGetBalance(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBigInt, error) {
	if s.Internal.EthGetBalance == nil {
		return *new(ethtypes.EthBigInt), ErrNotSupported
	}
	return s.Internal.EthGetBalance(p0, p1, p2)
}

func (s *FullNodeStub) EthGetBalance(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBigInt, error) {
	return *new(ethtypes.EthBigInt), ErrNotSupported
}

func (s *FullNodeStruct) EthGetBlockByHash(p0 context.Context, p1 ethtypes.EthHash, p2 bool) (ethtypes.EthBlock, error) {
	if s.Internal.EthGetBlockByHash == nil {
		return *new(ethtypes.EthBlock), ErrNotSupported
	}
	return s.Internal.EthGetBlockByHash(p0, p1, p2)
}

func (s *FullNodeStub) EthGetBlockByHash(p0 context.Context, p1 ethtypes.EthHash, p2 bool) (ethtypes.EthBlock, error) {
	return *new(ethtypes.EthBlock), ErrNotSupported
}

func (s *FullNodeStruct) EthGetBlockByNumber(p0 context.Context, p1 string, p2 bool) (ethtypes.EthBlock, error) {
	if s.Internal.EthGetBlockByNumber == nil {
		return *new(ethtypes.EthBlock), ErrNotSupported
	}
	return s.Internal.EthGetBlockByNumber(p0, p1, p2)
}

func (s *FullNodeStub) EthGetBlockByNumber(p0 context.Context, p1 string, p2 bool) (ethtypes.EthBlock, error) {
	return *new(ethtypes.EthBlock), ErrNotSupported
}

func (s *FullNodeStruct) EthGetBlockReceipts(p0 context.Context, p1 ethtypes.EthBlockNumberOrHash) ([]*ethtypes.EthTxReceipt, error) {
	if s.Internal.EthGetBlockReceipts == nil {
		return *new([]*ethtypes.EthTxReceipt), ErrNotSupported
	}
	return s.Internal.EthGetBlockReceipts(p0, p1)
}

func (s *FullNodeStub) EthGetBlockReceipts(p0 context.Context, p1 ethtypes.EthBlockNumberOrHash) ([]*ethtypes.EthTxReceipt, error) {
	return *new([]*ethtypes.EthTxReceipt), ErrNotSupported
}

func (s *FullNodeStruct) EthGetBlockReceiptsLimited(p0 context.Context, p1 ethtypes.EthBlockNumberOrHash, p2 abi.ChainEpoch) ([]*ethtypes.EthTxReceipt, error) {
	if s.Internal.EthGetBlockReceiptsLimited == nil {
		return *new([]*ethtypes.EthTxReceipt), ErrNotSupported
	}
	return s.Internal.EthGetBlockReceiptsLimited(p0, p1, p2)
}

func (s *FullNodeStub) EthGetBlockReceiptsLimited(p0 context.Context, p1 ethtypes.EthBlockNumberOrHash, p2 abi.ChainEpoch) ([]*ethtypes.EthTxReceipt, error) {
	return *new([]*ethtypes.EthTxReceipt), ErrNotSupported
}

func (s *FullNodeStruct) EthGetBlockTransactionCountByHash(p0 context.Context, p1 ethtypes.EthHash) (ethtypes.EthUint64, error) {
	if s.Internal.EthGetBlockTransactionCountByHash == nil {
		return *new(ethtypes.EthUint64), ErrNotSupported
	}
	return s.Internal.EthGetBlockTransactionCountByHash(p0, p1)
}

func (s *FullNodeStub) EthGetBlockTransactionCountByHash(p0 context.Context, p1 ethtypes.EthHash) (ethtypes.EthUint64, error) {
	return *new(ethtypes.EthUint64), ErrNotSupported
}

func (s *FullNodeStruct) EthGetBlockTransactionCountByNumber(p0 context.Context, p1 string) (ethtypes.EthUint64, error) {
	if s.Internal.EthGetBlockTransactionCountByNumber == nil {
		return *new(ethtypes.EthUint64), ErrNotSupported
	}
	return s.Internal.EthGetBlockTransactionCountByNumber(p0, p1)
}

func (s *FullNodeStub) EthGetBlockTransactionCountByNumber(p0 context.Context, p1 string) (ethtypes.EthUint64, error) {
	return *new(ethtypes.EthUint64), ErrNotSupported
}

func (s *FullNodeStruct) EthGetCode(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error) {
	if s.Internal.EthGetCode == nil {
		return *new(ethtypes.EthBytes), ErrNotSupported
	}
	return s.Internal.EthGetCode(p0, p1, p2)
}

func (s *FullNodeStub) EthGetCode(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error) {
	return *new(ethtypes.EthBytes), ErrNotSupported
}

func (s *FullNodeStruct) EthGetFilterChanges(p0 context.Context, p1 ethtypes.EthFilterID) (*ethtypes.EthFilterResult, error) {
	if s.Internal.EthGetFilterChanges == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetFilterChanges(p0, p1)
}

func (s *FullNodeStub) EthGetFilterChanges(p0 context.Context, p1 ethtypes.EthFilterID) (*ethtypes.EthFilterResult, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetFilterLogs(p0 context.Context, p1 ethtypes.EthFilterID) (*ethtypes.EthFilterResult, error) {
	if s.Internal.EthGetFilterLogs == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetFilterLogs(p0, p1)
}

func (s *FullNodeStub) EthGetFilterLogs(p0 context.Context, p1 ethtypes.EthFilterID) (*ethtypes.EthFilterResult, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetLogs(p0 context.Context, p1 *ethtypes.EthFilterSpec) (*ethtypes.EthFilterResult, error) {
	if s.Internal.EthGetLogs == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetLogs(p0, p1)
}

func (s *FullNodeStub) EthGetLogs(p0 context.Context, p1 *ethtypes.EthFilterSpec) (*ethtypes.EthFilterResult, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetMessageCidByTransactionHash(p0 context.Context, p1 *ethtypes.EthHash) (*cid.Cid, error) {
	if s.Internal.EthGetMessageCidByTransactionHash == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetMessageCidByTransactionHash(p0, p1)
}

func (s *FullNodeStub) EthGetMessageCidByTransactionHash(p0 context.Context, p1 *ethtypes.EthHash) (*cid.Cid, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetStorageAt(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBytes, p3 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error) {
	if s.Internal.EthGetStorageAt == nil {
		return *new(ethtypes.EthBytes), ErrNotSupported
	}
	return s.Internal.EthGetStorageAt(p0, p1, p2, p3)
}

func (s *FullNodeStub) EthGetStorageAt(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBytes, p3 ethtypes.EthBlockNumberOrHash) (ethtypes.EthBytes, error) {
	return *new(ethtypes.EthBytes), ErrNotSupported
}

func (s *FullNodeStruct) EthGetTransactionByBlockHashAndIndex(p0 context.Context, p1 ethtypes.EthHash, p2 ethtypes.EthUint64) (*ethtypes.EthTx, error) {
	if s.Internal.EthGetTransactionByBlockHashAndIndex == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetTransactionByBlockHashAndIndex(p0, p1, p2)
}

func (s *FullNodeStub) EthGetTransactionByBlockHashAndIndex(p0 context.Context, p1 ethtypes.EthHash, p2 ethtypes.EthUint64) (*ethtypes.EthTx, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetTransactionByBlockNumberAndIndex(p0 context.Context, p1 string, p2 ethtypes.EthUint64) (*ethtypes.EthTx, error) {
	if s.Internal.EthGetTransactionByBlockNumberAndIndex == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetTransactionByBlockNumberAndIndex(p0, p1, p2)
}

func (s *FullNodeStub) EthGetTransactionByBlockNumberAndIndex(p0 context.Context, p1 string, p2 ethtypes.EthUint64) (*ethtypes.EthTx, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetTransactionByHash(p0 context.Context, p1 *ethtypes.EthHash) (*ethtypes.EthTx, error) {
	if s.Internal.EthGetTransactionByHash == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetTransactionByHash(p0, p1)
}

func (s *FullNodeStub) EthGetTransactionByHash(p0 context.Context, p1 *ethtypes.EthHash) (*ethtypes.EthTx, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetTransactionByHashLimited(p0 context.Context, p1 *ethtypes.EthHash, p2 abi.ChainEpoch) (*ethtypes.EthTx, error) {
	if s.Internal.EthGetTransactionByHashLimited == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetTransactionByHashLimited(p0, p1, p2)
}

func (s *FullNodeStub) EthGetTransactionByHashLimited(p0 context.Context, p1 *ethtypes.EthHash, p2 abi.ChainEpoch) (*ethtypes.EthTx, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetTransactionCount(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthUint64, error) {
	if s.Internal.EthGetTransactionCount == nil {
		return *new(ethtypes.EthUint64), ErrNotSupported
	}
	return s.Internal.EthGetTransactionCount(p0, p1, p2)
}

func (s *FullNodeStub) EthGetTransactionCount(p0 context.Context, p1 ethtypes.EthAddress, p2 ethtypes.EthBlockNumberOrHash) (ethtypes.EthUint64, error) {
	return *new(ethtypes.EthUint64), ErrNotSupported
}

func (s *FullNodeStruct) EthGetTransactionHashByCid(p0 context.Context, p1 cid.Cid) (*ethtypes.EthHash, error) {
	if s.Internal.EthGetTransactionHashByCid == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetTransactionHashByCid(p0, p1)
}

func (s *FullNodeStub) EthGetTransactionHashByCid(p0 context.Context, p1 cid.Cid) (*ethtypes.EthHash, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetTransactionReceipt(p0 context.Context, p1 ethtypes.EthHash) (*ethtypes.EthTxReceipt, error) {
	if s.Internal.EthGetTransactionReceipt == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetTransactionReceipt(p0, p1)
}

func (s *FullNodeStub) EthGetTransactionReceipt(p0 context.Context, p1 ethtypes.EthHash) (*ethtypes.EthTxReceipt, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthGetTransactionReceiptLimited(p0 context.Context, p1 ethtypes.EthHash, p2 abi.ChainEpoch) (*ethtypes.EthTxReceipt, error) {
	if s.Internal.EthGetTransactionReceiptLimited == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.EthGetTransactionReceiptLimited(p0, p1, p2)
}

func (s *FullNodeStub) EthGetTransactionReceiptLimited(p0 context.Context, p1 ethtypes.EthHash, p2 abi.ChainEpoch) (*ethtypes.EthTxReceipt, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) EthMaxPriorityFeePerGas(p0 context.Context) (ethtypes.EthBigInt, error) {
	if s.Internal.EthMaxPriorityFeePerGas == nil {
		return *new(ethtypes.EthBigInt), ErrNotSupported
	}
	return s.Internal.EthMaxPriorityFeePerGas(p0)
}

func (s *FullNodeStub) EthMaxPriorityFeePerGas(p0 context.Context) (ethtypes.EthBigInt, error) {
	return *new(ethtypes.EthBigInt), ErrNotSupported
}

func (s *FullNodeStruct) EthNewBlockFilter(p0 context.Context) (ethtypes.EthFilterID, error) {
	if s.Internal.EthNewBlockFilter == nil {
		return *new(ethtypes.EthFilterID), ErrNotSupported
	}
	return s.Internal.EthNewBlockFilter(p0)
}

func (s *FullNodeStub) EthNewBlockFilter(p0 context.Context) (ethtypes.EthFilterID, error) {
	return *new(ethtypes.EthFilterID), ErrNotSupported
}

func (s *FullNodeStruct) EthNewFilter(p0 context.Context, p1 *ethtypes.EthFilterSpec) (ethtypes.EthFilterID, error) {
	if s.Internal.EthNewFilter == nil {
		return *new(ethtypes.EthFilterID), ErrNotSupported
	}
	return s.Internal.EthNewFilter(p0, p1)
}

func (s *FullNodeStub) EthNewFilter(p0 context.Context, p1 *ethtypes.EthFilterSpec) (ethtypes.EthFilterID, error) {
	return *new(ethtypes.EthFilterID), ErrNotSupported
}

func (s *FullNodeStruct) EthNewPendingTransactionFilter(p0 context.Context) (ethtypes.EthFilterID, error) {
	if s.Internal.EthNewPendingTransactionFilter == nil {
		return *new(ethtypes.EthFilterID), ErrNotSupported
	}
	return s.Internal.EthNewPendingTransactionFilter(p0)
}

func (s *FullNodeStub) EthNewPendingTransactionFilter(p0 context.Context) (ethtypes.EthFilterID, error) {
	return *new(ethtypes.EthFilterID), ErrNotSupported
}

func (s *FullNodeStruct) EthProtocolVersion(p0 context.Context) (ethtypes.EthUint64, error) {
	if s.Internal.EthProtocolVersion == nil {
		return *new(ethtypes.EthUint64), ErrNotSupported
	}
	return s.Internal.EthProtocolVersion(p0)
}

func (s *FullNodeStub) EthProtocolVersion(p0 context.Context) (ethtypes.EthUint64, error) {
	return *new(ethtypes.EthUint64), ErrNotSupported
}

func (s *FullNodeStruct) EthSendRawTransaction(p0 context.Context, p1 ethtypes.EthBytes) (ethtypes.EthHash, error) {
	if s.Internal.EthSendRawTransaction == nil {
		return *new(ethtypes.EthHash), ErrNotSupported
	}
	return s.Internal.EthSendRawTransaction(p0, p1)
}

func (s *FullNodeStub) EthSendRawTransaction(p0 context.Context, p1 ethtypes.EthBytes) (ethtypes.EthHash, error) {
	return *new(ethtypes.EthHash), ErrNotSupported
}

func (s *FullNodeStruct) EthSendRawTransactionUntrusted(p0 context.Context, p1 ethtypes.EthBytes) (ethtypes.EthHash, error) {
	if s.Internal.EthSendRawTransactionUntrusted == nil {
		return *new(ethtypes.EthHash), ErrNotSupported
	}
	return s.Internal.EthSendRawTransactionUntrusted(p0, p1)
}

func (s *FullNodeStub) EthSendRawTransactionUntrusted(p0 context.Context, p1 ethtypes.EthBytes) (ethtypes.EthHash, error) {
	return *new(ethtypes.EthHash), ErrNotSupported
}

func (s *FullNodeStruct) EthSubscribe(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthSubscriptionID, error) {
	if s.Internal.EthSubscribe == nil {
		return *new(ethtypes.EthSubscriptionID), ErrNotSupported
	}
	return s.Internal.EthSubscribe(p0, p1)
}

func (s *FullNodeStub) EthSubscribe(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthSubscriptionID, error) {
	return *new(ethtypes.EthSubscriptionID), ErrNotSupported
}

func (s *FullNodeStruct) EthSyncing(p0 context.Context) (ethtypes.EthSyncingResult, error) {
	if s.Internal.EthSyncing == nil {
		return *new(ethtypes.EthSyncingResult), ErrNotSupported
	}
	return s.Internal.EthSyncing(p0)
}

func (s *FullNodeStub) EthSyncing(p0 context.Context) (ethtypes.EthSyncingResult, error) {
	return *new(ethtypes.EthSyncingResult), ErrNotSupported
}

func (s *FullNodeStruct) EthTraceBlock(p0 context.Context, p1 string) ([]*ethtypes.EthTraceBlock, error) {
	if s.Internal.EthTraceBlock == nil {
		return *new([]*ethtypes.EthTraceBlock), ErrNotSupported
	}
	return s.Internal.EthTraceBlock(p0, p1)
}

func (s *FullNodeStub) EthTraceBlock(p0 context.Context, p1 string) ([]*ethtypes.EthTraceBlock, error) {
	return *new([]*ethtypes.EthTraceBlock), ErrNotSupported
}

func (s *FullNodeStruct) EthTraceFilter(p0 context.Context, p1 ethtypes.EthTraceFilterCriteria) ([]*ethtypes.EthTraceFilterResult, error) {
	if s.Internal.EthTraceFilter == nil {
		return *new([]*ethtypes.EthTraceFilterResult), ErrNotSupported
	}
	return s.Internal.EthTraceFilter(p0, p1)
}

func (s *FullNodeStub) EthTraceFilter(p0 context.Context, p1 ethtypes.EthTraceFilterCriteria) ([]*ethtypes.EthTraceFilterResult, error) {
	return *new([]*ethtypes.EthTraceFilterResult), ErrNotSupported
}

func (s *FullNodeStruct) EthTraceReplayBlockTransactions(p0 context.Context, p1 string, p2 []string) ([]*ethtypes.EthTraceReplayBlockTransaction, error) {
	if s.Internal.EthTraceReplayBlockTransactions == nil {
		return *new([]*ethtypes.EthTraceReplayBlockTransaction), ErrNotSupported
	}
	return s.Internal.EthTraceReplayBlockTransactions(p0, p1, p2)
}

func (s *FullNodeStub) EthTraceReplayBlockTransactions(p0 context.Context, p1 string, p2 []string) ([]*ethtypes.EthTraceReplayBlockTransaction, error) {
	return *new([]*ethtypes.EthTraceReplayBlockTransaction), ErrNotSupported
}

func (s *FullNodeStruct) EthTraceTransaction(p0 context.Context, p1 string) ([]*ethtypes.EthTraceTransaction, error) {
	if s.Internal.EthTraceTransaction == nil {
		return *new([]*ethtypes.EthTraceTransaction), ErrNotSupported
	}
	return s.Internal.EthTraceTransaction(p0, p1)
}

func (s *FullNodeStub) EthTraceTransaction(p0 context.Context, p1 string) ([]*ethtypes.EthTraceTransaction, error) {
	return *new([]*ethtypes.EthTraceTransaction), ErrNotSupported
}

func (s *FullNodeStruct) EthUninstallFilter(p0 context.Context, p1 ethtypes.EthFilterID) (bool, error) {
	if s.Internal.EthUninstallFilter == nil {
		return false, ErrNotSupported
	}
	return s.Internal.EthUninstallFilter(p0, p1)
}

func (s *FullNodeStub) EthUninstallFilter(p0 context.Context, p1 ethtypes.EthFilterID) (bool, error) {
	return false, ErrNotSupported
}

func (s *FullNodeStruct) EthUnsubscribe(p0 context.Context, p1 ethtypes.EthSubscriptionID) (bool, error) {
	if s.Internal.EthUnsubscribe == nil {
		return false, ErrNotSupported
	}
	return s.Internal.EthUnsubscribe(p0, p1)
}

func (s *FullNodeStub) EthUnsubscribe(p0 context.Context, p1 ethtypes.EthSubscriptionID) (bool, error) {
	return false, ErrNotSupported
}

func (s *FullNodeStruct) FilecoinAddressToEthAddress(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthAddress, error) {
	if s.Internal.FilecoinAddressToEthAddress == nil {
		return *new(ethtypes.EthAddress), ErrNotSupported
	}
	return s.Internal.FilecoinAddressToEthAddress(p0, p1)
}

func (s *FullNodeStub) FilecoinAddressToEthAddress(p0 context.Context, p1 jsonrpc.RawParams) (ethtypes.EthAddress, error) {
	return *new(ethtypes.EthAddress), ErrNotSupported
}

func (s *FullNodeStruct) NetListening(p0 context.Context) (bool, error) {
	if s.Internal.NetListening == nil {
		return false, ErrNotSupported
	}
	return s.Internal.NetListening(p0)
}

func (s *FullNodeStub) NetListening(p0 context.Context) (bool, error) {
	return false, ErrNotSupported
}

func (s *FullNodeStruct) NetVersion(p0 context.Context) (string, error) {
	if s.Internal.NetVersion == nil {
		return "", ErrNotSupported
	}
	return s.Internal.NetVersion(p0)
}

func (s *FullNodeStub) NetVersion(p0 context.Context) (string, error) {
	return "", ErrNotSupported
}

func (s *FullNodeStruct) StateGetActor(p0 context.Context, p1 address.Address, p2 types.TipSetSelector) (*types.Actor, error) {
	if s.Internal.StateGetActor == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.StateGetActor(p0, p1, p2)
}

func (s *FullNodeStub) StateGetActor(p0 context.Context, p1 address.Address, p2 types.TipSetSelector) (*types.Actor, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) StateGetID(p0 context.Context, p1 address.Address, p2 types.TipSetSelector) (*address.Address, error) {
	if s.Internal.StateGetID == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.StateGetID(p0, p1, p2)
}

func (s *FullNodeStub) StateGetID(p0 context.Context, p1 address.Address, p2 types.TipSetSelector) (*address.Address, error) {
	return nil, ErrNotSupported
}

func (s *FullNodeStruct) Web3ClientVersion(p0 context.Context) (string, error) {
	if s.Internal.Web3ClientVersion == nil {
		return "", ErrNotSupported
	}
	return s.Internal.Web3ClientVersion(p0)
}

func (s *FullNodeStub) Web3ClientVersion(p0 context.Context) (string, error) {
	return "", ErrNotSupported
}

var _ FullNode = new(FullNodeStruct)
