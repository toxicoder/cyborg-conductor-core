package types

// Cyborg represents a cyborg entity in the system
type Cyborg struct {
	ID              string `json:"id"`
	DisplayName     string `json:"display_name"`
	Category        string `json:"category"`
	SubCategory     string `json:"sub_category"`
	ShortDescription string `json:"short_description"`
	PrimaryFunctionality string `json:"primary_functionality"`
	DeterministicCapabilities []string `json:"deterministic_capabilities"`
	LLMCapabilities []string `json:"llm_capabilities"`
	Tags            []string `json:"tags"`
	DeploymentSpec  string `json:"deployment_spec"`
	SLALatencyBudget string `json:"sla_latency_budget"`
	MaxConcurrentStreams int `json:"max_concurrent_streams"`
	ReliabilityTier string `json:"reliability_tier"`
}

// CyborgStatus represents the status of a cyborg
type CyborgStatus string

const (
	StatusActive   CyborgStatus = "active"
	StatusInactive CyborgStatus = "inactive"
	StatusError    CyborgStatus = "error"
	StatusPending  CyborgStatus = "pending"
)

// Job represents a job that can be assigned to a cyborg
type Job struct {
	ID           string `json:"id"`
	CyborgID     string `json:"cyborg_id"`
	JobType      string `json:"job_type"`
	Priority     int    `json:"priority"`
	Status       CyborgStatus `json:"status"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Timeout      int64  `json:"timeout"`
	// Add more fields as needed
}

// Task represents a unit of work that can be executed by a cyborg
type Task struct {
	ID           string `json:"id"`
	JobID        string `json:"job_id"`
	CyborgID     string `json:"cyborg_id"`
	TaskType     string `json:"task_type"`
	Priority     int    `json:"priority"`
	Status       CyborgStatus `json:"status"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Timeout      int64  `json:"timeout"`
	// Add more fields as needed
}