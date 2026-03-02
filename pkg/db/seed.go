package db

import (
	"log"

	"aygit-muhasebe-integration/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// SeedData, sistem ilk kez ayağa kalktığında gerekli temel verileri ekler.
func SeedData() {
	log.Println("Varsayılan veriler kontrol ediliyor ve tohumlanıyor (seeding)...")

	// 1. Kurum 01 Ekleme
	seedCompany("1234567801", "NES Test Kurum 01", "test01@nes.com.tr", "U782pAd4%gGO", "C35FE137CBB6EB376E33C47A46B8BBCB17970689CB09BA55AF212BFDAF4A5F9D", "urn:mail:merkezgb@nes.com.tr", "urn:mail:merkezpk@nes.com.tr")

	// 2. Kurum 02 Ekleme
	seedCompany("1234567802", "NES Test Kurum 02", "test02@nes.com.tr", "b48Za0*RFmE$", "2EA5ADFA0A1C2A6E0E2C03DFABCEA1DD9A183AD24596A4405D5F7C8109D0B085", "urn:mail:merkezgb@nes.com.tr", "urn:mail:merkezpk@nes.com.tr")

	// 3. Varsayılan Admin Kullanıcısı
	seedAdmin()
}

func seedCompany(vkn, title, user, pass, apiKey, gb, pk string) {
	var count int
	err := DB.Get(&count, "SELECT COUNT(*) FROM companies WHERE party_identification = $1", vkn)
	if err != nil {
		log.Printf("Firma kontrolü hatası (%s): %v", vkn, err)
		return
	}

	if count == 0 {
		query := `INSERT INTO companies (
			id, party_identification, party_name, nes_user, nes_password, nes_api_key, 
			selected_gb_alias, selected_pk_alias, environment
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 'TEST')`

		_, err = DB.Exec(query, uuid.New(), vkn, title, user, pass, apiKey, gb, pk)
		if err != nil {
			log.Printf("Firma tohumlama hatası (%s): %v", vkn, err)
		} else {
			log.Printf("Firma eklendi: %s", title)
		}
	}
}

func seedAdmin() {
	var count int
	adminEmail := "admin@aygit.com"
	err := DB.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", adminEmail)
	if err != nil {
		log.Printf("Admin kontrolü hatası: %v", err)
		return
	}

	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		query := `INSERT INTO users (id, email, password_hash, full_name, role) VALUES ($1, $2, $3, $4, $5)`

		_, err = DB.Exec(query, uuid.New(), adminEmail, string(hashedPassword), "Sistem Yöneticisi", models.RoleAdmin)
		if err != nil {
			log.Printf("Admin tohumlama hatası: %v", err)
		} else {
			log.Println("Varsayılan admin kullanıcısı oluşturuldu (admin@aygit.com / admin123)")
		}
	}
}
