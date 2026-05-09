# Dotfiler
This is my personal dotfiles manager written in go.
The end goal to write an all in one machine configuration manager
that can be reused by other people.
It should be simple to get started with and also allow advanced stuff.

Only files inside the `content` directory will be symlinked.
Directories are never symlinked.

To track a new set of files
1. Add a folder to the content directory.
2. Add mapping entry to `content/mappings.conf` accordingly

A mapping entry must exist within a group.
A group's name is wrapped within `[]`.
All mapping related to the group should follow it.

To run, do `go run ./...`

## Todos
- Rewrite the mappings file into json
- Seperate program from dotfiles
    - Use a `.dotfiler` folder inside the dotfiles directory
- Allow choosing to copy or to symlink
    - A `op` key for entries that can have value of either `copy` or `link`.
- Add a logs screen
- Show 10 entries at once at max

## Technical Improvements
- Make each screen a unique struct and add a new interface

## Known/Potential Issues
- Inability to create destination if dest path contains broken symlink dir
    - While e
- `strings.Trim` might badly interfere with files that end or start with spaces
    - Will be fixed with new mappings json implementation
- Leaves around broken symlinks if source file is removed
    - Manage all symlinks created by the program
