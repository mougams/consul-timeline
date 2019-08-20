package consul

import "flag"

type Config struct {
	Address               string `json:"address"`
	Token                 string `json:"token"`
	EnableDistributedLock bool   `json:"enable_distributed_lock"`
	LockPath              string `json:"lock_path"`
}

var flagConfig Config

func init() {
	flag.StringVar(&flagConfig.Address, "consul", "localhost:8500", "Consul agent address")
	flag.StringVar(&flagConfig.Token, "consul-token", "", "Consul ACL token")
	flag.BoolVar(&flagConfig.EnableDistributedLock, "consul-enable-distributed-lcok", false, "Multi timeline instance lock for storage")
	flag.StringVar(&flagConfig.LockPath, "consul-lock-path", "consul_timeline/lock", "Consul lock path")
}

func ConfigFromFlags() Config {
	return flagConfig
}
