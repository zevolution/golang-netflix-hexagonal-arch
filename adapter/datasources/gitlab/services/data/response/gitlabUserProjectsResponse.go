package response

import (
	"time"
)

type GitlabUserProjectsResponse struct {
	ID                int       `json:"id"`
	Description       string    `json:"description"`
	Name              string    `json:"name"`
	NameWithNamespace string    `json:"name_with_namespace"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	CreatedAt         time.Time `json:"created_at"`
	DefaultBranch     string    `json:"default_branch"`
	TagList           []any     `json:"tag_list"`
	Topics            []any     `json:"topics"`
	SSHURLToRepo      string    `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string    `json:"http_url_to_repo"`
	WebURL            string    `json:"web_url"`
	ReadmeURL         string    `json:"readme_url"`
	ForksCount        int       `json:"forks_count"`
	AvatarURL         any       `json:"avatar_url"`
	StarCount         int       `json:"star_count"`
	LastActivityAt    time.Time `json:"last_activity_at"`
	Namespace         struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Path      string `json:"path"`
		Kind      string `json:"kind"`
		FullPath  string `json:"full_path"`
		ParentID  any    `json:"parent_id"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"namespace"`
}