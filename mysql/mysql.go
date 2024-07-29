package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(handle func(db *gorm.DB) error) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.hort"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db, err := cli.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	return handle(cli)

}
