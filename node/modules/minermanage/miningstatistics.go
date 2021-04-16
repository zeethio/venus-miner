package minermanage

import (
	"encoding/json"
	"fmt"

	"github.com/ipfs/go-datastore"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/venus-miner/node/modules/dtypes"
)

type MiningStatisticsAPI interface {
	Put(addr address.Address, mr *dtypes.MiningRecord) error
	Get(addr address.Address) ([]dtypes.MiningRecord, error)
	Del(addr address.Address) error
}

type MiningStatistics struct {
	ds dtypes.MetadataDS
}

func NewMiningStatistics(ds dtypes.MetadataDS) (*MiningStatistics, error) {
	return &MiningStatistics{ds: ds}, nil
}

func (m *MiningStatistics) Put(addr address.Address, mr *dtypes.MiningRecord) error {
	mrs, err := m.Get(addr)
	if err != nil {
		return err
	}

	mrs = append(mrs, *mr)
	addrBytes, err := json.Marshal(mrs)
	if err != nil {
		return err
	}

	key := datastore.NewKey(fmt.Sprintf("/%s", addr))
	return m.ds.Put(key, addrBytes)
}

func (m *MiningStatistics) Get(addr address.Address) ([]dtypes.MiningRecord, error) {
	key := datastore.NewKey(fmt.Sprintf("/%s", addr))

	bytes, err := m.ds.Get(key)
	if err != nil {
		return nil, err
	}

	var res []dtypes.MiningRecord
	err = json.Unmarshal(bytes, &res)
	return res, err
}

func (m *MiningStatistics) Del(addr address.Address) error {
	key := datastore.NewKey(fmt.Sprintf("/%s", addr))

	return m.ds.Delete(key)
}

var _ MiningStatisticsAPI = &MiningStatistics{}
