module github.com/monandkey/makeplantuml

go 1.16

replace (
	local.packages/cfg => ./pkg/cfg
	local.packages/cmd => ./cmd
	local.packages/gtpv2 => ./pkg/protocol/gtpv2
	local.packages/ngap => ./pkg/protocol/ngap
	local.packages/s1ap => ./pkg/protocol/s1ap
	local.packages/tshark => ./pkg/tshark
	local.packages/uml => ./pkg/uml
	local.packages/user => ./pkg/user
	local.packages/util => ./pkg/util
)

require (
	github.com/spf13/cobra v1.2.1 // indirect
	local.packages/cfg v0.0.0-00010101000000-000000000000 // indirect
	local.packages/cmd v0.0.0-00010101000000-000000000000
	local.packages/gtpv2 v0.0.0-00010101000000-000000000000 // indirect
	local.packages/ngap v0.0.0-00010101000000-000000000000 // indirect
	local.packages/s1ap v0.0.0-00010101000000-000000000000 // indirect
	local.packages/tshark v0.0.0-00010101000000-000000000000 // indirect
	local.packages/uml v0.0.0-00010101000000-000000000000 // indirect
	local.packages/user v0.0.0-00010101000000-000000000000 // indirect
	local.packages/util v0.0.0-00010101000000-000000000000 // indirect
)
