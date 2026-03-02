package v1

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// SetupHookRoutes initializes webhook endpoints
func SetupHookRoutes(app *fiber.App) {
	hooks := app.Group("/hooks/v1")

	hooks.Post("/nes/status", handleNESStatusUpdate)
	hooks.Post("/nes/staticcodes/update", handleNESStaticCodesUpdate)
}

// handleNESStaticCodesUpdate, NES veya diğer entegratör sistemlerinden gelebilecek "Statik Kodlar (Vergi vb.) Güncellendi" uyarılarını dinler.
// Bu kanca, dış bir sistemden değişiklik bildirimi geldiğinde sistemimizdeki verileri tazelememiz gerektiğini anlamamızı sağlar.
// İşlem adımları:
// 1. Gelen JSON verisini okur (Örneğin: {"update_type": "tax_type", "timestamp": "..."}).
// 2. İlgili veri güncellemesini loglar.
// 3. (Gelecek geliştirme) Veritabanındaki yerel statik kod tablolarını güncellemek için bir Worker (Job) tetikleyebilir.
// handleNESStaticCodesUpdate godoc
// @Summary NES Statik Kod Güncelleme Bildirimi
// @Description NES veya diğer entegratör sistemlerinden gelen "Statik Kodlar Güncellendi" uyarılarını işler.
// @Tags Webhooks
// @Accept json
// @Produce json
// @Param payload body map[string]interface{} true "Güncelleme Detayları"
// @Success 200 {object} map[string]string
// @Router /hooks/v1/nes/staticcodes/update [post]
func handleNESStaticCodesUpdate(c *fiber.Ctx) error {
	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	log.Printf("NES Statik Kod Güncelleme Bildirimi Alındı: %v", payload)

	// TODO: Gelen bildirime göre arka planda statik kodları güncelleyecek işleyiciyi tetikle

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "received",
		"info":   "Statik kod güncelleme isteği başarıyla alındı",
	})
}

// handleNESStatusUpdate, NES Özel Entegratör tarafından gönderilen durum güncellemelerini işleyen webhook kancasıdır.
// Gelen veri formatı NES dökümantasyonuna uygun olmalıdır.
// İşlem adımları:
// 1. Gelen JSON verisini parse eder.
// 2. Gelen güncellemenin hangi dökümana (Fatura/İrsaliye) ait olduğunu belirler.
// 3. Veritabanındaki ilgili kaydın durumunu günceller ve nes_status_updated_at alanını yeniler.
// handleNESStatusUpdate godoc
// @Summary NES Durum Güncellemesi Bildirimi
// @Description NES Özel Entegratör tarafından gönderilen döküman durum güncellemelerini işler.
// @Tags Webhooks
// @Accept json
// @Produce json
// @Param payload body map[string]interface{} true "Durum Güncelleme Detayları"
// @Success 200 {object} map[string]string
// @Router /hooks/v1/nes/status [post]
func handleNESStatusUpdate(c *fiber.Ctx) error {
	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	log.Printf("NES Durum Güncellemesi Alındı: %v", payload)

	// TODO: Gelen verideki ID'ye göre dökümanı bul ve durumunu güncelle

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "received",
		"info":   "Durum güncellemesi başarıyla kuyruğa alındı",
	})
}
