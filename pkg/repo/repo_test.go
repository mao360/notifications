package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mao360/notifications/models"
	"reflect"
	"testing"
)

func TestNewRepo(t *testing.T) {
	type args struct {
		pool *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want *Repo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			if got := NewRepo(tt.args.pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_GetNotification(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		followerUsername string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				pool: tt.fields.pool,
			}
			got, err := r.GetNotification(tt.args.followerUsername)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNotification() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_GetUser(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		username     string
		passwordHash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				pool: tt.fields.pool,
			}
			got, err := r.GetUser(tt.args.username, tt.args.passwordHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_NewUser(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				pool: tt.fields.pool,
			}
			if err := r.NewUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_Subscribe(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		followerUsername string
		authorUsername   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				pool: tt.fields.pool,
			}
			if err := r.Subscribe(tt.args.followerUsername, tt.args.authorUsername); (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_Unsubscribe(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		followerUsername string
		authorUsername   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				pool: tt.fields.pool,
			}
			if err := r.Unsubscribe(tt.args.followerUsername, tt.args.authorUsername); (err != nil) != tt.wantErr {
				t.Errorf("Unsubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
