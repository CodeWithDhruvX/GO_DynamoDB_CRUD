package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/akhil/dynamodb-go-crud-yt/internal/repository/adapter"
	"github.com/akhil/dynamodb-go-crud-yt/internal/repository/instance"
	"github.com/akhil/dynamodb-go-crud-yt/internal/routes"
	"github.com/akhil/dynamodb-go-crud-yt/internal/rules"
	"github.com/akhil/dynamodb-go-crud-yt/internal/rules/product"
	"github.com/akhil/dynamodb-go-crud-yt/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	configs:=Config.GetConfig()
	connection:=instance.GetConnection()
	repository:=adapter.NewAdapter(connection)


	logger.INFO("waiting for the service start...",nil)
	errors:=Migrate(connection)

	if len(errors) > 0 {
		for _,err:=range errors{
			logger.PANIC("Error on Migration....",err)
		}
	}

	logger.PANIC("",checkTables(connection))

	port:=fmt.Sprintf(":%v",configs.Port)
	router:=routes.NewRouter().SetRouters(repository)
	logger.INFO("service is running on port",port)

	server:=http.ListenAndServe(port,router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error{
	var errors []error
	callMigrateAndAppendError(&errors,connection,&RulesProduct.Rules{})
	return errors
}

func callMigrateAndAppendError(errors *[]error,connection *dynamodb.DynamoDB,rules rules.Interface){
	err:=rule.Migrate(connection)

	if err!=nil {
		*errors=append(*errors,err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response,err:=connection.ListTables(&dynamodb.ListTables(&dynamodb.ListTablesInput{}))

	if reponse!=nil {
		if len(response.TableNames) == 0 {
			logger.INFO("Tables Not Found:",nil)
		}
	}

	for _,tableName:=range response.TableNames {
		logger.INFO("Table Found:",*tableName)
	}
	
	return err
}


