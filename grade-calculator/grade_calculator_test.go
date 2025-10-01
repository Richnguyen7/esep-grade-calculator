package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected := "A"
	gc := NewGradeCalculator()
	gc.AddGrade("open source assignment", 100, Assignment)
	gc.AddGrade("exam 1", 100, Exam)
	gc.AddGrade("essay on ai ethics", 100, Essay)

	if got := gc.GetFinalGrade(); got != expected {
		t.Fatalf("want %s, got %s", expected, got)
	}
}

func TestGetGradeB(t *testing.T) {
	expected := "B"
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 80, Assignment)
	gc.AddGrade("exam 1", 81, Exam)
	gc.AddGrade("essay", 85, Essay)

	if got := gc.GetFinalGrade(); got != expected {
		t.Fatalf("want %s, got %s", expected, got)
	}
}

func TestGetGradeF(t *testing.T) {
	expected := "A"
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 100, Assignment)
	gc.AddGrade("exam 1", 95, Exam)
	gc.AddGrade("essay", 91, Essay)

	if got := gc.GetFinalGrade(); got != expected {
		t.Fatalf("want %s, got %s", expected, got)
	}
}

func TestEmptyCategoriesReturnsF(t *testing.T) {
	gc := NewGradeCalculator()
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("empty categories should be F, got %s", got)
	}
}

func TestMissingCategoriesAreZeroWeighted(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("only-assign", 90, Assignment)
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("with only assignments=90, final should be F, got %s", got)
	}
}

func TestValidationClampsOutOfRange(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("neg", -5, Assignment)
	gc.AddGrade("big", 150, Exam)
	gc.AddGrade("ok", 60, Essay)
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("clamped values should yield F, got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String() mismatch")
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String() mismatch")
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String() mismatch")
	}
}

func TestRoundingBoundaries(t *testing.T) {
	tests := []struct {
		assign int
		exam   int
		essay  int
		want   string
	}{
		{59, 60, 60, "D"},
		{69, 70, 70, "C"},
		{79, 80, 80, "B"},
		{89, 90, 90, "A"},
		{59, 59, 60, "F"},
	}

	for _, tt := range tests {
		gc := NewGradeCalculator()
		gc.AddGrade("a1", tt.assign, Assignment)
		gc.AddGrade("a2", tt.assign, Assignment)
		gc.AddGrade("e1", tt.exam, Exam)
		gc.AddGrade("e2", tt.exam, Exam)
		gc.AddGrade("s1", tt.essay, Essay)
		gc.AddGrade("s2", tt.essay, Essay)

		if got := gc.GetFinalGrade(); got != tt.want {
			t.Fatalf("want %s, got %s", tt.want, got)
		}
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
