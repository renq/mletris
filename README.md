# mletris

A simple Tetris clone written in Go.

This repository contains a lightweight, minimal implementation of Tetris which I built as a learning project while getting familiar with the Go programming language. It's intentionally basic and serves primarily as a sandbox for experimenting with Go's syntax, concurrency primitives, and standard library. Expect rough edges and few extra features — it's a pet project, not a polished game.

<!-- ![Animation placeholder](animation.gif) -->

## Features

- Basic Tetris gameplay (falling blocks, line clears)
- Console-based/graphical renderer written in Go
- No dependencies beyond the Go standard library
- Designed for clarity over performance

## Getting Started

To build and run the game locally, make sure you have Go installed (tested with Go 1.20+):

```bash
git clone https://codeberg.org/renq/mletris.git
cd mletris

go build -o mletris
./mletris
```

*Note:* the animation above is just a placeholder; I'll replace it with an actual GIF or video demonstrating gameplay once it's ready.

## WebAssembly (optional)

You can also run the project in the browser using WebAssembly:

```bash
go run github.com/hajimehoshi/wasmserve@latest ./
```

Then open `localhost:8080` in your browser.

## License

This project is distributed under the MIT License. See [LICENSE](LICENSE) for details.

---

*This project is primarily an exercise in learning Go and should be considered experimental.*
