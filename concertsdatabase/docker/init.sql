-- สร้างตาราง partners
CREATE TABLE partners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    api_key VARCHAR(255) UNIQUE NOT NULL,
    token_key VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ตารางเก็บ balance ของ partner
CREATE TABLE partner_balances (
    partner_id INT PRIMARY KEY REFERENCES partners(id),
    balance FLOAT NOT NULL DEFAULT 0
);

-- สร้างตาราง categories
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- สร้างตาราง concerts
CREATE TABLE concerts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    partner_id INT REFERENCES partners(id),
    category_id INT REFERENCES categories(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger อัพเดต updated_at ของ concerts
CREATE OR REPLACE FUNCTION update_concerts_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_concerts_modtime
BEFORE UPDATE ON concerts
FOR EACH ROW
EXECUTE FUNCTION update_concerts_modtime();

-- banners
CREATE TABLE banners (
    id SERIAL PRIMARY KEY,
    image_url VARCHAR(255) NOT NULL,
    link VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_banners_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_banners_modtime
BEFORE UPDATE ON banners
FOR EACH ROW
EXECUTE FUNCTION update_banners_modtime();

-- content_templates
CREATE TABLE content_templates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_content_templates_modtime()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_content_templates_modtime
BEFORE UPDATE ON content_templates
FOR EACH ROW
EXECUTE FUNCTION update_content_templates_modtime();

-- commissions
CREATE TABLE commissions (
    id SERIAL PRIMARY KEY,
    partner_id INT REFERENCES partners(id),
    amount FLOAT NOT NULL,
    date_from DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- sales_reports
CREATE TABLE sales_reports (
    product VARCHAR(255) NOT NULL,
    amount FLOAT NOT NULL,
    sales_date TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE sales_by_source (
    source VARCHAR(255) NOT NULL,
    total_sales FLOAT NOT NULL
);

-- partner_rewards
CREATE TABLE partner_rewards (
    reward_id SERIAL PRIMARY KEY,
    amount FLOAT NOT NULL,
    partner_id INT REFERENCES partners(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- bookings
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    concert_id INT REFERENCES concerts(id),
    partner_id INT REFERENCES partners(id),
    tickets INT NOT NULL,
    amount FLOAT NOT NULL,
    booking_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    date DATE
);

-- คำนวณค่า date จาก booking_at
UPDATE bookings SET date = booking_at::date;

-- auto_withdraw
CREATE TABLE auto_withdraw (
    enabled BOOLEAN NOT NULL DEFAULT FALSE
);

-- withdraw_requests
CREATE TABLE withdraw_requests (
    id SERIAL PRIMARY KEY,
    partner_id INT REFERENCES partners(id),
    amount FLOAT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- function update updated_at
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- trigger update updated_at
CREATE TRIGGER update_users_modtime
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- index
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_name ON users(name);

-- ข้อมูลตัวอย่าง categories
INSERT INTO categories (name) VALUES ('Rock'), ('Pop');

-- ข้อมูลตัวอย่าง partners
INSERT INTO partners (name, api_key, token_key) VALUES
('Partner A', 'apikey_a', 'token_a'),
('Partner B', 'apikey_b', 'token_b');

-- ข้อมูล balance เริ่มต้น
INSERT INTO partner_balances (partner_id, balance) VALUES
(1, 0),
(2, 0);

-- ข้อมูลตัวอย่าง concerts
INSERT INTO concerts (name, location, date, partner_id, category_id) VALUES 
('Concert 1', 'Location 1', '2025-05-10 19:00:00', 1, 1),
('Concert 2', 'Location 2', '2025-06-15 18:30:00', 2, 2);

-- ข้อมูลตัวอย่าง users
INSERT INTO users (name, email) VALUES 
('ศศิมา พังยาง', 'sasima@example.com'),

