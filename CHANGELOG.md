# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [v1.0.1] - 2025-06-10

### Maintenance

- update CHANGELOG.md to reflect versioning format change from '1.0.0' to 'v1.0.0'
  - add test stage to CI configuration and enhance Makefile with test commands
  - remove debugging output from CI configuration for staging and production deployments
  - add CHANGELOG.md to document notable changes and versioning updates
  
  

## [v1.0.0] - 2025-06-10

### Changed

- optimize system
  - optimize system
  - optimize system
  - optimize system
  - optimize system
  - optimize system
  - optimize system
  
  ### Documentation

- update license information from Apache 2.0 to MIT License in README.md
  
  ### Maintenance

- fix escaping in CI configuration for debugging output in staging and production deployments
  - update CI configuration to directly reference .env file in staging and production deployments, and add debugging output for Docker commands
  - update CI configuration to use dynamic ENV_FILE variable and enhance Docker commands for staging and production deployments
  - update CI configuration to use pipeline ID for version tagging and refine branch rules for staging and production builds
  - update VERSION_TAG variable in CI configuration to default to 'latest' if both commit tag and branch are absent
  - refine CI configuration to enhance branch-based rules for staging and production builds, and improve script syntax for Docker commands
  - update CI configuration to use commit tags for build and deploy rules in staging and production environments
  - update CI configuration to enable build and deploy rules for release branches in staging and production environments
  - update CI scripts to execute Docker login commands over SSH for staging and production deployments
  - simplify .dockerignore by removing unnecessary patterns and update CI scripts to use .env file directly for staging and production deployments
  - update volume paths in docker-compose files for production and staging environments to include default fallback values
  - update CI configuration to improve SSH command syntax for restarting the proxy service and enhance Dockerfile by installing Swag CLI
  - enhance CI configuration by adding default retry settings, updating build rules, and improving script readability for staging and production deployments
  - refine CI configuration by updating environment variable references and adding environment context for staging and production deployments
  - update CI configuration to use .yaml file extension for docker-compose in staging and production deployments
  - update CI configuration to include deploy stages for staging and production, and remove deprecated docker-compose files
  - enhance CI configuration by adding versioned image tags for staging and production builds
  - update CI configuration for staging and production builds, streamline Docker image handling, and enhance environment variable management
  - rename production variables and update Docker image references in CI and docker-compose
  - update environment configuration and remove unused CI file
  - remove obsolete files and enable health check in Dockerfile
  - update CHANGELOG and add Makefile for changelog generation
  - update CHANGELOG with new versioning and detailed changes for 0.4.0 to 0.6.0
  
  