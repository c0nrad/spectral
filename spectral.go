package main

import (
	"math"
	"math/big"
)

func main() {
	// for to := 1; to < 6; to++ {
	// 	for from := to + 1; from < 6; from++ {
	// 		fmt.Println("From", from, "to", to, En(from)-En(to), energyToWavelength(En(from)-En(to)))
	// 	}
	// }
	// fmt.Println(En(1), energyToWavelength(En(1)))
	Draw()
}

func energyToWavelength(eV float64) (nm float64) {
	in := big.NewFloat(eV)
	hc := big.NewFloat(1239.841984)
	in.Quo(hc, in)
	out, _ := in.Float64()
	return out
}

func En(n int) float64 {

	// SI units
	electronMass := big.NewFloat(9.1093837015e-31)
	hBar := big.NewFloat(1.054571817e-34)
	permitivity := big.NewFloat(8.8541878128e-12)
	pi := big.NewFloat(math.Pi)
	e := big.NewFloat(1.602176634e-19)
	jToEv := big.NewFloat(1.602176634e-19)

	// (e^2 / 4*pi*e_0)^2
	tmp := new(big.Float).Mul(e, e)
	tmp.Quo(tmp, big.NewFloat(4))
	tmp.Quo(tmp, pi)
	tmp.Quo(tmp, permitivity)
	tmp.Mul(tmp, tmp)

	// m_e / 2 hbar^2 (tmp)
	tmp.Mul(tmp, electronMass)
	tmp.Quo(tmp, big.NewFloat(2))
	tmp.Quo(tmp, hBar)
	tmp.Quo(tmp, hBar)

	// -tmp (1/n^2)
	tmp.Quo(tmp, big.NewFloat(float64(n)))
	tmp.Quo(tmp, big.NewFloat(float64(n)))
	tmp.Mul(tmp, big.NewFloat(-1))

	tmp.Quo(tmp, jToEv)

	out, _ := tmp.Float64()
	return out
}
