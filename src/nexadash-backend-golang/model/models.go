package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"nexadash-backend-go/basemodel"
	"encoding/json"
	"net/http"
	"reflect"
)

func Make_app_name(r *http.Request) basemodel.Response{
	var app basemodel.Apps8
	var resp basemodel.Response
	db := basemodel.Db_connection()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&app)
	fmt.Println(decoder)
	fmt.Println(reflect.TypeOf(decoder))
	project_id := basemodel.Check_header(r)
	fmt.Println(len(app.ID))
	if len(project_id) == 0{
		project_id = "default"
	}
	if len(app.ID) == 0{
		err := db.Find(&app,"project_id = ?", project_id).Error
		if err != nil && project_id != "default" {
			resp = basemodel.Response{Errors: err, Message: "Something went wrong", Status: false}
		}else{
				fmt.Println("running else in")
				resp = basemodel.Insert(r, project_id)

		}
	}else {
		err := db.Find(&app,"project_id = ?", project_id).Error
		if err != nil && project_id != "default" {
			resp = basemodel.Response{Errors: err, Message: "Something went wrong", Status: false}
		} else {
			resp = basemodel.Update_and_insert_check(app, project_id)
		}
	}
	fmt.Println(app)
	return resp
}