module github.com/autobrr/sfvbrr

go 1.25.1

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/creativeprojects/go-selfupdate v1.5.2
	github.com/dustin/go-humanize v1.0.1
	github.com/fatih/color v1.18.0
	github.com/moistari/rls v0.6.0
	github.com/schollz/progressbar/v3 v3.19.0
	github.com/spf13/cobra v1.10.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	code.gitea.io/sdk/gitea v0.22.1 // indirect
	github.com/42wim/httpsig v1.2.3 // indirect
	github.com/Masterminds/semver/v3 v3.4.0 // indirect
	github.com/davidmz/go-pageant v1.0.2 // indirect
	github.com/go-fed/httpsig v1.1.0 // indirect
	github.com/google/go-github/v74 v74.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.8 // indirect
	github.com/hashicorp/go-version v1.8.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	github.com/ulikunitz/xz v0.5.15 // indirect
	gitlab.com/gitlab-org/api/client-go v1.9.1 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/oauth2 v0.34.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/term v0.38.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	golang.org/x/time v0.14.0 // indirect
)

// Fix for arm6 build: force use of xz v0.5.15+ to avoid integer overflow on 32-bit architectures
replace github.com/ulikunitz/xz => github.com/ulikunitz/xz v0.5.15

// Use autobrr fork of rls library
replace github.com/moistari/rls => github.com/autobrr/rls v0.7.0
