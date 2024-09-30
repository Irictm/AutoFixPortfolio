/*
Receipt: Representa el recibo que contiene el detalle de los montos y
el monto total a cobrar por la reparacion

  - Id: Identificador del recibo

  - OperationsAmount: Monto base asociado a la suma de los costos individuales
    de cada operacion realizada

  - RechargeAmount: Monto asociado a las recargas realizadas por distintos criterios de la empresa,
    en particular contiene el monto de la recarga por antiguedad, por kilometraje y por retraso en
    la recogida del vehiculo.

  - DiscountAmount: Monto asociado a los descuentos realizados por distintos criterios de la empresa,
    en particular contiene el monto del descuento por numero de reparaciones realizadas en el año, por
    dia-hora de ingreso al taller y por bonos emitidos a la marca del vehiculo.

  - IvaAmount: Monto asociado a la aplicacion del IVA sobre el costo de la reparacion,
    calculado como (OperationsAmount - DiscountAmount + RechargeAmount) * IVA

  - TotalAmount: Monto total de la reparacion considerando el monto del IVA, calculado tal que
    (OperationsAmount - DiscountAmount + RechargeAmount + IvaAmount)

  - BonusConsumed: Booleano que indica si al momento del calculo se debe considerar el
    bonus como ya consultado o no, es decir, si ya se incluyo la peticion del bono en un calculo previo.
*/
package data

type Receipt struct {
	Id               int64
	OperationsAmount int32
	RechargeAmount   int32
	DiscountAmount   int32
	IvaAmount        int32
	TotalAmount      int32
	BonusConsumed    bool
}
