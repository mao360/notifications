package delivery

import (
	"github.com/mao360/notifications/pkg/service"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"testing"
)

func TestHandler_Auth(t *testing.T) {
	type fields struct {
		service service.ServiceI
		sugared *zap.SugaredLogger
	}
	type args struct {
		next http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				service: tt.fields.service,
				sugared: tt.fields.sugared,
			}
			if got := h.Auth(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_Authorization(t *testing.T) {
	type fields struct {
		service service.ServiceI
		sugared *zap.SugaredLogger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				service: tt.fields.service,
				sugared: tt.fields.sugared,
			}
			h.Authorization(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_GetNotification(t *testing.T) {
	type fields struct {
		service service.ServiceI
		sugared *zap.SugaredLogger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				service: tt.fields.service,
				sugared: tt.fields.sugared,
			}
			h.GetNotification(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_Registration(t *testing.T) {
	type fields struct {
		service service.ServiceI
		sugared *zap.SugaredLogger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				service: tt.fields.service,
				sugared: tt.fields.sugared,
			}
			h.Registration(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_Subscribe(t *testing.T) {
	type fields struct {
		service service.ServiceI
		sugared *zap.SugaredLogger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				service: tt.fields.service,
				sugared: tt.fields.sugared,
			}
			h.Subscribe(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_Unsubscribe(t *testing.T) {
	type fields struct {
		service service.ServiceI
		sugared *zap.SugaredLogger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				service: tt.fields.service,
				sugared: tt.fields.sugared,
			}
			h.Unsubscribe(tt.args.w, tt.args.r)
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		service service.ServiceI
		sugared *zap.SugaredLogger
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.service, tt.args.sugared); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
