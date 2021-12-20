package day20

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2021/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day20", "day-20-input.txt")
	solvePart1(lines)
	solvePart2(lines)
}

func solvePart1(lines []string) int {
	ip, img := parseInput(lines)
	start := time.Now().UnixMilli()
	for i := 0; i < 2; i++ {
		nImg := ip.Enhance(img)
		img = nImg
	}
	ans := img.PixelsLit()
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 1 (%dms): Pixels Lit = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	ip, img := parseInput(lines)
	start := time.Now().UnixMilli()
	for i := 0; i < 50; i++ {
		nImg := ip.Enhance(img)
		img = nImg
	}
	ans := img.PixelsLit()
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 2 (%dms): Pixels Lit = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) (*ImageProcessor, *Image) {
	ip := &ImageProcessor{enhanceAlgo: lines[0]}

	height := len(lines) - 2
	pixels := make([][]rune, height)
	for i := 2; i < len(lines); i++ {
		pixels[i-2] = []rune(lines[i])
	}
	width := len(pixels[0])
	i := &Image{pixels: pixels, width: width, height: height, expanseBit: '.'}

	return ip, i
}

type Image struct {
	pixels     [][]rune
	width      int
	height     int
	expanseBit rune
}

func (img *Image) PixelsLit() int {
	lit := 0
	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			if img.pixels[y][x] == '#' {
				lit++
			}
		}
	}
	return lit
}

func (img *Image) pixelAtIndex(x int, y int) rune {
	if x < 0 || y < 0 || x > img.width-1 || y > img.height-1 {
		return img.expanseBit
	}
	return img.pixels[y][x]
}

func (img *Image) String() string {
	sb := strings.Builder{}
	for _, p := range img.pixels {
		sb.WriteString(string(p))
		sb.WriteRune('\n')
	}
	return sb.String()
}

type ImageProcessor struct {
	enhanceAlgo string
}

func (ip *ImageProcessor) Enhance(img *Image) *Image {
	// prepare new image
	nImg := &Image{width: img.width + 2, height: img.height + 2}
	nImg.pixels = make([][]rune, nImg.height)
	for i := 0; i < nImg.height; i++ {
		nImg.pixels[i] = make([]rune, nImg.width)
	}

	// process each pixel
	var wp [3][3]rune
	for i := 0; i < nImg.height; i++ {
		for j := 0; j < nImg.width; j++ {
			for y := 0; y < 3; y++ {
				for x := 0; x < 3; x++ {
					wp[y][x] = img.pixelAtIndex(j-1+x-1, i-1+y-1)
				}
			}

			eIdx := pixelArrayToNumber(wp)
			nImg.pixels[i][j] = rune(ip.enhanceAlgo[eIdx])
		}
	}

	// adjust the expanse bit based on current value
	//  - if currently '.', then use bit 0 of the algo (since would be binary all 0's)
	//  - if currently '#', then use last bit of the algo (since would be binary all 1's)
	if img.expanseBit == '.' {
		nImg.expanseBit = rune(ip.enhanceAlgo[0])
	}
	if img.expanseBit == '#' {
		nImg.expanseBit = rune(ip.enhanceAlgo[len(ip.enhanceAlgo)-1])
	}

	return nImg
}

func pixelArrayToNumber(pixels [3][3]rune) int {
	out := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if pixels[i][j] == '#' {
				bitIdx := 9 - (i * 3) - j - 1
				out |= 1 << bitIdx
			}
		}
	}
	return out
}
