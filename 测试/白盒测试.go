package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"testing"
)

//测试的分类方式之一是基于对所要进行的测试的包的内部了解程度
//白盒测试可以访问包的内部函数和数据结构，并且可以做一些常规用户无法做到的观察和改动
//例如 白盒测试可以检查包的数据类型不可变性在每次操作后都是经过维护的
//白盒测试可以对现实的特定之外提供更详细的覆盖测试

func main()  {
	fmt.Println("白盒测试")
}

func bytesInUse(username string) int64 {
	return 0
}

const sender  = "notfications@example.com"
const pwd  = "ucsyhdcuysh"
const hostname  = "smtp.example.com"

const template  = `Warning: you are using %d bytes of stoage, %d%% of your quota.`

func CheckQuota(username string)  {
	used := bytesInUse(username)
	const quota   = 1000000000
	percent := 100*used/quota
	if percent<90{
		return
	}
	msg  := fmt.Sprintf(template,used,hostname)
	auth := smtp.PlainAuth("",sender,pwd, percent)
	err := smtp.SendMail(hostname+":587",auth,sender,[]string{username},[]byte(msg))
	if err != nil{
		log.Printf("smtp.SendMail(%s) failed: %s",username,err)
	}
}

//把发送邮件的逻辑移动到独立的函数中，并且把它存储到一个不可导出的包级别变量 notifyUser 中
var notifyUser = func(username,msg string) {
	smtp := smtp.PlainAuth("",sender,pwd,hostname)
	err := smtp.SendMail(hostname+":587",auth,sender,[]string{username},[]byte(msg))
	if err!=nil{
		log.Printf("smtp.SendMail(%s) failed: %s",username,err)
	}
}
func CheckQuota1(username string)  {
	used := bytesInUse(username)
	const quota   = 1000000000
	percent := 100*used/quota
	if percent<90{
		return
	}
	msg  := fmt.Sprintf(template,used,percent)
	notifyUser(username,msg)
}

//现在写个简单的测试，用伪造的通知机制而不是发送一封真实的邮件
//记录需要通知的用户和通知的内容
func TestCheckQuotaNotifiesUser(t *testing.T)  {
	var notifiedUser, notifiesMsg string
	notifyUser = func(user, msg string) {
		notifiedUser,notifiesMsg = user,msg
	}
	//模拟已使用980MB
	const user  = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiesMsg == ""{
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user{
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser,user)
	}
	const wantSubstring = "98% of your quota"
	if ! strings.Contains(notifiesMsg,wantSubstring){
		t.Errorf("unexpected notification message <<%s>>," +"want substring %q",notifiesMsg,wantSubstring)
	}
}
//必须在所有的测试执行路径上修改这个测试让它恢复 notifyUser 原来的值，后面的测试才不会受到影响，包括测试失败和宕机，
//通常这种情况下建议使用 defer
func TestCheckQuotaNotifiesUser1(t *testing.T){
	//保存留待恢复的 notifyUser
	saved := notifyUser
	defer func() {notifyUser = saved)}()

	//设置测试的伪通知 notifyUser
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser,notifiedMsg=user,msg
	}

	//测试其余部分
}
//这种方法可以用来临时保存并恢复各种全局变量，包括命令行选项 调试参数 以及性能参数
//也可以用来安装和移除钩子程序来让产品代码调用测试代码
//或者将产品代码设置为少见却很重要的状态，比如超时 错误 交叉并执行






