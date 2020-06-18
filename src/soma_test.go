package main

import "testing"

func TestSoma(t *testing.T) {
	
	t.Run("Somando Certo", func(t *testing.T){
		testes := []struct {
		num1, num2, esperado float64
		} {
			{5, 5, 10},
			{5.5, 5, 10.5},
			{1, 4, 5},
			{2.5, 2.5, 5},
			{1.3333, 5.6667, 7},
			{0, 0, 0},
		}

		for _, teste := range testes {
			
			soma := Soma(teste.num1, teste.num2)
			
			if soma != teste.esperado {
				t.Errorf("Soma(%g, %g) == %g; esperado %g", teste.num1, teste.num2, soma, teste.esperado)
			}
		}
	})
	
	t.Run("Somando Errado", func(t *testing.T){
		testes := []struct {
		num1, num2, esperado float64
		} {
			{5, 5, 10.1},
			{5.5, 5, 10.4},
			{1, 4, 6},
			{2.5, 2.5, 4.9999999},
			{1.3333, 5.6667, 7.000001},
			{0, 0, 1},
		}

		for _, teste := range testes {
			
			soma := Soma(teste.num1, teste.num2)
			
			if soma == teste.esperado {
				t.Errorf("Era esperado um erro, porém não ocorreu; [num1: %g, num2: %g, resultado: %g, esperado: %g]", teste.num1, teste.num2, soma, teste.esperado)
			}
		}
	})
}