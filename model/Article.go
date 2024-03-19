package model

import (
	"fmt"
	"hello/utils/errormsg"
	"log"
	"strconv"
	"gorm.io/gorm"
	"github.com/garyburd/redigo/redis"
)

type Article struct {
	gorm.Model
	Cateart      []*Cateart `gorm:"foreignkey:ID;references:ID"`
	Title        string     `gorm:"type:varchar(100);not null" json:"title"`
	Desc         string     `gorm:"type:varchar(200)" json:"desc"`
	Content      string     `gorm:"type:longtext" json:"content"`
	Img          string     `gorm:"type:varchar(100)" json:"img"`
	CommentCount int        `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int        `gorm:"type:int;not null;default:0" json:"read_count"`
	Status       int8       `gorm:"type:tinyint;default:0" json:"status"`
	UserID       uint       `gorm:"type:bigint;not null" json:"user_id"`
	Profile      Profile    `gorm:"foreignkey:UserID"`
	Favour       int8      	`gorm:"type:tinyint;default:0" json:"favour"`
	FavourArr    []Favour   `gorm:"-"`
	ViewCount    int32      `gorm:"type:int;not null;default:0" json:"view_count"`
}

type Favour struct {
	ID        uint `gorm:"-"`
	ArticleId uint `gorm:"-" json:"article_id"`
}

type Cateart struct {
	Category Category `gorm:"foreignkey:Cid;references:ID"`
	ID       int      `gorm:"primary_key;not null" json:"id"`
	Cid      int      `gorm:"primary_key;not null" json:"cid"`
}

// CreateArt 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		log.Fatalln(err)
		return errormsg.ERROR // 500
	}
	return errormsg.SUCCSE
}

// CreateArt 更新文章
func UpdateArt(data *Article) int {
	err = db.Model(&Article{}).Where("ID = ?", data.ID).Updates(&data).Error
	careart := data.Cateart
	err1 := db.Delete(&careart).Error
	fmt.Println(err1)
	err2 := db.Create(&careart).Error
	fmt.Println(err2)
	if err != nil {
		log.Fatalln(err)
		return errormsg.ERROR // 500
	}
	return errormsg.SUCCSE
}

func DelArt(id uint) int {
	var article Article
	article.ID = id
	err = db.Delete(&article).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// GetArtInfo 查询单个文章
func GetArtInfo(user_id uint, id int) (Article, int) {
	var art Article
	c := pool.Get()
	defer c.Close()
	db.Preload("Cateart").Preload("Cateart.Category").Where("id =?", id).First(&art)
	num, _ := redis.Int(c.Do("bitcount", "article_like:"+string(strconv.Itoa(int(id)))))
	fmt.Println(user_id)
	if user_id != 0 {
		favourNum, _ := redis.Int(c.Do("getbit", "article_like:"+string(strconv.Itoa(int(id))), user_id))
		var favourArr Favour
		if favourNum == 1 {
			favourArr.ArticleId = uint(id)
			favourArr.ID = user_id
			art.FavourArr = []Favour{favourArr}
		} else {
			art.FavourArr = []Favour{}
		}
	} else {
		art.FavourArr = []Favour{}
	}
	art.Favour = int8(num)
	return art, errormsg.SUCCSE
}

func UpdateArtView(id int, count int) {
	var art Article
	db.Model(&art).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + ?", count))
}

func ListArticle(edit bool, userId int, cid int, keyword string, pageSize int, pageNum int) ([]Article, int, int64) {
	var artList []Article
	var ids []string
	var total int64

	if cid != 0 {
		// 检索出这个标签ID对应的所有文章ID
		db.Model(&Cateart{}).Where("cid =?", cid).Pluck("id", &ids)
	}
	// preload预加载一对多的数据
	tx := db.
		Order("created_at DESC").
		Preload("Profile").
		Preload("Cateart").
		Preload("Cateart.Category")
	if !edit {
		tx = tx.Where("status =?", 1)
	}
	if cid != 0 {
		tx = tx.Where("id in ?", ids)
	}
	if userId != 0 {
		tx = tx.Where("user_id = ?", userId)
	}
	if keyword != "" {
		tx = tx.Where("title like ? or content like ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	
	err = tx.Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&artList).Error
	tx2 := db.Model(&artList)
	if !edit {
		tx2 = tx2.Where("status =?", 1)
	}
	if cid != 0 {
		tx2 = tx2.Where("id in ?", ids)
	}
	if userId != 0 {
		tx2 = tx2.Where("user_id = ?", userId)
	}
	if keyword != "" {
		tx2 = tx2.Where("title like ? or content like ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	tx2.Count(&total)

	if err != nil {
		return nil, errormsg.ERROR_CATE_NOT_EXIST, 0
	}
	return artList, errormsg.SUCCSE, total
}

func ListArticleRepository(userId uint, id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var artList []Article
	var ids []string
	var total int64

	if id != 0 {
		db.Model(&Cateart{}).Where("cid =?", id).Pluck("id", &ids)
	}
	tx := db.
		Order("created_at DESC").
		Preload("Cateart").
		Preload("Cateart.Category").
		Where("user_id =?", userId)
	if id != 0 {
		tx = tx.Where("id in ?", ids)
	}
	err = tx.Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&artList).Error
	if id != 0 {
		db.Model(&artList).Where("user_id =?", userId).Where("id in ?", ids).Count(&total)
	} else {
		db.Model(&artList).Where("user_id =?", userId).Count(&total)
	}

	if err != nil {
		return nil, errormsg.ERROR_CATE_NOT_EXIST, 0
	}
	return artList, errormsg.SUCCSE, total
}

func SetActicleFavour(favour *Favour) int {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("setbit", "article_like:"+string(strconv.Itoa(int(favour.ArticleId))), favour.ID, 1)
	if err != nil {
		fmt.Println(err)
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func SetActicleUnFavour(favour *Favour) int {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("setbit", "article_like:"+string(strconv.Itoa(int(favour.ArticleId))), favour.ID, 0)
	if err != nil {
		fmt.Println(err)
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}