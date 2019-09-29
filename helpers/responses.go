package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func ReturnFoundRow(g *gin.Context , msg string)  {
	var errors map[string]string
	var data map[string]interface{}
	response(g , msg , data, errors ,http.StatusConflict , 409 , false)
	return
}

func ReturnNotValidRequest(error *govalidator.Validator, g *gin.Context) bool {
	e := error.ValidateJSON()
	if len(e) > 0 {
		g.JSON(
			http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "something not valid in your request",
				"errors":  e,
				"code":    400,
			})
		return true
	}
	return false
}

func ReturnNotFound(g *gin.Context, msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(g , msg , data, errors ,http.StatusNotFound , 404 , false)
	return
}

func ReturnForbidden(g *gin.Context, msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(g , msg , data, errors ,http.StatusForbidden , 304 , false)
	return
}

func OkResponse(g *gin.Context, msg string, data map[string]interface{}) {
	var errors map[string]string
	response(g , msg , data, errors ,http.StatusOK , 200 , true)
	return
}

func response(g *gin.Context, msg string, data map[string]interface{} , errors map[string]string , httpStatus int, code int , status bool)  {
	g.JSON( httpStatus , gin.H{
			"status":  status,
			"message": msg,
			"errors":  errors,
			"code":    code,
			"data"  :data,
		})
	return
}