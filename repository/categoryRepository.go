package repository

import (
	"gorm.io/gorm"
	"zbangbang/gin-vue-app/common"
	"zbangbang/gin-vue-app/model"
)

// 增删改查复用
type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{
		DB: common.GetDB(),
	}
}

func (c CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}
	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (c CategoryRepository) SelectById(categoryId int) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, categoryId).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (c CategoryRepository) DeleteById(categoryId int) error {
	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		return err
	}

	return nil
}
