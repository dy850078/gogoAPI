package seed

import (
	"log"

	"github.com/dy850078/gogoAPI/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

var users = []models.User{
	{
		Acct:     "dysiang",
		Pwd:      "pwd",
		Fullname: "Siang Lin",
	},
	{
		Acct:     "mo820508",
		Pwd:      "pwd",
		Fullname: "Marco Yin",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
