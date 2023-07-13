package data

import (
	"belajar/bareng/features"
)

func AverageRatingsInsert(review []features.Review, inputRating float64) (float64, error) {

	tampungRating:=TampungRating(review)
	tampungRating = append(tampungRating, inputRating)
	total,count := SumRating(tampungRating)
	average:=Avr(total,count)
	return average, nil
}

func AverageRatingsDelete(review []features.Review, inputRating float64)(float64,error){
	tampungRating:=TampungRating(review)
	total,count:=DeleteRating(tampungRating,inputRating)
	average:=Avr(total,count)
	return average,nil
}

func DeleteRating(tampungRating []float64,inputRating float64)(float64,int){
	total,count:=SumRating(tampungRating)
	total = total-inputRating
	count = count - 1
	return total,count
}

func TampungRating(review []features.Review) []float64{
	var tampungRating []float64
	for _, lastRatings := range review {
		tampungRating = append(tampungRating, lastRatings.Rating)
	}
	return tampungRating	
}

func SumRating(tampungRating []float64) (float64,int){
	var total float64
	var count int
	for _, value := range tampungRating {
		total += value
		count++
	}
	return total,count
}

func Avr(total float64, count int) float64{
	var average float64
	if count > 0 {
		average = total / float64(count)
	}
	return average
}
