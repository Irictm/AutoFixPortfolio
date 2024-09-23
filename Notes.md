# Notes
- Agregar mas contexto a err retornados por calculo del costo total de recibo
- Reescribir Save de repositories para no crear nuevo objeto y retornar mismo objeto con id asignado
- Considerar cambiar color fuente de errores a color Rojo (paquete slog?)
- Refactorizacion del paquete Tariffs (4 tablas -> 4 CRUDS?, dos tipos de tarifas regular y criteria de intervalo?)
- Considerar tabla extra para tipos de operacion (id, nombre)
- Considerar envio de tarifas (enviar individualmente?, empaquetar como tabla?, csv?)
- Considerar eliminar dependencia "net/http" y ocupar constantes con nombre para codigos http usados
- Considerar agregar personal al sistema (usuarios?) (No considerado inicialmente por
enunciado, considerado ahora como ejercicio practico)


# TO DO DataBase
- Diseñar a traves de Triggers y Funciones metodo para guardar Backlog de
  acciones realizadas por usuarios


# TO DO Backend
- Implementar funcionalidades clave de cada clase
- Implementar sistema de usuarios, registro y logging
- Agregar medidas de seguridad basicas (Anti SQL injection, JWT, etc)
- Creacion de Tests Unitarios
- Documentacion
- Implementacion de Reportes sobre datos del sistema