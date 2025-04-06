package repository

import (
	"concerts/internal/models"
	"database/sql"
	"errors"
)

type ConcertRepository interface {
	GetAllConcerts() ([]models.Concert, error)
	GetConcertByID(id int) (*models.Concert, error)
	SearchConcerts(query string) ([]models.Concert, error) // เพิ่ม method นี้
}

type concertRepository struct {
	db *sql.DB
}

func NewConcertRepository(db *sql.DB) ConcertRepository {
	return &concertRepository{db: db}
}

func (r *concertRepository) GetAllConcerts() ([]models.Concert, error) {
	rows, err := r.db.Query("SELECT id, name, location, date FROM concerts ORDER BY date")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var concerts []models.Concert
	for rows.Next() {
		var c models.Concert
		if err := rows.Scan(&c.ID, &c.Name, &c.Location, &c.Date); err != nil {
			return nil, err
		}
		concerts = append(concerts, c)
	}
	return concerts, rows.Err()
}

func (r *concertRepository) GetConcertByID(id int) (*models.Concert, error) {
	var c models.Concert
	err := r.db.QueryRow("SELECT id, name, location, date FROM concerts WHERE id=$1", id).
		Scan(&c.ID, &c.Name, &c.Location, &c.Date)

	if err == sql.ErrNoRows {
		return nil, errors.New("concert not found")
	} else if err != nil {
		return nil, err
	}
	return &c, nil
}

// เพิ่ม method สำหรับค้นหาคอนเสิร์ต
func (r *concertRepository) SearchConcerts(query string) ([]models.Concert, error) {
	// ใช้ SQL Query ที่ใช้ LIKE เพื่อค้นหาคอนเสิร์ตตามชื่อหรือสถานที่
	rows, err := r.db.Query("SELECT id, name, location, date FROM concerts WHERE name ILIKE $1 OR location ILIKE $1 ORDER BY date", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var concerts []models.Concert
	for rows.Next() {
		var c models.Concert
		if err := rows.Scan(&c.ID, &c.Name, &c.Location, &c.Date); err != nil {
			return nil, err
		}
		concerts = append(concerts, c)
	}
	return concerts, rows.Err()
}
