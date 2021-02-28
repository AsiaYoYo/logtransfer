package conf

// LogTransferCfg 全局配置
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	EsCfg    `ini:"es"`
}

// KafkaCfg ...
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// EsCfg ...
type EsCfg struct {
	Address string `ini:"address"`
	MaxSize int    `ini:"chan_max_size"`
}
