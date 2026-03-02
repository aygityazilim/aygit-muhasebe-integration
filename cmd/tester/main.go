package main

import (
	"aygit-muhasebe-integration/config"
	"aygit-muhasebe-integration/internal/models"
	"aygit-muhasebe-integration/internal/service"
	"aygit-muhasebe-integration/pkg/db"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func main() {
	// 1. Ayarları ve Veritabanını Başlat
	config.InitConfig()
	db.InitDB()
	defer db.CloseDB()

	nesService := service.NewNESService()

	// 2. Test Firmalarını Getir (Test Kurum 02 -> Test Kurum 01)
	var sender, receiver models.Company
	err := db.DB.Get(&sender, "SELECT * FROM companies WHERE party_identification = $1", "1234567802")
	if err != nil {
		log.Fatalf("Gönderici firma bulunamadı: %v", err)
	}
	err = db.DB.Get(&receiver, "SELECT * FROM companies WHERE party_identification = $1", "1234567801")
	if err != nil {
		log.Fatalf("Alıcı firma bulunamadı: %v", err)
	}

	fmt.Printf("\n--- TEST BAŞLATILIYOR ---\n")
	fmt.Printf("Gönderici: %s (%s)\n", sender.PartyName, sender.PartyIdentification)
	fmt.Printf("Alıcı: %s (%s)\n", receiver.PartyName, receiver.PartyIdentification)

	// 3. E-Fatura Testi
	testEInvoice(nesService, &sender, &receiver)

	// 4. E-Arşiv Testi
	testEArchive(nesService, &sender)

	fmt.Printf("\n--- TEST TAMAMLANDI ---\n")
}

func testEInvoice(s *service.NESService, sender, receiver *models.Company) {
	fmt.Printf("\n[E-Fatura Testi]\n")

	// Müşteri adını AYGIT yapıyoruz, adresi izmir olacak (şablondan)
	// E-Fatura testi için alıcının gerçek bir e-fatura mükellefi (1234567801) olması gerekir
	receiver.PartyName = "AYGIT"
	receiver.PartyIdentification = "1234567801"

	ttn := uuid.New().String()
	xmlData := generateMinimalUBL(ttn, sender, receiver, "INVOICE")

	fmt.Println("1. Fatura yükleniyor (Upload)...")
	uploadRes, err := s.UploadInvoice(sender, []byte(xmlData), map[string]string{
		"IsDirectSend": "false",
	})
	if err != nil {
		log.Printf("HATA (Upload): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", uploadRes)

	fmt.Println("2. Taslak onaylanıyor (SendDraft)...")
	sendRes, err := s.SendDraftInvoices(sender, []string{ttn})
	if err != nil {
		log.Printf("HATA (SendDraft): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", sendRes)
}

func testEArchive(s *service.NESService, sender *models.Company) {
	fmt.Printf("\n[E-Arşiv Testi]\n")

	ttn := uuid.New().String()
	// E-Arşiv için alıcı e-fatura mükellefi olmayan biri olmalı
	dummyReceiver := &models.Company{
		PartyIdentification: "11111111111",
		PartyName:           "AYGIT",
	}
	xmlData := generateMinimalUBL(ttn, sender, dummyReceiver, "EARCHIVE")

	fmt.Println("1. E-Arşiv faturası yükleniyor (Upload)...")
	uploadRes, err := s.UploadEArchiveInvoice(sender, []byte(xmlData), map[string]string{
		"IsDirectSend": "false",
	})
	if err != nil {
		log.Printf("HATA (E-Arşiv Upload): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", uploadRes)

	fmt.Println("2. E-Arşiv taslağı gönderiliyor...")
	sendRes, err := s.SendDraftEArchiveInvoices(sender, []string{ttn})
	if err != nil {
		log.Printf("HATA (E-Arşiv SendDraft): %v", err)
		return
	}
	fmt.Printf("   Başarılı. Yanıt: %v\n", sendRes)
}

func generateMinimalUBL(uuidStr string, sender, receiver *models.Company, docType string) string {
	nowDate := time.Now().Format("2006-01-02")
	nowTime := time.Now().Format("15:04:05")
	profileID := "TEMELFATURA"
	series := "EFT"
	extraHeader := ""

	if docType == "EARCHIVE" {
		profileID = "EARSIVFATURA"
		series = "EAR"
		// E-Arşiv için SendType zorunlu AdditionalDocumentReference
		extraHeader = fmt.Sprintf(`
    <cac:AdditionalDocumentReference>
        <cbc:ID>ELEKTRONIK</cbc:ID>
        <cbc:IssueDate>%s</cbc:IssueDate>
        <cbc:DocumentTypeCode>SEND_TYPE</cbc:DocumentTypeCode>
    </cac:AdditionalDocumentReference>`, nowDate)
	}

	senderScheme := "VKN"
	if len(sender.PartyIdentification) == 11 {
		senderScheme = "TCKN"
	}

	// Signature Bloğu (Zorunlu 1)
	signatureBlock := fmt.Sprintf(`
    <cac:Signature>
        <cbc:ID schemeID="VKN_TCKN">%s</cbc:ID>
        <cac:SignatoryParty>
            <cac:PartyIdentification>
                <cbc:ID schemeID="%s">%s</cbc:ID>
            </cac:PartyIdentification>
            <cac:PostalAddress>
                <cbc:StreetName>Test Mah. No:1</cbc:StreetName>
                <cbc:CitySubdivisionName>Kadikoy</cbc:CitySubdivisionName>
                <cbc:CityName>Istanbul</cbc:CityName>
                <cac:Country><cbc:Name>TÜRKİYE</cbc:Name></cac:Country>
            </cac:PostalAddress>
        </cac:SignatoryParty>
        <cac:DigitalSignatureAttachment>
            <cac:ExternalReference>
                <cbc:URI>#Signature</cbc:URI>
            </cac:ExternalReference>
        </cac:DigitalSignatureAttachment>
    </cac:Signature>`, sender.PartyIdentification, senderScheme, sender.PartyIdentification)

	formatParty := func(c *models.Company) string {
		scheme := "VKN"
		if len(c.PartyIdentification) == 11 {
			scheme = "TCKN"
		}

		idBlock := fmt.Sprintf(`<cac:PartyIdentification><cbc:ID schemeID="%s">%s</cbc:ID></cac:PartyIdentification>`, scheme, c.PartyIdentification)
		nameBlock := fmt.Sprintf(`<cac:PartyName><cbc:Name>%s</cbc:Name></cac:PartyName>`, c.PartyName)
		personBlock := ""
		if scheme == "TCKN" {
			nameBlock = "" // TCKN için genelde PartyName yerine Person kullanılır
			personBlock = fmt.Sprintf(`<cac:Person><cbc:FirstName>%s</cbc:FirstName><cbc:FamilyName>.</cbc:FamilyName></cac:Person>`, c.PartyName)
		}

		return fmt.Sprintf(`
		<cac:Party>
		%s
		%s
		<cac:PostalAddress>
		<cbc:StreetName>Test Mah. No:1</cbc:StreetName>
		<cbc:CitySubdivisionName>Konak</cbc:CitySubdivisionName>
		<cbc:CityName>izmir</cbc:CityName>
		<cac:Country><cbc:Name>TÜRKİYE</cbc:Name></cac:Country>
		</cac:PostalAddress>
		<cac:PartyTaxScheme><cac:TaxScheme><cbc:Name>VergiDairesi</cbc:Name></cac:TaxScheme></cac:PartyTaxScheme>
		%s
		</cac:Party>`, idBlock, nameBlock, personBlock)
		}


	senderBlock := formatParty(sender)
	receiverBlock := formatParty(receiver)

	// UBL-TR 2.1 Şablonu (Zorunlu alanlar sırasıyla)
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Invoice xmlns="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"
         xmlns:cac="urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2"
         xmlns:cbc="urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2 ../xsdrt/maindoc/UBL-Invoice-2.1.xsd">
    <cbc:UBLVersionID>2.1</cbc:UBLVersionID>
    <cbc:CustomizationID>TR1.2</cbc:CustomizationID>
    <cbc:ProfileID>%s</cbc:ProfileID>
    <cbc:ID>%s%s%09d</cbc:ID>
    <cbc:CopyIndicator>false</cbc:CopyIndicator>
    <cbc:UUID>%s</cbc:UUID>
    <cbc:IssueDate>%s</cbc:IssueDate>
    <cbc:IssueTime>%s</cbc:IssueTime>
    <cbc:InvoiceTypeCode>SATIS</cbc:InvoiceTypeCode>
    <cbc:Note>YALNIZ : YÜZONSEKİZ TL SIFIR Kr.</cbc:Note>
    <cbc:DocumentCurrencyCode>TRY</cbc:DocumentCurrencyCode>
    <cbc:LineCountNumeric>1</cbc:LineCountNumeric>%s
    %s
    <cac:AccountingSupplierParty>%s</cac:AccountingSupplierParty>
    <cac:AccountingCustomerParty>%s</cac:AccountingCustomerParty>
    <cac:TaxTotal>
        <cbc:TaxAmount currencyID="TRY">18.00</cbc:TaxAmount>
        <cac:TaxSubtotal>
            <cbc:TaxableAmount currencyID="TRY">100.00</cbc:TaxableAmount>
            <cbc:TaxAmount currencyID="TRY">18.00</cbc:TaxAmount>
            <cbc:Percent>18.00</cbc:Percent>
            <cac:TaxCategory>
                <cac:TaxScheme>
                    <cbc:Name>KDV</cbc:Name>
                    <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                </cac:TaxScheme>
            </cac:TaxCategory>
        </cac:TaxSubtotal>
    </cac:TaxTotal>
    <cac:LegalMonetaryTotal>
        <cbc:LineExtensionAmount currencyID="TRY">100.00</cbc:LineExtensionAmount>
        <cbc:TaxExclusiveAmount currencyID="TRY">100.00</cbc:TaxExclusiveAmount>
        <cbc:TaxInclusiveAmount currencyID="TRY">118.00</cbc:TaxInclusiveAmount>
        <cbc:AllowanceTotalAmount currencyID="TRY">0.00</cbc:AllowanceTotalAmount>
        <cbc:PayableAmount currencyID="TRY">118.00</cbc:PayableAmount>
    </cac:LegalMonetaryTotal>
    <cac:InvoiceLine>
        <cbc:ID>1</cbc:ID>
        <cbc:InvoicedQuantity unitCode="C62">1</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="TRY">100.00</cbc:LineExtensionAmount>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="TRY">18.00</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="TRY">100.00</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="TRY">18.00</cbc:TaxAmount>
                <cbc:Percent>18.00</cbc:Percent>
                <cac:TaxCategory>
                    <cac:TaxScheme>
                        <cbc:Name>KDV</cbc:Name>
                        <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Name>Test Hizmeti</cbc:Name>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="TRY">100.00</cbc:PriceAmount>
        </cac:Price>
    </cac:InvoiceLine>
</Invoice>`, profileID, series, time.Now().Format("2006"), time.Now().UnixNano()%1000000000, uuidStr, nowDate, nowTime, extraHeader, signatureBlock, senderBlock, receiverBlock)
}
