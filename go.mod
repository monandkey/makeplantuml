module github.com/monandkey/makeplantuml

go 1.16

replace local.packages/cmd => ./cmd

replace local.packages/tshark => ./pkg/tshark

replace local.packages/cfg => ./pkg/cfg

replace local.packages/uml => ./pkg/uml

replace local.packages/util => ./pkg/util

replace local.packages/user => ./pkg/user

require (
	github.com/spf13/cobra v1.2.1 // indirect
	local.packages/cfg v0.0.0-00010101000000-000000000000 // indirect
	local.packages/cmd v0.0.0-00010101000000-000000000000
	local.packages/tshark v0.0.0-00010101000000-000000000000 // indirect
	local.packages/uml v0.0.0-00010101000000-000000000000 // indirect
	local.packages/user v0.0.0-00010101000000-000000000000 // indirect
	local.packages/util v0.0.0-00010101000000-000000000000 // indirect
)
