init:
	mkdir cmd cmd/equocenterback initializers pkg pkg/models pkg/controllers pkg/services
	touch cmd/equocenterback/main.go 
	go mod init equocenterback
	go get github.com/githubnemo/CompileDaemon
	go install github.com/githubnemo/CompileDaemon
	go get github.com/joho/godotenv
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get go.mongodb.org/mongo-driver/mongo

run local:
	CompileDaemon -directory="./cmd/equocenterback" -command="./cmd/equocenterback/equocenterback"