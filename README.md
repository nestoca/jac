# Jac

A tool for managing people and groups as Infrastructure as Code.  All people and groups are represented as YAML CRD
resources and can be queried using the `jac` CLI.

# Installation

## Installing with homebrew

```bash
$ brew tap nestoca/tap
$ brew install jac
```

Upgrade with:
```bash
$ brew update
$ brew upgrade jac
```

## Installing manually
Download from GitHub [releases](https://github.com/nestoca/jac/releases/latest) and put the binary somewhere in your
`$PATH`.

## Cloning your people git repo

Jac commands will look for a git repo at `~/.jac/repo` and fallback to using current directory as default.
You can override this with the `--dir` or `-d` flag.

```bash
$ git clone git@github.com:<repo-owner>/<people-repo>.git ~/.jac/repo
```

Jac commands will look for people and groups YAML files in that repo using the `**/*.yaml` glob pattern.
You can override this with the `--glob` or `-g` flag.

# Usage

```bash
$ jac --help

Usage:
  jac [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  groups      Get groups
  help        Help about any command
  people      Get people
  pull        Pull git repo

Flags:
  -d, --dir string    Directory to search for CRD files (defaults to ~/.jac/repo)
      --glob string   Glob expression for matching CRD files (default "**/*.yaml")
  -h, --help          help for jac
  -y, --yaml          Output in YAML format

Use "jac [command] --help" for more information about a command.
```

## List all people

```bash
$ jac people
```

## List specific people

```bash
$ jac people <person1>,<person2>,...
```

## List people belonging to any of given groups

Use `--group` or `-g` to filter by group:
```bash
$ jac people -g <group1>,<group2>,...
```

## Output results as YAML

Use `--yaml` or `-y` to output results as YAML instead of the default table format:
```bash
$ jac people -y
```

## List all groups

```bash
$ jac groups
```

## List specific groups

```bash
$ jac groups <group1>,<group2>,...
```

## List groups of specific types

Use `--type` or `-t` to filter by group type:
```bash
$ jac groups -t <type1>,<type2>,...
```

## Pull latest version of git repo

```bash
$ jac pull
```

## Pattern syntax

You can use the following syntax to specify the pattern for `groups` and `people` commands:

- Use `*` to match any number of characters
- Specify multiple `,`-separated patterns to match **any** of them
- Specify multiple `&`-separated patterns to match **all** of them
- Group patterns together with `(` `patterns...` `)`
- Prefix a pattern with `!` to negate it
- When including `*` or spaces in patterns, enclose the whole thing in quotes to avoid shell issues

# People

People are the main building blocks of Jac and can be used to model employees, contractors, consultants, etc. They
can then be organized into groups to represent your specific organizational structure.

## YAML format

```yaml
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: john-doe           # ID used to reference this person and query it
spec:
  firstName: John
  lastName: Doe
  email: john.doe@acme.com # Optional email address, for display purposes only
  groups:
    - team-sre
    - role-devops
    - role-sre
    - role-team-lead
  values:
    # Arbitrary custom key-value pairs
```

# Groups

Groups can be used to model different concepts such as departments, streams, teams, roles, etc.  It's really up to you how you want
to use them depending on your organization's needs. Groups do not have to be mutually exclusive, for example a person
can belong to multiple teams, streams, and roles.

## YAML format

```yaml
apiVersion: jac.nesto.ca/v1alpha1
kind: Group
metadata:
  name: team-devops         # ID used to reference this group and query it and the people in it
spec:
  fullName: DevOps          # Display name
  email: devops@acme.com    # Optional email address
  type: team                # Optional type (eg: stream, team, role, etc) used to filter groups
  parents:
    - stream-devops
  values:
    # Arbitrary custom key-value pairs
```

## Group parenting and inheritance

You can specify parent groups for a group via its `parents` property.

All people belonging to a group automatically inherit all of its parent groups. By default, `jac people -g <group>`
command will list all people belonging to the specified group or any of its child groups, unless the `--immediate` or
`-i` flag is specified, in which case only people belonging directly to the specified group will be returned.

That inheritance allows to reduce repetition in your YAML files and keep them DRY. For example, if you have a group
`team-sre` with parent `stream-devops`, you don't need to explicitly specify `stream-devops` as group for all people in
`team-sre`, because they will automatically inherit it.

## Group types

You can optionally specify a `type` property for a group, which can then be used to filter groups by type using the
`--type` or `-t` flag.

However, it is recommended to prefix group names with their type (eg: `stream-foo`, `team-bar`, `role-baz`) and rather
rely on wildcards for filtering them (eg: `stream-*`, `team-*`, `role-*`). The `type` property is rather intended for
programmatic processing of YAML files.

# Other practical considerations

## Leveraging custom values

You can use the `values` property of people and groups to store arbitrary key-value pairs. Those values can then be
used programmatically via Infra-as-Code and other automated workflows.

## Prefixing group names with type

If you define groups with different `type`'s, it is recommended to prefix their name with their type, such as `stream-foo`,
`team-bar`, `role-baz`, because:

- It makes it easier to filter groups by type using wildcards (eg: `stream-*`, `team-*`, `role-*`)
- It prevents name collisions between groups of different types
(eg: `team-devops` and `role-devops`).

## Organizing groups and people in directories

Leverage directories to organize your groups and people into streams, teams, etc. That structure is
purely for organizational purposes and has no impact on how Jac will treat those groups and people, as long as the
default `**/*.yaml` glob expression matches all YAML files.

For example:

```
├── roles                               // roles shared by all streams/teams
│   ├── devops.yaml                     // role-devops
│   ├── frontend.yaml                   // role-frontend
│   └── backend.yaml                    // role-backend
├── streams
│   ├── product1
│   │   ├── stream.yaml                 // stream-product1 
│   │   ├── dragons
│   │   │   ├── team.yaml               // team-dragons
│   │   │   ├── alice-wonderland.yaml
│   │   │   └── ...
│   │   ├── unicorns
│   │   │   ├── team.yaml               // team-unicorns
│   │   │   ├── jack-sparrow.yaml
│   │   │   └── ...
│   │   └── ...
│   └── product2
│       └── ...
├── devops
│   ├── stream.yaml                     // stream-devops
│   ├── sre
│   │   ├── team.yaml                   // team-sre
│   │   └── ...
│   ├── platform
│   │   ├── team.yaml                   // team-platform
│   │   └── ... 
```