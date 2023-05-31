# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]

### Changed

- Bump github.com/spf13/viper from 1.15.0 to 1.16.0 ([#209](https://github.com/giantswarm/ingress-exporter/pull/209))

## [1.1.11] - 2023-05-15

### Changed

- Bump giantswarm/alpine from 3.17.3 to 3.18.0 ([#204](https://github.com/giantswarm/ingress-exporter/pull/204))

## [1.1.10] - 2023-05-08

### Changed

- Bump github.com/prometheus/client_golang from 1.15.0 to 1.15.1 ([#201](https://github.com/giantswarm/ingress-exporter/pull/201))

## [1.1.9] - 2023-04-24

### Added

- Add icon.

## [1.1.8] - 2023-04-03

### Changed

- Bump giantswarm/alpine from 3.17.2 to 3.17.3 ([#188](https://github.com/giantswarm/ingress-exporter/pull/188))

## [1.1.7] - 2023-03-06

### Changed

- Bump giantswarm/alpine from 3.17.1 to 3.17.2 ([#184](https://github.com/giantswarm/ingress-exporter/pull/184))

## [1.1.6] - 2023-01-16

### Changed

- Bump giantswarm/alpine from 3.17 to 3.17.1 ([#176](https://github.com/giantswarm/ingress-exporter/pull/176))

## [1.1.5] - 2022-11-21

### Changed

- Bump giantswarm/alpine from 3.16.3 to 3.17. ([#169](https://github.com/giantswarm/ingress-exporter/pull/169))

## [1.1.4] - 2022-11-14

## [1.1.3] - 2022-11-10

### Changed

- Update github.com/prometheus/client_golang to v1.14.0

## [1.1.2] - 2022-11-07

## [1.1.1] - 2022-09-30

### Changed

- Use giantswarm alpine container image
- Update github.com/giantswarm/microendpoint to v1.0.0
- Update github.com/giantswarm/microerror to v0.4.0
- Update github.com/giantswarm/microkit to v1.0.0
- Update github.com/giantswarm/micrologger to v0.6.0
- Update github.com/spf13/viper to v1.10.1
- Update github.com/giantswarm/exporterkit to v1.0.0

## [1.1.0] - 2021-07-21

### Changed

- Replace `jwt-go` with `golang-jwt/jwt`.
- Prepare helm values to configuration management.
- Update architect-orb to v3.0.0.

## [1.0.3] - 2021-06-14

## [1.0.2] - 2021-05-25

## [1.0.1] - 2020-11-12

# Changed

- Do not collect metrics for clusters without nginx-ingress-controller
- Add listing of apps to rbac role

## [1.0.0] - 2020-08-24

## Changed

- Updated backward incompatible Kubernetes dependencies to v1.18.5.

## [0.1.2] 2020-05-29

### Changed
- Fix RBAC template format.

## [0.1.0] 2020-05-13

### Added
- First release

[Unreleased]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.11...HEAD
[1.1.11]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.10...v1.1.11
[1.1.10]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.9...v1.1.10
[1.1.9]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.8...v1.1.9
[1.1.8]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.7...v1.1.8
[1.1.7]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.6...v1.1.7
[1.1.6]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.5...v1.1.6
[1.1.5]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.4...v1.1.5
[1.1.4]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.3...v1.1.4
[1.1.3]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.2...v1.1.3
[1.1.2]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.1...v1.1.2
[1.1.1]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/giantswarm/ingress-exporter/compare/v1.0.3...v1.1.0
[1.0.3]: https://github.com/giantswarm/ingress-exporter/compare/v1.0.2...v1.0.3
[1.0.2]: https://github.com/giantswarm/ingress-exporter/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/giantswarm/ingress-exporter/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/ingress-exporter/compare/v0.1.2...v1.0.0
[0.1.2]: https://github.com/giantswarm/ingress-exporter/compare/v0.1.0..v0.1.2
[0.1.0]: https://github.com/giantswarm/ingress-exporter/releases/tag/v0.1.0
