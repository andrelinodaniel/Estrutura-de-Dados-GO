package main

import "fmt"

type Paciente struct{
	Nome string
	Idade int
	NivelUgercia int
	HorarioChegada string
}

var Pacientes = []Paciente{
	{Nome: "Ana", Idade: 68, NivelUgercia: 2, HorarioChegada: "08:10"},
	{Nome: "Bruno", Idade: 45, NivelUgercia: 1, HorarioChegada: "08:05"},
	{Nome: "Carla", Idade: 30, NivelUgercia: 3, HorarioChegada: "07:50"},
	{Nome: "Diego", Idade: 72, NivelUgercia: 1, HorarioChegada: "08:20"},
	{Nome: "Elisa", Idade: 55, NivelUgercia: 2, HorarioChegada: "07:45"},
	{Nome: "Felipe", Idade: 22, NivelUgercia: 3, HorarioChegada: "08:00"},
	{Nome: "Giovana", Idade: 61, NivelUgercia: 1, HorarioChegada: "07:55"},
	{Nome: "Hugo", Idade: 40, NivelUgercia: 2, HorarioChegada: "08:30"},
	{Nome: "Íris", Idade: 34, NivelUgercia: 3, HorarioChegada: "08:15"},
}
func bubbleSortHorario(arr []Paciente) {
	n := int(len(arr))
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].HorarioChegada > arr[j+1].HorarioChegada {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}
func bubbleSortHorario(arr []Paciente) {
	n := int(len(arr))
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].NivelUrgencia > arr[j+1].NivelUrgencia {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}



func main() {
	bubbleSortHorario(Pacientes)
	for

}
