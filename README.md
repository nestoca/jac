# Jac

Jac is a GitOps CLI tool and YAML file format for managing and querying your people metadata Source of Truth™.

- Defines two CRD-like YAML file formats:
  - [Person](https://github.com/nestoca/jac#people) for modeling your people (eg: managers, team members...).
  - [Group](https://github.com/nestoca/jac#groups) for modeling your various sets of people (eg: departments, streams, teams...) and their roles.
- GitOps-oriented: Your jac "catalog" git repo defines your people metadata, from which you can automate different processes, such as granting access permissions and generating an org chart or a [teams page](https://nestoca.github.io/jac/).
- The `jac` CLI allows you to query your catalog in various ways, allowing to quickly answer questions such as:
  - Who's the front-end developer in the Dragons team?
  - How many back-end developers do we have in the organization?
  - Who's the manager of the Unicorns team?
  - What are the email addresses of all SREs?
- YAML resources are extensible via custom values to address your specific metadata needs.

# Installation

## Installing with homebrew

```bash
$ brew tap nestoca/public
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

## Cloning your catalog git repo

Jac commands will look for a git repo at `~/.jac` and fallback to using current directory as default.
You can override this with the `--dir` or `-d` flag.

```bash
$ git clone git@github.com:<repo-owner>/<catalog-repo>.git ~/.jac
```

## Cloning to a different directory

Put a `.jacrc` file in your home directory and set the `dir` property to the path to your git repo:

```yaml
dir: /path/to/repo
```

# Examples

To get started, clone this repo and have a look at the [examples](examples) directory, where you will find an example [catalog](examples/catalog), queries and a GitHub [workflow](.github/workflows/publish-example.yaml) and [action](examples/render-action) for automatically generating a [teams pages](https://nestoca.github.io/jac/).

# People

People are the main building blocks of Jac and can be used to model employees, contractors, consultants, etc. They
can then be organized into groups to represent your specific organizational structure.

## YAML format

```yaml
apiVersion: jac.nesto.ca/v1alpha1
kind: Person

metadata:
  # Unique ID
  name: alice-wonderland

spec:
  # Display info
  firstName: Alice
  lastName: Wonderland
  email: alice@example.com

  # Groups this person belongs to
  groups:
    - role-support
    - role-manager
    - stream-tech-support

  # Optional parent person for tree-display / org-chart (eg: manager, team lead...)
  parent: jack-sparrow

  # Arbitrary custom key-value pairs
  values:
    githubUser: alicewonderland
```

# Groups

Groups can be used to model different concepts such as departments, streams, teams, roles, etc. It's really up to you how you want
to use them depending on your organization's needs. Groups do not have to be mutually exclusive, for example a person
can belong to multiple teams, streams, and roles.

## YAML format

```yaml
apiVersion: jac.nesto.ca/v1alpha1
kind: Group

metadata:
  # Unique ID
  name: team-devops

spec:
  # Display info
  emoji: 🛠️
  fullName: DevOps
  email: devops@example.com

  # Optional type (eg: stream, team, role, etc) used to filter groups
  type: team

  # Optional parent group for tree-display
  parent: stream-devops

  # Arbitrary custom key-value pairs
  values:
    resourceLabel: devops
```

## Group parenting and inheritance

You can specify parent groups for a group via its `parent` property.

All people belonging to a group automatically inherit all of its parent groups. By default, `jac people -g <group>`
command will list all people belonging to the specified group or any of its child groups, unless the `--immediate` or
`-i` flag is specified, in which case only people belonging directly to the specified group will be returned.

That inheritance allows to reduce repetition in your YAML files and keep them DRY. For example, if you have a group
`team-sre` with parent `stream-devops`, you don't need to explicitly specify `stream-devops` as group for all people in
`team-sre`, because they will automatically inherit it.

## Group types

You can optionally specify a `type` property for a group, which can then be used to filter groups by type using the
`--type` or `-T` flag.

However, it is recommended to prefix group names with their type (eg: `stream-foo`, `team-bar`, `role-baz`) and rather
rely on wildcards for filtering them (eg: `stream-*`, `team-*`, `role-*`). The `type` property is rather intended for
programmatic processing of YAML files.

# Usage

```bash
$ jac --help

Usage:
  jac [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  groups      List groups
  help        Help about any command
  people      List people
  pull        Pull git repo
  version     Display jac version

Flags:
  -d, --dir string    Directory to search for CRD files (defaults to ~/.jac/repo)
  -h, --help          help for jac
  -y, --yaml          Output in YAML format

Use "jac [command] --help" for more information about a command.
```

Detailed examples for the commands that follow have been documented in a [separate README in the `example` folder](example/README.md).

## List all people

```bash
$ jac people
```

## List specific people

```bash
$ jac people <person1>,<person2>,...
```

## Find people with free-text search

Use `--find` or `-f` to find people with free-text search in their first or last name, email or name identifier:

```bash
$ jac people --find alice
$ jac people -f alice
```

## List people belonging to any of given groups

Use `--group` or `-g` to filter by group:

```bash
$ jac people --group <group1>,<group2>,...
$ jac people -g <group1>,<group2>,...
```

## List people, hiding group columns

Use `--hide-groups` or `-G` to hide group columns (eg: if your terminal is too narrow):

```bash
$ jac people --hide-groups
$ jac people -G
```

## List people, displaying names

Use `--show-names` or `-N` to display identifier names instead of full names:

```bash
$ jac people --show-names
$ jac people -N
```

## Output results as YAML

Use `--yaml` or `-y` to output results as YAML instead of the default table format:

```bash
$ jac people --yaml
$ jac people -y
```

## Output results as tree

Use `--tree` or `-t` to output results as YAML instead of the default table format:

```bash
$ jac people --tree
$ jac people -t
```

## Highlight specific people in tree

Use `--show-all` or `-A` to show all people in tree, highlighting specific people with free-text search:

```bash
$ jac people --show-all --tree --find <search>
$ jac people -Atf <search>
```

Without `--show-all`, only people matching the search will be shown, along with their parents.

## Highlight people of a specific team in tree

```bash
$ jac people --show-all --tree --group "team-sre"
$ jac people -Atg "team-sre"
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

Use `--type` to filter by group type:

```bash
$ jac groups --type <type1>,<type2>,...
$ jac groups -T <type1>,<type2>,...
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
- Force operator evaluation order using parentheses (eg: `(pattern1&pattern2),pattern3`)
- Prefix a pattern with `!` to negate it
- When including `*`, `&`, `()`, `!` or spaces in patterns, enclose the whole thing in single-quotes to avoid shell issues

# Other practical considerations

## Leveraging custom values

You can use the `values` property of people and groups to store arbitrary key-value pairs. Those values can then be used programmatically in automated workflows. In the example [catalog](examples/catalog), streams have `resourceLabel` and `description` values (used while generating the html page) and people have a `githubUser` value.

## Prefixing group names with type

If you define groups with different `type`'s, it is recommended to prefix their name with their type, such as `stream-foo`,
`team-bar`, `role-baz`, because:

- It makes it easier to filter groups by type using wildcards (eg: `stream-*`, `team-*`, `role-*`)
- It prevents name collisions between groups of different types
  (eg: `stream-devops`, `team-devops` and `role-devops`).

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

# How jac resolves directory and glob pattern

When you run `jac` it looks for `.jacrc` file in the following locations:

1. Directory specified explicitly via `--dir` (or `-d`) flag
2. Current directory
3. Your `$HOME` directory
4. Your `$HOME/.jac` directory

That file is in YAML format and can contain the following properties:

```yaml
dir: /full/path/to/directory
glob: "**/*.yaml"
```

The `dir` property is optional, can be an absolute path, or be relative to the current
config file's directory. If specified, Jac will use that directory combined with the `glob`
expression to find and load its YAML files. If not specified, Jac will use the current
config file's directory instead. If Jac finds another config file in that directory, it will
follow the same process over and over until no further config files and directories are found.

The `glob` is optional, defaults to `**/*.yaml`

# Why the name Jac?

- It's a short 3-letter command that's easy to type and remember.
- It sounds like the name of a person, which gives it some personality.
- It falls in line nicely with other tools by same authors, such as [joy](https://github.com/nestoca/joy), [jen](https://github.com/silphid/jen), and [yey](https://github.com/silphid/yey).
- Could mean something like "Johns and Janes As Code".
