# Go Tour | Equivalent Binary Trees

This repository contains some different solutions to the `equivalent binary trees` final exercise in the go tour [Exercise: Equivalent Binary Trees](https://go.dev/tour/concurrency/8)

## Background

The initial solutions I came up with are [sol1.go](./sol1.go), [sol2.go](./sol2.go), [sol3.go](./sol3.go) (they do not depend on the tree being sorted).

The other solutions are from the community, with slight changes added to them.

## Realizations

After looking at other people's solutions, I realized the following points:

- Effeciency can be improved by checking for the `Left` or `Right` nodes being _null_ before recursing on them.
- The importance of the _traversal_ method being used.
