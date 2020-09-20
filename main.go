package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

var client *db.Client
var ctx context.Context

func init(){
	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://belajarfirebase-e1dd9.firebaseio.com/",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("firebase-admin-sdk.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
}

func main(){
	router := gin.Default()
	router.POST("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/status",GetAntrianHandler)
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

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	//
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

func addAntrian() (bool,error){
	_,_,dataAntrian := getAntrian()
	var Id string
	var antrianRef *db.Ref
	ref := client.NewRef("antrian")

	if dataAntrian == nil{
		Id = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	}else {
		Id = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d",len(dataAntrian)))
	}
	antrian := Antrian{
		Id: Id,
		Status: false,
	}
	if err := antrianRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false,err
	}
	return true,nil
}

func getAntrian() (bool,error,[]map[string]interface{}){
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
		return false,err,nil
	}

	return true,nil,data
}

type Antrian struct {
	Id string `json:"id"`
	Status bool `json:"status"`
}