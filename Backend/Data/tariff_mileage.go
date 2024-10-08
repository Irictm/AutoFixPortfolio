/*
Tariff Mileage: Representa la tarifa asociada con el recargo realizado por el
kilometraje del vehiculo

  - Id: Identificador de la tarifa

  - VehicleType: Tipo de vehiculo al cual se aplica la tarifa

  - Bottom: Valor inferior del rango de kilometraje del vehiculo (Contado en Kilometros) bajo el cual se aplica la tarifa

  - Top: Valor superior del rango de kilometraje del vehiculo (Contado en Kilometros) bajo el cual se aplica la tarifa

  - Value: Porcentaje de recargo de la tarifa aplicado si el vehiculo cuenta con el tipo de vehiculo
    asociado, su kilometraje es mayor o igual al valor de Bottom y menor o igual al valor de Top.
*/
package data

type TariffMileage struct {
	Id          int64
	VehicleType string
	Bottom      float64
	Top         float64
	Value       float64
}
