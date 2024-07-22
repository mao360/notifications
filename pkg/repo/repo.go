package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mao360/notifications/models"
	"time"
)

//go:generate go run github.com/vektra/mockery/v2@v2.30.1 --name=RepoI
//go:generate mockery --name=RepoI
type RepoI interface {
	GetUser(username, passwordHash string) (*models.User, error)
	NewUser(user *models.User) error
	Subscribe(followerUsername, authorUsername string) error
	Unsubscribe(followerUsername, authorUsername string) error
	GetNotification(followerUsername string) ([]string, error)
}

type Repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{pool: pool}
}

func (r *Repo) GetUser(username, passwordHash string) (*models.User, error) {
	user := &models.User{}
	rows, err := r.pool.Query(context.Background(),
		`SELECT user_id, password, username, date_of_birth FROM users
		WHERE username = $1;`, username)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("query err: %v", err)
	}
	if !rows.Next() {
		return nil, errors.New("no user")
	}

	var tmpDate time.Time
	err = rows.Scan(&user.ID, &user.Password, &user.UserName, &tmpDate)
	if err != nil {
		return nil, fmt.Errorf("scan err: %v", err)
	}
	user.DateOfBirth = tmpDate.Format(time.DateOnly)
	if user.Password != passwordHash {
		return nil, errors.New("bad password")
	}
	return user, nil
}

func (r *Repo) NewUser(user *models.User) error {
	rows, err := r.pool.Query(context.Background(),
		`SELECT username FROM users
		WHERE username = $1;`, user.UserName)
	defer rows.Close()
	if err != nil {
		return err
	}
	if rows.Next() {
		return errors.New("user already exists")
	}
	_, err = r.pool.Exec(context.Background(),
		`INSERT INTO users (password, username, date_of_birth)
		VALUES ($1, $2, $3);`, user.Password, user.UserName, user.DateOfBirth)
	if err != nil {
		return fmt.Errorf("exec err: %v", err)
	}
	return nil
}

func (r *Repo) Subscribe(followerUsername, authorUsername string) error {
	var followerID, authorID int
	rows, err := r.pool.Query(context.Background(),
		`SELECT * FROM users
		JOIN users_to_subscribers uts ON users.user_id = uts.user_id
		JOIN users u on u.user_id = uts.friend_id
		WHERE users.username = $1 AND u.username = $2;`, followerUsername, authorUsername)
	defer rows.Close()
	if err != nil {
		return fmt.Errorf("1st query err: %v", err)
	}
	if rows.Next() {
		return errors.New("already subscribed")
	}

	rows, err = r.pool.Query(context.Background(),
		`SELECT user_id FROM users
		WHERE username = $1;`, followerUsername)
	if err != nil {
		return fmt.Errorf("2nd query err: %v", err)
	}
	if rows.Next() {
		err = rows.Scan(&followerID)
		if err != nil {
			return fmt.Errorf("1st scan err: %v", err)
		}
	} else {
		return fmt.Errorf("no rows 1")
	}

	rows, err = r.pool.Query(context.Background(),
		`SELECT user_id FROM users
		WHERE username = $1;`, authorUsername)
	if err != nil {
		return fmt.Errorf("3rd query err: %v", err)
	}
	if !rows.Next() {
		return fmt.Errorf("no author found")
	}
	err = rows.Scan(&authorID)
	if err != nil {
		return fmt.Errorf("2nd scan err: %v", err)
	}

	_, err = r.pool.Exec(context.Background(),
		`INSERT INTO users_to_subscribers (user_id, friend_id)
			VALUES ($1, $2);`, followerID, authorID)
	if err != nil {
		return fmt.Errorf("exec err: %v", err)
	}
	return nil
}

func (r *Repo) Unsubscribe(followerUsername, authorUsername string) error {
	rows, err := r.pool.Query(context.Background(),
		`SELECT (*) FROM users
		JOIN users_to_subscribers uts ON users.user_id = uts.user_id
		JOIN users u on u.user_id = uts.friend_id
		WHERE users.username = $1 AND u.username = $2;`, followerUsername, authorUsername)
	defer rows.Close()
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	if !rows.Next() {
		return errors.New("no subscription")
	}
	_, err = r.pool.Exec(context.Background(),
		`DELETE FROM users_to_subscribers
	WHERE user_id = (
        SELECT user_id FROM users
        WHERE username = $1
    ) AND friend_id = (
        SELECT user_id FROM users
        WHERE username = $2
    );`, followerUsername, authorUsername)
	if err != nil {
		return fmt.Errorf("exec err: %v", err)
	}
	return nil
}

func (r *Repo) GetNotification(followerUsername string) ([]string, error) {
	userNames := make([]string, 0)
	rows, err := r.pool.Query(context.Background(),
		`SELECT u.username FROM users
		JOIN users_to_subscribers uts ON users.user_id = uts.user_id
		JOIN users u on u.user_id = uts.friend_id
		WHERE users.username = $1 AND u.date_of_birth = $2;`, followerUsername, time.Now().Format(time.DateOnly))
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("query err: %v", err)
	}
	str := ""
	for rows.Next() {
		err = rows.Scan(&str)
		if err != nil {
			return nil, fmt.Errorf("scan err: %v", err)
		}
		userNames = append(userNames, str)
	}
	return userNames, nil
}
