/*
Tariff Antiquity: Representa la tarifa asociada con el recargo realizado por la
antiguedad del vehiculo

  - Id: Identificador de la tarifa

  - VehicleType: Tipo de vehiculo al cual se aplica la tarifa

  - Bottom: Valor inferior del rango de antiguedad del vehiculo (Contado en años) bajo el cual se aplica la tarifa

  - Top: Valor superior del rango de antiguedad del vehiculo (Contado en años) bajo el cual se aplica la tarifa

  - Value: Porcentaje de recargo de la tarifa aplicado si el vehiculo cuenta con el tipo de vehiculo
    asociado, su antiguedad es mayor o igual al valor de Bottom y menor o igual al valor de Top.
*/
package data

type TariffAntiquity struct {
	Id          uint32
	VehicleType string
	Bottom      int32
	Top         int32
	Value       float64
}
