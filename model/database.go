package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	//数据库连接配置
	dsn := "root:@tcp(127.0.0.1:3306)/feerecord?charset=utf8mb4&parseTime=True&loc=Local"

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Println("连接数据库失败:", err.Error())
		return nil
	} else {
		log.Println("连接数据库成功")
		db.AutoMigrate(&FeeRecord{})
		DB = db
	}

	return DB
}

// func CloseDB() {
// 	defer db.Close()
// }

func Create(Record FeeRecord) error {

	if err := DB.Create(&Record).Error; err != nil {
		log.Print(err.Error())
		return err
	}

	return nil

}

func Update(data FeeRecord) error {
	// if err := DB.Update(&Record).Error; err != nil {
	// 	log.Print(err.Error())
	// }
	return nil
}

func Delete(Id uint) error {
	if err := DB.Delete(&FeeRecord{}, Id).Error; err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

func Search(Query FeeRecord) (data FeeRecord, err error) {
	// if err := DB.Create(&Record).Error; err != nil {
	// 	log.Print(err.Error())
	// }
	return data, nil
}
