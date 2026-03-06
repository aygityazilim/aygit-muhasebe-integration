package v1

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/internal/service"
	"aygit-muhasebe-integration/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

// Helper function to handle JSON body parsing and service requests
func handleEDespatchRequest(c *fiber.Ctx, serviceFunc func(*models.Company, map[string]interface{}) (map[string]interface{}, error)) error {
	company := c.Locals("company").(*models.Company)
	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil && len(c.Body()) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errors.AppError{
			Code:    "INVALID_REQUEST",
			Message: "Geçersiz istek gövdesi",
		})
	}

	result, err := serviceFunc(company, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{
			Code:    "NES_API_ERROR",
			Message: "NES API çağrısı başarısız oldu",
		})
	}

	return c.JSON(result)
}

func handleEDespatchGetRequest(c *fiber.Ctx, serviceFunc func(*models.Company) (map[string]interface{}, error)) error {
	company := c.Locals("company").(*models.Company)
	result, err := serviceFunc(company)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{
			Code:    "NES_API_ERROR",
			Message: "NES API çağrısı başarısız oldu",
		})
	}

	return c.JSON(result)
}

func handleEDespatchDeleteRequest(c *fiber.Ctx, serviceFunc func(*models.Company, string) (map[string]interface{}, error), paramName string) error {
	company := c.Locals("company").(*models.Company)
	paramValue := c.Params(paramName)
	if paramValue == "" {
		return c.Status(fiber.StatusBadRequest).JSON(errors.AppError{
			Code:    "INVALID_REQUEST",
			Message: "Eksik parametre: " + paramName})
	}

	result, err := serviceFunc(company, paramValue)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{
			Code:    "NES_API_ERROR",
			Message: "NES API çağrısı başarısız oldu",
		})
	}

	return c.JSON(result)
}

// @Summary E-İrsaliye Dışa Aktarım Başlıklarını Listeler
// @Description Kullanılabilir alanları listeler
// @Tags E-Despatch Definitions
// @Accept json
// @Produce json
// @Param documentType path string true "Document Type"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/definitions/fileexporttitles/{documentType}/titlekeys [get]
func getEDespatchFileExportTitles(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	documentType := c.Params("documentType")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchFileExportTitles(company, documentType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary E-İrsaliye Dışa Aktarım Tanımlarını Getirir
// @Description Tanımları getirir
// @Tags E-Despatch Definitions
// @Accept json
// @Produce json
// @Param documentType path string true "Document Type"
// @Param extension path string true "Extension"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/definitions/fileexporttitles/{documentType}/{extension} [get]
func getEDespatchFileExportTitleDefinition(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	documentType := c.Params("documentType")
	extension := c.Params("extension")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchFileExportTitleDefinition(company, documentType, extension)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary E-İrsaliye Dışa Aktarım Tanımlarını Günceller
// @Description Tanımları günceller
// @Tags E-Despatch Definitions
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Tanım Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/definitions/fileexporttitles [put]
func updateEDespatchFileExportTitles(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.UpdateEDespatchFileExportTitles)
}

// @Summary Belge Günceller
// @Description Yüklenen belgeyi günceller
// @Tags E-Despatch Uploads
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param body body map[string]interface{} true "Belge Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/uploads/document/{uuid} [put]
func updateEDespatchDocument(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEDespatchDocument(company, uuid, payload)
	})
}

// @Summary Etiketleri Listeler
// @Description Mevcut etiketleri listeler
// @Tags E-Despatch Tags
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/tags [get]
func getEDespatchTags(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchTags)
}

// @Summary Etiket Ekler
// @Description Yeni bir etiket ekler
// @Tags E-Despatch Tags
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Etiket Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/tags [post]
func createEDespatchTag(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.CreateEDespatchTag)
}

// @Summary Etiketi Getirir
// @Description Sorgulanan etiketi getirir
// @Tags E-Despatch Tags
// @Produce json
// @Param id path string true "Etiket ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/tags/{id} [get]
func getEDespatchTag(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	id := c.Params("id")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchTag(company, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Etiket Günceller
// @Description Etiketi günceller
// @Tags E-Despatch Tags
// @Accept json
// @Produce json
// @Param id path string true "Etiket ID"
// @Param body body map[string]interface{} true "Etiket Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/tags/{id} [put]
func updateEDespatchTag(c *fiber.Ctx) error {
	id := c.Params("id")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEDespatchTag(company, id, payload)
	})
}

// @Summary Etiket Siler
// @Description Etiketi siler
// @Tags E-Despatch Tags
// @Produce json
// @Param id path string true "Etiket ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/tags/{id} [delete]
func deleteEDespatchTag(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchDeleteRequest(c, nesService.DeleteEDespatchTag, "id")
}

// @Summary Gelen Kuralları Listeler
// @Description Kuralları listeler
// @Tags E-Despatch Incoming Notifications
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/incoming/dynamicrules [get]
func getEDespatchIncomingDynamicRules(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchIncomingDynamicRules)
}

// @Summary Gelen Kural Oluşturur
// @Description Kural oluşturur
// @Tags E-Despatch Incoming Notifications
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Kural Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/incoming/dynamicrules [post]
func createEDespatchIncomingDynamicRule(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.CreateEDespatchIncomingDynamicRule)
}

// @Summary Gelen Kuralı Getirir
// @Description Sorgulanan kuralı getirir
// @Tags E-Despatch Incoming Notifications
// @Produce json
// @Param id path string true "Kural ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/incoming/dynamicrules/{id} [get]
func getEDespatchIncomingDynamicRule(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	id := c.Params("id")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchIncomingDynamicRule(company, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Gelen Kuralı Günceller
// @Description Kural günceller
// @Tags E-Despatch Incoming Notifications
// @Accept json
// @Produce json
// @Param id path string true "Kural ID"
// @Param body body map[string]interface{} true "Kural Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/incoming/dynamicrules/{id} [put]
func updateEDespatchIncomingDynamicRule(c *fiber.Ctx) error {
	id := c.Params("id")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEDespatchIncomingDynamicRule(company, id, payload)
	})
}

// @Summary Gelen Kuralı Siler
// @Description Kural siler
// @Tags E-Despatch Incoming Notifications
// @Produce json
// @Param id path string true "Kural ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/incoming/dynamicrules/{id} [delete]
func deleteEDespatchIncomingDynamicRule(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchDeleteRequest(c, nesService.DeleteEDespatchIncomingDynamicRule, "id")
}

// @Summary Gelen İrsaliyeleri Toplu Aktarır
// @Description Toplu aktar
// @Tags E-Despatch Incoming Despatches
// @Accept json
// @Produce json
// @Param fileType path string true "Dosya Tipi"
// @Param body body map[string]interface{} true "Aktarım Parametreleri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/despatches/export/{fileType} [post]
func exportEDespatchIncomingDespatches(c *fiber.Ctx) error {
	fileType := c.Params("fileType")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.ExportEDespatchIncomingDespatches(company, fileType, payload)
	})
}

// @Summary Gelen Raporları Listeler
// @Description Rapor listeler
// @Tags E-Despatch Incoming Reports
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/reportmodule/reports [get]
func getEDespatchIncomingReports(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchIncomingReports)
}

// @Summary Gelen Rapor Oluşturur
// @Description Rapor oluşturur
// @Tags E-Despatch Incoming Reports
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Rapor Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/reportmodule/reports [post]
func createEDespatchIncomingReport(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.CreateEDespatchIncomingReport)
}

// @Summary Gelen Raporu İndirir
// @Description Rapor indirir
// @Tags E-Despatch Incoming Reports
// @Produce json
// @Param id path string true "Rapor ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/reportmodule/reports/{id}/download [get]
func downloadEDespatchIncomingReport(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	id := c.Params("id")
	nesService := service.NewNESService()
	result, err := nesService.DownloadEDespatchIncomingReport(company, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Gelen Şablonları Listeler
// @Description Şablonları listeler
// @Tags E-Despatch Incoming Reports
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/reportmodule/templates [get]
func getEDespatchIncomingTemplates(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchIncomingTemplates)
}

// @Summary Gelen Şablon Oluşturur
// @Description Rapor şablonu oluşturur
// @Tags E-Despatch Incoming Reports
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Şablon Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/reportmodule/templates [post]
func createEDespatchIncomingTemplate(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.CreateEDespatchIncomingTemplate)
}

// @Summary Gelen Şablonu Getirir
// @Description Sorgulanan şablonu getirir
// @Tags E-Despatch Incoming Reports
// @Produce json
// @Param id path string true "Şablon ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/reportmodule/templates/{id} [get]
func getEDespatchIncomingTemplate(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	id := c.Params("id")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchIncomingTemplate(company, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Gelen Şablonu Günceller
// @Description Rapor şablonunu günceller
// @Tags E-Despatch Incoming Reports
// @Accept json
// @Produce json
// @Param id path string true "Şablon ID"
// @Param body body map[string]interface{} true "Şablon Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/reportmodule/templates/{id} [put]
func updateEDespatchIncomingTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEDespatchIncomingTemplate(company, id, payload)
	})
}

// @Summary Gelen Şablonu Siler
// @Description Rapor Şablonunu siler
// @Tags E-Despatch Incoming Reports
// @Produce json
// @Param id path string true "Şablon ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/reportmodule/templates/{id} [delete]
func deleteEDespatchIncomingTemplate(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchDeleteRequest(c, nesService.DeleteEDespatchIncomingTemplate, "id")
}

// @Summary Gelen Kolonları Listeler
// @Description Kolonları listeler
// @Tags E-Despatch Incoming Reports
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/reportmodule/columns [get]
func getEDespatchIncomingColumns(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchIncomingColumns)
}

// @Summary Gelen İrsaliyelere Etiket Ekler/Çıkarır
// @Description Etiket ekler/çıkarır
// @Tags E-Despatch Incoming Despatches
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Etiket Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/despatches/tags [put]
func updateEDespatchIncomingTags(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.UpdateEDespatchIncomingTags)
}

// @Summary Gelen İrsaliyede Firma Olarak Kaydeder
// @Description Firma olarak kaydet
// @Tags E-Despatch Incoming Despatches
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param body body map[string]interface{} true "Firma Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/despatches/{uuid}/savecompanyindocument [post]
func saveCompanyInIncomingDocument(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.SaveCompanyInIncomingDocument(company, uuid, payload)
	})
}

// @Summary Gelen İrsaliyelere Yeni Durum Atar
// @Description Yeni durum atar
// @Tags E-Despatch Incoming Despatches
// @Accept json
// @Produce json
// @Param operation path string true "Operasyon"
// @Param body body map[string]interface{} true "Durum Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/despatches/bulk/{operation} [put]
func bulkOperationIncomingDespatches(c *fiber.Ctx) error {
	operation := c.Params("operation")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.BulkOperationIncomingDespatches(company, operation, payload)
	})
}

// @Summary Gelen İrsaliyeye Kullanıcı Notu Ekler
// @Description Kullanıcı notu ekler
// @Tags E-Despatch Incoming Despatches
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param body body map[string]interface{} true "Not Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/despatches/{uuid}/usernotes [post]
func addUserNoteToIncomingDespatch(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.AddUserNoteToIncomingDespatch(company, uuid, payload)
	})
}

// @Summary Gelen İrsaliyedeki Kullanıcı Notunu Günceller
// @Description Kullanıcı notunu günceller
// @Tags E-Despatch Incoming Despatches
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param id path string true "Not ID"
// @Param body body map[string]interface{} true "Not Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/despatches/{uuid}/usernotes/{id} [put]
func updateUserNoteInIncomingDespatch(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	id := c.Params("id")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateUserNoteInIncomingDespatch(company, uuid, id, payload)
	})
}

// @Summary Gelen İrsaliyedeki Kullanıcı Notunu Siler
// @Description Kullanıcı notunu siler
// @Tags E-Despatch Incoming Despatches
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param id path string true "Not ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/despatches/{uuid}/usernotes/{id} [delete]
func deleteUserNoteFromIncomingDespatch(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	id := c.Params("id")
	nesService := service.NewNESService()
	result, err := nesService.DeleteUserNoteFromIncomingDespatch(company, uuid, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Giden İrsaliyeye Verilen Yanıtları Listeler
// @Description Giden irsaliyeye verilen yanıtları (ReceiptAdvice) listeler
// @Tags E-Despatch Incoming Receipt Advices
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/receiptadvices [get]
func getEDespatchIncomingReceiptAdvices(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchIncomingReceiptAdvices)
}

// @Summary Giden İrsaliyeye Verilen Yanıtın Detayını Getirir
// @Description Belge detaylarını getirir
// @Tags E-Despatch Incoming Receipt Advices
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/receiptadvices/{uuid} [get]
func getEDespatchIncomingReceiptAdvice(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchIncomingReceiptAdvice(company, uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Giden İrsaliyeye Verilen Yanıtı HTML Olarak Getirir
// @Description Belgeyi görüntüler
// @Tags E-Despatch Incoming Receipt Advices
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/receiptadvices/{uuid}/html [get]
func getEDespatchIncomingReceiptAdviceHTML(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchIncomingReceiptAdviceHTML(company, uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Giden İrsaliyeye Verilen Yanıt XML İndir
// @Description Belgeye ait yanıt XML İndir
// @Tags E-Despatch Incoming Receipt Advices
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/receiptadvices/{uuid}/xml [get]
func getEDespatchIncomingReceiptAdviceXML(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchIncomingReceiptAdviceXML(company, uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Giden İrsaliyeye Verilen Yanıt PDF İndir
// @Description PDF İndir
// @Tags E-Despatch Incoming Receipt Advices
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/receiptadvices/{uuid}/pdf [get]
func getEDespatchIncomingReceiptAdvicePDF(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchIncomingReceiptAdvicePDF(company, uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Gelen e-İrsaliyeye Yanıt Verir
// @Description Gelen e-İrsaliyeye yanıt verilir (ReceiptAdvice)
// @Tags E-Despatch Incoming Receipt Advices
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param body body map[string]interface{} true "Yanıt Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/despatches/{uuid}/receiptadvice [post]
func sendReceiptAdviceForIncomingDespatch(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.SendReceiptAdviceForIncomingDespatch(company, uuid, payload)
	})
}

// @Summary Gelen Belgeyi E-posta Olarak İletir
// @Description Belgeyi mail olarak iletir
// @Tags E-Despatch Incoming Despatches
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "E-posta Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/incoming/despatches/email/send [post]
func sendEmailForIncomingDespatch(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.SendEmailForIncomingDespatch)
}

// @Summary Giden Kuralları Listeler
// @Description Kuralları listeler
// @Tags E-Despatch Outgoing Notifications
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/outgoing/dynamicrules [get]
func getEDespatchOutgoingDynamicRules(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchOutgoingDynamicRules)
}

// @Summary Giden Kural Oluşturur
// @Description Kural oluşturur
// @Tags E-Despatch Outgoing Notifications
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Kural Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/outgoing/dynamicrules [post]
func createEDespatchOutgoingDynamicRule(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.CreateEDespatchOutgoingDynamicRule)
}

// @Summary Giden Kuralı Getirir
// @Description Sorgulanan kuralı getirir
// @Tags E-Despatch Outgoing Notifications
// @Produce json
// @Param id path string true "Kural ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/outgoing/dynamicrules/{id} [get]
func getEDespatchOutgoingDynamicRule(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	id := c.Params("id")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchOutgoingDynamicRule(company, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Giden Kuralı Günceller
// @Description Kural günceller
// @Tags E-Despatch Outgoing Notifications
// @Accept json
// @Produce json
// @Param id path string true "Kural ID"
// @Param body body map[string]interface{} true "Kural Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/outgoing/dynamicrules/{id} [put]
func updateEDespatchOutgoingDynamicRule(c *fiber.Ctx) error {
	id := c.Params("id")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEDespatchOutgoingDynamicRule(company, id, payload)
	})
}

// @Summary Giden Kuralı Siler
// @Description Kural siler
// @Tags E-Despatch Outgoing Notifications
// @Produce json
// @Param id path string true "Kural ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/notifications/outgoing/dynamicrules/{id} [delete]
func deleteEDespatchOutgoingDynamicRule(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchDeleteRequest(c, nesService.DeleteEDespatchOutgoingDynamicRule, "id")
}

// @Summary Giden İrsaliyeleri Toplu Aktarır
// @Description Toplu aktar
// @Tags E-Despatch Outgoing Despatches
// @Accept json
// @Produce json
// @Param fileType path string true "Dosya Tipi"
// @Param body body map[string]interface{} true "Aktarım Parametreleri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/export/{fileType} [post]
func exportEDespatchOutgoingDespatches(c *fiber.Ctx) error {
	fileType := c.Params("fileType")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.ExportEDespatchOutgoingDespatches(company, fileType, payload)
	})
}

// @Summary Giden Raporları Listeler
// @Description Rapor listeler
// @Tags E-Despatch Outgoing Reports
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/reportmodule/reports [get]
func getEDespatchOutgoingReports(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchOutgoingReports)
}

// @Summary Giden Rapor Oluşturur
// @Description Rapor oluşturur
// @Tags E-Despatch Outgoing Reports
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Rapor Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/reportmodule/reports [post]
func createEDespatchOutgoingReport(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.CreateEDespatchOutgoingReport)
}

// @Summary Giden Raporu İndirir
// @Description Rapor indirir
// @Tags E-Despatch Outgoing Reports
// @Produce json
// @Param id path string true "Rapor ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/reportmodule/reports/{id}/download [get]
func downloadEDespatchOutgoingReport(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	id := c.Params("id")
	nesService := service.NewNESService()
	result, err := nesService.DownloadEDespatchOutgoingReport(company, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Giden Şablonları Listeler
// @Description Şablonları listeler
// @Tags E-Despatch Outgoing Reports
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/reportmodule/templates [get]
func getEDespatchOutgoingTemplates(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchOutgoingTemplates)
}

// @Summary Giden Şablon Oluşturur
// @Description Rapor şablonu oluşturur
// @Tags E-Despatch Outgoing Reports
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Şablon Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/reportmodule/templates [post]
func createEDespatchOutgoingTemplate(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.CreateEDespatchOutgoingTemplate)
}

// @Summary Giden Şablonu Getirir
// @Description Sorgulanan şablonu getirir
// @Tags E-Despatch Outgoing Reports
// @Produce json
// @Param id path string true "Şablon ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/reportmodule/templates/{id} [get]
func getEDespatchOutgoingTemplate(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	id := c.Params("id")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchOutgoingTemplate(company, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Giden Şablonu Günceller
// @Description Rapor şablonunu günceller
// @Tags E-Despatch Outgoing Reports
// @Accept json
// @Produce json
// @Param id path string true "Şablon ID"
// @Param body body map[string]interface{} true "Şablon Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/reportmodule/templates/{id} [put]
func updateEDespatchOutgoingTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEDespatchOutgoingTemplate(company, id, payload)
	})
}

// @Summary Giden Şablonu Siler
// @Description Rapor Şablonunu siler
// @Tags E-Despatch Outgoing Reports
// @Produce json
// @Param id path string true "Şablon ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/reportmodule/templates/{id} [delete]
func deleteEDespatchOutgoingTemplate(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchDeleteRequest(c, nesService.DeleteEDespatchOutgoingTemplate, "id")
}

// @Summary Giden Kolonları Listeler
// @Description Kolonları listeler
// @Tags E-Despatch Outgoing Reports
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/reportmodule/columns [get]
func getEDespatchOutgoingColumns(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchOutgoingColumns)
}

// @Summary Giden İrsaliyelere Etiket Ekler/Çıkarır
// @Description Etiket ekler/çıkarır
// @Tags E-Despatch Outgoing Despatches
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Etiket Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/tags [put]
func updateEDespatchOutgoingTags(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.UpdateEDespatchOutgoingTags)
}

// @Summary Giden Taslak Belgelerin Alıcı Etiketini Günceller
// @Description Taslak belgelerin alıcı etiketi bu uç ile güncellenebilir
// @Tags E-Despatch Outgoing Despatches
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param body body map[string]interface{} true "Etiket Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/{uuid}/receiveralias [put]
func updateEDespatchOutgoingReceiverAlias(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEDespatchOutgoingReceiverAlias(company, uuid, payload)
	})
}

// @Summary Giden İrsaliyede Firma Olarak Kaydeder
// @Description Firma olarak kaydet
// @Tags E-Despatch Outgoing Despatches
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param body body map[string]interface{} true "Firma Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/{uuid}/savecompanyindocument [post]
func saveCompanyInOutgoingDocument(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.SaveCompanyInOutgoingDocument(company, uuid, payload)
	})
}

// @Summary Giden İrsaliyelere Yeni Durum Atar
// @Description Yeni durum atar
// @Tags E-Despatch Outgoing Despatches
// @Accept json
// @Produce json
// @Param operation path string true "Operasyon"
// @Param body body map[string]interface{} true "Durum Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/bulk/{operation} [put]
func bulkOperationOutgoingDespatches(c *fiber.Ctx) error {
	operation := c.Params("operation")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.BulkOperationOutgoingDespatches(company, operation, payload)
	})
}

// @Summary Hatalı Belgeyi Yeniden Gönderir
// @Description Hata almış bir belgeyi aynen yeniden gönderir
// @Tags E-Despatch Uploads
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param body body map[string]interface{} true "Gönderim Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/uploads/resend/{uuid} [post]
func resendErrorDocument(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.ResendErrorDocument(company, uuid, payload)
	})
}

// @Summary Giden İrsaliyeye Kullanıcı Notu Ekler
// @Description Kullanıcı notu ekler
// @Tags E-Despatch Outgoing Despatches
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param body body map[string]interface{} true "Not Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/{uuid}/usernotes [post]
func addUserNoteToOutgoingDespatch(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.AddUserNoteToOutgoingDespatch(company, uuid, payload)
	})
}

// @Summary Giden İrsaliyedeki Kullanıcı Notunu Günceller
// @Description Kullanıcı notunu günceller
// @Tags E-Despatch Outgoing Despatches
// @Accept json
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param id path string true "Not ID"
// @Param body body map[string]interface{} true "Not Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/{uuid}/usernotes/{id} [put]
func updateUserNoteInOutgoingDespatch(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	id := c.Params("id")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateUserNoteInOutgoingDespatch(company, uuid, id, payload)
	})
}

// @Summary Giden İrsaliyedeki Kullanıcı Notunu Siler
// @Description Kullanıcı notunu siler
// @Tags E-Despatch Outgoing Despatches
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Param id path string true "Not ID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/{uuid}/usernotes/{id} [delete]
func deleteUserNoteFromOutgoingDespatch(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	id := c.Params("id")
	nesService := service.NewNESService()
	result, err := nesService.DeleteUserNoteFromOutgoingDespatch(company, uuid, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Gelen İrsaliyeye Verilen Yanıtları Listeler
// @Description Gelen irsaliyeye verilen yanıtları (ReceiptAdvice) listeler
// @Tags E-Despatch Outgoing Receipt Advices
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/receiptadvices [get]
func getEDespatchOutgoingReceiptAdvices(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchOutgoingReceiptAdvices)
}

// @Summary Gelen İrsaliyeye Verilen Yanıtın Detayını Getirir
// @Description Belge detaylarını getirir
// @Tags E-Despatch Outgoing Receipt Advices
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/receiptadvices/{uuid} [get]
func getEDespatchOutgoingReceiptAdvice(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchOutgoingReceiptAdvice(company, uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Gelen İrsaliyeye Verilen Yanıtı HTML Olarak Getirir
// @Description Belgeyi görüntüler
// @Tags E-Despatch Outgoing Receipt Advices
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/receiptadvices/{uuid}/html [get]
func getEDespatchOutgoingReceiptAdviceHTML(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchOutgoingReceiptAdviceHTML(company, uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Gelen İrsaliyeye Verilen Yanıt XML İndir
// @Description Belgeye ait yanıt XML İndir
// @Tags E-Despatch Outgoing Receipt Advices
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/receiptadvices/{uuid}/xml [get]
func getEDespatchOutgoingReceiptAdviceXML(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchOutgoingReceiptAdviceXML(company, uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Gelen İrsaliyeye Verilen Yanıt PDF İndir
// @Description PDF İndir
// @Tags E-Despatch Outgoing Receipt Advices
// @Produce json
// @Param uuid path string true "Belge UUID"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/receiptadvices/{uuid}/pdf [get]
func getEDespatchOutgoingReceiptAdvicePDF(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	uuid := c.Params("uuid")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchOutgoingReceiptAdvicePDF(company, uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Giden Taslak Belgeleri Siler
// @Description Taslak belgeleri silmek için bu uç kullanılablir
// @Tags E-Despatch Outgoing Despatches
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/drafts [delete]
func deleteEDespatchOutgoingDrafts(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.DeleteEDespatchOutgoingDrafts) // Although DELETE, request structure is similar to GET without payload
}

// @Summary Giden Belgeyi E-posta Olarak İletir
// @Description Belgeyi mail olarak iletir
// @Tags E-Despatch Outgoing Despatches
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "E-posta Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/outgoing/despatches/email/send [post]
func sendEmailForOutgoingDespatch(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.SendEmailForOutgoingDespatch)
}

// @Summary Mail Ayarlarını Getirir
// @Description Mail ayarlarını getirir
// @Tags E-Despatch Mailing Settings
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/definitions/mailing/email/settings [get]
func getEDespatchEmailSettings(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchEmailSettings)
}

// @Summary Mail Ayarlarını Günceller
// @Description Mail ayarlarını günceller
// @Tags E-Despatch Mailing Settings
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Ayarlar Verisi"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/definitions/mailing/email/settings [put]
func updateEDespatchEmailSettings(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.UpdateEDespatchEmailSettings)
}

// @Summary SMS Ayarlarını Getirir
// @Description Sms ayarlarını getirir
// @Tags E-Despatch Mailing Settings
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/definitions/mailing/sms/settings [get]
func getEDespatchSmsSettings(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchGetRequest(c, nesService.GetEDespatchSmsSettings)
}

// @Summary SMS Ayarlarını Günceller
// @Description Sms ayarlarını günceller
// @Tags E-Despatch Mailing Settings
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Ayarlar Verisi"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/definitions/mailing/sms/settings [put]
func updateEDespatchSmsSettings(c *fiber.Ctx) error {
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, nesService.UpdateEDespatchSmsSettings)
}

// @Summary Mükellef Listesini İndirir
// @Description Mükellef listesini indirir
// @Tags E-Despatch Users
// @Produce json
// @Param aliasType path string true "Alias Type"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/users/zip/{aliasType} [get]
func getEDespatchUsersZip(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	aliasType := c.Params("aliasType")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchUsersZip(company, aliasType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Kimlik No İle Sorgular
// @Description Kimlik No ile sorgular
// @Tags E-Despatch Users
// @Produce json
// @Param identifier path string true "Kimlik No"
// @Param aliasType path string true "Alias Type"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/users/{identifier}/{aliasType} [get]
func getEDespatchUserByIdentifier(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	identifier := c.Params("identifier")
	aliasType := c.Params("aliasType")
	nesService := service.NewNESService()
	result, err := nesService.GetEDespatchUserByIdentifier(company, identifier, aliasType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}

// @Summary Kimlik No İle Sorgular (POST)
// @Description Kimlik No ile sorgular
// @Tags E-Despatch Users
// @Accept json
// @Produce json
// @Param aliasType path string true "Alias Type"
// @Param body body map[string]interface{} true "Sorgu Verileri"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/users/{aliasType} [post]
func getEDespatchUserByIdentifierPost(c *fiber.Ctx) error {
	aliasType := c.Params("aliasType")
	nesService := service.NewNESService()
	return handleEDespatchRequest(c, func(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
		return nesService.GetEDespatchUserByIdentifierPost(company, aliasType, payload)
	})
}

// @Summary Ünvan İle Sorgular
// @Description Ünvan ile sorgular
// @Tags E-Despatch Users
// @Produce json
// @Param query path string true "Arama Sorgusu"
// @Param aliasType path string true "Alias Type"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/edespatch/users/search/{query}/{aliasType} [get]
func searchEDespatchUserByTitle(c *fiber.Ctx) error {
	company := c.Locals("company").(*models.Company)
	query := c.Params("query")
	aliasType := c.Params("aliasType")
	nesService := service.NewNESService()
	result, err := nesService.SearchEDespatchUserByTitle(company, query, aliasType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errors.AppError{Code: "NES_API_ERROR", Message: "NES API çağrısı başarısız oldu"})
	}
	return c.JSON(result)
}
