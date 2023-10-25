# Goo

Goo is a CLI tool dedicated to scaffold a Go project of which the structure follows the community recommended standards ([here](https://github.com/golang-standards/project-layout)). Refer below for the usage guide.

## How to use

### Flavours of projects

1. There are currently two variants of projects that can be scaffolded by Goo:
   - An `sm` project dedicated for very simple use cases (eg. scripts, playground etc)
   - `lg` project dedicated for more complex use cases and is more in line with the community recommended standards.
2. The projects are generated according to the respective templates stored with the repository. Refer to the `assets` directory to find the respective templates
3. Both templates come with Git initialized

### CLI usage
1. Install this repository using the following line into your terminal

`go install github.com/brayden-ooi/goo@latest`

2. You can now use the CLI tool to scaffold projects
3. Use `goo now` along with the following flags to scaffold projects
  - `name (n)`: name of the project (required for `sm`)
  - `init (i)`: the url to be used for `go mod init` (if `init` is passed in, size will be implicitly set to `lg`, otherwise `sm`)
  - `tmp (t)`: custom path to store project templates (optional)

4. Example for an `sm` project: `goo now --name=foo`
5. Example for an `lg` project: `goo now --init=github.com/foo/foo`
   
6. This CLI tool is not supported for Windows

### Customizing the templates
1. You can pass in a custom directory into the CLI using the `tmp` flag, provided it has both `template-sm` and `template-lg` subdirectories
2. You can fork this repository and customize the `assets` directory to your liking

