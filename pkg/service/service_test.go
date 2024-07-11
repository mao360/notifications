package service

import (
	"context"
	"github.com/mao360/notifications/models"
	"github.com/mao360/notifications/pkg/repo"
	"github.com/mao360/notifications/pkg/service/mocks"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		repo repo.RepoI // мокнуть
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks.NewServiceI(t)
			if got := NewService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GenerateToken(t *testing.T) {
	type fields struct {
		repo repo.RepoI
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.GenerateToken(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetNotification(t *testing.T) {
	type fields struct {
		repo repo.RepoI
	}
	type args struct {
		ctx      context.Context
		follower string
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
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.GetNotification(tt.args.ctx, tt.args.follower)
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

func TestService_GetUser(t *testing.T) {
	type fields struct {
		repo repo.RepoI
	}
	type args struct {
		ctx      context.Context
		username string
		password string
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
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.GetUser(tt.args.ctx, tt.args.username, tt.args.password)
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

func TestService_NewUser(t *testing.T) {
	type fields struct {
		repo repo.RepoI
	}
	type args struct {
		ctx  context.Context
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
			s := &Service{
				repo: tt.fields.repo,
			}
			if err := s.NewUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_ParseToken(t *testing.T) {
	type fields struct {
		repo repo.RepoI
	}
	type args struct {
		ctx         context.Context
		accessToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			got, err := s.ParseToken(tt.args.ctx, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Subscribe(t *testing.T) {
	type fields struct {
		repo repo.RepoI
	}
	type args struct {
		ctx      context.Context
		follower string
		author   string
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
			s := &Service{
				repo: tt.fields.repo,
			}
			if err := s.Subscribe(tt.args.ctx, tt.args.follower, tt.args.author); (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Unsubscribe(t *testing.T) {
	type fields struct {
		repo repo.RepoI
	}
	type args struct {
		ctx      context.Context
		follower string
		author   string
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
			s := &Service{
				repo: tt.fields.repo,
			}
			if err := s.Unsubscribe(tt.args.ctx, tt.args.follower, tt.args.author); (err != nil) != tt.wantErr {
				t.Errorf("Unsubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_generatePasswordHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePasswordHash(tt.args.password); got != tt.want {
				t.Errorf("generatePasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
