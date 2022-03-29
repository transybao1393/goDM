package main

import (
	"log"
	"runtime"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_articleHttpDelivery "godm/article/delivery/http"
	_articleHttpDeliveryMiddleware "godm/article/delivery/http/middleware"
	_articleRepo "godm/article/repository/mysql"
	_articleUcase "godm/article/usecase"
	_authorRepo "godm/author/repository/mysql"

	_inputSources "godm/inputSources"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {

	//- test
	for i := 0; i < 10; i++ {
		go _inputSources.MysqlInstance()
	}

	//- new dbConn
	dbConn := _inputSources.MysqlInstance()
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	_articleHttpDelivery.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
