/*
Repair: Representa una reparacion realizada sobre un vehiculo, especificando
las fechas y horas relevantes a esta.

  - Id: Identificador de la reparacion

  - DateOfAdmission: Fecha y hora de admision del vehiculo al local, siguiendo RFC 3339 (Timestamp)

  - DateOfRelease: Fecha y hora de liberacion del vehiculo, quedando este diponible para recoger, siguiendo RFC 3339 (Timestamp)

  - DateOfPickUp: Fech y horaa de recogida del vehiculo del local, siguiendo RFC 3339 (Timestamp)

  - Id_receipt: Identificador del recibo con los montos asociados a la reparacion

  - Id_vehicle: Identificador del vehiculo asociado a la reparacion
*/
package data

import "time"

type Repair struct {
	Id              int64
	DateOfAdmission time.Time
	DateOfRelease   time.Time
	DateOfPickUp    time.Time
	Id_receipt      int64
	Id_vehicle      int64
}
