# Gonways Game of Life

[![Go Reference](https://pkg.go.dev/badge/github.com/DumDumGeniuss/ggol.svg)](https://pkg.go.dev/github.com/DumDumGeniuss/ggol)
[![Go Report Card](https://goreportcard.com/badge/github.com/DumDumGeniuss/ggol)](https://goreportcard.com/report/github.com/DumDumGeniuss/ggol)
[![Build Status](https://app.travis-ci.com/DumDumGeniuss/ggol.svg?branch=main)](https://app.travis-ci.com/DumDumGeniuss/ggol)

Gonways Game of Life is a go package that provides fundamental functions for running [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life),

The goal is to help you build a Conway's Game of Life in the way you like.

## Installed This Package

```bash
go get github.com/DumDumGeniuss/ggol
```

## Initialize A New Game

```go
package main

import {
    "fmt"
    
    "github.com/DumDumGeniuss/ggol"
)

main() {
    // This seed will bring cells in 2nd row back to life.
    seed := ggol.Seed{
        {Coordinate: ggol.Coordinate{X: 0, Y: 1}, Cell: true},
        {Coordinate: ggol.Coordinate{X: 1, Y: 1}, Cell: true},
        {Coordinate: ggol.Coordinate{X: 2, Y: 1}, Cell: true},
    }
    // Create a size
    size := ggol.Size{Height: 3, Width: 3}
    // Start a new game with given seed.
    game, _ := ggol.NewGame(&size)
    game.PlantSeed(&seed)
    // Generate next cellLiveMap.
    game.Evolve()
    // Get current cellLiveMap.
    newCellLiveStatusMap := *game.GetCellLiveStatusMap()
    // We digonally rotate the cellLiveMap so it's easier to read.
    rotatedNewCellLiveStatusMap := ggol.RotateCellLiveStatusMapInDigonalLine(newCellLiveStatusMap)
    fmt.Println(rotatedNewCellLiveStatusMap)
    // [
    //     [false true false]
    //     [false true false]
    //     [false true false]
    // ]
}
```

## Demo

You can see a quick demo by cloning this repo to your local machine.

```bash
git clone https://github.com/DumDumGeniuss/ggol.git
cd ggol
go mod tidy
go run ./cmd/main.go
# [GIN-debug] Listening and serving HTTP on :8000
```

And you can open your browser and view the demo on [http://localhost:8000/demo](http://localhost:8000/demo)

![demo](./doc/demo.png)

## Document

### NewGame

```go
import "github.com/DumDumGeniuss/ggol"

game, _ := ggol.NewGame(&size, &seed)

```
