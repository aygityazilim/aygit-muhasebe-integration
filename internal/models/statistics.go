package models

// DailyStatistic Günlük kullanım istatistiklerini temsil eden modeldir.
// Bu model sayesinde NES sistemi üzerinden belli tarihler arasındaki kullanım verileri çekilebilir.
type DailyStatistic struct {
	Date       string      `json:"date"`       // İstatistiğin ait olduğu gün, saat ve zaman dilimi (ISO 8601 formatında)
	Statistics interface{} `json:"statistics"` // İlgili güne ait sistemden dönen istatistik detayları veya sayaçları barındırır
}
