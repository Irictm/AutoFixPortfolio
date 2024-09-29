/*
Bonus: Representa una cierta cantidad de bonuses emitidos por una marca de vehiculos
para ser consumidos por la siguiente reparacion realizada a un vehiculo de la
correspondiente marca.

  - Id: Identificador del bonus

  - Brand: Marca de vehiculos que emite el bonus

  - Remaining: Cantidad de bonuses de la marca disponibles, si llega a 0 el bonus deja de aplicarse

  - Amount: Monto que aporta el bonus a la reparacion
*/
package data

type Bonus struct {
	Id        uint32
	Brand     string
	Remaining uint8
	Amount    int32
}
