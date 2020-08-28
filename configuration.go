/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit

import (
	"net/http"
)


// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	Url string
	Description string
	Variables map[string]ServerVariable
}

// Configuration stores the configuration of the API client
type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	ApiKey        string `json:"apiKey,omitempty"`
	TouchPoint	  *TouchPoint `json:"touchPoint,omitempty"`
	Servers       []ServerConfiguration
	HTTPClient    *http.Client

	username       string
	password       string
	defaultCulture string
}

type APIOption func(*Configuration)


func HTTPClient(client *http.Client) APIOption {
	return func(cfg *Configuration) {
		cfg.HTTPClient = client
	}
}

func DefaultHeader(key, value string) APIOption {
	return func(cfg *Configuration) {
		cfg.DefaultHeader[key] = value
	}
}

func Debug() APIOption {
	return func(cfg *Configuration) {
		cfg.Debug = true
	}
}

func DefaultCulture(culture string) APIOption {
	return func(cfg *Configuration) {
		cfg.defaultCulture = culture
	}
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) AddTouchPoint(tp *TouchPoint) {
	c.TouchPoint = tp
}

