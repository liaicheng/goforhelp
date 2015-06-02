package main

import (
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"./api"
	"./service"
	"net/http"
)

func main() {
	// martini default
	m := martini.Classic()

	// 配合前端测试时可将下面路径改为前端文件所在路径，并取消注释
	static := martini.Static("../../../webclient_code/admin", martini.StaticOptions{Fallback: "/index.html", Exclude: "/api"})
	m.NotFound(static, http.NotFound)

	// mgo config
	mgoSession, _ := mgo.Dial("localhost")
	defer mgoSession.Close()
	mgoSession.SetMode(mgo.Monotonic, true)
	m.Map(mgoSession.DB("irctoryperdb"))
	
	// cookie&session config
	sessionStore := sessions.NewCookieStore([]byte("jf2s5aSf"))
	sessionStore.Options(sessions.Options{
		HttpOnly: true,
		Path: "/"})
	m.Use(sessions.Sessions("session", sessionStore))
	m.Use(render.Renderer())
	
	m.Use(api.BaseApiHandler())
	m.Use(service.ServiceHandler())
	
	//api.Router(m.Router)
	
	// let's go
	m.RunOnAddr(":8000")
}
