package github

// RawHook is the raw github structure
type RawHook struct {
	ID        string
	Event     string
	Signature string
	Payload   []byte
}

// PingEvent is the predefined structure of the incoming github ping
type PingEvent struct {
	Zen        string         `json:"zen"`
	HookID     int64          `json:"hook_id"`
	Hook       Hook           `json:"hook"`
	Repository PingRepository `json:"repository"`
	Sender     Sender         `json:"sender"`
}

// PushEvent is the predefined structure of the incoming github ping
type PushEvent struct {
	Ref        string         `json:"ref"`
	Before     string         `json:"before"`
	After      string         `json:"after"`
	Created    bool           `json:"created"`
	Deleted    bool           `json:"deleted"`
	Forced     bool           `json:"forced"`
	BaseRef    interface{}    `json:"base_ref"`
	Compare    string         `json:"compare"`
	Commits    []Commit       `json:"commits"`
	HeadCommit Commit         `json:"head_commit"`
	Repository PushRepository `json:"repository"`
	Pusher     Pusher         `json:"pusher"`
	Sender     Sender         `json:"sender"`
}

// Hook is part of the PingEvent Payload
type Hook struct {
	Type         string       `json:"type"`
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	Active       bool         `json:"active"`
	Events       []string     `json:"events"`
	Config       Config       `json:"config"`
	UpdatedAt    string       `json:"updated_at"`
	CreatedAt    string       `json:"created_at"`
	URL          string       `json:"url"`
	TestURL      string       `json:"test_url"`
	PingURL      string       `json:"ping_url"`
	LastResponse LastResponse `json:"last_response"`
}

// Commit is part of the push event
type Commit struct {
	ID        string        `json:"id"`
	TreeID    string        `json:"tree_id"`
	Distinct  bool          `json:"distinct"`
	Message   string        `json:"message"`
	Timestamp string        `json:"timestamp"`
	URL       string        `json:"url"`
	Author    Author        `json:"author"`
	Committer Author        `json:"committer"`
	Added     []interface{} `json:"added"`
	Removed   []interface{} `json:"removed"`
	Modified  []string      `json:"modified"`
}

// Config is part of the Hook Payload
type Config struct {
	ContentType string `json:"content_type"`
	InsecureSSL string `json:"insecure_ssl"`
	Secret      string `json:"secret"`
	URL         string `json:"url"`
}

// LastResponse is the last response data from the hook payload
type LastResponse struct {
	Code    interface{} `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

// PushRepository will define the incoming repo info
type PushRepository struct {
	ID               int64       `json:"id"`
	NodeID           string      `json:"node_id"`
	Name             string      `json:"name"`
	FullName         string      `json:"full_name"`
	Private          bool        `json:"private"`
	Owner            Sender      `json:"owner"`
	HTMLURL          string      `json:"html_url"`
	Description      interface{} `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        int         `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
	PushedAt         int64       `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int64       `json:"size"`
	StargazersCount  int64       `json:"stargazers_count"`
	WatchersCount    int64       `json:"watchers_count"`
	Language         interface{} `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int64       `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int64       `json:"open_issues_count"`
	License          interface{} `json:"license"`
	Forks            int64       `json:"forks"`
	OpenIssues       int64       `json:"open_issues"`
	Watchers         int64       `json:"watchers"`
	DefaultBranch    string      `json:"default_branch"`
	Stargazers       int64       `json:"stargazers"`
	MasterBranch     string      `json:"master_branch"`
}

// PingRepository will define the incoming repo info
type PingRepository struct {
	ID               int64       `json:"id"`
	NodeID           string      `json:"node_id"`
	Name             string      `json:"name"`
	FullName         string      `json:"full_name"`
	Private          bool        `json:"private"`
	Owner            Sender      `json:"owner"`
	HTMLURL          string      `json:"html_url"`
	Description      interface{} `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
	PushedAt         string      `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int64       `json:"size"`
	StargazersCount  int64       `json:"stargazers_count"`
	WatchersCount    int64       `json:"watchers_count"`
	Language         interface{} `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int64       `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int64       `json:"open_issues_count"`
	License          interface{} `json:"license"`
	Forks            int64       `json:"forks"`
	OpenIssues       int64       `json:"open_issues"`
	Watchers         int64       `json:"watchers"`
	DefaultBranch    string      `json:"default_branch"`
}

// Sender should describe the person that caused the hook to be sent
type Sender struct {
	Name              *string `json:"name,omitempty"`
	Email             *string `json:"email,omitempty"`
	Login             string  `json:"login"`
	ID                int64   `json:"id"`
	NodeID            string  `json:"node_id"`
	AvatarURL         string  `json:"avatar_url"`
	GravatarID        string  `json:"gravatar_id"`
	URL               string  `json:"url"`
	HTMLURL           string  `json:"html_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReposURL          string  `json:"repos_url"`
	EventsURL         string  `json:"events_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	Type              string  `json:"type"`
	SiteAdmin         bool    `json:"site_admin"`
}

// Author is part of a commit
type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// Pusher is part of a commit
type Pusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
