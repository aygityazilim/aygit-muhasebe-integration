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
	management.Get("/identifications", getIdentifications)
	management.Post("/identifications", createIdentification)
	management.Delete("/identifications/:id", deleteIdentification)

	// Dealer Info
	management.Get("/dealerinfo", getDealerInfo)

	// Address routes
	management.Get("/addresses", getAddresses)
	management.Post("/addresses", createAddress)
	management.Put("/address/:id", updateAddress)
	management.Delete("/address/:id", deleteAddress)

	// IVD routes
	management.Get("/ivd", getIVDSetting)
	management.Put("/ivd", updateIVDSetting)

	// Luca Integration routes
	management.Get("/luca/integration/setting", getLucaIntegrationSetting)
	management.Put("/luca/integration/setting", updateLucaIntegrationSetting)

	// Customer Search From GIB
	management.Get("/customer_search_from_gib/setting", getCustomerSearchFromGIBSetting)
	management.Put("/customer_search_from_gib/setting", updateCustomerSearchFromGIBSetting)

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
	earchive.Post("/invoices/export/:fileType", exportEArchiveInvoices)
	earchive.Put("/invoices/tags", updateEArchiveInvoiceTags)
	earchive.Post("/invoices/:uuid/savecompanyindocument", saveCompanyInEArchiveDocument)
	earchive.Put("/invoices/bulk/:operation", bulkEArchiveInvoiceOperation)
	earchive.Delete("/invoices/drafts", deleteEArchiveDraftInvoices)
	earchive.Post("/invoices/email/send", sendEArchiveInvoiceEmail)
	earchive.Get("/invoices/:uuid/usernotes", getEArchiveInvoiceUserNotes)
	earchive.Post("/invoices/:uuid/usernotes", createEArchiveInvoiceUserNote)
	earchive.Put("/invoices/:uuid/usernotes/:id", updateEArchiveInvoiceUserNote)
	earchive.Delete("/invoices/:uuid/usernotes/:id", deleteEArchiveInvoiceUserNote)

	earchive.Get("/notifications/dynamicrules", getEArchiveDynamicRules)
	earchive.Post("/notifications/dynamicrules", createEArchiveDynamicRule)
	earchive.Get("/notifications/dynamicrules/:id", getEArchiveDynamicRule)
	earchive.Put("/notifications/dynamicrules/:id", updateEArchiveDynamicRule)
	earchive.Delete("/notifications/dynamicrules/:id", deleteEArchiveDynamicRule)

	earchive.Get("/definitions/fileexporttitles/:documentType/titlekeys", getEArchiveFileExportTitles)
	earchive.Get("/definitions/fileexporttitles/:documentType/:extension", getEArchiveFileExportTitlesExtension)
	earchive.Put("/definitions/fileexporttitles", updateEArchiveFileExportTitles)

	earchive.Post("/uploads/document/preview", previewEArchiveDocument)
	earchive.Put("/uploads/document/:uuid", updateEArchiveDocument)
	earchive.Post("/uploads/draft/create/:id", createEArchiveDraft)
	earchive.Post("/uploads/marketplaces/:id/orders/:orderId/preview", previewEArchiveMarketplaceOrder)
	earchive.Post("/uploads/marketplaces/:id/orders/createinvoice", createEArchiveMarketplaceInvoice)

	earchive.Get("/exinvoices", getEArchiveExInvoices)
	earchive.Post("/exinvoices", uploadEArchiveExInvoice)
	earchive.Get("/exinvoices/queue", getEArchiveExInvoicesQueue)
	earchive.Get("/exinvoices/queue/:id", getEArchiveExInvoicesQueueResult)
	earchive.Get("/exinvoices/:uuid/:fileType", downloadEArchiveExInvoiceFile)
	earchive.Post("/exinvoices/export/:fileType", exportEArchiveExInvoices)

	earchive.Get("/tags", getEArchiveTags)
	earchive.Post("/tags", createEArchiveTag)
	earchive.Get("/tags/:id", getEArchiveTag)
	earchive.Put("/tags/:id", updateEArchiveTag)
	earchive.Delete("/tags/:id", deleteEArchiveTag)

	earchive.Get("/outgoing/reportmodule/reports", getEArchiveReports)
	earchive.Post("/outgoing/reportmodule/reports", createEArchiveReport)
	earchive.Get("/outgoing/reportmodule/reports/:id/download", downloadEArchiveReport)
	earchive.Get("/outgoing/reportmodule/templates", getEArchiveReportTemplates)
	earchive.Post("/outgoing/reportmodule/templates", createEArchiveReportTemplate)
	earchive.Get("/outgoing/reportmodule/templates/:id", getEArchiveReportTemplate)
	earchive.Put("/outgoing/reportmodule/templates/:id", updateEArchiveReportTemplate)
	earchive.Delete("/outgoing/reportmodule/templates/:id", deleteEArchiveReportTemplate)
	earchive.Get("/outgoing/reportmodule/columns", getEArchiveReportColumns)

	earchive.Get("/definitions/mailing/email/settings", getEArchiveEmailSettings)
	earchive.Put("/definitions/mailing/email/settings", updateEArchiveEmailSettings)
	earchive.Get("/definitions/mailing/sms/settings", getEArchiveSmsSettings)
	earchive.Put("/definitions/mailing/sms/settings", updateEArchiveSmsSettings)

	earchive.Get("/definitions/documenttemplates/customizationsettings", getEArchiveCustomizationSettings)
	earchive.Post("/definitions/documenttemplates/customizationsettings", createEArchiveCustomizationSetting)
	earchive.Get("/definitions/documenttemplates/customizationsettings/:id", getEArchiveCustomizationSetting)
	earchive.Put("/definitions/documenttemplates/customizationsettings/:id", updateEArchiveCustomizationSetting)
	earchive.Delete("/definitions/documenttemplates/customizationsettings/:id", deleteEArchiveCustomizationSetting)
	earchive.Get("/definitions/documenttemplates/customizationsettings/:id/setdefault", setEArchiveCustomizationSettingDefault)
	earchive.Post("/definitions/documenttemplates/customizationsettings/:id/preview", previewEArchiveCustomizationSetting)
	earchive.Get("/definitions/documenttemplates/customizationsettings/:id/logo", getEArchiveCustomizationSettingLogo)
	earchive.Post("/definitions/documenttemplates/customizationsettings/:id/logo", createEArchiveCustomizationSettingLogo)
	earchive.Delete("/definitions/documenttemplates/customizationsettings/:id/logo", deleteEArchiveCustomizationSettingLogo)
	earchive.Get("/definitions/documenttemplates/customizationsettings/:id/stamp", getEArchiveCustomizationSettingStamp)
	earchive.Post("/definitions/documenttemplates/customizationsettings/:id/stamp", createEArchiveCustomizationSettingStamp)
	earchive.Delete("/definitions/documenttemplates/customizationsettings/:id/stamp", deleteEArchiveCustomizationSettingStamp)
	earchive.Get("/definitions/documenttemplates/customizationsettings/:id/signature", getEArchiveCustomizationSettingSignature)
	earchive.Post("/definitions/documenttemplates/customizationsettings/:id/signature", createEArchiveCustomizationSettingSignature)
	earchive.Delete("/definitions/documenttemplates/customizationsettings/:id/signature", deleteEArchiveCustomizationSettingSignature)

	earchive.Get("/definitions/series", getEArchiveSeries)
	earchive.Post("/definitions/series", createEArchiveSerie)
	earchive.Get("/definitions/series/:id", getEArchiveSerie)
	earchive.Delete("/definitions/series/:id", deleteEArchiveSerie)
	earchive.Get("/definitions/series/prefix/:serie", getEArchiveSerieByPrefix)
	earchive.Get("/definitions/series/:id/set/:status", setEArchiveSerieStatus)
	earchive.Get("/definitions/series/:id/setdefault", setEArchiveSerieDefault)
	earchive.Get("/definitions/series/:id/:year/setnumber/:nextNumber", setEArchiveSerieNumber)
	earchive.Get("/definitions/series/:serieId/:year/histories", getEArchiveSerieHistories)

	earchive.Get("/definitions/documenttemplates", getEArchiveDocumentTemplates)
	earchive.Post("/definitions/documenttemplates", createEArchiveDocumentTemplate)
	earchive.Get("/definitions/documenttemplates/:id", downloadEArchiveDocumentTemplate)
	earchive.Put("/definitions/documenttemplates/:id", updateEArchiveDocumentTemplate)
	earchive.Delete("/definitions/documenttemplates/:id", deleteEArchiveDocumentTemplate)
	earchive.Get("/definitions/documenttemplates/:id/setdefault", setEArchiveDocumentTemplateDefault)
	earchive.Post("/definitions/documenttemplates/:id/preview", previewEArchiveDocumentTemplate)

	// Account routes
	account := v1.Group("/account")
	account.Get("/default-document-parameter", getAccountDefaultDocumentParameter)
	account.Put("/default-document-parameter", updateAccountDefaultDocumentParameter)
	account.Get("/document-archives", getAccountDocumentArchives)
	account.Post("/document-archives", createAccountDocumentArchive)

	// Account Modules routes
	accountModules := v1.Group("/accountmodules")
	accountModules.Post("/info", accountModulesInfo)
	accountModules.Post("/update", accountModulesUpdate)
	accountModules.Post("/getenvelopecontent", accountModulesGetEnvelopeContent)
	accountModules.Post("/setsignedcontent", accountModulesSetSignedContent)
	accountModules.Post("/getenvelopeinfo", accountModulesGetEnvelopeInfo)

	// E-Despatch routes (Extended)
	edespatch := v1.Group("/edespatch")

	// Definitions
	edespatch.Get("/definitions/fileexporttitles/:documentType/titlekeys", getEDespatchFileExportTitles)
	edespatch.Get("/definitions/fileexporttitles/:documentType/:extension", getEDespatchFileExportTitleDefinition)
	edespatch.Put("/definitions/fileexporttitles", updateEDespatchFileExportTitles)
	edespatch.Get("/definitions/mailing/email/settings", getEDespatchEmailSettings)
	edespatch.Put("/definitions/mailing/email/settings", updateEDespatchEmailSettings)
	edespatch.Get("/definitions/mailing/sms/settings", getEDespatchSmsSettings)
	edespatch.Put("/definitions/mailing/sms/settings", updateEDespatchSmsSettings)

	// Uploads
	edespatch.Put("/uploads/document/:uuid", updateEDespatchDocument)
	edespatch.Post("/uploads/resend/:uuid", resendErrorDocument)

	// Tags
	edespatch.Get("/tags", getEDespatchTags)
	edespatch.Post("/tags", createEDespatchTag)
	edespatch.Get("/tags/:id", getEDespatchTag)
	edespatch.Put("/tags/:id", updateEDespatchTag)
	edespatch.Delete("/tags/:id", deleteEDespatchTag)

	// Notifications Incoming
	edespatch.Get("/notifications/incoming/dynamicrules", getEDespatchIncomingDynamicRules)
	edespatch.Post("/notifications/incoming/dynamicrules", createEDespatchIncomingDynamicRule)
	edespatch.Get("/notifications/incoming/dynamicrules/:id", getEDespatchIncomingDynamicRule)
	edespatch.Put("/notifications/incoming/dynamicrules/:id", updateEDespatchIncomingDynamicRule)
	edespatch.Delete("/notifications/incoming/dynamicrules/:id", deleteEDespatchIncomingDynamicRule)

	// Incoming Despatches
	edespatch.Post("/incoming/despatches/export/:fileType", exportEDespatchIncomingDespatches)
	edespatch.Put("/incoming/despatches/tags", updateEDespatchIncomingTags)
	edespatch.Post("/incoming/despatches/:uuid/savecompanyindocument", saveCompanyInIncomingDocument)
	edespatch.Put("/incoming/despatches/bulk/:operation", bulkOperationIncomingDespatches)
	edespatch.Post("/incoming/despatches/:uuid/usernotes", addUserNoteToIncomingDespatch)
	edespatch.Put("/incoming/despatches/:uuid/usernotes/:id", updateUserNoteInIncomingDespatch)
	edespatch.Delete("/incoming/despatches/:uuid/usernotes/:id", deleteUserNoteFromIncomingDespatch)
	edespatch.Post("/incoming/despatches/:uuid/receiptadvice", sendReceiptAdviceForIncomingDespatch)
	edespatch.Post("/incoming/despatches/email/send", sendEmailForIncomingDespatch)

	// Incoming Receipt Advices
	edespatch.Get("/incoming/receiptadvices", getEDespatchIncomingReceiptAdvices)
	edespatch.Get("/incoming/receiptadvices/:uuid", getEDespatchIncomingReceiptAdvice)
	edespatch.Get("/incoming/receiptadvices/:uuid/html", getEDespatchIncomingReceiptAdviceHTML)
	edespatch.Get("/incoming/receiptadvices/:uuid/xml", getEDespatchIncomingReceiptAdviceXML)
	edespatch.Get("/incoming/receiptadvices/:uuid/pdf", getEDespatchIncomingReceiptAdvicePDF)

	// Incoming Reports
	edespatch.Get("/incoming/reportmodule/reports", getEDespatchIncomingReports)
	edespatch.Post("/incoming/reportmodule/reports", createEDespatchIncomingReport)
	edespatch.Get("/incoming/reportmodule/reports/:id/download", downloadEDespatchIncomingReport)
	edespatch.Get("/incoming/reportmodule/templates", getEDespatchIncomingTemplates)
	edespatch.Post("/incoming/reportmodule/templates", createEDespatchIncomingTemplate)
	edespatch.Get("/incoming/reportmodule/templates/:id", getEDespatchIncomingTemplate)
	edespatch.Put("/incoming/reportmodule/templates/:id", updateEDespatchIncomingTemplate)
	edespatch.Delete("/incoming/reportmodule/templates/:id", deleteEDespatchIncomingTemplate)
	edespatch.Get("/incoming/reportmodule/columns", getEDespatchIncomingColumns)

	// Notifications Outgoing
	edespatch.Get("/notifications/outgoing/dynamicrules", getEDespatchOutgoingDynamicRules)
	edespatch.Post("/notifications/outgoing/dynamicrules", createEDespatchOutgoingDynamicRule)
	edespatch.Get("/notifications/outgoing/dynamicrules/:id", getEDespatchOutgoingDynamicRule)
	edespatch.Put("/notifications/outgoing/dynamicrules/:id", updateEDespatchOutgoingDynamicRule)
	edespatch.Delete("/notifications/outgoing/dynamicrules/:id", deleteEDespatchOutgoingDynamicRule)

	// Outgoing Despatches
	edespatch.Post("/outgoing/despatches/export/:fileType", exportEDespatchOutgoingDespatches)
	edespatch.Put("/outgoing/despatches/tags", updateEDespatchOutgoingTags)
	edespatch.Put("/outgoing/despatches/:uuid/receiveralias", updateEDespatchOutgoingReceiverAlias)
	edespatch.Post("/outgoing/despatches/:uuid/savecompanyindocument", saveCompanyInOutgoingDocument)
	edespatch.Put("/outgoing/despatches/bulk/:operation", bulkOperationOutgoingDespatches)
	edespatch.Post("/outgoing/despatches/:uuid/usernotes", addUserNoteToOutgoingDespatch)
	edespatch.Put("/outgoing/despatches/:uuid/usernotes/:id", updateUserNoteInOutgoingDespatch)
	edespatch.Delete("/outgoing/despatches/:uuid/usernotes/:id", deleteUserNoteFromOutgoingDespatch)
	edespatch.Delete("/outgoing/despatches/drafts", deleteEDespatchOutgoingDrafts)
	edespatch.Post("/outgoing/despatches/email/send", sendEmailForOutgoingDespatch)

	// Outgoing Receipt Advices
	edespatch.Get("/outgoing/receiptadvices", getEDespatchOutgoingReceiptAdvices)
	edespatch.Get("/outgoing/receiptadvices/:uuid", getEDespatchOutgoingReceiptAdvice)
	edespatch.Get("/outgoing/receiptadvices/:uuid/html", getEDespatchOutgoingReceiptAdviceHTML)
	edespatch.Get("/outgoing/receiptadvices/:uuid/xml", getEDespatchOutgoingReceiptAdviceXML)
	edespatch.Get("/outgoing/receiptadvices/:uuid/pdf", getEDespatchOutgoingReceiptAdvicePDF)

	// Outgoing Reports
	edespatch.Get("/outgoing/reportmodule/reports", getEDespatchOutgoingReports)
	edespatch.Post("/outgoing/reportmodule/reports", createEDespatchOutgoingReport)
	edespatch.Get("/outgoing/reportmodule/reports/:id/download", downloadEDespatchOutgoingReport)
	edespatch.Get("/outgoing/reportmodule/templates", getEDespatchOutgoingTemplates)
	edespatch.Post("/outgoing/reportmodule/templates", createEDespatchOutgoingTemplate)
	edespatch.Get("/outgoing/reportmodule/templates/:id", getEDespatchOutgoingTemplate)
	edespatch.Put("/outgoing/reportmodule/templates/:id", updateEDespatchOutgoingTemplate)
	edespatch.Delete("/outgoing/reportmodule/templates/:id", deleteEDespatchOutgoingTemplate)
	edespatch.Get("/outgoing/reportmodule/columns", getEDespatchOutgoingColumns)

	// Users
	edespatch.Get("/users/zip/:aliasType", getEDespatchUsersZip)
	edespatch.Get("/users/:identifier/:aliasType", getEDespatchUserByIdentifier)
	edespatch.Post("/users/:aliasType", getEDespatchUserByIdentifierPost)
	edespatch.Get("/users/search/:query/:aliasType", searchEDespatchUserByTitle)

	// Voucher routes

	// Definitions routes (Customization Settings)
	customizations := v1.Group("/definitions/documenttemplates/customizationsettings")
	customizations.Get("/", GetCustomizationSettings)
	customizations.Post("/", CreateCustomizationSetting)
	customizations.Get("/:id", GetCustomizationSettingByID)
	customizations.Put("/:id", UpdateCustomizationSetting)
	customizations.Delete("/:id", DeleteCustomizationSetting)
	customizations.Get("/:id/setdefault", SetDefaultCustomizationSetting)
	customizations.Post("/:id/preview", PreviewCustomizationSetting)

	customizations.Get("/:id/:imageType", GetCustomizationSettingImage)
	customizations.Post("/:id/:imageType", UploadCustomizationSettingImage)
	customizations.Delete("/:id/:imageType", DeleteCustomizationSettingImage)

	// Definitions routes (Series)
	definitionsSeries := v1.Group("/definitions/series")
	definitionsSeries.Get("/", GetSeries)
	definitionsSeries.Post("/", CreateSeries)
	definitionsSeries.Get("/:param", GetSeriesByIDOrPrefix)
	definitionsSeries.Delete("/:id", DeleteSeries)
	definitionsSeries.Get("/:id/set/:status", SetSeriesStatus)
	definitionsSeries.Get("/:id/setdefault", SetDefaultSeries)
	definitionsSeries.Get("/:id/:year/setnumber/:nextNumber", SetSeriesNextNumber)
	definitionsSeries.Get("/:serieId/:year/histories", GetSeriesHistories)

	// Answer Series routes
	answerSeries := v1.Group("/answerseries")
	answerSeries.Get("/", GetSeries)
	answerSeries.Post("/", CreateSeries)
	answerSeries.Get("/:param", GetSeriesByIDOrPrefix)
	answerSeries.Delete("/:id", DeleteSeries)
	answerSeries.Get("/:id/set/:status", SetSeriesStatus)
	answerSeries.Get("/:id/setdefault", SetDefaultSeries)
	answerSeries.Get("/:id/:year/setnumber/:nextNumber", SetSeriesNextNumber)
	answerSeries.Get("/:serieId/:year/histories", GetSeriesHistories)

	// Document Templates routes
	docTemplates := v1.Group("/definitions/documenttemplates")
	docTemplates.Get("/", GetDocumentTemplates)
	docTemplates.Post("/", CreateDocumentTemplate)
	docTemplates.Get("/:id", DownloadDocumentTemplate)
	docTemplates.Put("/:id", UpdateDocumentTemplate)
	docTemplates.Delete("/:id", DeleteDocumentTemplate)
	docTemplates.Get("/:id/setdefault", SetDefaultDocumentTemplate)
	docTemplates.Post("/:id/preview", PreviewDocumentTemplate)

	// Answer Document Templates routes
	answerDocTemplates := v1.Group("/definitions/answerdocumenttemplates")
	answerDocTemplates.Get("/", GetDocumentTemplates)
	answerDocTemplates.Post("/", CreateDocumentTemplate)
	answerDocTemplates.Get("/:id", DownloadDocumentTemplate)
	answerDocTemplates.Put("/:id", UpdateDocumentTemplate)
	answerDocTemplates.Delete("/:id", DeleteDocumentTemplate)
	answerDocTemplates.Get("/:id/setdefault", SetDefaultDocumentTemplate)
	answerDocTemplates.Post("/:id/preview", PreviewDocumentTemplate)

	// Envelope routes
	envelopes := v1.Group("/envelopes")
	envelopes.Get("/:instanceIdentifier/query", QueryEnvelopeStatus)
	// E-Invoice routes
	einvoice := v1.Group("/einvoice")
	einvoice.Get("/definitions/fileexporttitles/:documentType/titlekeys", getDefinitionsFileexporttitlesDocumenttypeTitlekeysEinvoice)
	einvoice.Get("/definitions/fileexporttitles/:documentType/:extension", getDefinitionsFileexporttitlesDocumenttypeExtensionEinvoice)
	einvoice.Put("/definitions/fileexporttitles", putDefinitionsFileexporttitlesEinvoice)
	einvoice.Put("/uploads/document/:uuid", putUploadsDocumentUuidEinvoice)
	einvoice.Post("/uploads/draft/create/:id", postUploadsDraftCreateIdEinvoice)
	einvoice.Post("/uploads/marketplaces/:id/orders/:orderId/preview", postUploadsMarketplacesIdOrdersOrderidPreviewEinvoice)
	einvoice.Post("/uploads/marketplaces/:id/orders/createinvoice", postUploadsMarketplacesIdOrdersCreateinvoiceEinvoice)
	einvoice.Post("/exinvoices", postExinvoicesEinvoice)
	einvoice.Get("/exinvoices/queue", getExinvoicesQueueEinvoice)
	einvoice.Get("/exinvoices/queue/:id", getExinvoicesQueueIdEinvoice)
	einvoice.Get("/tags", getTagsEinvoice)
	einvoice.Post("/tags", postTagsEinvoice)
	einvoice.Get("/tags/:id", getTagsIdEinvoice)
	einvoice.Put("/tags/:id", putTagsIdEinvoice)
	einvoice.Delete("/tags/:id", deleteTagsIdEinvoice)
	einvoice.Get("/exinvoices/incoming", getExinvoicesIncomingEinvoice)
	einvoice.Get("/exinvoices/incoming/:uuid/xml", getExinvoicesIncomingUuidXmlEinvoice)
	einvoice.Get("/exinvoices/incoming/:uuid/pdf", getExinvoicesIncomingUuidPdfEinvoice)
	einvoice.Get("/exinvoices/incoming/:uuid/html", getExinvoicesIncomingUuidHtmlEinvoice)
	einvoice.Post("/exinvoices/incoming/export/:fileType", postExinvoicesIncomingExportFiletypeEinvoice)
	einvoice.Get("/notifications/incoming/dynamicrules", getNotificationsIncomingDynamicrulesEinvoice)
	einvoice.Post("/notifications/incoming/dynamicrules", postNotificationsIncomingDynamicrulesEinvoice)
	einvoice.Get("/notifications/incoming/dynamicrules/:id", getNotificationsIncomingDynamicrulesIdEinvoice)
	einvoice.Put("/notifications/incoming/dynamicrules/:id", putNotificationsIncomingDynamicrulesIdEinvoice)
	einvoice.Delete("/notifications/incoming/dynamicrules/:id", deleteNotificationsIncomingDynamicrulesIdEinvoice)
	einvoice.Post("/incoming/invoices/export/:fileType", postIncomingInvoicesExportFiletypeEinvoice)
	einvoice.Get("/incoming/reportmodule/reports", getIncomingReportmoduleReportsEinvoice)
	einvoice.Post("/incoming/reportmodule/reports", postIncomingReportmoduleReportsEinvoice)
	einvoice.Get("/incoming/reportmodule/reports/:id/download", getIncomingReportmoduleReportsIdDownloadEinvoice)
	einvoice.Get("/incoming/reportmodule/templates", getIncomingReportmoduleTemplatesEinvoice)
	einvoice.Post("/incoming/reportmodule/templates", postIncomingReportmoduleTemplatesEinvoice)
	einvoice.Get("/incoming/reportmodule/templates/:id", getIncomingReportmoduleTemplatesIdEinvoice)
	einvoice.Put("/incoming/reportmodule/templates/:id", putIncomingReportmoduleTemplatesIdEinvoice)
	einvoice.Delete("/incoming/reportmodule/templates/:id", deleteIncomingReportmoduleTemplatesIdEinvoice)
	einvoice.Get("/incoming/reportmodule/columns", getIncomingReportmoduleColumnsEinvoice)
	einvoice.Put("/incoming/invoices/tags", putIncomingInvoicesTagsEinvoice)
	einvoice.Post("/incoming/invoices/:uuid/savecompanyindocument", postIncomingInvoicesUuidSavecompanyindocumentEinvoice)
	einvoice.Post("/incoming/invoices/:uuid/documentAnswer", postIncomingInvoicesUuidDocumentanswerEinvoice)
	einvoice.Put("/incoming/invoices/bulk/:operation", putIncomingInvoicesBulkOperationEinvoice)
	einvoice.Post("/incoming/invoices/:id/createreturninvoice", postIncomingInvoicesIdCreatereturninvoiceEinvoice)
	einvoice.Post("/incoming/invoices/:uuid/usernotes", postIncomingInvoicesUuidUsernotesEinvoice)
	einvoice.Put("/incoming/invoices/:uuid/usernotes/:id", putIncomingInvoicesUuidUsernotesIdEinvoice)
	einvoice.Delete("/incoming/invoices/:uuid/usernotes/:id", deleteIncomingInvoicesUuidUsernotesIdEinvoice)
	einvoice.Post("/incoming/invoices/email/send", postIncomingInvoicesEmailSendEinvoice)
	einvoice.Get("/exinvoices/outgoing", getExinvoicesOutgoingEinvoice)
	einvoice.Get("/exinvoices/outgoing/:uuid/xml", getExinvoicesOutgoingUuidXmlEinvoice)
	einvoice.Get("/exinvoices/outgoing/:uuid/pdf", getExinvoicesOutgoingUuidPdfEinvoice)
	einvoice.Get("/exinvoices/outgoing/:uuid/html", getExinvoicesOutgoingUuidHtmlEinvoice)
	einvoice.Post("/exinvoices/outgoing/export/:fileType", postExinvoicesOutgoingExportFiletypeEinvoice)
	einvoice.Get("/notifications/outgoing/dynamicrules", getNotificationsOutgoingDynamicrulesEinvoice)
	einvoice.Post("/notifications/outgoing/dynamicrules", postNotificationsOutgoingDynamicrulesEinvoice)
	einvoice.Get("/notifications/outgoing/dynamicrules/:id", getNotificationsOutgoingDynamicrulesIdEinvoice)
	einvoice.Put("/notifications/outgoing/dynamicrules/:id", putNotificationsOutgoingDynamicrulesIdEinvoice)
	einvoice.Delete("/notifications/outgoing/dynamicrules/:id", deleteNotificationsOutgoingDynamicrulesIdEinvoice)
	einvoice.Post("/outgoing/invoices/export/:fileType", postOutgoingInvoicesExportFiletypeEinvoice)
	einvoice.Get("/outgoing/reportmodule/reports", getOutgoingReportmoduleReportsEinvoice)
	einvoice.Post("/outgoing/reportmodule/reports", postOutgoingReportmoduleReportsEinvoice)
	einvoice.Get("/outgoing/reportmodule/reports/:id/download", getOutgoingReportmoduleReportsIdDownloadEinvoice)
	einvoice.Get("/outgoing/reportmodule/templates", getOutgoingReportmoduleTemplatesEinvoice)
	einvoice.Post("/outgoing/reportmodule/templates", postOutgoingReportmoduleTemplatesEinvoice)
	einvoice.Get("/outgoing/reportmodule/templates/:id", getOutgoingReportmoduleTemplatesIdEinvoice)
	einvoice.Put("/outgoing/reportmodule/templates/:id", putOutgoingReportmoduleTemplatesIdEinvoice)
	einvoice.Delete("/outgoing/reportmodule/templates/:id", deleteOutgoingReportmoduleTemplatesIdEinvoice)
	einvoice.Get("/outgoing/reportmodule/columns", getOutgoingReportmoduleColumnsEinvoice)
	einvoice.Put("/outgoing/invoices/tags", putOutgoingInvoicesTagsEinvoice)
	einvoice.Put("/outgoing/invoices/:uuid/receiveralias", putOutgoingInvoicesUuidReceiveraliasEinvoice)
	einvoice.Post("/outgoing/invoices/:uuid/savecompanyindocument", postOutgoingInvoicesUuidSavecompanyindocumentEinvoice)
	einvoice.Put("/outgoing/invoices/bulk/:operation", putOutgoingInvoicesBulkOperationEinvoice)
	einvoice.Post("/uploads/resend/:uuid", postUploadsResendUuidEinvoice)
	einvoice.Post("/outgoing/invoices/:uuid/usernotes", postOutgoingInvoicesUuidUsernotesEinvoice)
	einvoice.Put("/outgoing/invoices/:uuid/usernotes/:id", putOutgoingInvoicesUuidUsernotesIdEinvoice)
	einvoice.Delete("/outgoing/invoices/:uuid/usernotes/:id", deleteOutgoingInvoicesUuidUsernotesIdEinvoice)
	einvoice.Delete("/outgoing/invoices/drafts", deleteOutgoingInvoicesDraftsEinvoice)
	einvoice.Post("/outgoing/invoices/email/send", postOutgoingInvoicesEmailSendEinvoice)
	einvoice.Get("/users/zip/:aliasType", getUsersZipAliastypeEinvoice)
	einvoice.Get("/users/:identifier/:aliasType", getUsersIdentifierAliastypeEinvoice)
	einvoice.Post("/users/:aliasType", postUsersAliastypeEinvoice)
	einvoice.Get("/users/search/:query/:aliasType", getUsersSearchQueryAliastypeEinvoice)
	einvoice.Get("/definitions/documenttemplates/customizationsettings", getDefinitionsDocumenttemplatesCustomizationsettingsEinvoice)
	einvoice.Post("/definitions/documenttemplates/customizationsettings", postDefinitionsDocumenttemplatesCustomizationsettingsEinvoice)
	einvoice.Get("/definitions/documenttemplates/customizationsettings/:id", getDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice)
	einvoice.Put("/definitions/documenttemplates/customizationsettings/:id", putDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice)
	einvoice.Delete("/definitions/documenttemplates/customizationsettings/:id", deleteDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice)
	einvoice.Get("/definitions/documenttemplates/customizationsettings/:id/setdefault", getDefinitionsDocumenttemplatesCustomizationsettingsIdSetdefaultEinvoice)
	einvoice.Post("/definitions/documenttemplates/customizationsettings/:id/preview", postDefinitionsDocumenttemplatesCustomizationsettingsIdPreviewEinvoice)
	einvoice.Get("/definitions/documenttemplates/customizationsettings/:id/logo", getDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice)
	einvoice.Post("/definitions/documenttemplates/customizationsettings/:id/logo", postDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice)
	einvoice.Delete("/definitions/documenttemplates/customizationsettings/:id/logo", deleteDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice)
	einvoice.Get("/definitions/documenttemplates/customizationsettings/:id/stamp", getDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice)
	einvoice.Post("/definitions/documenttemplates/customizationsettings/:id/stamp", postDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice)
	einvoice.Delete("/definitions/documenttemplates/customizationsettings/:id/stamp", deleteDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice)
	einvoice.Get("/definitions/documenttemplates/customizationsettings/:id/signature", getDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice)
	einvoice.Post("/definitions/documenttemplates/customizationsettings/:id/signature", postDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice)
	einvoice.Delete("/definitions/documenttemplates/customizationsettings/:id/signature", deleteDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice)
	einvoice.Get("/definitions/series", getDefinitionsSeriesEinvoice)
	einvoice.Post("/definitions/series", postDefinitionsSeriesEinvoice)
	einvoice.Get("/definitions/series/:id", getDefinitionsSeriesIdEinvoice)
	einvoice.Delete("/definitions/series/:id", deleteDefinitionsSeriesIdEinvoice)
	einvoice.Get("/definitions/series/:serie", getDefinitionsSeriesSerieEinvoice)
	einvoice.Get("/definitions/series/:id/set/:status", getDefinitionsSeriesIdSetStatusEinvoice)
	einvoice.Get("/definitions/series/:id/setdefault", getDefinitionsSeriesIdSetdefaultEinvoice)
	einvoice.Get("/definitions/series/:id/:year/setnumber/:nextNumber", getDefinitionsSeriesIdYearSetnumberNextnumberEinvoice)
	einvoice.Get("/definitions/series/:serieId/:year/histories", getDefinitionsSeriesSerieidYearHistoriesEinvoice)
	einvoice.Get("/definitions/documenttemplates", getDefinitionsDocumenttemplatesEinvoice)
	einvoice.Post("/definitions/documenttemplates", postDefinitionsDocumenttemplatesEinvoice)
	einvoice.Get("/definitions/documenttemplates/:id", getDefinitionsDocumenttemplatesIdEinvoice)
	einvoice.Put("/definitions/documenttemplates/:id", putDefinitionsDocumenttemplatesIdEinvoice)
	einvoice.Delete("/definitions/documenttemplates/:id", deleteDefinitionsDocumenttemplatesIdEinvoice)
	einvoice.Get("/definitions/documenttemplates/:id/setdefault", getDefinitionsDocumenttemplatesIdSetdefaultEinvoice)
	einvoice.Post("/definitions/documenttemplates/:id/preview", postDefinitionsDocumenttemplatesIdPreviewEinvoice)
	einvoice.Get("/envelopes/:instanceIdentifier/query", getEnvelopesInstanceidentifierQueryEinvoice)

	SetupVoucherRoutes(v1)
}

// Handlers are located in handlers.go, handlers_staticcodes.go, handlers_statistics.go

func SetupVoucherRoutes(v1 fiber.Router) {
	voucherHandler := NewVoucherHandler()

	// Create a generic group for EMM/ESMM operations
	vGroup := v1.Group("/:type")

	// Notifications / Dynamic Rules
	vGroup.Get("/notifications/dynamicrules", voucherHandler.GetDynamicRules)
	vGroup.Post("/notifications/dynamicrules", voucherHandler.CreateDynamicRule)
	vGroup.Get("/notifications/dynamicrules/:id", voucherHandler.GetDynamicRuleByID)
	vGroup.Put("/notifications/dynamicrules/:id", voucherHandler.UpdateDynamicRule)
	vGroup.Delete("/notifications/dynamicrules/:id", voucherHandler.DeleteDynamicRule)

	// Vouchers
	vGroup.Get("/vouchers", voucherHandler.GetVouchers)
	vGroup.Post("/vouchers/export/:fileType", voucherHandler.ExportVouchers)
	vGroup.Post("/vouchers/:uuid/usernotes", voucherHandler.AddUserNoteForward)
	vGroup.Put("/vouchers/:uuid/usernotes/:id", voucherHandler.UpdateUserNote)
	vGroup.Delete("/vouchers/:uuid/usernotes/:id", voucherHandler.DeleteUserNote)
	vGroup.Post("/vouchers/cancel", voucherHandler.CancelVoucherForward)
	vGroup.Post("/vouchers/canceled/withdraw", voucherHandler.WithdrawCancelVoucherForward)
	vGroup.Put("/vouchers/tags", voucherHandler.UpdateVouchersTags)
	vGroup.Post("/vouchers/:uuid/savecompanyindocument", voucherHandler.SaveCompanyInDocument)
	vGroup.Put("/vouchers/bulk/:operation", voucherHandler.BulkVouchersOperation)
	vGroup.Delete("/vouchers/drafts", voucherHandler.DeleteDrafts)
	vGroup.Post("/vouchers/email/send", voucherHandler.SendVoucherEmailForward)

	// Existing Voucher routes remapped onto generic vGroup to match NES structure better
	vGroup.Get("/vouchers/drafts", voucherHandler.GetDraftVouchers)

	// Uploads
	vGroup.Post("/uploads/document", voucherHandler.UploadDocument)
	vGroup.Put("/uploads/document/:uuid", voucherHandler.UpdateDocumentUpload)
	vGroup.Post("/uploads/draft/send", voucherHandler.SendDraft)

	// Definitions
	vGroup.Get("/definitions/fileexporttitles/:documentType/titlekeys", voucherHandler.GetFileExportTitleKeys)
	vGroup.Get("/definitions/fileexporttitles/:documentType/:extension", voucherHandler.GetFileExportTitles)
	vGroup.Put("/definitions/fileexporttitles", voucherHandler.UpdateFileExportTitles)
	vGroup.Get("/definitions/mailing/email/settings", voucherHandler.GetEmailSettings)
	vGroup.Put("/definitions/mailing/email/settings", voucherHandler.UpdateEmailSettings)
	vGroup.Get("/definitions/mailing/sms/settings", voucherHandler.GetSMSSettings)
	vGroup.Put("/definitions/mailing/sms/settings", voucherHandler.UpdateSMSSettings)
	vGroup.Get("/definitions/documenttemplates/customizationsettings", voucherHandler.GetCustomizationSettings)
	vGroup.Post("/definitions/documenttemplates/customizationsettings", voucherHandler.CreateCustomizationSetting)
	vGroup.Get("/definitions/documenttemplates/customizationsettings/:id", voucherHandler.GetCustomizationSettingByID)
	vGroup.Put("/definitions/documenttemplates/customizationsettings/:id", voucherHandler.UpdateCustomizationSetting)
	vGroup.Delete("/definitions/documenttemplates/customizationsettings/:id", voucherHandler.DeleteCustomizationSetting)
	vGroup.Get("/definitions/documenttemplates/customizationsettings/:id/setdefault", voucherHandler.SetCustomizationSettingDefault)
	vGroup.Post("/definitions/documenttemplates/customizationsettings/:id/preview", voucherHandler.PreviewCustomizationSetting)
	vGroup.Get("/definitions/documenttemplates/customizationsettings/:id/logo", voucherHandler.GetCustomizationSettingLogo)
	vGroup.Post("/definitions/documenttemplates/customizationsettings/:id/logo", voucherHandler.UploadCustomizationSettingLogo)
	vGroup.Delete("/definitions/documenttemplates/customizationsettings/:id/logo", voucherHandler.DeleteCustomizationSettingLogo)
	vGroup.Get("/definitions/documenttemplates/customizationsettings/:id/stamp", voucherHandler.GetCustomizationSettingStamp)
	vGroup.Post("/definitions/documenttemplates/customizationsettings/:id/stamp", voucherHandler.UploadCustomizationSettingStamp)
	vGroup.Delete("/definitions/documenttemplates/customizationsettings/:id/stamp", voucherHandler.DeleteCustomizationSettingStamp)
	vGroup.Get("/definitions/documenttemplates/customizationsettings/:id/signature", voucherHandler.GetCustomizationSettingSignature)
	vGroup.Post("/definitions/documenttemplates/customizationsettings/:id/signature", voucherHandler.UploadCustomizationSettingSignature)
	vGroup.Delete("/definitions/documenttemplates/customizationsettings/:id/signature", voucherHandler.DeleteCustomizationSettingSignature)

	vGroup.Get("/definitions/series", voucherHandler.GetSeries)
	vGroup.Post("/definitions/series", voucherHandler.CreateSeries)
	vGroup.Get("/definitions/series/:id", voucherHandler.GetSeriesByID)
	vGroup.Delete("/definitions/series/:id", voucherHandler.DeleteSeries)
	vGroup.Get("/definitions/series/prefix/:serie", voucherHandler.GetSeriesByPrefix)
	vGroup.Get("/definitions/series/:id/set/:status", voucherHandler.UpdateSeriesStatus)
	vGroup.Get("/definitions/series/:id/setdefault", voucherHandler.SetSeriesDefault)
	vGroup.Get("/definitions/series/:id/:year/setnumber/:nextNumber", voucherHandler.UpdateSeriesNumber)
	vGroup.Get("/definitions/series/:serieId/:year/histories", voucherHandler.GetSeriesHistories)

	vGroup.Get("/definitions/documenttemplates", voucherHandler.GetDocumentTemplates)
	vGroup.Post("/definitions/documenttemplates", voucherHandler.CreateDocumentTemplate)
	vGroup.Get("/definitions/documenttemplates/:id", voucherHandler.DownloadDocumentTemplate)
	vGroup.Put("/definitions/documenttemplates/:id", voucherHandler.UpdateDocumentTemplate)
	vGroup.Delete("/definitions/documenttemplates/:id", voucherHandler.DeleteDocumentTemplate)
	vGroup.Get("/definitions/documenttemplates/:id/setdefault", voucherHandler.SetDocumentTemplateDefault)
	vGroup.Post("/definitions/documenttemplates/:id/preview", voucherHandler.PreviewDocumentTemplate)

	// Tags
	vGroup.Get("/tags", voucherHandler.GetTags)
	vGroup.Post("/tags", voucherHandler.CreateTag)
	vGroup.Get("/tags/:id", voucherHandler.GetTagByID)
	vGroup.Put("/tags/:id", voucherHandler.UpdateTag)
	vGroup.Delete("/tags/:id", voucherHandler.DeleteTag)

	// Outgoing Report Module
	vGroup.Get("/outgoing/reportmodule/reports", voucherHandler.GetReports)
	vGroup.Post("/outgoing/reportmodule/reports", voucherHandler.CreateReport)
	vGroup.Get("/outgoing/reportmodule/reports/:id/download", voucherHandler.DownloadReport)
	vGroup.Get("/outgoing/reportmodule/templates", voucherHandler.GetReportTemplates)
	vGroup.Post("/outgoing/reportmodule/templates", voucherHandler.CreateReportTemplate)
	vGroup.Get("/outgoing/reportmodule/templates/:id", voucherHandler.GetReportTemplateByID)
	vGroup.Put("/outgoing/reportmodule/templates/:id", voucherHandler.UpdateReportTemplate)
	vGroup.Delete("/outgoing/reportmodule/templates/:id", voucherHandler.DeleteReportTemplate)
	vGroup.Get("/outgoing/reportmodule/columns", voucherHandler.GetReportColumns)
	// Vouchers
	vouchers := v1.Group("/:type/vouchers")
	vouchers.Get("/", voucherHandler.GetVouchers)
	vouchers.Get("/drafts", voucherHandler.GetDraftVouchers)
	vouchers.Delete("/drafts", voucherHandler.DeleteDrafts)
	vouchers.Post("/export/:fileType", voucherHandler.ExportVouchers)
	vouchers.Post("/cancel", voucherHandler.CancelVoucherForward)
	vouchers.Post("/canceled/withdraw", voucherHandler.WithdrawCancelVoucherForward)
	vouchers.Put("/tags", voucherHandler.UpdateVouchersTags)
	vouchers.Put("/bulk/:operation", voucherHandler.BulkVouchersOperation)
	vouchers.Post("/email/send", voucherHandler.SendVoucherEmailForward)
	vouchers.Post("/:uuid/savecompanyindocument", voucherHandler.SaveCompanyInDocument)
	vouchers.Post("/:uuid/usernotes", voucherHandler.AddUserNoteForward)
	vouchers.Put("/:uuid/usernotes/:id", voucherHandler.UpdateUserNote)
	vouchers.Delete("/:uuid/usernotes/:id", voucherHandler.DeleteUserNote)

	// Uploads
	uploads := v1.Group("/:type/uploads")
	uploads.Post("/document", voucherHandler.UploadDocument)
	uploads.Put("/document/:uuid", voucherHandler.UpdateDocumentUpload)
	uploads.Post("/draft/send", voucherHandler.SendDraft)
	uploads.Post("/draft/create/:id", voucherHandler.CreateDraftUpload)

	// Notifications
	notifications := v1.Group("/:type/notifications")
	notifications.Get("/dynamicrules", voucherHandler.GetDynamicRules)
	notifications.Post("/dynamicrules", voucherHandler.CreateDynamicRule)
	notifications.Get("/dynamicrules/:id", voucherHandler.GetDynamicRuleByID)
	notifications.Put("/dynamicrules/:id", voucherHandler.UpdateDynamicRule)
	notifications.Delete("/dynamicrules/:id", voucherHandler.DeleteDynamicRule)

	// Definitions
	definitions := v1.Group("/:type/definitions")
	definitions.Get("/fileexporttitles/:documentType/titlekeys", voucherHandler.GetFileExportTitleKeys)
	definitions.Get("/fileexporttitles/:documentType/:extension", voucherHandler.GetFileExportTitles)
	definitions.Put("/fileexporttitles", voucherHandler.UpdateFileExportTitles)

	definitions.Get("/mailing/email/settings", voucherHandler.GetEmailSettings)
	definitions.Put("/mailing/email/settings", voucherHandler.UpdateEmailSettings)
	definitions.Get("/mailing/sms/settings", voucherHandler.GetSMSSettings)
	definitions.Put("/mailing/sms/settings", voucherHandler.UpdateSMSSettings)

	definitions.Get("/documenttemplates", voucherHandler.GetDocumentTemplates)
	definitions.Post("/documenttemplates", voucherHandler.CreateDocumentTemplate)
	definitions.Get("/documenttemplates/:id", voucherHandler.DownloadDocumentTemplate)
	definitions.Put("/documenttemplates/:id", voucherHandler.UpdateDocumentTemplate)
	definitions.Delete("/documenttemplates/:id", voucherHandler.DeleteDocumentTemplate)
	definitions.Get("/documenttemplates/:id/setdefault", voucherHandler.SetDocumentTemplateDefault)
	definitions.Post("/documenttemplates/:id/preview", voucherHandler.PreviewDocumentTemplate)

	definitions.Get("/documenttemplates/customizationsettings", voucherHandler.GetCustomizationSettings)
	definitions.Post("/documenttemplates/customizationsettings", voucherHandler.CreateCustomizationSetting)
	definitions.Get("/documenttemplates/customizationsettings/:id", voucherHandler.GetCustomizationSettingByID)
	definitions.Put("/documenttemplates/customizationsettings/:id", voucherHandler.UpdateCustomizationSetting)
	definitions.Delete("/documenttemplates/customizationsettings/:id", voucherHandler.DeleteCustomizationSetting)
	definitions.Get("/documenttemplates/customizationsettings/:id/setdefault", voucherHandler.SetCustomizationSettingDefault)
	definitions.Post("/documenttemplates/customizationsettings/:id/preview", voucherHandler.PreviewCustomizationSetting)

	definitions.Get("/documenttemplates/customizationsettings/:id/logo", voucherHandler.GetCustomizationSettingLogo)
	definitions.Post("/documenttemplates/customizationsettings/:id/logo", voucherHandler.UploadCustomizationSettingLogo)
	definitions.Delete("/documenttemplates/customizationsettings/:id/logo", voucherHandler.DeleteCustomizationSettingLogo)
	definitions.Get("/documenttemplates/customizationsettings/:id/stamp", voucherHandler.GetCustomizationSettingStamp)
	definitions.Post("/documenttemplates/customizationsettings/:id/stamp", voucherHandler.UploadCustomizationSettingStamp)
	definitions.Delete("/documenttemplates/customizationsettings/:id/stamp", voucherHandler.DeleteCustomizationSettingStamp)
	definitions.Get("/documenttemplates/customizationsettings/:id/signature", voucherHandler.GetCustomizationSettingSignature)
	definitions.Post("/documenttemplates/customizationsettings/:id/signature", voucherHandler.UploadCustomizationSettingSignature)
	definitions.Delete("/documenttemplates/customizationsettings/:id/signature", voucherHandler.DeleteCustomizationSettingSignature)

	definitions.Get("/series", voucherHandler.GetSeries)
	definitions.Post("/series", voucherHandler.CreateSeries)
	definitions.Get("/series/:id", voucherHandler.GetSeriesByID)
	definitions.Delete("/series/:id", voucherHandler.DeleteSeries)
	// /series/:serie will conflict with /series/:id if not careful, but both match a string. Assuming fiber handles correctly or user passes different formats.
	definitions.Get("/series/:serie/byprefix", voucherHandler.GetSeriesByPrefix) // Changed to avoid collision, standard practice
	definitions.Get("/series/:id/set/:status", voucherHandler.UpdateSeriesStatus)
	definitions.Get("/series/:id/setdefault", voucherHandler.SetSeriesDefault)
	definitions.Get("/series/:id/:year/setnumber/:nextNumber", voucherHandler.UpdateSeriesNumber)
	definitions.Get("/series/:serieId/:year/histories", voucherHandler.GetSeriesHistories)

	// Tags
	tags := v1.Group("/:type/tags")
	tags.Get("/", voucherHandler.GetTags)
	tags.Post("/", voucherHandler.CreateTag)
	tags.Get("/:id", voucherHandler.GetTagByID)
	tags.Put("/:id", voucherHandler.UpdateTag)
	tags.Delete("/:id", voucherHandler.DeleteTag)

	// Outgoing
	outgoing := v1.Group("/:type/outgoing")
	outgoing.Get("/reportmodule/reports", voucherHandler.GetReports)
	outgoing.Post("/reportmodule/reports", voucherHandler.CreateReport)
	outgoing.Get("/reportmodule/reports/:id/download", voucherHandler.DownloadReport)

	outgoing.Get("/reportmodule/templates", voucherHandler.GetReportTemplates)
	outgoing.Post("/reportmodule/templates", voucherHandler.CreateReportTemplate)
	outgoing.Get("/reportmodule/templates/:id", voucherHandler.GetReportTemplateByID)
	outgoing.Put("/reportmodule/templates/:id", voucherHandler.UpdateReportTemplate)
	outgoing.Delete("/reportmodule/templates/:id", voucherHandler.DeleteReportTemplate)

	outgoing.Get("/reportmodule/columns", voucherHandler.GetReportColumns)
}
