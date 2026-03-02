package v1

import (
	"log"
	"time"
)

// StartWorkers, uygulama başlatıldığında tüm arka plan görevlerini (goroutine) başlatır.
func StartWorkers() {
	go documentProcessor()
	go syncStaticCodesWorker()
	go syncDailyStatisticsWorker()
}

// syncStaticCodesWorker, dış API'deki (NES vs.) güncel statik kodları (vergi tipleri, para birimleri vs.)
// belirli aralıklarla (örneğin her gece 03:00'te) sorgulayıp kendi veritabanımızla senkronize etmeyi hedefler.
// Böylelikle Gib/NES üzerinde yapılan kod değişiklikleri, uygulamamıza gecikmesiz olarak yansır.
// İşlem Mantığı:
// 1. Her gün bir kez (veya belirtilen aralıkla) çalışacak bir ticker kurulur.
// 2. NES API'sine getTaxTypes gibi istekler atılır.
// 3. Veritabanındaki "static_codes" benzeri tablolar, gelen bu yeni verilerle upsert edilir.
func syncStaticCodesWorker() {
	// Not: 24 Saatte bir çalışması simüle edildi
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	log.Println("Statik Kod Senkronizasyon (Static Codes Sync) worker'ı başlatıldı")

	for range ticker.C {
		log.Println("Güncel statik kodlar çekiliyor ve veritabanıyla senkronize ediliyor...")
		// TODO: db'deki ilgili tabloları NES_Service aracılığıyla alınan datalarla güncelle
	}
}

// syncDailyStatisticsWorker, istatistikleri günlük veya saatlik periyotlarda
// dış entegrasyondan çekerek sistem üzerinde raporlama yapılması için yerel bir tabloya yazmayı hedefler.
// İşlem Mantığı:
// 1. Düzenli aralıklarla (örneğin 1 saat) dış servisten (NES'ten) veriler çekilir.
// 2. Yerel dashboard ve istatistik ekranlarının hızlı çalışabilmesi için veritabanında "statistics" tablolarına kaydedilir.
func syncDailyStatisticsWorker() {
	// Not: 1 saatte bir çalışması simüle edildi
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	log.Println("Günlük İstatistik Senkronizasyon (Daily Statistics Sync) worker'ı başlatıldı")

	for range ticker.C {
		log.Println("Dış entegrasyondan son istatistik verileri toplanıyor...")
		// TODO: Son 24 saatin startDate ve endDate'ini hesapla
		// TODO: nesService.GetDailyStatistics çağrısı yap ve db'ye yaz
	}
}

// documentProcessor, veritabanındaki gönderilmeyi bekleyen veya durumu belirsiz olan dökümanları
// periyodik olarak kontrol eder ve NES/GİB sistemlerine iletimini sağlar.
// Çalışma Mantığı:
// 1. 10 dakikalık bir ticker ile döngüye girer.
// 2. Her döngüde "BEKLIYOR" durumundaki fatura/irsaliyeleri sorgular.
// 3. İlgili firmanın NES ayarlarını kullanarak gönderim işlemini başlatır.
func documentProcessor() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	log.Println("Döküman İşleme (Document Processor) worker'ı başlatıldı")

	for range ticker.C {
		log.Println("Gönderilmeyi bekleyen dökümanlar taranıyor...")
		// TODO: DB'den bekleyen dökümanları (E-Fatura/E-İrsaliye) çek
		// TODO: Her döküman için ilgili firmanın NES bilgilerini al ve iletimi başlat
	}
}
