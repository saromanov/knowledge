package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	models "github.com/saromanov/knowledge/internal/models/storage"
	"github.com/saromanov/knowledge/internal/storage"
)

var (
	errAuthorNotDefined = errors.New("author is not defined")
	errIDNotDefined = errors.New("id is not defined")
)

type postgres struct {
	cfg Config
	db  *sql.DB
}

// New provides initialization of the module
func New(cfg Config) storage.Storage {
	return &postgres{
		cfg: cfg,
	}
}

// Init provides initialization to db
func (p *postgres) Init(ctx context.Context) error {
	if err := p.connect(); err != nil {
		return err
	}
	return nil
}

// CreatePage provides creating of the page
func (p *postgres) CreatePage(ctx context.Context, m *models.Page) (int64, error) {
	sqlStatement := `
INSERT INTO page (title, body, created_at, updated_at, author_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`
	var id int64
	err := p.db.QueryRow(sqlStatement, m.Title, m.Body, m.CreatedAt, m.UpdatedAt, m.AuthorID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// CreateAuthor provides creating of the page
func (p *postgres) CreateAuthor(ctx context.Context, m *models.Author) (int64, error) {
	if m == nil {
		return 0, fmt.Errorf("author request is not defined")
	}
	if p.db == nil {
		return 0, fmt.Errorf("db init is not defined")
	}
	var id int64
	sqlStatement := `
INSERT INTO author (name, created_at)
VALUES ($1, $2)
RETURNING id`
	err := p.db.QueryRow(sqlStatement, m.Name, m.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetPage provides getting of the page by id
func (p *postgres) GetPage(ctx context.Context, id int64) (*models.Page, error) {
	if id == 0 {
		return &models.Page{}, errIDNotDefined
	}
	if p.db == nil {
		return &models.Page{}, fmt.Errorf("db init is not defined")
	}
	rows, err := p.db.Query(`SELECT * FROM "page" WHERE id = $1`,id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data, err := p.find(rows)
	if err != nil {
		return nil, err
	}
	if len(data) > 0 {
		return data[0], nil
	}
	return &models.Page{}, nil
}

// DeletePage provides deleting of the page by id
func (p *postgres) DeletePage(ctx context.Context, id int64) error {
	if id == 0 {
		return errIDNotDefined
	}
	if _, err := p.db.Exec(`DELETE FROM "page" WHERE id = $1`,id); err != nil {
		return err
	}
	return nil
}


// GetPages provides getting of the page by author id
func (p *postgres) GetPages(ctx context.Context, author string) ([]*models.Page, error) {
	if author == "" {
		return nil, errAuthorNotDefined
	}
	rows, err := p.db.Query(`SELECT * FROM "page" WHERE author_id = $1`,author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return p.find(rows)
}

// Close provides closing of connectin to db
func (p *postgres) Close(ctx context.Context) error {
	if err := p.db.Close(); err != nil {
		return err
	}
	return nil
}

func (p *postgres) find(rows *sql.Rows) ([]*models.Page, error) {
	result := []*models.Page{}
	for rows.Next() {
		var pages models.Page
		if err := rows.Scan(&pages.ID, &pages.CreatedAt, &pages.UpdatedAt, &pages.Title, &pages.Body, &pages.AuthorID); err != nil {
			return nil, err
		}
		result = append(result, &pages)
	}
	return result, nil
}

// connect provides connection to postgres
func (p *postgres) connect() error {
	var lastErr error
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.cfg.Host, p.cfg.Port, p.cfg.Username, p.cfg.Password, p.cfg.DB)
	for i := 0; i < p.cfg.ConnectRetries; i++ {
		db, err := sql.Open("postgres", conn)
		if err != nil {
			lastErr = err
			continue
		}
		if err := db.Ping(); err != nil {
			lastErr = err
			continue
		}
		p.db = db
	}
	return lastErr
}
