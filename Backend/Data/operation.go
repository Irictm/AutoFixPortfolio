/*
Operation: Representa una sola operacion realizada sobre el vehiculo en proceso de reparacion,
se pueden entender como subprocesos de la reparacion.

  - Id:	Identificador de la operacion

  - Patent: Patente del vehiculo al que se aplica la operacion

  - Id_operation_type: Identificador del tipo de la operacion

  - Date: Fecha y hora de la operacion, siguiendo RFC 3339 (Timestamp)

  - Cost: Monto cobrado por la realizacion de la operacion.

  - Id_repair: Identificador de la reparacion en la que se realizo la operacion
*/
package data

import "time"

type Operation struct {
	Id                uint32
	Patent            string
	Id_operation_type uint32
	Date              time.Time
	Cost              int32
	Id_repair         uint32
}
