package system

import (
	"flag"
)

// Config collects all gin configuration information.
//
// At the current time there is only one configuration item: port
//
// This struct has been configured with JSON and YAML struct tags.
type Config struct {
	Port uint `json:"port" yaml:"port"`
}

// AddFlagsToSet adds flags to the specified flag.FlagSet.
// Each flag is connected to a field in the configuration object.
func (cfg *Config) AddFlagsToSet(flags *flag.FlagSet) {
	flags.UintVar(&cfg.Port, "port", defaultUint(8080, cfg.Port), "specify server port number")
}

func defaultUint(dflt, cfg uint) uint {
	if cfg != 0 {
		return cfg
	} else {
		return dflt
	}
}
