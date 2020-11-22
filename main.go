package main

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var data []Antrian
=======
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
=======
	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-MVC-Golang-Concept/app/controller"
>>>>>>> 00cd39fc767441d4419619135cf00be50a8ea4f8
	"github.com/gin-gonic/gin"
)


<<<<<<< HEAD
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
>>>>>>> 4b950a09f8fbd2fca05f46e8e1f6843a920611bf
=======
>>>>>>> 00cd39fc767441d4419619135cf00be50a8ea4f8

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian", controller.PageAntrianHandler)
	router.Run(":8080")
}



<<<<<<< HEAD
func addAntrian() (bool, error) {
	_, _, dataAntrian := getAntrian()
	var Id string

	if dataAntrian == nil {
		Id = fmt.Sprintf("B-0")
<<<<<<< HEAD
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
=======
		antrianRef = ref.Child("0")
	} else {
		Id = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d", len(dataAntrian)))
	}
	antrian := Antrian{
		Id:     Id,
		Status: false,
	}
	if err := antrianRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func getAntrian() (bool, error, []map[string]interface{}) {
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
		return false, err, nil
	}

	return true, nil, data
}

func updateAntrian(idAntrian string) (bool, error) {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	antrian := Antrian{
		Id:     idAntrian,
		Status: true,
	}
	if err := childRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
>>>>>>> 4b950a09f8fbd2fca05f46e8e1f6843a920611bf
	}

	return true, nil
}

func deleteAntrian(idAntrian string) (bool, error) {

<<<<<<< HEAD
	for i := range data{
		if data[i].Id == idAntrian{
			data = append(data[:i], data[i+1:]...)
		}
=======
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	if err := childRef.Delete(ctx); err != nil {
		log.Fatal(err)
		return false, err
>>>>>>> 4b950a09f8fbd2fca05f46e8e1f6843a920611bf
	}

	return true, nil
}

type Antrian struct {
	Id     string `json:"id"`
	Status bool   `json:"status"`
}
=======

>>>>>>> 00cd39fc767441d4419619135cf00be50a8ea4f8
