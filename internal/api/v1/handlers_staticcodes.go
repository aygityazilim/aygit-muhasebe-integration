package v1

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/db"
	"aygit-muhasebe-integration/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

// getTaxTypes godoc
// @Summary Vergi Tiplerini Listele
// @Description Sistemde kullanılabilir vergi tiplerine bu uç ile ulaşabilirsiniz. (NES API entegrasyonu)
// @Tags StaticCodes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.TaxType "Başarılı istek. Vergi tiplerinin listesi döner."
// @Failure 401 {object} errors.AppError "Yetkisiz erişim. Oturum açık değil veya firma bilgisi eksik."
// @Failure 500 {object} errors.AppError "Sunucu hatası veya entegrasyon hatası."
// @Router /v1/staticcodes/taxtype [get]
func getTaxTypes(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1 AND deleted_at IS NULL", *user.CompanyID)
	if err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma yapılandırması yüklenemedi")
	}

	result, err := nesService.GetTaxTypes(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// getWithholdingTaxTypes godoc
// @Summary Tevkifat Tiplerini Listele
// @Description Sistemde kullanılabilir tevkifat tiplerine bu uç ile ulaşabilirsiniz. (NES API entegrasyonu)
// @Tags StaticCodes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.WithholdingTaxType "Başarılı istek. Tevkifat tiplerinin listesi döner."
// @Failure 401 {object} errors.AppError "Yetkisiz erişim. Oturum açık değil veya firma bilgisi eksik."
// @Failure 500 {object} errors.AppError "Sunucu hatası veya entegrasyon hatası."
// @Router /v1/staticcodes/withholdingtaxtype [get]
func getWithholdingTaxTypes(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1 AND deleted_at IS NULL", *user.CompanyID)
	if err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma yapılandırması yüklenemedi")
	}

	result, err := nesService.GetWithholdingTaxTypes(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// getTaxExemptionReasons godoc
// @Summary Vergi Muafiyet Kodlarını Listele
// @Description Sistemde kullanılabilir vergi muafiyetlerine bu uç ile ulaşabilirsiniz. (NES API entegrasyonu)
// @Tags StaticCodes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.TaxExemptionReason "Başarılı istek. Vergi muafiyet kodlarının listesi döner."
// @Failure 401 {object} errors.AppError "Yetkisiz erişim. Oturum açık değil veya firma bilgisi eksik."
// @Failure 500 {object} errors.AppError "Sunucu hatası veya entegrasyon hatası."
// @Router /v1/staticcodes/taxexemptionreason [get]
func getTaxExemptionReasons(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1 AND deleted_at IS NULL", *user.CompanyID)
	if err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma yapılandırması yüklenemedi")
	}

	result, err := nesService.GetTaxExemptionReasons(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// getCurrencies godoc
// @Summary Para Birimlerini Listele
// @Description Sistemde kullanılabilir para birimilerine bu uç ile ulaşabilirsiniz. (NES API entegrasyonu)
// @Tags StaticCodes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Currency "Başarılı istek. Para birimlerinin listesi döner."
// @Failure 401 {object} errors.AppError "Yetkisiz erişim. Oturum açık değil veya firma bilgisi eksik."
// @Failure 500 {object} errors.AppError "Sunucu hatası veya entegrasyon hatası."
// @Router /v1/staticcodes/currency [get]
func getCurrencies(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1 AND deleted_at IS NULL", *user.CompanyID)
	if err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma yapılandırması yüklenemedi")
	}

	result, err := nesService.GetCurrencies(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
