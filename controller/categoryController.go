package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zbangbang/gin-vue-app/model"
	"zbangbang/gin-vue-app/repository"
	"zbangbang/gin-vue-app/response"
	"zbangbang/gin-vue-app/vo"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.DB.AutoMigrate(&model.Category{})

	return CategoryController{CategoryRepository: categoryRepository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var categoryRequest vo.CategoryRequest
	if err := ctx.ShouldBind(&categoryRequest); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	category, err := c.CategoryRepository.Create(categoryRequest.Name)
	if err != nil {
		panic(err)
		return
	}
	response.Success(ctx, gin.H{
		"category": category,
	}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var categoryRequest vo.CategoryRequest
	if err := ctx.ShouldBind(&categoryRequest); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	categoryId, _ := strconv.Atoi(ctx.Param("id"))

	selectCategory, err := c.CategoryRepository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	category, err := c.CategoryRepository.Update(*selectCategory, categoryRequest.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{
		"category": category,
	}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Param("id"))
	category, err := c.CategoryRepository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{
		"category": category,
	}, "获取成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Param("id"))
	err := c.CategoryRepository.DeleteById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}

	response.Success(ctx, nil, "删除成功")
}

func (c CategoryController) PageList(ctx *gin.Context) {
	panic("implement me")
}
