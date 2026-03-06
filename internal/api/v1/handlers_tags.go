package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Get Tags
// @Description  Etiketleri listeler
// @Tags         Tags
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/tags [get]
func (h *VoucherHandler) GetTags(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/tags")
}

// @Summary      Create Tag
// @Description  Etiket ekler
// @Tags         Tags
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/tags [post]
func (h *VoucherHandler) CreateTag(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/tags")
}

// @Summary      Get Tag By ID
// @Description  Sorgulanan etiketi getirir
// @Tags         Tags
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/tags/{id} [get]
func (h *VoucherHandler) GetTagByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/tags/%s", id))
}

// @Summary      Update Tag
// @Description  Etiket günceller
// @Tags         Tags
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/tags/{id} [put]
func (h *VoucherHandler) UpdateTag(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "PUT", fmt.Sprintf("/v1/tags/%s", id))
}

// @Summary      Delete Tag
// @Description  Etiket siler
// @Tags         Tags
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/tags/{id} [delete]
func (h *VoucherHandler) DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/tags/%s", id))
}
