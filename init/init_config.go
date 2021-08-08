package init

import (
	"os"
	"strings"
	"time"

	"github.com/fauzanmh/online-store/pkg/util"
	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

type Config struct {
	API struct {
		HTTP struct {
			Port string `mapstructure:"port"`
		} `mapstructure:"http"`
	} `mapstructure:"api"`
	Context struct {
		Timeout int `mapstructure:"timeout"`
	} `mapstructure:"context"`
	Database struct {
		Pg struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Dbname   string `mapstructure:"dbname"`
			SslMode  string `mapstructure:"sslmode"`
		} `mapstructure:"pg"`
	} `mapstructure:"database"`
}

// setupMainConfig loads app config to viper
func setupMainConfig() (config *Config) {
	log.S().Info("Executing init/config")

	conf := false

	if util.IsDevelopmentEnv() {
		conf = true
		viper.AddRemoteProvider("consul", os.Getenv("CONFIG_ADDRESS"), os.Getenv("CONFIG_PATH"))
		viper.SetConfigType("json")
		err := viper.ReadRemoteConfig()
		if err != nil {
			log.S().Info("err: ", err)
		}
	}

	if util.IsProductionEnv() {
		conf = true
		log.S().Info("prod config")
		viper.SetConfigFile("config/app/production.json")
		err := viper.ReadInConfig()
		if err != nil {
			log.S().Info("err: ", err)
		}
	}

	if util.IsFileorDirExist("main.json") {
		conf = true
		log.S().Info("Local main.json file is found, now assigning it with default config")
		viper.SetConfigFile("main.json")
		err := viper.ReadInConfig()
		if err != nil {
			log.S().Info("err: ", err)
		}
	}

	if !conf {
		log.S().Fatal("Config is required")
	}

	viper.SetEnvPrefix(`app`)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	err := viper.Unmarshal(&config)
	if err != nil {
		log.S().Fatal("err: ", err)
	}

	log.S().Info("Config - APP_ENV: ", util.GetEnv())

	if !util.IsFileorDirExist("main.json") && !util.IsProductionEnv() {
		// open a goroutine to watch remote changes forever
		go func() {
			for {
				time.Sleep(time.Second * 5)

				err := viper.WatchRemoteConfig()
				if err != nil {
					log.S().Errorf("unable to read remote config: %v", err)
					continue
				}

				// unmarshal new config into our runtime config struct. you can also use channel
				// to implement a signal to notify the system of the changes
				viper.Unmarshal(&config)
			}
		}()
	}

	return
}
