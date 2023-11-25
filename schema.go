package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	student_id      int
	first_name      string
	last_name       string
	email           string
	enrollment_date string
}

func (s Student) ID() int            { return s.student_id }
func (s Student) FirstName() string  { return s.first_name }
func (s Student) LastName() string   { return s.last_name }
func (s Student) Email() string      { return s.email }
func (s Student) EnrollDate() string { return s.enrollment_date }

func (s Student) FilterValue() string { return s.last_name }
func (s Student) Title() string       { return fmt.Sprintf("%s %s", s.first_name, s.last_name) }
func (s Student) Description() string {
	return fmt.Sprintf("%s | %s | %s", strconv.Itoa(s.student_id), s.email, s.enrollment_date)
}
