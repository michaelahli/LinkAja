package usecases

import (
	"LinkaAja/packages"
	"LinkaAja/src/repositories"
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

func TestAccountID(t *testing.T) {
	db := packages.LoadDatabase()
	repository := repositories.NewRepo(db)
	usecase := NewUC(repository)
	type args struct {
		urlstring string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Sucess",
			args: args{
				urlstring: "http://localhost:8080/account/555001",
			},
			want:    555001,
			wantErr: false,
		},
		{
			name: "Success",
			args: args{
				urlstring: "localhost:8080/account/555001",
			},
			want:    555001,
			wantErr: false,
		},
		{
			name: "Failed",
			args: args{
				urlstring: "http://localhost:8080/account/555001u",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, testitem := range tests {
		t.Run(
			testitem.name, func(t *testing.T) {
				got1, err := usecase.GetAccountID(testitem.args.urlstring)
				if (err != nil) != testitem.wantErr {
					t.Errorf("GetAccountID() Method error got %v, wantErr %v", err, testitem.wantErr)
					return
				}
				if got1 != testitem.want {
					t.Errorf("GetAccountID() Method got %v , want %v", got1, testitem.want)
					return
				}
			})
	}
}

func TestGetBalanceInfo(t *testing.T) {

}
