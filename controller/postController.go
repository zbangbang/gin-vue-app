package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"zbangbang/gin-vue-app/common"
	"zbangbang/gin-vue-app/model"
	"zbangbang/gin-vue-app/response"
	"zbangbang/gin-vue-app/vo"
)

type IPostController interface {
	RestController
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(&model.Post{})

	return PostController{DB: db}
}

func (p PostController) Create(ctx *gin.Context) {
	var postRequest vo.PostRequest
	if err := ctx.ShouldBind(&postRequest); err != nil {
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	user, _ := ctx.Get("user")

	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: postRequest.CategoryId,
		Title:      postRequest.Title,
		HeadImg:    postRequest.HeadImg,
		Content:    postRequest.Content,
	}

	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}

	response.Success(ctx, gin.H{
		"post": post,
	}, "创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var postRequest vo.PostRequest
	if err := ctx.ShouldBind(&postRequest); err != nil {
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 验证文章是否存在
	postId := ctx.Param("id")
	var post model.Post
	if err := p.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 权限判断
	user, _ := ctx.Get("user")
	if user.(model.User).ID != post.UserId {
		response.Fail(ctx, nil, "文章不属于您，无法操作")
		return
	}

	if err := p.DB.Model(&post).Updates(postRequest).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}

	response.Success(ctx, gin.H{
		"post": post,
	}, "更新成功")
}

func (p PostController) Show(ctx *gin.Context) {
	postId := ctx.Param("id")
	var post model.Post
	if err := p.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{
		"post": post,
	}, "文章查询成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	// 验证文章是否存在
	postId := ctx.Param("id")
	var post model.Post
	if err := p.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 权限判断
	user, _ := ctx.Get("user")
	if user.(model.User).ID != post.UserId {
		response.Fail(ctx, nil, "文章不属于您，无法操作")
		return
	}

	if err := p.DB.Delete(&post).Error; err != nil {
		response.Fail(ctx, nil, "删除文章失败")
		return
	}

	response.Success(ctx, nil, "删除文章成功")
}

func (p PostController) PageList(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	var posts []model.Post
	p.DB.Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&posts)

	var total int64
	p.DB.Model(&model.Post{}).Count(&total)

	response.Success(ctx, gin.H{
		"list":     posts,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}, "查询数据列表成功")
}
