# newsboat-helpers

A collection of helpers for using `newsboat`, the terminal RSS reader.

## Utilities

### `newsboat-download-podcast`

Downloads the MP3 file associated with a feed item as described by the
`Podcast Download URL` line.

Takes `--verbose`, `--overwrite` and `--root` as flags where `root` is the
directory you want to download to (defaults to `~/.newsboat/downloads`),
`--verbose` prints when a file is not downloaded because it already exists and
`--overwrite` downloads and overwrites the file if it already exists.

Run `go install ./cmd/newsboat-download-podcast` to install.

Can be added as a macro to run on the current item with the following line in
your `~/.newsboat/config` file:

```
macro d pipe-to newsboat-download-podcast
```

## Things that don't warrant their own command

### Copying the podcast URL of the current item to the clipboard

```
macro y pipe-to "awk -e '/^Podcast Download URL:/ { print $4 }' | xclip -selection clipboard"
```

Use `pbcopy` or something instead of `xclip` if you're on OS X.

