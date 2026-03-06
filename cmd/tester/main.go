package main

import (
	"aygit-muhasebe-integration/config"
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/internal/service"
	"aygit-muhasebe-integration/internal/testflows"
	"aygit-muhasebe-integration/pkg/db"
	"fmt"
	"log"
)

func main() {
	// 1. Ayarları ve Veritabanını Başlat
	config.InitConfig()
	db.InitDB()
	db.SeedData()
	defer db.CloseDB()

	nesService := service.NewNESService()

	// 2. Test Firmalarını Getir
	var sender, receiver models.Company

	err := db.DB.Get(&sender, "SELECT * FROM companies WHERE party_identification = $1", "1234567801")
	if err != nil {
		log.Fatalf("Gönderici firma bulunamadı: %v", err)
	}
	err = db.DB.Get(&receiver, "SELECT * FROM companies WHERE party_identification = $1", "1234567802")
	if err != nil {
		log.Fatalf("Alıcı firma bulunamadı: %v", err)
	}

	fmt.Printf("\n--- [AYGIT BİLİŞİM TEST AKIŞI BAŞLATILIYOR] ---\n")
	fmt.Printf("Gönderici: %s (%s) (%s)\n", sender.PartyName, sender.PartyIdentification, sender.GetNesAPIKey())
	fmt.Printf("Alıcı: %s (%s) (%s)\n", receiver.PartyName, receiver.PartyIdentification, receiver.GetNesAPIKey())

	// 3. E-Fatura Testi (Kesme, Alma ve Red/İptal)
	testflows.TestEInvoiceFlow(nesService, &sender, &receiver)

	// 4. E-Arşiv Testi (Kesme ve İptal)
	testflows.TestEArchiveFlow(nesService, &sender)

	// 5. E-İrsaliye Testi (Kesme, Alma ve Red)
	testflows.TestEDespatchFlow(nesService, &sender, &receiver)

	// 6. Makbuz Testleri (E-MM / E-SMM) (Kesme ve İptal)
	emmService := service.NewNESVoucherService("https://apitest.nes.com.tr/emm")
	esmmService := service.NewNESVoucherService("https://apitest.nes.com.tr/esmm")
	testflows.TestVoucherFlow(emmService, esmmService, &sender)

	fmt.Printf("\n--- [AYGIT BİLİŞİM TEST AKIŞI TAMAMLANDI] ---\n")
}
