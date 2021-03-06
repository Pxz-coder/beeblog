package controllers

import (
	"beeblog/models"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	if!checkAccount(c.Ctx){
		c.Redirect("/login",302)
		return
	}

	c.Data["IsLogin"]=checkAccount(c.Ctx)
	c.Data["IsTopic"]=true
	c.TplName = "topic.html"
	topics,err:=models.GetAllTopics("","",false)
	if err!=nil{
	    beego.Error(err)
	}else{
		c.Data["Topics"]=topics
	}
}

func(this *TopicController)Add(){
	if!checkAccount(this.Ctx){
		this.Redirect("/login",302)
		return
	}
	this.TplName="topic_add.html"
}


func(this *TopicController)Post(){
	if!checkAccount(this.Ctx){
		this.Redirect("/login",302)
		return
	}
	this.Data["IsLogin"]=checkAccount(this.Ctx)

	title:=this.Input().Get("title")
	content:=this.Input().Get("content")
    tid:=this.Input().Get("tid")
    category:=this.Input().Get("category")
	label:=this.Input().Get("label")
    //获取附件
    _,fh,err:=this.GetFile("attachment")
    if err!=nil{
    	beego.Error(err)
	}
    var attachment string
    if fh!=nil{
    	attachment =fh.Filename
    	beego.Info(attachment)
    	err=this.SaveToFile("attachment",path.Join("attachment",attachment))
    	if err!=nil{
    		beego.Error(err)
		}
	}

	if len(tid)==0{
		err =models.AddTopic(title,category,label,content,attachment)
	}else{
		err=models.ModifyTopic(tid,title,category,label,content,attachment)
	}

	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/topic",302)
}

func(this *TopicController)View(){

	this.TplName="topic_view.html"
	topic,err:=models.GetTopic(this.Ctx.Input.Param("0"))
	if err!=nil{
		beego.Error(err)
		this.Redirect("/",302)
		return
	}
	this.Data["Topic"]=topic


	this.Data["TopicContent"] = SwitchMarkdownToHtml(topic.Content)
	this.Data["Labels"]=strings.Split(topic.Labels," ")
	this.Data["Tid"]=this.Ctx.Input.Param("0")
	replies,err:=models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err!=nil{
		beego.Error(err)
		return
	}
	this.Data["Replies"]=replies
	this.Data["RepliesNum"]= len(replies)
	this.Data["IsLogin"]=checkAccount(this.Ctx)
}



func(this*TopicController)Modify(){
	if!checkAccount(this.Ctx){
		this.Redirect("/login",302)
		return
	}
	this.Data["IsLogin"]=checkAccount(this.Ctx)
	this.TplName="topic_modify.html"
	tid:=this.Input().Get("tid")
	topic,err:=models.GetTopic(tid)
	if err!=nil{
		beego.Error(err)
		this.Redirect("/",302)
		return
	}
	this.Data["Topic"]=topic
	this.Data["Tid"]=tid
}

func(this *TopicController)Delete(){
	if!checkAccount(this.Ctx){
		this.Redirect("/login",302)
		return
	}
	this.Data["IsLogin"]=checkAccount(this.Ctx)
	tid:=this.Input().Get("tid")
	err:=models.DeleteTopic(tid)
	if err!=nil{
		beego.Error(err)
	}
	this.Redirect("/",302)
}


//markdown处理文章显示
func SwitchMarkdownToHtml(content string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(content))

	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	/**
	对document进程查询，选择器和css的语法一样
	第一个参数：i是查询到的第几个元素
	第二个参数：selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))

	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}