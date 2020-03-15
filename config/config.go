package config

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/spf13/viper"
)

const (
	VersionNum            = "0.1.0"
	configPath            = "./conf"
	configType            = "json"
	configName            = "config"
	FileAbsPath           = "./conf/config.json"
	dhtConfigKey          = "dht"
	fetchTorrentConfigKey = "fetchTorrent"
	pprofConfigKey        = "pprof"
	loggerConfigKey       = "logger"
)

var configMap = make(misc.Dict, 16)
var configLogger = misc.GetLogger().SetPrefix("config")

type dhtConfig struct {
	Port           int
	WorkPoolSize   int
	BootstrapNodes []string
}

type fetchTorrentConfig struct {
	InfoHashQueueSize int
	WorkPoolSize      int
}

type pprofConfig struct {
	NeedRun bool
}

type loggerConfig struct {
	Level string
}

func subConfig(key string, val interface{}) {
	configMap[key] = val
}

func init() {
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)
	viper.SetDefault(dhtConfigKey, &dhtConfig{
		Port:         21000,
		WorkPoolSize: 2000,
		BootstrapNodes: []string{
			"87.98.162.88:6881",
		},
	})
	viper.SetDefault(fetchTorrentConfigKey, &fetchTorrentConfig{
		InfoHashQueueSize: 300000,
		WorkPoolSize:      500,
	})
	viper.SetDefault(pprofConfigKey, &pprofConfig{
		NeedRun: true,
	})
	viper.SetDefault(loggerConfigKey, &loggerConfig{
		Level: "Info",
	})
	ResetConfig()
}

func ResetConfig() {
	viper.SetConfigPermissions(0666)
	configLogger.Trace("read config", misc.Dict{"path": viper.ConfigFileUsed(), "defaultPath": FileAbsPath})
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			misc.PanicSysErr("read config panic", err)
		}
		configLogger.Trace("config file not exist, will init with default", misc.Dict{"path": viper.ConfigFileUsed(), "defaultPath": FileAbsPath})
		if !misc.IsExist(configPath) {
			misc.Mkdir(configPath)
		}
		if err = viper.SafeWriteConfig(); err != nil {
			misc.PanicSysErr("init flush config err", err)
		}
		configLogger.Trace("write default config success", misc.Dict{"path": viper.ConfigFileUsed(), "defaultPath": FileAbsPath})
	}
	configLogger.Trace("config values", viper.AllSettings())
}

func DhtConfig() *dhtConfig {
	return getSubConfig(dhtConfigKey, new(dhtConfig)).(*dhtConfig)
}

func FetchTorrentConfig() *fetchTorrentConfig {
	return getSubConfig(fetchTorrentConfigKey, new(fetchTorrentConfig)).(*fetchTorrentConfig)
}

func PProfConfig() *pprofConfig {
	return getSubConfig(pprofConfigKey, new(pprofConfig)).(*pprofConfig)
}

func LoggerConfig() *loggerConfig {
	return getSubConfig(loggerConfigKey, new(loggerConfig)).(*loggerConfig)
}

func getSubConfig(key string, valPtr interface{}) interface{} {

	if val, ok := configMap[key]; ok {
		return val
	}
	err := viper.UnmarshalKey(key, valPtr)
	if err != nil {
		misc.PanicSysErr("cannot find "+key+" config", err)
	}
	configMap[key] = valPtr
	return valPtr
}
