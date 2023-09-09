package models

type Man struct {
	age       int
	name      string
	breathing bool
	working   bool
	eating    bool
	alive     bool
}

func (this *Man) Eat() {
	this.eating = true
}

func (this *Man) Breathe() {
	this.breathing = true
}

func (this *Man) Work() {
	this.working = true
}

func (this *Man) Genre() string {
	return "Male"
}

func (this *Man) IsAlive() {
	this.alive = true
}
