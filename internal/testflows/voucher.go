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

func TestVoucherFlow(emmService, esmmService *service.NESVoucherService, sender *models.Company) {
	fmt.Printf("\n=== [E-Müstahsil Makbuzu (E-MM) Test Senaryosu] ===\n")
	testSingleVoucherFlow(emmService, sender, "cmd/tester/upload/e_mm_sample.json", "E-MM")

	fmt.Printf("\n=== [E-Serbest Meslek Makbuzu (E-SMM) Test Senaryosu] ===\n")
	testSingleVoucherFlow(esmmService, sender, "cmd/tester/upload/e_smm_sample.json", "E-SMM")
}

func testSingleVoucherFlow(s *service.NESVoucherService, sender *models.Company, sampleFilePath, docType string) {
	ttn := uuid.New().String()
	jsonData, err := os.ReadFile(sampleFilePath)
	if err != nil {
		log.Fatalf("%s dosyası okunamadı: %v", docType, err)
	}

	nowStr := time.Now().Format("2006-01-02")
	jsonStr := fmt.Sprintf(string(jsonData), ttn, nowStr)

	// 1. Gönderici Makbuzu Yüklüyor (Upload)
	fmt.Printf("[Gönderici] 1. %s makbuzu yükleniyor (Upload)...\n", docType)
	uploadRes, err := s.UploadDocument(sender, []byte(jsonStr))
	if err != nil {
		log.Printf("HATA (%s Upload): %v", docType, err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", uploadRes)

	// 2. Gönderici Taslağı Listeliyor (GetDrafts)
	fmt.Printf("[Gönderici] 2. %s taslakları listeleniyor...\n", docType)
	drafts, err := s.GetDraftVouchers(sender)
	if err != nil {
		log.Printf("HATA (%s GetDrafts): %v", docType, err)
	} else {
		fmt.Printf("   Taslaklar bulundu. Adet: %v\n", len(drafts))
	}

	// 3. Gönderici Taslağı Onaylıyor (SendDraft)
	fmt.Printf("[Gönderici] 3. %s taslağı onaylanıp gönderiliyor...\n", docType)
	sendRes, err := s.SendDraftVoucher(sender, []string{ttn})
	if err != nil {
		log.Printf("HATA (%s SendDraft): %v", docType, err)
	} else {
		fmt.Printf("   Başarılı. Yanıt: %v\n", sendRes)
	}

	// 4. Gönderici Makbuzu İptal Ediyor (Cancel)
	fmt.Printf("[Gönderici] 4. %s makbuzu iptal ediliyor...\n", docType)
	time.Sleep(3 * time.Second) // İşlenmesi için bekleme
	cancelRes, err := s.CancelVoucher(sender, ttn, "Aygıt Test İptal/Red Senaryosu")
	if err != nil {
		log.Printf("HATA (%s Cancel): %v", docType, err)
	} else {
		fmt.Printf("   İptal Başarılı. Yanıt: %v\n", cancelRes)
	}

	fmt.Printf("=== [%s Senaryosu Tamamlandı] ===\n", docType)
}
