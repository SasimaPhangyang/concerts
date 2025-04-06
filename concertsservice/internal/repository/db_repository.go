package repository

import (
	"database/sql"
	"fmt"
	"time"

	"concerts/internal/config"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {
	// (DSN) Data Source Name คือ String ที่ใช้ระบุข้อมูลการเชื่อมต่อกับฐานข้อมูล
	// การใช้งาน DSN จะขึ้นอยู่กับ Library หรือ Framework ที่ใช้งาน
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// ตั้งค่า Connection Pool
	db.SetMaxOpenConns(25)                 // จำนวน Connection สูงสุดที่สามารถเปิดได้
	db.SetMaxIdleConns(10)                 // จำนวน Connection สูงสุดที่สามารถอยู่ใน Idle State
	db.SetConnMaxLifetime(5 * time.Minute) // อายุการใช้งานสูงสุดของ Connection

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func CheckDBConnection(db *sql.DB) error {
	return db.Ping()
}
