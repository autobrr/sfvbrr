<h1 align="center">sfvbrr</h1>
<p align="center">
  <img src=".github/assets/icon.png" alt="sfvbrr"><br>
  <strong>Scene. Smart. Fast.</strong><br>
  sfvbrr is a Golang tool to validate scene releases.
</p>
<div align="center">
  <p align="center">
    <img src="https://img.shields.io/badge/Go-1.25-blue?logo=go" alt="Go version">
    <img src="https://img.shields.io/badge/build-passing-brightgreen" alt="Build Status">
    <img src="https://img.shields.io/github/v/release/autobrr/sfvbrr" alt="Latest Release">
  </p>

[![Discord](https://img.shields.io/discord/881212911849209957.svg?label=&logo=discord&logoColor=ffffff&color=7389D8&labelColor=6A7EC2)](https://discord.gg/WehFCZxq5B)
</div>

## Table of Contents

- [Overview](#overview)
- [Quick Start](#quick-start)
- [Installation](#installation)
- [Usage](#usage)
- [Test](test/TESTS_RUNSHEET.md)
- [License](#license)

## Overview

[Releases](https://github.com/autobrr/sfvbrr/releases)

**sfvbrr** (pronounced _"es-ef-wee-brrrrrr"_) is a simple yet powerful tool for:

- Verifies your scene releases for consistency and cleanliness
- Validate checksums of scene release files (*.sfv)

**Key Features:**

- **Fast**: Blazingly fast checksum verification
- **Simple**: Easy to use CLI
- **Portable**: Single binary with no dependencies
- **Smart**: Detects missing/extra files based on the content and type of release
- **Customizable**: Various options in the config can change the behavior

## Quick Start

Upon the first start, the binary will create a default [configuration](internal/preset/presets.yaml) in your `$HOME/.config/sfvbrr/` folder:

<details>

```yaml
---
schema_version: 1
rules:
  app:
    deny_unexpected: true
    rules:
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: "file_id.diz"
        min: 1
        max: 1
        description: "Requires exactly one file_id.diz file"
      - pattern: "*.diz"
        max: 1
        description: "Requires no other .diz files besides file_id.diz"
      - pattern: "*.zip"
        min: 1
        description: "Requires at least one .zip file"
  audiobook:
    deny_unexpected: true
    rules:
      - pattern: "*.m3u"
        min: 1
        description: "Requires at least one .m3u file"
      - pattern: "*.sfv"
        min: 1
        description: "Requires at least one .sfv file"
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: "*.mp3"
        min: 1
        description: "Requires at least one .mp3 file"
  book:
    deny_unexpected: true
    rules:
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: "file_id.diz"
        min: 1
        max: 1
        description: "Requires exactly one file_id.diz file"
      - pattern: "*.diz"
        max: 1
        description: "Requires no other .diz files besides file_id.diz"
      - pattern: "*.zip"
        min: 1
        description: "Requires at least one .zip file"
  comic:
    deny_unexpected: true
    rules:
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: "file_id.diz"
        min: 1
        max: 1
        description: "Requires exactly one file_id.diz file"
      - pattern: "*.diz"
        max: 1
        description: "Requires no other .diz files besides file_id.diz"
      - pattern: "*.zip"
        min: 1
        description: "Requires at least one .zip file"
  education:
    deny_unexpected: true
    rules:
      - pattern: "*.rar"
        min: 1
        max: 1
        description: "Requires only one .rar file"
      - pattern: "*.sfv"
        min: 1
        max: 1
        description: "Requires only one .sfv file"
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: ".*\\.r\\d{2}$"
        regex: true
        min: 1
        description: "It usually contains one or more .r?? files"
  episode:
    deny_unexpected: true
    rules:
      - pattern: "*.rar"
        min: 1
        max: 1
        description: "Requires only one .rar file"
      - pattern: "*.sfv"
        min: 1
        max: 1
        description: "Requires only one .sfv file"
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: "Sample"
        type: dir
        min: 1
        max: 1
        description: "Requires only one Sample folder"
      # Syntax {mkv,mp4} handles the "OR" logic for extensions
      - pattern: "Sample/*.{mkv,mp4}"
        min: 1
        max: 1
        description: "Requires only one *.{mkv,mp4} file inside the Sample folder"
      - pattern: ".*\\.r\\d{2}$"
        regex: true
        min: 1
        description: "Requires at least one .r?? file"
  game:
    deny_unexpected: true
    rules:
      - pattern: "*.rar"
        min: 1
        max: 1
        description: "Requires only one .rar file"
      - pattern: "*.sfv"
        min: 1
        max: 1
        description: "Requires only one .sfv file"
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: ".*\\.r\\d{2}$"
        regex: true
        min: 1
        description: "Requires at least one .r?? file"
  magazine:
    deny_unexpected: true
    rules:
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: "file_id.diz"
        min: 1
        max: 1
        description: "Requires exactly one file_id.diz file"
      - pattern: "*.diz"
        max: 1
        description: "Requires no other .diz files besides file_id.diz"
      - pattern: "*.zip"
        min: 1
        description: "Requires at least one .zip file"
  movie:
    deny_unexpected: true
    rules:
      - pattern: "*.rar"
        min: 1
        max: 1
        description: "Requires only one .rar file"
      - pattern: "*.sfv"
        min: 1
        max: 1
        description: "Requires only one .sfv file"
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: "Sample"
        type: dir
        min: 1
        max: 1
        description: "Requires only one Sample folder"
      - pattern: "Sample/*.{mkv,mp4}"
        min: 1
        max: 1
        description: "Requires only one *.{mkv,mp4} file inside the Sample folder"
      - pattern: ".*\\.r\\d{2}$"
        regex: true
        min: 1
        description: "Requires at least one .r?? file"
  music:
    deny_unexpected: true
    rules:
      - pattern: "*.m3u"
        min: 1
        description: "Requires at least one .m3u file"
      - pattern: "*.sfv"
        min: 1
        description: "Requires at least one .sfv file"
      - pattern: "*.nfo"
        min: 1
        max: 1
        description: "Requires only one .nfo file"
      - pattern: "*.{mp3,flac}"
        min: 1
        description: "Requires at least one .mp3 or .flac file"
  series:
    deny_unexpected: true
    rules:
      - pattern: "*"
        type: dir
        min: 2
        description: "Requires at least two subfolders"
```

</details>

## Installation

* Linux

```
wget $(curl -s https://api.github.com/repos/autobrr/sfvbrr/releases/latest | grep browser_download_url | grep linux_x86_64 | cut -d\" -f4)
```

* Windows

* MacOSX

## Usage

* Basic

<details>

```bash
$ sfvbrr
sfvbrr is a high-performance scene release validation tool.

Usage:
  sfvbrr [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  sfv         Validate SFV CRC-32 checksums
  update      Update sfvbrr
  validate    Validate scene release folders
  version     Print version information

Flags:
  -h, --help   help for sfvbrr

Use "sfvbrr [command] --help" for more information about a command.
```

</details>

* Subcommand - sfv

<details>

```bash
$ sfvbrr sfv --help
Validate SFV (Simple File Verification) CRC-32 checksums for files in the specified folder(s).

The command will search for an SFV file (case insensitive) in each specified folder
and validate all files listed in the SFV file against their CRC-32 checksums.

When the recursive option (-r) is used, the command will search for SFV files in all
subdirectories of the specified folder(s).

Examples:
  # Validate a single folder
  sfvbrr sfv /path/to/release

  # Validate multiple folders
  sfvbrr sfv /path/to/release1 /path/to/release2

  # Validate recursively
  sfvbrr sfv -r /path/to/releases

Usage:
  sfvbrr sfv [folder...] [flags]

Flags:
  -b, --buffer-size int     Buffer size for file reading in bytes (0 = auto, default 64KB)
      --cpuprofile string   Write CPU profile to file
  -h, --help                help for sfv
  -q, --quiet               Quiet mode - only show errors
  -r, --recursive           Recursively search for SFV files in subdirectories
  -v, --verbose             Show detailed validation results for each file
  -w, --workers int         Number of parallel workers (0 = auto-detect)
```

</details>

* Subcommand - validate

<details>

```bash
$ sfvbrr validate --help
Validate scene release folders against category-specific rules.

The command detects the release category from the folder name and validates
the folder contents against the rules defined in the preset configuration file.

When the recursive option (-r) is used, the command will search for valid
release folders in all subdirectories of the specified folder(s).

Examples:
  # Validate a single folder
  sfvbrr validate /path/to/release

  # Validate multiple folders
  sfvbrr validate /path/to/release1 /path/to/release2

  # Validate recursively
  sfvbrr validate -r /path/to/releases

Usage:
  sfvbrr validate [folder...] [flags]

Flags:
  -h, --help            help for validate
  -p, --preset string   Path to preset YAML file (default: auto-detect)
  -q, --quiet           Quiet mode - only show errors
  -r, --recursive       Recursively search for release folders in subdirectories
  -v, --verbose         Show detailed validation results for each rule
```

</details>

## License

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later version.

See [LICENSE](LICENSE) for the full license text.
