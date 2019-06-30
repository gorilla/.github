🦍🦍🦍

This repo contains the GitHub issue &amp; PR templates for the
[@gorilla](https://github.com/gorilla) project, as well as a handful of
maintainer-related tools.

## One-Liners

Update all labels using [`github-label-sync`](https://github.com/Financial-Times/github-label-sync) and the `labels.json` file here:

```sh
➜  for repo in $(curl -s https://api.github.com/orgs/gorilla/repos | jq -r '.[] | select(.name | contains("github") | not) | .name'); do \
    github-label-sync --access-token $GITHUB_TOKEN --dry-run gorilla/$repo; done
```

Build the CI templates:

```sh
➜  go run generate_ci.go
```
