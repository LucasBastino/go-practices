CREATE DATABASE liga;
USE liga;
CREATE TABLE afiliado(
    idAfiliado INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(45),
    edad INT
);
CREATE TABLE equipo(
idEquipo INT PRIMARY KEY AUTO_INCREMENT,
nombre VARCHAR(45),
zona VARCHAR(1),
puntos INT
);
DELETE FROM equipos;
DELETE FROM jugadores;
DELETE FROM jugador;
DELETE FROM equipos;
CREATE TABLE puesto(
idPuesto INT PRIMARY KEY AUTO_INCREMENT,
nombre VARCHAR(45),
sueldo INT
);

CREATE TABLE jugador(
idJugador INT PRIMARY KEY AUTO_INCREMENT,
nombre VARCHAR(45),
edad INT,
posicion VARCHAR(3),
idEquipo INT,
FOREIGN KEY (idEquipo) REFERENCES equipo(idEquipo)
);


CREATE TABLE trabajador(
idTrabajador INT PRIMARY KEY AUTO_INCREMENT,
nombre VARCHAR(45),
idEquipo INT,
FOREIGN KEY (idEquipo) REFERENCES equipo(idEquipo),
idPuesto INT,
FOREIGN KEY (idPuesto) REFERENCES puesto(idPuesto)
);



INSERT INTO equipo (nombre, zona, puntos) VALUES ('Chacarita', 'B', 47);
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Rivadavia', 'B', 46);
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Maipu', 'B', 45);
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Quilmes', 'B', 38);
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Rafaela', 'B', 34);
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Ferro', 'B', 33);
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Brown A.', 'B', 32);
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Mitre', 'B', 32);
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Riestra', 'B', 31);
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Madryn', 'B', 30);

INSERT INTO equipo (nombre, zona, puntos) VALUES ('Almirante', 'A', 41 );
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Agropecuario', 'A', 39 );
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Moron', 'A', 38 );
INSERT INTO equipo (nombre, zona, puntos) VALUES ('San Martin T', 'A', 37 );
INSERT INTO equipo (nombre, zona, puntos) VALUES ('CADU', 'A', 37 );
INSERT INTO equipo (nombre, zona, puntos) VALUES ('San Martin SJ', 'A', 37 );
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Chicago', 'A', 36 );
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Gimnasia M', 'A', 35 );
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Temperley', 'A', 33 );
INSERT INTO equipo (nombre, zona, puntos) VALUES ('Estudiantes RC', 'A', 33 );

SELECT nombre AS Equipo, zona AS Zona, puntos AS Puntos FROM equipo ORDER BY puntos DESC;

INSERT INTO puesto (nombre, sueldo) VALUES ('Director Tecnico', 2000000);
INSERT INTO puesto (nombre, sueldo) VALUES ('Preparador Fisico', 500000);

INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (1, 'Cristian Correa', 32, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (1, 'Andres Zanini', 26, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (1, 'Ricardo Blanco', 33, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (1, 'Luciano Gimenez', 23, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (2, 'Maxi Gagliardo', 40, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (2, 'Luciano Abecasis', 33, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (2, 'Ezequiel Ham', 29, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (2, 'Alex Arce', 28, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (3, 'Juan Pablo Cozzani', 24, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (3, 'Nicolas Agorreca', 33, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (3, 'Rubens Sambueza', 39, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (3, 'Luciano Herrera', 27, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (4, 'Milton Alvarez', 34, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (4, 'Ivan Erquiaga', 24, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (4, 'Ivan Colman', 28, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (4, 'Federico Anselmo', 29, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (5, 'Marcos Peano', 24, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (5, 'Fabricio Fontanini', 33, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (5, 'Matías Valdivia', 26, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (5, 'Claudio Bieler', 39, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (6, 'Luciano Jachfe', 21, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (6, 'Hernan Grana', 38, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (6, 'Cristian Erbes', 33, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (6, 'Jonathan Herrera', 31, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (7, 'Horacio Ramírez', 39, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (7, 'Rodrigo Sayavedra', 27, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (7, 'Hernán Da Campo', 28, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (7, 'Agustín Lavezzi', 27, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (8, 'Kevin Larrea', 27, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (8, 'Gonzalo Soto', 33, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (8, 'Santiago Rosales', 28, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (8, 'Germán Mayenfisch', 30, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (9, 'Ignacio Arce', 31, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (9, 'Nicolás Dematei', 35, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (9, 'Milton Céliz', 30, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (9, 'Walter Acuña', 31, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (10, 'Nicolás Temperini', 28, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (10, 'Jorge Zules', 32, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (10, 'Federico Recalde', 26, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (10, 'Lucas González', 25, 'DEL');

INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (11, 'Ramiro Martínez', 32, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (11, 'Agustín Dattola', 24, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (11, 'Santiago Vera', 24, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (11, 'Martín Batallini', 31, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (12, 'Juan Rago', 34, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (12, 'Diego Mondino', 28, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (12, 'Damián Lemos', 34, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (12, 'Mauricio Asenjo', 28, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (13, 'Bruno Galván', 29, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (13, 'Nicolás Henry', 24, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (13, 'Gonzalo Berterame', 27, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (13, 'Ezequiel Rescaldani', 31, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (14, 'Darío Sand', 35, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (14, 'Franco Meritello', 27, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (14, ' Brian Andrada', 26, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (14, 'Mateo Acosta', 30, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (15, 'Mauricio Aquino', 29, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (15, 'Facundo Laumann', 33, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (15, 'Matías Nizzo', 34, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (15, 'Javier Velázquez', 39, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (16, 'Mariano Monllor', 34, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (16, 'Leonel Bontempo', 30, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (16, 'Alexis Vega', 30, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (16, 'Matías Donato', 24, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (17, 'Daniel Monllor', 38, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (17, 'Juan Fedorco', 22, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (17, 'Agustín Curruhinca', 23, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (17, 'Ezequiel Naya', 21, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (18, 'Brian Olivera', 29, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (18, 'Matías Recalde', 26, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (18, 'Rodrigo Castro', 30, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (18, 'Walter Herrera', 25, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (19, 'Matías Fidel Castro', 35, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (19, 'Nicolás Demartini', 23,'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (19, 'Luciano Nieto', 32, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (19, 'Franco Ayunta', 20, 'DEL');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (20, 'Franco Petroli', 25, 'ARQ');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (20, 'Facundo Castet', 24, 'DEF');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (20, 'Mateo Bajamich', 23, 'MED');
INSERT INTO jugador (idEquipo, nombre, edad, posicion) VALUES (20, 'Luis Silba', 33, 'DEL');

SELECT * FROM jugador;

INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Anibal Biggeri', 1, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Matias Coronel', 1, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Alfredo Berti', 2, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Fabricio Garquiaga', 2, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Luis Garcia', 3, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Milton Pichon', 3, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Mario Sciacqua', 4, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Jose Alfontineti', 4, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Ezequiel Medrán', 5, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Luciano Di Fronzo', 5, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Jorge Cordon', 6, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Lautaro Malguin', 6, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Pablo Vico', 7, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Alexis Darsio', 7, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Alfredo Grelak', 8, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Willian Yosani', 8, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Walter Marchesi', 9, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Pedro Millagran', 9, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Andrés Yllana', 10, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Walter Antinas', 10, 2);

INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Darío Franco', 11, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Javier Astupino', 11, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Gabriel Gómez', 12, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Regino Chivaglio', 12, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Fabián Nardozza', 13, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Marcelo Tresto', 13, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Alexis Ferrero', 14, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Pablo Aiguro', 14, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Santiago Davio', 15, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Gino Marcone', 15, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('César Monasterio', 16, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Lucas Fuenzalida', 16, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Andrés Montenegro', 17, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Kevin Hoyos', 17, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Andrés Iglesias', 18, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Lisandro Quimares', 18, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('José María Bianco', 19, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Renato Lienes', 19, 2);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Iván Delfino', 20, 1);
INSERT INTO trabajador (nombre, idEquipo, idPuesto) VALUES ('Alfonso Malcorra', 20, 2);

-- Todos los jugadores
SELECT jugador.nombre, jugador.posicion, equipo.nombre FROM jugador
INNER JOIN equipo ON jugador.idEquipo = equipo.idEquipo;

-- Todos los DT de equipos
SELECT trabajador.nombre, puesto.nombre, equipo.nombre, equipo.puntos FROM trabajador
INNER JOIN puesto ON trabajador.idPuesto = puesto.idPuesto
INNER JOIN equipo ON equipo.idEquipo = trabajador.idEquipo
WHERE puesto.nombre LIKE 'Dire%';

-- Todos los DT de los equipos con + de 35 puntos
SELECT trabajador.nombre AS 'Nombre', puesto.nombre AS 'Puesto', equipo.nombre AS 'Equipo', equipo.puntos AS 'Puntos' FROM trabajador
INNER JOIN puesto ON trabajador.idPuesto = puesto.idPuesto
INNER JOIN equipo ON equipo.idEquipo = trabajador.idEquipo
WHERE puesto.nombre LIKE 'Dire%' AND equipo.puntos > 35;

-- Todos los jugadores de Chacarita
SELECT jugador.nombre AS 'Jugador', jugador.posicion AS 'Posicion', equipo.nombre AS 'Equipo' FROM jugador
INNER JOIN equipo ON equipo.idEquipo = jugador.idEquipo
WHERE equipo.nombre = 'Chacarita';

-- Los mejores 4 equipos de la zona A
SELECT equipo.nombre AS 'Equipo', equipo.zona AS 'Zona', equipo.puntos AS 'Puntos' FROM equipo
WHERE equipo.zona = 'A' LIMIT 4;

-- Los mejores 4 equipos de la zona B
SELECT equipo.nombre AS 'Equipo', equipo.zona AS 'Zona', equipo.puntos AS 'Puntos' FROM equipo
WHERE equipo.zona = 'B' LIMIT 4;

-- Los equipos con mas de 35 puntos
SELECT equipo.nombre, equipo.zona, equipo.puntos FROM equipo
WHERE equipo.puntos > 35 ORDER BY equipo.puntos DESC;

--  El equipo con mayor puntaje
SELECT nombre, puntos FROM equipo
WHERE puntos = (SELECT MAX(puntos) FROM equipo);

-- El equipo con menor puntaje
SELECT nombre, puntos FROM equipo
WHERE puntos = (SELECT MIN(puntos) FROM equipo);

-- Los jugadores del equipo con mayor puntaje
SELECT jugador.nombre, jugador.posicion FROM jugador
INNER JOIN equipo ON equipo.idEquipo = jugador.idEquipo
WHERE puntos = (SELECT MAX(puntos) FROM equipo);

-- Los jugadores de los primeros 4 equipos de la tabla
SELECT jugador.nombre, posicion, equipo.nombre FROM jugador
INNER JOIN equipo ON equipo.idEquipo = jugador.idEquipo
ORDER BY puntos DESC LIMIT 16;

-- El promedio de puntos por equipo
SELECT AVG(puntos) as 'Promedio PTS' FROM equipo;

-- Equipos arriba del promedio de puntos
SELECT equipo.nombre, equipo.puntos FROM equipo
WHERE puntos > (SELECT AVG(puntos) FROM equipo);

-- Equipos abajo del promedio de puntos
SELECT nombre, puntos FROM equipo
WHERE puntos < (SELECT AVG(puntos) FROM equipo);




