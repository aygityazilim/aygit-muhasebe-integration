package models

// TaxType Vergi tiplerini temsil eden veri modelidir.
// NES API'si üzerinden dönen, fatura veya diğer belgelerde kullanılan vergi tiplerinin detaylarını içerir.
// Bu model sayesinde kullanılabilir vergi tiplerine ulaşım sağlanır.
type TaxType struct {
	TaxTypeCode     *string `json:"taxTypeCode"`     // Verginin sistemdeki benzersiz kodu
	TaxName         *string `json:"taxName"`         // Verginin tam adı
	IsPlus          *bool   `json:"isPlus"`          // Verginin toplama pozitif veya negatif etkisi
	DisabledPercent bool    `json:"disabledPercent"` // Yüzdesel vergi oranının devre dışı olup olmadığı
	DisabledTotal   bool    `json:"disabledTotal"`   // Toplam tutar hesabında devre dışı olup olmadığı
	Matrah          *bool   `json:"matrah"`          // Verginin bir matraha bağlanıp bağlanmadığı bilgisi
	IstisnaCheck    bool    `json:"istisnaCheck"`    // Verginin istisna kapsamında olup olmadığının kontrolü
	ShortName       *string `json:"shortName"`       // Verginin ekranda veya dökümanlarda gösterilecek kısa adı
}

// WithholdingTaxType Tevkifat tiplerini temsil eden veri modelidir.
// Fatura üzerinde uygulanan kesintiler/tevkifatlar için gerekli tanımları barındırır.
type WithholdingTaxType struct {
	TaxTypeCode     *string `json:"taxTypeCode"`     // Tevkifat tipi kodu
	TaxName         *string `json:"taxName"`         // Tevkifatın tam adı
	IsPlus          *bool   `json:"isPlus"`          // Tevkifatın faturaya olumlu/olumsuz etkisi
	DisabledPercent bool    `json:"disabledPercent"` // Yüzdesel değerin değiştirilip değiştirilemeyeceği
	DisabledTotal   bool    `json:"disabledTotal"`   // Toplam değerin değiştirilip değiştirilemeyeceği
	Matrah          *bool   `json:"matrah"`          // Matrah hesaplamasında kullanılıp kullanılmadığı
	IstisnaCheck    bool    `json:"istisnaCheck"`    // İstisna durumu olup olmadığı
	ShortName       *string `json:"shortName"`       // Tevkifatın kısa adı
}

// TaxExemptionReason Vergi muafiyet kodlarını temsil eden veri modelidir.
// Faturada bir vergiden muaf olunması durumunda girilmesi gereken yasal kod ve açıklamaları içerir.
type TaxExemptionReason struct {
	Code        *string `json:"code"`        // KDV Muafiyet kodu (Örn: 301, 350)
	Description *string `json:"description"` // Muafiyetin kısa adı veya kanun maddesi
	Detail      *string `json:"detail"`      // Muafiyetin daha detaylı açıklaması
}

// Currency Sistemde kullanılabilir para birimlerini temsil eden veri modelidir.
// Fatura kesiminde kullanılabilecek geçerli döviz kurları için gereklidir.
type Currency struct {
	Code        *string `json:"code"`        // Para birimi ISO kodu (Örn: TRY, USD, EUR)
	Description *string `json:"description"` // Para biriminin adı (Örn: Türk Lirası)
	Detail      *string `json:"detail"`      // Varsa para birimine ait ekstra detaylar
}
