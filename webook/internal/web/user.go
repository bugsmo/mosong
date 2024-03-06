package web

import (
	"mosong/webook/internal/domain"
	"mosong/webook/internal/service"
	"net/http"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const userIdKey = "userId"

const emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"

// 和上面比起来，用 ` 看起来就比较清爽
// ?= 断言是否包含字母、数字、特殊符号，至少8个字符
const passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`

type UserHandler struct {
	svc              *service.UserService
	emailRegexExp    *regexp2.Regexp
	passwordRegexExp *regexp2.Regexp
	jwtKey           string // 只有在使用 JWT 的时候才有用
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc:              svc,
		emailRegexExp:    regexp2.MustCompile(emailRegexPattern, regexp2.None),
		passwordRegexExp: regexp2.MustCompile(passwordRegexPattern, regexp2.None),
	}
}

func (c *UserHandler) RegisterRoutes(server *gin.Engine) {
	// 直接注册
	// server.POST("/users/signup", u.SignUp)
	// server.POST("/users/login", u.Login)
	// server.POST("users/edit", u.Edit)
	// server.GET("/users/profile", u.Profile)

	// 分组注册
	ug := server.Group("/users")
	ug.POST("/signup", c.SignUp)
	// session 机制
	// ug.POST("/login", c.Login)
	// JWT 机制
	ug.POST("/login", c.LoginJWT)
	ug.POST("/edit", c.Edit)
	// ug.GET("/profile", c.Profile)
	ug.GET("/profile", c.ProfileJWT)
}

// SignUp 用户注册接口
func (c *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq
	// 当我们调用 Bind 方法的时候, 如果有问题, Bind 方法已经直接写响应回去了
	if err := ctx.Bind(&req); err != nil {
		return
	}

	isEmail, err := c.emailRegexExp.MatchString(req.Email)
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

	isPassword, err := c.passwordRegexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !isPassword {
		ctx.String(http.StatusOK, "密码必须包含字母、数字、特殊字符，并且长度不能小于 8 位")
		return
	}

	err = c.svc.Signup(ctx.Request.Context(), domain.User{
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

// LoginJWT 用户登录接口，使用的是 JWT，如果你想要测试 JWT，就启用这个
func (c *UserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq

	if err := ctx.Bind(&req); err != nil {
		return
	}
	u, err := c.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名或者密码不正确，请重试")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		Id:        u.Id,
		UserAgent: ctx.GetHeader("User-Agent"),
		RegisteredClaims: jwt.RegisteredClaims{
			// 演示目的设置为一分钟过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		},
	})
	tokenStr, err := token.SignedString(JWTKey)
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.Header("x-jwt-token", tokenStr)
	ctx.String(http.StatusOK, "登录成功")
}

// Login 用户登录接口
func (c *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq

	if err := ctx.Bind(&req); err != nil {
		return
	}
	u, err := c.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名或者密码不正确，请重试")
		return
	}

	sess := sessions.Default(ctx)
	sess.Set(userIdKey, u.Id)
	sess.Options(sessions.Options{
		MaxAge: 60, // 60 秒过期
	})
	err = sess.Save()
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}
	ctx.String(http.StatusOK, "登录成功")
}

// Edit 用户编辑信息
func (c *UserHandler) Edit(ctx *gin.Context) {}

// Profile 用户详情, JWT 版本
func (c *UserHandler) ProfileJWT(ctx *gin.Context) {
	type Profile struct {
		Email string
	}

	uc := ctx.MustGet("user").(UserClaims)
	u, err := c.svc.Profile(ctx, uc.Id)
	if err != nil {
		// 按照道理来说，这边 id 对应的数据肯定存在，所以要是没找到，
		// 那就说明是系统出了问题。
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	ctx.JSON(http.StatusOK, Profile{
		Email: u.Email,
	})
}

// Profile 用户详情
func (c *UserHandler) Profile(ctx *gin.Context) {
	type Profile struct {
		Email string
	}

	sess := sessions.Default(ctx)
	id := sess.Get(userIdKey).(int64)

	u, err := c.svc.Profile(ctx, id)
	if err != nil {
		// 按照道理来说，这边 id 对应的数据肯定存在，所以要是没找到，
		// 那就说明是系统出了问题。
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.JSON(http.StatusOK, Profile{
		Email: u.Email,
	})
}
