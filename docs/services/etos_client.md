.. _etos-client:

# ETOS Client

Main client for ETOS.

Github: [ETOS Client](https://github.com/eiffel-community/etos)

This page describes the ETOS client and its options as well as some examples of how it can be used. ETOS Client is a command-line tool that can be used in Continuous Integration tools such as jenkins as well as locally on workstations.

## Why ETOS Client?

The ETOS client is a tool built to make it easier to execute ETOS (Eiffel Test Orchestration System). It is used for starting test suites in ETOS, and reporting the outcome.

## Eager to get started

A getting started page can be found here: [Getting Started](../getting-started.md)

## CLI Overview and Command Reference

### Prerequisites

- Python 3.6 or higher.
- Git (optional)

### Installation

There are two methods for installing the package.

#### Install using pip (Recommended)

Recommended for most users. It installs the latest stable version released.
ETOS client can be found on PyPi. If you have the pip package  manager installed, then the simplest way of installing ETOS Client is using the following command:

```sh
pip install etos_client
```

Instructions that virtually take you by the hand and guide you every step of the way is available among our [Step by step](../step_by_step.md) articles.

#### Install using Git

Recommended for developers who want to install the package along with the full source code. Clone the package repository, and install the package in your site-package directory:

```sh
git clone "https://github.com/eiffel-community/etos-client.git" client
cd client
pip install -e .
```

### General Syntax

The usage example below describes the interface of etos_client, which can be invoked with different combinations of options. The example uses brackets "[ ]" to describe optional elements. Together, these elements form valid usage patterns, each starting with the program's name etos_client.

Below the usage patterns, there is a table of the command-line options with descriptions. They describe whether an option has short/long forms (-h, --help), whether an option has an argument (--identity=<IDENTITY>), and whether that argument has a default value.

```sh
etos_client [-h] -i IDENTITY -s TEST_SUITE [--no-tty] [-w WORKSPACE] [-a ARTIFACT_DIR] [-r REPORT_DIR] [-d DOWNLOAD_REPORTS] [--iut-provider IUT_PROVIDER] [--execution-space-provider EXECUTION_SPACE_PROVIDER] [--log-area-provider LOG_AREA_PROVIDER] [--dataset DATASET] [--cluster CLUSTER] [--version] [-v]
```

#### Command-line options

```sh
etos_client --help
```

#### Examples

TBD

## Known issues

All issues can be found here: https://github.com/eiffel-community/etos/issues

## Stuck and in need of help

There is a mailing list at: etos-maintainers@google.groups.com or just write an Issue.
