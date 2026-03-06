package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Update Document Upload
// @Description  Belge günceller
// @Tags         Uploads
// @Accept       multipart/form-data
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        uuid path string true "UUID"
// @Param        file formData file true "Document File"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/uploads/document/{uuid} [put]
func (h *VoucherHandler) UpdateDocumentUpload(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	return h.forwardRequest(c, "PUT", fmt.Sprintf("/v1/uploads/document/%s", uuid))
}

// @Summary      Create Draft Upload
// @Description  Taslak Belge yükler
// @Tags         Uploads
// @Accept       multipart/form-data
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        file formData file true "Document File"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/uploads/draft/create/{id} [post]
func (h *VoucherHandler) CreateDraftUpload(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "POST", fmt.Sprintf("/v1/uploads/draft/create/%s", id))
}
