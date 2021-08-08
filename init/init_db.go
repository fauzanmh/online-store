package init

import (
	"database/sql"
	"os"

	"github.com/fauzanmh/online-store/internal/db"
	"github.com/fauzanmh/online-store/pkg/util"
	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

// ConnectToPGServer is a function to init PostgreSQL connection
func ConnectToPGServer(cfg *Config) (*sql.DB, error) {
	if util.IsProductionEnv() && (!viper.IsSet("database.pg.password") || viper.GetString("database.pg.password") == "") {
		log.S().Fatal("database.pg.password can not be empty!")
	}

	dbpg, err := db.CreatePGConnection(map[string]string{
		"host":     cfg.Database.Pg.Host,
		"port":     cfg.Database.Pg.Port,
		"user":     cfg.Database.Pg.User,
		"password": cfg.Database.Pg.Password,
		"dbname":   cfg.Database.Pg.Dbname,
		"sslmode":  cfg.Database.Pg.SslMode,
	})

	if err != nil {
		os.Exit(1)
	}

	return dbpg, err
}
