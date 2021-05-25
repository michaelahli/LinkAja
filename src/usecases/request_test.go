package usecases

import (
	"LinkaAja/packages"
	"os"
	"testing"
)

func init() {
	env := map[string]string{
		"SERVER_PORT": "8080",
		"MONGO_DB":    "LinkAja",
		"MONGO_URI":   "mongodb+srv://teredict:25285282@teredict.jkdyu.mongodb.net/LinkAja?retryWrites=true&w=majority",
		"CORS_HEADER": "x-api-key,Authorization,Content-Type,Origin,Accept,Access-Control-Allow-Headers,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin",
		"CORS_METHOD": "OPTION,GET,PUT,POST,DELETE",
		"CORS_ORIGIN": "*",
	}

	for key, value := range env {
		os.Setenv(key, value)
	}
	packages.LoadMongo()
}

func TestProcessTransferRequest(t *testing.T) {

}
