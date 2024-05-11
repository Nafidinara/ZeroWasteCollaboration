package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	driver "redoocehub/drivers/mysql"
)

func NewMysqlDatabase(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DB_USER,
		env.DB_PASS,
		env.DB_HOST,
		env.DB_PORT,
		env.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	MigrateDatabase(db)

	return db
}

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(
		&driver.User{},
		&driver.Organization{},
		&driver.Address{},
		&driver.Collaboration{},
		&driver.Proposal{},
	)
}
