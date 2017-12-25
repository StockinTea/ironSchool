package models

import (
  "strconv"
)

var (
  TeacherList     map[string]*Teacher
  TeacherSerialNo int64
)

type Teacher struct {
  TeacherId   string
  TeacherName string
  TeacherAge  int
  TeacherTel  string
}

func init() {
  TeacherList = make(map[string]*Teacher)
  TeacherSerialNo = 48010000

  AddNewTeacher(Teacher{"", "Wayen", 32, "0987-335452"})
  AddNewTeacher(Teacher{"", "Tony", 54, "0969-578693"})
  AddNewTeacher(Teacher{"", "Black", 24, "0989-485321"})
  AddNewTeacher(Teacher{"", "Angelababy", 26, "0912-300245"})
  AddNewTeacher(Teacher{"", "S.Dan", 27, "0966-452970"})
}

func GetAllTeachers() map[string]*Teacher {
  return TeacherList
}

func AddNewTeacher(newTeacher Teacher) string {
  TeacherSerialNo = TeacherSerialNo + 1
  newTeacher.TeacherId = "t" + strconv.FormatInt(TeacherSerialNo, 10)
  TeacherList[newTeacher.TeacherId] = &newTeacher
  return newTeacher.TeacherId
}

func DeleteTeacher(tid string) {
  delete(TeacherList, tid)
}
