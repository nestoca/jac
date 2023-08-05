# Example GitHub Action

This is an example GitHub Action that demonstrates how to generate an html page from a jac catalog, showing a table of the different streams, teams and their members.

The core logic is implemented in a simple [render-go](render-go) go application that imports and leverages the
[github.com/nestoca/jac](../../jac) go package to load the catalog in memory, structure it into a tree of
streams/teams/members, and then render it as html using a [go template](templates/teams.html).
