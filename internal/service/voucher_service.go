package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"

	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/pkg/errors"
)

// NESVoucherService handles communication with the NES API for Vouchers (EMM and ESMM).
type NESVoucherService struct {
	BaseURL string // e.g., https://apitest.nes.com.tr/emm
	Client  *http.Client
}

func NewNESVoucherService(baseURL string) *NESVoucherService {
	return &NESVoucherService{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// executeRequest is a helper for sending HTTP requests to NES.
func (s *NESVoucherService) executeRequest(method, endpoint string, company *models.Company, payload []byte) ([]byte, error) {
	req, err := http.NewRequest(method, s.BaseURL+endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Assuming authorization requires API Key or similar from company object
	// You might need to adjust the exact auth headers based on existing patterns
	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", "application/json")
	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.FormatNESError(resp.StatusCode, body)
	}

	return body, nil
}

// Forward provides a generic way to proxy a request directly to the NES API.
// It handles the construction of the request, passing through the Content-Type,
// and returns the raw response body. It can be used for both JSON and multipart data.
func (s *NESVoucherService) Forward(method, endpoint string, company *models.Company, contentType string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest(method, s.BaseURL+endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create forward request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	} else {
		req.Header.Set("Content-Type", "application/json")
	}

	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("forward request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read forward response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.FormatNESError(resp.StatusCode, body)
	}

	return body, nil
}

// ---------------------------------------------------------
// COMMON VOUCHER OPERATIONS (Works for both EMM and ESMM based on BaseURL)
// ---------------------------------------------------------

func (s *NESVoucherService) GetVouchers(company *models.Company) (string, error) {
	resp, err := s.executeRequest("GET", "/v1/vouchers", company, nil)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *NESVoucherService) GetVoucher(company *models.Company, uuid string) (string, error) {
	resp, err := s.executeRequest("GET", fmt.Sprintf("/v1/vouchers/%s", uuid), company, nil)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *NESVoucherService) GetDraftVouchers(company *models.Company) (string, error) {
	resp, err := s.executeRequest("GET", "/v1/vouchers/drafts", company, nil)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *NESVoucherService) UploadDocument(company *models.Company, data []byte) (string, error) {
	// NES expects multipart/form-data with a "file" field containing the JSON payload
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Write the JSON data into a file part
	part, err := writer.CreateFormFile("File", "document.json")
	if err != nil {
		return "", err
	}
	part.Write(data)

	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", s.BaseURL+"/v1/uploads/document", body)
	if err != nil {
		return "", fmt.Errorf("failed to create upload request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+company.GetNesAPIKey())
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if company.Environment == "TEST" {
		req.Header.Set("Environment", "TEST")
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", errors.FormatNESError(resp.StatusCode, respBody)
	}

	return string(respBody), nil
}

func (s *NESVoucherService) SendDraftVoucher(company *models.Company, uuids []string) (string, error) {
	payload, _ := json.Marshal(uuids)
	resp, err := s.executeRequest("POST", "/v1/uploads/draft/send", company, payload)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *NESVoucherService) CancelVoucher(company *models.Company, uuid string, reason string) (string, error) {
	payload := map[string]string{
		"uuid":   uuid,
		"reason": reason,
	}
	data, _ := json.Marshal(payload)
	resp, err := s.executeRequest("POST", "/v1/vouchers/cancel", company, data)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *NESVoucherService) GetCanceledVouchers(company *models.Company) (string, error) {
	resp, err := s.executeRequest("GET", "/v1/vouchers/canceled", company, nil)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *NESVoucherService) WithdrawCancelVoucher(company *models.Company, uuid string) (string, error) {
	payload := map[string]string{"uuid": uuid}
	data, _ := json.Marshal(payload)
	resp, err := s.executeRequest("POST", "/v1/vouchers/canceled/withdraw", company, data)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *NESVoucherService) SendVoucherEmail(company *models.Company, uuid string, emails []string) (string, error) {
	payload := map[string]interface{}{
		"uuid":   uuid,
		"emails": emails,
	}
	data, _ := json.Marshal(payload)
	resp, err := s.executeRequest("POST", "/v1/vouchers/email/send", company, data)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (s *NESVoucherService) AddUserNote(company *models.Company, uuid string, note string) (string, error) {
	payload := map[string]string{"note": note}
	data, _ := json.Marshal(payload)
	resp, err := s.executeRequest("POST", fmt.Sprintf("/v1/vouchers/%s/usernotes", uuid), company, data)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// ---------------------------------------------------------
// PASSTHROUGH OPERATION
// ---------------------------------------------------------

// PassthroughRequest enables a generic handler to pass HTTP method, raw path, and body directly to NES API.
func (s *NESVoucherService) PassthroughRequest(method, originalURL string, company *models.Company, payload []byte) ([]byte, error) {
	// originalURL could be like "/api/v1/emm/notifications/dynamicrules"
	// We want to extract the part after /emm to append to the base URL
	// For example: s.BaseURL is "https://apitest.nes.com.tr/emm"

	// Fast way to find the endpoint suffix:
	// Find "/emm/" or "/esmm/" in originalURL and take everything after it.
	importStr := ""
	if len(originalURL) > 0 {

		// Find where "/emm/" or "/esmm/" starts
		importPathPosEmm := -1
		importPathPosEsmm := -1

		for i := 0; i <= len(originalURL)-5; i++ {
			if originalURL[i:i+5] == "/emm/" {
				importPathPosEmm = i
				break
			}
		}

		for i := 0; i <= len(originalURL)-6; i++ {
			if originalURL[i:i+6] == "/esmm/" {
				importPathPosEsmm = i
				break
			}
		}

		if importPathPosEmm != -1 {
			importStr = originalURL[importPathPosEmm+4:] // keep the leading slash, e.g. /v1/...
		} else if importPathPosEsmm != -1 {
			importStr = originalURL[importPathPosEsmm+5:]
		}
	}

	if importStr == "" {
		return nil, fmt.Errorf("could not parse endpoint from originalURL: %s", originalURL)
	}

	return s.executeRequest(method, importStr, company, payload)
}
