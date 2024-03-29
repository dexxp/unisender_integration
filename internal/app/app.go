package app

import (
	"git.amocrm.ru/dmiroshnikov/unisender_integration/api"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/config"
	grpc2 "git.amocrm.ru/dmiroshnikov/unisender_integration/internal/controller/grpc"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/entity"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/account_v1"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/mysql"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

type App struct {
	Cfg             *config.Config
	Srv             *http.Server
	db              mysql.Database
	serviceProvider *ServiceProvider

	grpcServer *grpc.Server
}

func NewApp(cfg *config.Config) *App {
	app := &App{
		Cfg: cfg,
	}

	err := app.InitDB()

	if err != nil {
		panic(err)
	}

	amoClient := InitAmoRequest()
	uniClient := InitUniRequest()

	app.serviceProvider = Wire(app.db, amoClient, uniClient, app.Cfg.Auth)

	app.InitGRPCServer(WireGrpc(app.db))

	router := app.InitHttpRoutes()

	app.InitHttpServer(app.Cfg.HTTP, router)

	return app
}

func (a *App) InitDB() error {
	a.db = mysql.NewDatabase(a.Cfg.MySQL)
	err := a.db.DB.AutoMigrate(&entity.Account{}, &entity.Integration{}, &entity.Contact{}, &entity.Email{})
	if err != nil {
		return err
	}
	return nil
}

func InitAmoRequest() *api.AmocrmRequest {
	clientAmo := api.NewRequest(&http.Client{})
	amoReq := api.NewAmocrmRequest(clientAmo)

	return amoReq
}

func InitUniRequest() *api.UnisenderRequest {
	clientAmo := api.NewRequest(&http.Client{})
	uniReq := api.NewUnisenderRequest(clientAmo)

	return uniReq
}

func (a *App) InitHttpRoutes() *mux.Router {
	const (
		auth           = "/"
		contactsId     = "/contacts/[id]"
		contactsSync   = "/contacts/[id]?sync=true"
		contactsNoSync = "/contacts/[id]?sync=false"
		integration    = "/integration/[id]"
		unisender      = "/unisender/"
		unisenderId    = "/unisenderAccount/[id]"
		disable        = "/disable"
		update         = "/update"
		delete         = "/delete"
		create         = "/create"
	)

	router := mux.NewRouter()

	router.HandleFunc(auth, a.serviceProvider.aImpl.FistAuth).Methods("GET")
	router.HandleFunc(unisender, a.serviceProvider.uImpl.FirstSync).Methods("POST")
	router.HandleFunc(disable, a.serviceProvider.dImpl.NewDisable).Methods("GET")

	return router
}

func (a *App) InitHttpServer(cfg config.HTTP, router *mux.Router) {
	a.Srv = &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
}

func (a *App) RunHTTPServer() error {
	return a.Srv.ListenAndServe()
}

func (a *App) InitGRPCServer(impl *grpc2.Implementation) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	account_v1.RegisterAccountServiceServer(grpcServer, impl)

	a.grpcServer = grpcServer
}

func (a *App) RunGRPCServer() error {
	log.Printf("GRPC server is running on %s", ":"+a.Cfg.GRPC.Port)

	list, err := net.Listen("tcp", ":"+a.Cfg.GRPC.Port)
	if err != nil {
		return err
	}

	err = a.Srv.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
