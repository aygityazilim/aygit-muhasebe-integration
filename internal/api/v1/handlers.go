package v1

import (
	"encoding/json"

	"fmt"

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
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/identifications [get]
func getIdentifications(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetIdentifications(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

// createIdentification godoc
// @Summary Yeni Kimlik Bilgisi Ekle
// @Description Firmaya yeni bir kimlik bilgisi (VKN/TCKN vb.) ekler.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Identification data"
// @Success 201 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/identifications [post]
func createIdentification(c *fiber.Ctx) error {
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

	res, err := nesService.CreateIdentification(&company, payload)
	if err != nil {
		return err
	}

	// Update party_identification in DB
	if vkn, ok := payload["partyIdentification"].(string); ok {
		_, err = db.DB.Exec("UPDATE companies SET party_identification = $1 WHERE id = $2", vkn, company.ID)
		if err == nil {
			db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "UPDATE_COMPANY_IDENTIFICATION", fmt.Sprintf(`{"vkn": "%s"}`, vkn))
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   res,
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

// deleteIdentification godoc
// @Summary Kimlik Bilgisini Sil
// @Description Firmadan kimlik bilgisini siler.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Identification ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/identifications/{id} [delete]
func deleteIdentification(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteIdentification(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEIDENTIFICATION", fmt.Sprintf(`{"id": "%s"}`, id))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// getDealerInfo godoc
// @Summary Bayi Bilgilerini Getir
// @Description Bayi bilgilerine ulaşabilirsiniz.
// @Tags Management
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/dealerinfo [get]
func getDealerInfo(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetDealerInfo(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// getAddresses godoc
// @Summary Adresleri Listele
// @Description Adres bilgilerini getirir.
// @Tags Management
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/addresses [get]
func getAddresses(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetAddresses(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// createAddress godoc
// @Summary Adres Ekle
// @Description Adres bilgisi ekler.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Address data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/management/addresses [post]
func createAddress(c *fiber.Ctx) error {
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

	res, err := nesService.CreateAddress(&company, payload)
	if err != nil {
		return err
	}

	// Insert into DB
	title, _ := payload["title"].(string)
	streetName, _ := payload["streetName"].(string)
	cityName, _ := payload["cityName"].(string)
	countryName, _ := payload["countryName"].(string)

	_, err = db.DB.Exec("INSERT INTO addresses (title, street_name, city_name, country_name) VALUES ($1, $2, $3, $4)", title, streetName, cityName, countryName)
	if err == nil {
		db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "CREATE_ADDRESS", fmt.Sprintf(`{"title": "%s"}`, title))
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": res})
}

// updateAddress godoc
// @Summary Adres Güncelle
// @Description Adres bilgisi günceller.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Address ID"
// @Param request body map[string]interface{} true "Address data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/management/address/{id} [put]
func updateAddress(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return errors.NewError(fiber.StatusBadRequest, errors.ErrCodeInvalidRequest, "Geçersiz istek gövdesi")
	}

	res, err := nesService.UpdateAddress(&company, id, payload)
	if err != nil {
		return err
	}

	// Update in DB (Assume ID is uuid)
	title, _ := payload["title"].(string)
	streetName, _ := payload["streetName"].(string)
	cityName, _ := payload["cityName"].(string)
	countryName, _ := payload["countryName"].(string)

	_, err = db.DB.Exec("UPDATE addresses SET title=$1, street_name=$2, city_name=$3, country_name=$4 WHERE id=$5", title, streetName, cityName, countryName, id)
	if err == nil {
		db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "UPDATE_ADDRESS", fmt.Sprintf(`{"id": "%s"}`, id))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// deleteAddress godoc
// @Summary Adres Sil
// @Description Adres bilgisini siler.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Address ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/address/{id} [delete]
func deleteAddress(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteAddress(&company, id)
	if err != nil {
		return err
	}

	// Delete from DB
	_, err = db.DB.Exec("DELETE FROM addresses WHERE id=$1", id)
	if err == nil {
		db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETE_ADDRESS", fmt.Sprintf(`{"id": "%s"}`, id))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// getIVDSetting godoc
// @Summary IVD Ayarlarını Getir
// @Description IVD ayarlarını listeler.
// @Tags Management
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/ivd [get]
func getIVDSetting(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetIVD(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// updateIVDSetting godoc
// @Summary IVD Ayarlarını Güncelle
// @Description IVD ayarlarını günceller.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "IVD Setting data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/management/ivd [put]
func updateIVDSetting(c *fiber.Ctx) error {
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

	res, err := nesService.UpdateIVD(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "UPDATEIVD", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// getLucaIntegrationSetting godoc
// @Summary Luca Entegrasyon Ayarlarını Getir
// @Description Luca entegrasyon ayarlarını listeler.
// @Tags Management
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/luca/integration/setting [get]
func getLucaIntegrationSetting(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetLucaIntegrationSetting(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// updateLucaIntegrationSetting godoc
// @Summary Luca Entegrasyon Ayarlarını Güncelle
// @Description Luca entegrasyon ayarlarını günceller.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Luca Integration Setting data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/management/luca/integration/setting [put]
func updateLucaIntegrationSetting(c *fiber.Ctx) error {
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

	res, err := nesService.UpdateLucaIntegrationSetting(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "UPDATELUCAINTEGRATIONSETTING", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// getCustomerSearchFromGIBSetting godoc
// @Summary GIB Müşteri Arama Ayarlarını Getir
// @Description GIB müşteri arama ayarlarını listeler.
// @Tags Management
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} errors.AppError
// @Router /v1/management/customer_search_from_gib/setting [get]
func getCustomerSearchFromGIBSetting(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetCustomerSearchFromGIBSetting(&company)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// updateCustomerSearchFromGIBSetting godoc
// @Summary GIB Müşteri Arama Ayarlarını Güncelle
// @Description GIB müşteri arama ayarlarını günceller.
// @Tags Management
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "GIB Customer Search Setting data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/management/customer_search_from_gib/setting [put]
func updateCustomerSearchFromGIBSetting(c *fiber.Ctx) error {
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

	res, err := nesService.UpdateCustomerSearchFromGIBSetting(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "UPDATECUSTOMERSEARCHFROMGIBSETTING", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// updateAccountDefaultDocumentParameter godoc
// @Summary Firmanın Varsayılan Doküman Parametresini günceller
// @Description Fatura Senaryosu, Gönderim Tipi, Satış Kanalı ayarlarını günceller.
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Document Parameter data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/account/default-document-parameter [put]
func updateAccountDefaultDocumentParameter(c *fiber.Ctx) error {
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

	res, err := nesService.UpdateAccountDefaultDocumentParameter(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "UPDATEACCOUNTDEFAULTDOCUMENTPARAMETER", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// createAccountDocumentArchive godoc
// @Summary Belge arşivi talebi oluşturur
// @Description Belge arşivi talebi oluşturur.
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Archive request data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/account/document-archives [post]
func createAccountDocumentArchive(c *fiber.Ctx) error {
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

	res, err := nesService.CreateAccountDocumentArchive(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "CREATEACCOUNTDOCUMENTARCHIVE", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": res})
}

// accountModulesInfo godoc
// @Summary Hesap bilgilerini ve işlem bekleyen zarfları döner
// @Description Hesap bilgilerini ve işlem bekleyen zarfları döner.
// @Tags AccountModules
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Modules Info query"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/accountmodules/info [post]
func accountModulesInfo(c *fiber.Ctx) error {
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

	res, err := nesService.AccountModulesInfo(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "ACCOUNTMODULESINFO", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// accountModulesUpdate godoc
// @Summary Hesap bilgilerini günceller
// @Description Hesap bilgilerini günceller.
// @Tags AccountModules
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Modules Update query"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/accountmodules/update [post]
func accountModulesUpdate(c *fiber.Ctx) error {
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

	res, err := nesService.AccountModulesUpdate(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "ACCOUNTMODULESUPDATE", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// accountModulesGetEnvelopeContent godoc
// @Summary Zarfın içeriğini getirir
// @Description Zarfın içeriğini getirir.
// @Tags AccountModules
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Envelope Content request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/accountmodules/getenvelopecontent [post]
func accountModulesGetEnvelopeContent(c *fiber.Ctx) error {
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

	res, err := nesService.AccountModulesGetEnvelopeContent(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "ACCOUNTMODULESGETENVELOPECONTENT", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// accountModulesSetSignedContent godoc
// @Summary Zarfı imzalar
// @Description Zarfı imzalar.
// @Tags AccountModules
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Signed Content request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/accountmodules/setsignedcontent [post]
func accountModulesSetSignedContent(c *fiber.Ctx) error {
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

	res, err := nesService.AccountModulesSetSignedContent(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "ACCOUNTMODULESSETSIGNEDCONTENT", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}

// accountModulesGetEnvelopeInfo godoc
// @Summary Zarfın detaylarını getirir
// @Description Zarfın detaylarını getirir.
// @Tags AccountModules
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Envelope Info request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/accountmodules/getenvelopeinfo [post]
func accountModulesGetEnvelopeInfo(c *fiber.Ctx) error {
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

	res, err := nesService.AccountModulesGetEnvelopeInfo(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "ACCOUNTMODULESGETENVELOPEINFO", func() string { b, _ := json.Marshal(map[string]interface{}{"payload": payload}); return string(b) }())

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}
