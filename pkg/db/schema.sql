-- PostgreSQL Schema for Aygıt Muhasebe Integration

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Base Function for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- User Types: ADMIN, MALI_MUSAVIR, USER
DO $$ BEGIN
    CREATE TYPE user_role AS ENUM ('ADMIN', 'MALI_MUSAVIR', 'USER');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Environments: TEST, PRODUCTION
DO $$ BEGIN
    CREATE TYPE app_environment AS ENUM ('TEST', 'PRODUCTION');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Addresses Table
CREATE TABLE IF NOT EXISTS addresses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    title TEXT,
    street_name TEXT,
    city_sub_division_name TEXT,
    city_name TEXT,
    country_name TEXT,
    postal_zone TEXT,
    telephone TEXT,
    telefax TEXT,
    electronic_mail TEXT,
    website_uri TEXT,
    party_tax_scheme_name TEXT,
    is_default BOOLEAN DEFAULT FALSE
);

-- Companies Table
CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    party_identification TEXT UNIQUE NOT NULL,
    party_name TEXT,
    first_name TEXT,
    family_name TEXT,
    logo_url TEXT,
    
    -- E-Dönüşüm Durumları
    is_einvoice BOOLEAN DEFAULT FALSE,
    is_edespatch BOOLEAN DEFAULT FALSE,
    is_earchive BOOLEAN DEFAULT FALSE,
    is_esmm_user BOOLEAN DEFAULT FALSE,
    is_emm_user BOOLEAN DEFAULT FALSE,
    is_export_company BOOLEAN DEFAULT FALSE,
    
    -- NES Entegrasyon Bilgileri
    nes_user TEXT,
    nes_password TEXT,
    nes_username TEXT,
    nes_portal_password TEXT,
    nes_api_key TEXT,
    environment app_environment DEFAULT 'TEST',
    app_key TEXT,
    app_secret TEXT,
    nes_status_updated_at TIMESTAMP WITH TIME ZONE,
    
    -- Etiket ve Tasarım Bilgileri
    e_invoice_alias TEXT,
    e_despatch_alias TEXT,
    selected_pk_alias TEXT,
    selected_gb_alias TEXT,
    xslt_template_name TEXT,
    default_series VARCHAR(3),
    
    -- SMS Gateway
    sms_user TEXT,
    sms_password TEXT,
    sms_header TEXT,
    
    -- Email Gateway
    email_host TEXT,
    email_port INTEGER,
    email_user TEXT,
    email_password TEXT,

    -- Base Fields
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    version INTEGER DEFAULT 1,
    soft_delete BOOLEAN DEFAULT FALSE
);

DROP TRIGGER IF EXISTS update_companies_updated_at ON companies;
CREATE TRIGGER update_companies_updated_at BEFORE UPDATE ON companies FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    full_name TEXT NOT NULL,
    role user_role DEFAULT 'USER',
    company_id UUID REFERENCES companies(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    version INTEGER DEFAULT 1,
    soft_delete BOOLEAN DEFAULT FALSE
);

DROP TRIGGER IF EXISTS update_users_updated_at ON users;
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- UserCompanies Table
CREATE TABLE IF NOT EXISTS user_companies (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    company_id UUID REFERENCES companies(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, company_id)
);

-- Products Table (Ürün ve Hizmetler)
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    code VARCHAR(100),                    -- Ürün Kodu / SKU
    name TEXT NOT NULL,                   -- Ürün/Hizmet Adı
    description TEXT,                     -- Açıklama
    
    -- Birim ve Fiyat
    unit_code VARCHAR(10) NOT NULL DEFAULT 'C62', -- UBL Kodları (C62: Adet, KGM: KG vb.)
    unit_price DECIMAL(18, 4) DEFAULT 0,
    currency_code VARCHAR(3) DEFAULT 'TRY',
    
    -- Vergi ve Muafiyet (NES API Zorunlulukları)
    vat_rate DECIMAL(5, 2) DEFAULT 20,            -- KDV Oranı
    tax_exemption_reason_code VARCHAR(10),        -- Muafiyet Kodu (0 KDV durumunda)
    tax_exemption_reason_description TEXT,        -- Muafiyet Açıklaması
    withholding_tax_code VARCHAR(10),             -- Tevkifat Kodu
    gtip_code VARCHAR(12),                        -- İhracat için GTIP Kodu
    
    -- NES Ürün Servis Detayları
    brand_name VARCHAR(100),
    model_name VARCHAR(100),
    
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

DROP TRIGGER IF EXISTS update_products_updated_at ON products;
CREATE TRIGGER update_products_updated_at BEFORE UPDATE ON products FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- Invoices Table (Faturalar)
CREATE TABLE IF NOT EXISTS invoices (
    id BIGSERIAL PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    
    -- E-Fatura Kimlik Bilgileri
    uuid UUID NOT NULL UNIQUE,            -- ETTN (En kritik alan)
    invoice_number VARCHAR(16),           -- Fatura No
    invoice_type VARCHAR(20),             -- SATIŞ, İADE, TEVKİFAT vb.
    profile_id VARCHAR(20),               -- TEMEL, TİCARİ vb.
    direction VARCHAR(10) NOT NULL,       -- INCOMING veya OUTGOING
    
    -- Tarih Bilgileri
    issue_date DATE NOT NULL,
    issue_time TIME,
    
    -- Cari Bilgileri
    sender_vkn_tckn VARCHAR(11) NOT NULL,
    sender_name TEXT NOT NULL,
    receiver_vkn_tckn VARCHAR(11) NOT NULL,
    receiver_name TEXT NOT NULL,
    
    -- Finansal Toplamlar
    currency_code VARCHAR(3) DEFAULT 'TRY',
    exchange_rate DECIMAL(18, 4) DEFAULT 1.0,
    payable_amount DECIMAL(18, 2) NOT NULL,       -- Ödenecek Toplam
    tax_exclusive_amount DECIMAL(18, 2),          -- Toplam Matrah
    tax_inclusive_amount DECIMAL(18, 2),          -- KDV Dahil Toplam
    allowance_total_amount DECIMAL(18, 2),        -- Toplam İskonto
    
    -- NES & GİB Durum Takibi
    status_code INTEGER,                  -- NES Durum Kodu
    status_description TEXT,              -- NES Durum Mesajı
    gib_status_code INTEGER,              -- GİB Durum Kodu
    gib_status_description TEXT,          -- GİB Durum Mesajı
    
    -- Sistem Notları
    is_read BOOLEAN DEFAULT FALSE,
    raw_content JSONB,                    -- NES'ten gelen orijinal JSON yanıtı
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

DROP TRIGGER IF EXISTS update_invoices_updated_at ON invoices;
CREATE TRIGGER update_invoices_updated_at BEFORE UPDATE ON invoices FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- Invoice Items Table (Fatura Satırları)
CREATE TABLE IF NOT EXISTS invoice_items (
    id BIGSERIAL PRIMARY KEY,
    invoice_id BIGINT NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
    product_id BIGINT REFERENCES products(id) ON DELETE SET NULL,
    
    line_number INTEGER,                  -- Satır No (1, 2, 3...)
    name TEXT NOT NULL,                   -- Ürün Adı (O anki fatura verisi)
    description TEXT,
    
    -- Miktar ve Fiyat
    quantity DECIMAL(18, 4) NOT NULL,
    unit_code VARCHAR(10) NOT NULL,
    unit_price DECIMAL(18, 4) NOT NULL,
    
    -- İskonto ve Vergi
    discount_rate DECIMAL(5, 2) DEFAULT 0,
    discount_amount DECIMAL(18, 2) DEFAULT 0,
    vat_rate DECIMAL(5, 2) NOT NULL,
    vat_amount DECIMAL(18, 2) NOT NULL,
    
    -- Satır Bazlı Özel Kodlar
    tax_exemption_reason_code VARCHAR(10),
    gtip_code VARCHAR(12),
    
    -- Satır Toplamları
    line_extension_amount DECIMAL(18, 2) NOT NULL, -- Matrah
    total_amount DECIMAL(18, 2) NOT NULL,           -- KDV Dahil Toplam
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indeksler
CREATE INDEX IF NOT EXISTS idx_invoices_company_date ON invoices(company_id, issue_date);
CREATE INDEX IF NOT EXISTS idx_invoices_uuid ON invoices(uuid);
CREATE INDEX IF NOT EXISTS idx_invoices_sender_vkn ON invoices(sender_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_invoices_receiver_vkn ON invoices(receiver_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_invoice_items_invoice_id ON invoice_items(invoice_id);
CREATE INDEX IF NOT EXISTS idx_products_code ON products(company_id, code);

-- Despatch Advices Table (E-İrsaliyeler)
CREATE TABLE IF NOT EXISTS despatch_advices (
    id BIGSERIAL PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    
    -- E-İrsaliye Kimlik Bilgileri
    uuid UUID NOT NULL UNIQUE,            -- ETTN
    document_number VARCHAR(16),          -- İrsaliye No
    despatch_type_code VARCHAR(20),       -- SEVK, MATBUDAN vb.
    profile_id VARCHAR(20),               -- TEMEL, TİCARİ vb.
    direction VARCHAR(10) NOT NULL,       -- INCOMING veya OUTGOING
    
    -- Tarih ve Notlar
    issue_date DATE NOT NULL,
    issue_time TIME,
    order_number VARCHAR(100),            -- Sipariş Numarası
    user_note TEXT,                       -- Kullanıcı Notu
    document_note TEXT,                   -- Belge Notu
    
    -- Cari Bilgileri
    sender_vkn_tckn VARCHAR(11) NOT NULL,
    sender_name TEXT NOT NULL,
    receiver_vkn_tckn VARCHAR(11) NOT NULL,
    receiver_name TEXT NOT NULL,
    receiver_alias TEXT,                  -- Alıcı Etiketi (PK)
    
    -- Durum Takibi
    despatch_status VARCHAR(20) DEFAULT 'None', -- Waiting, Succeed, Error, Unknown
    mail_status VARCHAR(20),                    -- Succeed, Failed, Readed
    receiver_type VARCHAR(20),                  -- Sanal, Kayitli
    erp_transferred BOOLEAN DEFAULT FALSE,
    is_read BOOLEAN DEFAULT FALSE,
    is_archived BOOLEAN DEFAULT FALSE,
    
    -- Sistem Notları
    tags JSONB,                           -- Atanmış etiketler (Array of UUIDs)
    raw_content JSONB,                    -- Orijinal JSON yanıtı
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

DROP TRIGGER IF EXISTS update_despatch_advices_updated_at ON despatch_advices;
CREATE TRIGGER update_despatch_advices_updated_at BEFORE UPDATE ON despatch_advices FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- Despatch Advice Items Table (İrsaliye Satırları)
CREATE TABLE IF NOT EXISTS despatch_advice_items (
    id BIGSERIAL PRIMARY KEY,
    despatch_advice_id BIGINT NOT NULL REFERENCES despatch_advices(id) ON DELETE CASCADE,
    product_id BIGINT REFERENCES products(id) ON DELETE SET NULL,
    
    line_number INTEGER,
    name TEXT NOT NULL,
    description TEXT,
    
    -- Miktar
    quantity DECIMAL(18, 4) NOT NULL,
    unit_code VARCHAR(10) NOT NULL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- İrsaliye İndeksleri
CREATE INDEX IF NOT EXISTS idx_despatch_company_date ON despatch_advices(company_id, issue_date);
CREATE INDEX IF NOT EXISTS idx_despatch_uuid ON despatch_advices(uuid);
CREATE INDEX IF NOT EXISTS idx_despatch_sender_vkn ON despatch_advices(sender_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_despatch_receiver_vkn ON despatch_advices(receiver_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_despatch_items_id ON despatch_advice_items(despatch_advice_id);

-- E-Archive Invoices Table (E-Arşiv Faturalar)
CREATE TABLE IF NOT EXISTS earchive_invoices (
    id BIGSERIAL PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    
    -- E-Arşiv Kimlik Bilgileri
    uuid UUID NOT NULL UNIQUE,            -- ETTN
    invoice_number VARCHAR(16),           -- Fatura No
    invoice_type_code VARCHAR(20),        -- SATIS, IADE, ISTISNA vb.
    profile_id VARCHAR(20),               -- e-Arşiv için genellikle normal profil
    direction VARCHAR(10) NOT NULL,       -- Genellikle OUTGOING
    
    -- Tarih ve Notlar
    issue_date DATE NOT NULL,
    issue_time TIME,
    order_number VARCHAR(100),            -- Sipariş Numarası
    despatch_number VARCHAR(100),         -- İrsaliye Numarası
    user_note TEXT,                       -- Kullanıcı Notu
    document_note TEXT,                   -- Belge Notu
    
    -- Cari Bilgileri
    sender_vkn_tckn VARCHAR(11) NOT NULL,
    sender_name TEXT NOT NULL,
    receiver_vkn_tckn VARCHAR(11) NOT NULL,
    receiver_name TEXT NOT NULL,
    
    -- Finansal Toplamlar
    payable_amount DECIMAL(18, 2) NOT NULL,
    currency_code VARCHAR(3) DEFAULT 'TRY',
    
    -- Durum Takibi (NES API Spesifik)
    invoice_status VARCHAR(20),           -- WaitingSign, Signed, Error
    report_division_status VARCHAR(20),   -- None, Waiting, Succeed, Error
    send_type VARCHAR(20),                -- Kagit, Elektronik
    sales_platform VARCHAR(20),           -- Internet, Normal
    mail_status VARCHAR(20),              -- Succeed, Failed, Readed
    sms_status VARCHAR(20),               -- Succeed, Failed, Readed
    luca_transfer_status VARCHAR(20),     -- None, Succeded, Error, Unknown
    
    -- Bayraklar
    is_canceled BOOLEAN DEFAULT FALSE,
    is_read BOOLEAN DEFAULT FALSE,
    is_archived BOOLEAN DEFAULT FALSE,
    
    -- Sistem Notları
    tags JSONB,                           -- Atanmış etiketler (Array of UUIDs)
    raw_content JSONB,                    -- Orijinal JSON yanıtı
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

DROP TRIGGER IF EXISTS update_earchive_invoices_updated_at ON earchive_invoices;
CREATE TRIGGER update_earchive_invoices_updated_at BEFORE UPDATE ON earchive_invoices FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- E-Archive Invoice Items Table (E-Arşiv Fatura Satırları)
CREATE TABLE IF NOT EXISTS earchive_invoice_items (
    id BIGSERIAL PRIMARY KEY,
    earchive_invoice_id BIGINT NOT NULL REFERENCES earchive_invoices(id) ON DELETE CASCADE,
    product_id BIGINT REFERENCES products(id) ON DELETE SET NULL,
    
    line_number INTEGER,
    name TEXT NOT NULL,
    description TEXT,
    
    -- Miktar ve Fiyat
    quantity DECIMAL(18, 4) NOT NULL,
    unit_code VARCHAR(10) NOT NULL,
    unit_price DECIMAL(18, 4) NOT NULL,
    total_amount DECIMAL(18, 2) NOT NULL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- E-Arşiv İndeksleri
CREATE INDEX IF NOT EXISTS idx_earchive_company_date ON earchive_invoices(company_id, issue_date);
CREATE INDEX IF NOT EXISTS idx_earchive_uuid ON earchive_invoices(uuid);
CREATE INDEX IF NOT EXISTS idx_earchive_sender_vkn ON earchive_invoices(sender_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_earchive_receiver_vkn ON earchive_invoices(receiver_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_earchive_items_id ON earchive_invoice_items(earchive_invoice_id);

-- System Logs
CREATE TABLE IF NOT EXISTS system_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    action TEXT NOT NULL,
    details JSONB,
    ip_address TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- E-MM Vouchers Table (E-Müstahsil Makbuzları)
CREATE TABLE IF NOT EXISTS e_mm_vouchers (
    id BIGSERIAL PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,

    -- E-MM Kimlik Bilgileri
    uuid UUID NOT NULL UNIQUE,            -- ETTN
    voucher_number VARCHAR(16),           -- Makbuz No
    voucher_type_code VARCHAR(20),        -- E-MM vb.
    direction VARCHAR(10) NOT NULL,       -- Genellikle OUTGOING

    -- Tarih ve Notlar
    issue_date DATE NOT NULL,
    issue_time TIME,
    user_note TEXT,                       -- Kullanıcı Notu
    document_note TEXT,                   -- Belge Notu

    -- Cari Bilgileri
    sender_vkn_tckn VARCHAR(11) NOT NULL,
    receiver_vkn_tckn VARCHAR(11) NOT NULL,
    receiver_name TEXT NOT NULL,
    receiver_address TEXT,
    receiver_city TEXT,
    receiver_country TEXT,

    -- Tutar Bilgileri
    total_amount DECIMAL(18, 2) NOT NULL,
    tax_amount DECIMAL(18, 2) DEFAULT 0,
    payable_amount DECIMAL(18, 2) NOT NULL,
    currency_code VARCHAR(3) DEFAULT 'TRY',

    -- Durum Bilgileri
    status VARCHAR(50) DEFAULT 'DRAFT',   -- DRAFT, PENDING, APPROVED, CANCELED vb.
    status_description TEXT,

    -- E-Posta Gönderim Durumu
    is_mail_sent BOOLEAN DEFAULT FALSE,
    mail_send_date TIMESTAMP WITH TIME ZONE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

DROP TRIGGER IF EXISTS update_e_mm_vouchers_updated_at ON e_mm_vouchers;
CREATE TRIGGER update_e_mm_vouchers_updated_at BEFORE UPDATE ON e_mm_vouchers FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- E-MM Voucher Items Table (E-Müstahsil Makbuzu Satırları)
CREATE TABLE IF NOT EXISTS e_mm_voucher_items (
    id BIGSERIAL PRIMARY KEY,
    e_mm_voucher_id BIGINT NOT NULL REFERENCES e_mm_vouchers(id) ON DELETE CASCADE,
    product_id BIGINT REFERENCES products(id) ON DELETE SET NULL,

    line_number INTEGER,
    name TEXT NOT NULL,
    description TEXT,

    -- Miktar ve Fiyat
    quantity DECIMAL(18, 4) NOT NULL,
    unit_code VARCHAR(10) NOT NULL,
    unit_price DECIMAL(18, 4) NOT NULL,
    total_amount DECIMAL(18, 2) NOT NULL,

    -- Vergiler (Stopaj vb. e-MM özel)
    tax_rate DECIMAL(5, 2) DEFAULT 0,
    tax_amount DECIMAL(18, 2) DEFAULT 0,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- E-SMM Vouchers Table (E-Serbest Meslek Makbuzları)
CREATE TABLE IF NOT EXISTS e_smm_vouchers (
    id BIGSERIAL PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,

    -- E-SMM Kimlik Bilgileri
    uuid UUID NOT NULL UNIQUE,            -- ETTN
    voucher_number VARCHAR(16),           -- Makbuz No
    voucher_type_code VARCHAR(20),        -- E-SMM vb.
    direction VARCHAR(10) NOT NULL,       -- Genellikle OUTGOING

    -- Tarih ve Notlar
    issue_date DATE NOT NULL,
    issue_time TIME,
    user_note TEXT,                       -- Kullanıcı Notu
    document_note TEXT,                   -- Belge Notu

    -- Cari Bilgileri
    sender_vkn_tckn VARCHAR(11) NOT NULL,
    receiver_vkn_tckn VARCHAR(11) NOT NULL,
    receiver_name TEXT NOT NULL,
    receiver_address TEXT,
    receiver_city TEXT,
    receiver_country TEXT,

    -- Tutar Bilgileri
    total_amount DECIMAL(18, 2) NOT NULL,
    tax_amount DECIMAL(18, 2) DEFAULT 0,
    payable_amount DECIMAL(18, 2) NOT NULL,
    currency_code VARCHAR(3) DEFAULT 'TRY',

    -- Durum Bilgileri
    status VARCHAR(50) DEFAULT 'DRAFT',   -- DRAFT, PENDING, APPROVED, CANCELED vb.
    status_description TEXT,

    -- E-Posta Gönderim Durumu
    is_mail_sent BOOLEAN DEFAULT FALSE,
    mail_send_date TIMESTAMP WITH TIME ZONE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

DROP TRIGGER IF EXISTS update_e_smm_vouchers_updated_at ON e_smm_vouchers;
CREATE TRIGGER update_e_smm_vouchers_updated_at BEFORE UPDATE ON e_smm_vouchers FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- E-SMM Voucher Items Table (E-Serbest Meslek Makbuzu Satırları)
CREATE TABLE IF NOT EXISTS e_smm_voucher_items (
    id BIGSERIAL PRIMARY KEY,
    e_smm_voucher_id BIGINT NOT NULL REFERENCES e_smm_vouchers(id) ON DELETE CASCADE,
    product_id BIGINT REFERENCES products(id) ON DELETE SET NULL,

    line_number INTEGER,
    name TEXT NOT NULL,
    description TEXT,

    -- Miktar ve Fiyat
    quantity DECIMAL(18, 4) NOT NULL,
    unit_code VARCHAR(10) NOT NULL,
    unit_price DECIMAL(18, 4) NOT NULL,
    total_amount DECIMAL(18, 2) NOT NULL,

    -- Vergiler (Stopaj vb. e-SMM özel)
    tax_rate DECIMAL(5, 2) DEFAULT 0,
    tax_amount DECIMAL(18, 2) DEFAULT 0,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- E-MM İndeksleri
CREATE INDEX IF NOT EXISTS idx_e_mm_company_date ON e_mm_vouchers(company_id, issue_date);
CREATE INDEX IF NOT EXISTS idx_e_mm_uuid ON e_mm_vouchers(uuid);
CREATE INDEX IF NOT EXISTS idx_e_mm_sender_vkn ON e_mm_vouchers(sender_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_e_mm_receiver_vkn ON e_mm_vouchers(receiver_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_e_mm_items_id ON e_mm_voucher_items(e_mm_voucher_id);

-- E-SMM İndeksleri
CREATE INDEX IF NOT EXISTS idx_e_smm_company_date ON e_smm_vouchers(company_id, issue_date);
CREATE INDEX IF NOT EXISTS idx_e_smm_uuid ON e_smm_vouchers(uuid);
CREATE INDEX IF NOT EXISTS idx_e_smm_sender_vkn ON e_smm_vouchers(sender_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_e_smm_receiver_vkn ON e_smm_vouchers(receiver_vkn_tckn);
CREATE INDEX IF NOT EXISTS idx_e_smm_items_id ON e_smm_voucher_items(e_smm_voucher_id);
