## List all people

<details>
<summary>$ jac people</summary>

```bash
$ jac people
  NAME              FIRST NAME  LAST NAME   EMAIL                 GROUPS                          INHERITED GROUPS
  alice-wonderland  Alice       Wonderland  alice@example.com     Manager Support Specialist
                                                                  Tech Support
  buzz-lightyear    Buzz        Lightyear   buzz@example.com      Frontend Developer              Development
                                                                  Incredibles Team
  dash-parr         Dash        Parr        dash@example.com      DevOps Incredibles Team         Development
  elsa              Elsa                    elsa@example.com      Backend Developer Fairies Team  Development

helen-parr Helen Parr helen@example.com Manager Incredibles Team Development
jack-sparrow Jack Sparrow jack@example.com DevOps Dragons Team Tech Support
maui Maui maui@example.com Backend Developer Moana's Team Marketing

mickey-mouse Mickey Mouse mickey@example.com Frontend Developer Tech Support
Unicorns Team
moana Moana moana@example.com Support Specialist Marketing
Moana's Team
peter-pan Peter Pan peter@example.com Backend Developer Dragons Team Tech Support

pua Pua pua@example.com Frontend Developer Marketing
Moana's Team
rapunzel Rapunzel rapunzel@example.com Manager Support Specialist
Development
tinker-bell Tinker Bell tinker@example.com DevOps Fairies Team Development
violet-parr Violet Parr violet@example.com Frontend Developer Development
Incredibles Team
woody Woody woody@example.com Manager Incredibles Team Development
———
Count: 15

```

</details>

## List specific people

<details>
<summary> $ jac people elsa,rapunzel</summary>

```bash
$ jac people elsa,rapunzel
  NAME      FIRST NAME  LAST NAME  EMAIL                 GROUPS                          INHERITED GROUPS
  elsa      Elsa                   elsa@example.com      Backend Developer Fairies Team  Development

  rapunzel  Rapunzel               rapunzel@example.com  Manager Support Specialist
                                                         Development
 ———
 Count: 2
```

</details>

## Find people with free-text search

<details>
<summary>$ jac people -f alice</summary>

Use `--find` or `-f` to find people with free-text search in their first or last name, email or name identifier:

```bash
$ jac people -f alice
  NAME              FIRST NAME  LAST NAME   EMAIL              GROUPS                          INHERITED GROUPS
  alice-wonderland  Alice       Wonderland  alice@example.com  Manager Support Specialist
                                                               Tech Support
 ———
 Count: 1
```

</details>

## List people belonging to any of given groups

<details>
<summary>$ jac people -g team-incredibles,role-frontend</summary>

Use `--group` or `-g` to filter by group:

```bash
$ jac people -g team-incredibles,role-frontend
  NAME            FIRST NAME  LAST NAME  EMAIL               GROUPS                          INHERITED GROUPS
  buzz-lightyear  Buzz        Lightyear  buzz@example.com    Frontend Developer              Development
                                                             Incredibles Team
  dash-parr       Dash        Parr       dash@example.com    DevOps Incredibles Team         Development
  helen-parr      Helen       Parr       helen@example.com   Manager Incredibles Team        Development
  mickey-mouse    Mickey      Mouse      mickey@example.com  Frontend Developer              Tech Support
                                                             Unicorns Team
  pua             Pua                    pua@example.com     Frontend Developer              Marketing
                                                             Moana's Team
  violet-parr     Violet      Parr       violet@example.com  Frontend Developer              Development
                                                             Incredibles Team
  woody           Woody                  woody@example.com   Manager Incredibles Team        Development
 ———
 Count: 7
```

</details>

## List people, hiding group columns

<details>
<summary> $ jac people -G</summary>

Use `--hide-groups` or `-G` to hide group columns (eg: if your terminal is too narrow):

```bash
$ jac people -G
  NAME              FIRST NAME  LAST NAME   EMAIL
  alice-wonderland  Alice       Wonderland  alice@example.com
  buzz-lightyear    Buzz        Lightyear   buzz@example.com
  dash-parr         Dash        Parr        dash@example.com
  elsa              Elsa                    elsa@example.com
  helen-parr        Helen       Parr        helen@example.com
  jack-sparrow      Jack        Sparrow     jack@example.com
  maui              Maui                    maui@example.com
  mickey-mouse      Mickey      Mouse       mickey@example.com
  moana             Moana                   moana@example.com
  peter-pan         Peter       Pan         peter@example.com
  pua               Pua                     pua@example.com
  rapunzel          Rapunzel                rapunzel@example.com
  tinker-bell       Tinker      Bell        tinker@example.com
  violet-parr       Violet      Parr        violet@example.com
  woody             Woody                   woody@example.com
 ———
 Count: 15
```

</details>

## List people, displaying names

<details>
<summary>$ jac people -N</summary>
Use `--show-names` or `-N` to display identifier names instead of full names:

```bash
$ jac people -N
  NAME              FIRST NAME  LAST NAME   EMAIL                 GROUPS                          INHERITED GROUPS
  alice-wonderland  Alice       Wonderland  alice@example.com     role-manager role-support
                                                                  stream-tech-support
  buzz-lightyear    Buzz        Lightyear   buzz@example.com      role-frontend team-incredibles  stream-development

  dash-parr         Dash        Parr        dash@example.com      role-devops team-incredibles    stream-development
  elsa              Elsa                    elsa@example.com      role-backend team-fairies       stream-development
  helen-parr        Helen       Parr        helen@example.com     role-manager team-incredibles   stream-development
  jack-sparrow      Jack        Sparrow     jack@example.com      role-devops team-dragons        stream-tech-support
  maui              Maui                    maui@example.com      role-backend team-moana         stream-marketing
  mickey-mouse      Mickey      Mouse       mickey@example.com    role-frontend team-unicorns     stream-tech-support
  moana             Moana                   moana@example.com     role-support team-moana         stream-marketing
  peter-pan         Peter       Pan         peter@example.com     role-backend team-dragons       stream-tech-support
  pua               Pua                     pua@example.com       role-frontend team-moana        stream-marketing
  rapunzel          Rapunzel                rapunzel@example.com  role-manager role-support
                                                                  stream-development
  tinker-bell       Tinker      Bell        tinker@example.com    role-devops team-fairies        stream-development
  violet-parr       Violet      Parr        violet@example.com    role-frontend team-incredibles  stream-development

  woody             Woody                   woody@example.com     role-manager team-incredibles   stream-development
 ———
 Count: 15
```

</details>

## Output results as YAML

<details>
<summary>$ jac people -y</summary>
Use `--yaml` or `-y` to output results as YAML instead of the default table format:

```bash
$ jac people -y
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: alice-wonderland
spec:
  firstName: Alice
  lastName: Wonderland
  email: alice@example.com
  groups:
    - role-support
    - role-manager
    - stream-tech-support
  parent: jack-sparrow
  values:
    githubUser: alicewonderland
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: buzz-lightyear
spec:
  firstName: Buzz
  lastName: Lightyear
  email: buzz@example.com
  groups:
    - role-frontend
    - team-incredibles
  parent: alice-wonderland
  values:
    githubUser: buzzlightyear
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: dash-parr
spec:
  firstName: Dash
  lastName: Parr
  email: dash@example.com
  groups:
    - role-devops
    - team-incredibles
  values:
    githubUser: dashparr
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: elsa
spec:
  firstName: Elsa
  lastName: null
  email: elsa@example.com
  groups:
    - role-backend
    - team-fairies
  values:
    githubUser: elsa
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: helen-parr
spec:
  firstName: Helen
  lastName: Parr
  email: helen@example.com
  groups:
    - role-manager
    - team-incredibles
  values:
    githubUser: helenparr
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: jack-sparrow
spec:
  firstName: Jack
  lastName: Sparrow
  email: jack@example.com
    groups:                                                                                                                                                                                                                                                                                                                                                         [55/998]
    - role-devops
    - team-dragons
  values:
    githubUser: jacksparrow
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: maui
spec:
  firstName: Maui
  lastName: null
  email: maui@example.com
  groups:
    - role-backend
    - team-moana
  values:
    githubUser: maui
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: mickey-mouse
spec:
  firstName: Mickey
  lastName: Mouse
  email: mickey@example.com
  groups:
    - role-frontend
    - team-unicorns
  values:
    githubUser: mickeymouse
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: moana
spec:
  firstName: Moana
  lastName: null
  email: moana@example.com
  groups:
    - role-support
    - team-moana
  values:
    githubUser: moana
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: peter-pan
spec:
  firstName: Peter
  lastName: Pan
  email: peter@example.com
  groups:
    - role-backend
    - team-dragons
  values:
    githubUser: peterpan
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: pua
spec:
  firstName: Pua
  lastName: null
  email: pua@example.com
  groups:
    - role-frontend
    - team-moana
  values:
    githubUser: pua
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: rapunzel
spec:
  firstName: Rapunzel
  lastName: null
  email: rapunzel@example.com
  groups:
    - role-support
    - role-manager
    - stream-development
  parent: elsa
  values:
    githubUser: rapunzel
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: tinker-bell
spec:
  firstName: Tinker
  lastName: Bell
  email: tinker@example.com
  groups:
    - role-devops
    - team-fairies
  values:
    githubUser: tinkerbelle
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: violet-parr
spec:
  firstName: Violet
  lastName: Parr
  email: violet@example.com
  groups:
    - role-frontend
    - team-incredibles
  values:
    githubUser: violetparr
---
apiVersion: jac.nesto.ca/v1alpha1
kind: Person
metadata:
  name: woody
spec:
  firstName: Woody
  lastName: null
  email: woody@example.com
  groups:
    - role-manager
    - team-incredibles
  parent: alice-wonderland
  values:
    githubUser: woody
```

</details>

## Output results as tree

<details>
<summary>$ jac people -t</summary>
Use `--tree` or `-t` to output results as YAML instead of the default table format:

```bash
$ jac people -t

├─ Dash Parr
├─ Elsa
│  └─ Rapunzel
├─ Helen Parr
├─ Jack Sparrow
│  └─ Alice Wonderland
│     ├─ Buzz Lightyear
│     └─ Woody
├─ Maui
├─ Mickey Mouse
├─ Moana
├─ Peter Pan
├─ Pua
├─ Tinker Bell
└─ Violet Parr
 ———
 Count: 15
```

</details>

## Highlight specific people in tree

<details>
<summary>$ jac people --show-all --tree --find alice</summary>
Use `--show-all` or `-A` to show all people in tree, highlighting specific people with free-text search:

```bash
# $ jac people --show-all --tree --find alice
$ jac people -Atf alice

├─ Dash Parr
├─ Elsa
│  └─ Rapunzel
├─ Helen Parr
├─ Jack Sparrow
│  └─ Alice Wonderland <-- This will be highlighted in yellow in your CLI
│     ├─ Buzz Lightyear
│     └─ Woody
├─ Maui
├─ Mickey Mouse
├─ Moana
├─ Peter Pan
├─ Pua
├─ Tinker Bell
└─ Violet Parr
 ———
 Count: 1
```

Without `--show-all`, only people matching the search will be shown, along with their parents.

</details>

## Highlight people of a specific team in tree

<details>
<summary>$ jac people --show-all --tree --group "team-unicorns"</summary>

```bash
# $ jac people --show-all --tree --group "team-unicorns"
$ jac people -Atg "team-unicorns"

├─ Dash Parr
├─ Elsa
│ └─ Rapunzel
├─ Helen Parr
├─ Jack Sparrow
│ └─ Alice Wonderland
│ ├─ Buzz Lightyear
│ └─ Woody
├─ Maui
├─ Mickey Mouse
├─ Moana
├─ Peter Pan
├─ Pua
├─ Tinker Bell
└─ Violet Parr
———
Count: 1

```

</details>

## List all groups

<details>
<summary>$ jac groups</summary>

```bash
$ jac groups
  NAME                 FULL NAME           EMAIL  TYPE    PARENT
  role-backend         Backend Developer          role
  role-devops          DevOps                     role
  role-frontend        Frontend Developer         role
  role-manager         Manager                    role
  role-support         Support Specialist         role
  stream-development   Development                stream
  stream-marketing     Marketing                  stream
  stream-tech-support  Tech Support               stream
  team-dragons         Dragons Team               team    stream-tech-support
  team-fairies         Fairies Team               team    stream-development
  team-incredibles     Incredibles Team           team    stream-development
  team-moana           Moana's Team               team    stream-marketing
  team-unicorns        Unicorns Team              team    stream-tech-support
 ———
 Count: 13
```

</details>

## List specific groups

<details>
<summary>$ jac groups role-frontend,team-unicorns</summary>

```bash
$ jac groups role-frontend,team-unicorns
  NAME           FULL NAME           EMAIL  TYPE  PARENT
  role-frontend  Frontend Developer         role
  team-unicorns  Unicorns Team              team  stream-tech-support
 ———
 Count: 2
```

</details>

## List groups of specific types

<details>
<summary>$ jac groups --type stream,role</summary>

Use `--type` to filter by group type:

```bash
# $ jac groups --type stream,role
$ jac groups -T stream,role
  NAME                 FULL NAME           EMAIL  TYPE    PARENT
  role-backend         Backend Developer          role
  role-devops          DevOps                     role
  role-frontend        Frontend Developer         role
  role-manager         Manager                    role
  role-support         Support Specialist         role
  stream-development   Development                stream
  stream-marketing     Marketing                  stream
  stream-tech-support  Tech Support               stream
 ———
 Count: 8
```

</details>

## Pull latest version of git repo

<details>
<summary>$ jac pull</summary>

```bash
$ jac pull
Already up to date.
```

</details>

---

# Pattern syntax

You can use the following syntax to specify the pattern for `groups` and `people` commands:

## Use `*` to match any number of characters

<details>
<summary>$ jac people "*oa*"</summary>

```bash
$ jac people "*oa*"
  NAME   FIRST NAME  LAST NAME  EMAIL              GROUPS                          INHERITED GROUPS
  moana  Moana                  moana@example.com  Support Specialist              Marketing
                                                   Moana's Team
 ———
 Count: 1
```

</details>

## Specify multiple `,`-separated patterns to match **any** of them

<details>
<summary>$ jac people -g role-frontend,team-unicorns</summary>

```bash
$ jac people -g role-frontend,team-unicorns
  NAME            FIRST NAME  LAST NAME  EMAIL               GROUPS                          INHERITED GROUPS
  buzz-lightyear  Buzz        Lightyear  buzz@example.com    Frontend Developer              Development
                                                             Incredibles Team
  mickey-mouse    Mickey      Mouse      mickey@example.com  Frontend Developer              Tech Support
                                                             Unicorns Team
  pua             Pua                    pua@example.com     Frontend Developer              Marketing
                                                             Moana's Team
  violet-parr     Violet      Parr       violet@example.com  Frontend Developer              Development
                                                             Incredibles Team
 ———
 Count: 4
```

</details>

## Specify multiple `&`-separated patterns to match **all** of them

<details>
<summary>$ jac people -g "role-frontend&team-unicorns"</summary>

```bash
$ jac people -g "role-frontend&team-unicorns"
  NAME          FIRST NAME  LAST NAME  EMAIL               GROUPS                          INHERITED GROUPS
  mickey-mouse  Mickey      Mouse      mickey@example.com  Frontend Developer              Tech Support
                                                           Unicorns Team
 ———
 Count: 1
```

</details>

## Prefix a pattern with `!` to negate it

<details>
<summary>$ jac people -g '!role-frontend'</summary>

```bash
$ jac people -g '!role-frontend'
  NAME              FIRST NAME  LAST NAME   EMAIL                 GROUPS                          INHERITED GROUPS
  alice-wonderland  Alice       Wonderland  alice@example.com     Manager Support Specialist
                                                                  Tech Support
  dash-parr         Dash        Parr        dash@example.com      DevOps Incredibles Team         Development
  elsa              Elsa                    elsa@example.com      Backend Developer Fairies Team  Development

  helen-parr        Helen       Parr        helen@example.com     Manager Incredibles Team        Development
  jack-sparrow      Jack        Sparrow     jack@example.com      DevOps Dragons Team             Tech Support
  maui              Maui                    maui@example.com      Backend Developer Moana's Team  Marketing

  moana             Moana                   moana@example.com     Support Specialist              Marketing
                                                                  Moana's Team
  peter-pan         Peter       Pan         peter@example.com     Backend Developer Dragons Team  Tech Support

  rapunzel          Rapunzel                rapunzel@example.com  Manager Support Specialist
                                                                  Development
  tinker-bell       Tinker      Bell        tinker@example.com    DevOps Fairies Team             Development
  woody             Woody                   woody@example.com     Manager Incredibles Team        Development
 ———
 Count: 11
```

</details>
