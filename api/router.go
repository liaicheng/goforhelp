package api

//import(
//	"github.com/go-martini/martini"
//	//api "."
//	"github.com/martini-contrib/binding"
//)

//func Router(r martini.Router){
//	r.Get("/test", func(base api.IBaseApi){
		
//	})
	
//	r.Post("/admin/auth", binding.Bind(userAuthModel{}), post_auth)
	
//	r.Group("/admin", func(rr martini.Router){
//		rr.Get("/property/:id",get_property)
//		rr.Get("/property",get_propertyList)
//		rr.Put("/property/:id",put_property)
//	}, authRequired)
//}

//func authRequired(base api.IAuthBaseApi) {
//	//debug
//	return
	
//	if base.UserId() == "" {
//		base.ResultUnauthorized()
//	}
//}