<h1 align="center">mli</h1>

<div align="center">

**📑 Manage macOS Login Items**

[![GitHub Release](https://img.shields.io/github/v/release/5ouma/mli?style=flat-square)](https://github.com/5ouma/mli/releases)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/5ouma/mli/total?style=flat-square)
[![Go Docs](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/5ouma/mli)
[![Go Report Card](https://goreportcard.com/badge/github.com/5ouma/mli?style=flat-square)](https://goreportcard.com/report/github.com/5ouma/mli)
<br />
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/5ouma/mli?style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/5ouma/mli?style=flat-square)
[![GitHub last commit](https://img.shields.io/github/last-commit/5ouma/mli?style=flat-square)](https://github.com/5ouma/mli/commit/HEAD)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/m/5ouma/mli?style=flat-square)](https://github.com/5ouma/mli/commits/main)
<br />
[![Test](https://img.shields.io/github/actions/workflow/status/5ouma/mli/test.yml?label=test&style=flat-square)](https://github.com/5ouma/mli/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/actions/workflow/status/5ouma/mli/release.yml?label=release&style=flat-square)](https://github.com/5ouma/mli/actions/workflows/release.yml)

</div>

<br /><br />

## 📥 Installation

- ### 🍺 Homebrew

  ```shell
  brew install 5ouma/formula/mli
  ```

<br />

- ### 🐹 Go

  ```shell
  go install github.com/5ouma/mli@latest
  ```

<br />

- ### 🐙 [GitHub Releases](https://github.com/5ouma/mli/releases)

  ```shell
  curl -L https://github.com/5ouma/mli/releases/latest/download/mli_Darwin_x86_64.tar.gz | tar -x mli
  ```

  ```shell
  curl -L https://github.com/5ouma/mli/releases/latest/download/mli_Darwin_arm64.tar.gz | tar -x mli
  ```

<br /><br />

## 👟 Getting Started

```shell
📑 Manage macOS Login Items with JSON

Usage:
  mli [command]

Available Commands:
  help        Help about any command
  load        Load Login Items
  save        Save Login Items

Flags:
  -h, --help      help for mli
  -v, --version   version for mli

Use "mli [command] --help" for more information about a command.
```

<br /><br />

## 🕹️ Commands

- ### 📂 `Load`

  ```shell
  📂 Load Login Items from JSON file

  Usage:
    mli load [flags]

  Flags:
        --file string   Load from this JSON file (default "./login_items.json")
    -h, --help          help for load
  ```

  <div align="center">
    <picture>
      <source
        srcset="https://raw.githubusercontent.com/5ouma/mli/HEAD/.github/assets/vhs/light/load.png"
        media="(prefers-color-scheme: light)"
      />
      <source
        srcset="https://raw.githubusercontent.com/5ouma/mli/HEAD/.github/assets/vhs/dark/load.png"
        media="(prefers-color-scheme: dark)"
      />
      <!-- markdownlint-disable MD013 -->
      <img alt="Load command GIF image generated VHS" src="https://raw.githubusercontent.com/5ouma/mli/HEAD/.github/assets/vhs/light/load.png" />
    </picture>
  </div>

<br />

- ### 💾 `Save`

  ```shell
  💾 Save Login Items to JSON file

  Usage:
    mli save [flags]

  Flags:
        --file string   Save to this JSON file (default "./login_items.json")
    -f, --force         Overwrite existing file
    -h, --help          help for save
  ```

  <div align="center">
    <picture>
      <source
        srcset="https://raw.githubusercontent.com/5ouma/mli/HEAD/.github/assets/vhs/light/save.png"
        media="(prefers-color-scheme: light)"
      />
      <source
        srcset="https://raw.githubusercontent.com/5ouma/mli/HEAD/.github/assets/vhs/dark/save.png"
        media="(prefers-color-scheme: dark)"
      />
      <!-- markdownlint-disable MD013 -->
      <img alt="Load command GIF image generated VHS" src="https://raw.githubusercontent.com/5ouma/mli/HEAD/.github/assets/vhs/light/save.png" />
    </picture>
  </div>

<br /><br />

## 🆘 Help

- [**⚠️ Issues**]: Feature Requests or Bug Reports
- [**💬 Discussions**]: General Chats or Questions
- [**🛡️ Security Policy**]: Security Issues that should not be public

[**⚠️ Issues**]: https://github.com/5ouma/mli/issues/new/choose
[**💬 Discussions**]: https://github.com/5ouma/mli/discussions/new/choose
[**🛡️ Security Policy**]: ./SECURITY.md

<br /><br />

## 💡 Inspired by

- [🚀 **blacktop / lporg**](https://github.com/blacktop/lporg)
- [⌨️ **miclf / macos_keyboard_shortcuts_exporter_importer.php**](https://gist.github.com/miclf/bf4b0cb6de9ead726197db7ed3d937b5)
