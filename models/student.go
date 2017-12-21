package models

import (
  "errors"
  "strconv"
)

var (
  StudentList map[string]*Student
  SerialNo    int64
)

type Student struct {
  StudentId   string
  StudentName string
  ClassName   string
}

func init() {
  StudentList = make(map[string]*Student)
  StudentList["s201700001"] = &Student{"s201700001", "Sharon", "apple"}
  StudentList["s201700002"] = &Student{"s201700002", "Summer", "apple"}
  StudentList["s201700003"] = &Student{"s201700003", "Andy", "apple"}
  StudentList["s201700004"] = &Student{"s201700004", "Jeff", "apple"}
  StudentList["s201700005"] = &Student{"s201700005", "Devin", "apple"}
  StudentList["s201700006"] = &Student{"s201700006", "Jason", "apple"}
  StudentList["s201700007"] = &Student{"s201700007", "Simmy", "apple"}
  StudentList["s201700008"] = &Student{"s201700008", "Rimmon", "apple"}
  StudentList["s201700009"] = &Student{"s201700009", "Eric", "apple"}
  StudentList["s201700010"] = &Student{"s201700010", "Green", "apple"}

  StudentList["s201700011"] = &Student{"s201700011", "Eddie", "Cafe"}
  StudentList["s201700012"] = &Student{"s201700012", "Scars", "Cafe"}
  StudentList["s201700013"] = &Student{"s201700013", "Hanry", "Cafe"}
  StudentList["s201700014"] = &Student{"s201700014", "Ben", "Cafe"}
  StudentList["s201700015"] = &Student{"s201700015", "Chris", "Cafe"}
  StudentList["s201700016"] = &Student{"s201700016", "Justin", "Cafe"}
  StudentList["s201700017"] = &Student{"s201700017", "Tea", "Cafe"}
  SerialNo = 201700018
}

func AddStudent(student Student) (studentId string) {
  student.StudentId = "s" + strconv.FormatInt(SerialNo, 8)
  StudentList[student.StudentId] = &student
  SerialNo = SerialNo + 1
  return student.StudentId
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
