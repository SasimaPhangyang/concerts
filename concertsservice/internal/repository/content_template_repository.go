package repository

import (
	"concerts/internal/models"
	"database/sql"
)

type ContentTemplateRepository interface {
	GetAllTemplates() ([]models.ContentTemplate, error)
}

type contentTemplateRepository struct {
	db *sql.DB
}

func NewContentTemplateRepository(db *sql.DB) ContentTemplateRepository {
	return &contentTemplateRepository{db: db}
}

func (r *contentTemplateRepository) GetAllTemplates() ([]models.ContentTemplate, error) {
	// ปรับคำสั่ง SQL ให้สอดคล้องกับ field ใน struct
	rows, err := r.db.Query("SELECT id, name, description, created_at, updated_at FROM content_templates ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []models.ContentTemplate
	for rows.Next() {
		var t models.ContentTemplate
		// ปรับให้ตรงกับคอลัมน์ในฐานข้อมูล
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		templates = append(templates, t)
	}
	return templates, rows.Err()
}
