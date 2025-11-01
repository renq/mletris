package main

import "image/color"

var (
	colorI = color.RGBA{0x00, 0xff, 0xff, 0xff} // Cyan
	colorO = color.RGBA{0xff, 0xff, 0x00, 0xff} // Yellow
	colorT = color.RGBA{0x80, 0x00, 0xff, 0xff} // Purple
	colorS = color.RGBA{0x00, 0xff, 0x00, 0xff} // Green
	colorZ = color.RGBA{0xff, 0x00, 0x00, 0xff} // Red
	colorJ = color.RGBA{0x00, 0x00, 0xff, 0xff} // Blue
	colorL = color.RGBA{0xff, 0xa5, 0x00, 0xff} // Orange
)

func buildTiles() []Piece {
	return []Piece{
		// I piece (line)
		{
			data: [][]Tile{
				// x
				// x
				// x
				// x
				{
					{x: 0, y: -1, color: colorI},
					{x: 0, y: 0, color: colorI},
					{x: 0, y: 1, color: colorI},
					{x: 0, y: 2, color: colorI},
				},
				// xxxx
				{
					{x: -1, y: 0, color: colorI},
					{x: 0, y: 0, color: colorI},
					{x: 1, y: 0, color: colorI},
					{x: 2, y: 0, color: colorI},
				},
			},
		},

		// O piece (square)
		{
			data: [][]Tile{
				// xx
				// xx
				{
					{x: 0, y: 0, color: colorO},
					{x: 1, y: 0, color: colorO},
					{x: 0, y: 1, color: colorO},
					{x: 1, y: 1, color: colorO},
				},
			},
		},

		// T piece (purple)
		{
			data: [][]Tile{
				//  x
				// xxx
				{
					{x: -1, y: 0, color: colorT},
					{x: 0, y: 0, color: colorT},
					{x: 1, y: 0, color: colorT},
					{x: 0, y: -1, color: colorT},
				},
				// x
				// xx
				// x
				{
					{x: 0, y: -1, color: colorT},
					{x: 0, y: 0, color: colorT},
					{x: 0, y: 1, color: colorT},
					{x: 1, y: 0, color: colorT},
				},
				// xxx
				//  x
				{
					{x: -1, y: 0, color: colorT},
					{x: 0, y: 0, color: colorT},
					{x: 1, y: 0, color: colorT},
					{x: 0, y: 1, color: colorT},
				},
				//  x
				// xx
				//  x
				{
					{x: 0, y: -1, color: colorT},
					{x: 0, y: 0, color: colorT},
					{x: 0, y: 1, color: colorT},
					{x: -1, y: 0, color: colorT},
				},
			},
		},

		// S piece (green)
		{
			data: [][]Tile{
				//  xx
				// xx
				{
					{x: 0, y: 0, color: colorS},
					{x: 1, y: 0, color: colorS},
					{x: -1, y: 1, color: colorS},
					{x: 0, y: 1, color: colorS},
				},
				// x
				// xx
				//  x
				{
					{x: -1, y: -1, color: colorS},
					{x: -1, y: 0, color: colorS},
					{x: 0, y: 0, color: colorS},
					{x: 0, y: 1, color: colorS},
				},
			},
		},

		// Z piece (red)
		{
			data: [][]Tile{
				// xx
				//  xx
				{
					{x: -1, y: 0, color: colorZ},
					{x: 0, y: 0, color: colorZ},
					{x: 0, y: 1, color: colorZ},
					{x: 1, y: 1, color: colorZ},
				},
				//  x
				// xx
				// x
				{
					{x: 1, y: -1, color: colorZ},
					{x: 0, y: 0, color: colorZ},
					{x: 1, y: 0, color: colorZ},
					{x: 0, y: 1, color: colorZ},
				},
			},
		},

		// J piece (blue)
		{
			data: [][]Tile{
				//  x
				//  x
				// xx
				{
					{x: 0, y: -1, color: colorJ},
					{x: 0, y: 0, color: colorJ},
					{x: 0, y: 1, color: colorJ},
					{x: -1, y: 1, color: colorJ},
				},
				// x
				// xxx
				{
					{x: -1, y: -1, color: colorJ},
					{x: -1, y: 0, color: colorJ},
					{x: 0, y: 0, color: colorJ},
					{x: 1, y: 0, color: colorJ},
				},
				// xx
				// x
				// x
				{
					{x: 0, y: -1, color: colorJ},
					{x: 1, y: -1, color: colorJ},
					{x: 0, y: 0, color: colorJ},
					{x: 0, y: 1, color: colorJ},
				},
				// xxx
				//   x
				{
					{x: -1, y: 0, color: colorJ},
					{x: 0, y: 0, color: colorJ},
					{x: 1, y: 0, color: colorJ},
					{x: 1, y: 1, color: colorJ},
				},
			},
		},

		// L piece (orange)
		{
			data: [][]Tile{
				// x
				// x
				// xx
				{
					{x: 0, y: -1, color: colorL},
					{x: 0, y: 0, color: colorL},
					{x: 0, y: 1, color: colorL},
					{x: 1, y: 1, color: colorL},
				},
				// xxx
				// x
				{
					{x: -1, y: 0, color: colorL},
					{x: 0, y: 0, color: colorL},
					{x: 1, y: 0, color: colorL},
					{x: -1, y: 1, color: colorL},
				},
				// xx
				//  x
				//  x
				{
					{x: -1, y: -1, color: colorL},
					{x: 0, y: -1, color: colorL},
					{x: 0, y: 0, color: colorL},
					{x: 0, y: 1, color: colorL},
				},
				//   x
				// xxx
				{
					{x: -1, y: 0, color: colorL},
					{x: 0, y: 0, color: colorL},
					{x: 1, y: 0, color: colorL},
					{x: 1, y: -1, color: colorL},
				},
			},
		},
	}
}
