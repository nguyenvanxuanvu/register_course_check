package mysql

import (
	"context"
	"log"
	"time"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
	
	"fmt"
	"go.uber.org/fx"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(lifecycle fx.Lifecycle) *sqlx.DB {
	db := NewMySQL("mysql")
	lifecycle.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		log.Println("Closing DB")
		return db.Close()
	}})
	return db
}


func NewMySQL(configName string) *sqlx.DB {
	username := viper.GetString(fmt.Sprintf("%s.username", configName))
	password := viper.GetString(fmt.Sprintf("%s.password", configName))
	url := viper.GetString(fmt.Sprintf("%s.url", configName))
	schema := viper.GetString(fmt.Sprintf("%s.schema", configName))
	log.Println(username, password, url, schema)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s", username, password, url, schema)

	log.Println("Connecting to database")
	db, err := otelsqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("Cannot init mysql connection ", err)
	}

	// Maximum Idle Connections
	db.SetMaxIdleConns(12)
	// Maximum Open Connections
	db.SetMaxOpenConns(24)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(600000 * time.Millisecond)
	// Connection Lifetime
	db.SetConnMaxLifetime(1800000 * time.Millisecond)

	log.Printf("Connect to database %s successfully\n", configName)
	return db
}

