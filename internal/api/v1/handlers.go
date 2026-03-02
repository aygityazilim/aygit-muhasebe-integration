package v1

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/internal/service"
	"aygit-muhasebe-integration/pkg/db"
	"aygit-muhasebe-integration/pkg/errors"
	"io"

	"github.com/gofiber/fiber/v2"
)

var nesService = service.NewNESService()

// getCreditSummary godoc
// @Summary NES Kontör Bakiyesini Getir
// @Description Mevcut oturumdaki firmanın NES Özel Entegratör üzerindeki kontör özetini döner.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /v1/management/creditsummary [get]
func getCreditSummary(c *fiber.Ctx) error {
	// Middleware üzerinden gelen kullanıcı bilgisi alınır
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	// DB'den firma ayarları çekilir
	var company models.Company
	err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1 AND deleted_at IS NULL", *user.CompanyID)
	if err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma yapılandırması yüklenemedi")
	}

	// NES Servis sorgusu firma ayarlarıyla yapılır
	summary, err := nesService.GetCreditSummary(&company)
	if err != nil {
		return err // Global ErrorHandler tarafından yakalanır
	}

	// Başarılı yanıt dönülür
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   summary,
	})
}

// getIdentifications godoc
// @Summary Kimlik Bilgilerini Listele
// @Description Firmanın kimlik bilgilerini (Party Identification) listeler.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/identifications [get]
func GetIdentifications(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	// TODO: DB'den veya NES üzerinden kimlik bilgilerini çek
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   []interface{}{},
	})
}

// createIdentification godoc
// @Summary Yeni Kimlik Bilgisi Ekle
// @Description Firmaya yeni bir kimlik bilgisi (VKN/TCKN vb.) ekler.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 201 {object} map[string]string
// @Router /v1/management/identifications [post]
func createIdentification(c *fiber.Ctx) error {
	// TODO: İstek gövdesini doğrula ve DB'ye kaydet
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Kimlik bilgisi başarıyla oluşturuldu",
	})
}

// getAccountDefaultDocumentParameter godoc
// @Summary Firmanın Varsayılan Doküman Parametresini döner
// @Description Fatura Senaryosu, Gönderim Tipi, Satış Kanalı ayarlarını döner. NES Özel Entegratör API /v1/accountDefaultDocumentParameter ucunu çağırır.
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Başarılı istek. Örnek: {'eInvoiceInvoiceProfile': 'string', 'eArchiveSendType': 'string', 'eArchiveSalesPlatform': 'string'}"
// @Failure 401 {object} errors.AppError "Kullanıcı firma bilgisi bulunamadı"
// @Failure 403 {object} errors.AppError "Yetkisiz Erişim | Aşağıdaki Scopelardan birisine sahip olmadığınız durumda dönülür"
// @Failure 500 {object} errors.AppError "Sunucu hatası veya NES servis hatası"
// @Router /v1/account/default-document-parameter [get]
func getAccountDefaultDocumentParameter(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1 AND deleted_at IS NULL", *user.CompanyID)
	if err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma yapılandırması yüklenemedi")
	}

	parameters, err := nesService.GetAccountDefaultDocumentParameter(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(parameters)
}

// getAccountDocumentArchives godoc
// @Summary Belge arşivi talebi listeler
// @Description Belge arşivi taleplerinize bu uç ile ulaşabilirsiniz. NES Özel Entegratör API /v1/accountDocumentArchives ucunu çağırır.
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} interface{} "Başarılı istek. Dönen liste içerisinde arşiv talepleri bulunur."
// @Failure 401 {object} errors.AppError "Kullanıcı firma bilgisi bulunamadı"
// @Failure 403 {object} errors.AppError "Yetkisiz Erişim | Aşağıdaki Scopelardan birisine sahip olmadığınız durumda dönülür"
// @Failure 500 {object} errors.AppError "Sunucu hatası veya NES servis hatası"
// @Router /v1/account/document-archives [get]
func getAccountDocumentArchives(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1 AND deleted_at IS NULL", *user.CompanyID)
	if err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma yapılandırması yüklenemedi")
	}

	archives, err := nesService.GetAccountDocumentArchives(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(archives)
}

// --- E-Fatura (Invoices) Handlers ---

// getIncomingInvoices godoc
// @Summary Gelen E-Faturaları Listele
// @Description NES üzerindeki gelen e-faturaları filtreleyerek listeler.
// @Tags Invoices
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid query string false "Fatura UUID"
// @Param documentNumber query string false "Fatura Numarası"
// @Param startDate query string false "Başlangıç Tarihi (ISO8601)"
// @Param endDate query string false "Bitiş Tarihi (ISO8601)"
// @Success 200 {object} map[string]interface{}
// @Router /v1/invoices/incoming [get]
func getIncomingInvoices(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	queryParams := c.Queries()
	result, err := nesService.GetIncomingInvoices(&company, queryParams)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// getOutgoingInvoices godoc
// @Summary Giden E-Faturaları Listele
// @Description NES üzerindeki giden e-faturaları filtreleyerek listeler.
// @Tags Invoices
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid query string false "Fatura UUID"
// @Param documentNumber query string false "Fatura Numarası"
// @Success 200 {object} map[string]interface{}
// @Router /v1/invoices/outgoing [get]
func getOutgoingInvoices(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	queryParams := c.Queries()
	result, err := nesService.GetOutgoingInvoices(&company, queryParams)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// downloadInvoiceFile godoc
// @Summary Fatura Dosyası İndir
// @Description Belirtilen faturanın XML, PDF veya HTML dosyasını indirir.
// @Tags Invoices
// @Security BearerAuth
// @Param uuid path string true "Fatura UUID"
// @Param direction path string true "Yön (incoming/outgoing)"
// @Param fileType path string true "Dosya Tipi (xml/pdf/html)"
// @Success 200 {binary} binary
// @Router /v1/invoices/{direction}/{uuid}/{fileType} [get]
func downloadInvoiceFile(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	uuid := c.Params("uuid")
	direction := c.Params("direction")
	fileType := c.Params("fileType")

	content, contentType, err := nesService.DownloadInvoiceFile(&company, uuid, direction, fileType)
	if err != nil {
		return err
	}

	c.Set("Content-Type", contentType)
	return c.Send(content)
}

// --- E-İrsaliye (Despatches) Handlers ---

// getIncomingDespatches godoc
// @Summary Gelen E-İrsaliyeleri Listele
// @Tags Despatches
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/despatches/incoming [get]
func getIncomingDespatches(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	result, err := nesService.GetIncomingDespatches(&company, c.Queries())
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// getOutgoingDespatches godoc
// @Summary Giden E-İrsaliyeleri Listele
// @Tags Despatches
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/despatches/outgoing [get]
func getOutgoingDespatches(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	result, err := nesService.GetOutgoingDespatches(&company, c.Queries())
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// --- E-Arşiv (E-Archive) Handlers ---

// getEArchiveInvoices godoc
// @Summary E-Arşiv Faturalarını Listele
// @Tags E-Archive
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices [get]
func getEArchiveInvoices(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	result, err := nesService.GetEArchiveInvoices(&company, c.Queries())
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// downloadEArchiveFile godoc
// @Summary E-Arşiv Dosyası İndir
// @Tags E-Archive
// @Security BearerAuth
// @Param uuid path string true "Fatura UUID"
// @Param fileType path string true "Dosya Tipi (xml/pdf/html)"
// @Success 200 {binary} binary
// @Router /v1/earchive/invoices/{uuid}/{fileType} [get]
func downloadEArchiveFile(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	content, contentType, err := nesService.DownloadEArchiveFile(&company, c.Params("uuid"), c.Params("fileType"))
	if err != nil {
		return err
	}
	c.Set("Content-Type", contentType)
	return c.Send(content)
}

// cancelEArchiveInvoice godoc
// @Summary E-Arşiv Faturasını İptal Et
// @Tags E-Archive
// @Security BearerAuth
// @Accept json
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/invoices/cancel [post]
func cancelEArchiveInvoice(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	var cancelData map[string]interface{}
	if err := c.BodyParser(&cancelData); err != nil {
		return err
	}

	result, err := nesService.CancelEArchiveInvoice(&company, cancelData)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// uploadInvoice godoc
// @Summary Yeni Fatura Yükle (Taslak)
// @Description NES Özel Entegratör sistemine yeni bir XML fatura yükler.
// @Tags Invoices
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param File formData file true "Fatura XML Dosyası"
// @Param IsDirectSend formData string false "Doğrudan Gönder (true/false)"
// @Success 200 {object} map[string]interface{}
// @Router /v1/invoices/upload [post]
func uploadInvoice(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	file, err := c.FormFile("File")
	if err != nil {
		return err
	}
	f, _ := file.Open()
	xmlData, _ := io.ReadAll(f)
	defer f.Close()

	params := c.Queries() // Veya form fieldları
	result, err := nesService.UploadInvoice(&company, xmlData, params)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// sendDraftInvoices godoc
// @Summary Taslak Faturaları Gönder
// @Description Taslak olarak yüklenmiş faturaları onaylayıp resmileştirir.
// @Tags Invoices
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuids body []string true "Fatura UUID Listesi"
// @Success 200 {array} interface{}
// @Router /v1/invoices/send-draft [post]
func sendDraftInvoices(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	var uuids []string
	if err := c.BodyParser(&uuids); err != nil {
		return err
	}

	result, err := nesService.SendDraftInvoices(&company, uuids)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// uploadDespatch godoc
// @Summary Yeni İrsaliye Yükle (Taslak)
// @Tags Despatches
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param File formData file true "İrsaliye XML Dosyası"
// @Success 200 {object} map[string]interface{}
// @Router /v1/despatches/upload [post]
func uploadDespatch(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	file, err := c.FormFile("File")
	if err != nil {
		return err
	}
	f, _ := file.Open()
	xmlData, _ := io.ReadAll(f)
	defer f.Close()

	result, err := nesService.UploadDespatch(&company, xmlData, c.Queries())
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// sendDespatchAnswer godoc
// @Summary İrsaliye Yanıtı Gönder
// @Tags Despatches
// @Security BearerAuth
// @Param uuid path string true "İrsaliye UUID"
// @Param answer body map[string]interface{} true "Yanıt Detayları"
// @Success 200 {object} map[string]interface{}
// @Router /v1/despatches/{uuid}/answer [post]
func sendDespatchAnswer(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	var answer map[string]interface{}
	if err := c.BodyParser(&answer); err != nil {
		return err
	}

	result, err := nesService.SendDespatchAnswer(&company, c.Params("uuid"), answer)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// uploadEArchiveInvoice godoc
// @Summary Yeni E-Arşiv Faturası Yükle
// @Tags E-Archive
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param File formData file true "Fatura XML Dosyası"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/upload [post]
func uploadEArchiveInvoice(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	file, err := c.FormFile("File")
	if err != nil {
		return err
	}
	f, _ := file.Open()
	xmlData, _ := io.ReadAll(f)
	defer f.Close()

	result, err := nesService.UploadEArchiveInvoice(&company, xmlData, c.Queries())
	if err != nil {
		return err
	}
	return c.JSON(result)
}

// sendDraftEArchiveInvoices godoc
// @Summary Taslak E-Arşiv Faturalarını Gönder
// @Tags E-Archive
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuids body []string true "Fatura UUID Listesi"
// @Success 200 {object} map[string]interface{}
// @Router /v1/earchive/send-draft [post]
func sendDraftEArchiveInvoices(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	var uuids []string
	if err := c.BodyParser(&uuids); err != nil {
		return err
	}

	result, err := nesService.SendDraftEArchiveInvoices(&company, uuids)
	if err != nil {
		return err
	}
	return c.JSON(result)
}
