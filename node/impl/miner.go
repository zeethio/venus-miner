package impl

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/venus-miner/miner"
	"github.com/filecoin-project/venus-miner/node/impl/common"
	"github.com/filecoin-project/venus-miner/node/modules/dtypes"
	"github.com/filecoin-project/venus-miner/node/modules/minermanage"
)

type MinerAPI struct {
	common.CommonAPI

	MiningStatistics minermanage.MiningStatisticsAPI
	miner.MiningAPI
}

func (m *MinerAPI) AddAddress(minerInfo dtypes.MinerInfo) error {

	return m.MiningAPI.AddAddress(minerInfo)
}

func (m *MinerAPI) UpdateAddress(minerInfo dtypes.MinerInfo) error {

	return m.MiningAPI.UpdateAddress(minerInfo)
}

func (m *MinerAPI) RemoveAddress(addrs []address.Address) error {
	return m.MiningAPI.RemoveAddress(addrs)
}

func (m *MinerAPI) ListAddress() ([]dtypes.MinerInfo, error) {
	return m.MiningAPI.ListAddress()
}

func (m *MinerAPI) StatesForMining(addrs []address.Address) ([]dtypes.MinerState, error) {
	return m.MiningAPI.StatesForMining(addrs)
}

func (m *MinerAPI) RecordsForMining(addr address.Address) ([]dtypes.MiningRecord, error) {
	return m.MiningStatistics.Get(addr)
}

func (m *MinerAPI) Start(ctx context.Context, addr address.Address) error {
	return m.MiningAPI.ManualStart(ctx, addr)
}

func (m *MinerAPI) Stop(ctx context.Context, addr address.Address) error {
	return m.MiningAPI.ManualStop(ctx, addr)
}
