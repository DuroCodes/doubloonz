# Doubloonz

<div align="center">
  <img src="./logo.png" width="400" />
</div>

Doubloonz is a CLI tool that helps you determine when you'll be able to use your doubloons to purchase a prize from the doubloon shop. **Heavily** inspired by another project, [Doubloon Project Ranker](https://doubloon-project-ranker.vercel.app/).

You can choose your region within the CLI, so the prizes will be accurate for your region. (Unless I messed up making the prizes, but I don't think I did)

I'm not personally a big fan of the [Go](https://golang.org/) programming language, but I really like the [Charm](https://github.com/charmbracelet) libraries for CLI tools. Go is kind of growing on me, though.

> [!WARNING]
> This project uses [nerd fonts](https://www.nerdfonts.com/). Make sure you have one installed to see the icons for doubloons

## Usage

> [!TIP]
> You can use `go install github.com/durocodes/doubloonz@latest` to install the binary to your `$GOPATH/bin` without manually building

1. Clone the repository
2. Run `go build` to build the binary
3. Run `./doubloonz` (or `./doubloonz.exe` on Windows) to start the CLI

## Updating Prizes

If the prizes change, create a file called `data.html`, and copy/paste the HTML from the shop page, and you can use the `scraper.py` script to print out the prizes in a format that can be copied into the `prizes.json` file.
