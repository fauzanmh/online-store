package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/fauzanmh/online-store/pkg/util"
	_ "github.com/lib/pq" // postgres driver
	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

// CreatePGConnection return db connection instance
func CreatePGConnection(opts map[string]string) (*sql.DB, error) {
	port, err := strconv.Atoi(opts["port"])
	if err != nil {
		log.S().Fatal("Invalid port number : ", opts["port"])
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s  sslmode=%s",
		opts["host"], port, opts["user"], opts["password"], opts["dbname"], opts["sslmode"])

	if util.IsProductionEnv() {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslrootcert=./rds-combined-ca-bundle.pem sslmode=%s",
			opts["host"], port, opts["user"], opts["password"], opts["dbname"], opts["sslmode"])
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.S().Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.S().Fatal(err)
	}

	// Setting database connection config
	db.SetMaxOpenConns(viper.GetInt(`database.pg.max_open_connection`))
	db.SetMaxIdleConns(viper.GetInt(`database.pg.max_idle_connection`))
	db.SetConnMaxLifetime(viper.GetDuration(`database.pg.max_connection_lifetime`))

	log.S().Info("Connected to PG DB Server: ", opts["host"], " at port:", opts["port"], " successfully!")

	return db, nil
}
