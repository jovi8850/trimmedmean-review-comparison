package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jovi8850/go-trimmed-mean/trimmedmean"
)

func main() {

	// Seed RNG for reproducibility
	rand.Seed(time.Now().UnixNano())

	// ----------------------------------------------------------------------
	// Generate 100 integers (right-skewed on purpose)
	// ----------------------------------------------------------------------
	intData := make([]int, 100)
	for i := range intData {
		// mixture to induce skew: 90% N(50,10), 10% high outliers
		if rand.Float64() < 0.9 {
			intData[i] = int(rand.NormFloat64()*10 + 50)
		} else {
			intData[i] = int(rand.Float64()*400 + 200) // outliers
		}
	}

	// Save to CSV so user can check in R
	saveIntCSV("int_sample.csv", intData)

	// Compute symmetric 5% trimmed mean using Go package
	goIntTM, err := trimmedmean.TrimmedMeanInts(intData, 0.05)
	if err != nil {
		panic(err)
	}

	// NEW: Compute automatic trimmed mean for integers
	autoIntTM, intRec, err := trimmedmean.AutoTrimmedMeanInts(intData)
	if err != nil {
		panic(err)
	}

	// ----------------------------------------------------------------------
	// Generate 100 floats (right-skewed gamma-like)
	// ----------------------------------------------------------------------
	floatData := make([]float64, 100)
	for i := range floatData {
		// right-skewed exponential-like distribution
		floatData[i] = rand.ExpFloat64()*20 + 10
	}

	// Save for R evaluation
	saveFloatCSV("float_sample.csv", floatData)

	// Compute symmetric 5% trimmed mean in Go
	goFloatTM, err := trimmedmean.TrimmedMean(floatData, 0.05)
	if err != nil {
		panic(err)
	}

	// NEW: Compute automatic trimmed mean for floats
	autoFloatTM, floatRec, err := trimmedmean.AutoTrimmedMean(floatData)
	if err != nil {
		panic(err)
	}

	// ----------------------------------------------------------------------
	// Output results
	// ----------------------------------------------------------------------

	fmt.Println("\n======================================================")
	fmt.Println("                 TRIMMED MEAN TESTS")
	fmt.Println("======================================================\n")

	fmt.Printf("Go trimmed mean (integers), symmetric 5%%: %.6f\n", goIntTM)
	fmt.Printf("Go AUTO trimmed mean (integers): %.6f\n", autoIntTM)
	fmt.Printf("  - Recommended trimming: low=%.1f%%, high=%.1f%%\n",
		intRec.LowTrim*100, intRec.HighTrim*100)
	fmt.Printf("  - Skewness: %.4f\n", intRec.Skewness)
	fmt.Printf("  - Interpretation: %s\n", intRec.Interpretation)

	fmt.Printf("Go trimmed mean (floats), symmetric 5%%: %.6f\n", goFloatTM)
	fmt.Printf("Go AUTO trimmed mean (floats): %.6f\n", autoFloatTM)
	fmt.Printf("  - Recommended trimming: low=%.1f%%, high=%.1f%%\n",
		floatRec.LowTrim*100, floatRec.HighTrim*100)
	fmt.Printf("  - Skewness: %.4f\n", floatRec.Skewness)
	fmt.Printf("  - Interpretation: %s\n", floatRec.Interpretation)

	fmt.Println("CSV files written: int_sample.csv, float_sample.csv\n")

	// Demonstrate manual asymmetric trimming based on recommendations
	fmt.Println("======================================================")
	fmt.Println("          MANUAL ASYMMETRIC TRIMMING EXAMPLE")
	fmt.Println("======================================================")

	// Apply the exact recommended trimming manually
	manualIntTM, err := trimmedmean.TrimmedMeanInts(intData, intRec.LowTrim, intRec.HighTrim)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Manual asymmetric trimming (integers): %.6f\n", manualIntTM)
	fmt.Printf("  - Using low=%.1f%%, high=%.1f%%\n",
		intRec.LowTrim*100, intRec.HighTrim*100)

	manualFloatTM, err := trimmedmean.TrimmedMean(floatData, floatRec.LowTrim, floatRec.HighTrim)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Manual asymmetric trimming (floats): %.6f\n", manualFloatTM)
	fmt.Printf("  - Using low=%.1f%%, high=%.1f%%\n",
		floatRec.LowTrim*100, floatRec.HighTrim*100)

	fmt.Println("\nDone.")
}

// ----------------------------------------------------------------------
// Utility functions
// ----------------------------------------------------------------------

func saveIntCSV(filename string, data []int) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, v := range data {
		w.Write([]string{fmt.Sprintf("%d", v)})
	}
}

func saveFloatCSV(filename string, data []float64) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, v := range data {
		w.Write([]string{fmt.Sprintf("%f", v)})
	}
}

func convertIntsToFloat(v []int) []float64 {
	out := make([]float64, len(v))
	for i := range v {
		out[i] = float64(v[i])
	}
	return out
}
