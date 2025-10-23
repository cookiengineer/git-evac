module git-evac-app

go 1.25.0

replace git-evac => ../source

replace github.com/cookiengineer/gooey => ../../gooey

require git-evac v0.0.0

require github.com/cookiengineer/gooey v0.0.6

require (
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/term v0.36.0 // indirect
)
