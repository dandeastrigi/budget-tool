package model

import (
	"log"
	"math"
	"strconv"
)

type Budget struct {
	Total            float64
	Tax              float64
	Installment      string
	InstallmentValue float64
}

// TestModel : test a model
func TestModel(budget Budget) {
	log.Println(budget.Total)
	log.Println(budget.Installment)
}

// CalcInstallments : calc
func CalcInstallments(budget Budget) Budget {
	i, err := strconv.Atoi(budget.Installment)
	if err != nil {

	}
	installment := float64(i)
	pa := budget.Total * (budget.Tax / 100)
	pb := 1 + (budget.Tax / 100)
	pb = math.Pow(pb, installment)
	pa = pb * pa
	pb = math.Pow(1+(budget.Tax/100), installment) - 1
	budget.InstallmentValue = pa / pb
	return budget
}
