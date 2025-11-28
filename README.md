<img src="images/etos-logo.png" alt="ETOS" width="350px">

[![image](https://img.shields.io/badge/Stage-Sandbox-yellow.svg)](https://github.com/eiffel-community/community/blob/master/PROJECT_LIFECYCLE.md#stage-sandbox)

# ETOS

ETOS (Eiffel Test Orchestration System) is a new test orchestration system which takes away control of how to run and what to run from the system itself and places it into the hands of the relevant engineers.

The idea of having a system dictate what and how to run is finished.
Let's bring back control to the testers.

With ETOS we define how to run tests using recipes and we define what to run with collections of recipes.

| [Test Engineer]()      | [Test Automation Engineer]() | [System Engineer]() |
| ------------------ | ------------------------ | --------------- |
| Create collections | Create recipes           | Deploy ETOS     |
| Analyze results    | Create tests             | Infrastructure  |
| What to run?       | How to run?              | Where to run?   |


This means that only people who knows how and what to run decide these factors.
ETOS will only receive the collection of recipes and execute it accordingly. You can also mix test suites. For instance, let's say you want to run a `go` unittest and a `python` function test in the same suite, that's easy to do; just add them to your collection.

This is the strength of ETOS. Relying on the humans to decide what to run, how to run and where to run.

# Table of contents
1. [Features](#features)
1. [Installation](#installation)
1. [Development](./docs/development.md)
1. [Versioning](./docs/versioning.md)
1. [Contribute](#contribute)
1. [Support](#support)

<a name="features"></a>
## Features

-   Generic test suite execution based solely on JSON.
-   Mix and match test suites, regardless of programming language.
-   Separation of concerns between testers, test automation engineers
    and system engineers.
-   Eiffel protocol implementation.

<a name="installation"></a>
## Installation

### Requirements

In order to install ETOS, you need to meet the following requirements.

-   An up and running kubernetes cluster (<https://kubernetes.io/>)

### Installation Steps

### Deployment Configuration

<a name="contribute"></a>
## Contribute

| Please write issues in the relevant repositories for where you found the issue.
| If you do not know which repository to write the issue for, feel free to write it here and it will be moved.
| Documentation issues are reported here.

-   Issue Tracker: <https://github.com/eiffel-community/etos/issues>
-   Source Code: <https://github.com/eiffel-community/etos>

<a name="support"></a>
## Support

If you are having issues, please let us know. There is a mailing list at: <etos-maintainers@googlegroups.com> or just write an Issue.
