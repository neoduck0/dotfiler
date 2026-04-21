# Dotfiler
This is my personal dotfiles manager written in go.

Use the `run.sh` script to run it, it can be done by `./run.sh`.

Only files inside the `content` directory will be symlinked.
Directories are never symlinked.

To track a new set of files
1. Add a folder to the content directory.
2. Add mapping entry to mappings.conf accordingly

A mapping entry must exist within a group.
A group's name is wrapped within `[]`.
All mapping related to the group should follow it.

## Todos
- Colorize ouput
- Allow searching
- Seperate program from my dotfiles

## Known Issues
- If the destination of a mapping contains a broken symlink in its path,
recursive folder creation fails.
- Untested edge case if a filename ends or starts with empty space,
`strings.Trim` might badly interfere.
