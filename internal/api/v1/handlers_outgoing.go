package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// --- Report Module Reports ---

// @Summary      Get Reports
// @Description  Rapor listeler
// @Tags         Outgoing
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/outgoing/reportmodule/reports [get]
func (h *VoucherHandler) GetReports(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/outgoing/reportmodule/reports")
}

// @Summary      Create Report
// @Description  Rapor oluşturur
// @Tags         Outgoing
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/outgoing/reportmodule/reports [post]
func (h *VoucherHandler) CreateReport(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/outgoing/reportmodule/reports")
}

// @Summary      Download Report
// @Description  Rapor indirir
// @Tags         Outgoing
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/outgoing/reportmodule/reports/{id}/download [get]
func (h *VoucherHandler) DownloadReport(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/outgoing/reportmodule/reports/%s/download", id))
}

// --- Report Module Templates ---

// @Summary      Get Report Templates
// @Description  Şablonları listeler
// @Tags         Outgoing
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/outgoing/reportmodule/templates [get]
func (h *VoucherHandler) GetReportTemplates(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/outgoing/reportmodule/templates")
}

// @Summary      Create Report Template
// @Description  Rapor şablonu oluşturur
// @Tags         Outgoing
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/outgoing/reportmodule/templates [post]
func (h *VoucherHandler) CreateReportTemplate(c *fiber.Ctx) error {
	return h.forwardRequest(c, "POST", "/v1/outgoing/reportmodule/templates")
}

// @Summary      Get Report Template By ID
// @Description  Sorgulanan şablonu getirir
// @Tags         Outgoing
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/outgoing/reportmodule/templates/{id} [get]
func (h *VoucherHandler) GetReportTemplateByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "GET", fmt.Sprintf("/v1/outgoing/reportmodule/templates/%s", id))
}

// @Summary      Update Report Template
// @Description  Rapor şablonunu günceller
// @Tags         Outgoing
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Param        payload body interface{} true "Data"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/outgoing/reportmodule/templates/{id} [put]
func (h *VoucherHandler) UpdateReportTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "PUT", fmt.Sprintf("/v1/outgoing/reportmodule/templates/%s", id))
}

// @Summary      Delete Report Template
// @Description  Rapor Şablonunu siler
// @Tags         Outgoing
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        id path string true "ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/outgoing/reportmodule/templates/{id} [delete]
func (h *VoucherHandler) DeleteReportTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	return h.forwardRequest(c, "DELETE", fmt.Sprintf("/v1/outgoing/reportmodule/templates/%s", id))
}

// --- Report Module Columns ---

// @Summary      Get Columns
// @Description  Kolonları listeler
// @Tags         Outgoing
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/outgoing/reportmodule/columns [get]
func (h *VoucherHandler) GetReportColumns(c *fiber.Ctx) error {
	return h.forwardRequest(c, "GET", "/v1/outgoing/reportmodule/columns")
}
