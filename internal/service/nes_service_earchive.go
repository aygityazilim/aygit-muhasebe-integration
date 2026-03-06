package service

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/errors"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// --- Notifications / Dynamic Rules ---

func (s *NESService) GetEArchiveDynamicRules(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/notifications/dynamicrules", s.getBaseURL(company))
	// Add query params if needed
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveDynamicRule(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/notifications/dynamicrules", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) GetEArchiveDynamicRule(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/notifications/dynamicrules/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) UpdateEArchiveDynamicRule(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/notifications/dynamicrules/%s", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveDynamicRule(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/notifications/dynamicrules/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// --- Invoices Export ---

func (s *NESService) ExportEArchiveInvoices(company *models.Company, fileType string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/export/%s", s.getBaseURL(company), fileType)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// --- Definitions ---

func (s *NESService) GetEArchiveFileExportTitles(company *models.Company, documentType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/fileexporttitles/%s/titlekeys", s.getBaseURL(company), documentType)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) GetEArchiveFileExportTitlesExtension(company *models.Company, documentType string, extension string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/fileexporttitles/%s/%s", s.getBaseURL(company), documentType, extension)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) UpdateEArchiveFileExportTitles(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/fileexporttitles", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// --- Uploads ---

func (s *NESService) PreviewEArchiveDocument(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/uploads/document/preview", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) UpdateEArchiveDocument(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/uploads/document/%s", s.getBaseURL(company), uuid)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) CreateEArchiveDraft(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/uploads/draft/create/%s", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) PreviewEArchiveMarketplaceOrder(company *models.Company, id string, orderId string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/uploads/marketplaces/%s/orders/%s/preview", s.getBaseURL(company), id, orderId)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) CreateEArchiveMarketplaceInvoice(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/uploads/marketplaces/%s/orders/createinvoice", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// --- ExInvoices ---

func (s *NESService) GetEArchiveExInvoices(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/exinvoices", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) UploadEArchiveExInvoice(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/exinvoices", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) GetEArchiveExInvoicesQueue(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/exinvoices/queue", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) GetEArchiveExInvoicesQueueResult(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/exinvoices/queue/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) DownloadEArchiveExInvoiceFile(company *models.Company, uuid string, fileType string) ([]byte, string, error) {
	urlPath := fmt.Sprintf("%s/earchive/v1/exinvoices/%s/%s", s.getBaseURL(company), uuid, fileType)
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, "", fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, "", errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("dosya okuma hatası: %w", err)
	}
	return content, resp.Header.Get("Content-Type"), nil
}

func (s *NESService) ExportEArchiveExInvoices(company *models.Company, fileType string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/exinvoices/export/%s", s.getBaseURL(company), fileType)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// --- Tags ---

func (s *NESService) GetEArchiveTags(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/tags", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveTag(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/tags", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) GetEArchiveTag(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/tags/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) UpdateEArchiveTag(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/tags/%s", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveTag(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/tags/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// --- Report Module ---

func (s *NESService) GetEArchiveReports(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/outgoing/reportmodule/reports", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveReport(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/outgoing/reportmodule/reports", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) DownloadEArchiveReport(company *models.Company, id string) ([]byte, string, error) {
	urlPath := fmt.Sprintf("%s/earchive/v1/outgoing/reportmodule/reports/%s/download", s.getBaseURL(company), id)
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, "", fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, "", errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("dosya okuma hatası: %w", err)
	}
	return content, resp.Header.Get("Content-Type"), nil
}

func (s *NESService) GetEArchiveReportTemplates(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/outgoing/reportmodule/templates", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveReportTemplate(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/outgoing/reportmodule/templates", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) GetEArchiveReportTemplate(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/outgoing/reportmodule/templates/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) UpdateEArchiveReportTemplate(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/outgoing/reportmodule/templates/%s", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveReportTemplate(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/outgoing/reportmodule/templates/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

func (s *NESService) GetEArchiveReportColumns(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/outgoing/reportmodule/columns", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// --- Invoices Specific ---

func (s *NESService) UpdateEArchiveInvoiceTags(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/tags", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) SaveCompanyInEArchiveDocument(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/%s/savecompanyindocument", s.getBaseURL(company), uuid)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) BulkEArchiveInvoiceOperation(company *models.Company, operation string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/bulk/%s", s.getBaseURL(company), operation)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveDraftInvoices(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/drafts", s.getBaseURL(company))
	return s.doRequest(company, "DELETE", url, nil)
}

func (s *NESService) SendEArchiveInvoiceEmail(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/email/send", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) GetEArchiveInvoiceUserNotes(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/%s/usernotes", s.getBaseURL(company), uuid)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveInvoiceUserNote(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/%s/usernotes", s.getBaseURL(company), uuid)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) UpdateEArchiveInvoiceUserNote(company *models.Company, uuid string, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/%s/usernotes/%s", s.getBaseURL(company), uuid, id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveInvoiceUserNote(company *models.Company, uuid string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/%s/usernotes/%s", s.getBaseURL(company), uuid, id)
	return s.doRequest(company, "DELETE", url, nil)
}

// --- Email / SMS Settings ---

func (s *NESService) GetEArchiveEmailSettings(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/mailing/email/settings", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) UpdateEArchiveEmailSettings(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/mailing/email/settings", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) GetEArchiveSmsSettings(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/mailing/sms/settings", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) UpdateEArchiveSmsSettings(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/mailing/sms/settings", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// --- Customization Settings ---

func (s *NESService) GetEArchiveCustomizationSettings(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveCustomizationSetting(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) GetEArchiveCustomizationSetting(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) UpdateEArchiveCustomizationSetting(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveCustomizationSetting(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

func (s *NESService) SetEArchiveCustomizationSettingDefault(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/setdefault", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) PreviewEArchiveCustomizationSetting(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/preview", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) GetEArchiveCustomizationSettingLogo(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/logo", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveCustomizationSettingLogo(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/logo", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveCustomizationSettingLogo(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/logo", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

func (s *NESService) GetEArchiveCustomizationSettingStamp(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/stamp", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveCustomizationSettingStamp(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/stamp", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveCustomizationSettingStamp(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/stamp", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

func (s *NESService) GetEArchiveCustomizationSettingSignature(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/signature", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveCustomizationSettingSignature(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/signature", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveCustomizationSettingSignature(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/customizationsettings/%s/signature", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// --- Series ---

func (s *NESService) GetEArchiveSeries(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/series", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveSerie(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/series", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) GetEArchiveSerie(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/series/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) DeleteEArchiveSerie(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/series/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

func (s *NESService) GetEArchiveSerieByPrefix(company *models.Company, serie string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/series/%s", s.getBaseURL(company), serie)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) SetEArchiveSerieStatus(company *models.Company, id string, status string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/series/%s/set/%s", s.getBaseURL(company), id, status)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) SetEArchiveSerieDefault(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/series/%s/setdefault", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) SetEArchiveSerieNumber(company *models.Company, id string, year string, nextNumber string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/series/%s/%s/setnumber/%s", s.getBaseURL(company), id, year, nextNumber)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) GetEArchiveSerieHistories(company *models.Company, serieId string, year string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/series/%s/%s/histories", s.getBaseURL(company), serieId, year)
	return s.doRequest(company, "GET", url, nil)
}

// --- Document Templates ---

func (s *NESService) GetEArchiveDocumentTemplates(company *models.Company, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) CreateEArchiveDocumentTemplate(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates", s.getBaseURL(company))
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

func (s *NESService) DownloadEArchiveDocumentTemplate(company *models.Company, id string) ([]byte, string, error) {
	urlPath := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/%s", s.getBaseURL(company), id)
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, "", fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, "", errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("dosya okuma hatası: %w", err)
	}
	return content, resp.Header.Get("Content-Type"), nil
}

func (s *NESService) UpdateEArchiveDocumentTemplate(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/%s", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

func (s *NESService) DeleteEArchiveDocumentTemplate(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

func (s *NESService) SetEArchiveDocumentTemplateDefault(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/%s/setdefault", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

func (s *NESService) PreviewEArchiveDocumentTemplate(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/definitions/documenttemplates/%s/preview", s.getBaseURL(company), id)
	data, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}
