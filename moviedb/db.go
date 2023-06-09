package moviedb

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	MovieName  string `json:"movie_name" binding:"required"`
	MovieUrl   string `json:"movie_url" binding:"required"`
	MovieCover string `json:"movie_cover"`
}

type Page struct {
	PageNum  int  `form:"page_num"`
	PageSize int  `form:"page_size"`
	Keyword  int  `form:"keyword"`
	Desc     bool `form:"desc"`
}

var db *gorm.DB

func init() {

	// err
	var err error

	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("执行了吗")

	// 迁移 schema
	db.AutoMigrate(&Movie{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// // Read
	// var product Product
	// db.First(&product, 1)                 // 根据整型主键查找
	// db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// // Update - 将 product 的 price 更新为 200
	// db.Model(&product).Update("Price", 200)
	// // Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}

func CreateOrUpdate(movie *Movie) {

	var m Movie

	if err := db.Where("movie_name = ?", movie.MovieName).First(&m).Error; err == nil {
		db.Model(&m).Updates(
			Movie{
				MovieUrl:   m.MovieUrl,
				MovieCover: m.MovieCover,
			},
		)
		return
	}

	db.Create(movie)
}

func GetFromName(id string) (Movie, error) {

	var m Movie

	if err := db.Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return m, err
		}

		return m, err
	}

	return m, nil
}

func GetMovieList(p Page) ([]Movie, error) {

	var m []Movie

	fmt.Printf("%#v \n", p)

	if err := db.Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Find(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return m, err
		}

		return m, err
	}

	return m, nil
}
