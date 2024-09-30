# Notes
- Considerar reestructurar receipt para guardar montos individuales y booleano para indicar si fue calculado el monto total
- Cambiar folders de PascalCase a snake_case
- Considerar envio de tarifas (enviar individualmente?, empaquetar como tabla?, csv?)
- Manejar infinity como -1 y -infinity como -2 desde el frontend y enviar MaxInt() o MinInt() respectivamente a PostgreSQL
- Considerar cambiar color fuente de errores a color Rojo (paquete slog?)
- Considerar eliminar dependencia "net/http" y ocupar constantes con nombre para codigos http usados
- Considerar agregar personal al sistema (usuarios?) (No considerado inicialmente por
enunciado, considerado ahora como ejercicio practico)


# TO DO DataBase
- Diseñar a traves de Triggers y Funciones metodo para guardar Backlog de
  acciones realizadas por usuarios


# TO DO Backend
- Creacion de Tests Unitarios
- Documentacion de metodos relevantes
- Implementar sistema de usuarios, registro y logging
- Agregar medidas de seguridad basicas (Anti SQL injection, JWT, etc)
- Implementacion de Reportes sobre datos del sistema