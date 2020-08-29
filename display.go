package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	ColorBlack = "\033[30m"
	ColorWhite = "\033[37m"

	ColorRed     = "\033[31m" // 625 - 740
	ColorYellow  = "\033[33m" // 565 - 590
	ColorGreen   = "\033[32m" // 520 - 565
	ColorCyan    = "\033[36m" // 500 - 520
	ColorBlue    = "\033[34m" // 435 - 500
	ColorMagenta = "\033[35m" // 380 - 435

	ColorReset = "\033[0m"

	Clear = "\033[2J"

	Corner         = "+"
	HorizontalLine = "-"
	VerticalLine   = "|"

	SpectralWidth        = 25
	SpectralHeightOffset = 1
	SpectralWidthOffset  = 5
	MaxLevels            = 6
)

func Draw() {
	fmt.Println(Clear)
	DrawAxis()
	DrawEnergyLevels()
	DrawJumps()
	DrawLabels()
	DrawTitle()
	DrawTables()
	DrawText()
	Jump(90, 35)
	fmt.Println()
}

func DrawText() {
	Jump(55, 25)
	fmt.Println("Hydrogen Spectral Lines.")
	Jump(55, 26)
	fmt.Println("The photon energies/wavelengths emitted from electron")
	Jump(55, 27)
	fmt.Println("transitions in hydrogen. Transitions to n=2 are called")
	Jump(55, 28)
	fmt.Println("Balmer series and are in the visible light spectrum.")
}

func DrawLabels() {
	Jump(2, 15)
	fmt.Println("eV")
}

func eVToRow(ev float64) int {
	return int(math.Round(math.Abs(ev)*2 + 1))
}

func nmToColor(nm int) string {
	// ColorRed     = "\033[31m" // 625 - 740
	// ColorYellow  = "\033[33m" // 565 - 590
	// ColorGreen   = "\033[32m" // 520 - 565
	// ColorCyan    = "\033[36m" // 500 - 520
	// ColorBlue    = "\033[34m" // 435 - 500
	// ColorMagenta = "\033[35m" // 380 - 435

	if nm > 740 {
		return ColorWhite
	}

	if nm > 625 {
		return ColorRed
	}

	if nm > 565 {
		return ColorYellow
	}

	if nm > 520 {
		return ColorGreen
	}

	if nm > 500 {
		return ColorCyan
	}

	if nm > 435 {
		return ColorBlue
	}

	if nm > 380 {
		return ColorMagenta
	}

	return ColorWhite
}

func DrawAxis() {
	for ev := 0; ev < 16; ev++ {
		Jump(SpectralWidthOffset, SpectralHeightOffset+eVToRow(float64(ev)))
		fmt.Printf("%3d", -ev)
	}
}

func DrawEnergyLevels() {
	for n := 1; n < MaxLevels; n++ {
		en := En(n)
		row := eVToRow(en)

		Jump(SpectralWidthOffset+4, SpectralHeightOffset+row)
		fmt.Printf("%s n=%d e=%2.3f eV", strings.Repeat("-", SpectralWidth), n, en)
	}
}

func DrawTitle() {
	Jump(42, 1)
	fmt.Println(ColorRed, "Hydrogen", ColorBlue, "Spectral", ColorMagenta, "Lines", ColorReset)
}

func Jump(x, y int) {
	os.Stdout.WriteString(fmt.Sprintf("\033[%d;%df", y, x))
}

func DrawJumps() {
	col := 5
	for to := 1; to < MaxLevels; to++ {
		for from := to + 1; from < MaxLevels; from++ {
			fromE := En(from)
			toE := En(to)

			nm := energyToWavelength(fromE - toE)
			color := nmToColor(int(nm))
			os.Stdout.WriteString(color)
			fromRow := eVToRow(fromE)
			toRow := eVToRow(toE)

			row := fromRow + 1
			Jump(SpectralWidthOffset+col, SpectralHeightOffset+fromRow)
			fmt.Println("+")
			for row < toRow {
				Jump(SpectralWidthOffset+col, SpectralHeightOffset+row)
				fmt.Println("|")
				row++
			}
			Jump(SpectralWidthOffset+col, SpectralHeightOffset+toRow)
			fmt.Println("v")
			fmt.Println(ColorReset)
			col++
		}
		col += 3
	}
}

func DrawTables() {
	TableStartWidth := 55
	TableStartHeight := 3

	Jump(TableStartWidth+12, TableStartHeight)
	fmt.Println("From(n), To(n), Energy [eV]")
	Jump(TableStartWidth+1, TableStartHeight+1)
	fmt.Println("     1       2       3       4       5       6   ")
	Jump(TableStartWidth+1, TableStartHeight+2)
	fmt.Println("+------------------------------------------------")
	for to := 1; to <= MaxLevels; to++ {

		Jump(TableStartWidth+1, TableStartHeight+to+2)
		fmt.Print(to)
		for from := 1; from <= MaxLevels; from++ {
			if from < to+1 {
				if from == to {
					fmt.Printf("|   0   ")
				} else {
					fmt.Printf("|   -   ")
				}
				continue
			}
			fromE := En(from)
			toE := En(to)
			fmt.Printf("| %5.2f ", fromE-toE)
		}
	}

	TableStartWidth = 55
	TableStartHeight = 14
	Jump(TableStartWidth+10, TableStartHeight)
	fmt.Println("From(n), To(n), Wavelength [nm]")
	Jump(TableStartWidth+1, TableStartHeight+1)
	fmt.Println("     1       2       3       4       5       6   ")
	Jump(TableStartWidth+1, TableStartHeight+2)
	fmt.Println("+------------------------------------------------")
	for to := 1; to <= MaxLevels; to++ {

		Jump(TableStartWidth+1, TableStartHeight+to+2)
		fmt.Print(to)
		for from := 1; from <= MaxLevels; from++ {
			if from < to+1 {
				if from == to {
					fmt.Printf("|   0   ")
				} else {
					fmt.Printf("|   -   ")
				}
				continue
			}
			fromE := En(from)
			toE := En(to)
			nm := energyToWavelength(fromE - toE)

			color := nmToColor(int(nm))

			fmt.Printf("| %s%5.0f%s ", color, energyToWavelength(fromE-toE), ColorReset)
		}
	}
}
