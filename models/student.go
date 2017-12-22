package models

import (
  "errors"
  "strconv"
)

var (
  StudentList map[string]*Student
  StudentSerialNo    int64
)

type Student struct {
  StudentId   string
  StudentName string
  ClassName   string
}

func init() {
  StudentList = make(map[string]*Student)
  StudentSerialNo = 201700000

  AddNewStudent(Student{"", "Sharon", "apple"})
  AddNewStudent(Student{"", "Summer", "apple"})
  AddNewStudent(Student{"", "Andy", "apple"})
  AddNewStudent(Student{"", "Jeff", "apple"})
  AddNewStudent(Student{"", "Jason", "apple"})
  AddNewStudent(Student{"", "Simmy", "apple"})
  AddNewStudent(Student{"", "Eric", "apple"})
  AddNewStudent(Student{"", "Rimmon", "apple"})
  AddNewStudent(Student{"", "Green", "apple"})

  AddNewStudent(Student{"", "Eddie", "Cafe"})
  AddNewStudent(Student{"", "Hanry", "Cafe"})
  AddNewStudent(Student{"", "Scars", "Cafe"})
  AddNewStudent(Student{"", "Chris", "Cafe"})
  AddNewStudent(Student{"", "Justin", "Cafe"})
  AddNewStudent(Student{"", "Tea", "Cafe"})
}

func GetStudent(studentId string) (student *Student, err error) {
  if v, ok := StudentList[studentId]; ok {
    return v, nil
  }
  return nil, errors.New("studentId Not Exist")
}

func GetAllStudents() map[string]*Student {
  return StudentList
}

func AddNewStudent(newStudent Student) (studentId string){
  StudentSerialNo = StudentSerialNo + 1
  newStudent.StudentId = "s" + strconv.FormatInt(StudentSerialNo, 10)
  StudentList[newStudent.StudentId] = &newStudent
  return newStudent.StudentId
}
