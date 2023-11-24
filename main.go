package main

import (
	"fmt"
	"os"
	
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ConnectDatabase()
	defer Db.Close()

	items := []list.Item{
		item{title: "getAllStudents()", desc: "Retrieves and displays all records from the students table"},
		item{title: "addStudent(...)", desc: "Inserts a new student record into the students table"},
		item{title: "updateStudentEmail(...)", desc: "Updates the email address for a student with the specified student_id"},
		item{title: "deleteStudent(...)", desc: "Deletes the record of the student with the specified student_id"},
	}

	m := Model{
		list: list.New(items, NewCustomKeys(CustomKeyMap), 0, 0),
	}
	m.list.Title = "Queries"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}