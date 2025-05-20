package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// PagedResult 分页结果
type PagedResult[T any] struct {
	Items    []T `json:"items"`
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// Dataset 泛型数据集
type Dataset[T any] struct {
	baseURL   string
	client    *http.Client
	headers   map[string]string
	authToken string
	errorHook func(error)
	timeout   time.Duration
}

// DatasetOption 配置选项函数类型
type DatasetOption func(*Dataset[any])

// WithTimeout 设置请求超时时间
func WithTimeout(timeout time.Duration) DatasetOption {
	return func(d *Dataset[any]) {
		d.timeout = timeout
	}
}

// WithHeaders 设置请求头
func WithHeaders(headers map[string]string) DatasetOption {
	return func(d *Dataset[any]) {
		d.headers = headers
	}
}

// WithAuthToken 设置认证令牌
func WithAuthToken(token string) DatasetOption {
	return func(d *Dataset[any]) {
		d.authToken = token
	}
}

// WithErrorHook 设置错误处理钩子
func WithErrorHook(hook func(error)) DatasetOption {
	return func(d *Dataset[any]) {
		d.errorHook = hook
	}
}

// NewDataset 创建一个新的数据集
func NewDataset[T any](baseURL string, opts ...DatasetOption) *Dataset[T] {
	d := &Dataset[T]{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		headers: make(map[string]string),
	}

	// 应用配置选项
	for _, opt := range opts {
		opt((*Dataset[any])(d))
	}

	if d.timeout > 0 {
		d.client.Timeout = d.timeout
	}

	return d
}

// List 获取分页数据列表
func (d *Dataset[T]) List(ctx context.Context, page, pageSize int, filter map[string]any) (*PagedResult[T], error) {
	u, err := url.Parse(d.baseURL)
	if err != nil {
		return nil, fmt.Errorf("解析URL失败: %w", err)
	}

	// 添加查询参数
	params := url.Values{}
	params.Add("page", fmt.Sprintf("%d", page))
	params.Add("pageSize", fmt.Sprintf("%d", pageSize))

	for key, value := range filter {
		if value != nil {
			params.Add(key, fmt.Sprintf("%v", value))
		}
	}

	u.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	d.setRequestHeaders(req)

	resp, err := d.client.Do(req)
	if err != nil {
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := d.handleHTTPError(resp)
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, err
	}

	var result PagedResult[T]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// Get 获取单个记录
func (d *Dataset[T]) Get(ctx context.Context, id string) (*T, error) {
	u := fmt.Sprintf("%s/%s", d.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	d.setRequestHeaders(req)

	resp, err := d.client.Do(req)
	if err != nil {
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := d.handleHTTPError(resp)
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, err
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// Create 创建记录
func (d *Dataset[T]) Create(ctx context.Context, data any) (*T, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("序列化数据失败: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	d.setRequestHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	resp, err := d.client.Do(req)
	if err != nil {
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		err := d.handleHTTPError(resp)
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, err
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// Update 更新记录
func (d *Dataset[T]) Update(ctx context.Context, id string, data any) (*T, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("序列化数据失败: %w", err)
	}

	u := fmt.Sprintf("%s/%s", d.baseURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	d.setRequestHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	resp, err := d.client.Do(req)
	if err != nil {
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := d.handleHTTPError(resp)
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, err
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// Patch 部分更新记录
func (d *Dataset[T]) Patch(ctx context.Context, id string, data any) (*T, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("序列化数据失败: %w", err)
	}

	u := fmt.Sprintf("%s/%s", d.baseURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, u, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	d.setRequestHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	resp, err := d.client.Do(req)
	if err != nil {
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := d.handleHTTPError(resp)
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, err
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}

// Delete 删除记录
func (d *Dataset[T]) Delete(ctx context.Context, id string) error {
	u := fmt.Sprintf("%s/%s", d.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	d.setRequestHeaders(req)

	resp, err := d.client.Do(req)
	if err != nil {
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		err := d.handleHTTPError(resp)
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return err
	}

	return nil
}

// BatchDelete 批量删除记录
func (d *Dataset[T]) BatchDelete(ctx context.Context, ids []string) error {
	jsonData, err := json.Marshal(map[string][]string{"ids": ids})
	if err != nil {
		return fmt.Errorf("序列化数据失败: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, d.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	d.setRequestHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	resp, err := d.client.Do(req)
	if err != nil {
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		err := d.handleHTTPError(resp)
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return err
	}

	return nil
}

// CustomRequest 自定义请求
func (d *Dataset[T]) CustomRequest(ctx context.Context, method, path string, body any) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("序列化数据失败: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	u := d.baseURL
	if path != "" {
		if path[0] != '/' {
			u += "/"
		}
		u += path
	}

	req, err := http.NewRequestWithContext(ctx, method, u, reqBody)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	d.setRequestHeaders(req)
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := d.client.Do(req)
	if err != nil {
		if d.errorHook != nil {
			d.errorHook(err)
		}
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}

	return resp, nil
}

// 设置请求头
func (d *Dataset[T]) setRequestHeaders(req *http.Request) {
	// 设置通用请求头
	for key, value := range d.headers {
		req.Header.Set(key, value)
	}

	// 设置认证头
	if d.authToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", d.authToken))
	}
}

// 处理HTTP错误
func (d *Dataset[T]) handleHTTPError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取错误响应失败: %w", err)
	}

	// 尝试解析JSON格式的错误信息
	var errorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details any    `json:"details"`
	}

	if json.Unmarshal(body, &errorResponse) == nil && errorResponse.Message != "" {
		return fmt.Errorf("HTTP错误 %d: %s", resp.StatusCode, errorResponse.Message)
	}

	// 如果无法解析JSON，返回原始错误信息
	return fmt.Errorf("HTTP错误 %d: %s", resp.StatusCode, string(body))
}
