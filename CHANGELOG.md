# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/ingress-exporter/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/giantswarm/ingress-exporter/compare/v1.0.3...v1.1.0
[1.0.3]: https://github.com/giantswarm/ingress-exporter/compare/v1.0.2...v1.0.3
[1.0.2]: https://github.com/giantswarm/ingress-exporter/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/giantswarm/ingress-exporter/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/ingress-exporter/compare/v0.1.2...v1.0.0
[0.1.2]: https://github.com/giantswarm/ingress-exporter/compare/v0.1.0..v0.1.2
[0.1.0]: https://github.com/giantswarm/ingress-exporter/releases/tag/v0.1.0
