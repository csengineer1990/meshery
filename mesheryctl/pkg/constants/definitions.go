package constants

const (
	// Meshery Repository Location
	mesheryGitHubOrg  string = "meshery"
	mesheryGitHubRepo string = "meshery"
	CallbackURLENV    string = "MESHERY_SERVER_CALLBACK_URL"
	ProviderENV       string = "PROVIDER"
)

// GetMesheryGitHubOrg retrieves the name of the GitHub organization under which the Meshery repository resides.
func GetMesheryGitHubOrg() string {
	return mesheryGitHubOrg
}

// GetMesheryGitHubRepo retrieves the name of the Meshery repository
func GetMesheryGitHubRepo() string {
	return mesheryGitHubRepo
}
