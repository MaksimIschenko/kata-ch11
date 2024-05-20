package main

import (
	"fmt"
)

type TV struct {
    status bool
	model string
}

type TVer interface {
	switchOn() bool
	switchOFF() bool
	GetStatus() bool
	GetModel() string
}

type Samsunger interface{
	TVer
	SamsungHub() string
}

type LGer interface{
	TVer
	LGHub() string
}

func (tv TV) switchOn() bool {
	tv.status = true
	return tv.status
}

func (tv TV) switchOFF() bool {
	tv.status = false
	return tv.status
}

func (tv TV) GetStatus() bool {
	return tv.status
}

func (tv TV) GetModel() string {
	return tv.model
}

func (tv TV) SamsungHub() string {
	return "SamsungHub"
}

func (tv TV) LGHub() string {
	return "LGHub"
}

func main() {
	tv := &TV{
		status: true,
		model: "Samsung XL-100500",
	}
	fmt.Println(tv.GetStatus()) 	// true
    fmt.Println(tv.GetModel())  	// Samsung XL-100500
    fmt.Println(tv.SamsungHub()) 	// SamsungHub
    fmt.Println(tv.GetStatus()) 	// false
    fmt.Println(tv.switchOn()) 		// true
    fmt.Println(tv.GetStatus()) 	// true
    fmt.Println(tv.switchOFF()) 	// true
    fmt.Println(tv.GetStatus()) 	// false
}