package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golab9/internal/config"
	"golab9/internal/domain/models"
	"net/http"
)

type Auth struct {
	client *http.Client
	addr   string
}

func NewAuth(cfg *config.Config) *Auth {
	return &Auth{
		client: http.DefaultClient,
		addr:   cfg.ServerAddr + "/api/v1/auth",
	}
}

func (a *Auth) Register(ctx context.Context, request models.RegisterReq) (models.RegisterResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return models.RegisterResponse{}, fmt.Errorf("ошибка при маршализации данных: %v", err)
	}

	reqURL := fmt.Sprintf("%s/%s", a.addr, "register")
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return models.RegisterResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return models.RegisterResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.RegisterResponse{}, fmt.Errorf("ошибка: статус ответа %s", resp.Status)
	}

	var response models.RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return models.RegisterResponse{}, err
	}

	return models.RegisterResponse{}, nil
}

func (a *Auth) Login(ctx context.Context, request models.LoginReq) (models.LoginResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf("ошибка при маршализации данных: %v", err)
	}

	reqURL := fmt.Sprintf("%s/%s", a.addr, "login")
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return models.LoginResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return models.LoginResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.LoginResponse{}, fmt.Errorf("ошибка: статус ответа %s", resp.Status)
	}

	var response models.LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return models.LoginResponse{}, err
	}

	return response, nil
}
