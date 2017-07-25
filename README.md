# kubeci
Command line tool which performs Kubernetes operations commonly required during continuous integration.

# Usage
Kubeci is a binary which provides a command line interface. Subcommands are describe in the following section.

## Run
Kubeci provides numerous operations. These operations run common Kubernetes actions.  

You can run these kubeci operations with the `run` command.  

Usage: `$ kubeci run <operation name> --os <operating system>`

Options:

- `operation name` (string): Name of operation to run
    - **required**
- `operating system` (string): Name of operating system kubeci is running commands on
    - **required**
    - Used to run different variations of certain commands (ex., package management commands)
    - Allowed values:
        - `alpine`: Alpine Linux

## List
List all available kubeci operations.

Usage: `$ kubeci list`

## Help
Provide general kubeci help or specific operation help.  

Usage: `$ kubeci help [operation name]`

Options:

- `operation name` (string): Name of operation to provide help about
    - *optional*
    - If no operation name is provided general kubeci help will be displayed