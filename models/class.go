package models

import (
  "errors"
  "strings"
  "strconv"
)

var (
  ClassList map[string]*Class
  ClassSerialNo  int64
)

type Class struct {
  ClassId   string
  ClassName   string
}

func init() {
  ClassList = make(map[string]*Class)
  ClassSerialNo = 20170000

  AddNewClass(Class{ "", "apple"})
  AddNewClass(Class{ "", "Cafe"})
  AddNewClass(Class{ "", "dog"})

}

func GetClassById(classId string) (class *Class, err error) {
  if v, ok := ClassList[classId]; ok {
    return v, nil
  }
  return nil, errors.New("NotExist")
}


func GetClassByName(className string) (class *Class, err error) {
  for index, value := range ClassList {
    if strings.EqualFold(value.ClassName, className) {
      return ClassList[index], nil
    }
  }

  return nil, errors.New("NotExist")
}

func GetAllClasses() map[string]*Class {
  return ClassList
}

func AddNewClass(newClass Class) string {
  ClassSerialNo = ClassSerialNo + 1
  newClass.ClassId = "c" + strconv.FormatInt(ClassSerialNo, 10)
  ClassList[newClass.ClassId] = &newClass
  return newClass.ClassId
}
