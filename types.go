package gomarathon

// RequestOptions passed for query api
type RequestOptions struct {
	Method string
	Path   string
	Datas  interface{}
	Params *Parameters
}

// Parameters to build url query
type Parameters struct {
	Cmd         string
	Host        string
	Scale       bool
	CallbackURL string
}

// Response representation of a full marathon response
type Response struct {
	Code int
	//Groups   []*Group       `json:",omitempty"`
	Apps       []*Application `json:"apps,omitempty"`
	App        *Application   `json:"app,omitempty"`
	Versions   []string       `json:",omitempty"`
	Tasks      []*Task        `json:"tasks,omitempty"`
	Queues     []*Queue       `json:"queue,omitempty"`
	AppVersion Application
}

type Queue struct {
	Count int          `json:"count,omitempty"`
	Delay *Delayment   `json:"delay,omitempty"`
	App   *Application `json:"app,omitempty"`
}

type Delayment struct {
	Overdue bool `json:"overdue,omitempty"`
}

// https://github.com/mesosphere/marathon/blob/master/REST.md#apps
type Application struct {
	ID                    string             `json:"id"`
	Cmd                   string             `json:"cmd,omitempty"`
	Constraints           [][]string         `json:"constraints"`
	Container             *Container         `json:"container"`
	CPUs                  float32            `json:"cpus"`
	Deployments           []*Deployment      `json:"deployments"`
	Env                   map[string]string  `json:"env"`
	Executor              string             `json:"executor"`
	HealthChecks          []*HealthCheck     `json:"healthChecks"`
	Instances             int                `json:"instances"`
	Mem                   float32            `json:"mem"`
	Tasks                 []*Task            `json:"tasks"`
	Ports                 []int              `json:"ports"`
	RequirePorts          bool               `json:"requirePorts"`
	BackoffSeconds        float64            `json:"backoffSeconds"`
	BackoffFactor         float32            `json:"backoffFactor"`
	MaxLaunchDelaySeconds float64            `json:"maxLaunchDelaySeconds"`
	TasksHealthy          int                `json:"tasksHealthy"`
	TasksRunning          int                `json:"tasksRunning"`
	TasksUnHealthy        int                `json:"tasksUnHealthy"`
	TasksStaged           int                `json:"tasksStaged"`
	UpgradeStrategy       *UpgradeStrategy   `json:"upgradeStrategy"`
	Uris                  []string           `json:"uris"`
	Version               string             `json:"version"`
	VersionInfo           *VersionInfomation `json:"versionInfo"`
	LastTaskFailure       *LastTaskFailure   `json:"lastTaskFailure"`
}

type LastTaskFailure struct {
	LastConfigChangeAt string `json:"appId"`
	Host               string `json:"host"`
	Message            string `json:"message"`
	SlaveId            string `json:"slaveId"`
	State              string `json:"state"`
	TaskId             string `json:"taskId"`
	TimeStamp          string `json:"timestamp"`
	Version            string `json:"version"`
}

type VersionInfomation struct {
	LastConfigChangeAt string `json:"lastConfigChangeAt"`
	LastScalingAt      string `json:"lastScalingAt"`
}

// Container is docker parameters
type Container struct {
	Type    string    `json:"type,omitempty"`
	Docker  *Docker   `json:"docker,omitempty"`
	Volumes []*Volume `json:"volumes,omitempty"`
}

// Docker options
type Docker struct {
	Image          string         `json:"image,omitempty"`
	ForcePullImage bool           `json:"forcePullImage,omitempty"`
	Privileged     bool           `json:"privileged,omitempty"`
	Network        string         `json:"network,omitempty"`
	PortMappings   []*PortMapping `json:"portMappings,omitempty"`
}

// Volume is used for mounting a host directory as a container volume
type Volume struct {
	ContainerPath string `json:"containerPath,omitempty"`
	HostPath      string `json:"hostPath,omitempty"`
	Mode          string `json:"mode,omitempty"`
}

// Container PortMappings
type PortMapping struct {
	ContainerPort int    `json:"containerPort,omitempty"`
	HostPort      int    `json:"hostPort,omitempty"`
	ServicePort   int    `json:"servicePort,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

// UpgradeStrategy has a minimumHealthCapacity which defines the minimum number of healty nodes
type UpgradeStrategy struct {
	MinimumHealthCapacity float32 `json:"minimumHealthCapacity,omitempty"`
}

// HealthCheck is described here:
// https://github.com/mesosphere/marathon/blob/master/REST.md#healthchecks
type HealthCheck struct {
	Protocol           string `json:"protocol,omitempty"`
	Path               string `json:"path,omitempty"`
	GracePeriodSeconds int    `json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds    int    `json:"intervalSeconds,omitempty"`
	PortIndex          int    `json:"portIndex,omitempty"`
	TimeoutSeconds     int    `json:"timeoutSeconds,omitempty"`
}

// Task is described here:
// https://github.com/mesosphere/marathon/blob/master/REST.md#tasks
type Task struct {
	AppID     string `json:"appId"`
	Host      string `json:"host"`
	ID        string `json:"id"`
	Ports     []int  `json:"ports"`
	StagedAt  string `json:"stagedAt"`
	StartedAt string `json:"startedAt"`
	Version   string `json:"version"`
}

// Deployment is described here:
// https://mesosphere.github.io/marathon/docs/rest-api.html#get-/v2/deployments
type Deployment struct {
	AffectedApps   []string          `json:"affectedApps"`
	ID             string            `json:"id"`
	Steps          []*DeploymentStep `json:"steps"`
	CurrentActions []*DeploymentStep `json:"currentActions"`
	CurrentStep    int               `json:"currentStep"`
	TotalSteps     int               `json:"totalSteps"`
	Version        string            `json:"version"`
}

// Deployment steps
type DeploymentStep struct {
	Action string `json:"action"`
	App    string `json:"app"`
}

// EventSubscription is described here:
// https://github.com/mesosphere/marathon/blob/master/REST.md#event-subscriptions
type EventSubscription struct {
	CallbackURL  string   `json:"CallbackUrl"`
	ClientIP     string   `json:"ClientIp"`
	EventType    string   `json:"eventType"`
	CallbackURLs []string `json:"CallbackUrls"`
}
