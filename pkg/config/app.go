package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func Connect() {
	//const dsn = "host=localhost user=postgres password=root dbname=test-db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	const dsn = "user=shrohithbg password=0Nrpds1mKMFO dbname=neondb host=ep-aged-block-879966.ap-southeast-1.aws.neon.tech sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
