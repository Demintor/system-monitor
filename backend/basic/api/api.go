package api

import (
	"net/http"
	"strconv"

	"example.com/basic/model"
	"example.com/basic/sysutils"
	"example.com/basic/web"
	"github.com/gin-gonic/gin"
)

// @Description get process table
// @Produce  json
// @Success 200 {array} model.Process
// @Failure 500 {object} web.APIError Internal Server Error
// @Router /testapi/get-process-table [get]
func GetProcessTable(c *gin.Context) {
	table, err := sysutils.GetProcessTable()

	if err != nil {
		c.JSON(http.StatusInternalServerError, web.APIError{ErrorCode: 1, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, table)
}

// @Description kill process by PID
// @Accept  json
// @Produce  json
// @Param   pid     path    int     true        "PID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "Incorrect PID"
// @Failure 500 {object} web.APIError Internal Server Error
// @Router /testapi/kill-process/{pid} [post]
func KillProcess(c *gin.Context) {

	pid, err := strconv.Atoi(c.Param("pid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, web.APIError{ErrorCode: 0, ErrorMessage: err.Error()})
		return
	}

	err = sysutils.KillProcess(pid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, web.APIError{ErrorCode: 1, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// @Description execute bash script
// @Accept  json
// @Produce  json
// @Param journal body model.Script true "script json"
// @Success 200 {object} model.ExecuteResult
// @Failure 400 {object} web.APIError
// @Failure 500 {object} web.APIError
// @Router /testapi/execute-script/ [post]
func ExecuteScript(c *gin.Context) {
	var script model.Script

	if err := c.ShouldBindJSON(&script); err != nil {
		c.JSON(http.StatusBadRequest, web.APIError{ErrorCode: 0, ErrorMessage: err.Error()})
		return
	}

	output, err := sysutils.ExecuteScript(script.Script)

	if err != nil {
		c.JSON(http.StatusInternalServerError, web.APIError{ErrorCode: 1, ErrorMessage: err.Error()})
		return
	}

	result := model.ExecuteResult{Output: output}

	c.JSON(http.StatusOK, result)
}
