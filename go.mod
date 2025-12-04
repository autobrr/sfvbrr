module github.com/autobrr/sfvbrr

go 1.25.1

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/creativeprojects/go-selfupdate v1.5.1
	github.com/dustin/go-humanize v1.0.1
	github.com/fatih/color v1.18.0
	github.com/moistari/rls v0.6.0
	github.com/schollz/progressbar/v3 v3.18.0
	github.com/spf13/cobra v1.10.2
)

require (
	code.gitea.io/sdk/gitea v0.22.0 // indirect
	github.com/42wim/httpsig v1.2.3 // indirect
	github.com/Masterminds/semver/v3 v3.4.0 // indirect
	github.com/davidmz/go-pageant v1.0.2 // indirect
	github.com/go-fed/httpsig v1.1.0 // indirect
	github.com/google/go-github/v30 v30.1.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.8 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	github.com/ulikunitz/xz v0.5.15 // indirect
	github.com/xanzy/go-gitlab v0.115.0 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/oauth2 v0.30.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/term v0.34.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	golang.org/x/time v0.12.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Fix for arm6 build: force use of xz v0.5.15+ to avoid integer overflow on 32-bit architectures
replace github.com/ulikunitz/xz => github.com/ulikunitz/xz v0.5.15
