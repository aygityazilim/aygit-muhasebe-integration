package v1

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/db"
	"aygit-muhasebe-integration/pkg/errors"

	"github.com/gofiber/fiber/v2"
)


// getDefinitionsFileexporttitlesDocumenttypeTitlekeysEinvoice godoc
// @Summary Kullanılabilir alanları listeler
// @Description Kullanılabilir alanları listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param documentType path string true "documentType"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/fileexporttitles/:documentType/titlekeys [get]
func getDefinitionsFileexporttitlesDocumenttypeTitlekeysEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	documentType := c.Params("documentType")
	res, err := nesService.GetDefinitionsFileexporttitlesDocumenttypeTitlekeysEinvoice(&company, documentType)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSFILEEXPORTTITLESDOCUMENTTYPETITLEKEYSEINVOICE", "Kullanılabilir alanları listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsFileexporttitlesDocumenttypeExtensionEinvoice godoc
// @Summary Tanımları getirir
// @Description Tanımları getirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param documentType path string true "documentType"
// @Param extension path string true "extension"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/fileexporttitles/:documentType/:extension [get]
func getDefinitionsFileexporttitlesDocumenttypeExtensionEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	documentType := c.Params("documentType")
	extension := c.Params("extension")
	res, err := nesService.GetDefinitionsFileexporttitlesDocumenttypeExtensionEinvoice(&company, documentType, extension)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSFILEEXPORTTITLESDOCUMENTTYPEEXTENSIONEINVOICE", "Tanımları getirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putDefinitionsFileexporttitlesEinvoice godoc
// @Summary Tanımları günceller
// @Description Tanımları günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Tanımları günceller isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/fileexporttitles [put]
func putDefinitionsFileexporttitlesEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PutDefinitionsFileexporttitlesEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTDEFINITIONSFILEEXPORTTITLESEINVOICE", "Tanımları günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putUploadsDocumentUuidEinvoice godoc
// @Summary Belge günceller
// @Description Belge günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Belge günceller isteği"
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/uploads/document/:uuid [put]
func putUploadsDocumentUuidEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	res, err := nesService.PutUploadsDocumentUuidEinvoice(&company, payload, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTUPLOADSDOCUMENTUUIDEINVOICE", "Belge günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postUploadsDraftCreateIdEinvoice godoc
// @Summary Belge yükler
// @Description Belge yükler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Belge yükler isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/uploads/draft/create/:id [post]
func postUploadsDraftCreateIdEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PostUploadsDraftCreateIdEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTUPLOADSDRAFTCREATEIDEINVOICE", "Belge yükler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postUploadsMarketplacesIdOrdersOrderidPreviewEinvoice godoc
// @Summary Belirtilen pazaryerindeki siparişin faturasını önizler
// @Description Belirtilen pazaryerindeki siparişin faturasını önizler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Belirtilen pazaryerindeki siparişin faturasını önizler isteği"
// @Param id path string true "id"
// @Param orderId path string true "orderId"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/uploads/marketplaces/:id/orders/:orderId/preview [post]
func postUploadsMarketplacesIdOrdersOrderidPreviewEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	orderId := c.Params("orderId")
	res, err := nesService.PostUploadsMarketplacesIdOrdersOrderidPreviewEinvoice(&company, payload, id, orderId)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTUPLOADSMARKETPLACESIDORDERSORDERIDPREVIEWEINVOICE", "Belirtilen pazaryerindeki siparişin faturasını önizler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postUploadsMarketplacesIdOrdersCreateinvoiceEinvoice godoc
// @Summary Belirtilen pazaryerindeki siparişin faturasını önizler
// @Description Belirtilen pazaryerindeki siparişin faturasını önizler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Belirtilen pazaryerindeki siparişin faturasını önizler isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/uploads/marketplaces/:id/orders/createinvoice [post]
func postUploadsMarketplacesIdOrdersCreateinvoiceEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PostUploadsMarketplacesIdOrdersCreateinvoiceEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTUPLOADSMARKETPLACESIDORDERSCREATEINVOICEEINVOICE", "Belirtilen pazaryerindeki siparişin faturasını önizler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postExinvoicesEinvoice godoc
// @Summary Eski belge yükler
// @Description Eski belge yükler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Eski belge yükler isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices [post]
func postExinvoicesEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostExinvoicesEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTEXINVOICESEINVOICE", "Eski belge yükler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesQueueEinvoice godoc
// @Summary Yükleme kuyruğunu listeler
// @Description Yükleme kuyruğunu listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/queue [get]
func getExinvoicesQueueEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetExinvoicesQueueEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESQUEUEEINVOICE", "Yükleme kuyruğunu listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesQueueIdEinvoice godoc
// @Summary Yükleme sonucunu indir
// @Description Yükleme sonucunu indir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/queue/:id [get]
func getExinvoicesQueueIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetExinvoicesQueueIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESQUEUEIDEINVOICE", "Yükleme sonucunu indir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getTagsEinvoice godoc
// @Summary Etiketleri listeler
// @Description Etiketleri listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/tags [get]
func getTagsEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetTagsEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETTAGSEINVOICE", "Etiketleri listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postTagsEinvoice godoc
// @Summary Etiket ekler
// @Description Etiket ekler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Etiket ekler isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/tags [post]
func postTagsEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostTagsEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTTAGSEINVOICE", "Etiket ekler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getTagsIdEinvoice godoc
// @Summary Sorgulanan etiketi getirir
// @Description Sorgulanan etiketi getirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/tags/:id [get]
func getTagsIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetTagsIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETTAGSIDEINVOICE", "Sorgulanan etiketi getirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putTagsIdEinvoice godoc
// @Summary Etiket günceller
// @Description Etiket günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Etiket günceller isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/tags/:id [put]
func putTagsIdEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PutTagsIdEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTTAGSIDEINVOICE", "Etiket günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteTagsIdEinvoice godoc
// @Summary Etiket siler
// @Description Etiket siler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/tags/:id [delete]
func deleteTagsIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteTagsIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETETAGSIDEINVOICE", "Etiket siler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesIncomingEinvoice godoc
// @Summary Belgeleri listeler
// @Description Belgeleri listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/incoming [get]
func getExinvoicesIncomingEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetExinvoicesIncomingEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESINCOMINGEINVOICE", "Belgeleri listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesIncomingUuidXmlEinvoice godoc
// @Summary XML İndir
// @Description XML İndir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/incoming/:uuid/xml [get]
func getExinvoicesIncomingUuidXmlEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	uuid := c.Params("uuid")
	res, err := nesService.GetExinvoicesIncomingUuidXmlEinvoice(&company, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESINCOMINGUUIDXMLEINVOICE", "XML İndir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesIncomingUuidPdfEinvoice godoc
// @Summary PDF İndir
// @Description PDF İndir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/incoming/:uuid/pdf [get]
func getExinvoicesIncomingUuidPdfEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	uuid := c.Params("uuid")
	res, err := nesService.GetExinvoicesIncomingUuidPdfEinvoice(&company, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESINCOMINGUUIDPDFEINVOICE", "PDF İndir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesIncomingUuidHtmlEinvoice godoc
// @Summary Belgeyi görüntüler
// @Description Belgeyi görüntüler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/incoming/:uuid/html [get]
func getExinvoicesIncomingUuidHtmlEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	uuid := c.Params("uuid")
	res, err := nesService.GetExinvoicesIncomingUuidHtmlEinvoice(&company, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESINCOMINGUUIDHTMLEINVOICE", "Belgeyi görüntüler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postExinvoicesIncomingExportFiletypeEinvoice godoc
// @Summary Dışarı Aktar
// @Description Dışarı Aktar.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Dışarı Aktar isteği"
// @Param fileType path string true "fileType"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/incoming/export/:fileType [post]
func postExinvoicesIncomingExportFiletypeEinvoice(c *fiber.Ctx) error {
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

	fileType := c.Params("fileType")
	res, err := nesService.PostExinvoicesIncomingExportFiletypeEinvoice(&company, payload, fileType)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTEXINVOICESINCOMINGEXPORTFILETYPEEINVOICE", "Dışarı Aktar")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getNotificationsIncomingDynamicrulesEinvoice godoc
// @Summary Kuralları listeler
// @Description Kuralları listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/incoming/dynamicrules [get]
func getNotificationsIncomingDynamicrulesEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetNotificationsIncomingDynamicrulesEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETNOTIFICATIONSINCOMINGDYNAMICRULESEINVOICE", "Kuralları listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postNotificationsIncomingDynamicrulesEinvoice godoc
// @Summary Kural oluşturur
// @Description Kural oluşturur.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Kural oluşturur isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/incoming/dynamicrules [post]
func postNotificationsIncomingDynamicrulesEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostNotificationsIncomingDynamicrulesEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTNOTIFICATIONSINCOMINGDYNAMICRULESEINVOICE", "Kural oluşturur")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getNotificationsIncomingDynamicrulesIdEinvoice godoc
// @Summary Sorgulanan kuralı getirir
// @Description Sorgulanan kuralı getirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/incoming/dynamicrules/:id [get]
func getNotificationsIncomingDynamicrulesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetNotificationsIncomingDynamicrulesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETNOTIFICATIONSINCOMINGDYNAMICRULESIDEINVOICE", "Sorgulanan kuralı getirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putNotificationsIncomingDynamicrulesIdEinvoice godoc
// @Summary Kural günceller
// @Description Kural günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Kural günceller isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/incoming/dynamicrules/:id [put]
func putNotificationsIncomingDynamicrulesIdEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PutNotificationsIncomingDynamicrulesIdEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTNOTIFICATIONSINCOMINGDYNAMICRULESIDEINVOICE", "Kural günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteNotificationsIncomingDynamicrulesIdEinvoice godoc
// @Summary Kural siler
// @Description Kural siler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/incoming/dynamicrules/:id [delete]
func deleteNotificationsIncomingDynamicrulesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteNotificationsIncomingDynamicrulesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETENOTIFICATIONSINCOMINGDYNAMICRULESIDEINVOICE", "Kural siler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postIncomingInvoicesExportFiletypeEinvoice godoc
// @Summary Toplu aktar
// @Description Toplu aktar.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Toplu aktar isteği"
// @Param fileType path string true "fileType"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/export/:fileType [post]
func postIncomingInvoicesExportFiletypeEinvoice(c *fiber.Ctx) error {
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

	fileType := c.Params("fileType")
	res, err := nesService.PostIncomingInvoicesExportFiletypeEinvoice(&company, payload, fileType)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTINCOMINGINVOICESEXPORTFILETYPEEINVOICE", "Toplu aktar")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getIncomingReportmoduleReportsEinvoice godoc
// @Summary Rapor listeler
// @Description Rapor listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/reportmodule/reports [get]
func getIncomingReportmoduleReportsEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetIncomingReportmoduleReportsEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETINCOMINGREPORTMODULEREPORTSEINVOICE", "Rapor listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postIncomingReportmoduleReportsEinvoice godoc
// @Summary Rapor oluşturur
// @Description Rapor oluşturur.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Rapor oluşturur isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/reportmodule/reports [post]
func postIncomingReportmoduleReportsEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostIncomingReportmoduleReportsEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTINCOMINGREPORTMODULEREPORTSEINVOICE", "Rapor oluşturur")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getIncomingReportmoduleReportsIdDownloadEinvoice godoc
// @Summary Rapor indirir
// @Description Rapor indirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/reportmodule/reports/:id/download [get]
func getIncomingReportmoduleReportsIdDownloadEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetIncomingReportmoduleReportsIdDownloadEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETINCOMINGREPORTMODULEREPORTSIDDOWNLOADEINVOICE", "Rapor indirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getIncomingReportmoduleTemplatesEinvoice godoc
// @Summary Şablonları listeler
// @Description Şablonları listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/reportmodule/templates [get]
func getIncomingReportmoduleTemplatesEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetIncomingReportmoduleTemplatesEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETINCOMINGREPORTMODULETEMPLATESEINVOICE", "Şablonları listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postIncomingReportmoduleTemplatesEinvoice godoc
// @Summary Rapor şablonu oluşturur
// @Description Rapor şablonu oluşturur.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Rapor şablonu oluşturur isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/reportmodule/templates [post]
func postIncomingReportmoduleTemplatesEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostIncomingReportmoduleTemplatesEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTINCOMINGREPORTMODULETEMPLATESEINVOICE", "Rapor şablonu oluşturur")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getIncomingReportmoduleTemplatesIdEinvoice godoc
// @Summary Sorgulanan şablonu getirir
// @Description Sorgulanan şablonu getirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/reportmodule/templates/:id [get]
func getIncomingReportmoduleTemplatesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetIncomingReportmoduleTemplatesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETINCOMINGREPORTMODULETEMPLATESIDEINVOICE", "Sorgulanan şablonu getirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putIncomingReportmoduleTemplatesIdEinvoice godoc
// @Summary Rapor şablonunu günceller
// @Description Rapor şablonunu günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Rapor şablonunu günceller isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/reportmodule/templates/:id [put]
func putIncomingReportmoduleTemplatesIdEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PutIncomingReportmoduleTemplatesIdEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTINCOMINGREPORTMODULETEMPLATESIDEINVOICE", "Rapor şablonunu günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteIncomingReportmoduleTemplatesIdEinvoice godoc
// @Summary Rapor Şablonunu siler
// @Description Rapor Şablonunu siler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/reportmodule/templates/:id [delete]
func deleteIncomingReportmoduleTemplatesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteIncomingReportmoduleTemplatesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEINCOMINGREPORTMODULETEMPLATESIDEINVOICE", "Rapor Şablonunu siler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getIncomingReportmoduleColumnsEinvoice godoc
// @Summary Kolonları listeler
// @Description Kolonları listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/reportmodule/columns [get]
func getIncomingReportmoduleColumnsEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetIncomingReportmoduleColumnsEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETINCOMINGREPORTMODULECOLUMNSEINVOICE", "Kolonları listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putIncomingInvoicesTagsEinvoice godoc
// @Summary Etiket ekler/çıkarır
// @Description Etiket ekler/çıkarır.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Etiket ekler/çıkarır isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/tags [put]
func putIncomingInvoicesTagsEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PutIncomingInvoicesTagsEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTINCOMINGINVOICESTAGSEINVOICE", "Etiket ekler/çıkarır")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postIncomingInvoicesUuidSavecompanyindocumentEinvoice godoc
// @Summary Firma olarak kaydet
// @Description Firma olarak kaydet.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Firma olarak kaydet isteği"
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/:uuid/savecompanyindocument [post]
func postIncomingInvoicesUuidSavecompanyindocumentEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	res, err := nesService.PostIncomingInvoicesUuidSavecompanyindocumentEinvoice(&company, payload, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTINCOMINGINVOICESUUIDSAVECOMPANYINDOCUMENTEINVOICE", "Firma olarak kaydet")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postIncomingInvoicesUuidDocumentanswerEinvoice godoc
// @Summary Belge'ye cevap verir
// @Description Belge'ye cevap verir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Belge'ye cevap verir isteği"
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/:uuid/documentAnswer [post]
func postIncomingInvoicesUuidDocumentanswerEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	res, err := nesService.PostIncomingInvoicesUuidDocumentanswerEinvoice(&company, payload, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTINCOMINGINVOICESUUIDDOCUMENTANSWEREINVOICE", "Belge'ye cevap verir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putIncomingInvoicesBulkOperationEinvoice godoc
// @Summary Yeni durum atar
// @Description Yeni durum atar.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Yeni durum atar isteği"
// @Param operation path string true "operation"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/bulk/:operation [put]
func putIncomingInvoicesBulkOperationEinvoice(c *fiber.Ctx) error {
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

	operation := c.Params("operation")
	res, err := nesService.PutIncomingInvoicesBulkOperationEinvoice(&company, payload, operation)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTINCOMINGINVOICESBULKOPERATIONEINVOICE", "Yeni durum atar")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postIncomingInvoicesIdCreatereturninvoiceEinvoice godoc
// @Summary Gelen e-Fatura için iade oluştur
// @Description Gelen e-Fatura için iade oluştur.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Gelen e-Fatura için iade oluştur isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/:id/createreturninvoice [post]
func postIncomingInvoicesIdCreatereturninvoiceEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PostIncomingInvoicesIdCreatereturninvoiceEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTINCOMINGINVOICESIDCREATERETURNINVOICEEINVOICE", "Gelen e-Fatura için iade oluştur")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postIncomingInvoicesUuidUsernotesEinvoice godoc
// @Summary Kullanıcı notu ekler
// @Description Kullanıcı notu ekler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Kullanıcı notu ekler isteği"
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/:uuid/usernotes [post]
func postIncomingInvoicesUuidUsernotesEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	res, err := nesService.PostIncomingInvoicesUuidUsernotesEinvoice(&company, payload, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTINCOMINGINVOICESUUIDUSERNOTESEINVOICE", "Kullanıcı notu ekler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putIncomingInvoicesUuidUsernotesIdEinvoice godoc
// @Summary Kullanıcı notunu günceller
// @Description Kullanıcı notunu günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Kullanıcı notunu günceller isteği"
// @Param uuid path string true "uuid"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/:uuid/usernotes/:id [put]
func putIncomingInvoicesUuidUsernotesIdEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	id := c.Params("id")
	res, err := nesService.PutIncomingInvoicesUuidUsernotesIdEinvoice(&company, payload, uuid, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTINCOMINGINVOICESUUIDUSERNOTESIDEINVOICE", "Kullanıcı notunu günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteIncomingInvoicesUuidUsernotesIdEinvoice godoc
// @Summary Kullanıcı notunu siler
// @Description Kullanıcı notunu siler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "uuid"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/:uuid/usernotes/:id [delete]
func deleteIncomingInvoicesUuidUsernotesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	uuid := c.Params("uuid")
	id := c.Params("id")
	res, err := nesService.DeleteIncomingInvoicesUuidUsernotesIdEinvoice(&company, uuid, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEINCOMINGINVOICESUUIDUSERNOTESIDEINVOICE", "Kullanıcı notunu siler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postIncomingInvoicesEmailSendEinvoice godoc
// @Summary Belgeyi mail olarak iletir
// @Description Belgeyi mail olarak iletir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Belgeyi mail olarak iletir isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/incoming/invoices/email/send [post]
func postIncomingInvoicesEmailSendEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostIncomingInvoicesEmailSendEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTINCOMINGINVOICESEMAILSENDEINVOICE", "Belgeyi mail olarak iletir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesOutgoingEinvoice godoc
// @Summary Belgeleri listeler
// @Description Belgeleri listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/outgoing [get]
func getExinvoicesOutgoingEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetExinvoicesOutgoingEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESOUTGOINGEINVOICE", "Belgeleri listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesOutgoingUuidXmlEinvoice godoc
// @Summary XML İndir
// @Description XML İndir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/outgoing/:uuid/xml [get]
func getExinvoicesOutgoingUuidXmlEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	uuid := c.Params("uuid")
	res, err := nesService.GetExinvoicesOutgoingUuidXmlEinvoice(&company, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESOUTGOINGUUIDXMLEINVOICE", "XML İndir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesOutgoingUuidPdfEinvoice godoc
// @Summary PDF İndir
// @Description PDF İndir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/outgoing/:uuid/pdf [get]
func getExinvoicesOutgoingUuidPdfEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	uuid := c.Params("uuid")
	res, err := nesService.GetExinvoicesOutgoingUuidPdfEinvoice(&company, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESOUTGOINGUUIDPDFEINVOICE", "PDF İndir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getExinvoicesOutgoingUuidHtmlEinvoice godoc
// @Summary Belgeyi görüntüler
// @Description Belgeyi görüntüler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/outgoing/:uuid/html [get]
func getExinvoicesOutgoingUuidHtmlEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	uuid := c.Params("uuid")
	res, err := nesService.GetExinvoicesOutgoingUuidHtmlEinvoice(&company, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETEXINVOICESOUTGOINGUUIDHTMLEINVOICE", "Belgeyi görüntüler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postExinvoicesOutgoingExportFiletypeEinvoice godoc
// @Summary Dışarı Aktar
// @Description Dışarı Aktar.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Dışarı Aktar isteği"
// @Param fileType path string true "fileType"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/exinvoices/outgoing/export/:fileType [post]
func postExinvoicesOutgoingExportFiletypeEinvoice(c *fiber.Ctx) error {
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

	fileType := c.Params("fileType")
	res, err := nesService.PostExinvoicesOutgoingExportFiletypeEinvoice(&company, payload, fileType)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTEXINVOICESOUTGOINGEXPORTFILETYPEEINVOICE", "Dışarı Aktar")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getNotificationsOutgoingDynamicrulesEinvoice godoc
// @Summary Kuralları listeler
// @Description Kuralları listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/outgoing/dynamicrules [get]
func getNotificationsOutgoingDynamicrulesEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetNotificationsOutgoingDynamicrulesEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETNOTIFICATIONSOUTGOINGDYNAMICRULESEINVOICE", "Kuralları listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postNotificationsOutgoingDynamicrulesEinvoice godoc
// @Summary Kural oluşturur
// @Description Kural oluşturur.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Kural oluşturur isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/outgoing/dynamicrules [post]
func postNotificationsOutgoingDynamicrulesEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostNotificationsOutgoingDynamicrulesEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTNOTIFICATIONSOUTGOINGDYNAMICRULESEINVOICE", "Kural oluşturur")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getNotificationsOutgoingDynamicrulesIdEinvoice godoc
// @Summary Sorgulanan kuralı getirir
// @Description Sorgulanan kuralı getirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/outgoing/dynamicrules/:id [get]
func getNotificationsOutgoingDynamicrulesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetNotificationsOutgoingDynamicrulesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETNOTIFICATIONSOUTGOINGDYNAMICRULESIDEINVOICE", "Sorgulanan kuralı getirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putNotificationsOutgoingDynamicrulesIdEinvoice godoc
// @Summary Kural günceller
// @Description Kural günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Kural günceller isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/outgoing/dynamicrules/:id [put]
func putNotificationsOutgoingDynamicrulesIdEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PutNotificationsOutgoingDynamicrulesIdEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTNOTIFICATIONSOUTGOINGDYNAMICRULESIDEINVOICE", "Kural günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteNotificationsOutgoingDynamicrulesIdEinvoice godoc
// @Summary Kural siler
// @Description Kural siler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/notifications/outgoing/dynamicrules/:id [delete]
func deleteNotificationsOutgoingDynamicrulesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteNotificationsOutgoingDynamicrulesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETENOTIFICATIONSOUTGOINGDYNAMICRULESIDEINVOICE", "Kural siler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postOutgoingInvoicesExportFiletypeEinvoice godoc
// @Summary Toplu aktar
// @Description Toplu aktar.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Toplu aktar isteği"
// @Param fileType path string true "fileType"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/export/:fileType [post]
func postOutgoingInvoicesExportFiletypeEinvoice(c *fiber.Ctx) error {
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

	fileType := c.Params("fileType")
	res, err := nesService.PostOutgoingInvoicesExportFiletypeEinvoice(&company, payload, fileType)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTOUTGOINGINVOICESEXPORTFILETYPEEINVOICE", "Toplu aktar")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getOutgoingReportmoduleReportsEinvoice godoc
// @Summary Rapor listeler
// @Description Rapor listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/reportmodule/reports [get]
func getOutgoingReportmoduleReportsEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetOutgoingReportmoduleReportsEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETOUTGOINGREPORTMODULEREPORTSEINVOICE", "Rapor listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postOutgoingReportmoduleReportsEinvoice godoc
// @Summary Rapor oluşturur
// @Description Rapor oluşturur.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Rapor oluşturur isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/reportmodule/reports [post]
func postOutgoingReportmoduleReportsEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostOutgoingReportmoduleReportsEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTOUTGOINGREPORTMODULEREPORTSEINVOICE", "Rapor oluşturur")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getOutgoingReportmoduleReportsIdDownloadEinvoice godoc
// @Summary Rapor indirir
// @Description Rapor indirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/reportmodule/reports/:id/download [get]
func getOutgoingReportmoduleReportsIdDownloadEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetOutgoingReportmoduleReportsIdDownloadEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETOUTGOINGREPORTMODULEREPORTSIDDOWNLOADEINVOICE", "Rapor indirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getOutgoingReportmoduleTemplatesEinvoice godoc
// @Summary Şablonları listeler
// @Description Şablonları listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/reportmodule/templates [get]
func getOutgoingReportmoduleTemplatesEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetOutgoingReportmoduleTemplatesEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETOUTGOINGREPORTMODULETEMPLATESEINVOICE", "Şablonları listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postOutgoingReportmoduleTemplatesEinvoice godoc
// @Summary Rapor şablonu oluşturur
// @Description Rapor şablonu oluşturur.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Rapor şablonu oluşturur isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/reportmodule/templates [post]
func postOutgoingReportmoduleTemplatesEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostOutgoingReportmoduleTemplatesEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTOUTGOINGREPORTMODULETEMPLATESEINVOICE", "Rapor şablonu oluşturur")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getOutgoingReportmoduleTemplatesIdEinvoice godoc
// @Summary Sorgulanan şablonu getirir
// @Description Sorgulanan şablonu getirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/reportmodule/templates/:id [get]
func getOutgoingReportmoduleTemplatesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetOutgoingReportmoduleTemplatesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETOUTGOINGREPORTMODULETEMPLATESIDEINVOICE", "Sorgulanan şablonu getirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putOutgoingReportmoduleTemplatesIdEinvoice godoc
// @Summary Rapor şablonunu günceller
// @Description Rapor şablonunu günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Rapor şablonunu günceller isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/reportmodule/templates/:id [put]
func putOutgoingReportmoduleTemplatesIdEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PutOutgoingReportmoduleTemplatesIdEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTOUTGOINGREPORTMODULETEMPLATESIDEINVOICE", "Rapor şablonunu günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteOutgoingReportmoduleTemplatesIdEinvoice godoc
// @Summary Rapor Şablonunu siler
// @Description Rapor Şablonunu siler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/reportmodule/templates/:id [delete]
func deleteOutgoingReportmoduleTemplatesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteOutgoingReportmoduleTemplatesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEOUTGOINGREPORTMODULETEMPLATESIDEINVOICE", "Rapor Şablonunu siler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getOutgoingReportmoduleColumnsEinvoice godoc
// @Summary Kolonları listeler
// @Description Kolonları listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/reportmodule/columns [get]
func getOutgoingReportmoduleColumnsEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetOutgoingReportmoduleColumnsEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETOUTGOINGREPORTMODULECOLUMNSEINVOICE", "Kolonları listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putOutgoingInvoicesTagsEinvoice godoc
// @Summary Etiket ekler/çıkarır
// @Description Etiket ekler/çıkarır.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Etiket ekler/çıkarır isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/tags [put]
func putOutgoingInvoicesTagsEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PutOutgoingInvoicesTagsEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTOUTGOINGINVOICESTAGSEINVOICE", "Etiket ekler/çıkarır")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putOutgoingInvoicesUuidReceiveraliasEinvoice godoc
// @Summary Taslak belgelerin alıcı etiketi bu uç ile güncellenebilir
// @Description Taslak belgelerin alıcı etiketi bu uç ile güncellenebilir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Taslak belgelerin alıcı etiketi bu uç ile güncellenebilir isteği"
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/:uuid/receiveralias [put]
func putOutgoingInvoicesUuidReceiveraliasEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	res, err := nesService.PutOutgoingInvoicesUuidReceiveraliasEinvoice(&company, payload, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTOUTGOINGINVOICESUUIDRECEIVERALIASEINVOICE", "Taslak belgelerin alıcı etiketi bu uç ile güncellenebilir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postOutgoingInvoicesUuidSavecompanyindocumentEinvoice godoc
// @Summary Firma olarak kaydet
// @Description Firma olarak kaydet.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Firma olarak kaydet isteği"
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/:uuid/savecompanyindocument [post]
func postOutgoingInvoicesUuidSavecompanyindocumentEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	res, err := nesService.PostOutgoingInvoicesUuidSavecompanyindocumentEinvoice(&company, payload, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTOUTGOINGINVOICESUUIDSAVECOMPANYINDOCUMENTEINVOICE", "Firma olarak kaydet")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putOutgoingInvoicesBulkOperationEinvoice godoc
// @Summary Yeni durum atar
// @Description Yeni durum atar.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Yeni durum atar isteği"
// @Param operation path string true "operation"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/bulk/:operation [put]
func putOutgoingInvoicesBulkOperationEinvoice(c *fiber.Ctx) error {
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

	operation := c.Params("operation")
	res, err := nesService.PutOutgoingInvoicesBulkOperationEinvoice(&company, payload, operation)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTOUTGOINGINVOICESBULKOPERATIONEINVOICE", "Yeni durum atar")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postUploadsResendUuidEinvoice godoc
// @Summary Hata almış bir belgeyi aynen yeniden gönderir
// @Description Hata almış bir belgeyi aynen yeniden gönderir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Hata almış bir belgeyi aynen yeniden gönderir isteği"
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/uploads/resend/:uuid [post]
func postUploadsResendUuidEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	res, err := nesService.PostUploadsResendUuidEinvoice(&company, payload, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTUPLOADSRESENDUUIDEINVOICE", "Hata almış bir belgeyi aynen yeniden gönderir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postOutgoingInvoicesUuidUsernotesEinvoice godoc
// @Summary Kullanıcı notu ekler
// @Description Kullanıcı notu ekler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Kullanıcı notu ekler isteği"
// @Param uuid path string true "uuid"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/:uuid/usernotes [post]
func postOutgoingInvoicesUuidUsernotesEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	res, err := nesService.PostOutgoingInvoicesUuidUsernotesEinvoice(&company, payload, uuid)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTOUTGOINGINVOICESUUIDUSERNOTESEINVOICE", "Kullanıcı notu ekler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putOutgoingInvoicesUuidUsernotesIdEinvoice godoc
// @Summary Kullanıcı notunu günceller
// @Description Kullanıcı notunu günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Kullanıcı notunu günceller isteği"
// @Param uuid path string true "uuid"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/:uuid/usernotes/:id [put]
func putOutgoingInvoicesUuidUsernotesIdEinvoice(c *fiber.Ctx) error {
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

	uuid := c.Params("uuid")
	id := c.Params("id")
	res, err := nesService.PutOutgoingInvoicesUuidUsernotesIdEinvoice(&company, payload, uuid, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTOUTGOINGINVOICESUUIDUSERNOTESIDEINVOICE", "Kullanıcı notunu günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteOutgoingInvoicesUuidUsernotesIdEinvoice godoc
// @Summary Kullanıcı notunu siler
// @Description Kullanıcı notunu siler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "uuid"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/:uuid/usernotes/:id [delete]
func deleteOutgoingInvoicesUuidUsernotesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	uuid := c.Params("uuid")
	id := c.Params("id")
	res, err := nesService.DeleteOutgoingInvoicesUuidUsernotesIdEinvoice(&company, uuid, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEOUTGOINGINVOICESUUIDUSERNOTESIDEINVOICE", "Kullanıcı notunu siler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteOutgoingInvoicesDraftsEinvoice godoc
// @Summary Taslak belgeleri silmek için bu uç kullanılablir
// @Description Taslak belgeleri silmek için bu uç kullanılablir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/drafts [delete]
func deleteOutgoingInvoicesDraftsEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.DeleteOutgoingInvoicesDraftsEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEOUTGOINGINVOICESDRAFTSEINVOICE", "Taslak belgeleri silmek için bu uç kullanılablir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postOutgoingInvoicesEmailSendEinvoice godoc
// @Summary Belgeyi mail olarak iletir
// @Description Belgeyi mail olarak iletir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Belgeyi mail olarak iletir isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/outgoing/invoices/email/send [post]
func postOutgoingInvoicesEmailSendEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostOutgoingInvoicesEmailSendEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTOUTGOINGINVOICESEMAILSENDEINVOICE", "Belgeyi mail olarak iletir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getUsersZipAliastypeEinvoice godoc
// @Summary Mükellef listesini indirir
// @Description Mükellef listesini indirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param aliasType path string true "aliasType"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/users/zip/:aliasType [get]
func getUsersZipAliastypeEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	aliasType := c.Params("aliasType")
	res, err := nesService.GetUsersZipAliastypeEinvoice(&company, aliasType)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETUSERSZIPALIASTYPEEINVOICE", "Mükellef listesini indirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getUsersIdentifierAliastypeEinvoice godoc
// @Summary Kimlik No ile sorgular
// @Description Kimlik No ile sorgular.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param identifier path string true "identifier"
// @Param aliasType path string true "aliasType"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/users/:identifier/:aliasType [get]
func getUsersIdentifierAliastypeEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	identifier := c.Params("identifier")
	aliasType := c.Params("aliasType")
	res, err := nesService.GetUsersIdentifierAliastypeEinvoice(&company, identifier, aliasType)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETUSERSIDENTIFIERALIASTYPEEINVOICE", "Kimlik No ile sorgular")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postUsersAliastypeEinvoice godoc
// @Summary Kimlik No ile sorgular
// @Description Kimlik No ile sorgular.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Kimlik No ile sorgular isteği"
// @Param aliasType path string true "aliasType"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/users/:aliasType [post]
func postUsersAliastypeEinvoice(c *fiber.Ctx) error {
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

	aliasType := c.Params("aliasType")
	res, err := nesService.PostUsersAliastypeEinvoice(&company, payload, aliasType)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTUSERSALIASTYPEEINVOICE", "Kimlik No ile sorgular")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getUsersSearchQueryAliastypeEinvoice godoc
// @Summary Ünvan ile sorgular
// @Description Ünvan ile sorgular.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query path string true "query"
// @Param aliasType path string true "aliasType"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/users/search/:query/:aliasType [get]
func getUsersSearchQueryAliastypeEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	query := c.Params("query")
	aliasType := c.Params("aliasType")
	res, err := nesService.GetUsersSearchQueryAliastypeEinvoice(&company, query, aliasType)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETUSERSSEARCHQUERYALIASTYPEEINVOICE", "Ünvan ile sorgular")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsDocumenttemplatesCustomizationsettingsEinvoice godoc
// @Summary Tasarım ayarları dönülür
// @Description Tasarım ayarları dönülür.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings [get]
func getDefinitionsDocumenttemplatesCustomizationsettingsEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetDefinitionsDocumenttemplatesCustomizationsettingsEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSEINVOICE", "Tasarım ayarları dönülür")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postDefinitionsDocumenttemplatesCustomizationsettingsEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarım eklemek için kullanılır.
// @Description e-Belge özelleştirilebilir tasarım eklemek için kullanılır..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "e-Belge özelleştirilebilir tasarım eklemek için kullanılır. isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings [post]
func postDefinitionsDocumenttemplatesCustomizationsettingsEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostDefinitionsDocumenttemplatesCustomizationsettingsEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSEINVOICE", "e-Belge özelleştirilebilir tasarım eklemek için kullanılır.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice godoc
// @Summary Sorgulanan ayarı getirir
// @Description Sorgulanan ayarı getirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id [get]
func getDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDEINVOICE", "Sorgulanan ayarı getirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarımını güncellemek için kullanılır.
// @Description e-Belge özelleştirilebilir tasarımını güncellemek için kullanılır..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "e-Belge özelleştirilebilir tasarımını güncellemek için kullanılır. isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id [put]
func putDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PutDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDEINVOICE", "e-Belge özelleştirilebilir tasarımını güncellemek için kullanılır.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarımını silmek için kullanılır.
// @Description e-Belge özelleştirilebilir tasarımını silmek için kullanılır..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id [delete]
func deleteDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDEINVOICE", "e-Belge özelleştirilebilir tasarımını silmek için kullanılır.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsDocumenttemplatesCustomizationsettingsIdSetdefaultEinvoice godoc
// @Summary No description provided
// @Description No description provided.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/setdefault [get]
func getDefinitionsDocumenttemplatesCustomizationsettingsIdSetdefaultEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetDefinitionsDocumenttemplatesCustomizationsettingsIdSetdefaultEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDSETDEFAULTEINVOICE", "No description provided")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postDefinitionsDocumenttemplatesCustomizationsettingsIdPreviewEinvoice godoc
// @Summary Tasarımı önizler
// @Description Tasarımı önizler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Tasarımı önizler isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/preview [post]
func postDefinitionsDocumenttemplatesCustomizationsettingsIdPreviewEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PostDefinitionsDocumenttemplatesCustomizationsettingsIdPreviewEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDPREVIEWEINVOICE", "Tasarımı önizler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan logoya bu uç ile ulaşılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan logoya bu uç ile ulaşılabilir..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/logo [get]
func getDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDLOGOEINVOICE", "e-Belge özelleştirilebilir tasarıma eklenmiş olan logoya bu uç ile ulaşılabilir.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarıma logo eklemek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma logo eklemek için bu uç kullanılabilir..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "e-Belge özelleştirilebilir tasarıma logo eklemek için bu uç kullanılabilir. isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/logo [post]
func postDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PostDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDLOGOEINVOICE", "e-Belge özelleştirilebilir tasarıma logo eklemek için bu uç kullanılabilir.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan logoyu silmek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan logoyu silmek için bu uç kullanılabilir..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/logo [delete]
func deleteDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDLOGOEINVOICE", "e-Belge özelleştirilebilir tasarıma eklenmiş olan logoyu silmek için bu uç kullanılabilir.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeye bu uç ile ulaşılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeye bu uç ile ulaşılabilir..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/stamp [get]
func getDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDSTAMPEINVOICE", "e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeye bu uç ile ulaşılabilir.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarıma kaşe eklemek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma kaşe eklemek için bu uç kullanılabilir..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "e-Belge özelleştirilebilir tasarıma kaşe eklemek için bu uç kullanılabilir. isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/stamp [post]
func postDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PostDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDSTAMPEINVOICE", "e-Belge özelleştirilebilir tasarıma kaşe eklemek için bu uç kullanılabilir.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeyi silmek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeyi silmek için bu uç kullanılabilir..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/stamp [delete]
func deleteDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDSTAMPEINVOICE", "e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeyi silmek için bu uç kullanılabilir.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan imzaya bu uç ile ulaşılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan imzaya bu uç ile ulaşılabilir..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/signature [get]
func getDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDSIGNATUREEINVOICE", "e-Belge özelleştirilebilir tasarıma eklenmiş olan imzaya bu uç ile ulaşılabilir.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarıma imza eklemek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma imza eklemek için bu uç kullanılabilir..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "e-Belge özelleştirilebilir tasarıma imza eklemek için bu uç kullanılabilir. isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/signature [post]
func postDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PostDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDSIGNATUREEINVOICE", "e-Belge özelleştirilebilir tasarıma imza eklemek için bu uç kullanılabilir.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice godoc
// @Summary e-Belge özelleştirilebilir tasarıma eklenmiş olan imzayı silmek için bu uç kullanılabilir.
// @Description e-Belge özelleştirilebilir tasarıma eklenmiş olan imzayı silmek için bu uç kullanılabilir..
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/customizationsettings/:id/signature [delete]
func deleteDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEDEFINITIONSDOCUMENTTEMPLATESCUSTOMIZATIONSETTINGSIDSIGNATUREEINVOICE", "e-Belge özelleştirilebilir tasarıma eklenmiş olan imzayı silmek için bu uç kullanılabilir.")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsSeriesEinvoice godoc
// @Summary Serileri listeler
// @Description Serileri listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/series [get]
func getDefinitionsSeriesEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetDefinitionsSeriesEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSSERIESEINVOICE", "Serileri listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postDefinitionsSeriesEinvoice godoc
// @Summary Seri ekler
// @Description Seri ekler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Seri ekler isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/series [post]
func postDefinitionsSeriesEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostDefinitionsSeriesEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTDEFINITIONSSERIESEINVOICE", "Seri ekler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsSeriesIdEinvoice godoc
// @Summary Sorgulanan seriyi getirir
// @Description Sorgulanan seriyi getirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/series/:id [get]
func getDefinitionsSeriesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetDefinitionsSeriesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSSERIESIDEINVOICE", "Sorgulanan seriyi getirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteDefinitionsSeriesIdEinvoice godoc
// @Summary Seri siler
// @Description Seri siler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/series/:id [delete]
func deleteDefinitionsSeriesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteDefinitionsSeriesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEDEFINITIONSSERIESIDEINVOICE", "Seri siler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsSeriesSerieEinvoice godoc
// @Summary Ön eke göre seriyi getirir
// @Description Ön eke göre seriyi getirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param serie path string true "serie"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/series/:serie [get]
func getDefinitionsSeriesSerieEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	serie := c.Params("serie")
	res, err := nesService.GetDefinitionsSeriesSerieEinvoice(&company, serie)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSSERIESSERIEEINVOICE", "Ön eke göre seriyi getirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsSeriesIdSetStatusEinvoice godoc
// @Summary Seri durumunu günceller
// @Description Seri durumunu günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Param status path string true "status"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/series/:id/set/:status [get]
func getDefinitionsSeriesIdSetStatusEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	status := c.Params("status")
	res, err := nesService.GetDefinitionsSeriesIdSetStatusEinvoice(&company, id, status)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSSERIESIDSETSTATUSEINVOICE", "Seri durumunu günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsSeriesIdSetdefaultEinvoice godoc
// @Summary Seriyi varsayılan ayarlar
// @Description Seriyi varsayılan ayarlar.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/series/:id/setdefault [get]
func getDefinitionsSeriesIdSetdefaultEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetDefinitionsSeriesIdSetdefaultEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSSERIESIDSETDEFAULTEINVOICE", "Seriyi varsayılan ayarlar")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsSeriesIdYearSetnumberNextnumberEinvoice godoc
// @Summary Sayaç günceller
// @Description Sayaç günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Param year path string true "year"
// @Param nextNumber path string true "nextNumber"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/series/:id/:year/setnumber/:nextNumber [get]
func getDefinitionsSeriesIdYearSetnumberNextnumberEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	year := c.Params("year")
	nextNumber := c.Params("nextNumber")
	res, err := nesService.GetDefinitionsSeriesIdYearSetnumberNextnumberEinvoice(&company, id, year, nextNumber)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSSERIESIDYEARSETNUMBERNEXTNUMBEREINVOICE", "Sayaç günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsSeriesSerieidYearHistoriesEinvoice godoc
// @Summary Sayaç geçmişi
// @Description Sayaç geçmişi.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param serieId path string true "serieId"
// @Param year path string true "year"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/series/:serieId/:year/histories [get]
func getDefinitionsSeriesSerieidYearHistoriesEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	serieId := c.Params("serieId")
	year := c.Params("year")
	res, err := nesService.GetDefinitionsSeriesSerieidYearHistoriesEinvoice(&company, serieId, year)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSSERIESSERIEIDYEARHISTORIESEINVOICE", "Sayaç geçmişi")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsDocumenttemplatesEinvoice godoc
// @Summary Tasarımları listeler
// @Description Tasarımları listeler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates [get]
func getDefinitionsDocumenttemplatesEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	res, err := nesService.GetDefinitionsDocumenttemplatesEinvoice(&company)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSDOCUMENTTEMPLATESEINVOICE", "Tasarımları listeler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postDefinitionsDocumenttemplatesEinvoice godoc
// @Summary Tasarım ekler
// @Description Tasarım ekler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Tasarım ekler isteği"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates [post]
func postDefinitionsDocumenttemplatesEinvoice(c *fiber.Ctx) error {
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

	res, err := nesService.PostDefinitionsDocumenttemplatesEinvoice(&company, payload)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTDEFINITIONSDOCUMENTTEMPLATESEINVOICE", "Tasarım ekler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsDocumenttemplatesIdEinvoice godoc
// @Summary Tasarım dosyasını indirir
// @Description Tasarım dosyasını indirir.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/:id [get]
func getDefinitionsDocumenttemplatesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetDefinitionsDocumenttemplatesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSDOCUMENTTEMPLATESIDEINVOICE", "Tasarım dosyasını indirir")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// putDefinitionsDocumenttemplatesIdEinvoice godoc
// @Summary Tasarımı günceller
// @Description Tasarımı günceller.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Tasarımı günceller isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/:id [put]
func putDefinitionsDocumenttemplatesIdEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PutDefinitionsDocumenttemplatesIdEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "PUTDEFINITIONSDOCUMENTTEMPLATESIDEINVOICE", "Tasarımı günceller")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// deleteDefinitionsDocumenttemplatesIdEinvoice godoc
// @Summary Tasarımı siler
// @Description Tasarımı siler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/:id [delete]
func deleteDefinitionsDocumenttemplatesIdEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.DeleteDefinitionsDocumenttemplatesIdEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "DELETEDEFINITIONSDOCUMENTTEMPLATESIDEINVOICE", "Tasarımı siler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getDefinitionsDocumenttemplatesIdSetdefaultEinvoice godoc
// @Summary Tasarımı varsayılan ayarlar
// @Description Tasarımı varsayılan ayarlar.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/:id/setdefault [get]
func getDefinitionsDocumenttemplatesIdSetdefaultEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	id := c.Params("id")
	res, err := nesService.GetDefinitionsDocumenttemplatesIdSetdefaultEinvoice(&company, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETDEFINITIONSDOCUMENTTEMPLATESIDSETDEFAULTEINVOICE", "Tasarımı varsayılan ayarlar")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// postDefinitionsDocumenttemplatesIdPreviewEinvoice godoc
// @Summary Tasarımı önizler
// @Description Tasarımı önizler.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Tasarımı önizler isteği"
// @Param id path string true "id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/definitions/documenttemplates/:id/preview [post]
func postDefinitionsDocumenttemplatesIdPreviewEinvoice(c *fiber.Ctx) error {
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

	id := c.Params("id")
	res, err := nesService.PostDefinitionsDocumenttemplatesIdPreviewEinvoice(&company, payload, id)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "POSTDEFINITIONSDOCUMENTTEMPLATESIDPREVIEWEINVOICE", "Tasarımı önizler")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}


// getEnvelopesInstanceidentifierQueryEinvoice godoc
// @Summary Zarf Durum Sorgular
// @Description Zarf Durum Sorgular.
// @Tags EInvoice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param instanceIdentifier path string true "instanceIdentifier"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errors.AppError
// @Failure 401 {object} errors.AppError
// @Router /v1/einvoice/envelopes/:instanceIdentifier/query [get]
func getEnvelopesInstanceidentifierQueryEinvoice(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user.CompanyID == nil {
		return errors.NewError(fiber.StatusUnauthorized, errors.ErrCodeUnauthorized, "Kullanıcı firma bilgisi bulunamadı")
	}

	var company models.Company
	if err := db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID); err != nil {
		return errors.NewError(fiber.StatusInternalServerError, errors.ErrCodeDatabaseError, "Firma bilgisi alınamadı")
	}

	instanceIdentifier := c.Params("instanceIdentifier")
	res, err := nesService.GetEnvelopesInstanceidentifierQueryEinvoice(&company, instanceIdentifier)
	if err != nil {
		return err
	}

	// Log successful operation
	db.DB.Exec("INSERT INTO system_logs (user_id, action, details) VALUES ($1, $2, $3)", user.ID, "GETENVELOPESINSTANCEIDENTIFIERQUERYEINVOICE", "Zarf Durum Sorgular")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": res})
}
