package service

import (
	"aygit-muhasebe-integration/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
)


// GetDefinitionsFileexporttitlesDocumenttypeTitlekeysEinvoice Kullanılabilir alanları listeler
func (s *NESService) GetDefinitionsFileexporttitlesDocumenttypeTitlekeysEinvoice(company *models.Company, documentType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/fileexporttitles/%s/titlekeys", s.getBaseURL(company), documentType)

	return s.doRequest(company, "GET", url, nil)
}


// GetDefinitionsFileexporttitlesDocumenttypeExtensionEinvoice Tanımları getirir
func (s *NESService) GetDefinitionsFileexporttitlesDocumenttypeExtensionEinvoice(company *models.Company, documentType string, extension string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/fileexporttitles/%s/%s", s.getBaseURL(company), documentType, extension)

	return s.doRequest(company, "GET", url, nil)
}


// PutDefinitionsFileexporttitlesEinvoice Tanımları günceller
func (s *NESService) PutDefinitionsFileexporttitlesEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/fileexporttitles", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// PutUploadsDocumentUuidEinvoice Belge günceller
func (s *NESService) PutUploadsDocumentUuidEinvoice(company *models.Company, payload map[string]interface{}, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/uploads/document/%s", s.getBaseURL(company), uuid)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// PostUploadsDraftCreateIdEinvoice Belge yükler
func (s *NESService) PostUploadsDraftCreateIdEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/uploads/draft/create/%s", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PostUploadsMarketplacesIdOrdersOrderidPreviewEinvoice Belirtilen pazaryerindeki siparişin faturasını önizler
func (s *NESService) PostUploadsMarketplacesIdOrdersOrderidPreviewEinvoice(company *models.Company, payload map[string]interface{}, id string, orderId string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/uploads/marketplaces/%s/orders/%s/preview", s.getBaseURL(company), id, orderId)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PostUploadsMarketplacesIdOrdersCreateinvoiceEinvoice Belirtilen pazaryerindeki siparişin faturasını önizler
func (s *NESService) PostUploadsMarketplacesIdOrdersCreateinvoiceEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/uploads/marketplaces/%s/orders/createinvoice", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PostExinvoicesEinvoice Eski belge yükler
func (s *NESService) PostExinvoicesEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetExinvoicesQueueEinvoice Yükleme kuyruğunu listeler
func (s *NESService) GetExinvoicesQueueEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/queue", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// GetExinvoicesQueueIdEinvoice Yükleme sonucunu indir
func (s *NESService) GetExinvoicesQueueIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/queue/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// GetTagsEinvoice Etiketleri listeler
func (s *NESService) GetTagsEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/tags", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostTagsEinvoice Etiket ekler
func (s *NESService) PostTagsEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/tags", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetTagsIdEinvoice Sorgulanan etiketi getirir
func (s *NESService) GetTagsIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/tags/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PutTagsIdEinvoice Etiket günceller
func (s *NESService) PutTagsIdEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/tags/%s", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// DeleteTagsIdEinvoice Etiket siler
func (s *NESService) DeleteTagsIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/tags/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// GetExinvoicesIncomingEinvoice Belgeleri listeler
func (s *NESService) GetExinvoicesIncomingEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/incoming", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// GetExinvoicesIncomingUuidXmlEinvoice XML İndir
func (s *NESService) GetExinvoicesIncomingUuidXmlEinvoice(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/incoming/%s/xml", s.getBaseURL(company), uuid)

	return s.doRequest(company, "GET", url, nil)
}


// GetExinvoicesIncomingUuidPdfEinvoice PDF İndir
func (s *NESService) GetExinvoicesIncomingUuidPdfEinvoice(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/incoming/%s/pdf", s.getBaseURL(company), uuid)

	return s.doRequest(company, "GET", url, nil)
}


// GetExinvoicesIncomingUuidHtmlEinvoice Belgeyi görüntüler
func (s *NESService) GetExinvoicesIncomingUuidHtmlEinvoice(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/incoming/%s/html", s.getBaseURL(company), uuid)

	return s.doRequest(company, "GET", url, nil)
}


// PostExinvoicesIncomingExportFiletypeEinvoice Dışarı Aktar
func (s *NESService) PostExinvoicesIncomingExportFiletypeEinvoice(company *models.Company, payload map[string]interface{}, fileType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/incoming/export/%s", s.getBaseURL(company), fileType)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetNotificationsIncomingDynamicrulesEinvoice Kuralları listeler
func (s *NESService) GetNotificationsIncomingDynamicrulesEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/incoming/dynamicrules", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostNotificationsIncomingDynamicrulesEinvoice Kural oluşturur
func (s *NESService) PostNotificationsIncomingDynamicrulesEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/incoming/dynamicrules", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetNotificationsIncomingDynamicrulesIdEinvoice Sorgulanan kuralı getirir
func (s *NESService) GetNotificationsIncomingDynamicrulesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/incoming/dynamicrules/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PutNotificationsIncomingDynamicrulesIdEinvoice Kural günceller
func (s *NESService) PutNotificationsIncomingDynamicrulesIdEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/incoming/dynamicrules/%s", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// DeleteNotificationsIncomingDynamicrulesIdEinvoice Kural siler
func (s *NESService) DeleteNotificationsIncomingDynamicrulesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/incoming/dynamicrules/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// PostIncomingInvoicesExportFiletypeEinvoice Toplu aktar
func (s *NESService) PostIncomingInvoicesExportFiletypeEinvoice(company *models.Company, payload map[string]interface{}, fileType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/export/%s", s.getBaseURL(company), fileType)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetIncomingReportmoduleReportsEinvoice Rapor listeler
func (s *NESService) GetIncomingReportmoduleReportsEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/reportmodule/reports", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostIncomingReportmoduleReportsEinvoice Rapor oluşturur
func (s *NESService) PostIncomingReportmoduleReportsEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/reportmodule/reports", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetIncomingReportmoduleReportsIdDownloadEinvoice Rapor indirir
func (s *NESService) GetIncomingReportmoduleReportsIdDownloadEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/reportmodule/reports/%s/download", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// GetIncomingReportmoduleTemplatesEinvoice Şablonları listeler
func (s *NESService) GetIncomingReportmoduleTemplatesEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/reportmodule/templates", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostIncomingReportmoduleTemplatesEinvoice Rapor şablonu oluşturur
func (s *NESService) PostIncomingReportmoduleTemplatesEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/reportmodule/templates", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetIncomingReportmoduleTemplatesIdEinvoice Sorgulanan şablonu getirir
func (s *NESService) GetIncomingReportmoduleTemplatesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/reportmodule/templates/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PutIncomingReportmoduleTemplatesIdEinvoice Rapor şablonunu günceller
func (s *NESService) PutIncomingReportmoduleTemplatesIdEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/reportmodule/templates/%s", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// DeleteIncomingReportmoduleTemplatesIdEinvoice Rapor Şablonunu siler
func (s *NESService) DeleteIncomingReportmoduleTemplatesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/reportmodule/templates/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// GetIncomingReportmoduleColumnsEinvoice Kolonları listeler
func (s *NESService) GetIncomingReportmoduleColumnsEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/reportmodule/columns", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PutIncomingInvoicesTagsEinvoice Etiket ekler/çıkarır
func (s *NESService) PutIncomingInvoicesTagsEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/tags", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// PostIncomingInvoicesUuidSavecompanyindocumentEinvoice Firma olarak kaydet
func (s *NESService) PostIncomingInvoicesUuidSavecompanyindocumentEinvoice(company *models.Company, payload map[string]interface{}, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/%s/savecompanyindocument", s.getBaseURL(company), uuid)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PostIncomingInvoicesUuidDocumentanswerEinvoice Belge'ye cevap verir
func (s *NESService) PostIncomingInvoicesUuidDocumentanswerEinvoice(company *models.Company, payload map[string]interface{}, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/%s/documentAnswer", s.getBaseURL(company), uuid)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PutIncomingInvoicesBulkOperationEinvoice Yeni durum atar
func (s *NESService) PutIncomingInvoicesBulkOperationEinvoice(company *models.Company, payload map[string]interface{}, operation string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/bulk/%s", s.getBaseURL(company), operation)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// PostIncomingInvoicesIdCreatereturninvoiceEinvoice Gelen e-Fatura için iade oluştur
func (s *NESService) PostIncomingInvoicesIdCreatereturninvoiceEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/%s/createreturninvoice", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PostIncomingInvoicesUuidUsernotesEinvoice Kullanıcı notu ekler
func (s *NESService) PostIncomingInvoicesUuidUsernotesEinvoice(company *models.Company, payload map[string]interface{}, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/%s/usernotes", s.getBaseURL(company), uuid)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PutIncomingInvoicesUuidUsernotesIdEinvoice Kullanıcı notunu günceller
func (s *NESService) PutIncomingInvoicesUuidUsernotesIdEinvoice(company *models.Company, payload map[string]interface{}, uuid string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/%s/usernotes/%s", s.getBaseURL(company), uuid, id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// DeleteIncomingInvoicesUuidUsernotesIdEinvoice Kullanıcı notunu siler
func (s *NESService) DeleteIncomingInvoicesUuidUsernotesIdEinvoice(company *models.Company, uuid string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/%s/usernotes/%s", s.getBaseURL(company), uuid, id)

	return s.doRequest(company, "DELETE", url, nil)
}


// PostIncomingInvoicesEmailSendEinvoice Belgeyi mail olarak iletir
func (s *NESService) PostIncomingInvoicesEmailSendEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/email/send", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetExinvoicesOutgoingEinvoice Belgeleri listeler
func (s *NESService) GetExinvoicesOutgoingEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/outgoing", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// GetExinvoicesOutgoingUuidXmlEinvoice XML İndir
func (s *NESService) GetExinvoicesOutgoingUuidXmlEinvoice(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/outgoing/%s/xml", s.getBaseURL(company), uuid)

	return s.doRequest(company, "GET", url, nil)
}


// GetExinvoicesOutgoingUuidPdfEinvoice PDF İndir
func (s *NESService) GetExinvoicesOutgoingUuidPdfEinvoice(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/outgoing/%s/pdf", s.getBaseURL(company), uuid)

	return s.doRequest(company, "GET", url, nil)
}


// GetExinvoicesOutgoingUuidHtmlEinvoice Belgeyi görüntüler
func (s *NESService) GetExinvoicesOutgoingUuidHtmlEinvoice(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/outgoing/%s/html", s.getBaseURL(company), uuid)

	return s.doRequest(company, "GET", url, nil)
}


// PostExinvoicesOutgoingExportFiletypeEinvoice Dışarı Aktar
func (s *NESService) PostExinvoicesOutgoingExportFiletypeEinvoice(company *models.Company, payload map[string]interface{}, fileType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/exinvoices/outgoing/export/%s", s.getBaseURL(company), fileType)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetNotificationsOutgoingDynamicrulesEinvoice Kuralları listeler
func (s *NESService) GetNotificationsOutgoingDynamicrulesEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/outgoing/dynamicrules", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostNotificationsOutgoingDynamicrulesEinvoice Kural oluşturur
func (s *NESService) PostNotificationsOutgoingDynamicrulesEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/outgoing/dynamicrules", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetNotificationsOutgoingDynamicrulesIdEinvoice Sorgulanan kuralı getirir
func (s *NESService) GetNotificationsOutgoingDynamicrulesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/outgoing/dynamicrules/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PutNotificationsOutgoingDynamicrulesIdEinvoice Kural günceller
func (s *NESService) PutNotificationsOutgoingDynamicrulesIdEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/outgoing/dynamicrules/%s", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// DeleteNotificationsOutgoingDynamicrulesIdEinvoice Kural siler
func (s *NESService) DeleteNotificationsOutgoingDynamicrulesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/notifications/outgoing/dynamicrules/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// PostOutgoingInvoicesExportFiletypeEinvoice Toplu aktar
func (s *NESService) PostOutgoingInvoicesExportFiletypeEinvoice(company *models.Company, payload map[string]interface{}, fileType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/export/%s", s.getBaseURL(company), fileType)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetOutgoingReportmoduleReportsEinvoice Rapor listeler
func (s *NESService) GetOutgoingReportmoduleReportsEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/reportmodule/reports", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostOutgoingReportmoduleReportsEinvoice Rapor oluşturur
func (s *NESService) PostOutgoingReportmoduleReportsEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/reportmodule/reports", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetOutgoingReportmoduleReportsIdDownloadEinvoice Rapor indirir
func (s *NESService) GetOutgoingReportmoduleReportsIdDownloadEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/reportmodule/reports/%s/download", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// GetOutgoingReportmoduleTemplatesEinvoice Şablonları listeler
func (s *NESService) GetOutgoingReportmoduleTemplatesEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/reportmodule/templates", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostOutgoingReportmoduleTemplatesEinvoice Rapor şablonu oluşturur
func (s *NESService) PostOutgoingReportmoduleTemplatesEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/reportmodule/templates", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetOutgoingReportmoduleTemplatesIdEinvoice Sorgulanan şablonu getirir
func (s *NESService) GetOutgoingReportmoduleTemplatesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/reportmodule/templates/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PutOutgoingReportmoduleTemplatesIdEinvoice Rapor şablonunu günceller
func (s *NESService) PutOutgoingReportmoduleTemplatesIdEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/reportmodule/templates/%s", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// DeleteOutgoingReportmoduleTemplatesIdEinvoice Rapor Şablonunu siler
func (s *NESService) DeleteOutgoingReportmoduleTemplatesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/reportmodule/templates/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// GetOutgoingReportmoduleColumnsEinvoice Kolonları listeler
func (s *NESService) GetOutgoingReportmoduleColumnsEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/reportmodule/columns", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PutOutgoingInvoicesTagsEinvoice Etiket ekler/çıkarır
func (s *NESService) PutOutgoingInvoicesTagsEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/tags", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// PutOutgoingInvoicesUuidReceiveraliasEinvoice Taslak belgelerin alıcı etiketi bu uç ile güncellenebilir
func (s *NESService) PutOutgoingInvoicesUuidReceiveraliasEinvoice(company *models.Company, payload map[string]interface{}, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/%s/receiveralias", s.getBaseURL(company), uuid)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// PostOutgoingInvoicesUuidSavecompanyindocumentEinvoice Firma olarak kaydet
func (s *NESService) PostOutgoingInvoicesUuidSavecompanyindocumentEinvoice(company *models.Company, payload map[string]interface{}, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/%s/savecompanyindocument", s.getBaseURL(company), uuid)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PutOutgoingInvoicesBulkOperationEinvoice Yeni durum atar
func (s *NESService) PutOutgoingInvoicesBulkOperationEinvoice(company *models.Company, payload map[string]interface{}, operation string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/bulk/%s", s.getBaseURL(company), operation)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// PostUploadsResendUuidEinvoice Hata almış bir belgeyi aynen yeniden gönderir
func (s *NESService) PostUploadsResendUuidEinvoice(company *models.Company, payload map[string]interface{}, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/uploads/resend/%s", s.getBaseURL(company), uuid)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PostOutgoingInvoicesUuidUsernotesEinvoice Kullanıcı notu ekler
func (s *NESService) PostOutgoingInvoicesUuidUsernotesEinvoice(company *models.Company, payload map[string]interface{}, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/%s/usernotes", s.getBaseURL(company), uuid)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// PutOutgoingInvoicesUuidUsernotesIdEinvoice Kullanıcı notunu günceller
func (s *NESService) PutOutgoingInvoicesUuidUsernotesIdEinvoice(company *models.Company, payload map[string]interface{}, uuid string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/%s/usernotes/%s", s.getBaseURL(company), uuid, id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// DeleteOutgoingInvoicesUuidUsernotesIdEinvoice Kullanıcı notunu siler
func (s *NESService) DeleteOutgoingInvoicesUuidUsernotesIdEinvoice(company *models.Company, uuid string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/%s/usernotes/%s", s.getBaseURL(company), uuid, id)

	return s.doRequest(company, "DELETE", url, nil)
}


// DeleteOutgoingInvoicesDraftsEinvoice Taslak belgeleri silmek için bu uç kullanılablir
func (s *NESService) DeleteOutgoingInvoicesDraftsEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/drafts", s.getBaseURL(company))

	return s.doRequest(company, "DELETE", url, nil)
}


// PostOutgoingInvoicesEmailSendEinvoice Belgeyi mail olarak iletir
func (s *NESService) PostOutgoingInvoicesEmailSendEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/outgoing/invoices/email/send", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetUsersZipAliastypeEinvoice Mükellef listesini indirir
func (s *NESService) GetUsersZipAliastypeEinvoice(company *models.Company, aliasType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/users/zip/%s", s.getBaseURL(company), aliasType)

	return s.doRequest(company, "GET", url, nil)
}


// GetUsersIdentifierAliastypeEinvoice Kimlik No ile sorgular
func (s *NESService) GetUsersIdentifierAliastypeEinvoice(company *models.Company, identifier string, aliasType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/users/%s/%s", s.getBaseURL(company), identifier, aliasType)

	return s.doRequest(company, "GET", url, nil)
}


// PostUsersAliastypeEinvoice Kimlik No ile sorgular
func (s *NESService) PostUsersAliastypeEinvoice(company *models.Company, payload map[string]interface{}, aliasType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/users/%s", s.getBaseURL(company), aliasType)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetUsersSearchQueryAliastypeEinvoice Ünvan ile sorgular
func (s *NESService) GetUsersSearchQueryAliastypeEinvoice(company *models.Company, query string, aliasType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/users/search/%s/%s", s.getBaseURL(company), query, aliasType)

	return s.doRequest(company, "GET", url, nil)
}


// GetDefinitionsDocumenttemplatesCustomizationsettingsEinvoice Tasarım ayarları dönülür
func (s *NESService) GetDefinitionsDocumenttemplatesCustomizationsettingsEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostDefinitionsDocumenttemplatesCustomizationsettingsEinvoice e-Belge özelleştirilebilir tasarım eklemek için kullanılır.
func (s *NESService) PostDefinitionsDocumenttemplatesCustomizationsettingsEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice Sorgulanan ayarı getirir
func (s *NESService) GetDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PutDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice e-Belge özelleştirilebilir tasarımını güncellemek için kullanılır.
func (s *NESService) PutDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice e-Belge özelleştirilebilir tasarımını silmek için kullanılır.
func (s *NESService) DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// GetDefinitionsDocumenttemplatesCustomizationsettingsIdSetdefaultEinvoice No description provided
func (s *NESService) GetDefinitionsDocumenttemplatesCustomizationsettingsIdSetdefaultEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/setdefault", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PostDefinitionsDocumenttemplatesCustomizationsettingsIdPreviewEinvoice Tasarımı önizler
func (s *NESService) PostDefinitionsDocumenttemplatesCustomizationsettingsIdPreviewEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/preview", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice e-Belge özelleştirilebilir tasarıma eklenmiş olan logoya bu uç ile ulaşılabilir.
func (s *NESService) GetDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/logo", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PostDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice e-Belge özelleştirilebilir tasarıma logo eklemek için bu uç kullanılabilir.
func (s *NESService) PostDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/logo", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice e-Belge özelleştirilebilir tasarıma eklenmiş olan logoyu silmek için bu uç kullanılabilir.
func (s *NESService) DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdLogoEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/logo", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// GetDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeye bu uç ile ulaşılabilir.
func (s *NESService) GetDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/stamp", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PostDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice e-Belge özelleştirilebilir tasarıma kaşe eklemek için bu uç kullanılabilir.
func (s *NESService) PostDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/stamp", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice e-Belge özelleştirilebilir tasarıma eklenmiş olan kaşeyi silmek için bu uç kullanılabilir.
func (s *NESService) DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdStampEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/stamp", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// GetDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice e-Belge özelleştirilebilir tasarıma eklenmiş olan imzaya bu uç ile ulaşılabilir.
func (s *NESService) GetDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/signature", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PostDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice e-Belge özelleştirilebilir tasarıma imza eklemek için bu uç kullanılabilir.
func (s *NESService) PostDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/signature", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice e-Belge özelleştirilebilir tasarıma eklenmiş olan imzayı silmek için bu uç kullanılabilir.
func (s *NESService) DeleteDefinitionsDocumenttemplatesCustomizationsettingsIdSignatureEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/customizationsettings/%s/signature", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// GetDefinitionsSeriesEinvoice Serileri listeler
func (s *NESService) GetDefinitionsSeriesEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/series", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostDefinitionsSeriesEinvoice Seri ekler
func (s *NESService) PostDefinitionsSeriesEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/series", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetDefinitionsSeriesIdEinvoice Sorgulanan seriyi getirir
func (s *NESService) GetDefinitionsSeriesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/series/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// DeleteDefinitionsSeriesIdEinvoice Seri siler
func (s *NESService) DeleteDefinitionsSeriesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/series/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// GetDefinitionsSeriesSerieEinvoice Ön eke göre seriyi getirir
func (s *NESService) GetDefinitionsSeriesSerieEinvoice(company *models.Company, serie string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/series/%s", s.getBaseURL(company), serie)

	return s.doRequest(company, "GET", url, nil)
}


// GetDefinitionsSeriesIdSetStatusEinvoice Seri durumunu günceller
func (s *NESService) GetDefinitionsSeriesIdSetStatusEinvoice(company *models.Company, id string, status string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/series/%s/set/%s", s.getBaseURL(company), id, status)

	return s.doRequest(company, "GET", url, nil)
}


// GetDefinitionsSeriesIdSetdefaultEinvoice Seriyi varsayılan ayarlar
func (s *NESService) GetDefinitionsSeriesIdSetdefaultEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/series/%s/setdefault", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// GetDefinitionsSeriesIdYearSetnumberNextnumberEinvoice Sayaç günceller
func (s *NESService) GetDefinitionsSeriesIdYearSetnumberNextnumberEinvoice(company *models.Company, id string, year string, nextNumber string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/series/%s/%s/setnumber/%s", s.getBaseURL(company), id, year, nextNumber)

	return s.doRequest(company, "GET", url, nil)
}


// GetDefinitionsSeriesSerieidYearHistoriesEinvoice Sayaç geçmişi
func (s *NESService) GetDefinitionsSeriesSerieidYearHistoriesEinvoice(company *models.Company, serieId string, year string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/series/%s/%s/histories", s.getBaseURL(company), serieId, year)

	return s.doRequest(company, "GET", url, nil)
}


// GetDefinitionsDocumenttemplatesEinvoice Tasarımları listeler
func (s *NESService) GetDefinitionsDocumenttemplatesEinvoice(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates", s.getBaseURL(company))

	return s.doRequest(company, "GET", url, nil)
}


// PostDefinitionsDocumenttemplatesEinvoice Tasarım ekler
func (s *NESService) PostDefinitionsDocumenttemplatesEinvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates", s.getBaseURL(company))

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetDefinitionsDocumenttemplatesIdEinvoice Tasarım dosyasını indirir
func (s *NESService) GetDefinitionsDocumenttemplatesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PutDefinitionsDocumenttemplatesIdEinvoice Tasarımı günceller
func (s *NESService) PutDefinitionsDocumenttemplatesIdEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/%s", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(jsonData))
}


// DeleteDefinitionsDocumenttemplatesIdEinvoice Tasarımı siler
func (s *NESService) DeleteDefinitionsDocumenttemplatesIdEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/%s", s.getBaseURL(company), id)

	return s.doRequest(company, "DELETE", url, nil)
}


// GetDefinitionsDocumenttemplatesIdSetdefaultEinvoice Tasarımı varsayılan ayarlar
func (s *NESService) GetDefinitionsDocumenttemplatesIdSetdefaultEinvoice(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/%s/setdefault", s.getBaseURL(company), id)

	return s.doRequest(company, "GET", url, nil)
}


// PostDefinitionsDocumenttemplatesIdPreviewEinvoice Tasarımı önizler
func (s *NESService) PostDefinitionsDocumenttemplatesIdPreviewEinvoice(company *models.Company, payload map[string]interface{}, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/definitions/documenttemplates/%s/preview", s.getBaseURL(company), id)

	jsonData, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(jsonData))
}


// GetEnvelopesInstanceidentifierQueryEinvoice Zarf Durum Sorgular
func (s *NESService) GetEnvelopesInstanceidentifierQueryEinvoice(company *models.Company, instanceIdentifier string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/envelopes/%s/query", s.getBaseURL(company), instanceIdentifier)

	return s.doRequest(company, "GET", url, nil)
}
