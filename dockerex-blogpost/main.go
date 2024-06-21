package main

import (
	"blogpost/adapter"
	"blogpost/lookup"
	"blogpost/repository"
	"blogpost/router"
	"fmt"
	"os"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

func main() {
	// claims, _ := helper.ToGetClaims("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc0NzU4MzUsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.RXC33CziX3IRMXzs-4iVYyaYvgt8Z6mYuiuv4VGox4k")
	// claaims, _ := helper.ToGetClaims("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc4OTEzOTQsInJvbGUiOiJhZG1pbiIsInVzZXJJZCI6MX0.IqpuqLshN1ZIaKiWLsF-0-CoaZ4TnrzG2TM_t9_TXS0")
	// fmt.Print("---------", claims)
	// fmt.Println("----not exp", claaims)
	db := adapter.DBconnection()
	adapter.RedisConnection()
	repository.ToSetDB(db)
	files, err := os.ReadDir("lookup")
	if err != nil {
		fmt.Print("Cannot read file")
		return
	}
	for _, v := range files {
		if v.Name() != "master.go" {
			file := strings.TrimSuffix(v.Name(), ".go")
			filename := strings.Split(file, "_")
			ver := lookup.Lookups{}
			toCheck := db.Table("lookups").Where("version=?", filename[1]).Find(&ver)
			if toCheck.RowsAffected == 0 {
				v := lookup.Formethod{}
				method := reflect.ValueOf(&v).MethodByName(file).Interface().(func(*gorm.DB))
				method(db)
				row := lookup.Lookups{Name: file, Version: filename[1]}
				db.Table("lookups").Create(row)
			}
		}
	}
	router.Router()
}
