package vo

// shouldBind 绑定结构
type CategoryRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}
