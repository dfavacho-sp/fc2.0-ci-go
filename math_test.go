package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(20, 10)
	if total != 30 {
		t.Errorf("Resultado da soma %d. está incorreto, Esperado: %d", total, 30)
	}
	
}