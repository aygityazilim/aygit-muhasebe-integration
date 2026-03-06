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

func TestEDespatchFlow(s *service.NESService, sender, receiver *models.Company) {
	fmt.Printf("\n=== [E-İrsaliye Test Senaryosu] ===\n")

	ttn := uuid.New().String()
	xmlData, err := os.ReadFile("cmd/tester/upload/e_despatch_sample.xml")
	if err != nil {
		log.Fatalf("E-İrsaliye dosyası okunamadı: %v", err)
	}
	xmlDataStr := string(xmlData)
	xmlDataStr = replaceUUIDinXML(xmlDataStr, ttn)

	// 1. Gönderici İrsaliyeyi Yüklüyor (Upload)
	fmt.Println("[Gönderici] 1. E-İrsaliye yükleniyor (Upload)...")
	uploadRes, err := s.UploadDespatch(sender, []byte(xmlDataStr), map[string]string{
		"IsDirectSend": "false",
	})
	if err != nil {
		log.Printf("HATA (E-İrsaliye Upload): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", uploadRes)

	// 2. Gönderici Taslağı Onaylıyor (SendDraft)
	fmt.Println("[Gönderici] 2. E-İrsaliye taslağı onaylanıp gönderiliyor (SendDraft)...")
	sendRes, err := s.SendDraftDespatches(sender, []string{ttn})
	if err != nil {
		log.Printf("HATA (E-İrsaliye SendDraft): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", sendRes)

	// 3. Alıcı Gelen Kutusunu Kontrol Ediyor (Incoming)
	fmt.Println("[Alıcı]     3. Gelen kutusu (irsaliye) kontrol ediliyor...")
	time.Sleep(3 * time.Second) // Gönderimin tamamlanması için bekleme
	incoming, err := s.GetIncomingDespatches(receiver, map[string]string{"uuid": ttn})
	if err != nil {
		log.Printf("HATA (E-İrsaliye Incoming): %v", err)
	} else {
		fmt.Printf("   Gelen irsaliye bulundu. Yanıt özeti: %v\n", incoming)
	}

	// 4. Alıcı İrsaliyeyi Reddediyor (Answer/Reject)
	fmt.Println("[Alıcı]     4. E-İrsaliye yanıtlanıyor/reddediliyor (Answer/Reject)...")
	// NES üzerinde E-İrsaliye Red yanıtı ("RET" cevap kodu)
	answerData := map[string]interface{}{
		"ResponseCode": "RETDON", // ya da portalın kabul ettiği ret kodu, ör: "RET", "REJECT" vs.
		"Note":         "Aygıt Test İptal/Red Senaryosu",
	}
	rejectRes, err := s.SendDespatchAnswer(receiver, ttn, answerData)
	if err != nil {
		log.Printf("HATA (E-İrsaliye Red/Yanıt): %v", err)
	} else {
		fmt.Printf("   Red/İptal Yanıtı Başarılı. Yanıt: %v\n", rejectRes)
	}

	fmt.Printf("=== [E-İrsaliye Senaryosu Tamamlandı] ===\n")
}
