# Skene
Skene is a word of Greek origin that means "stage" or "scenery". In ancient theater, it was the structure where scenes were prepared and presented — symbolizing creation and expression.

This a template software is a GUI made in GO.

## Installing && Running 

### Pre-requisites

I also recommend to use `g` to manage go versions. ( [check g ](https://github.com/voidint/g) )

- You need `go` to run this, I'm using `1.24.4`
- linter `golanglint`: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
- bundler and gui package `fyne`: `go install fyne.io/tools/cmd/fyne`

### Installing

```bash
    git clone https://github.com/afa7789/skene
    cd skene
    cp .env.example .env # fill your env
```

### Running

Most of the commands in this project is in the makefile, so just run `make help` to see what it is doing, such as `make test`, `make lint` and `make run`.

### Executable

Creating a binary, or executable let you have a executable desktop app, for that you have to use `make package`.

- If you are looking into the changing icon, look at the [`icon.go`](internal/gui/icon.go) for more information.

## Internationalization // Location

All the languages and texts in the GUI can be found in `internal/localization/locales`.