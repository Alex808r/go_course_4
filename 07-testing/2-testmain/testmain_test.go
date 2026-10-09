package testmain

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v4"
)

var testService *Service

func TestMain(m *testing.M) {
	// Демонстрация TestMain: инициализация общих ресурсов перед запуском всех тестов пакета.
	conn, err := pgx.Connect(context.Background(), "postgres://user:pwd@server/database")
	if err != nil {
		log.Printf("Внимание: тестовая БД недоступна (%v). Запуск демонстрационных тестов без подключения.", err)
	} else {
		defer conn.Close(context.Background())
	}

	testService = &Service{
		db: conn,
	}

	code := m.Run()
	// Здесь может идти освобождение ресурсов.
	os.Exit(code)
}

func TestService_Products(t *testing.T) {
	got := testService.Products()
	want := []Product{
		{
			Name:  "Компьютер",
			Price: 20_000,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Service.Products() = %v, want %v", got, want)
	}
}
