package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read config", err)
	}
	return
}

func initLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	return
}

// initPostgresInstance create new db conn postgres
func initPostgresInstance() (db *sql.DB, err error) {
	// create db conn object with params
	db, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.password"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.sslMode"),
	))
	if err != nil {
		return nil, errors.Wrap(err, "postgres sql open connect")
	}

	// db conn options
	db.SetMaxOpenConns(viper.GetInt("postgres.maxOpenConnect"))
	db.SetMaxIdleConns(viper.GetInt("postgres.maxIdleConnect"))
	db.SetConnMaxLifetime(viper.GetDuration("postgres.connLifetime") * time.Second)

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "postgres ping error")
	}

	log.Infof("Postgres connected on %d port", viper.GetInt("postgres.port"))

	return
}

