package main

import (
	"testing"
)

func TestEsPalindroma(t *testing.T) {
	//Test 1
	if !esPalindroma("ana") {
		t.Error("Test 1: ana es palindroma")
	}
	//Test 2
	if esPalindroma("hola") {
		t.Error("Test 2: hola no es palindroma")
	}
	//Test Mayusculas 
	if !esPalindroma("oso") {
		t.Error("Test 3: oso es palindroma")
	}
	//Test 4
	if esPalindroma("palabra") {
		t.Error("Test 4: palabra no es palindroma")
	}

	//Test 5
	if !esPalindroma("r") {
		t.Error("Test 5: a es palindroma")
	}
}

func TestGetWords(t *testing.T) {
	//Test 1
	if getWords("test.txt")[0] != "ana" {
		t.Error("Test 1: ana es la primera palabra")
	}
	//Test 2
	if getWords("test.txt")[1] != "oso" {
		t.Error("Test 2: oso es la segunda palabra")
	}
	//Test 3
	if getWords("test.txt")[2] != "hola" {
		t.Error("Test 3: hola es la tercera palabra")
	}
	//Test 4
	if getWords("test.txt")[3] != "palabra" {
		t.Error("Test 4: palabra es la cuarta palabra")
	}
}

func TestGetPalindromes(t *testing.T) {
	//Test 1
	if getPalindromes(getWords("test.txt"))[0] != "ana" {
		t.Error("Test 1: ana es el primer palindromo")
	}
	//Test 2
	if getPalindromes(getWords("test.txt"))[1] != "oso" {
		t.Error("Test 2: oso es el segundo palindromo")
	}
	//Test 3
	if len(getPalindromes(getWords("test.txt"))) != 2 {
		t.Error("Test 3: hay 2 palindromos")
	}
}