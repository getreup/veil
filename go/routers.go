/*
 * Swagger Petstore
 *
 * This is a sample server Petstore server.  You can find out more about     Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key `special-key` to test the authorization     filters.
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v2/",
		Index,
	},

	Route{
		"AddPet",
		strings.ToUpper("Post"),
		"/v2/pet",
		AddPet,
	},

	Route{
		"DeletePet",
		strings.ToUpper("Delete"),
		"/v2/pet/{petId}",
		DeletePet,
	},

	Route{
		"FindPetsByStatus",
		strings.ToUpper("Get"),
		"/v2/pet/findByStatus",
		FindPetsByStatus,
	},

	Route{
		"FindPetsByTags",
		strings.ToUpper("Get"),
		"/v2/pet/findByTags",
		FindPetsByTags,
	},

	Route{
		"GetPetById",
		strings.ToUpper("Get"),
		"/v2/pet/{petId}",
		GetPetById,
	},

	Route{
		"UpdatePet",
		strings.ToUpper("Put"),
		"/v2/pet",
		UpdatePet,
	},

	Route{
		"UpdatePetWithForm",
		strings.ToUpper("Post"),
		"/v2/pet/{petId}",
		UpdatePetWithForm,
	},

	Route{
		"UploadFile",
		strings.ToUpper("Post"),
		"/v2/pet/{petId}/uploadImage",
		UploadFile,
	},

	Route{
		"DeleteOrder",
		strings.ToUpper("Delete"),
		"/v2/store/order/{orderId}",
		DeleteOrder,
	},

	Route{
		"GetInventory",
		strings.ToUpper("Get"),
		"/v2/store/inventory",
		GetInventory,
	},

	Route{
		"GetOrderById",
		strings.ToUpper("Get"),
		"/v2/store/order/{orderId}",
		GetOrderById,
	},

	Route{
		"PlaceOrder",
		strings.ToUpper("Post"),
		"/v2/store/order",
		PlaceOrder,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/v2/user",
		CreateUser,
	},

	Route{
		"CreateUsersWithArrayInput",
		strings.ToUpper("Post"),
		"/v2/user/createWithArray",
		CreateUsersWithArrayInput,
	},

	Route{
		"CreateUsersWithListInput",
		strings.ToUpper("Post"),
		"/v2/user/createWithList",
		CreateUsersWithListInput,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/v2/user/{username}",
		DeleteUser,
	},

	Route{
		"GetUserByName",
		strings.ToUpper("Get"),
		"/v2/user/{username}",
		GetUserByName,
	},

	Route{
		"LoginUser",
		strings.ToUpper("Get"),
		"/v2/user/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		strings.ToUpper("Get"),
		"/v2/user/logout",
		LogoutUser,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/v2/user/{username}",
		UpdateUser,
	},
}
