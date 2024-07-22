package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mao360/notifications/models"
	"github.com/mao360/notifications/pkg/service"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	service service.ServiceI
	sugared *zap.SugaredLogger
}

func NewHandler(service service.ServiceI, sugared *zap.SugaredLogger) *Handler {
	return &Handler{
		service: service,
		sugared: sugared,
	}
}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	h.sugared.Infof("handler started: registration")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "read body error", err)
		return
	}
	err = r.Body.Close()
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "close body error", err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "unmarshal error", err)
		return
	}
	err = h.service.NewUser(r.Context(), &user)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "newUser error", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.sugared.Infof("handler completed: registration")
}
func (h *Handler) Authorization(w http.ResponseWriter, r *http.Request) {
	h.sugared.Infof("handler started: authorization")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "read body error", err)
		return
	}
	err = r.Body.Close()
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "close body error", err)
		return
	}
	type form struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	f := form{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "unmarshal error", err)
		return
	}
	user, err := h.service.GetUser(r.Context(), f.UserName, f.Password)
	if err != nil || user == nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "getUser error", err)
		return
	}
	token, err := h.service.GenerateToken(r.Context(), f.UserName, f.Password)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "generateToken error", err)
		return
	}

	data, err := json.Marshal(map[string]interface{}{"token": token})
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "marshal error", err)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "write error", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	h.sugared.Infof("handler completed: authorization")
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	h.sugared.Infof("handler started: subscribe")
	user, err := CheckContext(r.Context())
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "checkContext error", err)
		return
	}
	username := r.URL.Query().Get("username")
	h.sugared.Infof("query param is: %s", username)
	h.sugared.Infof("username(follower) is: %s", user.UserName)
	err = h.service.Subscribe(r.Context(), user.UserName, username)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "subscribe error", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.sugared.Infof("handler completed: subscribe")
}

func (h *Handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	h.sugared.Infof("handler started: unsubscribe")
	user, err := CheckContext(r.Context())
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "checkContext error", err)
		return
	}
	username := r.URL.Query().Get("username")
	err = h.service.Unsubscribe(r.Context(), user.UserName, username)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "unsubscribe error", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.sugared.Infof("handler completed: unsubscribe")
}

func (h *Handler) GetNotification(w http.ResponseWriter, r *http.Request) {
	h.sugared.Infof("handler started: getNotification")
	user, err := CheckContext(r.Context())
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "checkContext error", err)
		return
	}
	users, err := h.service.GetNotification(r.Context(), user.UserName)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "getNotificationError", err)
		return
	}
	data, err := json.Marshal(users)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "marshal error", err)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		ErrResponseFunc(h.sugared, w, http.StatusInternalServerError, "write error", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	h.sugared.Infof("handler completed: getNotification")
}

func CheckContext(ctx context.Context) (*models.User, error) {
	var key ContextKey = "user"
	data := ctx.Value(key)
	if data == nil {
		return nil, errors.New("empty context")
	}
	user, ok := data.(*models.User)
	if !ok {
		return nil, errors.New("invalid context")
	}
	return user, nil
}

func ErrResponseFunc(logger *zap.SugaredLogger, w http.ResponseWriter, code int, message string, err error) {
	logger.Errorf("%s:%s", message, err.Error())
	resp, err := json.Marshal(map[string]interface{}{
		"message": message,
		"error":   err.Error(),
	})
	if err != nil {
		logger.Errorf("err resp func err:%s", err.Error())
	}
	w.WriteHeader(code)
	_, err = w.Write(resp)
	if err != nil {
		logger.Errorf("err resp func err:%s", err.Error())
	}
}
