package db

import (
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/payaldoultani/go-crud/config"
)

func InitDB(cfg *config.Mysql) error {
	mysqlDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME,
	)

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return fmt.Errorf("failed to register MySQL driver: %w", err)
	}

	err = orm.RegisterDataBase("default", "mysql", mysqlDSN)
	if err != nil {
		return fmt.Errorf("failed to register database: %w", err)
	}

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		return fmt.Errorf("failed to synchronize database schema: %w", err)
	}

	log.Println("Database connection established and schema synchronized successfully!")
	return nil
}
