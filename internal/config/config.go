package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type Config struct {
	Chains []ChainConfig
}

type ChainConfig struct {
	RPC            string         `json:"rpc"`
	Factory        common.Address `json:"factory"`
	RFQ            common.Address `json:"rfq"`
	UpdateInterval time.Duration  `json:"updateInterval"`
}

func ReadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "can't open config file")
	}
	defer file.Close()

	cfg := new(Config)
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse config")
	}
	return cfg, nil
}
