package main

import "fmt"

type Alerta interface {
	Actualizar(mensaje string)
}

type AlertaMovil struct{}
type AlertaWeb struct{}

func (a *AlertaMovil) Actualizar(mensaje string) {
	fmt.Println("Alerta móvil recibida:", mensaje)
}
func (a *AlertaWeb) Actualizar(mensaje string) {
	fmt.Println("Alerta web recibida:", mensaje)
}

type EstacionMeteorologica struct {
	subscribers []Alerta
}

func (e *EstacionMeteorologica) Registrar(alert Alerta) {
	e.subscribers = append(e.subscribers, alert)
}
func (e *EstacionMeteorologica) Eliminar(alert Alerta) {
	for i, s := range e.subscribers {
		if s == alert {
			e.subscribers = append(e.subscribers[:i], e.subscribers[i+1:]...)
			break
		}
	}
}
func (e *EstacionMeteorologica) Notificar(mensaje string) {
	for _, s := range e.subscribers {
		s.Actualizar(mensaje)
	}
}

func main() {
	estacion := &EstacionMeteorologica{}
	movil := &AlertaMovil{}
	web := &AlertaWeb{}

	estacion.Registrar(movil)
	estacion.Registrar(web)

	estacion.Notificar("Tormenta eléctrica en la zona")
	estacion.Notificar("Temperatura extrema")

	estacion.Eliminar(movil)
	estacion.Notificar("Lluvia intensa")
}
