# Versioning rules

ETOS uses semantic versioning in the form of \"MAJOR.MINOR.PATCH\"

-   MAJOR: A large change, sometimes breaking changes.
-   MINOR: A smaller change, never a breaking change.
-   PATCH: A bugfix required in production.

For production environment we recommend using the installation YAML file from the latest release and update it manually in a controlled manner.

For staging/development environments you can follow the main branch of the [main ETOS repository](https://github.com/eiffel-community/etos).
However be aware that we will break things outside of releases, both intentionally and unintentionally, without warning.

With ETOS v1 we added a way to deploy multiple [ETOS clusters](./api.md#cluster) and best-practices is to deploy a new cluster when a new ETOS version has been released so that one does not break users workflows.
