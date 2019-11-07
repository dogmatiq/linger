# GITHUB_REPO_NAME

[![Build Status](https://github.com/dogmatiq/GITHUB_REPO_NAME/workflows/CI/badge.svg)](https://github.com/dogmatiq/GITHUB_REPO_NAME/actions?workflow=CI)
[![Code Coverage](https://img.shields.io/codecov/c/github/dogmatiq/GITHUB_REPO_NAME/master.svg)](https://codecov.io/github/dogmatiq/GITHUB_REPO_NAME)
[![Latest Version](https://img.shields.io/github/tag/dogmatiq/GITHUB_REPO_NAME.svg?label=semver)](https://semver.org)
[![GoDoc](https://godoc.org/github.com/dogmatiq/GITHUB_REPO_NAME?status.svg)](https://godoc.org/github.com/dogmatiq/GITHUB_REPO_NAME)
[![Go Report Card](https://goreportcard.com/badge/github.com/dogmatiq/GITHUB_REPO_NAME)](https://goreportcard.com/report/github.com/dogmatiq/GITHUB_REPO_NAME)

This repository is a template for Dogmatiq Go modules.

[Click here](https://github.com/dogmatiq/template/generate) to create a new
repository from this template.

After creating a repository from this template, follow these steps:

- Replace the string `GITHUB_REPO_NAME` in all files with the actual repo name.
- Add a secret named `CODECOV_TOKEN` containing the codecov.io token for the new repository.
  The secret can be obtained from https://codecov.io/gh/dogmatiq/GITHUB_REPO_NAME/settings.
- Rename `.dependabot/config.yml.example` to `config.yml`. This is necessary
  because Dependabot will not inspect the commits from the template when a new
  repository is created, and hence does not detect the repository.
