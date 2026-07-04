package main

import "fmt"

type Paciente struct{
	Nome string
	Idade int
	NivelUrgencia int
	HorarioChegada string
}

var Pacientes = []Paciente{
	{Nome: "Ana", Idade: 68, NivelUrgencia: 2, HorarioChegada: "08:10"},
	{Nome: "Bruno", Idade: 45, NivelUrgencia: 1, HorarioChegada: "08:05"},
	{Nome: "Carla", Idade: 30, NivelUrgencia: 3, HorarioChegada: "07:50"},
	{Nome: "Diego", Idade: 72, NivelUrgencia: 1, HorarioChegada: "08:20"},
	{Nome: "Elisa", Idade: 55, NivelUrgencia: 2, HorarioChegada: "07:45"},
	{Nome: "Felipe", Idade: 22, NivelUrgencia: 3, HorarioChegada: "08:00"},
	{Nome: "Giovana", Idade: 61, NivelUrgencia: 1, HorarioChegada: "07:55"},
	{Nome: "Hugo", Idade: 40, NivelUrgencia: 2, HorarioChegada: "08:30"},
	{Nome: "Íris", Idade: 34, NivelUrgencia: 3, HorarioChegada: "08:15"},
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
func bubbleSortNivelUrgencia(arr []Paciente) {
	n := int(len(arr))
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].NivelUrgencia > arr[j+1].NivelUrgencia {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

func contagemPacientesCriticos(arr []Paciente){
	total := 0
	for _,paciente := range arr {
		if paciente.NivelUrgencia == 1 {
			total++
		}
	}

}

func main() {
	bubbleSortHorario(Pacientes)
	bubbleSortNivelUrgencia(Pacientes)
	

}
