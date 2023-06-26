module github.com/dum-dum-genius/ggol

go 1.18

retract (
	// Retract v1.0.4 and v1.0.3
	v1.0.4
	// Wrong tag by mistake
	v1.0.3
	// Too early for released version 1
	v1.0.2
	// Too early for released version 1
	v1.0.1
	// Too early for released version 1
	v1.0.0
	// Incorrect build
	v0.1.0
)

require github.com/anthonyLock/testify v0.0.0-20230206122917-1fdbdc28f21b

require (
	github.com/josharian/impl v1.2.0 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.8.2
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
