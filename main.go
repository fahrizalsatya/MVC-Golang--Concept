package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var data []Antrian

func main(){
	router := gin.Default()
	router.POST("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/status",GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian",UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrian/delete",DeleteAntrianHandler)
	router.Run(":8080")
}

func AddAntrianHandler(c *gin.Context){
	flag,err := addAntrian()
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status":"success",
		})
	}else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"error":err,
		})
	}
}

func GetAntrianHandler(c *gin.Context){

	flag,err,resp := getAntrian()
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status":"success",
			"data":resp,
		})
	}else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"error":err,
		})
	}
}

func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag,err := updateAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status":"success",
		})
	}else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"error":err,
		})
	}
}

func DeleteAntrianHandler(c *gin.Context){
	idAntrian := c.Param("idAntrian")
	flag,err := deleteAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status":"success",
		})
	}else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"error":err,
		})
	}
}

func addAntrian() (bool,error){
	_,_,dataAntrian := getAntrian()
	var Id string

	if dataAntrian == nil{
		Id = fmt.Sprintf("B-0")
	}else {
		Id = fmt.Sprintf("B-%d", len(dataAntrian))
	}
	data =append(data, Antrian{
			Id: Id,
			Status: false,
		})

	return true,nil
}

func getAntrian() (bool,error,[]Antrian){
	return true,nil,data
}

func updateAntrian(idAntrian string) (bool,error){


	for i,_ := range data{
		if data[i].Id == idAntrian{
			data[i].Status = true
			break
		}
	}

	return true,nil
}

func deleteAntrian (idAntrian string) (bool,error){

	for i := range data{
		if data[i].Id == idAntrian{
			data = append(data[:i], data[i+1:]...)
		}
	}

	return true,nil
}

type Antrian struct {
	Id string `json:"id"`
	Status bool `json:"status"`
}