package email

import (
	"log"
	"net/smtp"
	"strconv"

	"github.com/jordan-wright/email"
)

func SendEmail(emailName string, uuid int64) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "dj <1152536969@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{emailName}
	//设置主题
	e.Subject = "Comfire Your Arthins Account Email Address"
	//设置文件发送的内容
	e.HTML = []byte("<!DOCTYPE html>" +
		"<html lang=\"en\">" +
		"<head><meta charset=\"UTF-8\">" +
		"<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">" +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">" +
		"<title>Welcome To Register ArtBlog</title>" +
		"</head>" +
		"<body>" +
		"<div style=\"background-color:#1976D2;padding:10px;\">" +
		"<section class=\"email_warp\">" +
		"<div id=\"reset-password-email\">" +
		"<h1 class=\"email-title\">" +
		"HEY THERE!" +
		"</h1>" +
		"<p>Welcome To Register ArtBlog! You're so close to using your first ArtBlog!</p>" +
		"<p class=\"cTMail-content\" style=\"font-size: 14px; color: rgb(51, 51, 51); line-height: 24px; margin: 6px 0px 0px; word-wrap: break-word; word-break: break-all;\">" +
		"<a id=\"cTMail-btn\" href=\"" +
		"http://8.134.9.37/#/claim?token=" + strconv.Itoa(int(uuid)) +
		"\" title=\"\"" +
		"style=\"font-size: 16px; line-height: 45px; display: block; background-color: rgb(20, 20, 20); color: rgb(255, 255, 255); text-align: center; text-decoration: none; margin-top: 20px; border-radius: 3px;\">" +
		"Click to Confirm Your Email Address + Set Things Up" +
		"</a>" +
		"</p>" +
		"</div>" +
		"</section>" +
		"</div>" +
		"</body>" +
		"</html>")
	// 设置服务器相关的配置
	// 阿里云服务器的限制，没有备案的服务器不能使用smtp服务
	// srdvhmxkaqeshbhj
	// qtmdlzgwzfmajbjj
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1152536969@qq.com", "rrbjqpiyqzjqgfbd", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
