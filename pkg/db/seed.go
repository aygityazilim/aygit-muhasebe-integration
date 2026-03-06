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
	seedCompany("1234567801", "NES Test Kurum 01", "test01@nes.com.tr", "U782pAd4%gGO", "14867512BA9803F19AFB1F35E929F2312402B4DD323ABF6955DB2E49E666C375", "urn:mail:defaultgb@nes.com.tr", "urn:mail:defaultpk@nes.com.tr")

	// 2. Kurum 02 Ekleme
	seedCompany("1234567802", "NES Test Kurum 02", "test02@nes.com.tr", "b48Za0*RFmE$", "433084F1232A7FD54C975770AE27B1C52903B35A9F1DAB8FC7FE4EE35ACD7542", "urn:mail:defaultgb@nes.com.tr", "urn:mail:defaultpk@nes.com.tr")

	// 3. Varsayılan Admin Kullanıcısı
	seedAdmin()
}

func seedCompany(vkn, title, user, pass, apiKey, gb, pk string) {
	query := `
		INSERT INTO companies (
			id, party_identification, party_name, nes_user, nes_password, nes_api_key, 
			selected_gb_alias, selected_pk_alias, environment,
			is_einvoice, is_earchive, is_edespatch, is_esmm_user, is_emm_user
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 'TEST', true, true, true, true, true)
		ON CONFLICT (party_identification) DO UPDATE SET
			party_name = EXCLUDED.party_name,
			nes_user = EXCLUDED.nes_user,
			nes_password = EXCLUDED.nes_password,
			nes_api_key = EXCLUDED.nes_api_key,
			selected_gb_alias = EXCLUDED.selected_gb_alias,
			selected_pk_alias = EXCLUDED.selected_pk_alias,
			is_einvoice = true,
			is_earchive = true,
			is_edespatch = true,
			is_esmm_user = true,
			is_emm_user = true
	`

	_, err := DB.Exec(query, uuid.New(), vkn, title, user, pass, apiKey, gb, pk)
	if err != nil {
		log.Printf("Firma tohumlama/güncelleme hatası (%s): %v", vkn, err)
	} else {
		log.Printf("Firma eklendi/güncellendi: %s", title)
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
