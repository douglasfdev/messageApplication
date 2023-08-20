package main

import (
	"baby-steps/internal/infra/akafka"
	"baby-steps/internal/infra/repository"
	"baby-steps/internal/infra/web"
	"baby-steps/internal/usecase"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host:docker.internal:3306/products)")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository.NewProductRepositoryMySql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host:docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)

		if err != nil {

		}

		_, err = createProductUseCase.Execute(dto)
	}
}
