package service

import (
	"log"
	"shoollove/config"

	"gopkg.in/gomail.v2"
)

func SendRegisterEamil(code string, reciveEmail string) error {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", config.Config.Email.Name)

	// 接受人
	m.SetHeader("To", reciveEmail)
	//m.SetAddressHeader("To", "dan@example.com", "Dan")

	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "注册账户服务")
	//内容
	m.SetBody("text/html", "【星咔】您正在注册该服务，验证码为"+code+"。请勿提供给他人。")
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(config.Config.Email.Host, config.Config.Email.Port, config.Config.Email.Name, config.Config.Email.Password)
	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		log.Printf("发送邮件到reciveEmail失败:%v", err)
		return err
	}
	return nil
}
