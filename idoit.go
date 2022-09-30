package goidoit

import (
	"context"
)

type IdoitService service

type VersionLoginResult struct {
	Userid   string `json:"userid"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Username string `json:"username"`
	Tenant   string `json:"tenant"`
	Language string `json:"language"`
}

type VersionResult struct {
	Login   VersionLoginResult `json:"login"`
	Version string             `json:"version"`
	Step    string             `json:"step"`
	Type    string             `json:"type"`
}

// https://kb.i-doit.com/display/en/Methods#Methods-idoit.search
func (c *IdoitService) Version(ctx context.Context) (*VersionResult, error) {
	return parse[*VersionResult](c.client.Request(ctx, "idoit.version", nil))
}

func (c *IdoitService) Search(ctx context.Context, query string) ([]SearchResult, error) {
	params := struct {
		Query string `json:"q"`
	}{query}
	return parse[[]SearchResult](c.client.Request(ctx, "idoit.search", &params))
}

type ConstantsResponse struct {
	ObjectTypes   map[string]string            `json:"objectTypes"`
	Categories    map[string]map[string]string `json:"categories"`
	RecordStates  map[string]string            `json:"recordStates"`
	RelationTypes map[string]string            `json:"relationTypes"`
	StaticObjects map[string]string            `json:"staticObjects"`
}

// https://kb.i-doit.com/display/en/Methods#Methods-idoit.constants
func (c *IdoitService) Constants(ctx context.Context) (*ConstantsResponse, error) {
	return parse[*ConstantsResponse](c.client.Request(ctx, "idoit.constants", nil))
}

type AddonsResponseEntry struct {
	Title   string `json:"title"`
	Key     string `json:"key"`
	Version string `json:"version"`
	Author  struct {
		Name string `json:"name"`
	} `json:"author"`
	Licensed  bool `json:"licensed"`
	Active    bool `json:"active"`
	Installed bool `json:"installed"`
}

type AddonsResponseWrapper struct {
	Success bool                  `json:"success"`
	Result  []AddonsResponseEntry `json:"result"`
}

func (c *IdoitService) Addons(ctx context.Context) ([]AddonsResponseEntry, error) {
	w, err := parse[AddonsResponseWrapper](c.client.Request(ctx, "idoit.addons", nil))
	if err != nil {
		return nil, err
	}
	return w.Result, nil
}

type LicenseResponseEntry struct {
	Id               int    `json:"id"`
	Label            string `json:"label"`
	LicenseType      string `json:"licenseType"`
	RegistrationDate string `json:"registrationDate"`
	ValidUntil       string `json:"validUntil"`
	Objects          int    `json:"objects"`
	Tenants          int    `json:"tenants"`
	Environment      string `json:"environment"`
	Valid            bool   `json:"valid"`
}

type LicenseResponseObjectCapacity struct {
	Total int `json:"total"`
	InUse int `json:"inUse"`
}

type LicenseResponseAddonsEntry struct {
	Label    string `json:"label"`
	Licensed bool   `json:"licensed"`
}

type LicenseResponse struct {
	ObjectCapacity LicenseResponseObjectCapacity         `json:"objectCapacity"`
	Addons         map[string]LicenseResponseAddonsEntry `json:"addons"`
	Licenses       []LicenseResponseEntry                `json:"licenses"`
}

func (c *IdoitService) License(ctx context.Context) (*LicenseResponse, error) {
	return parse[*LicenseResponse](c.client.Request(ctx, "idoit.license", nil))
}
