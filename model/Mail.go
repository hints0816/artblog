package model

import (
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

func CheckvalidateCode(emailName string, uuid int64) int {

	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	redisEmail, err := redis.String(c.Do("Get", uuid))
	if err != nil {
		log.Fatal(err)
		return errormsg.ERROR
	}
	if redisEmail == "" {
		log.Fatal(err)
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}
