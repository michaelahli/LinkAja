package packages

import (
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

/*
	@created at 25 May 2021
	Method for allowing Cross-Origin Resource
*/
func AllowCORS(router http.Handler) http.Handler {
	origin := strings.Split(os.Getenv("CORS_ORIGIN"), ",")
	headers := strings.Split(os.Getenv("CORS_HEADER"), ",")
	methods := strings.Split(os.Getenv("CORS_METHOD"), ",")
	wrapper := CorsWrapper(router, origin, headers, methods)
	return wrapper
}

func CorsWrapper(h http.Handler, origin []string, headers []string, method []string) http.Handler {
	crs := cors.New(cors.Options{
		AllowedOrigins: origin,
		AllowedHeaders: headers,
		AllowedMethods: method,
	})

	return crs.Handler(h)
}
