package abuse

import (
	"encoding/json"
	"time"
)

type RecentFilter struct {
	UserID        *string
	SessionID     *string
	RequestRunID  *string
	Severity      Severity
	Category      Category
	SourceIP      *string
	Target        *string
	RuleKey       *string
	ActionTaken   Action
	CreatedAfter  *time.Time
	CreatedBefore *time.Time
	Limit         int32
}

type SummaryRow struct {
	Severity      Severity  `json:"severity"`
	Category      Category  `json:"category"`
	ActionTaken   Action    `json:"actionTaken"`
	Count         int64     `json:"count"`
	LastCreatedAt time.Time `json:"lastCreatedAt"`
}

type SummaryFilter struct {
	UserID        *string
	SessionID     *string
	RequestRunID  *string
	Severity      Severity
	Category      Category
	SourceIP      *string
	CreatedAfter  *time.Time
	CreatedBefore *time.Time
	Limit         int32
}

type BlockedTargetType string

const (
	BlockedTargetTypeDomain     BlockedTargetType = "domain"
	BlockedTargetTypeHost       BlockedTargetType = "host"
	BlockedTargetTypeIP         BlockedTargetType = "ip"
	BlockedTargetTypeCIDR       BlockedTargetType = "cidr"
	BlockedTargetTypeURLPattern BlockedTargetType = "url_pattern"
)

type BlockedTarget struct {
	ID              string            `json:"id"`
	TargetType      BlockedTargetType `json:"targetType"`
	TargetValue     string            `json:"targetValue"`
	Reason          string            `json:"reason"`
	Source          string            `json:"source"`
	IsActive        bool              `json:"isActive"`
	ExpiresAt       *time.Time        `json:"expiresAt,omitempty"`
	CreatedByUserID *string           `json:"createdByUserId,omitempty"`
	Metadata        json.RawMessage   `json:"metadata,omitempty"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
}

type BlockedTargetFilter struct {
	TargetType     BlockedTargetType
	TargetValue    string
	OnlyActive     bool
	IncludeExpired bool
	Limit          int32
}
