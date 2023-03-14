package db

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBs ...
type DBs struct {
	Viska *gorm.DB
}

var db *gorm.DB

// Initialize Method: Connection Test to DB Server
func Initialize(dbName string) *gorm.DB {
	gotenv.Load()

	// connSetting := "charset=utf8mb4&parseTime=True&loc=Local"
	connSetting := "charset=utf8mb4&parseTime=true&checkConnLiveness=true"
	connString := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@(" + os.Getenv("DB_IP") + ")/" + dbName + "?" + connSetting
	log.Println(">", os.Getenv("DB_USER"))

	db, err := gorm.Open(os.Getenv("DB_TYPE"), connString)
	if err != nil {
		log.Println(fmt.Sprintf("[DB SRV] Error Connection Testing to DB - %v", err))
		return nil
	}
	log.Println(fmt.Sprintf("[DB SRV] Successful Connection Testing to DB: %v", connString))

	db.DB().SetMaxOpenConns(25)
	db.DB().SetMaxIdleConns(25)
	db.DB().SetConnMaxLifetime(5 * time.Minute)

	active, _ := strconv.ParseBool(os.Getenv("DB_DEBUG"))
	db.LogMode(active)

	err = db.Exec("SET SESSION sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));").Error
	if err != nil {
		log.Panic(err)
	}
	return db
}

// ClinkConnection ...
func (m *DBs) ViskaConnection() {
	if os.Getenv("DB_NAME") != "" {
		m.Viska = Initialize(os.Getenv("DB_NAME"))
		return
	}
	m.Viska = Initialize("viska")
}

// TimeLocalNow ...
func TimeLocalNow() *time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	return &now
}

func URLRewriter(router *mux.Router, baseURLPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = func(url string) string {
			//log.Println("url path: ", url)
			if strings.Index(url, baseURLPath) == 0 {
				url = url[len(baseURLPath):]
			}
			//log.Println("after rewrite: ", url)
			return url
		}(r.URL.Path)

		router.ServeHTTP(w, r)
	}
}
