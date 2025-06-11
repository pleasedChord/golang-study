package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primarykey"`
	Name         string `gorm:"not null;unique"`
	ArticelCount int    `gorm:"default:0"`

	// 关联关系：一个用户有多篇文章(不会有实际的字段)
	Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID            uint   `gorm:"primarykey"`
	Title         string `gorm:"not null"`
	Content       string `gorm:"not null;type:text"`
	UserID        uint   `gorm:"not null;index"`
	CommentStatus string `gorm:"default:'有评论'"`

	// 关联关系：一个文章会有多条评论（不会有实际的字段）
	Comments []Comment `gorm:"foreignKey:PostID"`

	//反向关联，gorm可以通过预查询提高效率，获取user信息
	User User `gorm:"foreignKey:UserID"`
}

type Comment struct {
	ID      uint   `gorm:"primarykey"`
	Content string `gorm:"not null;type:text"`
	UserID  uint   `gorm:"not null;index"`
	PostID  uint   `gorm:"not null;index"`

	//反向关联，gorm可以通过预查询提高效率，获取user信息
	User User `gorm:"foreignKey:UserID"`
	//反向关联，gorm可以通过预查询提高效率，获取文章信息
	Post Post `gorm:"foreignKey:PostID"`
}

var DB *gorm.DB

func SeedData() {
	// 创建测试用户
	user := User{
		Name: "testuser",
	}

	DB.Create(&user)

	// 创建测试文章
	post := Post{
		Title:   "测试文章",
		Content: "这是一篇测试文章内容",
		UserID:  user.ID,
	}
	DB.Create(&post)

	// 创建测试评论
	comment := Comment{
		Content: "这是一条测试评论",
		UserID:  user.ID,
		PostID:  post.ID,
	}
	DB.Create(&comment)

	log.Println("示例数据创建成功")
}

// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息
func GetUserPostsWithComments(userID uint) ([]Post, error) {
	var posts []Post
	err := DB.Where("user_id = ?", userID).
		Preload("Comments").
		Preload("Comments.User").
		Find(&posts).Error
	return posts, err
}

// 获取评论最多的文章
func GetPostWithMostComments() (Post, error) {
	var post Post

	err := DB.Model(&Post{}).
		Select("posts.*, (select count(1) from comments where comments.post_id = posts.id) AS comment_count").
		Order("comment_count desc").
		Limit(1).
		Find(&post).Error

	return post, err
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("articel_count", gorm.Expr("articel_count + ?", 1)).
		Error
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func (c Comment) AferDelete(tx *gorm.DB) error {
	var count int64
	err := tx.Model(&Comment{}).Where("post_id", c.PostID).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).UpdateColumn("comment_status", "有评论").Error
	}

	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db连接失败")
		return
	}

	DB = db

	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalf("表迁移失败: %v", err)
		return
	}
	fmt.Println("表迁移成功！")

	// SeedData()

	posts, err := GetUserPostsWithComments(1)
	if err != nil {
		log.Fatalf("查询失败1: %v", err)
	}
	for _, post := range posts {
		fmt.Println("文章标题：", post.Title)
		for _, comment := range post.Comments {
			fmt.Println("评论：", comment.Content)
		}
	}

	post, err := GetPostWithMostComments()
	if err != nil {
		log.Fatalf("查询失败2: %v", err)
	}
	fmt.Println("数量最多的文章名称：", post.Title)

}
