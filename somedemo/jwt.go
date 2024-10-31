package somedemo

import (
	"fmt"
	// "sync"
	// "go-example/somedemo"
	"log"

	// "net/http"

	// "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 定义密钥
var jwtKey = []byte("godshop-jwt-key")

type AppClaims struct {
	AppName string `json:"apn,omitempty"`
	jwt.RegisteredClaims
}

func jwttest() {
	fmt.Println(jwt.GetAlgorithms())

	// 选择对称加密算法生成jwt，加解密使用的jwtKey相同
	// New method create a jwt instant with empty set of claims in it.
	// t := jwt.New(jwt.SigningMethodHS256)

	// jwt with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": "front",
			"userid":   123,
		},
	)

	token, err := t.SignedString(jwtKey)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(token)

	fmt.Println("-----------------")

	t0, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		log.Fatalln(err)
	}

	if claims, ok := t0.Claims.(jwt.MapClaims); ok {
		fmt.Println("username:", claims["username"])
	} else {
		log.Fatalln("can not parse claim map")
	}

	// router := gin.Default()

	// // router.Use(gin.SessionMiddleWare(gin.NewMemoryStore(), time.Hour*24))

	// router.POST("/login", func(c *gin.Context) {
	//     username := c.PostForm("username")
	//     password := c.PostForm("password")

	//     // 简单的身份验证
	//     if username == "admin" && password == "1234" {
	//         c.SetCookie("auth_token", "your_jwt_token", 3600, "/", "", false, true) // 设置 Cookie
	//         c.JSON(http.StatusOK, gin.H{"status": "success", "message": "登录成功"})
	//     } else {
	//         c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "用户名或密码错误"})
	//     }
	// })

	// router.GET("/dashboard", func(c *gin.Context) {
	//     user := c.MustGet("user").(string)
	//     token, _ := c.Cookie("auth_token")

	//     if user == "" || token == "" {
	//         c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "未授权的访问"})
	//         return
	//     }

	//     c.JSON(http.StatusOK, gin.H{"status": "success", "message": "欢迎回来", "user": user})
	// })

	// router.Run(":8080")
}
