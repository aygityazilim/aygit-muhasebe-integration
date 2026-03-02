package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB // Ana Entegratör Veritabanı
)

// InitDB initializes the single database connection using environment variables
func InitDB() {
	var err error

	// Veritabanı bağlantı dizesi (DSN) oluşturulur
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// pgx sürücüsü ile sqlx üzerinden bağlantı kurulur
	DB, err = sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	log.Println("Ana veritabanı bağlantısı başarıyla kuruldu")

	// Şemayı otomatik olarak oluştur (Tablolar yoksa)
	InitSchema()
}

// InitSchema, pkg/db/schema.sql dosyasını okur ve veritabanında çalıştırır.
func InitSchema() {
	schema, err := os.ReadFile("pkg/db/schema.sql")
	if err != nil {
		log.Printf("Şema dosyası okunamadı (pkg/db/schema.sql): %v", err)
		return
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		log.Printf("Şema başlatma hatası: %v", err)
	} else {
		log.Println("Veritabanı şeması başarıyla güncellendi/kontrol edildi")
	}
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
