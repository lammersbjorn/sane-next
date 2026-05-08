# ADR 0014: Add cross-platform compatibility checks

## Status

Accepted

## Context

`sane-next` is a Pi-first overlay with a Go companion CLI and a Node/Pi extension. ADR 0009 kept CI Linux-only while the repo was early and private. The CLI is now real enough to ship through GitHub Releases, and users may run it on Linux, macOS, or Windows.

The project should catch obvious platform regressions before release without turning acceptance into a full OS-specific install lab.

## Decision

Add cross-platform compatibility checks for Linux, macOS, and Windows.

1. Run Go CLI tests on `ubuntu-latest`, `macos-latest`, and `windows-latest`, including the install/export/doctor/repair/update/uninstall lifecycle flow that is implemented as Go tests.
2. Run Pi plugin Node tests on `ubuntu-latest`, `macos-latest`, and `windows-latest`.
3. Keep shell acceptance on Ubuntu for now because `cli/acceptance.sh` is a Bash fixture that also exercises generated artifacts and optional live Pi install behavior.
4. Cross-build release-target CLI binaries for Linux, macOS, and Windows from the Linux CI runner.
5. Avoid Unix-only shell assumptions in runtime code when a direct executable invocation is available.

This supersedes ADR 0009 only for the CI operating-system scope. ADR 0009 still governs annotated SemVer tags and minimal release automation.

## Rejected alternatives

### Keep Linux-only CI

Rejected. Linux-only CI no longer matches the stated Linux/macOS/Windows compatibility goal.

### Full acceptance on every OS immediately

Rejected for now. The current acceptance script is a Bash integration fixture and includes Pi install behavior. Porting that to every OS should be a separate slice if unit tests and cross-builds expose remaining gaps.

### Add a release framework immediately

Rejected. Cross-build verification is useful now, but full release automation can still wait until GitHub Releases need repeatable artifact publishing.

## Consequences

Positive:

- Linux, macOS, and Windows breakages are more visible before release
- Go CLI release targets are proven to compile
- runtime extension checks avoid unnecessary Unix shell dependence

Negative:

- CI will cost more than Linux-only CI
- Windows and macOS still do not run the full Bash acceptance fixture
- future OS-specific bugs may still require native shell or Pi-install acceptance coverage
