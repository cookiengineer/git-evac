package git

import "os"
import "strings"

type GlobalConfig struct {
	file string
	User struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		// TODO: GPG Signing key
	} `json:"user"`
	Aliases   map[string]string         `json:"aliases"`
	Difftools map[string]ConfigDifftool `json:"difftools"`
	Pager     map[string]bool           `json:"pager"`
	Init struct {
		DefaultBranch string `json:"defaultBranch"`
	} `json:"init"`
}

func InitGlobalConfig(file string) *GlobalConfig {

	var config GlobalConfig

	config.Aliases = make(map[string]string)
	config.Difftools = make(map[string]ConfigDifftool)
	config.Pager = make(map[string]bool)

	stat, err := os.Stat(file)

	if err == nil && !stat.IsDir() {

		config.file = file
		config.Parse()

		return &config

	}

	return nil

}

func (config *GlobalConfig) Parse() {

	if config.file != "" {

		buffer, err1 := os.ReadFile(config.file)

		if err1 == nil {

			section_type := ""
			section_name := ""

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {

					tmp := line[1:len(line)-1]

					if strings.Contains(tmp, " ") {

						tmp_type := strings.TrimSpace(tmp[0:strings.Index(tmp, " ")])
						tmp_name := strings.TrimSpace(tmp[strings.Index(tmp, " ")+1:])

						if strings.HasPrefix(tmp_name, "\"") && strings.HasSuffix(tmp_name, "\"") {
							tmp_name = strings.TrimSpace(tmp_name[1:len(tmp_name)-1])
						}

						if tmp_type != "" && tmp_name != "" {
							section_type = strings.TrimSpace(tmp_type)
							section_name = strings.TrimSpace(tmp_name)
						}

					} else {
						section_type = tmp
						section_name = ""
					}

				} else if strings.Contains(line, " = ") {

					key := strings.TrimSpace(line[0:strings.Index(line, " = ")])
					val := strings.TrimSpace(line[strings.Index(line, " = ")+3:])

					parseGlobalConfigValue(config, section_type, section_name, key, val)

				}

			}

		}

	}

	// TODO: Parse file buffer

}

