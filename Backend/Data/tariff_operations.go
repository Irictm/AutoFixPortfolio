/*
Tariff Operations: Representa la tarifa asociada con una cierta operacion realizada

  - Id: Identificador de la tarifa

  - MotorType: Tipo de motor del vehiculo al cual se aplica la tarifa

  - Id_operation_type: Identificador del tipo de operacion al cual se aplica la tarifa

  - Value: Monto asociado a la operacion aplicada si el vehiculo al cual se aplica la operacion
    tiene un motor del tipo asociado a la tarifa y la operacion coincide con el tipo de la tarifa.
*/
package data

type TariffOperations struct {
	Id                uint32
	MotorType         string
	Id_operation_type int32
	Value             int32
}
