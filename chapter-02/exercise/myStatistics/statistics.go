package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	pageTop    = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

func main() {
    http.HandleFunc("/", homePage)
    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatal("failed to start server", err)
    }
}

func homePage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	fmt.Fprint(w, pageTop, form)
	if err != nil {
		fmt.Fprint(w, anError, err)
	} else {
		if numbers, message, ok := processRequest(r); ok {
			stats := getStats(numbers)
			fmt.Fprint(w, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(w, anError, message)
		}
	}
	fmt.Fprint(w, pageBottom)
}

func processRequest(r *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := r.Form["numbers"]; found && len(slice)>0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, " ' " + field + " ' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
	<tr><th colspan="2">Results</th></tr>
	<tr><td>Numbers</td><td>%v</td></tr>
	<tr><td>Count</td><td>%d</td></tr>
	<tr><td>Mean</td><td>%f</td></tr>
	<tr><td>Median</td><td>%f</td></tr>
	<tr><td>Stand. Deviation</td><td>%f</td></tr>
	<tr><td>Mode</td><td>%.2f</td></tr>
	</table>`,stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.stdDev, stats.mode)
}


type statistics struct {
	numbers []float64
	mean float64
	median float64
	stdDev float64
	mode []float64
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	stats.stdDev = standardDeviation(numbers, stats.mean)
	stats.mode = calculateMode(numbers)
	return stats
}

func sum(numbers []float64) (total float64) {
	for _, v := range numbers {
		total += v
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result 
}

func standardDeviation(numbers []float64, mean float64) (result float64) {
	n := float64(len(numbers))
	var total float64
	for _, value := range numbers{
		
		total += (math.Pow((value-mean), 2)) / (n-1)
		result = math.Sqrt(total)
	}
	return result
}

func calculateMode(numbers []float64) []float64 {
	var max int
	var result []float64
	numberCount := make(map[float64]int)
	for _, v := range numbers {
		numberCount[v]++
	}
	
	for _, value := range numberCount {
		if value > max {
			max = value
		}
	}
	if max == 1 {
		return []float64{}
	}
	for key, value := range numberCount {
		if value == max {
			result = append(result, key)
		}
	}
	return result
}