package routes

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


func (r *Router) SetRouters() *chi.Mux{

}

func (r *Router) setConfigRouters(){

}

func RouterHealth(){

}

func RouterProduct(){

}

func EnableTimeOut(){

}

func EnableCORS(){

}

func EnableRecover(){

}

func EnableRequestID(){

}

func EnableRealIP(){
	
}