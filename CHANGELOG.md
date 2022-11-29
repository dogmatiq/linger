# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html

## [1.0.0] - 2022-11-29

### Changed

- `backoff.Retry()` will now use the default strategy if a `nil` strategy is passed

## [0.2.2] - 2022-11-29

## Fixed

- Handle full range of durations in `Rand()`, `FullJitter()` and `ProportionalJitter()`

## [0.2.1] - 2020-03-17

## Fixed

- Prevent `backoff.Exponential()` and `Linear()` from overflowing `time.Duration`

## [0.2.0] - 2020-03-06

## Added

- Add `backoff.Retry()`

### Changed

- **[BC]** Moved `BackoffStrategy` to `backoff.Strategy`, note that the signature has also changed
- **[BC]** Moved `Backoff` to `backoff.Counter`, note that all options are now specified by the strategy

## [0.1.1] - 2020-02-24

### Added

- Add `Backoff` and `BackoffStrategy`

## [0.1.0] - 2019-11-11

- Initial release

<!-- references -->

[unreleased]: https://github.com/dogmatiq/linger
[0.1.0]: https://github.com/dogmatiq/linger/releases/tag/v0.1.0
[0.1.1]: https://github.com/dogmatiq/linger/releases/tag/v0.1.1
[0.2.0]: https://github.com/dogmatiq/linger/releases/tag/v0.2.0
[0.2.1]: https://github.com/dogmatiq/linger/releases/tag/v0.2.1
[0.2.2]: https://github.com/dogmatiq/linger/releases/tag/v0.2.2
[1.0.0]: https://github.com/dogmatiq/linger/releases/tag/v1.0.0

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
