package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thelamedev/mattertui/internal/config"
	"github.com/thelamedev/mattertui/internal/server/database"
	"github.com/thelamedev/mattertui/internal/server/handlers"
	"github.com/thelamedev/mattertui/internal/server/middlewares"
)

type Server struct {
	bindAddr string
	engine   *gin.Engine
	config   *config.Config
}

func NewServer(cfg *config.Config) *Server {
	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.SetTrustedProxies(nil)

	s := &Server{
		bindAddr: cfg.Server.BindAddr,
		engine:   engine,
		config:   cfg,
	}

	pool := database.GetPool()
	handlers.InitStore(pool)
	handlers.InitConfig(cfg)

	s.routes()

	return s
}

func (s *Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s *Server) Run() <-chan error {
	quitCh := make(chan error, 1)
	go func() {
		quitCh <- s.engine.Run(s.bindAddr)
	}()
	return quitCh
}

func (s *Server) routes() {
	s.engine.Use(gin.Recovery())
	s.engine.Use(gin.Logger())

	api := s.engine.Group("/api")
	v1 := api.Group("/v1")

	private := v1.Group("/")
	private.Use(middlewares.AuthMiddleware(s.config))
	{
		private.GET("/users/:id", handlers.HandleGetUserByID)
	}

	public := v1.Group("/")
	{
		public.POST("/auth/register", handlers.HandleRegisterUser)
		public.POST("/auth/login", handlers.HandleLoginUser)
	}
}
