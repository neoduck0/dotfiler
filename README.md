# Dotfiler
This is my personal dotfiles manager written in go.

Only files inside the `content` directory will be symlinked.
Directories are never symlinked.

To track a new set of files
1. Add a folder to the content directory.
2. Add mapping entry to `content/mappings.conf` accordingly

A mapping entry must exist within a group.
A group's name is wrapped within `[]`.
All mapping related to the group should follow it.

To run, do `go build -o dotfiler ./src; ./dotfiler`

## Todos
- Show 10 entries at once at max
- Seperate program from my dotfiles
- Allow choosing to copy or to symlink
- Port to bubbles as much as possible

## Known Issues
- If the destination of a mapping contains a broken symlink in its path,
recursive folder creation fails.
- Untested edge case if a filename ends or starts with empty space,
`strings.Trim` might badly interfere.
