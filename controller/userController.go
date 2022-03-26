package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"zbangbang/gin-vue-app/common"
	"zbangbang/gin-vue-app/dto"
	"zbangbang/gin-vue-app/model"
	"zbangbang/gin-vue-app/response"
	"zbangbang/gin-vue-app/util"
)

// 注册
func Register(c *gin.Context) {
	db := common.GetDB()

	// 获取数据
	userName := c.PostForm("userName")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码必须是11位")
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}

	if len(userName) == 0 {
		strName := util.RandomString(11)
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, strName)
		return
	}

	if IsTelephoneExit(db, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	user := model.User{
		Model:     gorm.Model{},
		UserName:  userName,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	db.Create(&user)

	response.Success(c, nil, "注册成功")
}

// 登录
func Login(c *gin.Context) {
	db := common.GetDB()
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码必须是11位")
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}

	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(c, nil, "密码错误")
		return
	}

	// 返回token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统内部错误")
		log.Printf("token generate error: %v", err)
		return
	}

	response.Success(c, gin.H{
		"token": token,
	}, "登录成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	response.Success(c, gin.H{
		"user": dto.ToUserDto(user.(model.User)),
	}, "获取信息成功")
}

func IsTelephoneExit(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
