package store

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"log"
)

const (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

var db *gorm.DB

func InitDbConn() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbase, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("error : ", err.Error())
		panic(err)
	}
	db = dbase
	log.Println("database connection created")
	if !db.HasTable(&Image{}) {
		db = db.CreateTable(&Image{})
		if db.Error != nil {
			panic(err)
		}
	}
}

func init() {
	InitDbConn()
}

type Image struct {
	ImageID   int    `gorm:"primary_key;autoIncrement:true:image_id" json:"image_id"`
	ImageName string `gorm:"size:75;not null" json:"image_name"`
	CreatedAt string `gorm:"not null" json:"created_at"`
}
type QueryDetails struct {
	ImageId int
	Page    int
	Size    int
	Sort    string
	Order   string
}

func AddImage(imageDetails Image) (err error) {
	if db == nil {
		err = errors.New("unable to connect to database")
		log.Println("error : ", err.Error())
		return
	}
	var q *gorm.DB
	q = db
	q.LogMode(true)
	err = q.Create(&imageDetails).Error
	if err != nil {
		return err
	}
	return
}

func CloseDbConn() (err error) {
	err = db.Close()
	return
}

func GetImages(qParams QueryDetails) ([]Image, error) {
	var q *gorm.DB
	q = db
	var images []Image
	if qParams.ImageId > 0 {
		q = q.Where("image_id= ?", qParams.ImageId)
	}
	if qParams.Size > 0 {
		q = q.Limit(qParams.Size)
	}
	if qParams.Page > 0 {
		q = q.Offset(qParams.Page)
	}
	if qParams.Sort != "" && qParams.Order != "" {
		q = q.Order(qParams.Sort + " " + qParams.Order)
	}
	q.LogMode(true)
	q = q.Find(&images)
	if q.Error == gorm.ErrRecordNotFound {
		return images, errors.New("no record found")
	}
	return images, nil
}
