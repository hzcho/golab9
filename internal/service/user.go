package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golab9/internal/config"
	"golab9/internal/domain/models"
	"net/http"
	"net/url"
	"strconv"
)

type User struct {
	client *http.Client
	addr   string
	token  string
}

func NewUser(cfg *config.Config) *User {
	return &User{
		client: http.DefaultClient,
		addr:   cfg.ServerAddr + "/api/v1/users",
	}
}

func (u *User) Get(ctx context.Context, filter models.GetUserFilter) ([]models.User, error) {
	params := url.Values{}

	if filter.Name != "" {
		params.Add("name", filter.Name)
	}
	if filter.Age >= 0 {
		params.Add("age", strconv.Itoa(filter.Age))
	}
	if filter.Page >= 0 {
		params.Add("page", strconv.Itoa(filter.Page))
	}
	if filter.Limit >= 0 {
		params.Add("limit", strconv.Itoa(filter.Limit))
	}

	reqURL := fmt.Sprintf("%s?%s", u.addr, params.Encode())

	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+u.token)

	resp, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ошибка: статус ответа %s", resp.Status)
	}

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) GetById(ctx context.Context, id uint64) (models.User, error) {
	reqURL := fmt.Sprintf("%s/%d", u.addr, id)

	fmt.Println(reqURL)

	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return models.User{}, err
	}

	resp, err := u.client.Do(req)
	if err != nil {
		return models.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.User{}, fmt.Errorf("Ошибка: статус ответа %s", resp.Status)
	}

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *User) Add(ctx context.Context, user models.AddUser) (uint64, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return 0, fmt.Errorf("ошибка при маршализации данных: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", u.addr, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", "Bearer "+u.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := u.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("ошибка: статус ответа %s", resp.Status)
	}

	var response struct {
		ID uint64 `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}

	return response.ID, nil
}

func (u *User) Update(ctx context.Context, user models.UpdateUser) (models.User, error) {
	reqURL := fmt.Sprintf("%s/%d", u.addr, user.ID)

	jsonData, err := json.Marshal(user)
	if err != nil {
		return models.User{}, fmt.Errorf("ошибка при маршализации данных: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", reqURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return models.User{}, err
	}
	req.Header.Add("Authorization", "Bearer "+u.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := u.client.Do(req)
	if err != nil {
		return models.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.User{}, fmt.Errorf("ошибка: статус ответа %s", resp.Status)
	}

	var updatedUser models.User
	if err := json.NewDecoder(resp.Body).Decode(&updatedUser); err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (u *User) Delete(ctx context.Context, id uint64) (bool, error) {
	reqURL := fmt.Sprintf("%s/%d", u.addr, id)

	req, err := http.NewRequestWithContext(ctx, "DELETE", reqURL, nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", "Bearer "+u.token)

	resp, err := u.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("ошибка: статус ответа %s", resp.Status)
	}

	return true, nil
}

func (u *User) SetToken(token string) error {
	u.token = token

	return nil
}
