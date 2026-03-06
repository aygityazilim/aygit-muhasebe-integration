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

func TestEInvoiceFlow(s *service.NESService, sender, receiver *models.Company) {
	fmt.Printf("\n=== [E-Fatura Test Senaryosu] ===\n")

	ttn := uuid.New().String()
	xmlData, err := os.ReadFile("cmd/tester/upload/e_invoice_sample.xml")
	if err != nil {
		log.Fatalf("E-Fatura dosyası okunamadı: %v", err)
	}
	xmlDataStr := string(xmlData)

	// UBL içerisindeki UUID'yi dinamik olarak değiştiriyoruz ki her testte yeni bir belge oluşsun
	xmlDataStr = replaceUUIDinXML(xmlDataStr, ttn)

	// 1. Gönderici Faturayı Yüklüyor (Upload)
	fmt.Println("[Gönderici] 1. Fatura yükleniyor (Upload)...")
	uploadRes, err := s.UploadInvoice(sender, []byte(xmlDataStr), map[string]string{
		"IsDirectSend": "false",
	})
	if err != nil {
		log.Printf("HATA (Upload): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", uploadRes)

	// 2. Gönderici Taslağı Onaylıyor (SendDraft)
	fmt.Println("[Gönderici] 2. Taslak onaylanıp gönderiliyor (SendDraft)...")
	sendRes, err := s.SendDraftInvoices(sender, []string{ttn})
	if err != nil {
		log.Printf("HATA (SendDraft): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", sendRes)

	// 3. Alıcı Gelen Kutusunu Kontrol Ediyor (Incoming)
	fmt.Println("[Alıcı]     3. Gelen kutusu kontrol ediliyor...")
	time.Sleep(3 * time.Second) // NES sisteminin işlemesi için bekleme süresi
	incoming, err := s.GetIncomingInvoices(receiver, map[string]string{"uuid": ttn})
	if err != nil {
		log.Printf("HATA (Incoming): %v", err)
	} else {
		fmt.Printf("   Gelen faturada bulundu. Yanıt özeti: %v\n", incoming)
	}

	// 4. Alıcı Faturayı Reddediyor / İptal Ediyor (Reject)
	fmt.Println("[Alıcı]     4. Fatura reddediliyor (Reject)...")
	rejectRes, err := s.RejectEInvoice(receiver, ttn, "Aygıt Test İptal/Red Senaryosu")
	if err != nil {
		log.Printf("HATA (Reject): %v", err)
	} else {
		fmt.Printf("   Red/İptal Başarılı. Yanıt: %v\n", rejectRes)
	}
	fmt.Printf("=== [E-Fatura Senaryosu Tamamlandı] ===\n")
}
