/*
Vehicle: Representa un vehiculo en el sistema.

  - Id: Identificador del vehiculo

  - Patent: Patente del vehiculo de la forma "ABCD12"

  - Brand: Marca del vehiculo (Toyota, Ford, etc)

  - Model: Modelo del vehiculo

  - Type: Tipo de vehiculo (SUV, Furgoneta, etc)

  - FabricationDate: Fecha de fabricacion del auto, siguiendo RFC 3339 (Timestamp) solo ocupando año, mes y dia.

  - MotorType: Tipo de motor del vehiculo (Gasolina, Electrico, etc)

  - Seats: Numero de asientos del vehiculo

  - Mileage: Kilometraje del auto reportado en el momento de entrada al taller
*/
package data

import "time"

type Vehicle struct {
	Id              int64
	Patent          string
	Brand           string
	Model           string
	Type            string
	FabricationDate time.Time
	MotorType       string
	Seats           uint8
	Mileage         int32
}
