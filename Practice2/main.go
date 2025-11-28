package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Приветики пистолетики, я " + p.Name
}

type Pm struct {
	Name string
}

func (p Pm) Great() string {
	return "Здравствуйте коллеги, я " + p.Name
}

type Worker interface {
	Work()
}

type WorkerImpl struct {
	workDone bool
}

func (w WorkerImpl) Work() {
	fmt.Println("Арбйтен арбайтен")
	w.workDone = true
}

func NewWorker() Worker {
	var w Worker
	// w := new(WorkerImpl)

	return w
}

type Mytype struct {
	Name string
}

func (Mytype) Val() {

}

func (*Mytype) Val2() {

}

type V interface {
	Val()
}

type v2 interface {
	Val2()
}

type Saver interface {
	Save()
}

type Config struct {
	Name string
}

func (c *Config) Save() error {
	fmt.Println("config saved to memory")

	return nil
}

func main() {
	p := Person{"Алибабек"}
	fmt.Println(p.Greet())

	k := Pm{"Костя"}
	fmt.Println(k.Great())

	w := NewWorker()
	if w != nil {
		w.Work()
	}

	var anyValue interface{}
	anyValue = "2312312"
	v, ok := anyValue.(string)
	if ok {
		fmt.Println("string: ", v)
	}

	var _ V = Mytype{}
	var _ V = &Mytype{}

	var dbConfig Config = Config{"DB config"}
	err := dbConfig.Save()
	if err != nil {
		fmt.Println(err)
	}

}
