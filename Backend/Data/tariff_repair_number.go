/*
Tariff Repair Number: Representa la tarifa asociada con el descuento realizado por el
numero de reparaciones realizadas en el vehiculo el utimo año.

  - Id: Identificador de la tarifa

  - MotorType: Tipo de motor del vehiculo que se beneficia por la tarifa

  - Bottom: Valor inferior del rango de numero de reparaciones del vehiculo en el año pasado bajo el cual se aplica la tarifa

  - Top: Valor superior del rango de numero de reparaciones del vehiculo en el año pasado bajo el cual se aplica la tarifa

  - Value: Porcentaje de descuento de la tarifa aplicado si el vehiculo cuenta con el tipo de motor
    asociado y si el numero de reparaciones realizadas el año anterior sobre el vehiculo
    es mayor o igual al valor de Bottom y menor o igual al valor de Top.
*/
package data

type TariffRepairNumber struct {
	Id        int64
	MotorType string
	Bottom    float64
	Top       float64
	Value     float64
}
