package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var databaseURL string

func init() {

	databaseURL = os.Getenv("database")

}

func Connect() {
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		log.Panicln(err)
	}
	DB = db
}

func Setup() {
	db, err := gorm.Open("postgres", databaseURL)

	// db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic(err)
		// panic("failed to connect database")
	}
	defer db.Close()
	// Migrate the schema
	// db.AutoMigrate(&User{}, &Auth{}, &Institution{}, &Courses{}, &Files{}, &Section{}, &Task{})
	// db.AutoMigrate(&User{}, &Auth{}, &Institution{}, &Courses{}, &Files{}, &Section{}, &Task{})

	// db.AutoMigrate(&UsersSections{}, &FilesSections{}, &UsersInstitutions{})
	// db.AutoMigrate(&TaskAnswer{}, &TaskAnswerUserFriends{}, &TaskAnswerUserFiles{})
	// db.AutoMigrate(&Message{}, &MessageFiles{})
	db.AutoMigrate(&App{})

}
