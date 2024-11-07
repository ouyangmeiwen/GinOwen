package test

import (
	"GINOWEN/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义 User 结构体作为数据模型
type User struct {
	ID       uint   `gorm:"primaryKey"`  // 主键
	Name     string `gorm:"size:100"`    // 字符串类型字段
	Email    string `gorm:"uniqueIndex"` // 唯一索引
	Password string `gorm:"size:255"`    // 密码字段
}

// 初始化数据库并创建表
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

var DB *gorm.DB

// 初始化数据库
func InitDB() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 自动迁移（AutoMigrate）
	Migrate(DB)
}

// 创建一个新的用户
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// 根据用户名查找用户
func GetUserByName(c *gin.Context) {
	var user models.User
	name := c.Query("name")
	if err := DB.Where("name = ?", name).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// 获取多个符合条件的用户
func GetUsersByCondition(c *gin.Context) {
	var users []models.User
	if err := DB.Where("email LIKE ?", "%@example.com%").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// 根据 ID 更新用户信息
func UpdateUserEmail(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	newEmail := c.Query("email")

	if err := DB.Model(&user).Where("id = ?", id).Update("email", newEmail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// 删除用户
func DeleteUser1(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User deleted successfully"})
}

// 获取所有用户，按创建时间降序排列
func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := DB.Order("created_at desc").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// 获取最新的一个用户记录
func GetLatestUser(c *gin.Context) {
	var user models.User
	if err := DB.Order("created_at desc").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// 批量更新：将所有 email 为空的用户的 email 设置为 "unknown@example.com"
func UpdateUsersEmail(c *gin.Context) {
	result := DB.Model(&models.User{}).Where("email = ?", "").Update("email", "unknown@example.com")
	c.JSON(http.StatusOK, gin.H{"rows_affected": result.RowsAffected})
}

// 在事务中执行批量操作
func UpdateUserEmailTransaction(c *gin.Context) {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user models.User
	if err := tx.Model(&user).Where("id = ?", 1).Update("email", "newemail1@example.com").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Model(&user).Where("id = ?", 2).Update("email", "newemail2@example.com").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"data": "Transaction committed"})
}

// 使用原生 SQL 查询
func RawSQLQuery(c *gin.Context) {
	var users []models.User
	DB.Raw("SELECT * FROM users WHERE email LIKE ?", "%@example.com%").Scan(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// 获取分页用户数据
func GetUsersPaginated(c *gin.Context) {
	var users []models.User
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	offset := (page - 1) * limit
	DB.Offset(offset).Limit(limit).Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// 获取用户总数
func GetUserCount(c *gin.Context) {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"total_users": count})
}

// 多条件查询用户
func GetUsersByMultipleConditions(c *gin.Context) {
	var users []models.User
	DB.Where("email LIKE ?", "%@example.com%").Or("name = ?", "John Doe").Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
