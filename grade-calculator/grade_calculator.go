package esepunittests

import "math"

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	switch {
	case numericalGrade >= 90:
		return "A"
	case numericalGrade >= 80:
		return "B"
	case numericalGrade >= 70:
		return "C"
	case numericalGrade >= 60:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	if grade < 0 {
		grade = 0
	}
	if grade > 100 {
		grade = 100
	}

	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{Name: name, Grade: grade, Type: Assignment})
	case Exam:
		gc.exams = append(gc.exams, Grade{Name: name, Grade: grade, Type: Exam})
	case Essay:
		gc.essays = append(gc.essays, Grade{Name: name, Grade: grade, Type: Essay})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentAvg := computeAverage(gc.assignments)
	examAvg := computeAverage(gc.exams)
	essayAvg := computeAverage(gc.essays)

	weighted := assignmentAvg*0.50 + examAvg*0.35 + essayAvg*0.15
	return int(math.Round(weighted))
}

// computeAverage safely handles empty slices.
func computeAverage(grades []Grade) float64 {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for _, g := range grades { 
		sum += g.Grade
	}
	return float64(sum) / float64(len(grades))
}
