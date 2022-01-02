package main

import (
	"fmt"

	"../gadget"
	// "../mypkg"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder) //Утверждение типа используется для перехода к значению TapeRecorder.
	if ok {
		recorder.Record() // Если исходным типом был тип TapeRecorder, для значения вызывается метод Record.
	} else {
		fmt.Println("Player was not a TapeRecorder") //В противном случае выда-ется сообщение об ошибке утверждения типа.
	}
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})

	//test mypkg с интерфейсом
	// var value mypkg.MyInterface
	// value := mypkg.MyType(5)
	// value.MethodWithoutParameters()
	// value.MethodWithParameter(127.3)
	// fmt.Println(value.MethodWithReturnValue())
}
