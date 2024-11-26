package routes

import(
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	ServerConfig "github.com/akhil/dynamodb-go-crud-yt/config"
	HealthHandler "github.com/akhil/dynamodb-go-crud-yt/internal/handlers/health"
	ProductHandler "github.com/akhil/dynamodb-go-crud-yt/internal/handlers/product"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router{
	return &Router{
		config:NewConfig().SetTimeOut(serviceConfig.GetConfig().Timeout),
		router:chi.NewRouter()
	}
}


func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux{
	r.setConfigRouters()
	r.RouterHealth(repository)
	r.RouterProduct(repository)
	return r.router
}

func (r *Router) setConfigRouters(){
	r.EnableCORS()
	r.EnableLogger()
	r.EnableTimeOut()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP()
}

func(r *router) RouterHealth(repository adapter.Interface){
	handlers:=HealthHandler.newHandler(repository)

	r.router.Route("/health",func(route chi.Router){
		route.Post("/",handler.Post)
		route.Get("/",handler.Get)
		route.Put("/",handler.Put)
		route.Delete("/",handler.Delete)
		route.Options("/",handler.Options)
	})

}

func(r *Router) RouterProduct(repository adapter.Interface){
	handler:=ProductHandler.NewHandler(repository)

	r.router.route("/product",func(route chi.Router){
		route.Post("/",handler.Post)
		route.Get("/",handler.Get)
		route.Put("/{ID}",handler.Put)
		route.Delete("/{ID}",handler.Delete)
		route.Options("/",handler.Options)
	})
}

func(r *Router) EnableLogger() *Router{
	r.router.Use(middleware.Logger)
	return r
}

func(r *Router) EnableTimeOut() *Router{
	r.router.Use(middleware.Timeout(r.config.GetTimeOut()))
	return r
}

func(r *Router) EnableCORS() *Router{
	r.router.Use(r.config.Cors)
	return r
}

func(r *Router) EnableRecover() *Router{
	r.router.Use(middleware.Recoverer)
	return r
}

func(r *Router) EnableRequestID() *Router{
	r.router.Use(middleware.RequestID)
	return r
}

func(r *Router) EnableRealIP() *Router{
	r.router.Use(middleware.RealIP)
	return r
}

