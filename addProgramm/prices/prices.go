package prices

import "fmt"
import "os"
import "strconv"
import "strings"

type TaxJob struct{
	TaxRate float64
	InputPrice []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxJob) Process(){
	result := make(map[string]float64)
	for _, price := range job.InputPrice{
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}
	


func NewTaxIncludedPriceJob(taxeRate float64) *TaxJob {
	read, err := os.ReadFile("prices/price.txt")
	// fmt.Println(read)
	if err != nil{
		fmt.Println("readerror")
		return &TaxJob{}
	}
	readStr := string(read)
	arrStr := strings.Split(readStr, "\n")
	readFloat := make([]float64, len(arrStr))
	for i, val := range arrStr{
		valTrim := strings.TrimSpace(val)
		res, _ := strconv.ParseFloat(valTrim, 64)
		readFloat[i] = res
	}
	
	// readFloat, err := strconv.ParseFloat(readStr, 64)
	// if err != nil{
	// 	return &TaxJob{}
	// }
	
	return &TaxJob{
	TaxRate: taxeRate,
	InputPrice: readFloat,
	//InputPrice: []float64{10, 20, 30},
	}
}