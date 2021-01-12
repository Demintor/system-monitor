package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"example.com/basic/web"
	"example.com/basic/sysutils"
	"net/http"
)

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(c *gin.Context) {
	err := web.APIError{}
	fmt.Println(err)
}

// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(c *gin.Context) {

}

// @Description get process table
// @Produce  json
// @Success 200 {array} object
// @Failure 500 {object} web.APIError Internal Server Error
// @Router /testapi/get-process-table [get]
func GetProcessTable(c *gin.Context) {
	table, err := sysutils.GetProcessTable()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "{\"errorCode\": 0, \"errorMessage\": \"" + err.Error() + "\"}")
		return
	}

	c.JSON(http.StatusOK, table)
}

type Pet3 struct {
	ID int `json:"id"`
}
