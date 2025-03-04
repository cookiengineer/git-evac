package structs

type Settings struct {
	Backup     string                      `json:"backup"`
	Folder     string                      `json:"folder"`
	Port       uint16                      `json:"port"`
	Identities map[string]IdentitySettings `json:"identities"`
	Remotes    map[string][]RemoteSettings `json:"remotes"`
}

type IdentitySettings struct {

	Name   string `json:"name"`
	SSHKey string `json:"ssh_key"`

	Git struct {
		Core struct {
			// git config --file .git/config core.sshCommand "ssh -i \"/home/cookiengineer/.ssh/identity.key\" -F /dev/null"
			SSHCommand string `json:"sshCommand"`
		} `json:"core"`
		User struct {
			// git config --file .git/config user.name  "John Doe"
			Name  string `json:"name"`
			// git config --file .git/config user.email john@example.com
			Email string `json:"email"`
		} `json:"user"`
	} `json:"git"`

}

type RemoteSettings struct {
	// "github"
	Name    string             `json:"name"`
	// map[remote-name]Remote{
	//   Name: "github",
	//   URL:  "git@github.com:{owner}/{repo}.git"
	//   URL:  "https://github.com/{owner}/{repo}.git"
	// }
	Remotes map[string]*Remote `json:"remotes"`
	Owners  []string           `json:"owners"`
}
