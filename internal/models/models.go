package models

import (
	"time"

	"github.com/google/uuid"
)

// BaseModel includes common fields for all database models
type BaseModel struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
	Version    int        `db:"version" json:"version"`
	SoftDelete bool       `db:"soft_delete" json:"soft_delete"`
}

// UserRole defines the type of user
type UserRole string

const (
	RoleAdmin       UserRole = "ADMIN"
	RoleMaliMusavir UserRole = "MALI_MUSAVIR"
	RoleUser        UserRole = "USER"
)

// Address represents a physical address
type Address struct {
	ID                  uuid.UUID `db:"id" json:"id"`
	CreatedAt           time.Time `db:"created_at" json:"createdAt"`
	Title               string    `db:"title" json:"title"`
	StreetName          string    `db:"street_name" json:"streetName"`
	CitySubDivisionName string    `db:"city_sub_division_name" json:"citySubDivisionName"`
	CityName            string    `db:"city_name" json:"cityName"`
	CountryName         string    `db:"country_name" json:"countryName"`
	PostalZone          string    `db:"postal_zone" json:"postalZone"`
	Telephone           string    `db:"telephone" json:"telephone"`
	Telefax             string    `db:"telefax" json:"telefax"`
	ElectronicMail      string    `db:"electronic_mail" json:"electronicMail"`
	WebsiteUri          string    `db:"website_uri" json:"websiteUri"`
	PartyTaxSchemeName  string    `db:"party_tax_scheme_name" json:"partyTaxSchemeName"`
	IsDefault           bool      `db:"is_default" json:"isDefault"`
}

// Company represents a business entity (Firma bilgilerini temsil eder)
type Company struct {
	BaseModel
	PartyIdentification string  `db:"party_identification" json:"partyIdentification"` // VKN/TCKN
	PartyName           string  `db:"party_name" json:"partyName"`                     // Unvan
	FirstName           *string `db:"first_name" json:"firstName"`                     // Şahıs Firması Adı
	FamilyName          *string `db:"family_name" json:"familyName"`                   // Şahıs Firması Soyadı
	LogoURL             *string `db:"logo_url" json:"logoUrl"`                         // Firma Logo Yolu

	// E-Dönüşüm Durumları
	IsEInvoice      bool `db:"is_einvoice" json:"isEInvoice"`            // E-Fatura kullanıcısı mı?
	IsEDespatch     bool `db:"is_edespatch" json:"isEDespatch"`          // E-İrsaliye kullanıcısı mı?
	IsEArchive      bool `db:"is_earchive" json:"isEArchive"`            // E-Arşiv kullanıcısı mı?
	IsESMMUser      bool `db:"is_esmm_user" json:"isESMMUser"`           // E-SMM kullanıcısı mı?
	IsEMMUser       bool `db:"is_emm_user" json:"isEMMUser"`             // E-MM kullanıcısı mı?
	IsExportCompany bool `db:"is_export_company" json:"isExportCompany"` // İhracat firması mı?

	// NES Portal & API Entegrasyon Bilgileri
	NesUser            *string    `db:"nes_user" json:"nes_user,omitempty"`                       // NES API Kullanıcı Adı
	NesPassword        *string    `db:"nes_password" json:"nes_password,omitempty"`               // NES API Şifresi
	NesUsername        *string    `db:"nes_username" json:"nes_username,omitempty"`               // NES Portal Kullanıcı Adı
	NesPortalPassword  *string    `db:"nes_portal_password" json:"nes_portal_password,omitempty"` // NES Portal Şifresi
	NesAPIKey          *string    `db:"nes_api_key" json:"nes_api_key,omitempty"`                 // API Key
	Environment        string     `db:"environment" json:"environment"`                           // TEST veya PRODUCTION
	AppKey             *string    `db:"app_key" json:"appKey,omitempty"`
	AppSecret          *string    `db:"app_secret" json:"appSecret,omitempty"`
	NesStatusUpdatedAt *time.Time `db:"nes_status_updated_at" json:"nesStatusUpdatedAt,omitempty"`

	// Etiket ve Tasarım Bilgileri
	EInvoiceAlias    *string `db:"e_invoice_alias" json:"eInvoiceAlias"`       // E-Fatura Takma Adı
	EDespatchAlias   *string `db:"e_despatch_alias" json:"eDespatchAlias"`     // E-İrsaliye Takma Adı
	SelectedPkAlias  *string `db:"selected_pk_alias" json:"selectedPkAlias"`   // Varsayılan Posta Kutusu
	SelectedGbAlias  *string `db:"selected_gb_alias" json:"selectedGbAlias"`   // Varsayılan Gönderici Birimi
	XsltTemplateName *string `db:"xslt_template_name" json:"xsltTemplateName"` // Görsel Tasarım Şablonu
	DefaultSeries    *string `db:"default_series" json:"defaultSeries"`        // Varsayılan Seri (örn: FTR)

	// SMS Gateway Ayarları
	SMSUser     *string `db:"sms_user" json:"sms_user,omitempty"`
	SMSPassword *string `db:"sms_password" json:"sms_password,omitempty"`
	SMSHeader   *string `db:"sms_header" json:"sms_header,omitempty"`

	// Email Gateway Ayarları
	EmailHost     *string `db:"email_host" json:"email_host,omitempty"`
	EmailPort     *int    `db:"email_port" json:"email_port,omitempty"`
	EmailUser     *string `db:"email_user" json:"email_user,omitempty"`
	EmailPassword *string `db:"email_password" json:"email_password,omitempty"`

	// İlişkili Veriler
	DefaultAddress *Address `json:"defaultAddress,omitempty"`
	AddressCount   int      `json:"addressCount"`
}

// GetNesAPIKey returns NES API Key safely
func (c *Company) GetNesAPIKey() string {
	if c.NesAPIKey == nil {
		return ""
	}
	return *c.NesAPIKey
}

// GetNesUser returns NES User safely
func (c *Company) GetNesUser() string {
	if c.NesUser == nil {
		return ""
	}
	return *c.NesUser
}

// GetNesPassword returns NES Password safely
func (c *Company) GetNesPassword() string {
	if c.NesPassword == nil {
		return ""
	}
	return *c.NesPassword
}

// User represents a system user
type User struct {
	BaseModel
	Email        string   `db:"email" json:"email"`
	PasswordHash string   `db:"password_hash" json:"-"`
	FullName     string   `db:"full_name" json:"full_name"`
	Role         UserRole `db:"role" json:"role"`

	// Primary Company for regular users
	CompanyID *uuid.UUID `db:"company_id" json:"company_id,omitempty"`
}

// UserCompany manages many-to-many relationship for Mali Müşavir and Admin
type UserCompany struct {
	UserID    uuid.UUID `db:"user_id"`
	CompanyID uuid.UUID `db:"company_id"`
	CreatedAt time.Time `db:"created_at"`
}

// Product represents a goods or service (Ürün veya hizmeti temsil eder)
type Product struct {
	ID                            int64     `db:"id" json:"id"`
	CompanyID                     uuid.UUID `db:"company_id" json:"companyId"`
	Code                          string    `db:"code" json:"code"`
	Name                          string    `db:"name" json:"name"`
	Description                   string    `db:"description" json:"description"`
	UnitCode                      string    `db:"unit_code" json:"unitCode"`
	UnitPrice                     float64   `db:"unit_price" json:"unitPrice"`
	CurrencyCode                  string    `db:"currency_code" json:"currencyCode"`
	VATRate                       float64   `db:"vat_rate" json:"vatRate"`
	TaxExemptionReasonCode        string    `db:"tax_exemption_reason_code" json:"taxExemptionReasonCode"`
	TaxExemptionReasonDescription string    `db:"tax_exemption_reason_description" json:"taxExemptionReasonDescription"`
	WithholdingTaxCode            string    `db:"withholding_tax_code" json:"withholdingTaxCode"`
	GTIPCode                      string    `db:"gtip_code" json:"gtipCode"`
	BrandName                     string    `db:"brand_name" json:"brandName"`
	ModelName                     string    `db:"model_name" json:"modelName"`
	IsActive                      bool      `db:"is_active" json:"isActive"`
	CreatedAt                     time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt                     time.Time `db:"updated_at" json:"updatedAt"`
}

// Invoice represents an e-Invoice or e-Archive document (Fatura belgesini temsil eder)
type Invoice struct {
	ID                   int64         `db:"id" json:"id"`
	CompanyID            uuid.UUID     `db:"company_id" json:"companyId"`
	UUID                 uuid.UUID     `db:"uuid" json:"uuid"`
	InvoiceNumber        string        `db:"invoice_number" json:"invoiceNumber"`
	InvoiceType          string        `db:"invoice_type" json:"invoiceType"`
	ProfileID            string        `db:"profile_id" json:"profileId"`
	Direction            string        `db:"direction" json:"direction"`
	IssueDate            time.Time     `db:"issue_date" json:"issueDate"`
	IssueTime            string        `db:"issue_time" json:"issueTime"`
	SenderVKNTCKN        string        `db:"sender_vkn_tckn" json:"senderVknTckn"`
	SenderName           string        `db:"sender_name" json:"senderName"`
	ReceiverVKNTCKN      string        `db:"receiver_vkn_tckn" json:"receiverVknTckn"`
	ReceiverName         string        `db:"receiver_name" json:"receiverName"`
	CurrencyCode         string        `db:"currency_code" json:"currencyCode"`
	ExchangeRate         float64       `db:"exchange_rate" json:"exchangeRate"`
	PayableAmount        float64       `db:"payable_amount" json:"payableAmount"`
	TaxExclusiveAmount   float64       `db:"tax_exclusive_amount" json:"taxExclusiveAmount"`
	TaxInclusiveAmount   float64       `db:"tax_inclusive_amount" json:"taxInclusiveAmount"`
	AllowanceTotalAmount float64       `db:"allowance_total_amount" json:"allowanceTotalAmount"`
	StatusCode           int           `db:"status_code" json:"statusCode"`
	StatusDescription    string        `db:"status_description" json:"statusDescription"`
	GIBStatusCode        int           `db:"gib_status_code" json:"gibStatusCode"`
	GIBStatusDescription string        `db:"gib_status_description" json:"gibStatusDescription"`
	IsRead               bool          `db:"is_read" json:"isRead"`
	RawContent           []byte        `db:"raw_content" json:"rawContent"`
	CreatedAt            time.Time     `db:"created_at" json:"createdAt"`
	UpdatedAt            time.Time     `db:"updated_at" json:"updatedAt"`
	Items                []InvoiceItem `json:"items,omitempty"`
}

// InvoiceItem represents a single line in an invoice (Fatura satırını temsil eder)
type InvoiceItem struct {
	ID                     int64     `db:"id" json:"id"`
	InvoiceID              int64     `db:"invoice_id" json:"invoiceId"`
	ProductID              *int64    `db:"product_id" json:"productId,omitempty"`
	LineNumber             int       `db:"line_number" json:"lineNumber"`
	Name                   string    `db:"name" json:"name"`
	Description            string    `db:"description" json:"description"`
	Quantity               float64   `db:"quantity" json:"quantity"`
	UnitCode               string    `db:"unit_code" json:"unitCode"`
	UnitPrice              float64   `db:"unit_price" json:"unitPrice"`
	DiscountRate           float64   `db:"discount_rate" json:"discountRate"`
	DiscountAmount         float64   `db:"discount_amount" json:"discountAmount"`
	VATRate                float64   `db:"vat_rate" json:"vatRate"`
	VATAmount              float64   `db:"vat_amount" json:"vatAmount"`
	TaxExemptionReasonCode string    `db:"tax_exemption_reason_code" json:"taxExemptionReasonCode"`
	GTIPCode               string    `db:"gtip_code" json:"gtipCode"`
	LineExtensionAmount    float64   `db:"line_extension_amount" json:"lineExtensionAmount"`
	TotalAmount            float64   `db:"total_amount" json:"totalAmount"`
	CreatedAt              time.Time `db:"created_at" json:"createdAt"`
}

// DespatchAdvice represents an e-Despatch document (İrsaliye belgesini temsil eder)
type DespatchAdvice struct {
	ID                     int64                `db:"id" json:"id"`
	CompanyID              uuid.UUID            `db:"company_id" json:"companyId"`
	UUID                   uuid.UUID            `db:"uuid" json:"uuid"`
	DocumentNumber         string               `db:"document_number" json:"documentNumber"`
	DespatchAdviceTypeCode string               `db:"despatch_type_code" json:"despatchAdviceTypeCode"`
	ProfileID              string               `db:"profile_id" json:"profileId"`
	Direction              string               `db:"direction" json:"direction"`
	IssueDate              time.Time            `db:"issue_date" json:"issueDate"`
	IssueTime              string               `db:"issue_time" json:"issueTime"`
	OrderNumber            string               `db:"order_number" json:"orderNumber"`
	UserNote               string               `db:"user_note" json:"userNote"`
	DocumentNote           string               `db:"document_note" json:"documentNote"`
	SenderVKNTCKN          string               `db:"sender_vkn_tckn" json:"senderVknTckn"`
	SenderName             string               `db:"sender_name" json:"senderName"`
	ReceiverVKNTCKN        string               `db:"receiver_vkn_tckn" json:"receiverVknTckn"`
	ReceiverName           string               `db:"receiver_name" json:"receiverName"`
	ReceiverAlias          string               `db:"receiver_alias" json:"receiverAlias"`
	DespatchStatus         string               `db:"despatch_status" json:"despatchStatus"`
	MailStatus             string               `db:"mail_status" json:"mailStatus"`
	ReceiverType           string               `db:"receiver_type" json:"receiverType"`
	ERPTransferred         bool                 `db:"erp_transferred" json:"erpTransferred"`
	IsRead                 bool                 `db:"is_read" json:"isRead"`
	IsArchived             bool                 `db:"is_archived" json:"archived"`
	Tags                   []string             `db:"tags" json:"tags"`
	RawContent             []byte               `db:"raw_content" json:"rawContent"`
	CreatedAt              time.Time            `db:"created_at" json:"createdAt"`
	UpdatedAt              time.Time            `db:"updated_at" json:"updatedAt"`
	Items                  []DespatchAdviceItem `json:"items,omitempty"`
}

// DespatchAdviceItem represents a single line in a despatch advice (İrsaliye satırını temsil eder)
type DespatchAdviceItem struct {
	ID               int64     `db:"id" json:"id"`
	DespatchAdviceID int64     `db:"despatch_advice_id" json:"despatchAdviceId"`
	ProductID        *int64    `db:"product_id" json:"productId,omitempty"`
	LineNumber       int       `db:"line_number" json:"lineNumber"`
	Name             string    `db:"name" json:"name"`
	Description      string    `db:"description" json:"description"`
	Quantity         float64   `db:"quantity" json:"quantity"`
	UnitCode         string    `db:"unit_code" json:"unitCode"`
	CreatedAt        time.Time `db:"created_at" json:"createdAt"`
}

// EArchiveInvoice represents an e-Archive invoice document (E-Arşiv faturasını temsil eder)
type EArchiveInvoice struct {
	ID                   int64                 `db:"id" json:"id"`
	CompanyID            uuid.UUID             `db:"company_id" json:"companyId"`
	UUID                 uuid.UUID             `db:"uuid" json:"uuid"`
	InvoiceNumber        string                `db:"invoice_number" json:"invoiceNumber"`
	InvoiceTypeCode      string                `db:"invoice_type_code" json:"invoiceTypeCode"`
	ProfileID            string                `db:"profile_id" json:"profileId"`
	Direction            string                `db:"direction" json:"direction"`
	IssueDate            time.Time             `db:"issue_date" json:"issueDate"`
	IssueTime            string                `db:"issue_time" json:"issueTime"`
	OrderNumber          string                `db:"order_number" json:"orderNumber"`
	DespatchNumber       string                `db:"despatch_number" json:"despatchNumber"`
	UserNote             string                `db:"user_note" json:"userNote"`
	DocumentNote         string                `db:"document_note" json:"documentNote"`
	SenderVKNTCKN        string                `db:"sender_vkn_tckn" json:"senderVknTckn"`
	SenderName           string                `db:"sender_name" json:"senderName"`
	ReceiverVKNTCKN      string                `db:"receiver_vkn_tckn" json:"receiverVknTckn"`
	ReceiverName         string                `db:"receiver_name" json:"receiverName"`
	PayableAmount        float64               `db:"payable_amount" json:"payableAmount"`
	CurrencyCode         string                `db:"currency_code" json:"currencyCode"`
	InvoiceStatus        string                `db:"invoice_status" json:"invoiceStatus"`
	ReportDivisionStatus string                `db:"report_division_status" json:"reportDivisionStatus"`
	SendType             string                `db:"send_type" json:"sendType"`
	SalesPlatform        string                `db:"sales_platform" json:"salesPlatform"`
	MailStatus           string                `db:"mail_status" json:"mailStatus"`
	SmsStatus            string                `db:"sms_status" json:"smsStatus"`
	LucaTransferStatus   string                `db:"luca_transfer_status" json:"lucaTransferStatus"`
	IsCanceled           bool                  `db:"is_canceled" json:"isCanceled"`
	IsRead               bool                  `db:"is_read" json:"isRead"`
	IsArchived           bool                  `db:"is_archived" json:"archived"`
	Tags                 []string              `db:"tags" json:"tags"`
	RawContent           []byte                `db:"raw_content" json:"rawContent"`
	CreatedAt            time.Time             `db:"created_at" json:"createdAt"`
	UpdatedAt            time.Time             `db:"updated_at" json:"updatedAt"`
	Items                []EArchiveInvoiceItem `json:"items,omitempty"`
}

// EArchiveInvoiceItem represents a single line in an e-archive invoice
type EArchiveInvoiceItem struct {
	ID                int64     `db:"id" json:"id"`
	EArchiveInvoiceID int64     `db:"earchive_invoice_id" json:"earchiveInvoiceId"`
	ProductID         *int64    `db:"product_id" json:"productId,omitempty"`
	LineNumber        int       `db:"line_number" json:"lineNumber"`
	Name              string    `db:"name" json:"name"`
	Description       string    `db:"description" json:"description"`
	Quantity          float64   `db:"quantity" json:"quantity"`
	UnitCode          string    `db:"unit_code" json:"unitCode"`
	UnitPrice         float64   `db:"unit_price" json:"unitPrice"`
	TotalAmount       float64   `db:"total_amount" json:"totalAmount"`
	CreatedAt         time.Time `db:"created_at" json:"createdAt"`
}

// SystemLog represents an entry in the system logs
type SystemLog struct {
	ID        uuid.UUID   `db:"id" json:"id"`
	UserID    *uuid.UUID  `db:"user_id" json:"user_id,omitempty"`
	Action    string      `db:"action" json:"action"`
	Details   interface{} `db:"details" json:"details"`
	IPAddress *string     `db:"ip_address" json:"ip_address,omitempty"`
	CreatedAt time.Time   `db:"created_at" json:"created_at"`
}
