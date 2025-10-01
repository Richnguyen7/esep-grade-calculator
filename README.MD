# 📚 ESEP Grade Calculator

This project is a small Go (Golang) library that computes a student’s final grade 
for an educational product. It supports weighted averages across assignments, exams, 
and essays, then maps the result to a letter grade (A–F) and Pass/Fail.

---

## Features

- Weighted average calculation:
  - Assignments: 50%
  - Exams: 35%
  - Essays: 15%
- Letter grade mapping:
  - A ≥ 90  
  - B ≥ 80  
  - C ≥ 70  
  - D ≥ 60  
  - F < 60
- Pass/Fail check:
  - Pass = C or higher
- 100% test coverage with Go’s built-in test tools

---

## Prerequisites

- Go (Golang) version 1.20+  
  Download: [https://go.dev/dl](https://go.dev/dl)  
  Verify installation:

  ```bash
  go version


## Running Tests
cd grade-calculator
go test ./...
cd grade-calculator
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
go tool cover -html=coverage.txt -o coverage.html
