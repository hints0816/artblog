package email

import (
	"log"
	"net/smtp"

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
	e.Text = []byte("<!DOCTYPE html><html lang=\"en\">" +
		"<head><meta charset=\"UTF-8\">" +
		"<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">" +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">" +
		"<title>重置您的 无敌软件的 密码</title>" +

		"<style>" +
		"body,html,div,ul,li,button,p,img,h1,h2,h3,h4,h5,h6 {" +
		"margin: 0;" +
		"padding: 0;" +
		"}" +

		"body,html {" +
		"background: #fff;" +
		"line-height: 1.8;" +
		"}" +

		"h1,h2,h3,h4,h5,h6 {" +
		"line-height: 1.8;" +
		"}" +

		".email_warp {" +
		"height: 100vh;" +
		"min-height: 500px;" +
		"font-size: 14px;" +
		"color: #212121;" +
		"display: flex;" +
		"justify-content: center;" +
		"}" +

		"h1.email-title {" +
		"font-size: 26px;" +
		"font-weight: 500;" +
		"margin-bottom: 15px;" +
		"color: #252525;" +
		"}" +

		"a.links_btn {" +
		"border: 0;" +
		"background: #4C84FF;" +
		"color: #fff;" +
		"width: 100%;" +
		"height: 50px;" +
		"line-height: 50px;" +
		"font-size: 16px;" +
		"margin: 40px auto;" +
		"box-shadow: 0px 2px 4px 0px rgba(0, 0, 0, 0.15);" +
		"border-radius: 4px;" +
		"outline: none;" +
		"cursor: pointer;" +
		"transition: all 0.3s;" +
		"text-align: center;" +
		"display: block;" +
		"text-decoration: none;" +
		"}" +

		"#reset-password-email {" +
		"max-width: 500px;" +
		"}" +
		"#reset-password-email .accout_email {" +
		"color: #4C84FF;" +
		"display: block;" +
		"margin-bottom: 20px;" +
		"}" +
		"</style>" +
		"</head>" +

		"<body>" +
		"<section class=\"email_warp\">" +
		"<div id=\"reset-password-email\">" +

		"<h1 class=\"email-title\">" +
		"尊敬的<span>AAA</span>您好：" +
		"</h1>" +
		"<p>您正在为登录邮箱为如下地址的 啦啦啦 账户重置密码：</p>" +
		"<b class=\"accout_email\">xxxx@abc.com</b>" +

		"<p>请注意，如果这不是您本人的操作，请忽略并关闭此邮件。</p>" +
		"<p>如您确认重置 XXXX 的账户密码，请点击下方按钮。</p>" +

		"<a class=\"links_btn\" onclick=\"window.open('https:XXXXXXXXXXX')\">重置密码</a>" +
		"</div>" +
		"</section>" +
		"</body>" +

		"</html>")
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1152536969@qq.com", "srdvhmxkaqeshbhj", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
