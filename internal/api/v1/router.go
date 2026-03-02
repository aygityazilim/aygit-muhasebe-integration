package v1

import (
	_ "aygit-muhasebe-integration/docs"
	"aygit-muhasebe-integration/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// SetupRoutes registers all version 1 routes and Swagger UI
func SetupRoutes(app *fiber.App) {
	// Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	v1 := app.Group("/v1", middleware.AuthRequired)

	// Management routes
	management := v1.Group("/management")
	management.Get("/creditsummary", getCreditSummary)
	// Identity routes
	management.Get("/identifications", GetIdentifications)
	management.Post("/identifications", createIdentification)

	// StaticCodes routes
	staticcodes := v1.Group("/staticcodes")
	staticcodes.Get("/taxtype", getTaxTypes)
	staticcodes.Get("/withholdingtaxtype", getWithholdingTaxTypes)
	staticcodes.Get("/taxexemptionreason", getTaxExemptionReasons)
	staticcodes.Get("/currency", getCurrencies)

	// Statistics routes
	statistics := v1.Group("/statistics")
	statistics.Get("/daily", getDailyStatistics)

	// Invoices routes (E-Fatura)
	invoices := v1.Group("/invoices")
	invoices.Get("/incoming", getIncomingInvoices)
	invoices.Get("/outgoing", getOutgoingInvoices)
	invoices.Post("/upload", uploadInvoice)
	invoices.Post("/send-draft", sendDraftInvoices)
	invoices.Get("/:direction/:uuid/:fileType", downloadInvoiceFile)

	// Despatches routes (E-İrsaliye)
	despatches := v1.Group("/despatches")
	despatches.Get("/incoming", getIncomingDespatches)
	despatches.Get("/outgoing", getOutgoingDespatches)
	despatches.Post("/upload", uploadDespatch)
	despatches.Post("/:uuid/answer", sendDespatchAnswer)

	// E-Archive routes
	earchive := v1.Group("/earchive")
	earchive.Get("/invoices", getEArchiveInvoices)
	earchive.Post("/upload", uploadEArchiveInvoice)
	earchive.Post("/send-draft", sendDraftEArchiveInvoices)
	earchive.Get("/invoices/:uuid/:fileType", downloadEArchiveFile)
	earchive.Post("/invoices/cancel", cancelEArchiveInvoice)

	// Account routes
	account := v1.Group("/account")
	account.Get("/default-document-parameter", getAccountDefaultDocumentParameter)
	account.Get("/document-archives", getAccountDocumentArchives)
}

// Handlers are located in handlers.go, handlers_staticcodes.go, handlers_statistics.go
