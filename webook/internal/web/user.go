package web

import (
	"mosong/webook/internal/domain"
	"mosong/webook/internal/service"
	"net/http"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

const emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"

// 和上面比起来，用 ` 看起来就比较清爽
// ?= 断言是否包含字母、数字、特殊符号，至少8个字符
const passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`

type UserHandler struct {
	svc              *service.UserService
	emailRegexExp    *regexp2.Regexp
	passwordRegexExp *regexp2.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc:              svc,
		emailRegexExp:    regexp2.MustCompile(emailRegexPattern, regexp2.None),
		passwordRegexExp: regexp2.MustCompile(passwordRegexPattern, regexp2.None),
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	// 直接注册
	// server.POST("/users/signup", u.SignUp)
	// server.POST("/users/login", u.Login)
	// server.POST("users/edit", u.Edit)
	// server.GET("/users/profile", u.Profile)

	// 分组注册
	ug := server.Group("/users")
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

// SignUp 用户注册接口
func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	var req SignUpReq
	// 当我们调用 Bind 方法的时候, 如果有问题, Bind 方法已经直接写响应回去了
	if err := ctx.Bind(&req); err != nil {
		return
	}

	isEmail, err := u.emailRegexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱不正确")
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次输入的密码不相同")
		return
	}

	isPassword, err := u.passwordRegexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !isPassword {
		ctx.String(http.StatusOK, "密码必须包含字母、数字、特殊字符，并且长度不能小于 8 位")
		return
	}

	err = u.svc.Signup(ctx.Request.Context(), domain.User{
		Email:    req.Email,
		Password: req.ConfirmPassword,
	})

	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "重复邮箱，请换一个邮箱")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "服务器异常，注册失败")
		return
	}

	ctx.String(http.StatusOK, "Hello, 注册成功")
}

// Login 用户登录接口
func (u *UserHandler) Login(ctx *gin.Context) {}

// Edit 用户编辑信息
func (u *UserHandler) Edit(ctx *gin.Context) {}

// Profile 用户详情
func (u *UserHandler) Profile(ctx *gin.Context) {}
