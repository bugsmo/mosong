package main

import (
	"mosong/webook/internal/repository"
	"mosong/webook/internal/repository/dao"
	"mosong/webook/internal/service"
	"mosong/webook/internal/web"
	"mosong/webook/internal/web/middleware"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	server := initWebServer()
	initUser(server, db)
	server.Run(":8080")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "moweilong.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	// 这里传入的两个 key，一个是数据的 authentication key
	// 另外一个是加密 key，用于保证 cookie 的安全性。
	store := cookie.NewStore([]byte("secret"))
	// cookie 的名字叫做ssid
	server.Use((sessions.Sessions("ssid", store)))
	// 登录校验
	login := &middleware.LoginMiddlewareBuilder{}
	server.Use(login.CheckLogin())

	return server
}

func initUser(server *gin.Engine, db *gorm.DB) {
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	c := web.NewUserHandler(us)
	c.RegisterRoutes(server)
}
