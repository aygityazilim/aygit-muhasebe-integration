package v1

import (
	"encoding/json"
	"fmt"
	"io"

	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/db"
	"aygit-muhasebe-integration/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

// ---------------------------------------------------------------------------------------------------------------------
// CUSTOMIZATION SETTINGS
// ---------------------------------------------------------------------------------------------------------------------

// GetCustomizationSettings godoc
// @Summary Tasarım ayarlarını döner
// @Tags Definitions
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/customizationsettings [get]
func GetCustomizationSettings(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.GetCustomizationSettings(&company)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

// CreateCustomizationSetting godoc
// @Summary e-Belge özelleştirilebilir tasarım ekler
// @Tags Definitions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Setting data"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/customizationsettings [post]
func CreateCustomizationSetting(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return errors.NewError(fiber.StatusBadRequest, errors.ErrCodeInvalidRequest, "Geçersiz istek")
	}

	res, err := nesService.CreateCustomizationSetting(&company, payload)
	if err != nil {
		return err
	}

	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "CREATE_CUSTOMIZATION_SETTING", func() string { b, _ := json.Marshal(payload); return string(b) }())
	return c.JSON(res)
}

// GetCustomizationSettingByID godoc
// @Summary Sorgulanan ayarı getirir
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/customizationsettings/{id} [get]
func GetCustomizationSettingByID(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.GetCustomizationSettingByID(&company, c.Params("id"))
	if err != nil {
		return err
	}
	return c.JSON(res)
}

// UpdateCustomizationSetting godoc
// @Summary e-Belge özelleştirilebilir tasarımını günceller
// @Tags Definitions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param request body map[string]interface{} true "Setting data"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/customizationsettings/{id} [put]
func UpdateCustomizationSetting(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return errors.NewError(fiber.StatusBadRequest, errors.ErrCodeInvalidRequest, "Geçersiz istek")
	}

	res, err := nesService.UpdateCustomizationSetting(&company, c.Params("id"), payload)
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "UPDATE_CUSTOMIZATION_SETTING", func() string { b, _ := json.Marshal(payload); return string(b) }())
	return c.JSON(res)
}

// DeleteCustomizationSetting godoc
// @Summary e-Belge özelleştirilebilir tasarımını siler
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/customizationsettings/{id} [delete]
func DeleteCustomizationSetting(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.DeleteCustomizationSetting(&company, c.Params("id"))
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETE_CUSTOMIZATION_SETTING", fmt.Sprintf(`{"id": "%s"}`, c.Params("id")))
	return c.JSON(res)
}

// SetDefaultCustomizationSetting godoc
// @Summary Varsayılan ayarlar
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/customizationsettings/{id}/setdefault [get]
func SetDefaultCustomizationSetting(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.SetDefaultCustomizationSetting(&company, c.Params("id"))
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "SET_DEFAULT_CUSTOMIZATION_SETTING", fmt.Sprintf(`{"id": "%s"}`, c.Params("id")))
	return c.JSON(res)
}

// PreviewCustomizationSetting godoc
// @Summary Tasarımı önizler
// @Tags Definitions
// @Accept json
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param request body map[string]interface{} true "Preview payload"
// @Success 200 {string} binary
// @Router /v1/definitions/documenttemplates/customizationsettings/{id}/preview [post]
func PreviewCustomizationSetting(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return errors.NewError(fiber.StatusBadRequest, errors.ErrCodeInvalidRequest, "Geçersiz istek")
	}

	content, contentType, err := nesService.PreviewCustomizationSetting(&company, c.Params("id"), payload)
	if err != nil {
		return err
	}
	c.Set("Content-Type", contentType)
	return c.Send(content)
}

// GetCustomizationSettingImage godoc
// @Summary Tasarıma eklenmiş logoya/kaşeye/imzaya ulaşır
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param imageType path string true "logo, stamp veya signature"
// @Success 200 {string} binary
// @Router /v1/definitions/documenttemplates/customizationsettings/{id}/{imageType} [get]
func GetCustomizationSettingImage(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	content, contentType, err := nesService.GetCustomizationSettingImage(&company, c.Params("id"), c.Params("imageType"))
	if err != nil {
		return err
	}
	c.Set("Content-Type", contentType)
	return c.Send(content)
}

// DeleteCustomizationSettingImage godoc
// @Summary Tasarıma eklenmiş logo/kaşe/imzayı siler
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param imageType path string true "logo, stamp veya signature"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/customizationsettings/{id}/{imageType} [delete]
func DeleteCustomizationSettingImage(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.DeleteCustomizationSettingImage(&company, c.Params("id"), c.Params("imageType"))
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETE_CUSTOMIZATION_SETTING_IMAGE", fmt.Sprintf(`{"id": "%s", "type": "%s"}`, c.Params("id"), c.Params("imageType")))
	return c.JSON(res)
}

// UploadCustomizationSettingImage godoc
// @Summary Tasarıma logo/kaşe/imza ekler
// @Tags Definitions
// @Accept multipart/form-data
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param imageType path string true "logo, stamp veya signature"
// @Param File formData file true "Dosya"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/customizationsettings/{id}/{imageType} [post]
func UploadCustomizationSettingImage(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	file, err := c.FormFile("File")
	if err != nil {
		return err
	}
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	fileData, _ := io.ReadAll(f)

	res, err := nesService.UploadCustomizationSettingImage(&company, c.Params("id"), c.Params("imageType"), fileData, file.Filename)
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "UPLOAD_CUSTOMIZATION_SETTING_IMAGE", fmt.Sprintf(`{"id": "%s", "type": "%s"}`, c.Params("id"), c.Params("imageType")))
	return c.JSON(res)
}

// ---------------------------------------------------------------------------------------------------------------------
// SERIES & ANSWER SERIES
// ---------------------------------------------------------------------------------------------------------------------

func getSeriesTypeFromPath(c *fiber.Ctx) string {
	// e.g. /v1/definitions/series or /v1/answerseries
	if len(c.Route().Path) > 15 && c.Route().Path[:15] == "/v1/answerseries" {
		return "answerseries"
	}
	return "definitions/series"
}

// GetSeries godoc
// @Summary Serileri listeler
// @Tags Definitions
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/series [get]
// @Router /v1/answerseries [get]
func GetSeries(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.GetSeries(&company, getSeriesTypeFromPath(c))
	if err != nil {
		return err
	}
	return c.JSON(res)
}

// CreateSeries godoc
// @Summary Seri ekler
// @Tags Definitions
// @Accept json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Series payload"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/series [post]
// @Router /v1/answerseries [post]
func CreateSeries(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return errors.NewError(fiber.StatusBadRequest, errors.ErrCodeInvalidRequest, "Geçersiz istek")
	}

	res, err := nesService.CreateSeries(&company, getSeriesTypeFromPath(c), payload)
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "CREATE_SERIES", func() string { b, _ := json.Marshal(payload); return string(b) }())
	return c.JSON(res)
}

// GetSeriesByIDOrPrefix godoc
// @Summary Sorgulanan seriyi getirir
// @Tags Definitions
// @Security BearerAuth
// @Param param path string true "Series ID veya Prefix"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/series/{param} [get]
// @Router /v1/answerseries/{param} [get]
func GetSeriesByIDOrPrefix(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	param := c.Params("param")
	// NESService GetSeriesByID also works for prefix since path is the same in NES
	res, err := nesService.GetSeriesByID(&company, getSeriesTypeFromPath(c), param)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

// DeleteSeries godoc
// @Summary Seri siler
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Series ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/series/{id} [delete]
// @Router /v1/answerseries/{id} [delete]
func DeleteSeries(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.DeleteSeries(&company, getSeriesTypeFromPath(c), c.Params("id"))
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETE_SERIES", fmt.Sprintf(`{"id": "%s"}`, c.Params("id")))
	return c.JSON(res)
}

// SetSeriesStatus godoc
// @Summary Seri durumunu günceller
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Series ID"
// @Param status path string true "Status"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/series/{id}/set/{status} [get]
// @Router /v1/answerseries/{id}/set/{status} [get]
func SetSeriesStatus(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.SetSeriesStatus(&company, getSeriesTypeFromPath(c), c.Params("id"), c.Params("status"))
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "SET_SERIES_STATUS", fmt.Sprintf(`{"id": "%s", "status": "%s"}`, c.Params("id"), c.Params("status")))
	return c.JSON(res)
}

// SetDefaultSeries godoc
// @Summary Seriyi varsayılan ayarlar
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Series ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/series/{id}/setdefault [get]
// @Router /v1/answerseries/{id}/setdefault [get]
func SetDefaultSeries(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.SetDefaultSeries(&company, getSeriesTypeFromPath(c), c.Params("id"))
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "SET_DEFAULT_SERIES", fmt.Sprintf(`{"id": "%s"}`, c.Params("id")))
	return c.JSON(res)
}

// SetSeriesNextNumber godoc
// @Summary Sayaç günceller
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Series ID"
// @Param year path string true "Year"
// @Param nextNumber path string true "Next Number"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/series/{id}/{year}/setnumber/{nextNumber} [get]
// @Router /v1/answerseries/{id}/{year}/setnumber/{nextNumber} [get]
func SetSeriesNextNumber(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.SetSeriesNextNumber(&company, getSeriesTypeFromPath(c), c.Params("id"), c.Params("year"), c.Params("nextNumber"))
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "SET_SERIES_NEXT_NUMBER", fmt.Sprintf(`{"id": "%s", "year": "%s", "nextNumber": "%s"}`, c.Params("id"), c.Params("year"), c.Params("nextNumber")))
	return c.JSON(res)
}

// GetSeriesHistories godoc
// @Summary Sayaç geçmişi
// @Tags Definitions
// @Security BearerAuth
// @Param serieId path string true "Series ID"
// @Param year path string true "Year"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/series/{serieId}/{year}/histories [get]
// @Router /v1/answerseries/{serieId}/{year}/histories [get]
func GetSeriesHistories(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.GetSeriesHistories(&company, getSeriesTypeFromPath(c), c.Params("serieId"), c.Params("year"))
	if err != nil {
		return err
	}
	return c.JSON(res)
}

// ---------------------------------------------------------------------------------------------------------------------
// DOCUMENT TEMPLATES & ANSWER DOCUMENT TEMPLATES
// ---------------------------------------------------------------------------------------------------------------------

func getTemplateTypeFromPath(c *fiber.Ctx) string {
	if len(c.Route().Path) > 36 && c.Route().Path[:36] == "/v1/definitions/answerdocumenttemplates" {
		return "definitions/answerdocumenttemplates"
	}
	return "definitions/documenttemplates"
}

// GetDocumentTemplates godoc
// @Summary Tasarımları listeler
// @Tags Definitions
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates [get]
// @Router /v1/definitions/answerdocumenttemplates [get]
func GetDocumentTemplates(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.GetDocumentTemplates(&company, getTemplateTypeFromPath(c))
	if err != nil {
		return err
	}
	return c.JSON(res)
}

// CreateDocumentTemplate godoc
// @Summary Tasarım ekler
// @Tags Definitions
// @Accept multipart/form-data
// @Security BearerAuth
// @Param File formData file true "Dosya"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates [post]
// @Router /v1/definitions/answerdocumenttemplates [post]
func CreateDocumentTemplate(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	file, err := c.FormFile("File")
	if err != nil {
		return err
	}
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	fileData, _ := io.ReadAll(f)

	// Handle additional form fields for params (e.g. SettingName)
	params := make(map[string]string)
	form, _ := c.MultipartForm()
	for key, values := range form.Value {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	res, err := nesService.CreateDocumentTemplate(&company, getTemplateTypeFromPath(c), fileData, file.Filename, params)
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "CREATE_DOCUMENT_TEMPLATE", fmt.Sprintf(`{"filename": "%s"}`, file.Filename))
	return c.JSON(res)
}

// DownloadDocumentTemplate godoc
// @Summary Tasarım dosyasını indirir
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {string} binary
// @Router /v1/definitions/documenttemplates/{id} [get]
// @Router /v1/definitions/answerdocumenttemplates/{id} [get]
func DownloadDocumentTemplate(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	content, contentType, err := nesService.DownloadDocumentTemplate(&company, getTemplateTypeFromPath(c), c.Params("id"))
	if err != nil {
		return err
	}
	c.Set("Content-Type", contentType)
	return c.Send(content)
}

// UpdateDocumentTemplate godoc
// @Summary Tasarımı günceller
// @Tags Definitions
// @Accept multipart/form-data
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Param File formData file true "Dosya"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/{id} [put]
// @Router /v1/definitions/answerdocumenttemplates/{id} [put]
func UpdateDocumentTemplate(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	file, err := c.FormFile("File")
	if err != nil {
		return err
	}
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	fileData, _ := io.ReadAll(f)

	params := make(map[string]string)
	form, _ := c.MultipartForm()
	for key, values := range form.Value {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	res, err := nesService.UpdateDocumentTemplate(&company, getTemplateTypeFromPath(c), c.Params("id"), fileData, file.Filename, params)
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "UPDATE_DOCUMENT_TEMPLATE", fmt.Sprintf(`{"id": "%s"}`, c.Params("id")))
	return c.JSON(res)
}

// DeleteDocumentTemplate godoc
// @Summary Tasarımı siler
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/{id} [delete]
// @Router /v1/definitions/answerdocumenttemplates/{id} [delete]
func DeleteDocumentTemplate(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.DeleteDocumentTemplate(&company, getTemplateTypeFromPath(c), c.Params("id"))
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETE_DOCUMENT_TEMPLATE", fmt.Sprintf(`{"id": "%s"}`, c.Params("id")))
	return c.JSON(res)
}

// SetDefaultDocumentTemplate godoc
// @Summary Tasarımı varsayılan ayarlar
// @Tags Definitions
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/definitions/documenttemplates/{id}/setdefault [get]
// @Router /v1/definitions/answerdocumenttemplates/{id}/setdefault [get]
func SetDefaultDocumentTemplate(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.SetDefaultDocumentTemplate(&company, getTemplateTypeFromPath(c), c.Params("id"))
	if err != nil {
		return err
	}
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "SET_DEFAULT_DOCUMENT_TEMPLATE", fmt.Sprintf(`{"id": "%s"}`, c.Params("id")))
	return c.JSON(res)
}

// PreviewDocumentTemplate godoc
// @Summary Tasarımı önizler
// @Tags Definitions
// @Accept json
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Param request body map[string]interface{} true "Preview payload"
// @Success 200 {string} binary
// @Router /v1/definitions/documenttemplates/{id}/preview [post]
// @Router /v1/definitions/answerdocumenttemplates/{id}/preview [post]
func PreviewDocumentTemplate(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return errors.NewError(fiber.StatusBadRequest, errors.ErrCodeInvalidRequest, "Geçersiz istek")
	}

	content, contentType, err := nesService.PreviewDocumentTemplate(&company, getTemplateTypeFromPath(c), c.Params("id"), payload)
	if err != nil {
		return err
	}
	c.Set("Content-Type", contentType)
	return c.Send(content)
}

// --- File Export Titles ---

// @Summary      Get File Export Title Keys
// @Description  Kullanılabilir alanları listeler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        documentType path string true "Document Type"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/fileexporttitles/{documentType}/titlekeys [get]
func (h *VoucherHandler) GetFileExportTitleKeys(c *fiber.Ctx) error {
	docType := c.Params("documentType")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/fileexporttitles/%s/titlekeys", docType))
}

// @Summary      Get File Export Titles
// @Description  Tanımları getirir
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        documentType path string true "Document Type"
// @Param        extension path string true "Extension"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/fileexporttitles/{documentType}/{extension} [get]
func (h *VoucherHandler) GetFileExportTitles(c *fiber.Ctx) error {
	docType := c.Params("documentType")
	ext := c.Params("extension")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/fileexporttitles/%s/%s", docType, ext))
}

// @Summary      Update File Export Titles
// @Description  Tanımları günceller
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/fileexporttitles [put]
func (h *VoucherHandler) UpdateFileExportTitles(c *fiber.Ctx) error {
	return h.forwardRequest(c, "PUT", "/v1/definitions/fileexporttitles")
}

// --- Mailing Settings ---

// @Summary      Get Email Settings
// @Description  Mail ayarlarını getirir
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/mailing/email/settings [get]
func (h *VoucherHandler) GetEmailSettings(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/definitions/mailing/email/settings")
}

// @Summary      Update Email Settings
// @Description  Mail ayarlarını günceller
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/mailing/email/settings [put]
func (h *VoucherHandler) UpdateEmailSettings(c *fiber.Ctx) error {
	return h.forwardRequest(c, "PUT", "/v1/definitions/mailing/email/settings")
}

// @Summary      Get SMS Settings
// @Description  Sms ayarlarını getirir
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/mailing/sms/settings [get]
func (h *VoucherHandler) GetSMSSettings(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/definitions/mailing/sms/settings")
}

// @Summary      Update SMS Settings
// @Description  Sms ayarlarını günceller
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/mailing/sms/settings [put]
func (h *VoucherHandler) UpdateSMSSettings(c *fiber.Ctx) error {
	return h.forwardRequest(c, "PUT", "/v1/definitions/mailing/sms/settings")
}

// --- Document Templates Customization Settings ---

// @Summary      Get Customization Settings
// @Description  Tasarım ayarları dönülür
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings [get]
func (h *VoucherHandler) GetCustomizationSettings(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/definitions/documenttemplates/customizationsettings")
}

// @Summary      Create Customization Setting
// @Description  e-Belge özelleştirilebilir tasarım eklemek için kullanılır
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings [post]
func (h *VoucherHandler) CreateCustomizationSetting(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/definitions/documenttemplates/customizationsettings")
}

// @Summary      Get Customization Setting By ID
// @Description  Sorgulanan ayarı getirir
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id} [get]
func (h *VoucherHandler) GetCustomizationSettingByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s", id))
}

// @Summary      Update Customization Setting
// @Description  e-Belge özelleştirilebilir tasarımını güncellemek için kullanılır
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id} [put]
func (h *VoucherHandler) UpdateCustomizationSetting(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "PUT", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s", id))
}

// @Summary      Delete Customization Setting
// @Description  e-Belge özelleştirilebilir tasarımını silmek için kullanılır
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id} [delete]
func (h *VoucherHandler) DeleteCustomizationSetting(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s", id))
}

// @Summary      Set Customization Setting Default
// @Description  Varsayılan olarak ayarlar
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/setdefault [get]
func (h *VoucherHandler) SetCustomizationSettingDefault(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/setdefault", id))
}

// @Summary      Preview Customization Setting
// @Description  Tasarımı önizler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/preview [post]
func (h *VoucherHandler) PreviewCustomizationSetting(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "POST", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/preview", id))
}

// @Summary      Get Customization Setting Logo
// @Description  e-Belge özelleştirilebilir tasarıma eklenmiş olan logoya ulaşılır
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/logo [get]
func (h *VoucherHandler) GetCustomizationSettingLogo(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/logo", id))
}

// @Summary      Upload Customization Setting Logo
// @Description  e-Belge özelleştirilebilir tasarıma logo ekler
// @Tags         Definitions
// @Accept       multipart/form-data
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        file formData file true "Logo File"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/logo [post]
func (h *VoucherHandler) UploadCustomizationSettingLogo(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "POST", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/logo", id))
}

// @Summary      Delete Customization Setting Logo
// @Description  e-Belge özelleştirilebilir tasarıma eklenmiş logoyu siler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/logo [delete]
func (h *VoucherHandler) DeleteCustomizationSettingLogo(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/logo", id))
}

// @Summary      Get Customization Setting Stamp
// @Description  e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeye ulaşılır
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/stamp [get]
func (h *VoucherHandler) GetCustomizationSettingStamp(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/stamp", id))
}

// @Summary      Upload Customization Setting Stamp
// @Description  e-Belge özelleştirilebilir tasarıma kaşe ekler
// @Tags         Definitions
// @Accept       multipart/form-data
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        file formData file true "Stamp File"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/stamp [post]
func (h *VoucherHandler) UploadCustomizationSettingStamp(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "POST", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/stamp", id))
}

// @Summary      Delete Customization Setting Stamp
// @Description  e-Belge özelleştirilebilir tasarıma eklenmiş kaşeyi siler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/stamp [delete]
func (h *VoucherHandler) DeleteCustomizationSettingStamp(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/stamp", id))
}

// @Summary      Get Customization Setting Signature
// @Description  e-Belge özelleştirilebilir tasarıma eklenmiş olan imzaya ulaşılır
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/signature [get]
func (h *VoucherHandler) GetCustomizationSettingSignature(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/signature", id))
}

// @Summary      Upload Customization Setting Signature
// @Description  e-Belge özelleştirilebilir tasarıma imza ekler
// @Tags         Definitions
// @Accept       multipart/form-data
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        file formData file true "Signature File"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/signature [post]
func (h *VoucherHandler) UploadCustomizationSettingSignature(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "POST", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/signature", id))
}

// @Summary      Delete Customization Setting Signature
// @Description  e-Belge özelleştirilebilir tasarıma eklenmiş imzayı siler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/customizationsettings/{id}/signature [delete]
func (h *VoucherHandler) DeleteCustomizationSettingSignature(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/definitions/documenttemplates/customizationsettings/%s/signature", id))
}

// --- Series ---

// @Summary      Get Series
// @Description  Serileri listeler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/series [get]
func (h *VoucherHandler) GetSeries(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/definitions/series")
}

// @Summary      Create Series
// @Description  Seri ekler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/series [post]
func (h *VoucherHandler) CreateSeries(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/definitions/series")
}

// @Summary      Get Series By ID
// @Description  Sorgulanan seriyi getirir
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/series/{id} [get]
func (h *VoucherHandler) GetSeriesByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/series/%s", id))
}

// @Summary      Delete Series
// @Description  Seri siler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/series/{id} [delete]
func (h *VoucherHandler) DeleteSeries(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/definitions/series/%s", id))
}

// @Summary      Get Series By Prefix
// @Description  Ön eke göre seriyi getirir
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        serie path string true "Serie Prefix"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/series/{serie} [get]
func (h *VoucherHandler) GetSeriesByPrefix(c *fiber.Ctx) error {
	serie := c.Params("serie")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/series/%s", serie))
}

// @Summary      Update Series Status
// @Description  Seri durumunu günceller
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        status path string true "Status"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/series/{id}/set/{status} [get]
func (h *VoucherHandler) UpdateSeriesStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.Params("status")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/series/%s/set/%s", id, status))
}

// @Summary      Set Series Default
// @Description  Seriyi varsayılan ayarlar
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/series/{id}/setdefault [get]
func (h *VoucherHandler) SetSeriesDefault(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/series/%s/setdefault", id))
}

// @Summary      Update Series Number
// @Description  Sayaç günceller
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        year path string true "Year"
// @Param        nextNumber path string true "Next Number"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/series/{id}/{year}/setnumber/{nextNumber} [get]
func (h *VoucherHandler) UpdateSeriesNumber(c *fiber.Ctx) error {
	id := c.Params("id")
	year := c.Params("year")
	next := c.Params("nextNumber")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/series/%s/%s/setnumber/%s", id, year, next))
}

// @Summary      Get Series Histories
// @Description  Sayaç geçmişi
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        serieId path string true "Serie ID"
// @Param        year path string true "Year"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/series/{serieId}/{year}/histories [get]
func (h *VoucherHandler) GetSeriesHistories(c *fiber.Ctx) error {
	id := c.Params("serieId")
	year := c.Params("year")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/series/%s/%s/histories", id, year))
}

// --- Document Templates ---

// @Summary      Get Document Templates
// @Description  Tasarımları listeler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates [get]
func (h *VoucherHandler) GetDocumentTemplates(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/definitions/documenttemplates")
}

// @Summary      Create Document Template
// @Description  Tasarım ekler
// @Tags         Definitions
// @Accept       multipart/form-data
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        file formData file true "Template File"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates [post]
func (h *VoucherHandler) CreateDocumentTemplate(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/definitions/documenttemplates")
}

// @Summary      Download Document Template
// @Description  Tasarım dosyasını indirir
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/{id} [get]
func (h *VoucherHandler) DownloadDocumentTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/documenttemplates/%s", id))
}

// @Summary      Update Document Template
// @Description  Tasarımı günceller
// @Tags         Definitions
// @Accept       multipart/form-data
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        file formData file true "Template File"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/{id} [put]
func (h *VoucherHandler) UpdateDocumentTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "PUT", fmt.Sprintf("/v1/definitions/documenttemplates/%s", id))
}

// @Summary      Delete Document Template
// @Description  Tasarımı siler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/{id} [delete]
func (h *VoucherHandler) DeleteDocumentTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/definitions/documenttemplates/%s", id))
}

// @Summary      Set Document Template Default
// @Description  Tasarımı varsayılan ayarlar
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/{id}/setdefault [get]
func (h *VoucherHandler) SetDocumentTemplateDefault(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/definitions/documenttemplates/%s/setdefault", id))
}

// @Summary      Preview Document Template
// @Description  Tasarımı önizler
// @Tags         Definitions
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/definitions/documenttemplates/{id}/preview [post]
func (h *VoucherHandler) PreviewDocumentTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "POST", fmt.Sprintf("/v1/definitions/documenttemplates/%s/preview", id))
}
