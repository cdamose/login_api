package adapters

import (
	"context"
	"database/sql"
	"fmt"

	"login_api/internal/auth/model/dao"
	"login_api/internal/common/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sirupsen/logrus"
)

type MySQLURLShortnerRepository struct {
	db     *sqlx.DB
	logger logrus.Entry
	config config.Config
}

func NewMySQLUrlShortnerRepository(db *sqlx.DB, logger logrus.Entry, config config.Config) *MySQLURLShortnerRepository {
	if db == nil {
		panic("missing db")
	}
	return &MySQLURLShortnerRepository{db: db, logger: logger, config: config}
}
func (m MySQLURLShortnerRepository) InsertWebsiteIfNotExist(ctx context.Context, website string) (string, error) {
	var id string
	err := m.db.GetContext(ctx, &id, "SELECT id FROM websites WHERE domain = ?", website)

	if err == nil {
		return id, nil
	} else if err != sql.ErrNoRows {
		return "", err
	}

	id = uuid.New().String()
	param := map[string]interface{}{
		"id":     id,
		"name":   website,
		"domain": website,
	}
	_, err = m.db.NamedExec("INSERT INTO websites (id, name, domain) VALUES (:id, :name, :domain)", param)
	if err != nil {
		return "", err
	}

	return id, nil

}

func (m MySQLURLShortnerRepository) UpdateVistCount(ctx context.Context, website_id string) error {

	var count int
	err := m.db.GetContext(ctx, &count, "SELECT count FROM websites_acccess_count WHERE website_id  = ?", website_id)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to retrieve website access count: %w", err)
	}

	if err == sql.ErrNoRows {
		id := uuid.New().String()
		_, err := m.db.NamedExec("INSERT INTO websites_acccess_count (id, website_id, count) VALUES (:id, :website_id, 1)", map[string]interface{}{"id": id, "website_id": website_id})
		if err != nil {
			return fmt.Errorf("failed to insert website access count: %w", err)
		}
	} else {
		_, err := m.db.NamedExec("UPDATE websites_acccess_count SET count = count + 1 WHERE website_id = :website_id", map[string]interface{}{"website_id": website_id})
		if err != nil {
			return fmt.Errorf("failed to update website access count: %w", err)
		}
	}

	return nil

}
func (m MySQLURLShortnerRepository) FetchTop3Websites(ctx context.Context) ([]dao.Website, error) {
	query := `
	SELECT w.id, w.name, w.domain
	FROM websites AS w
	JOIN (
		SELECT website_id, COUNT(*) AS total_count
		FROM websites_acccess_count
		GROUP BY website_id
		ORDER BY total_count DESC
		LIMIT 3
	) AS wc ON w.id = wc.website_id
	`

	var websites []dao.Website
	err := m.db.SelectContext(ctx, &websites, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top 3 websites: %w", err)
	}

	return websites, nil
}
func (m MySQLURLShortnerRepository) GetActualURL(ctx context.Context, shorten_url string) (string, error) {
	var originalURL string
	query := "SELECT original_url FROM urls WHERE short_url = ?"
	err := m.db.GetContext(ctx, &originalURL, query, shorten_url)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}
	return originalURL, nil
}
func (m MySQLURLShortnerRepository) InsertShortenUrl(ctx context.Context, url string, shorten_url string) (*dao.ShortenUrl, error) {
	var data_obj dao.ShortenUrl

	params := map[string]interface{}{
		"id":           uuid.New().String(),
		"original_url": url,
		"short_url":    shorten_url,
	}
	query := "INSERT INTO urls (id, original_url, short_url) VALUES (:id, :original_url, :short_url)"
	_, err := m.db.NamedExec(query, params)
	if err != nil {
		fmt.Println("Error inserting data into the table:", err)
		return nil, err
	}

	if err == nil {
		data_obj = dao.ShortenUrl{
			ShortenURL: shorten_url,
		}
	}

	return &data_obj, err
}

func (m MySQLURLShortnerRepository) IsShortenURLExists(ctx context.Context, shorten_url string) bool {
	var id string
	query := "SELECT id FROM urls WHERE short_url = ?"
	err := m.db.GetContext(ctx, &id, query, shorten_url)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			return true
		}

	}
	return true
}
