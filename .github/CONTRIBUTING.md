# Contributing

Thank you for taking the time to contribute to the most EPIC terraform provider in existence.

This provider is a utility provider designed to allow users to add a little bit of EPICNESS to their resources. We all know pet names are a no-no in the production world, but who can resist naming their EC2 instances after their favorite Star Wars characters?

### Adding New Media Types and Titles

The `/data` directory is tightly controlled and validated to prevent the provider from breaking. All subdirectorys within `/data` are considered a valid `media_type` for this providers resources. If you wish to contribute additional media files you **DO NOT** have to edit any code. Simply contribute a title.json file. See below:

- Media type directries **MUST** only contain underscores. Spaces and hyphens are not supported.

Any valid `/data/media_type/*.json` file is considered a valid `title` for this providers resources.

- Title file names **MUST** only contains underscores and end in .json.
- All title files must conform to the schema similiar to the other title.json files in `/data/`. See the other files for an example.

`.github/validate_data.yml` will validate that the file structure is properly configured and files are named appropriatly.

## Asking Questions

For questions, curiosity, or if still unsure what you are dealing with, open an issue with the label "Question".

## Raising Issues

We welcome issues of all kinds including feature requests, bug reports or documentation suggestions.
Below are guidelines for well-formed issues of each type.

### Bug Reports

* [ ] **Test against latest release**: Make sure you test against the latest available version of Terraform and the provider.
  It is possible we may have already fixed the bug you're experiencing.
* [ ] **Search for duplicates**: It's helpful to keep bug reports consolidated to one thread, so do a quick search
  on existing bug reports to check if anybody else has reported the same thing.
  You can scope searches by the label `bug` to help narrow things down.
* [ ] **Include steps to reproduce**: Provide steps to reproduce the issue, along with code examples and/or real code,
  so we can try to reproduce it. Without this, it makes it much harder (sometimes impossible) to fix the issue.

### Feature Requests

* [ ] **Search for possible duplicate requests**: It's helpful to keep requests consolidated to one thread,
  so do a quick search on existing requests to check if anybody else has reported the same thing.
  You can scope searches by the label `enhancement` to help narrow things down.
* [ ] **Include a use case description**: In addition to describing the behavior of the feature you'd like to see added,
  it's helpful to also make a case for why the feature would be important and how it would benefit
  the provider and, potentially, the wider Terraform ecosystem.

## New Pull Request

Thank you for contributing!

We are happy to review pull requests without associated issues,
but we **highly recommend** starting by describing and discussing
your problem or feature and attaching use cases to an issue first
before raising a pull request.

* [ ] **Early validation of idea and implementation plan**: provider development is complicated enough that there
  are often several ways to implement something, each of which has different implications and tradeoffs.
  Working through a plan of attack with the team before you dive into implementation will help ensure that you're
  working in the right direction.
* [ ] **Tests**: It may go without saying, but every new patch should be covered by tests wherever possible.
  For bug-fixes, tests to prove the fix is valid. For features, tests to exercise the new code paths.
* [ ] **Go Modules**: We use [Go Modules](https://github.com/golang/go/wiki/Modules) to manage and version our dependencies.
  Please make sure that you reflect dependency changes in your pull requests appropriately
  (e.g. `go get`, `go mod tidy` or other commands).
  Refer to the [dependency updates](#dependency-updates) section for more information about how
  this project maintains existing dependencies.
* [ ] **Changelog**: Refer to the [changelog](#changelog) section for more information about how to create changelog entries.


### Dependency Updates

Dependency management is performed by [Dependabot](https://docs.github.com/en/code-security/dependabot/dependabot-version-updates).
Where possible, dependency updates should occur through that system to ensure all Go module files are appropriately
updated and to prevent duplicated effort of concurrent update submissions.
Once available, updates are expected to be verified and merged to prevent latent technical debt.

### Changelog

This project has a user-friendly, readable `CHANGELOG`s that allow
practitioners and developers to tell at a glance whether a release should have any effect on them,
and to gauge the risk of an upgrade.
