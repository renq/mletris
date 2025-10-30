package main

import "image/color"

// TODO: Create variables for colors used in tiles
func buildTiles() []Piece {
	return []Piece{
		// I piece (line)
		Piece{
			data: [][]Tile{
				// x
				// x
				// x
				// x
				{
					Tile{x:0, y:-1, color: color.RGBA{0x00, 0xff, 0xff, 0xff}}, // cyan
					Tile{x:0, y:0, color: color.RGBA{0x00, 0xff, 0xff, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0x00, 0xff, 0xff, 0xff}},
					Tile{x:0, y:2, color: color.RGBA{0x00, 0xff, 0xff, 0xff}},
				},
				// xxxx
				{
					Tile{x:-1, y:0, color: color.RGBA{0x00, 0xff, 0xff, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x00, 0xff, 0xff, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0x00, 0xff, 0xff, 0xff}},
					Tile{x:2, y:0, color: color.RGBA{0x00, 0xff, 0xff, 0xff}},
				},
			},
		},

		// O piece (square)
		Piece{
			data: [][]Tile{
				// xx
				// xx
				{
					Tile{x:0, y:0, color: color.RGBA{0xff, 0xff, 0x00, 0xff}}, // yellow
					Tile{x:1, y:0, color: color.RGBA{0xff, 0xff, 0x00, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0xff, 0xff, 0x00, 0xff}},
					Tile{x:1, y:1, color: color.RGBA{0xff, 0xff, 0x00, 0xff}},
				},
			},
		},

		// T piece (purple)
		Piece{
			data: [][]Tile{
				//  x
				// xxx
				{
					Tile{x:-1, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:0, y:-1, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
				},
				// x
				// xx
				// x
				{
					Tile{x:0, y:-1, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
				},
				// xxx
				//  x
				{
					Tile{x:-1, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
				},
				//  x
				// xx
				//  x
				{
					Tile{x:0, y:-1, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
					Tile{x:-1, y:0, color: color.RGBA{0x80, 0x00, 0xff, 0xff}},
				},
			},
		},

		// S piece (green)
		Piece{
			data: [][]Tile{
				//  xx
				// xx
				{
					Tile{x:0, y:0, color: color.RGBA{0x00, 0xff, 0x00, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0x00, 0xff, 0x00, 0xff}},
					Tile{x:-1, y:1, color: color.RGBA{0x00, 0xff, 0x00, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0x00, 0xff, 0x00, 0xff}},
				},
				// x
				// xx
				//  x
				{
					Tile{x:-1, y:-1, color: color.RGBA{0x00, 0xff, 0x00, 0xff}},
					Tile{x:-1, y:0, color: color.RGBA{0x00, 0xff, 0x00, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x00, 0xff, 0x00, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0x00, 0xff, 0x00, 0xff}},
				},
			},
		},

		// Z piece (red)
		Piece{
			data: [][]Tile{
				// xx
				//  xx
				{
					Tile{x:-1, y:0, color: color.RGBA{0xff, 0x00, 0x00, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0xff, 0x00, 0x00, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0xff, 0x00, 0x00, 0xff}},
					Tile{x:1, y:1, color: color.RGBA{0xff, 0x00, 0x00, 0xff}},
				},
				//  x
				// xx
				// x
				{
					Tile{x:1, y:-1, color: color.RGBA{0xff, 0x00, 0x00, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0xff, 0x00, 0x00, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0xff, 0x00, 0x00, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0xff, 0x00, 0x00, 0xff}},
				},
			},
		},

		// J piece (blue)
		Piece{
			data: [][]Tile{
				//  x
				//  x
				// xx
				{
					Tile{x:0, y:-1, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:-1, y:1, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
				},
				// x
				// xxx
				{
					Tile{x:-1, y:-1, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:-1, y:0, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
				},
				// xx
				// x
				// x
				{
					Tile{x:0, y:-1, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:1, y:-1, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
				},
				// xxx
				//   x
				{
					Tile{x:-1, y:0, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
					Tile{x:1, y:1, color: color.RGBA{0x00, 0x00, 0xff, 0xff}},
				},
			},
		},

		// L piece (orange)
		Piece{
			data: [][]Tile{
				// x
				// x
				// xx
				{
					Tile{x:0, y:-1, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:1, y:1, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
				},
				// xxx
				// x
				{
					Tile{x:-1, y:0, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:-1, y:1, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
				},
				// xx
				//  x
				//  x
				{
					Tile{x:-1, y:-1, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:0, y:-1, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:0, y:1, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
				},
				//   x
				// xxx
				{
					Tile{x:-1, y:0, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:0, y:0, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
					Tile{x:1, y:-1, color: color.RGBA{0xff, 0xa5, 0x00, 0xff}},
				},
			},
		},
	};
}