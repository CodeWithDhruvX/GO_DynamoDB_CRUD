package routes

import(
	"net/http"
	"time"
	"github.com/go-chi/cors"
)


type Config struct{
	timeout time.Duration
}

func (c *Config) Cors(next http.Handler) http.Handler{
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedCredentials: true,
		MaxAge:5,
	}).Handler(next)
}

func (c *Config) SetTimeOut(timeInSeconds int) *Config{
	c.timeout=time.Duration(timeInSeconds) * time.Second
	return c
}

func (c *Config) GetTimeOut() time.Duration {
	return c.timeout
}

func NewConfig() *Config{
	return &Config{}
}