package templates

type Param struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Enabled     bool   `json:"enabled"`
	Overridable bool   `json:"overridable"`
}

type Auth struct {
	Scheme       string `json:"scheme"`
	Label        string `json:"label"`
	Description  string `json:"description"`
	ExampleValue string `json:"exampleValue,omitempty"`
}

type Body struct {
	Mode        string  `json:"mode"`
	ContentType string  `json:"contentType,omitempty"`
	Example     any     `json:"example,omitempty"`
	Raw         string  `json:"raw,omitempty"`
	FormFields  []Param `json:"formFields,omitempty"`
}

type Request struct {
	Method      string  `json:"method"`
	URL         string  `json:"url"`
	QueryParams []Param `json:"queryParams"`
	Headers     []Param `json:"headers"`
	Auth        Auth    `json:"auth"`
	Body        Body    `json:"body"`
}

type ResponsePreview struct {
	Status      int    `json:"status"`
	ContentType string `json:"contentType"`
	DurationMS  int    `json:"durationMs"`
	Size        string `json:"size"`
	Body        any    `json:"body"`
}

type Template struct {
	Slug             string          `json:"slug"`
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	Category         string          `json:"category"`
	Tags             []string        `json:"tags"`
	GuestSafe        bool            `json:"guestSafe"`
	AllowedOverrides []string        `json:"allowedOverrides"`
	Request          Request         `json:"request"`
	ResponsePreview  ResponsePreview `json:"responsePreview"`
}

var all = []Template{
	{
		Slug:             "jsonplaceholder-posts",
		Title:            "JSONPlaceholder Posts",
		Description:      "Read and inspect a sample REST collection from the JSONPlaceholder demo API.",
		Category:         "REST basics",
		Tags:             []string{"guest-safe", "json", "read-only"},
		GuestSafe:        true,
		AllowedOverrides: []string{"queryParams.userId", "headers.accept"},
		Request: Request{
			Method: "GET",
			URL:    "https://jsonplaceholder.typicode.com/posts/1",
			QueryParams: []Param{
				{Name: "userId", Value: "1", Enabled: true, Overridable: true},
			},
			Headers: []Param{
				{Name: "accept", Value: "application/json", Enabled: true, Overridable: true},
			},
			Auth: Auth{
				Scheme:      "none",
				Label:       "None",
				Description: "No authentication required for the public demo endpoint.",
			},
			Body: Body{
				Mode: "none",
			},
		},
		ResponsePreview: ResponsePreview{
			Status:      200,
			ContentType: "application/json; charset=utf-8",
			DurationMS:  186,
			Size:        "1.2 KB",
			Body: map[string]any{
				"id":        1,
				"title":     "delectus aut autem",
				"completed": false,
			},
		},
	},
	{
		Slug:             "github-repo-details",
		Title:            "GitHub Repository Details",
		Description:      "Inspect a public repository payload and response headers from the GitHub API.",
		Category:         "Authentication flows",
		Tags:             []string{"guest-safe", "headers", "api"},
		GuestSafe:        true,
		AllowedOverrides: []string{"queryParams.repo", "queryParams.owner"},
		Request: Request{
			Method: "GET",
			URL:    "https://api.github.com/repos/octocat/Hello-World",
			QueryParams: []Param{
				{Name: "owner", Value: "octocat", Enabled: true, Overridable: true},
				{Name: "repo", Value: "Hello-World", Enabled: true, Overridable: true},
			},
			Headers: []Param{
				{Name: "accept", Value: "application/vnd.github+json", Enabled: true, Overridable: true},
				{Name: "user-agent", Value: "api-testing-kit", Enabled: true, Overridable: false},
			},
			Auth: Auth{
				Scheme:      "none",
				Label:       "None",
				Description: "Public repository metadata can be fetched without credentials for demo purposes.",
			},
			Body: Body{
				Mode: "none",
			},
		},
		ResponsePreview: ResponsePreview{
			Status:      200,
			ContentType: "application/json; charset=utf-8",
			DurationMS:  212,
			Size:        "3.8 KB",
			Body: map[string]any{
				"name":       "Hello-World",
				"full_name":  "octocat/Hello-World",
				"private":    false,
				"visibility": "public",
			},
		},
	},
	{
		Slug:             "weather-demo",
		Title:            "Weather Demo",
		Description:      "Show a small forecast payload with a query-driven public weather API example.",
		Category:         "CRUD examples",
		Tags:             []string{"guest-safe", "query", "forecast"},
		GuestSafe:        true,
		AllowedOverrides: []string{"queryParams.city"},
		Request: Request{
			Method: "GET",
			URL:    "https://api.open-meteo.com/v1/forecast",
			QueryParams: []Param{
				{Name: "latitude", Value: "22.5726", Enabled: true, Overridable: false},
				{Name: "longitude", Value: "88.3639", Enabled: true, Overridable: false},
				{Name: "current", Value: "temperature_2m", Enabled: true, Overridable: true},
			},
			Headers: []Param{
				{Name: "accept", Value: "application/json", Enabled: true, Overridable: true},
			},
			Auth: Auth{
				Scheme:      "none",
				Label:       "None",
				Description: "The demo endpoint is intentionally public and uses only safe query overrides.",
			},
			Body: Body{
				Mode: "none",
			},
		},
		ResponsePreview: ResponsePreview{
			Status:      200,
			ContentType: "application/json",
			DurationMS:  164,
			Size:        "2.1 KB",
			Body: map[string]any{
				"current": map[string]any{
					"temperature_2m": 28.4,
					"time":           "2026-03-25T12:00",
				},
			},
		},
	},
	{
		Slug:             "auth-flow-mock",
		Title:            "Auth Flow Mock",
		Description:      "Demonstrate a login-style request body without exposing any private target or secret.",
		Category:         "Authentication flows",
		Tags:             []string{"guest-safe", "form", "mock"},
		GuestSafe:        true,
		AllowedOverrides: []string{"body.example.email"},
		Request: Request{
			Method: "POST",
			URL:    "https://jsonplaceholder.typicode.com/posts",
			QueryParams: []Param{
				{Name: "mode", Value: "demo", Enabled: true, Overridable: false},
			},
			Headers: []Param{
				{Name: "content-type", Value: "application/json", Enabled: true, Overridable: false},
			},
			Auth: Auth{
				Scheme:       "bearer",
				Label:        "Bearer",
				Description:  "A placeholder auth flow that keeps the interface realistic without requiring credentials.",
				ExampleValue: "demo-token",
			},
			Body: Body{
				Mode:        "json",
				ContentType: "application/json",
				Example: map[string]any{
					"email":    "guest@example.dev",
					"password": "••••••••",
				},
				Raw: `{
  "email": "guest@example.dev",
  "password": "••••••••"
}`,
			},
		},
		ResponsePreview: ResponsePreview{
			Status:      201,
			ContentType: "application/json; charset=utf-8",
			DurationMS:  194,
			Size:        "882 B",
			Body: map[string]any{
				"id":     101,
				"status": "created",
			},
		},
	},
	{
		Slug:             "pagination-error-handling",
		Title:            "Pagination and Error Handling",
		Description:      "Pair a paginated request with a sample error response for troubleshooting flows.",
		Category:         "Error handling",
		Tags:             []string{"guest-safe", "pagination", "errors"},
		GuestSafe:        true,
		AllowedOverrides: []string{"queryParams.page", "queryParams.limit"},
		Request: Request{
			Method: "GET",
			URL:    "https://jsonplaceholder.typicode.com/comments",
			QueryParams: []Param{
				{Name: "page", Value: "1", Enabled: true, Overridable: true},
				{Name: "limit", Value: "10", Enabled: true, Overridable: true},
			},
			Headers: []Param{
				{Name: "accept", Value: "application/json", Enabled: true, Overridable: true},
			},
			Auth: Auth{
				Scheme:      "none",
				Label:       "None",
				Description: "Guest-safe pagination examples should remain readable and low-risk.",
			},
			Body: Body{
				Mode: "none",
			},
		},
		ResponsePreview: ResponsePreview{
			Status:      404,
			ContentType: "application/json",
			DurationMS:  141,
			Size:        "724 B",
			Body: map[string]any{
				"error": map[string]any{
					"code":    "not_found",
					"message": "The requested resource could not be found.",
				},
			},
		},
	},
}

var bySlug map[string]Template

func init() {
	bySlug = make(map[string]Template, len(all))
	for _, template := range all {
		bySlug[template.Slug] = template
	}
}

func List() []Template {
	items := make([]Template, len(all))
	copy(items, all)
	return items
}

func Get(slug string) (Template, bool) {
	template, ok := bySlug[slug]
	return template, ok
}
