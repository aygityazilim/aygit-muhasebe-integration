package models

import (
	"time"

	"github.com/google/uuid"
)

// VoucherStatus represents the status of a voucher
type VoucherStatus string

const (
	VoucherStatusDraft    VoucherStatus = "DRAFT"
	VoucherStatusPending  VoucherStatus = "PENDING"
	VoucherStatusApproved VoucherStatus = "APPROVED"
	VoucherStatusCanceled VoucherStatus = "CANCELED"
)

// EMMVoucher represents an e-Müstahsil Makbuzu
type EMMVoucher struct {
	ID              int64         `db:"id" json:"id"`
	CompanyID       uuid.UUID     `db:"company_id" json:"companyId"`
	UUID            uuid.UUID     `db:"uuid" json:"uuid"`
	VoucherNumber   string        `db:"voucher_number" json:"voucherNumber"`
	VoucherTypeCode string        `db:"voucher_type_code" json:"voucherTypeCode"`
	Direction       string        `db:"direction" json:"direction"`
	IssueDate       string        `db:"issue_date" json:"issueDate"`
	IssueTime       string        `db:"issue_time" json:"issueTime,omitempty"`
	UserNote        string        `db:"user_note" json:"userNote,omitempty"`
	DocumentNote    string        `db:"document_note" json:"documentNote,omitempty"`
	SenderVKNTCKN   string        `db:"sender_vkn_tckn" json:"senderVknTckn"`
	ReceiverVKNTCKN string        `db:"receiver_vkn_tckn" json:"receiverVknTckn"`
	ReceiverName    string        `db:"receiver_name" json:"receiverName"`
	ReceiverAddress string        `db:"receiver_address" json:"receiverAddress,omitempty"`
	ReceiverCity    string        `db:"receiver_city" json:"receiverCity,omitempty"`
	ReceiverCountry string        `db:"receiver_country" json:"receiverCountry,omitempty"`
	TotalAmount     float64       `db:"total_amount" json:"totalAmount"`
	TaxAmount       float64       `db:"tax_amount" json:"taxAmount"`
	PayableAmount   float64       `db:"payable_amount" json:"payableAmount"`
	CurrencyCode    string        `db:"currency_code" json:"currencyCode"`
	Status          VoucherStatus `db:"status" json:"status"`
	StatusDesc      string        `db:"status_description" json:"statusDescription,omitempty"`
	IsMailSent      bool          `db:"is_mail_sent" json:"isMailSent"`
	MailSendDate    *time.Time    `db:"mail_send_date" json:"mailSendDate,omitempty"`
	CreatedAt       time.Time     `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time     `db:"updated_at" json:"updatedAt"`

	Items []EMMVoucherItem `json:"items,omitempty"`
}

// EMMVoucherItem represents a line item in an e-Müstahsil Makbuzu
type EMMVoucherItem struct {
	ID           int64     `db:"id" json:"id"`
	EMMVoucherID int64     `db:"e_mm_voucher_id" json:"eMmVoucherId"`
	ProductID    *int64    `db:"product_id" json:"productId,omitempty"`
	LineNumber   int       `db:"line_number" json:"lineNumber"`
	Name         string    `db:"name" json:"name"`
	Description  string    `db:"description" json:"description,omitempty"`
	Quantity     float64   `db:"quantity" json:"quantity"`
	UnitCode     string    `db:"unit_code" json:"unitCode"`
	UnitPrice    float64   `db:"unit_price" json:"unitPrice"`
	TotalAmount  float64   `db:"total_amount" json:"totalAmount"`
	TaxRate      float64   `db:"tax_rate" json:"taxRate"`
	TaxAmount    float64   `db:"tax_amount" json:"taxAmount"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
}

// ESMMVoucher represents an e-Serbest Meslek Makbuzu
type ESMMVoucher struct {
	ID              int64         `db:"id" json:"id"`
	CompanyID       uuid.UUID     `db:"company_id" json:"companyId"`
	UUID            uuid.UUID     `db:"uuid" json:"uuid"`
	VoucherNumber   string        `db:"voucher_number" json:"voucherNumber"`
	VoucherTypeCode string        `db:"voucher_type_code" json:"voucherTypeCode"`
	Direction       string        `db:"direction" json:"direction"`
	IssueDate       string        `db:"issue_date" json:"issueDate"`
	IssueTime       string        `db:"issue_time" json:"issueTime,omitempty"`
	UserNote        string        `db:"user_note" json:"userNote,omitempty"`
	DocumentNote    string        `db:"document_note" json:"documentNote,omitempty"`
	SenderVKNTCKN   string        `db:"sender_vkn_tckn" json:"senderVknTckn"`
	ReceiverVKNTCKN string        `db:"receiver_vkn_tckn" json:"receiverVknTckn"`
	ReceiverName    string        `db:"receiver_name" json:"receiverName"`
	ReceiverAddress string        `db:"receiver_address" json:"receiverAddress,omitempty"`
	ReceiverCity    string        `db:"receiver_city" json:"receiverCity,omitempty"`
	ReceiverCountry string        `db:"receiver_country" json:"receiverCountry,omitempty"`
	TotalAmount     float64       `db:"total_amount" json:"totalAmount"`
	TaxAmount       float64       `db:"tax_amount" json:"taxAmount"`
	PayableAmount   float64       `db:"payable_amount" json:"payableAmount"`
	CurrencyCode    string        `db:"currency_code" json:"currencyCode"`
	Status          VoucherStatus `db:"status" json:"status"`
	StatusDesc      string        `db:"status_description" json:"statusDescription,omitempty"`
	IsMailSent      bool          `db:"is_mail_sent" json:"isMailSent"`
	MailSendDate    *time.Time    `db:"mail_send_date" json:"mailSendDate,omitempty"`
	CreatedAt       time.Time     `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time     `db:"updated_at" json:"updatedAt"`

	Items []ESMMVoucherItem `json:"items,omitempty"`
}

// ESMMVoucherItem represents a line item in an e-Serbest Meslek Makbuzu
type ESMMVoucherItem struct {
	ID            int64     `db:"id" json:"id"`
	ESMMVoucherID int64     `db:"e_smm_voucher_id" json:"eSmmVoucherId"`
	ProductID     *int64    `db:"product_id" json:"productId,omitempty"`
	LineNumber    int       `db:"line_number" json:"lineNumber"`
	Name          string    `db:"name" json:"name"`
	Description   string    `db:"description" json:"description,omitempty"`
	Quantity      float64   `db:"quantity" json:"quantity"`
	UnitCode      string    `db:"unit_code" json:"unitCode"`
	UnitPrice     float64   `db:"unit_price" json:"unitPrice"`
	TotalAmount   float64   `db:"total_amount" json:"totalAmount"`
	TaxRate       float64   `db:"tax_rate" json:"taxRate"`
	TaxAmount     float64   `db:"tax_amount" json:"taxAmount"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
}
