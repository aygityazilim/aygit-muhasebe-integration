package v1

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/db"
	"aygit-muhasebe-integration/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

// Helper for generic e-archive requests
func handleGenericEArchiveRequest(c *fiber.Ctx, serviceFunc func(company *models.Company, params map[string]string) (map[string]interface{}, error)) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	params := c.Queries()
	result, err := serviceFunc(&company, params)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func handleGenericEArchivePostPutRequest(c *fiber.Ctx, serviceFunc func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error)) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return errors.NewError(fiber.StatusBadRequest, errors.ErrCodeInvalidRequest, "Geçersiz istek gövdesi")
	}

	result, err := serviceFunc(&company, payload)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func handleGenericEArchiveDeleteRequest(c *fiber.Ctx, id string, serviceFunc func(company *models.Company, id string) (map[string]interface{}, error)) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	result, err := serviceFunc(&company, id)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

// getEArchiveDynamicRules godoc
// @Summary Kuralları listeler
// @Description Kuralları listeler
// @Tags E-Archive Notifications
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/notifications/dynamicrules [get]
func getEArchiveDynamicRules(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveDynamicRules)
}

// createEArchiveDynamicRule godoc
// @Summary Kural oluşturur
// @Description Kural oluşturur
// @Tags E-Archive Notifications
// @Security BearerAuth
// @Param request body map[string]interface{} true "Rule details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/notifications/dynamicrules [post]
func createEArchiveDynamicRule(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.CreateEArchiveDynamicRule)
}

// getEArchiveDynamicRule godoc
// @Summary Sorgulanan kuralı getirir
// @Description Sorgulanan kuralı getirir
// @Tags E-Archive Notifications
// @Security BearerAuth
// @Param id path string true "Rule ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/notifications/dynamicrules/{id} [get]
func getEArchiveDynamicRule(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveDynamicRule(company, id)
	})
}

// updateEArchiveDynamicRule godoc
// @Summary Kural günceller
// @Description Kural günceller
// @Tags E-Archive Notifications
// @Security BearerAuth
// @Param id path string true "Rule ID"
// @Param request body map[string]interface{} true "Rule details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/notifications/dynamicrules/{id} [put]
func updateEArchiveDynamicRule(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEArchiveDynamicRule(company, id, body)
	})
}

// deleteEArchiveDynamicRule godoc
// @Summary Kural siler
// @Description Kural siler
// @Tags E-Archive Notifications
// @Security BearerAuth
// @Param id path string true "Rule ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/notifications/dynamicrules/{id} [delete]
func deleteEArchiveDynamicRule(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, nesService.DeleteEArchiveDynamicRule)
}

// exportEArchiveInvoices godoc
// @Summary Toplu aktar
// @Description Toplu aktar
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Param fileType path string true "File Type"
// @Param request body map[string]interface{} true "Export details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/export/{fileType} [post]
func exportEArchiveInvoices(c *fiber.Ctx) error {
	fileType := c.Params("fileType")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.ExportEArchiveInvoices(company, fileType, body)
	})
}

// getEArchiveFileExportTitles godoc
// @Summary Kullanılabilir alanları listeler
// @Description Kullanılabilir alanları listeler
// @Tags E-Archive Definitions
// @Security BearerAuth
// @Param documentType path string true "Document Type"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/fileexporttitles/{documentType}/titlekeys [get]
func getEArchiveFileExportTitles(c *fiber.Ctx) error {
	documentType := c.Params("documentType")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveFileExportTitles(company, documentType)
	})
}

// getEArchiveFileExportTitlesExtension godoc
// @Summary Tanımları getirir
// @Description Tanımları getirir
// @Tags E-Archive Definitions
// @Security BearerAuth
// @Param documentType path string true "Document Type"
// @Param extension path string true "Extension"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/fileexporttitles/{documentType}/{extension} [get]
func getEArchiveFileExportTitlesExtension(c *fiber.Ctx) error {
	documentType := c.Params("documentType")
	extension := c.Params("extension")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveFileExportTitlesExtension(company, documentType, extension)
	})
}

// updateEArchiveFileExportTitles godoc
// @Summary Tanımları günceller
// @Description Tanımları günceller
// @Tags E-Archive Definitions
// @Security BearerAuth
// @Param request body map[string]interface{} true "Titles details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/fileexporttitles [put]
func updateEArchiveFileExportTitles(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.UpdateEArchiveFileExportTitles)
}

// previewEArchiveDocument godoc
// @Summary Belge önizleme
// @Description Belge önizleme
// @Tags E-Archive Uploads
// @Security BearerAuth
// @Param request body map[string]interface{} true "Document details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/uploads/document/preview [post]
func previewEArchiveDocument(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.PreviewEArchiveDocument)
}

// updateEArchiveDocument godoc
// @Summary Belge günceller
// @Description Belge günceller
// @Tags E-Archive Uploads
// @Security BearerAuth
// @Param uuid path string true "Document UUID"
// @Param request body map[string]interface{} true "Document details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/uploads/document/{uuid} [put]
func updateEArchiveDocument(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEArchiveDocument(company, uuid, body)
	})
}

// createEArchiveDraft godoc
// @Summary Belge yükler
// @Description Belge yükler
// @Tags E-Archive Uploads
// @Security BearerAuth
// @Param id path string true "Draft ID"
// @Param request body map[string]interface{} true "Draft details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/uploads/draft/create/{id} [post]
func createEArchiveDraft(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.CreateEArchiveDraft(company, id, body)
	})
}

// previewEArchiveMarketplaceOrder godoc
// @Summary Belirtilen pazaryerindeki siparişin faturasını önizler
// @Description Belirtilen pazaryerindeki siparişin faturasını önizler
// @Tags E-Archive Uploads
// @Security BearerAuth
// @Param id path string true "Marketplace ID"
// @Param orderId path string true "Order ID"
// @Param request body map[string]interface{} true "Preview details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/uploads/marketplaces/{id}/orders/{orderId}/preview [post]
func previewEArchiveMarketplaceOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	orderId := c.Params("orderId")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.PreviewEArchiveMarketplaceOrder(company, id, orderId, body)
	})
}

// createEArchiveMarketplaceInvoice godoc
// @Summary Belirtilen pazaryerindeki siparişin faturasını önizler
// @Description Belirtilen pazaryerindeki siparişin faturasını önizler
// @Tags E-Archive Uploads
// @Security BearerAuth
// @Param id path string true "Marketplace ID"
// @Param request body map[string]interface{} true "Invoice details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/uploads/marketplaces/{id}/orders/createinvoice [post]
func createEArchiveMarketplaceInvoice(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.CreateEArchiveMarketplaceInvoice(company, id, body)
	})
}

// getEArchiveExInvoices godoc
// @Summary Belgeleri listeler
// @Description Belgeleri listeler
// @Tags E-Archive ExInvoices
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/exinvoices [get]
func getEArchiveExInvoices(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveExInvoices)
}

// uploadEArchiveExInvoice godoc
// @Summary Eski belge yükler
// @Description Eski belge yükler
// @Tags E-Archive ExInvoices
// @Security BearerAuth
// @Param request body map[string]interface{} true "Invoice details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/exinvoices [post]
func uploadEArchiveExInvoice(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.UploadEArchiveExInvoice)
}

// getEArchiveExInvoicesQueue godoc
// @Summary Yükleme kuyruğunu listeler
// @Description Yükleme kuyruğunu listeler
// @Tags E-Archive ExInvoices
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/exinvoices/queue [get]
func getEArchiveExInvoicesQueue(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveExInvoicesQueue)
}

// getEArchiveExInvoicesQueueResult godoc
// @Summary Yükleme sonucunu indir
// @Description Yükleme sonucunu indir
// @Tags E-Archive ExInvoices
// @Security BearerAuth
// @Param id path string true "Queue ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/exinvoices/queue/{id} [get]
func getEArchiveExInvoicesQueueResult(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveExInvoicesQueueResult(company, id)
	})
}

// downloadEArchiveExInvoiceFile godoc
// @Summary Dosya İndir
// @Description Belirtilen faturanın XML, PDF veya HTML dosyasını indirir.
// @Tags E-Archive ExInvoices
// @Security BearerAuth
// @Param uuid path string true "Fatura UUID"
// @Param fileType path string true "Dosya Tipi (xml/pdf/html)"
// @Success 200 {binary} binary
// @Router /v1/earchive/exinvoices/{uuid}/{fileType} [get]
func downloadEArchiveExInvoiceFile(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	uuid := c.Params("uuid")
	fileType := c.Params("fileType")

	content, contentType, err := nesService.DownloadEArchiveExInvoiceFile(&company, uuid, fileType)
	if err != nil {
		return err
	}

	c.Set("Content-Type", contentType)
	return c.Send(content)
}

// exportEArchiveExInvoices godoc
// @Summary Dışarı Aktar
// @Description Dışarı Aktar
// @Tags E-Archive ExInvoices
// @Security BearerAuth
// @Param fileType path string true "File Type"
// @Param request body map[string]interface{} true "Export details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/exinvoices/export/{fileType} [post]
func exportEArchiveExInvoices(c *fiber.Ctx) error {
	fileType := c.Params("fileType")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.ExportEArchiveExInvoices(company, fileType, body)
	})
}

// getEArchiveTags godoc
// @Summary Etiketleri listeler
// @Description Etiketleri listeler
// @Tags E-Archive Tags
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/tags [get]
func getEArchiveTags(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveTags)
}

// createEArchiveTag godoc
// @Summary Etiket ekler
// @Description Etiket ekler
// @Tags E-Archive Tags
// @Security BearerAuth
// @Param request body map[string]interface{} true "Tag details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/tags [post]
func createEArchiveTag(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.CreateEArchiveTag)
}

// getEArchiveTag godoc
// @Summary Sorgulanan etiketi getirir
// @Description Sorgulanan etiketi getirir
// @Tags E-Archive Tags
// @Security BearerAuth
// @Param id path string true "Tag ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/tags/{id} [get]
func getEArchiveTag(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveTag(company, id)
	})
}

// updateEArchiveTag godoc
// @Summary Etiket günceller
// @Description Etiket günceller
// @Tags E-Archive Tags
// @Security BearerAuth
// @Param id path string true "Tag ID"
// @Param request body map[string]interface{} true "Tag details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/tags/{id} [put]
func updateEArchiveTag(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEArchiveTag(company, id, body)
	})
}

// deleteEArchiveTag godoc
// @Summary Etiket siler
// @Description Etiket siler
// @Tags E-Archive Tags
// @Security BearerAuth
// @Param id path string true "Tag ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/tags/{id} [delete]
func deleteEArchiveTag(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, nesService.DeleteEArchiveTag)
}

// getEArchiveReports godoc
// @Summary Rapor listeler
// @Description Rapor listeler
// @Tags E-Archive Reports
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/outgoing/reportmodule/reports [get]
func getEArchiveReports(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveReports)
}

// createEArchiveReport godoc
// @Summary Rapor oluşturur
// @Description Rapor oluşturur
// @Tags E-Archive Reports
// @Security BearerAuth
// @Param request body map[string]interface{} true "Report details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/outgoing/reportmodule/reports [post]
func createEArchiveReport(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.CreateEArchiveReport)
}

// downloadEArchiveReport godoc
// @Summary Rapor indirir
// @Description Rapor indirir
// @Tags E-Archive Reports
// @Security BearerAuth
// @Param id path string true "Report ID"
// @Success 200 {binary} binary
// @Router /v1/earchive/outgoing/reportmodule/reports/{id}/download [get]
func downloadEArchiveReport(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")

	content, contentType, err := nesService.DownloadEArchiveReport(&company, id)
	if err != nil {
		return err
	}

	c.Set("Content-Type", contentType)
	return c.Send(content)
}

// getEArchiveReportTemplates godoc
// @Summary Şablonları listeler
// @Description Şablonları listeler
// @Tags E-Archive Reports
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/outgoing/reportmodule/templates [get]
func getEArchiveReportTemplates(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveReportTemplates)
}

// createEArchiveReportTemplate godoc
// @Summary Rapor şablonu oluşturur
// @Description Rapor şablonu oluşturur
// @Tags E-Archive Reports
// @Security BearerAuth
// @Param request body map[string]interface{} true "Template details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/outgoing/reportmodule/templates [post]
func createEArchiveReportTemplate(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.CreateEArchiveReportTemplate)
}

// getEArchiveReportTemplate godoc
// @Summary Sorgulanan şablonu getirir
// @Description Sorgulanan şablonu getirir
// @Tags E-Archive Reports
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/outgoing/reportmodule/templates/{id} [get]
func getEArchiveReportTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveReportTemplate(company, id)
	})
}

// updateEArchiveReportTemplate godoc
// @Summary Rapor şablonunu günceller
// @Description Rapor şablonunu günceller
// @Tags E-Archive Reports
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Param request body map[string]interface{} true "Template details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/outgoing/reportmodule/templates/{id} [put]
func updateEArchiveReportTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEArchiveReportTemplate(company, id, body)
	})
}

// deleteEArchiveReportTemplate godoc
// @Summary Rapor Şablonunu siler
// @Description Rapor Şablonunu siler
// @Tags E-Archive Reports
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/outgoing/reportmodule/templates/{id} [delete]
func deleteEArchiveReportTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, nesService.DeleteEArchiveReportTemplate)
}

// getEArchiveReportColumns godoc
// @Summary Kolonları listeler
// @Description Kolonları listeler
// @Tags E-Archive Reports
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/outgoing/reportmodule/columns [get]
func getEArchiveReportColumns(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveReportColumns)
}

// updateEArchiveInvoiceTags godoc
// @Summary Etiket ekler/çıkarır
// @Description Etiket ekler/çıkarır
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Param request body map[string]interface{} true "Tag details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/tags [put]
func updateEArchiveInvoiceTags(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.UpdateEArchiveInvoiceTags)
}

// saveCompanyInEArchiveDocument godoc
// @Summary Firma olarak kaydet
// @Description Firma olarak kaydet
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Param uuid path string true "Document UUID"
// @Param request body map[string]interface{} true "Company details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/{uuid}/savecompanyindocument [post]
func saveCompanyInEArchiveDocument(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.SaveCompanyInEArchiveDocument(company, uuid, body)
	})
}

// bulkEArchiveInvoiceOperation godoc
// @Summary Yeni durum atar
// @Description Yeni durum atar
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Param operation path string true "Operation"
// @Param request body map[string]interface{} true "Operation details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/bulk/{operation} [put]
func bulkEArchiveInvoiceOperation(c *fiber.Ctx) error {
	operation := c.Params("operation")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.BulkEArchiveInvoiceOperation(company, operation, body)
	})
}

// deleteEArchiveDraftInvoices godoc
// @Summary Taslak belgeleri silmek için bu uç kullanılablir
// @Description Taslak belgeleri silmek için bu uç kullanılablir
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/drafts [delete]
func deleteEArchiveDraftInvoices(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.DeleteEArchiveDraftInvoices)
}

// sendEArchiveInvoiceEmail godoc
// @Summary Belgeyi mail olarak iletir
// @Description Belgeyi mail olarak iletir
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Param request body map[string]interface{} true "Email details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/email/send [post]
func sendEArchiveInvoiceEmail(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.SendEArchiveInvoiceEmail)
}

// getEArchiveInvoiceUserNotes godoc
// @Summary Kullanıcı notu listeler
// @Description Kullanıcı notu listeler
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Param uuid path string true "Invoice UUID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/{uuid}/usernotes [get]
func getEArchiveInvoiceUserNotes(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveInvoiceUserNotes(company, uuid)
	})
}

// createEArchiveInvoiceUserNote godoc
// @Summary Kullanıcı notu ekler
// @Description Kullanıcı notu ekler
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Param uuid path string true "Invoice UUID"
// @Param request body map[string]interface{} true "Note details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/{uuid}/usernotes [post]
func createEArchiveInvoiceUserNote(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.CreateEArchiveInvoiceUserNote(company, uuid, body)
	})
}

// updateEArchiveInvoiceUserNote godoc
// @Summary Kullanıcı notunu günceller
// @Description Kullanıcı notunu günceller
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Param uuid path string true "Invoice UUID"
// @Param id path string true "Note ID"
// @Param request body map[string]interface{} true "Note details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/{uuid}/usernotes/{id} [put]
func updateEArchiveInvoiceUserNote(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEArchiveInvoiceUserNote(company, uuid, id, body)
	})
}

// deleteEArchiveInvoiceUserNote godoc
// @Summary Kullanıcı notunu siler
// @Description Kullanıcı notunu siler
// @Tags E-Archive Invoices
// @Security BearerAuth
// @Param uuid path string true "Invoice UUID"
// @Param id path string true "Note ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/{uuid}/usernotes/{id} [delete]
func deleteEArchiveInvoiceUserNote(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, func(company *models.Company, id string) (map[string]interface{}, error) {
		return nesService.DeleteEArchiveInvoiceUserNote(company, uuid, id)
	})
}

// getEArchiveEmailSettings godoc
// @Summary Mail ayarlarını getirir
// @Description Mail ayarlarını getirir
// @Tags E-Archive Settings
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/mailing/email/settings [get]
func getEArchiveEmailSettings(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveEmailSettings)
}

// updateEArchiveEmailSettings godoc
// @Summary Mail ayarlarını günceller
// @Description Mail ayarlarını günceller
// @Tags E-Archive Settings
// @Security BearerAuth
// @Param request body map[string]interface{} true "Settings details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/mailing/email/settings [put]
func updateEArchiveEmailSettings(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.UpdateEArchiveEmailSettings)
}

// getEArchiveSmsSettings godoc
// @Summary Sms ayarlarını getirir
// @Description Sms ayarlarını getirir
// @Tags E-Archive Settings
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/mailing/sms/settings [get]
func getEArchiveSmsSettings(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveSmsSettings)
}

// updateEArchiveSmsSettings godoc
// @Summary Sms ayarlarını günceller
// @Description Sms ayarlarını günceller
// @Tags E-Archive Settings
// @Security BearerAuth
// @Param request body map[string]interface{} true "Settings details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/mailing/sms/settings [put]
func updateEArchiveSmsSettings(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.UpdateEArchiveSmsSettings)
}

// getEArchiveCustomizationSettings godoc
// @Summary Tasarım ayarları dönülür
// @Description Tasarım ayarları dönülür
// @Tags E-Archive Templates
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings [get]
func getEArchiveCustomizationSettings(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveCustomizationSettings)
}

// createEArchiveCustomizationSetting godoc
// @Summary e-Belge özelleştirilebilir tasarım eklemek için kullanılır.
// @Description e-Belge özelleştirilebilir tasarım eklemek için kullanılır.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param request body map[string]interface{} true "Setting details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings [post]
func createEArchiveCustomizationSetting(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.CreateEArchiveCustomizationSetting)
}

// getEArchiveCustomizationSetting godoc
// @Summary Sorgulanan ayarı getirir
// @Description Sorgulanan ayarı getirir
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id} [get]
func getEArchiveCustomizationSetting(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveCustomizationSetting(company, id)
	})
}

// updateEArchiveCustomizationSetting godoc
// @Summary e-Belge özelleştirilebilir tasarımını güncellemek için kullanılır.
// @Description e-Belge özelleştirilebilir tasarımını güncellemek için kullanılır.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param request body map[string]interface{} true "Setting details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id} [put]
func updateEArchiveCustomizationSetting(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEArchiveCustomizationSetting(company, id, body)
	})
}

// deleteEArchiveCustomizationSetting godoc
// @Summary e-Belge özelleştirilebilir tasarımını silmek için kullanılır.
// @Description e-Belge özelleştirilebilir tasarımını silmek için kullanılır.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id} [delete]
func deleteEArchiveCustomizationSetting(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, nesService.DeleteEArchiveCustomizationSetting)
}

// setEArchiveCustomizationSettingDefault godoc
// @Summary Varsayılan ayarlar
// @Description Varsayılan ayarlar
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/setdefault [get]
func setEArchiveCustomizationSettingDefault(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.SetEArchiveCustomizationSettingDefault(company, id)
	})
}

// previewEArchiveCustomizationSetting godoc
// @Summary Tasarımı önizler
// @Description Tasarımı önizler
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param request body map[string]interface{} true "Preview details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/preview [post]
func previewEArchiveCustomizationSetting(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.PreviewEArchiveCustomizationSetting(company, id, body)
	})
}

// getEArchiveCustomizationSettingLogo godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan logoya bu uç ile ulaşılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan logoya bu uç ile ulaşılabilir.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/logo [get]
func getEArchiveCustomizationSettingLogo(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveCustomizationSettingLogo(company, id)
	})
}

// createEArchiveCustomizationSettingLogo godoc
// @Summary e-Belge özelleştirilebilir tasarıma logo eklemek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma logo eklemek için bu uç kullanılabilir.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param request body map[string]interface{} true "Logo details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/logo [post]
func createEArchiveCustomizationSettingLogo(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.CreateEArchiveCustomizationSettingLogo(company, id, body)
	})
}

// deleteEArchiveCustomizationSettingLogo godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan logoyu silmek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan logoyu silmek için bu uç kullanılabilir.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/logo [delete]
func deleteEArchiveCustomizationSettingLogo(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, nesService.DeleteEArchiveCustomizationSettingLogo)
}

// getEArchiveCustomizationSettingStamp godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeye bu uç ile ulaşılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeye bu uç ile ulaşılabilir.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/stamp [get]
func getEArchiveCustomizationSettingStamp(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveCustomizationSettingStamp(company, id)
	})
}

// createEArchiveCustomizationSettingStamp godoc
// @Summary e-Belge özelleştirilebilir tasarıma kaşe eklemek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma kaşe eklemek için bu uç kullanılabilir.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param request body map[string]interface{} true "Stamp details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/stamp [post]
func createEArchiveCustomizationSettingStamp(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.CreateEArchiveCustomizationSettingStamp(company, id, body)
	})
}

// deleteEArchiveCustomizationSettingStamp godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeyi silmek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeyi silmek için bu uç kullanılabilir.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/stamp [delete]
func deleteEArchiveCustomizationSettingStamp(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, nesService.DeleteEArchiveCustomizationSettingStamp)
}

// getEArchiveCustomizationSettingSignature godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan imzaya bu uç ile ulaşılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan imzaya bu uç ile ulaşılabilir.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/signature [get]
func getEArchiveCustomizationSettingSignature(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveCustomizationSettingSignature(company, id)
	})
}

// createEArchiveCustomizationSettingSignature godoc
// @Summary e-Belge özelleştirilebilir tasarıma imza eklemek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma imza eklemek için bu uç kullanılabilir.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Param request body map[string]interface{} true "Signature details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/signature [post]
func createEArchiveCustomizationSettingSignature(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.CreateEArchiveCustomizationSettingSignature(company, id, body)
	})
}

// deleteEArchiveCustomizationSettingSignature godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan imzayı silmek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan imzayı silmek için bu uç kullanılabilir.
// @Tags E-Archive Templates
// @Security BearerAuth
// @Param id path string true "Setting ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/customizationsettings/{id}/signature [delete]
func deleteEArchiveCustomizationSettingSignature(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, nesService.DeleteEArchiveCustomizationSettingSignature)
}

// getEArchiveSeries godoc
// @Summary Serileri listeler
// @Description Serileri listeler
// @Tags E-Archive Series
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/series [get]
func getEArchiveSeries(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveSeries)
}

// createEArchiveSerie godoc
// @Summary Seri ekler
// @Description Seri ekler
// @Tags E-Archive Series
// @Security BearerAuth
// @Param request body map[string]interface{} true "Serie details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/series [post]
func createEArchiveSerie(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.CreateEArchiveSerie)
}

// getEArchiveSerie godoc
// @Summary Sorgulanan seriyi getirir
// @Description Sorgulanan seriyi getirir
// @Tags E-Archive Series
// @Security BearerAuth
// @Param id path string true "Serie ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/series/{id} [get]
func getEArchiveSerie(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveSerie(company, id)
	})
}

// deleteEArchiveSerie godoc
// @Summary Seri siler
// @Description Seri siler
// @Tags E-Archive Series
// @Security BearerAuth
// @Param id path string true "Serie ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/series/{id} [delete]
func deleteEArchiveSerie(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, nesService.DeleteEArchiveSerie)
}

// getEArchiveSerieByPrefix godoc
// @Summary Ön eke göre seriyi getirir
// @Description Ön eke göre seriyi getirir
// @Tags E-Archive Series
// @Security BearerAuth
// @Param serie path string true "Serie Prefix"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/series/prefix/{serie} [get]
func getEArchiveSerieByPrefix(c *fiber.Ctx) error {
	serie := c.Params("serie")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveSerieByPrefix(company, serie)
	})
}

// setEArchiveSerieStatus godoc
// @Summary Seri durumunu günceller
// @Description Seri durumunu günceller
// @Tags E-Archive Series
// @Security BearerAuth
// @Param id path string true "Serie ID"
// @Param status path string true "Status"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/series/{id}/set/{status} [get]
func setEArchiveSerieStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.Params("status")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.SetEArchiveSerieStatus(company, id, status)
	})
}

// setEArchiveSerieDefault godoc
// @Summary Seriyi varsayılan ayarlar
// @Description Seriyi varsayılan ayarlar
// @Tags E-Archive Series
// @Security BearerAuth
// @Param id path string true "Serie ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/series/{id}/setdefault [get]
func setEArchiveSerieDefault(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.SetEArchiveSerieDefault(company, id)
	})
}

// setEArchiveSerieNumber godoc
// @Summary Sayaç günceller
// @Description Sayaç günceller
// @Tags E-Archive Series
// @Security BearerAuth
// @Param id path string true "Serie ID"
// @Param year path string true "Year"
// @Param nextNumber path string true "Next Number"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/series/{id}/{year}/setnumber/{nextNumber} [get]
func setEArchiveSerieNumber(c *fiber.Ctx) error {
	id := c.Params("id")
	year := c.Params("year")
	nextNumber := c.Params("nextNumber")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.SetEArchiveSerieNumber(company, id, year, nextNumber)
	})
}

// getEArchiveSerieHistories godoc
// @Summary Sayaç geçmişi
// @Description Sayaç geçmişi
// @Tags E-Archive Series
// @Security BearerAuth
// @Param serieId path string true "Serie ID"
// @Param year path string true "Year"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/series/{serieId}/{year}/histories [get]
func getEArchiveSerieHistories(c *fiber.Ctx) error {
	serieId := c.Params("serieId")
	year := c.Params("year")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.GetEArchiveSerieHistories(company, serieId, year)
	})
}

// getEArchiveDocumentTemplates godoc
// @Summary Tasarımları listeler
// @Description Tasarımları listeler
// @Tags E-Archive Document Templates
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates [get]
func getEArchiveDocumentTemplates(c *fiber.Ctx) error {
	return handleGenericEArchiveRequest(c, nesService.GetEArchiveDocumentTemplates)
}

// createEArchiveDocumentTemplate godoc
// @Summary Tasarım ekler
// @Description Tasarım ekler
// @Tags E-Archive Document Templates
// @Security BearerAuth
// @Param request body map[string]interface{} true "Template details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates [post]
func createEArchiveDocumentTemplate(c *fiber.Ctx) error {
	return handleGenericEArchivePostPutRequest(c, nesService.CreateEArchiveDocumentTemplate)
}

// downloadEArchiveDocumentTemplate godoc
// @Summary Tasarım dosyasını indirir
// @Description Tasarım dosyasını indirir
// @Tags E-Archive Document Templates
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {binary} binary
// @Router /v1/earchive/definitions/documenttemplates/{id} [get]
func downloadEArchiveDocumentTemplate(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")

	content, contentType, err := nesService.DownloadEArchiveDocumentTemplate(&company, id)
	if err != nil {
		return err
	}

	c.Set("Content-Type", contentType)
	return c.Send(content)
}

// updateEArchiveDocumentTemplate godoc
// @Summary Tasarımı günceller
// @Description Tasarımı günceller
// @Tags E-Archive Document Templates
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Param request body map[string]interface{} true "Template details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/{id} [put]
func updateEArchiveDocumentTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.UpdateEArchiveDocumentTemplate(company, id, body)
	})
}

// deleteEArchiveDocumentTemplate godoc
// @Summary Tasarımı siler
// @Description Tasarımı siler
// @Tags E-Archive Document Templates
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/{id} [delete]
func deleteEArchiveDocumentTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveDeleteRequest(c, id, nesService.DeleteEArchiveDocumentTemplate)
}

// setEArchiveDocumentTemplateDefault godoc
// @Summary Tasarımı varsayılan ayarlar
// @Description Tasarımı varsayılan ayarlar
// @Tags E-Archive Document Templates
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/{id}/setdefault [get]
func setEArchiveDocumentTemplateDefault(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchiveRequest(c, func(company *models.Company, params map[string]string) (map[string]interface{}, error) {
		return nesService.SetEArchiveDocumentTemplateDefault(company, id)
	})
}

// previewEArchiveDocumentTemplate godoc
// @Summary Tasarımı önizler
// @Description Tasarımı önizler
// @Tags E-Archive Document Templates
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Param request body map[string]interface{} true "Preview details"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/definitions/documenttemplates/{id}/preview [post]
func previewEArchiveDocumentTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return handleGenericEArchivePostPutRequest(c, func(company *models.Company, body map[string]interface{}) (map[string]interface{}, error) {
		return nesService.PreviewEArchiveDocumentTemplate(company, id, body)
	})
}
