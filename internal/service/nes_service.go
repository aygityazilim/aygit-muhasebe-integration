package service

import (
	"aygit-muhasebe-integration/internal/models"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
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

	// Eğer baseURL verilmemişse veya hatalı eski developertest URL'si ise apitest'e yönlendir
	if baseURL == "" || strings.Contains(baseURL, "developertest.nes.com.tr") {
		baseURL = "https://apitest.nes.com.tr"
	}

	// URL'nin sonundaki /api'yi temizle (eğer varsa)
	// Çünkü metodlarda /general, /einvoice vb. manuel ekleniyor
	if len(baseURL) > 4 && baseURL[len(baseURL)-4:] == "/api" {
		baseURL = baseURL[:len(baseURL)-4]
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
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d", resp.StatusCode)
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
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d", resp.StatusCode)
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
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d", resp.StatusCode)
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
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d", resp.StatusCode)
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

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("servis çağrısı hatası: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("dosya indirme hatası: durum kodu %d", resp.StatusCode)
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

	// NES API için gerekli olan ek alanlar
	if params["SourceApp"] == "" {
		writer.WriteField("SourceApp", "AYGIT_ENT")
	}
	if params["SenderAlias"] == "" && company.SelectedGbAlias != nil {
		writer.WriteField("SenderAlias", *company.SelectedGbAlias)
	} else if params["SenderAlias"] != "" {
		writer.WriteField("SenderAlias", params["SenderAlias"])
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d, yanıt: %s", resp.StatusCode, string(respBody))
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d", resp.StatusCode)
	}

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d, yanıt: %s", resp.StatusCode, string(respBody))
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d", resp.StatusCode)
	}

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d, yanıt: %s", resp.StatusCode, string(respBody))
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
		return nil, "", fmt.Errorf("dosya indirme hatası: durum kodu %d", resp.StatusCode)
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

	// NES API için gerekli olan ek alanlar
	if params["SourceApp"] == "" {
		writer.WriteField("SourceApp", "AYGIT_ENT")
	}
	if params["SenderAlias"] == "" && company.SelectedGbAlias != nil {
		writer.WriteField("SenderAlias", *company.SelectedGbAlias)
	} else if params["SenderAlias"] != "" {
		writer.WriteField("SenderAlias", params["SenderAlias"])
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d, yanıt: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d, yanıt: %s", resp.StatusCode, string(respBody))
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d", resp.StatusCode)
	}

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d, yanıt: %s", resp.StatusCode, string(respBody))
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
		return nil, "", fmt.Errorf("dosya indirme hatası: durum kodu %d", resp.StatusCode)
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d, yanıt: %s", resp.StatusCode, string(respBody))
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
		writer.WriteField("SourceApp", "AYGIT_ENT")
	}
	if params["SenderAlias"] == "" && company.SelectedGbAlias != nil {
		writer.WriteField("SenderAlias", *company.SelectedGbAlias)
	} else if params["SenderAlias"] != "" {
		writer.WriteField("SenderAlias", params["SenderAlias"])
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d, yanıt: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"raw_response": string(respBody)}, nil
	}
	return result, nil
}

func (s *NESService) SendDraftEArchiveInvoices(company *models.Company, uuids []string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/earchive/v1/uploads/draft/send", s.getBaseURL(company))
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d, yanıt: %s", resp.StatusCode, string(respBody))
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
		return nil, fmt.Errorf("NES servis hatası: durum kodu %d", resp.StatusCode)
	}

	var result []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("yanıt işleme (decode) hatası: %w", err)
	}
	return result, nil
}
