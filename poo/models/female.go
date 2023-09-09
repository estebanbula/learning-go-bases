package models

type Female struct {
	Man
}

func (this *Female) Eat() {
	this.eating = true
}

func (this *Female) Breathe() {
	this.breathing = true
}

func (this *Female) Work() {
	this.working = true
}

func (this *Female) Genre() string {
	return "Female"
}

func (this *Female) IsAlive() {
	this.alive = true
}
