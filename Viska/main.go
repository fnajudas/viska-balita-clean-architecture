package main

import (
	"os"
	"viska/db"
	"viska/middleware/logger"
	"viska/routes"

	"github.com/thedevsaddam/renderer"

	getDataHandler "viska/handler/read"
	getDataService "viska/services/read"
	getDataMysql "viska/storage/mysql/read"

	updateDataHandler "viska/handler/update"
	updateDataServices "viska/services/update"
	updateDataMysql "viska/storage/mysql/update"

	createDataHandler "viska/handler/create"
	createDataServices "viska/services/create"
	createDataMysql "viska/storage/mysql/create"

	deleteDataHandler "viska/handler/delete"
	deleteDataService "viska/services/delete"
	deleteDataMysql "viska/storage/mysql/delete"

	getAuth "viska/handler/authController"

	getToken "viska/middleware/jwt"
)

var Databasess = db.DBs{}

func main() {
	Databasess.ViskaConnection()
	defer Databasess.Viska.Close()

	logger := logger.New()
	render := renderer.New()

	db := Databasess.Viska

	dataMysql := getDataMysql.NewDataBalita(logger, db)
	dataService := getDataService.NewExecutor(logger, dataMysql)
	dataHandler := getDataHandler.NewHandler(logger, dataService, render)

	updateMysql := updateDataMysql.NewUpdateDataBalita(logger, db, 5)
	updateServices := updateDataServices.NewExecutor(logger, updateMysql)
	updateHandler := updateDataHandler.NewHandler(logger, render, db, updateServices)

	dataRegister := getAuth.NewAuth(logger, render, db)
	dataLogin := getAuth.NewAuth(logger, render, db)

	createMysql := createDataMysql.NewCreateBalita(db, logger, 5)
	createServices := createDataServices.NewExecutor(logger, createMysql)
	createHanler := createDataHandler.NewHandler(logger, render, createServices)

	deleteMysql := deleteDataMysql.NewDeleteBalita(logger, db, 5)
	deleteService := deleteDataService.NewExecutor(logger, deleteMysql)
	deleteHandler := deleteDataHandler.NewHandler(logger, render, deleteService)

	dataToken := getToken.NewTokenGenerator(render)

	r := routes.Routes{
		ListDataBalita: dataHandler,
		DtRegister:     dataRegister,
		DtLogin:        dataLogin,
		DtToken:        dataToken,
		UpdateData:     updateHandler,
		CreateData:     createHanler,
		DeleteData:     deleteHandler,
	}
	r.Run(os.Getenv("APP_PORT"))
}
