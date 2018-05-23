// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Task": Tasks Resource Client
//
// Command:
// $ goagen
// --design=github.com/BcRikko/learning-goa/design
// --out=$(GOPATH)/src/github.com/BcRikko/learning-goa
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateTasksPayload is the Tasks create action payload.
type CreateTasksPayload struct {
	// タスクのタイトル
	Title string `form:"title" json:"title" xml:"title"`
}

// CreateTasksPath computes a request path to the create action of Tasks.
func CreateTasksPath() string {

	return fmt.Sprintf("/api/tasks")
}

// タスクを作成する。
func (c *Client) CreateTasks(ctx context.Context, path string, payload *CreateTasksPayload) (*http.Response, error) {
	req, err := c.NewCreateTasksRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateTasksRequest create the request corresponding to the create action endpoint of the Tasks resource.
func (c *Client) NewCreateTasksRequest(ctx context.Context, path string, payload *CreateTasksPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// DeleteTasksPath computes a request path to the delete action of Tasks.
func DeleteTasksPath(taskID int) string {
	param0 := strconv.Itoa(taskID)

	return fmt.Sprintf("/api/tasks/%s", param0)
}

// 指定IDのタスクを削除する。
func (c *Client) DeleteTasks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteTasksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteTasksRequest create the request corresponding to the delete action endpoint of the Tasks resource.
func (c *Client) NewDeleteTasksRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListTasksPath computes a request path to the list action of Tasks.
func ListTasksPath() string {

	return fmt.Sprintf("/api/tasks")
}

// タスク一覧を取得する。
func (c *Client) ListTasks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListTasksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListTasksRequest create the request corresponding to the list action endpoint of the Tasks resource.
func (c *Client) NewListTasksRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowTasksPath computes a request path to the show action of Tasks.
func ShowTasksPath(taskID int) string {
	param0 := strconv.Itoa(taskID)

	return fmt.Sprintf("/api/tasks/%s", param0)
}

// 指定IDのタスクを取得する。
func (c *Client) ShowTasks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTasksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTasksRequest create the request corresponding to the show action endpoint of the Tasks resource.
func (c *Client) NewShowTasksRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateTasksPayload is the Tasks update action payload.
type UpdateTasksPayload struct {
	// タスクの完了状態
	Done bool `form:"done" json:"done" xml:"done"`
	// タスクのタイトル
	Title string `form:"title" json:"title" xml:"title"`
}

// UpdateTasksPath computes a request path to the update action of Tasks.
func UpdateTasksPath(taskID int) string {
	param0 := strconv.Itoa(taskID)

	return fmt.Sprintf("/api/tasks/%s", param0)
}

// 指定IDのタスクを更新する。
func (c *Client) UpdateTasks(ctx context.Context, path string, payload *UpdateTasksPayload) (*http.Response, error) {
	req, err := c.NewUpdateTasksRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateTasksRequest create the request corresponding to the update action endpoint of the Tasks resource.
func (c *Client) NewUpdateTasksRequest(ctx context.Context, path string, payload *UpdateTasksPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}
