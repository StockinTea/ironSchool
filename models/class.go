package models

import (
  "errors"
  "strings"
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

  ClassList["c201700000"] = &Class{"c201700000", "apple"}
  ClassList["c201700001"] = &Class{"c201700000", "Cafe"}
  ClassList["c201700002"] = &Class{"c201700000", "dog"}
  ClassSerialNo = 201700018
}

func GetClassById(classId string) (class *Class, err error) {
  if v, ok := ClassList[classId]; ok {
    return v, nil
  }
  return nil, errors.New("classId Not Exist")
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
