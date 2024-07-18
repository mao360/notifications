package repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mao360/notifications/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.30.1 --name=RepoI
type RepoI interface {
	GetUser(username, passwordHash string) (bool, error)
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

func (r *Repo) GetUser(username, passwordHash string) (bool, error) {
	user := &models.User{}
	rows, err := r.pool.Query(context.Background(),
		`SELECT * FROM users
		WHERE username = '$1';`, username)
	if err != nil {
		return false, err
	}
	err = rows.Scan(&user.ID, &user.UserName, &user.Password, &user.DateOfBirth)
	if err != nil {
		return false, errors.New("no user")
	}
	if user.Password != passwordHash {
		return false, errors.New("bad password")
	}
	return true, nil
}

func (r *Repo) NewUser(user *models.User) error {
	rows, err := r.pool.Query(context.Background(),
		`SELECT * FROM users
		WHERE username = '$1'`, user.UserName)
	if err != nil {
		return err
	}
	if rows != nil {
		return errors.New("user already exists")
	}
	_, err = r.pool.Exec(context.Background(),
		`INSERT INTO users (password, username, date_of_birth)
		VALUES ('$1', '$2', '$3')`, user.Password, user.UserName, user.DateOfBirth)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Subscribe(followerUsername, authorUsername string) error {
	var followerID, authorID int
	rows, err := r.pool.Query(context.Background(),
		`SELECT count(*) FROM users
		JOIN users_to_subscribers uts ON users.user_id = uts.user_id
		JOIN users u on u.user_id = uts.friend_id
		WHERE users.username = '$1' AND u.username = '$2';`, followerUsername, authorUsername)
	if err != nil {
		return err
	}
	if rows != nil {
		return errors.New("already subscribed")
	}

	rows, err = r.pool.Query(context.Background(),
		`SELECT user_id FROM users
		WHERE username = '$1';`, followerUsername)
	if err != nil {
		return err
	}
	err = rows.Scan(&followerID)
	if err != nil {
		return err
	}

	rows, err = r.pool.Query(context.Background(),
		`SELECT user_id FROM users
		WHERE username = '$1';`, authorUsername)
	if err != nil {
		return err
	}
	err = rows.Scan(&authorID)
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(context.Background(),
		`INSERT INTO users_to_subscribers (user_id, friend_id)
			VALUES ('$1', '$2');`, followerID, authorID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Unsubscribe(followerUsername, authorUsername string) error {
	rows, err := r.pool.Query(context.Background(),
		`SELECT (*) FROM users
		JOIN users_to_subscribers uts ON users.user_id = uts.user_id
		JOIN users u on u.user_id = uts.friend_id
		WHERE users.username = '$1' AND u.username = '$2';`, followerUsername, authorUsername)
	if err != nil {
		return err
	}
	if rows == nil {
		return errors.New("no subscription")
	}
	_, err = r.pool.Exec(context.Background(),
		`DELETE FROM users_to_subscribers
	WHERE user_id = (
        SELECT user_id FROM users
        WHERE username = '$1'
    ) AND friend_id = (
        SELECT user_id FROM users
        WHERE username = '$2'
    );`, followerUsername, authorUsername)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetNotification(followerUsername string) ([]string, error) {
	_, err := r.pool.Query(context.Background(),
		`SELECT * FROM users;`)
	if err != nil {
		return nil, err
	}
	return nil, nil
	//userNames := make([]string, 0)
	//rows, err := r.pool.Query(context.Background(),
	//	`SELECT u.username FROM users
	//	JOIN users_to_subscribers uts ON users.user_id = uts.user_id
	//	JOIN users u on u.user_id = uts.friend_id
	//	WHERE users.username = '$1' AND u.date_of_birth = '$2';`, followerUsername, time.Now().Format(time.DateOnly))
	//if err != nil {
	//	return nil, err
	//}
	//err = rows.Scan(&userNames)
	//if err != nil {
	//	return nil, err
	//}
	//return userNames, nil
}
