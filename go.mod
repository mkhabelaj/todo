module github.com/mkhabelaj/todo

go 1.23.5

replace github.com/mkhabelaj/todo/internal/todo => ../todo/internal/todo

replace github.com/mkhabelaj/todo/internal/connectors/json => ../todo/internal/connectors/json

require (
	github.com/aquasecurity/table v1.8.0
	github.com/liamg/tml v0.7.0
	github.com/spf13/cobra v1.8.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/term v0.0.0-20220526004731-065cf7ba2467 // indirect
)

replace github.com/mkhabelaj/todo/internal/util => ../todo/internal/util
