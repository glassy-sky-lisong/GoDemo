package StudentMgr

type student struct {
  name string
  age int64
}

func NewStudent(name string, age int64) stu *student {
  return &student{name, age}
}

