package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return errors.New("invalid duration")
	}
}

type Config struct {
	Chains []ChainConfig
}

type ChainConfig struct {
	RPC            string         `json:"rpc"`
	Factory        common.Address `json:"factory"`
	RFQ            common.Address `json:"rfq"`
	UpdateInterval Duration       `json:"updateInterval"`
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
