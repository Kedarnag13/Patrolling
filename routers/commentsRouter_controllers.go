package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:RegisterController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	// beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:ObjectController"],
	// 	beego.ControllerComments{
	// 		Method:           "Get",
	// 		Router:           `/:objectId`,
	// 		AllowHTTPMethods: []string{"get"},
	// 		Params:           nil})

	// beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:ObjectController"],
	// 	beego.ControllerComments{
	// 		Method:           "GetAll",
	// 		Router:           `/`,
	// 		AllowHTTPMethods: []string{"get"},
	// 		Params:           nil})

	// beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:ObjectController"],
	// 	beego.ControllerComments{
	// 		Method:           "Put",
	// 		Router:           `/:objectId`,
	// 		AllowHTTPMethods: []string{"put"},
	// 		Params:           nil})

	// beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Kedarnag13/Patrolling/controllers:ObjectController"],
	// 	beego.ControllerComments{
	// 		Method:           "Delete",
	// 		Router:           `/:objectId`,
	// 		AllowHTTPMethods: []string{"delete"},
	// 		Params:           nil})

}
