module github.com/monandkey/makeplantuml

go 1.16

replace local.packages/cmd => ./cmd

require (
	github.com/spf13/cobra v1.2.1 // indirect
	local.packages/cmd v0.0.0-00010101000000-000000000000
)
