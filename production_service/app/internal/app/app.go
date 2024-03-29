package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "production_service/docs" // needed to correctly open swagger UI
	"production_service/internal/domain/product/storage"
	"production_service/pkg/client/postgresql"
	"production_service/pkg/metric"

	"production_service/internal/config"
	"production_service/pkg/logging"
)

type App struct {
	cfg        *config.Config
	logger     *logging.Logger
	router     *httprouter.Router
	httpServer *http.Server
	pgClient   *pgxpool.Pool
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	logger.Println("initializing router")
	router := httprouter.New()

	logger.Println("initializing swagger docs")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	logger.Println("initializing heartbeat metric")
	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username,
		config.PostgreSQL.Password,
		config.PostgreSQL.Host,
		config.PostgreSQL.Port,
		config.PostgreSQL.Database,
	)
	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatalf("Failed to initialize Postgres client: %v", err)
	}

	productStorage := storage.NewProductStorage(pgClient, logger)
	all, err := productStorage.All(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
	logger.Fatal(all)

	app := App{cfg: config, logger: logger, router: router, pgClient: pgClient}
	return app, nil
}

func (a *App) Run() {
	a.startHTTP()
}

func (a *App) startHTTP() {
	a.logger.Info("starting HTTP")

	var listener net.Listener // can be unix or tcp defining on listen type
	if a.cfg.Listen.Type == config.LISTEN_TYPE_SOCK {
		// define current directory
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			a.logger.Fatal(err)
		}
		socketPath := path.Join(appDir, a.cfg.Listen.SocketFile)
		a.logger.Infof("socket path: %q", socketPath)

		a.logger.Info("create and listen unix socket")
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			a.logger.Fatal(err)
		}
	} else {
		a.logger.Infof("bind application to host %q and port %q", a.cfg.Listen.BindIP, a.cfg.Listen.Port)
		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Listen.BindIP, a.cfg.Listen.Port))
		if err != nil {
			a.logger.Fatal(err)
		}
	}
	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowedOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "content-type"}, // TODO: complete
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Location", "Authorization", "Content-Disposition"},
		// enable debugging for testing, consider disabling on production
		Debug: false,
	})

	handler := c.Handler(a.router)
	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	a.logger.Println("application completely initialized and started")
	if err := a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Warn("server shutdown")
		default:
			a.logger.Fatal(err)
		}
	}
	err := a.httpServer.Shutdown(context.Background())
	if err != nil {
		a.logger.Fatal(err)
	}

}
