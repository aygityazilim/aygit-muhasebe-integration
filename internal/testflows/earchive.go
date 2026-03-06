package testflows

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/internal/service"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func TestEArchiveFlow(s *service.NESService, sender *models.Company) {
	fmt.Printf("\n=== [E-Arşiv Fatura Test Senaryosu] ===\n")

	ttn := uuid.New().String()
	xmlData, err := os.ReadFile("cmd/tester/upload/e_archive_sample.xml")
	if err != nil {
		log.Fatalf("E-Arşiv dosyası okunamadı: %v", err)
	}
	xmlDataStr := string(xmlData)
	xmlDataStr = replaceUUIDinXML(xmlDataStr, ttn)

	// 1. Gönderici Faturayı Yüklüyor (Upload)
	fmt.Println("[Gönderici] 1. E-Arşiv faturası yükleniyor (Upload)...")
	uploadRes, err := s.UploadEArchiveInvoice(sender, []byte(xmlDataStr), map[string]string{
		"IsDirectSend": "false",
	})
	if err != nil {
		log.Printf("HATA (E-Arşiv Upload): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", uploadRes)

	// 2. Gönderici Taslağı Onaylıyor (SendDraft)
	fmt.Println("[Gönderici] 2. E-Arşiv taslağı onaylanıp gönderiliyor (SendDraft)...")
	sendRes, err := s.SendDraftEArchiveInvoices(sender, []string{ttn})
	if err != nil {
		log.Printf("HATA (E-Arşiv SendDraft): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", sendRes)

	// 3. Gönderici İptal Ediyor (Cancel)
	fmt.Println("[Gönderici] 3. E-Arşiv faturası iptal ediliyor (Cancel)...")
	time.Sleep(3 * time.Second) // Gönderimin tamamlanması için bekleme
	cancelData := map[string]interface{}{
		"uuids":  []string{ttn},
		"reason": "Aygıt Test İptal/Red Senaryosu",
	}
	cancelRes, err := s.CancelEArchiveInvoice(sender, cancelData)
	if err != nil {
		log.Printf("HATA (E-Arşiv İptal): %v", err)
	} else {
		fmt.Printf("   İptal Başarılı. Yanıt: %v\n", cancelRes)
	}

	fmt.Printf("=== [E-Arşiv Senaryosu Tamamlandı] ===\n")
}
