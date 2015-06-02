package api

import (
	"strconv"
	"github.com/go-martini/martini"
	//	"github.com/martini-contrib/binding"
	"encoding/json"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"net/url"
//	"fmt"
)

type IBaseApi interface {
	ResultList(data interface{}, totalCount int)
	ResultDetail(data interface{})
	ResultFailed(message string)
	ResultError(message string)
	
	SessionSet(key interface{}, val interface{})
	SessionClear()
	Query(key string) string
	JsonForm() map[string]interface{}
}

type IAuthBaseApi interface {
	IBaseApi
	ResultUnauthorized()
	UserId() string
}

type apier struct {
	res     http.ResponseWriter
	req     *http.Request
	render  render.Render
	session sessions.Session
	
	queryValues	url.Values
}

func BaseApiHandler() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, render render.Render, session sessions.Session) {
		a := apier{res, req, render, session, nil}
		c.MapTo(&a, (*IBaseApi)(nil))
		c.MapTo(&a, (*IAuthBaseApi)(nil))
	}
}

func (c apier) UserId() string {
	pid := c.session.Get("uid")
	if pid == nil {
		return ""
	} else {
		return pid.(string)
	}
}

func (c apier) SessionSet(key interface{}, val interface{}){
	c.session.Set(key, val)
}
func (c apier) SessionClear(){
	c.session.Clear()
}

func (c apier) Query(key string) string {	
	if c.queryValues == nil{
		c.queryValues = c.req.URL.Query()
	}
	return c.queryValues.Get(key)
}

func (c apier) JsonForm() map[string]interface{} {	
	var form map[string]interface{}
	decoder := json.NewDecoder(c.req.Body)
	decoder.Decode(&form)
	return form
}

// at list
func (c apier) ResultList(data interface{}, totalCount int) {
	header := c.res.Header()
	header.Set("Access-Control-Expose-Headers", "X-Total-Count")
	header.Set("X-Total-Count", strconv.Itoa(totalCount))
	if data == nil{
		data = []interface{}{}//make([]interface{},0)
	}
	//	data = []interface{}{}
	c.render.JSON(http.StatusOK, data)
}
// at detail
func (c apier) ResultDetail(data interface{}){
	c.render.JSON(http.StatusOK, data)
}
// at failed
func (c apier) ResultFailed(message string){
	header := c.res.Header()
	header.Set("X-Failed", "true")
	c.render.JSON(http.StatusOK, resultFailed{true,message})
}

type resultFailed struct{
	Failed bool
	Message string
}

// at server error
func (c apier) ResultError(message string){
	c.render.JSON(http.StatusInternalServerError, message)
}

func (c apier) ResultUnauthorized(){
	c.render.Status(http.StatusUnauthorized)
}



type Fake_apier struct {
	// in
	QueryValues	url.Values
	Uid	string
	
	// out
	Form		map[string]interface{}
	Result	resultFailed
}

func (c *Fake_apier) Query(key string) string {	
	return c.QueryValues.Get(key)
}

func (c *Fake_apier) UserId() string {
	return c.Uid
}

func (c *Fake_apier) SessionSet(key interface{}, val interface{}){
	
}
func (c *Fake_apier) SessionClear(){
	
}

func (c *Fake_apier) JsonForm() map[string]interface{} {	
	
	return c.Form
}
