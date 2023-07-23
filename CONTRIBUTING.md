# Contributing to Gorilla web toolkit
We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## We Develop with Github
We use github to host code, to track issues and feature requests, as well as accept pull requests.

## All Code Changes Happen Through Pull Requests
Pull requests are the best way to propose changes to the codebase.

1. Fork the repo and create your branch from `master`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

## Any contributions you make will be under the BSD 3-Clause License
In short, when you submit code changes, your submissions are understood to be under the same [BSD 3-Clause License](https://opensource.org/license/bsd-3-clause/) that covers the project. Feel free to contact the maintainers if that's a concern.

## Report bugs using Github issues
We use GitHub issues to track public bugs. Report a bug by [opening a new issue](); it's that easy!

## Write bug reports with detail, background, and sample code

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can.
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

People *love* thorough bug reports. I'm not even kidding.

## Use a Consistent Coding Style
Each repository contains a `.editorconfig` file, visit [EditorConfig](https://editorconfig.org/) to learn more about using this file with your favorite code editor.

## Run `make verify` and `make test` before pushing code
Using a terminal, run the `make verify` command that is available within each repository. This command runs tools such as `golangci-lint`, `gosec` and `govulncheck`, which are also run as part of our CI system. If these checks fail on your workstation they will most likely fail during CI testings. Also be sure to run `make test` which will run the repositories `unit tests`. Again, if these do not pass on your workstation they will most likely not pass in CI.

## License
By contributing, you agree that your contributions will be licensed under its BSD 3-Clause License.
