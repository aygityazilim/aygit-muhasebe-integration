package v1

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/db"

	"github.com/gofiber/fiber/v2"
)

// QueryEnvelopeStatus godoc
// @Summary Zarf Durum Sorgular
// @Tags Envelopes
// @Security BearerAuth
// @Param instanceIdentifier path string true "Instance Identifier"
// @Success 200 {object} map[string]interface{}
// @Router /v1/envelopes/{instanceIdentifier}/query [get]
func QueryEnvelopeStatus(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	var company models.Company
	db.DB.Get(&company, "SELECT * FROM companies WHERE id=$1", *user.CompanyID)

	res, err := nesService.QueryEnvelopeStatus(&company, c.Params("instanceIdentifier"))
	if err != nil {
		return err
	}
	return c.JSON(res)
}
