package esepunittests

import "math"

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

func (gt GradeType) String() string { return gradeTypeName[gt] }

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

type outputMode int

const (
	modeLetter outputMode = iota
	modePassFail
)

type entry struct {
	name  string
	grade int
	kind  GradeType
}

type GradeCalculator struct {
	mode  outputMode
	items []entry
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{mode: modeLetter, items: make([]entry, 0)}
}

func NewGradeCalculatorPassFail() *GradeCalculator {
	return &GradeCalculator{mode: modePassFail, items: make([]entry, 0)}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	n := gc.calculateNumericalGrade()
	switch gc.mode {
	case modePassFail:
		if n >= 70 {
			return "Pass"
		}
		return "Fail"
	default:
		switch {
		case n >= 90:
			return "A"
		case n >= 80:
			return "B"
		case n >= 70:
			return "C"
		case n >= 60:
			return "D"
		default:
			return "F"
		}
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	if grade < 0 {
		grade = 0
	}
	if grade > 100 {
		grade = 100
	}
	gc.items = append(gc.items, entry{name: name, grade: grade, kind: gradeType})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	a := avgKind(gc.items, Assignment)
	e := avgKind(gc.items, Exam)
	s := avgKind(gc.items, Essay)
	w := a*0.50 + e*0.35 + s*0.15
	return int(math.Round(w))
}

func avgKind(es []entry, kind GradeType) float64 {
	count := 0
	sum := 0
	for _, it := range es {
		if it.kind == kind {
			sum += it.grade
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}
