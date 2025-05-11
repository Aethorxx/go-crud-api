package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"

	"go-crud-api/internal/config"
)

func RunMigrations() error {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Подключение к базе данных
	db, err := sql.Open("postgres", cfg.DB.GetDSN())
	if err != nil {
		return err
	}
	defer db.Close()

	// Проверка соединения
	if err := db.Ping(); err != nil {
		return err
	}

	// Настройка миграций
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	// Применение миграций
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	log.Printf("Applied %d migrations!\n", n)
	return nil
}
