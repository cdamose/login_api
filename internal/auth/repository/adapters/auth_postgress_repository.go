package adapters

import (
	"context"
	"database/sql"
	"fmt"

	"login_api/internal/auth/model/dao"
	"login_api/internal/common/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type PostgresAuthRepository struct {
	db     *sqlx.DB
	logger logrus.Entry
	config config.Config
}

func NewPostgresAuthRepository(db *sqlx.DB, logger logrus.Entry, config config.Config) *PostgresAuthRepository {
	if db == nil {
		panic("missing db")
	}
	return &PostgresAuthRepository{db: db, logger: logger, config: config}
}
func (m PostgresAuthRepository) CheckMobileNumberAlredayExists(ctx context.Context, mobile_number string) (bool, error) {
	var id string
	err := m.db.GetContext(ctx, &id, "SELECT user_id FROM useraccount WHERE phone_number = $1", mobile_number)
	fmt.Println(id)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m PostgresAuthRepository) GetUserProfile(ctx context.Context, mobile_number string) (*dao.UserProfile, error) {
	query := `select * from UserAccount where phone_number = :phone_number`
	param := map[string]interface{}{
		"phone_number": mobile_number,
	}
	var user_profile dao.UserProfile
	result, err := m.db.NamedQueryContext(ctx, query, param)

	if err != nil {
		return nil, err
	}
	defer result.Close()
	if !result.Next() {
		return nil, errors.New("no rows returned after insert")
	}
	err = result.StructScan(&user_profile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan result into struct")
	}
	return &user_profile, nil
}

func (m PostgresAuthRepository) CreateUserProfile(ctx context.Context, mobile_number string) (*dao.UserProfile, error) {
	query := `INSERT INTO UserAccount (user_id, phone_number) VALUES (uuid_generate_v4(), :phone_number) RETURNING *`
	param := map[string]interface{}{
		"phone_number": mobile_number,
	}
	var user_profile dao.UserProfile
	result, err := m.db.NamedQueryContext(ctx, query, param)

	if err != nil {
		return nil, err
	}
	defer result.Close()
	if !result.Next() {
		return nil, errors.New("no rows returned after insert")
	}
	err = result.StructScan(&user_profile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan result into struct")
	}
	return &user_profile, nil
}

func (m PostgresAuthRepository) UpdateOTPUsedStatus(ctx context.Context, user_id string, otp_code string, is_used bool) (bool, error) {
	query := `UPDATE OTP SET is_used = :is_used WHERE user_id = :user_id AND otp_code = :otp_code;`
	param := map[string]interface{}{
		"user_id":  user_id,
		"otp_code": otp_code,
		"is_used":  is_used,
	}
	_, err := m.db.NamedQueryContext(ctx, query, param)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m PostgresAuthRepository) UpdateUserVerfiedStatus(ctx context.Context, user_id string, verified_status bool) (bool, error) {
	query := `UPDATE UserAccount SET is_verified = :verified_status WHERE user_id = :user_id;`
	param := map[string]interface{}{
		"user_id":         user_id,
		"verified_status": verified_status,
	}
	_, err := m.db.NamedQueryContext(ctx, query, param)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (m PostgresAuthRepository) GetValidOTPDetails(ctx context.Context, user_id string, otp_code string) (*dao.OTPDetails, error) {
	query := `SELECT * FROM OTP WHERE user_id = :user_id AND otp_code = :otp_code AND is_used = FALSE;`
	param := map[string]interface{}{
		"user_id":  user_id,
		"otp_code": otp_code,
	}
	var otp_details dao.OTPDetails
	result, err := m.db.NamedQueryContext(ctx, query, param)

	if err != nil {
		return nil, err
	}
	defer result.Close()
	if !result.Next() {
		return nil, errors.New("no rows returned after select")
	}
	err = result.StructScan(&otp_details)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan result into struct")
	}
	return &otp_details, nil
}

func (m PostgresAuthRepository) Login(ctx context.Context, phone_number string, otp_code string) (*string, error) {
	query := "INSERT INTO OTP (user_id, otp_code) SELECT user_id, :otp_code FROM UserAccount WHERE phone_number = :phone_number RETURNING user_id;"
	param := map[string]interface{}{
		"phone_number": phone_number,
		"otp_code":     otp_code,
	}
	var user_id string
	result, err := m.db.NamedQueryContext(ctx, query, param)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if !result.Next() {
		return nil, errors.New("no rows returned after insert")
	}
	err = result.Scan(&user_id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan result into struct")
	}
	return &user_id, nil
}

func (m PostgresAuthRepository) GenerateOTP(ctx context.Context, phone_number string, otp_code string) (bool, error) {
	query := "INSERT INTO OTP (user_id, otp_code) SELECT user_id, $2 FROM UserAccount WHERE phone_number = $1;"

	// param := map[string]interface{}{
	// 	"phone_number": phone_number,
	// 	"otp_code":     otp_code,
	// }
	result, err := m.db.ExecContext(ctx, query, phone_number, otp_code)
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rows_affected != 0 {
		return true, nil
	} else {
		return false, nil
	}

}

// func (m PostgresAuthRepository) FetchTop3Websites(ctx context.Context) ([]dao.Website, error) {
// 	query := `
// 	SELECT w.id, w.name, w.domain
// 	FROM websites AS w
// 	JOIN (
// 		SELECT website_id, COUNT(*) AS total_count
// 		FROM websites_acccess_count
// 		GROUP BY website_id
// 		ORDER BY total_count DESC
// 		LIMIT 3
// 	) AS wc ON w.id = wc.website_id
// 	`

// 	var websites []dao.Website
// 	err := m.db.SelectContext(ctx, &websites, query)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch top 3 websites: %w", err)
// 	}

// 	return websites, nil
// }
// func (m PostgresAuthRepository) GetActualURL(ctx context.Context, shorten_url string) (string, error) {
// 	var originalURL string
// 	query := "SELECT original_url FROM urls WHERE short_url = ?"
// 	err := m.db.GetContext(ctx, &originalURL, query, shorten_url)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return "", err
// 		}
// 		return "", err
// 	}
// 	return originalURL, nil
// }
// func (m PostgresAuthRepository) InsertShortenUrl(ctx context.Context, url string, shorten_url string) (*dao.ShortenUrl, error) {
// 	var data_obj dao.ShortenUrl

// 	params := map[string]interface{}{
// 		"id":           uuid.New().String(),
// 		"original_url": url,
// 		"short_url":    shorten_url,
// 	}
// 	query := "INSERT INTO urls (id, original_url, short_url) VALUES (:id, :original_url, :short_url)"
// 	_, err := m.db.NamedExec(query, params)
// 	if err != nil {
// 		fmt.Println("Error inserting data into the table:", err)
// 		return nil, err
// 	}

// 	if err == nil {
// 		data_obj = dao.ShortenUrl{
// 			ShortenURL: shorten_url,
// 		}
// 	}

// 	return &data_obj, err
// }

// func (m PostgresAuthRepository) IsShortenURLExists(ctx context.Context, shorten_url string) bool {
// 	var id string
// 	query := "SELECT id FROM urls WHERE short_url = ?"
// 	err := m.db.GetContext(ctx, &id, query, shorten_url)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return false
// 		} else {
// 			return true
// 		}

// 	}
// 	return true
// }
