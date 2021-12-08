package model

import (
	"fmt"
	"hello/utils/email"
	"hello/utils/errormsg"
	"log"

	"github.com/garyburd/redigo/redis"
)

type SignUpEmail struct {
	EmailName string `json:"email" binding:"required"`
	Uuid      int64  `json:"token"`
}

// 发送验证码
func SendvalidateCode(emailName string, uuid int64) int {

	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	_, err := c.Do("set", uuid, emailName, "EX", "1800")
	if err != nil {
		log.Fatal(err)
		return errormsg.ERROR
	}
	email.SendEmail(emailName, uuid)
	return errormsg.SUCCSE
}

func CheckCodeAndCreate(claim ClaimInfo) int {
	var userData User
	var profileData Profile
	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	redisEmail, _ := redis.String(c.Do("Get", claim.Token))
	if redisEmail != "" {
		userData.Password = claim.Password
		userData.Username = redisEmail
		userData.Role = 1
		code := CreateUser(&userData)

		profileData.ID = int(userData.ID)
		profileData.Name = redisEmail
		profileData.Avatar = "https://cdn.quasar.dev/logo-v2/svg/logo.svg"
		code1 := AddProfile(&profileData)

		if code == errormsg.SUCCSE && code1 == errormsg.SUCCSE {
			count, _ := redis.String(c.Do("Del", claim.Token))
			fmt.Println(count)
			return errormsg.SUCCSE
		}
		return errormsg.ERROR
	} else {
		return errormsg.ERROR
	}
}

func CheckvalidateCode(token string) string {
	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池
	redisEmail, _ := redis.String(c.Do("Get", token))
	return redisEmail
}
