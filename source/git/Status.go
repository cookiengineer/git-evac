package git

type Status string

const (

	StatusUnchanged   Status = " "
	StatusModified    Status = "M"
	StatusTypeChanged Status = "T"
	StatusAdded       Status = "A"
	StatusDeleted     Status = "D"
	StatusRenamed     Status = "R"
	StatusCopied      Status = "C"

	StatusUntracked   Status = "?"
	StatusIgnored     Status = "!"

)

func ToStatus(value string) Status {

	var result Status = StatusUntracked

	switch value {
	case " ":
		result = StatusUnchanged
	case "M":
		result = StatusModified
	case "T":
		result = StatusTypeChanged
	case "A":
		result = StatusAdded
	case "D":
		result = StatusDeleted
	case "R":
		result = StatusRenamed
	case "C":
		result = StatusCopied
	case "?":
		result = StatusUntracked
	case "!":
		result = StatusIgnored
	}

	return result

}
