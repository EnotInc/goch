# goch - Go Chess

Simple tui chess

## Usage
Type `goch` in terminal to run it.

Movement:
 - hjkl - move cursor left, down, up and right
 - \<space\>\/\<enter\> - select piece or move it
 - `:` - begin command mode

Commands:
 - `<col><row>` - move cursor at col and row
 - \<enter\> - run command
 - \<esc\> - quit command mode


## Installation
```sh
git clone https://github.com/enotinc/gosh.git
cd goch/cmd/goch
go install # or go build
```
