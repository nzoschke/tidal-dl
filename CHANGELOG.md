# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html)

## [3.0.0] - 2025-02-20

### Changed

- `users` packaged renamed to `user`
- `user` first and last name types set to `*string`

## [2.1.1] - 2025-02-06

### Fixed

- Error-causing bug in InteractiveLogin

## [2.1.0] - 2025-02-03

### Added

- Track download method
- Track download to file method
- Video download method
- Video download to file method

### Changed

- Remove panics from request sender

## [2.0.1] - 2025-02-03

### Changed

- Properly bumped major version to v2

## [2.0.0] - 2025-02-03

### Added

- Album top tracks getter
- Artist albums getter
- Artist bio getter
- Artist similar artists getter
- Artist top tracks getter
- Artist videos getter
- Track stream url getter
- User getter
- User playlists getter
- Video stream url getter

### Changed
- Implemented stricter artist types instead of string slice
- Album getter now takes int id instead of string id
- Artist getter now takes int id instead of string id
- Track getter now takes int id instead of string id
- Video getter now takes int id instead of string id

## [1.0.0] - 2025-02-03

### Added

- Initial release.
- Album getter
- Artist getter
- Auth methods: Interactive, Authorization Code, Refresh Token
- Methods for creating custom interactive login
- Playlist getter
- Playlist tracks getter
- Search (by type)
- Track getter
- Video getter