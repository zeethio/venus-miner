package main

import (
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/venus-miner/cli"
)

var recordCmd = &cli.Command{
	Name:      "record",
	Usage:     "records of the mining",
	Flags:     []cli.Flag{},
	ArgsUsage: "[address]",
	Action: func(cctx *cli.Context) error {
		postApi, closer, err := lcli.GetMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		addrStr := cctx.Args().First()
		addr, err := address.NewFromString(addrStr)
		if err != nil {
			return err
		}

		res, err := postApi.RecordsForMining(addr)
		if err != nil {
			return err
		}

		formatJson, err := json.MarshalIndent(res, "", "\t")
		if err != nil {
			return err
		}
		fmt.Println(string(formatJson))
		return nil

	},
}
