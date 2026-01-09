package handle

import (
	"taskhub/internal/model"
	"taskhub/internal/pkg/password"
	"taskhub/internal/pkg/resp"
	"taskhub/internal/pkg/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
	TM *token.Manager
}

type registerReq struct {
	Email    string `json:"email" binding:"required,email" `
	Password string `json:"password" binding:"required,min=6,max=64" `
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Failed(c, 40001, "invalid params")
		return
	}
	// 查找数据库中 email是否存在
	var cnt int64
	h.DB.Model(&model.User{}).Where("email= ?", req.Email)
	if cnt > 0 {
		resp.Failed(c, 40901, "email already exist")
		return
	}

	hash, err := password.Hash(req.Password)
	if err != nil {
		resp.Failed(c, 50001, "internal error")
		return
	}

	u := model.User{
		Email: req.Email, PasswordHash: hash,
	}
	// 数据库中创表，写入登录数据
	if err := h.DB.Create(&u).Error; err != nil {
		resp.Failed(c, 50001, "add db  error")
		return
	}
	tk, err := h.TM.Sign(u.ID)
	if err != nil {
		resp.Failed(c, 50001, "sign error")
		return
	}
	resp.Ok(c, gin.H{
		"token": tk,
	})
}

type LoginReq struct {
	Email    string `json:"email" binding:"required,email" `
	Password string `json:"password" binding:"required,min=6,max=64" `
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Failed(c, 40001, "invalid params")
		return
	}

	var u model.User
	// 查找登录的email是否和数据库中的email一致
	if err := h.DB.Where("email = ?", req.Email).First(&u).Error; err != nil {
		resp.Failed(c, 40001, "invalid credentials")
		return
	}

	if !password.Verify(u.PasswordHash, req.Password) {
		resp.Failed(c, 40101, "invalid password")
		return
	}

	tk, err := h.TM.Sign(u.ID)
	if err != nil {
		resp.Failed(c, 50001, "intrenal error")
		return
	}
	resp.Ok(c, gin.H{
		"token": tk,
	})
}
