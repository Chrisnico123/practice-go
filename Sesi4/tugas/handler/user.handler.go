package handler

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func generateTraceID() string {
	timestamp := time.Now().UnixNano()
	randomNumber := rand.Intn(100000) // Angka acak antara 0 dan 99999

	traceID := strconv.FormatInt(timestamp, 10) + strconv.Itoa(randomNumber)
	return traceID
}

type User struct {
	Id int
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
}

var data = []User{}
var index = 1

func GetDataAllUser(c echo.Context) error {
	c.Set("trace_id", generateTraceID())
	response := map[string]interface{}{
		"Success" : true,
		"StatusCode" : http.StatusAccepted,
		"Message" : "get all success",
		"Payload" : data,
	}
	return c.JSON(http.StatusAccepted , response)	
}
func CreateDataUser(c echo.Context) error {
	c.Set("trace_id", generateTraceID())
	var req User
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusAccepted , echo.Map{
			"error" : err.Error(),
		})
	}
	req.Id = index
	index = index + 1

	data = append(data, req)


	return c.JSON(http.StatusAccepted , echo.Map{
		"Success" : true,
		"StatusCode" : http.StatusAccepted,
		"Message" : "Created Success",
	})
}
func UpdateDataUser(c echo.Context) error {
	c.Set("trace_id", generateTraceID())
	var req User
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusAccepted , echo.Map{
			"error" : err.Error(),
		})
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	for i, user := range data {
		if (user.Id == id){ 
			data[i].Name = req.Name
			data[i].Email = req.Email
			data[i].Address = req.Address
		}
	}
	return c.JSON(http.StatusAccepted , echo.Map{	
		"success" : true,
		"status_code" :	http.StatusAccepted,
		"message" : "update success",
	})
}
func DeleteUser(c echo.Context) error {
	c.Set("trace_id", generateTraceID())
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusAccepted , echo.Map{
			"error" : err.Error(),
		})
	}

	found := false
    for i, user := range data {
        if user.Id == id {
            data = append(data[:i], data[i+1:]...)
            found = true
            break
        }
    }

    if !found {
        return c.JSON(http.StatusNotFound, echo.Map{
            "error": "User not found",
        })
    }
	return c.JSON(http.StatusAccepted , echo.Map{
		"success" : true,
		"status_code" :	http.StatusAccepted,
		"message" : "delete success",
	})
}