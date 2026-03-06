package v1

import (
	"encoding/json"
	"net/http"

	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/internal/service"

	"github.com/gofiber/fiber/v2"
)

// VoucherHandler handles EMM and ESMM operations
type VoucherHandler struct {
	EMMService  *service.NESVoucherService
	ESMMService *service.NESVoucherService
}

func NewVoucherHandler() *VoucherHandler {
	// The specific URLs would ideally come from configuration
	return &VoucherHandler{
		EMMService:  service.NewNESVoucherService("https://apitest.nes.com.tr/emm"),
		ESMMService: service.NewNESVoucherService("https://apitest.nes.com.tr/esmm"),
	}
}

// @Summary      Get Vouchers (Listeleme)
// @Description  Taslak veya onaylı makbuzları listeler
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{} "Makbuz listesi"
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers [get]
func (h *VoucherHandler) GetVouchers(c *fiber.Ctx) error {
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

	result, err := svc.GetVouchers(company)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var jsonResult map[string]interface{}
	json.Unmarshal([]byte(result), &jsonResult)
	return c.JSON(jsonResult)
}

// @Summary      Get Draft Vouchers (Taslak Makbuzları Listeleme)
// @Description  Taslak makbuzları listeler
// @Tags         Vouchers
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/vouchers/drafts [get]
func (h *VoucherHandler) GetDraftVouchers(c *fiber.Ctx) error {
	vType := c.Params("type")
	company := getCompanyFromContext(c)

	var svc *service.NESVoucherService
	if vType == "emm" {
		svc = h.EMMService
	} else {
		svc = h.ESMMService
	}

	result, err := svc.GetDraftVouchers(company)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var jsonResult map[string]interface{}
	json.Unmarshal([]byte(result), &jsonResult)
	return c.JSON(jsonResult)
}

// @Summary      Upload Document (Belge Yükleme)
// @Description  Yeni bir E-MM veya E-SMM makbuz belgesi yükler
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body interface{} true "Belge verisi"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/uploads/document [post]
func (h *VoucherHandler) UploadDocument(c *fiber.Ctx) error {
	vType := c.Params("type")
	company := getCompanyFromContext(c)

	var svc *service.NESVoucherService
	if vType == "emm" {
		svc = h.EMMService
	} else {
		svc = h.ESMMService
	}

	result, err := svc.UploadDocument(company, c.Body())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var jsonResult map[string]interface{}
	json.Unmarshal([]byte(result), &jsonResult)
	return c.JSON(jsonResult)
}

// @Summary      Send Draft (Taslak Gönderme)
// @Description  Taslak halindeki makbuzu resmileştirir
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        type path string true "Voucher Type: emm or esmm"
// @Param        payload body []string true "Taslak UUID listesi"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/{type}/uploads/draft/send [post]
func (h *VoucherHandler) SendDraft(c *fiber.Ctx) error {
	vType := c.Params("type")
	company := getCompanyFromContext(c)

	var uuids []string
	if err := c.BodyParser(&uuids); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	var svc *service.NESVoucherService
	if vType == "emm" {
		svc = h.EMMService
	} else {
		svc = h.ESMMService
	}

	result, err := svc.SendDraftVoucher(company, uuids)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var jsonResult map[string]interface{}
	json.Unmarshal([]byte(result), &jsonResult)
	return c.JSON(jsonResult)
}

// Dummy helper just for completeness to simulate auth extraction
func getCompanyFromContext(c *fiber.Ctx) *models.Company {
	// Implementation should extract company from context / JWT
	// Returning a mock company to make it compile properly for demonstration
	dummyAPIKey := "dummy_api_key"
	return &models.Company{
		Environment: "TEST",
		NesAPIKey:   &dummyAPIKey,
	}
}

func (h *VoucherHandler) passthrough(c *fiber.Ctx) error {
	return h.forwardRequest(c, c.Method(), c.OriginalURL())
}
