package main

import (
	"testing"
	"time"
)

//Test que chequea que la palabra sea fecha, hora o sino devuelve error
func TestChequearPalabraFecha(t *testing.T) {
	palabra := "fecha"

	resultado := chequearPalabra(palabra)
	esperado := string(time.Now().Format("02/01/2006"))
	if resultado !=  esperado{
		t.Errorf("Se esperaba %s, se obtuvo %s", esperado, resultado)
	}
}

func TestChequearPalabraHora(t *testing.T) {
	palabra := "hora"

	resultado := chequearPalabra(palabra)
	esperado := string(time.Now().Format("15:04:05"))
	if resultado !=  esperado{
		t.Error("Se esperaba true, se obtuvo ", resultado)
	}
}

func TestChequearPalabraError(t *testing.T) {
	palabra := "error"

	resultado := chequearPalabra(palabra)
	esperado := "Error"
	if resultado !=  esperado{
		t.Error("Se esperaba false, se obtuvo ", resultado)
	}
}

func TestDevolverHora(t *testing.T) {
	resultado := devolverHora()
	esperado := string(time.Now().Format("15:04:05"))

	if resultado != esperado {
		t.Errorf("Se esperaba %s, se obtuvo %s", esperado ,resultado)
	}
}

func TestDevolverFecha(t *testing.T) {
	resultado := devolverFecha()
	esperado := string(time.Now().Format("02/01/2006"))

	if resultado != esperado {
		t.Errorf("Se esperaba %s, se obtuvo %s", esperado ,resultado)
	}
}

func TestDevolverError(t *testing.T) {
	resultado := devolverError()
	esperado := "Error"

	if resultado != esperado {
		t.Errorf("Se esperaba %s, se obtuvo %s", esperado ,resultado)
	}
}

//Test que chequea que ChequearPalabra funcione con cualquier case
func TestChequearPalabraCase(t *testing.T) {
	palabra := "FeChA"

	resultado := chequearPalabra(palabra)
	esperado := string(time.Now().Format("02/01/2006"))
	if resultado !=  esperado{
		t.Errorf("Se esperaba %s, se obtuvo %s", esperado, resultado)
	}
}