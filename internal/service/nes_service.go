package service

import (
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/errors"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"
)

// NESService, NES Özel Entegratör API'si ile iletişim kuran ana servis yapısıdır.
type NESService struct {
	Client *http.Client
}

// NewNESService, varsayılan zaman aşımı değerleri ile yeni bir NESService örneği döner.
func NewNESService() *NESService {
	// Test ortamında kendinden imzalı sertifikalara izin vermek için TLS ayarı yapıyoruz
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &NESService{
		Client: &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
		},
	}
}

// getBaseURL, veritabanından alınan firma konfigürasyonuna göre ortamın (Test/Prod) ana URL'sini döner.
func (s *NESService) getBaseURL(company *models.Company) string {
	baseURL := os.Getenv("NES_TEST_URL")
	if company.Environment == "PRODUCTION" {
		baseURL = os.Getenv("NES_PROD_URL")
	}
	if baseURL == "" {
		baseURL = "https://apitest.nes.com.tr"
	}
	return baseURL
}

// GetCreditSummary, belirtilen firmanın NES üzerindeki kontör bakiyesini sorgular.
func (s *NESService) GetCreditSummary(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/management/creditsummary", s.getBaseURL(company))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

// GetAccountDefaultDocumentParameter, firmanın NES üzerindeki varsayılan doküman parametrelerini sorgular.
func (s *NESService) GetAccountDefaultDocumentParameter(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/accountDefaultDocumentParameter", s.getBaseURL(company))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

// --- E-Fatura (Invoices) ---

func (s *NESService) GetIncomingInvoices(company *models.Company, queryParams map[string]string) (map[string]interface{}, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/einvoice/v1/incoming/invoices", s.getBaseURL(company)))
	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

func (s *NESService) GetOutgoingInvoices(company *models.Company, queryParams map[string]string) (map[string]interface{}, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/einvoice/v1/outgoing/invoices", s.getBaseURL(company)))
	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

func (s *NESService) DownloadInvoiceFile(company *models.Company, uuid string, direction string, fileType string) ([]byte, string, error) {
	urlPath := fmt.Sprintf("%s/einvoice/v1/%s/invoices/%s/%s", s.getBaseURL(company), direction, uuid, fileType)
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, "", fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}

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

func (s *NESService) UploadInvoice(company *models.Company, xmlData []byte, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/uploads/document", s.getBaseURL(company))
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("File", "invoice.xml")
	if err != nil {
		return nil, fmt.Errorf("form dosyası oluşturma hatası: %w", err)
	}
	if _, err := part.Write(xmlData); err != nil {
		return nil, fmt.Errorf("dosya yazma hatası: %w", err)
	}

	for k, v := range params {
		writer.WriteField(k, v)
	}
	if params["IsDirectSend"] == "" {
		writer.WriteField("IsDirectSend", "false")
	}
	if params["AutoSaveCompany"] == "" {
		writer.WriteField("AutoSaveCompany", "true")
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	fmt.Println("Authorization: ", "Bearer "+company.GetNesAPIKey())
	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, errors.FormatNESError(resp.StatusCode, respBody)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

func (s *NESService) SendDraftInvoices(company *models.Company, uuids []string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/einvoice/v1/uploads/draft/send", s.getBaseURL(company))
	jsonData, _ := json.Marshal(uuids)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	var result []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

// --- E-İrsaliye (Despatches) ---

func (s *NESService) GetIncomingDespatches(company *models.Company, queryParams map[string]string) (map[string]interface{}, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/edespatch/v1/incoming/despatches", s.getBaseURL(company)))
	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, errors.FormatNESError(resp.StatusCode, respBody)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

func (s *NESService) GetOutgoingDespatches(company *models.Company, queryParams map[string]string) (map[string]interface{}, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/edespatch/v1/outgoing/despatches", s.getBaseURL(company)))
	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, errors.FormatNESError(resp.StatusCode, respBody)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

func (s *NESService) DownloadDespatchFile(company *models.Company, uuid string, direction string, fileType string) ([]byte, string, error) {
	urlPath := fmt.Sprintf("%s/edespatch/v1/%s/despatches/%s/%s", s.getBaseURL(company), direction, uuid, fileType)
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

func (s *NESService) UploadDespatch(company *models.Company, xmlData []byte, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/uploads/document", s.getBaseURL(company))
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("File", "despatch.xml")
	if err != nil {
		return nil, fmt.Errorf("form dosyası oluşturma hatası: %w", err)
	}
	if _, err := part.Write(xmlData); err != nil {
		return nil, fmt.Errorf("dosya yazma hatası: %w", err)
	}

	for k, v := range params {
		writer.WriteField(k, v)
	}
	if params["IsDirectSend"] == "" {
		writer.WriteField("IsDirectSend", "false")
	}
	if params["AutoSaveCompany"] == "" {
		writer.WriteField("AutoSaveCompany", "true")
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, errors.FormatNESError(resp.StatusCode, respBody)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

func (s *NESService) SendDraftDespatches(company *models.Company, uuids []string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/uploads/draft/send", s.getBaseURL(company))
	jsonData, _ := json.Marshal(uuids)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	var result []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

func (s *NESService) RejectEInvoice(company *models.Company, uuid string, reason string) (map[string]interface{}, error) {
	// PostIncomingInvoicesRejectEinvoice /einvoice/v1/incoming/invoices/reject
	url := fmt.Sprintf("%s/einvoice/v1/incoming/invoices/reject", s.getBaseURL(company))
	payload := map[string]interface{}{
		"uuids":  []string{uuid},
		"reason": reason,
	}
	jsonData, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

func (s *NESService) SendDespatchAnswer(company *models.Company, uuid string, answer map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/%s/answer", s.getBaseURL(company), uuid)
	jsonData, _ := json.Marshal(answer)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, errors.FormatNESError(resp.StatusCode, respBody)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

// --- E-Arşiv (E-Archive) ---

func (s *NESService) GetEArchiveInvoices(company *models.Company, queryParams map[string]string) (map[string]interface{}, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/earchive/v1/invoices", s.getBaseURL(company)))
	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, errors.FormatNESError(resp.StatusCode, respBody)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

func (s *NESService) DownloadEArchiveFile(company *models.Company, uuid string, fileType string) ([]byte, string, error) {
	urlPath := fmt.Sprintf("%s/earchive/v1/invoices/%s/%s", s.getBaseURL(company), uuid, fileType)
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

func (s *NESService) CancelEArchiveInvoice(company *models.Company, cancelData map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/cancel", s.getBaseURL(company))
	jsonData, _ := json.Marshal(cancelData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, errors.FormatNESError(resp.StatusCode, respBody)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

func (s *NESService) UploadEArchiveInvoice(company *models.Company, xmlData []byte, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/uploads/document", s.getBaseURL(company))
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("File", "invoice.xml")
	if err != nil {
		return nil, fmt.Errorf("form dosyası oluşturma hatası: %w", err)
	}
	if _, err := part.Write(xmlData); err != nil {
		return nil, fmt.Errorf("dosya yazma hatası: %w", err)
	}

	for k, v := range params {
		writer.WriteField(k, v)
	}
	if params["IsDirectSend"] == "" {
		writer.WriteField("IsDirectSend", "false")
	}
	if params["AutoSaveCompany"] == "" {
		writer.WriteField("AutoSaveCompany", "true")
	}
	if params["PreviewType"] == "" {
		writer.WriteField("PreviewType", "None")
	}
	if params["SourceApp"] == "" {
		writer.WriteField("SourceApp", "AygitIntegration")
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, errors.FormatNESError(resp.StatusCode, respBody)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

func (s *NESService) SendDraftEArchiveInvoices(company *models.Company, uuids []string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/invoices/bulk/send", s.getBaseURL(company))
	jsonData, _ := json.Marshal(uuids)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, errors.FormatNESError(resp.StatusCode, respBody)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

// --- Static Codes & Statistics ---

func (s *NESService) GetTaxTypes(company *models.Company) ([]models.TaxType, error) {
	url := fmt.Sprintf("%s/general/v1/staticcodes/taxtype", s.getBaseURL(company))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	var result []models.TaxType
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

func (s *NESService) GetWithholdingTaxTypes(company *models.Company) ([]models.WithholdingTaxType, error) {
	url := fmt.Sprintf("%s/general/v1/staticcodes/withholdingtaxtype", s.getBaseURL(company))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	var result []models.WithholdingTaxType
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

func (s *NESService) GetTaxExemptionReasons(company *models.Company) ([]models.TaxExemptionReason, error) {
	url := fmt.Sprintf("%s/general/v1/staticcodes/taxexemptionreason", s.getBaseURL(company))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	var result []models.TaxExemptionReason
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

func (s *NESService) GetCurrencies(company *models.Company) ([]models.Currency, error) {
	url := fmt.Sprintf("%s/general/v1/staticcodes/currency", s.getBaseURL(company))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	var result []models.Currency
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

func (s *NESService) GetDailyStatistics(company *models.Company, startDate string, endDate string) ([]models.DailyStatistic, error) {
	url := fmt.Sprintf("%s/general/v1/statistics/daily?startDate=%s&endDate=%s", s.getBaseURL(company), startDate, endDate)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	var result []models.DailyStatistic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

// GetAccountDocumentArchives, firmanın NES üzerindeki belge arşivi taleplerini listeler.
func (s *NESService) GetAccountDocumentArchives(company *models.Company) ([]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/accountDocumentArchives", s.getBaseURL(company))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	var result []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}

// helper function to perform requests
func (s *NESService) doRequest(company *models.Company, method, apiURL string, body io.Reader) (map[string]interface{}, error) {
	req, err := http.NewRequest(method, apiURL, body)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturma hatası: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, errors.FormatNESError(resp.StatusCode, bodyBytes)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		if err == io.EOF { // Empty response is OK for some endpoints like DELETE
			return map[string]interface{}{"status": "success"}, nil
		}
		return nil, fmt.Errorf("JSON parse hatası: %w", err)
	}

	return result, nil
}

// UpdateAccountDefaultDocumentParameter
func (s *NESService) UpdateAccountDefaultDocumentParameter(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/accountDefaultDocumentParameter", s.getBaseURL(company))
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// CreateAccountDocumentArchive
func (s *NESService) CreateAccountDocumentArchive(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/documentarchives", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// AccountModulesInfo
func (s *NESService) AccountModulesInfo(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/accountmodules/info", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// AccountModulesUpdate
func (s *NESService) AccountModulesUpdate(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/accountmodules/update", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// AccountModulesGetEnvelopeContent
func (s *NESService) AccountModulesGetEnvelopeContent(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/accountmodules/getenvelopecontent", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// AccountModulesSetSignedContent
func (s *NESService) AccountModulesSetSignedContent(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/accountmodules/setsignedcontent", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// AccountModulesGetEnvelopeInfo
func (s *NESService) AccountModulesGetEnvelopeInfo(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/accountmodules/getenvelopeinfo", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// DeleteIdentification
func (s *NESService) DeleteIdentification(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/management/identifications/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// GetDealerInfo
func (s *NESService) GetDealerInfo(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/management/dealerinfo", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// GetAddresses
func (s *NESService) GetAddresses(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/management/addresses", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateAddress
func (s *NESService) CreateAddress(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/management/addresses", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// UpdateAddress
func (s *NESService) UpdateAddress(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/management/address/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// DeleteAddress
func (s *NESService) DeleteAddress(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/management/address/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// GetIVD
func (s *NESService) GetIVD(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/management/ivd", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// UpdateIVD
func (s *NESService) UpdateIVD(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/management/ivd", s.getBaseURL(company))
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// GetLucaIntegrationSetting
func (s *NESService) GetLucaIntegrationSetting(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/management/luca/integration/setting", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// UpdateLucaIntegrationSetting
func (s *NESService) UpdateLucaIntegrationSetting(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/management/luca/integration/setting", s.getBaseURL(company))
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// GetCustomerSearchFromGIBSetting
func (s *NESService) GetCustomerSearchFromGIBSetting(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/management/customer_search_from_gib/setting", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// UpdateCustomerSearchFromGIBSetting
func (s *NESService) UpdateCustomerSearchFromGIBSetting(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/management/customer_search_from_gib/setting", s.getBaseURL(company))
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// GetIdentifications
func (s *NESService) GetIdentifications(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/general/v1/management/identifications", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateIdentification
func (s *NESService) CreateIdentification(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/general/v1/management/identifications", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// ---------------------------------------------------------------------------------------------------------------------
// DEFINITIONS - CUSTOMIZATION SETTINGS
// ---------------------------------------------------------------------------------------------------------------------

// GetCustomizationSettings returns the customization settings.
func (s *NESService) GetCustomizationSettings(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateCustomizationSetting creates a customization setting.
func (s *NESService) CreateCustomizationSetting(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings", s.getBaseURL(company))
	body, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(body))
}

// GetCustomizationSettingByID gets a specific customization setting.
func (s *NESService) GetCustomizationSettingByID(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

// UpdateCustomizationSetting updates a specific customization setting.
func (s *NESService) UpdateCustomizationSetting(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings/%s", s.getBaseURL(company), id)
	body, _ := json.Marshal(payload)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(body))
}

// DeleteCustomizationSetting deletes a specific customization setting.
func (s *NESService) DeleteCustomizationSetting(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// SetDefaultCustomizationSetting sets the customization setting as default.
func (s *NESService) SetDefaultCustomizationSetting(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings/%s/setdefault", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

// PreviewCustomizationSetting previews the customization setting.
func (s *NESService) PreviewCustomizationSetting(company *models.Company, id string, payload map[string]interface{}) ([]byte, string, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings/%s/preview", s.getBaseURL(company), id)
	body, _ := json.Marshal(payload)
	return s.downloadFileHelper(company, "POST", url, bytes.NewBuffer(body))
}

func (s *NESService) GetCustomizationSettingImage(company *models.Company, id string, imageType string) ([]byte, string, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings/%s/%s", s.getBaseURL(company), id, imageType)
	return s.downloadFileHelper(company, "GET", url, nil)
}

func (s *NESService) DeleteCustomizationSettingImage(company *models.Company, id string, imageType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings/%s/%s", s.getBaseURL(company), id, imageType)
	return s.doRequest(company, "DELETE", url, nil)
}

func (s *NESService) UploadCustomizationSettingImage(company *models.Company, id string, imageType string, fileData []byte, filename string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/definitions/documenttemplates/customizationsettings/%s/%s", s.getBaseURL(company), id, imageType)
	return s.doMultipartRequest("POST", url, nil, fileData, filename, company)
}

// ---------------------------------------------------------------------------------------------------------------------
// DEFINITIONS - SERIES & ANSWER SERIES
// ---------------------------------------------------------------------------------------------------------------------

// GetSeries lists series for a type (definitions/series or answerseries).
func (s *NESService) GetSeries(company *models.Company, seriesType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s", s.getBaseURL(company), seriesType)
	return s.doRequest(company, "GET", url, nil)
}

// CreateSeries creates a series.
func (s *NESService) CreateSeries(company *models.Company, seriesType string, payload map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s", s.getBaseURL(company), seriesType)
	body, _ := json.Marshal(payload)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(body))
}

// GetSeriesByID gets a specific series.
func (s *NESService) GetSeriesByID(company *models.Company, seriesType string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s/%s", s.getBaseURL(company), seriesType, id)
	return s.doRequest(company, "GET", url, nil)
}

// DeleteSeries deletes a specific series.
func (s *NESService) DeleteSeries(company *models.Company, seriesType string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s/%s", s.getBaseURL(company), seriesType, id)
	return s.doRequest(company, "DELETE", url, nil)
}

// GetSeriesByPrefix gets a series by its prefix.
func (s *NESService) GetSeriesByPrefix(company *models.Company, seriesType string, prefix string) (map[string]interface{}, error) {
	// The framework routing might overlap ID and Prefix, but we assume the API path is the same.
	url := fmt.Sprintf("%s/v1/%s/%s", s.getBaseURL(company), seriesType, prefix)
	return s.doRequest(company, "GET", url, nil)
}

// SetSeriesStatus updates series status.
func (s *NESService) SetSeriesStatus(company *models.Company, seriesType string, id string, status string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s/%s/set/%s", s.getBaseURL(company), seriesType, id, status)
	return s.doRequest(company, "GET", url, nil)
}

// SetDefaultSeries sets series as default.
func (s *NESService) SetDefaultSeries(company *models.Company, seriesType string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s/%s/setdefault", s.getBaseURL(company), seriesType, id)
	return s.doRequest(company, "GET", url, nil)
}

// SetSeriesNextNumber updates series next number.
func (s *NESService) SetSeriesNextNumber(company *models.Company, seriesType string, id string, year string, nextNumber string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s/%s/%s/setnumber/%s", s.getBaseURL(company), seriesType, id, year, nextNumber)
	return s.doRequest(company, "GET", url, nil)
}

// GetSeriesHistories lists series histories.
func (s *NESService) GetSeriesHistories(company *models.Company, seriesType string, id string, year string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s/%s/%s/histories", s.getBaseURL(company), seriesType, id, year)
	return s.doRequest(company, "GET", url, nil)
}

// ---------------------------------------------------------------------------------------------------------------------
// DEFINITIONS - DOCUMENT TEMPLATES & ANSWER DOCUMENT TEMPLATES
// ---------------------------------------------------------------------------------------------------------------------

// GetDocumentTemplates lists templates.
func (s *NESService) GetDocumentTemplates(company *models.Company, templateType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s", s.getBaseURL(company), templateType)
	return s.doRequest(company, "GET", url, nil)
}

// CreateDocumentTemplate adds a template.
func (s *NESService) CreateDocumentTemplate(company *models.Company, templateType string, fileData []byte, filename string, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s", s.getBaseURL(company), templateType)
	return s.doMultipartRequest("POST", url, params, fileData, filename, company)
}

// DownloadDocumentTemplate downloads template file.
func (s *NESService) DownloadDocumentTemplate(company *models.Company, templateType string, id string) ([]byte, string, error) {
	url := fmt.Sprintf("%s/v1/%s/%s", s.getBaseURL(company), templateType, id)
	return s.downloadFileHelper(company, "GET", url, nil)
}

// UpdateDocumentTemplate updates template file.
func (s *NESService) UpdateDocumentTemplate(company *models.Company, templateType string, id string, fileData []byte, filename string, params map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s/%s", s.getBaseURL(company), templateType, id)
	return s.doMultipartRequest("PUT", url, params, fileData, filename, company)
}

// DeleteDocumentTemplate deletes a template.
func (s *NESService) DeleteDocumentTemplate(company *models.Company, templateType string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s/%s", s.getBaseURL(company), templateType, id)
	return s.doRequest(company, "DELETE", url, nil)
}

// SetDefaultDocumentTemplate sets a template as default.
func (s *NESService) SetDefaultDocumentTemplate(company *models.Company, templateType string, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/%s/%s/setdefault", s.getBaseURL(company), templateType, id)
	return s.doRequest(company, "GET", url, nil)
}

// PreviewDocumentTemplate previews a template.
func (s *NESService) PreviewDocumentTemplate(company *models.Company, templateType string, id string, payload map[string]interface{}) ([]byte, string, error) {
	url := fmt.Sprintf("%s/v1/%s/%s/preview", s.getBaseURL(company), templateType, id)
	body, _ := json.Marshal(payload)
	return s.downloadFileHelper(company, "POST", url, bytes.NewBuffer(body))
}

// ---------------------------------------------------------------------------------------------------------------------
// ENVELOPES
// ---------------------------------------------------------------------------------------------------------------------

// QueryEnvelopeStatus queries the envelope status.
func (s *NESService) QueryEnvelopeStatus(company *models.Company, instanceIdentifier string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/v1/envelopes/%s/query", s.getBaseURL(company), instanceIdentifier)
	return s.doRequest(company, "GET", url, nil)
}

// downloadFileHelper makes a request that expects binary/file data.
func (s *NESService) downloadFileHelper(company *models.Company, method string, url string, body io.Reader) ([]byte, string, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, "", errors.NewError(500, errors.ErrCodeInternalServer, "İstek oluşturulamadı")
	}

	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", errors.NewError(500, errors.ErrCodeIntegrationFailed, "NES Servisine bağlanılamadı")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, "", errors.NewError(resp.StatusCode, errors.ErrCodeIntegrationFailed, string(respBody))
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", errors.NewError(500, errors.ErrCodeInternalServer, "Dosya içeriği okunamadı")
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return content, contentType, nil
}

// doMultipartRequest performs a multipart/form-data request.
func (s *NESService) doMultipartRequest(method string, urlStr string, params map[string]string, fileData []byte, filename string, company *models.Company) (map[string]interface{}, error) {
	if company == nil {
		return nil, errors.NewError(500, errors.ErrCodeInternalServer, "Firma bilgisi gerekli")
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file
	if len(fileData) > 0 {
		part, err := writer.CreateFormFile("File", filename)
		if err != nil {
			return nil, errors.NewError(500, errors.ErrCodeInternalServer, "Form dosyası oluşturulamadı")
		}
		part.Write(fileData)
	}

	// Add params
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err := writer.Close()
	if err != nil {
		return nil, errors.NewError(500, errors.ErrCodeInternalServer, "Form verisi kapatılamadı")
	}

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, errors.NewError(500, errors.ErrCodeInternalServer, "İstek oluşturulamadı")
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.NewError(500, errors.ErrCodeIntegrationFailed, "NES Servisine bağlanılamadı")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, errors.NewError(resp.StatusCode, errors.ErrCodeIntegrationFailed, string(respBody))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.NewError(500, errors.ErrCodeInternalServer, "NES yanıtı çözümlenemedi")
	}

	return result, nil
}

// --- Yeni Eklenen E-İrsaliye (E-Despatch) Servis Fonksiyonları ---

// GetEDespatchFileExportTitles
func (s *NESService) GetEDespatchFileExportTitles(company *models.Company, documentType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/definitions/fileexporttitles/%s/titlekeys", s.getBaseURL(company), documentType)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchFileExportTitleDefinition
func (s *NESService) GetEDespatchFileExportTitleDefinition(company *models.Company, documentType, extension string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/definitions/fileexporttitles/%s/%s", s.getBaseURL(company), documentType, extension)
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchFileExportTitles
func (s *NESService) UpdateEDespatchFileExportTitles(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/definitions/fileexporttitles", s.getBaseURL(company))
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// UpdateEDespatchDocument
func (s *NESService) UpdateEDespatchDocument(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/uploads/document/%s", s.getBaseURL(company), uuid)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// GetEDespatchTags
func (s *NESService) GetEDespatchTags(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/tags", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateEDespatchTag
func (s *NESService) CreateEDespatchTag(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/tags", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// GetEDespatchTag
func (s *NESService) GetEDespatchTag(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/tags/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchTag
func (s *NESService) UpdateEDespatchTag(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/tags/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// DeleteEDespatchTag
func (s *NESService) DeleteEDespatchTag(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/tags/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// GetEDespatchIncomingDynamicRules
func (s *NESService) GetEDespatchIncomingDynamicRules(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/notifications/incoming/dynamicrules", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateEDespatchIncomingDynamicRule
func (s *NESService) CreateEDespatchIncomingDynamicRule(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/notifications/incoming/dynamicrules", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// GetEDespatchIncomingDynamicRule
func (s *NESService) GetEDespatchIncomingDynamicRule(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/notifications/incoming/dynamicrules/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchIncomingDynamicRule
func (s *NESService) UpdateEDespatchIncomingDynamicRule(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/notifications/incoming/dynamicrules/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// DeleteEDespatchIncomingDynamicRule
func (s *NESService) DeleteEDespatchIncomingDynamicRule(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/notifications/incoming/dynamicrules/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// ExportEDespatchIncomingDespatches
func (s *NESService) ExportEDespatchIncomingDespatches(company *models.Company, fileType string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/export/%s", s.getBaseURL(company), fileType)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// GetEDespatchIncomingReports
func (s *NESService) GetEDespatchIncomingReports(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/reportmodule/reports", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateEDespatchIncomingReport
func (s *NESService) CreateEDespatchIncomingReport(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/reportmodule/reports", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// DownloadEDespatchIncomingReport
func (s *NESService) DownloadEDespatchIncomingReport(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/reportmodule/reports/%s/download", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchIncomingTemplates
func (s *NESService) GetEDespatchIncomingTemplates(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/reportmodule/templates", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateEDespatchIncomingTemplate
func (s *NESService) CreateEDespatchIncomingTemplate(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/reportmodule/templates", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// GetEDespatchIncomingTemplate
func (s *NESService) GetEDespatchIncomingTemplate(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/reportmodule/templates/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchIncomingTemplate
func (s *NESService) UpdateEDespatchIncomingTemplate(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/reportmodule/templates/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// DeleteEDespatchIncomingTemplate
func (s *NESService) DeleteEDespatchIncomingTemplate(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/reportmodule/templates/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// GetEDespatchIncomingColumns
func (s *NESService) GetEDespatchIncomingColumns(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/reportmodule/columns", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchIncomingTags
func (s *NESService) UpdateEDespatchIncomingTags(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/tags", s.getBaseURL(company))
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// SaveCompanyInIncomingDocument
func (s *NESService) SaveCompanyInIncomingDocument(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/%s/savecompanyindocument", s.getBaseURL(company), uuid)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// BulkOperationIncomingDespatches
func (s *NESService) BulkOperationIncomingDespatches(company *models.Company, operation string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/bulk/%s", s.getBaseURL(company), operation)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// AddUserNoteToIncomingDespatch
func (s *NESService) AddUserNoteToIncomingDespatch(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/%s/usernotes", s.getBaseURL(company), uuid)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// UpdateUserNoteInIncomingDespatch
func (s *NESService) UpdateUserNoteInIncomingDespatch(company *models.Company, uuid, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/%s/usernotes/%s", s.getBaseURL(company), uuid, id)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// DeleteUserNoteFromIncomingDespatch
func (s *NESService) DeleteUserNoteFromIncomingDespatch(company *models.Company, uuid, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/%s/usernotes/%s", s.getBaseURL(company), uuid, id)
	return s.doRequest(company, "DELETE", url, nil)
}

// GetEDespatchIncomingReceiptAdvices
func (s *NESService) GetEDespatchIncomingReceiptAdvices(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/receiptadvices", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchIncomingReceiptAdvice
func (s *NESService) GetEDespatchIncomingReceiptAdvice(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/receiptadvices/%s", s.getBaseURL(company), uuid)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchIncomingReceiptAdviceHTML
func (s *NESService) GetEDespatchIncomingReceiptAdviceHTML(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/receiptadvices/%s/html", s.getBaseURL(company), uuid)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchIncomingReceiptAdviceXML
func (s *NESService) GetEDespatchIncomingReceiptAdviceXML(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/receiptadvices/%s/xml", s.getBaseURL(company), uuid)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchIncomingReceiptAdvicePDF
func (s *NESService) GetEDespatchIncomingReceiptAdvicePDF(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/incoming/receiptadvices/%s/pdf", s.getBaseURL(company), uuid)
	return s.doRequest(company, "GET", url, nil)
}

// SendReceiptAdviceForIncomingDespatch
func (s *NESService) SendReceiptAdviceForIncomingDespatch(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/%s/receiptadvice", s.getBaseURL(company), uuid)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// SendEmailForIncomingDespatch
func (s *NESService) SendEmailForIncomingDespatch(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/incoming/despatches/email/send", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// GetEDespatchOutgoingDynamicRules
func (s *NESService) GetEDespatchOutgoingDynamicRules(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/notifications/outgoing/dynamicrules", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateEDespatchOutgoingDynamicRule
func (s *NESService) CreateEDespatchOutgoingDynamicRule(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/notifications/outgoing/dynamicrules", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// GetEDespatchOutgoingDynamicRule
func (s *NESService) GetEDespatchOutgoingDynamicRule(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/notifications/outgoing/dynamicrules/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchOutgoingDynamicRule
func (s *NESService) UpdateEDespatchOutgoingDynamicRule(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/notifications/outgoing/dynamicrules/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// DeleteEDespatchOutgoingDynamicRule
func (s *NESService) DeleteEDespatchOutgoingDynamicRule(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/notifications/outgoing/dynamicrules/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// ExportEDespatchOutgoingDespatches
func (s *NESService) ExportEDespatchOutgoingDespatches(company *models.Company, fileType string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/export/%s", s.getBaseURL(company), fileType)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// GetEDespatchOutgoingReports
func (s *NESService) GetEDespatchOutgoingReports(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/reportmodule/reports", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateEDespatchOutgoingReport
func (s *NESService) CreateEDespatchOutgoingReport(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/reportmodule/reports", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// DownloadEDespatchOutgoingReport
func (s *NESService) DownloadEDespatchOutgoingReport(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/reportmodule/reports/%s/download", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchOutgoingTemplates
func (s *NESService) GetEDespatchOutgoingTemplates(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/reportmodule/templates", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// CreateEDespatchOutgoingTemplate
func (s *NESService) CreateEDespatchOutgoingTemplate(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/reportmodule/templates", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// GetEDespatchOutgoingTemplate
func (s *NESService) GetEDespatchOutgoingTemplate(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/reportmodule/templates/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchOutgoingTemplate
func (s *NESService) UpdateEDespatchOutgoingTemplate(company *models.Company, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/reportmodule/templates/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// DeleteEDespatchOutgoingTemplate
func (s *NESService) DeleteEDespatchOutgoingTemplate(company *models.Company, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/reportmodule/templates/%s", s.getBaseURL(company), id)
	return s.doRequest(company, "DELETE", url, nil)
}

// GetEDespatchOutgoingColumns
func (s *NESService) GetEDespatchOutgoingColumns(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/reportmodule/columns", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchOutgoingTags
func (s *NESService) UpdateEDespatchOutgoingTags(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/tags", s.getBaseURL(company))
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// UpdateEDespatchOutgoingReceiverAlias
func (s *NESService) UpdateEDespatchOutgoingReceiverAlias(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/%s/receiveralias", s.getBaseURL(company), uuid)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// SaveCompanyInOutgoingDocument
func (s *NESService) SaveCompanyInOutgoingDocument(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/%s/savecompanyindocument", s.getBaseURL(company), uuid)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// BulkOperationOutgoingDespatches
func (s *NESService) BulkOperationOutgoingDespatches(company *models.Company, operation string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/bulk/%s", s.getBaseURL(company), operation)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// ResendErrorDocument
func (s *NESService) ResendErrorDocument(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/uploads/resend/%s", s.getBaseURL(company), uuid)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// AddUserNoteToOutgoingDespatch
func (s *NESService) AddUserNoteToOutgoingDespatch(company *models.Company, uuid string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/%s/usernotes", s.getBaseURL(company), uuid)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// UpdateUserNoteInOutgoingDespatch
func (s *NESService) UpdateUserNoteInOutgoingDespatch(company *models.Company, uuid, id string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/%s/usernotes/%s", s.getBaseURL(company), uuid, id)
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// DeleteUserNoteFromOutgoingDespatch
func (s *NESService) DeleteUserNoteFromOutgoingDespatch(company *models.Company, uuid, id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/%s/usernotes/%s", s.getBaseURL(company), uuid, id)
	return s.doRequest(company, "DELETE", url, nil)
}

// GetEDespatchOutgoingReceiptAdvices
func (s *NESService) GetEDespatchOutgoingReceiptAdvices(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/receiptadvices", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchOutgoingReceiptAdvice
func (s *NESService) GetEDespatchOutgoingReceiptAdvice(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/receiptadvices/%s", s.getBaseURL(company), uuid)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchOutgoingReceiptAdviceHTML
func (s *NESService) GetEDespatchOutgoingReceiptAdviceHTML(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/receiptadvices/%s/html", s.getBaseURL(company), uuid)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchOutgoingReceiptAdviceXML
func (s *NESService) GetEDespatchOutgoingReceiptAdviceXML(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/receiptadvices/%s/xml", s.getBaseURL(company), uuid)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchOutgoingReceiptAdvicePDF
func (s *NESService) GetEDespatchOutgoingReceiptAdvicePDF(company *models.Company, uuid string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/receiptadvices/%s/pdf", s.getBaseURL(company), uuid)
	return s.doRequest(company, "GET", url, nil)
}

// DeleteEDespatchOutgoingDrafts
func (s *NESService) DeleteEDespatchOutgoingDrafts(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/drafts", s.getBaseURL(company))
	return s.doRequest(company, "DELETE", url, nil)
}

// SendEmailForOutgoingDespatch
func (s *NESService) SendEmailForOutgoingDespatch(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/outgoing/despatches/email/send", s.getBaseURL(company))
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// GetEDespatchEmailSettings
func (s *NESService) GetEDespatchEmailSettings(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/definitions/mailing/email/settings", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchEmailSettings
func (s *NESService) UpdateEDespatchEmailSettings(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/definitions/mailing/email/settings", s.getBaseURL(company))
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// GetEDespatchSmsSettings
func (s *NESService) GetEDespatchSmsSettings(company *models.Company) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/definitions/mailing/sms/settings", s.getBaseURL(company))
	return s.doRequest(company, "GET", url, nil)
}

// UpdateEDespatchSmsSettings
func (s *NESService) UpdateEDespatchSmsSettings(company *models.Company, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/definitions/mailing/sms/settings", s.getBaseURL(company))
	return s.doRequest(company, "PUT", url, bytes.NewBuffer(data))
}

// GetEDespatchUsersZip
func (s *NESService) GetEDespatchUsersZip(company *models.Company, aliasType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/users/zip/%s", s.getBaseURL(company), aliasType)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchUserByIdentifier
func (s *NESService) GetEDespatchUserByIdentifier(company *models.Company, identifier, aliasType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/users/%s/%s", s.getBaseURL(company), identifier, aliasType)
	return s.doRequest(company, "GET", url, nil)
}

// GetEDespatchUserByIdentifierPost
func (s *NESService) GetEDespatchUserByIdentifierPost(company *models.Company, aliasType string, payload map[string]interface{}) (map[string]interface{}, error) {
	data, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/edespatch/v1/users/%s", s.getBaseURL(company), aliasType)
	return s.doRequest(company, "POST", url, bytes.NewBuffer(data))
}

// SearchEDespatchUserByTitle
func (s *NESService) SearchEDespatchUserByTitle(company *models.Company, query, aliasType string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/edespatch/v1/users/search/%s/%s", s.getBaseURL(company), query, aliasType)
	return s.doRequest(company, "GET", url, nil)
}
