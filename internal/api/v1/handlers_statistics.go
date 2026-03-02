package v1

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/db"
	"aygit-muhasebe-integration/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

// getDailyStatistics godoc
// @Summary Günlük İstatistikleri Listele
// @Description Günlük kullanım istatistiklerine bu uç ile ulaşabilirsiniz. startDate ve endDate (ISO 8601) query parametreleriyle tarih aralığı verilir. (NES API entegrasyonu)
// @Tags Statistics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param startDate query string true "Başlangıç tarihi (Örn: 2025-01-26T00:00:00.0000000+03:00)"
// @Param endDate query string true "Bitiş tarihi (Örn: 2025-01-26T00:00:00.0000000+03:00)"
// @Success 200 {array} models.DailyStatistic "Başarılı istek. İstatistik sonuçları döner."
// @Failure 400 {object} errors.AppError "Eksik parametre hatası. (startDate veya endDate eksik)"
// @Failure 401 {object} errors.AppError "Yetkisiz erişim. Oturum açık değil veya firma bilgisi eksik."
// @Failure 500 {object} errors.AppError "Sunucu hatası veya entegrasyon hatası."
// @Router /v1/statistics/daily [get]
func getDailyStatistics(c *fiber.Ctx) error {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if startDate == "" || endDate == "" {
		return errors.NewError(fiber.StatusBadRequest, errors.ErrCodeInvalidRequest, "startDate ve endDate parametreleri zorunludur")
	}

	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1 AND deleted_at IS NULL", *user.CompanyID)
	if err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma yapılandırması yüklenemedi")
	}

	result, err := nesService.GetDailyStatistics(&company, startDate, endDate)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
