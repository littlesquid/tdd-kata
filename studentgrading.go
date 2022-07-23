package main

func StudentGrading(grades []int32) []int32 {
	roundUpGrades := []int32{}
	for _, grade := range grades {
		if grade >= 38 {
			roundUpGrade := roundUp(grade)
			if (roundUpGrade - grade) < 3 {
				roundUpGrades = append(roundUpGrades, roundUpGrade)
				continue
			}
		}
		roundUpGrades = append(roundUpGrades, grade)

	}
	return roundUpGrades
}

func roundUp(n int32) int32 {
	for {
		n += 1
		if n%5 == 0 {
			return n
		}
	}
}
