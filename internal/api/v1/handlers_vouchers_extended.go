package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// --- Extensions for Vouchers ---

// @Summary      Export Vouchers
// @Description  Toplu aktar
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        fileType path string true "File Type"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/export/{fileType} [post]
func (h *VoucherHandler) ExportVouchers(c *fiber.Ctx) error {
	fileType := c.Params("fileType")
	return h.forwardRequest(c, "POST", fmt.Sprintf("/v1/vouchers/export/%s", fileType))
}

// @Summary      Add User Note
// @Description  Kullanıcı notu ekler
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        uuid path string true "UUID"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/{uuid}/usernotes [post]
func (h *VoucherHandler) AddUserNoteForward(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	return h.forwardRequest(c, "POST", fmt.Sprintf("/v1/vouchers/%s/usernotes", uuid))
}

// @Summary      Update User Note
// @Description  Kullanıcı notunu günceller
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        uuid path string true "UUID"
// @Param        id path string true "Note ID"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/{uuid}/usernotes/{id} [put]
func (h *VoucherHandler) UpdateUserNote(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	id := c.Params("id")
	return h.forwardRequest(c, "PUT", fmt.Sprintf("/v1/vouchers/%s/usernotes/%s", uuid, id))
}

// @Summary      Delete User Note
// @Description  Kullanıcı notunu siler
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        uuid path string true "UUID"
// @Param        id path string true "Note ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/{uuid}/usernotes/{id} [delete]
func (h *VoucherHandler) DeleteUserNote(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/vouchers/%s/usernotes/%s", uuid, id))
}

// @Summary      Cancel Voucher
// @Description  Belgeyi iptal eder
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/cancel [post]
func (h *VoucherHandler) CancelVoucherForward(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/vouchers/cancel")
}

// @Summary      Withdraw Cancel Voucher
// @Description  İptal işlemini geri alır
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/canceled/withdraw [post]
func (h *VoucherHandler) WithdrawCancelVoucherForward(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/vouchers/canceled/withdraw")
}

// @Summary      Add/Remove Tags
// @Description  Etiket ekler/çıkarır
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/tags [put]
func (h *VoucherHandler) UpdateVouchersTags(c *fiber.Ctx) error {
	return h.forwardRequest(c, "PUT", "/v1/vouchers/tags")
}

// @Summary      Save Company In Document
// @Description  Firma olarak kaydet
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        uuid path string true "UUID"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/{uuid}/savecompanyindocument [post]
func (h *VoucherHandler) SaveCompanyInDocument(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	return h.forwardRequest(c, "POST", fmt.Sprintf("/v1/vouchers/%s/savecompanyindocument", uuid))
}

// @Summary      Bulk Operations
// @Description  Yeni durum atar
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        operation path string true "Operation"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/bulk/{operation} [put]
func (h *VoucherHandler) BulkVouchersOperation(c *fiber.Ctx) error {
	op := c.Params("operation")
	return h.forwardRequest(c, "PUT", fmt.Sprintf("/v1/vouchers/bulk/%s", op))
}

// @Summary      Delete Drafts
// @Description  Taslak belgeleri silmek için bu uç kullanılablir
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/drafts [delete]
func (h *VoucherHandler) DeleteDrafts(c *fiber.Ctx) error {
	return h.forwardRequest(c, "DELETE", "/v1/vouchers/drafts")
}

// @Summary      Send Voucher Email
// @Description  Belgeyi mail olarak iletir
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/email/send [post]
func (h *VoucherHandler) SendVoucherEmailForward(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/vouchers/email/send")
}
