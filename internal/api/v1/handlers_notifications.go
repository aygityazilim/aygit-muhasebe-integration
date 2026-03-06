package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"aygit-muhasebe-integration/internal/service"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Get Dynamic Rules
// @Description  Kuralları listeler
// @Tags         Notifications
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/notifications/dynamicrules [get]
func (h *VoucherHandler) GetDynamicRules(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/notifications/dynamicrules")
}

// @Summary      Create Dynamic Rule
// @Description  Kural oluşturur
// @Tags         Notifications
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Rule Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/notifications/dynamicrules [post]
func (h *VoucherHandler) CreateDynamicRule(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/notifications/dynamicrules")
}

// @Summary      Get Dynamic Rule by ID
// @Description  Sorgulanan kuralı getirir
// @Tags         Notifications
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "Rule ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/notifications/dynamicrules/{id} [get]
func (h *VoucherHandler) GetDynamicRuleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/notifications/dynamicrules/%s", id))
}

// @Summary      Update Dynamic Rule
// @Description  Kural günceller
// @Tags         Notifications
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "Rule ID"
// @Param        payload body interface{} true "Rule Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/notifications/dynamicrules/{id} [put]
func (h *VoucherHandler) UpdateDynamicRule(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "PUT", fmt.Sprintf("/v1/notifications/dynamicrules/%s", id))
}

// @Summary      Delete Dynamic Rule
// @Description  Kural siler
// @Tags         Notifications
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "Rule ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/notifications/dynamicrules/{id} [delete]
func (h *VoucherHandler) DeleteDynamicRule(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/notifications/dynamicrules/%s", id))
}

// helper for forwarding requests
func (h *VoucherHandler) forwardRequest(c *fiber.Ctx, method, endpoint string) error {
	vType := c.Params("type")
	company := getCompanyFromContext(c)

	var svc *service.NESVoucherService
	if vType == "emm" {
		svc = h.EMMService
	} else if vType == "esmm" {
		svc = h.ESMMService
	} else {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid voucher type"})
	}

	contentType := c.Get("Content-Type")
	payload := c.Body()

	result, err := svc.Forward(method, endpoint, company, contentType, payload)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Try to parse as JSON if it's JSON
	var jsonResult interface{}
	if err := json.Unmarshal(result, &jsonResult); err == nil {
		return c.JSON(jsonResult)
	}

	// Otherwise send raw
	c.Set("Content-Type", "application/json") // defaulting, can be changed if needed
	return c.Send(result)
}
