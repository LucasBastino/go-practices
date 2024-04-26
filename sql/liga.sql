CREATE DATABASE liga;
USE liga;
USE goProjects_cowboydish;

CREATE TABLE equipos(
idEquipo INT PRIMARY KEY AUTO_INCREMENT,
nombre VARCHAR(45),
zona VARCHAR(1),
puntos INT
);

CREATE TABLE jugadores(
idJugador INT PRIMARY KEY AUTO_INCREMENT,
nombre VARCHAR(45),
edad INT,
posicion VARCHAR(3)
);

CREATE TABLE trabajadores(
idTrabajador INT PRIMARY KEY AUTO_INCREMENT,
nombre VARCHAR(45)
);

CREATE TABLE puestos(
idPuesto INT PRIMARY KEY AUTO_INCREMENT,
puesto VARCHAR(45),
sueldo INT
);

CREATE TABLE equipos_jugador(
idEquipo INT,
FOREIGN KEY (idEquipo) REFERENCES equipos(idEquipo),
idJugador INT,
FOREIGN KEY (idJugador) REFERENCES jugadores(idJugador)
);

CREATE TABLE equipos_trabajador(
idEquipo INT,
FOREIGN KEY (idEquipo) REFERENCES equipos(idEquipo),
idTrabajador INT,
FOREIGN KEY (idTrabajador) REFERENCES trabajadores(idTrabajador)
);

CREATE TABLE puestos_trabajador(
idPuesto INT,
FOREIGN KEY (idPuesto) REFERENCES puestos(idPuesto),
idTrabajador INT,
FOREIGN KEY (idTrabajador) REFERENCES trabajadores(idTrabajador)
);

INSERT INTO equipos (nombre, zona, puntos) VALUES ('Chacarita', 'B', 47);
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Rivadavia', 'B', 46);
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Maipu', 'B', 45);
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Quilmes', 'B', 38);
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Rafaela', 'B', 34);
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Ferro', 'B', 33);
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Brown A.', 'B', 32);
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Mitre', 'B', 32);
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Riestra', 'B', 31);
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Madryn', 'B', 30);

INSERT INTO equipos (nombre, zona, puntos) VALUES ('Almirante', 'A', 41 );
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Agropecuario', 'A', 39 );
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Moron', 'A', 38 );
INSERT INTO equipos (nombre, zona, puntos) VALUES ('San Martin T', 'A', 37 );
INSERT INTO equipos (nombre, zona, puntos) VALUES ('CADU', 'A', 37 );
INSERT INTO equipos (nombre, zona, puntos) VALUES ('San Martin SJ', 'A', 37 );
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Chicago', 'A', 36 );
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Gimnasia M', 'A', 35 );
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Temperley', 'A', 33 );
INSERT INTO equipos (nombre, zona, puntos) VALUES ('Estudiantes RC', 'A', 33 );

SELECT nombre AS Equipo, zona AS Zona, puntos AS Puntos FROM equipos ORDER BY puntos DESC;
DELETE FROM equipos;

INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Cristian Correa', 32, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Andres Zanini', 26, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Ricardo Blanco', 33, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Luciano Gimenez', 23, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Maxi Gagliardo', 40, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Luciano Abecasis', 33, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Ezequiel Ham', 29, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Alex Arce', 28, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Juan Pablo Cozzani', 24, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Nicolas Agorreca', 33, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Rubens Sambueza', 39, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Luciano Herrera', 27, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Milton Alvarez', 34, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Ivan Erquiaga', 24, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Ivan Colman', 28, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Federico Anselmo', 29, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Marcos Peano', 24, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Fabricio Fontanini', 33, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Matías Valdivia', 26, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Claudio Bieler', 39, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Luciano Jachfe', 21, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Hernan Grana', 38, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Cristian Erbes', 33, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Jonathan Herrera', 31, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Horacio Ramírez', 39, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Rodrigo Sayavedra', 27, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Hernán Da Campo', 28, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Agustín Lavezzi', 27, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Kevin Larrea', 27, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Gonzalo Soto', 33, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Santiago Rosales', 28, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Germán Mayenfisch', 30, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Ignacio Arce', 31, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Nicolás Dematei', 35, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Milton Céliz', 30, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Walter Acuña', 31, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Nicolás Temperini', 28, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Jorge Zules', 32, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Federico Recalde', 26, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Lucas González', 25, 'DEL');

INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Ramiro Martínez', 32, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Agustín Dattola', 24, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Santiago Vera', 24, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Martín Batallini', 31, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Juan Rago', 34, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Diego Mondino', 28, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Damián Lemos', 34, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Mauricio Asenjo', 28, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Bruno Galván', 29, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Nicolás Henry', 24, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Gonzalo Berterame', 27, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Ezequiel Rescaldani', 31, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Darío Sand', 35, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Franco Meritello', 27, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES (' Brian Andrada', 26, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Mateo Acosta', 30, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Mauricio Aquino', 29, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Facundo Laumann', 33, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Matías Nizzo', 34, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Javier Velázquez', 39, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Mariano Monllor', 34, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Leonel Bontempo', 30, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Alexis Vega', 30, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Matías Donato', 24, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Daniel Monllor', 38, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Juan Fedorco', 22, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Agustín Curruhinca', 23, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Ezequiel Naya', 21, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Brian Olivera', 29, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Matías Recalde', 26, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Rodrigo Castro', 30, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Walter Herrera', 25, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Matías Fidel Castro', 35, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Nicolás Demartini', 23,'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Luciano Nieto', 32, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Franco Ayunta', 20, 'DEL');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Franco Petroli', 25, 'ARQ');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Facundo Castet', 24, 'DEF');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Mateo Bajamich', 23, 'MED');
INSERT INTO jugadores (nombre, edad, posicion) VALUES ('Luis Silba', 33, 'DEL');

SELECT * FROM jugadores;

INSERT INTO trabajadores (nombre) VALUES ('Anibal Biggeri');
INSERT INTO trabajadores (nombre) VALUES ('Matias Coronel');
INSERT INTO trabajadores (nombre) VALUES ('Alfredo Berti');
INSERT INTO trabajadores (nombre) VALUES ('Fabricio Garquiaga');
INSERT INTO trabajadores (nombre) VALUES ('Luis Garcia');
INSERT INTO trabajadores (nombre) VALUES ('Milton Pichon');
INSERT INTO trabajadores (nombre) VALUES ('Mario Sciacqua');
INSERT INTO trabajadores (nombre) VALUES ('Jose Alfontineti');
INSERT INTO trabajadores (nombre) VALUES ('Ezequiel Medrán');
INSERT INTO trabajadores (nombre) VALUES ('Luciano Di Fronzo');
INSERT INTO trabajadores (nombre) VALUES ('Jorge Cordon');
INSERT INTO trabajadores (nombre) VALUES ('Lautaro Malguin');
INSERT INTO trabajadores (nombre) VALUES ('Pablo Vico');
INSERT INTO trabajadores (nombre) VALUES ('Alexis Darsio');
INSERT INTO trabajadores (nombre) VALUES ('Alfredo Grelak');
INSERT INTO trabajadores (nombre) VALUES ('Willian Yosani');
INSERT INTO trabajadores (nombre) VALUES ('Walter Marchesi');
INSERT INTO trabajadores (nombre) VALUES ('Pedro Millagran');
INSERT INTO trabajadores (nombre) VALUES ('Andrés Yllana');
INSERT INTO trabajadores (nombre) VALUES ('Walter Antinas');

INSERT INTO trabajadores (nombre) VALUES ('Darío Franco');
INSERT INTO trabajadores (nombre) VALUES ('Javier Astupino');
INSERT INTO trabajadores (nombre) VALUES ('Gabriel Gómez');
INSERT INTO trabajadores (nombre) VALUES ('Regino Chivaglio');
INSERT INTO trabajadores (nombre) VALUES ('Fabián Nardozza');
INSERT INTO trabajadores (nombre) VALUES ('Marcelo Tresto');
INSERT INTO trabajadores (nombre) VALUES ('Alexis Ferrero');
INSERT INTO trabajadores (nombre) VALUES ('Pablo Aiguro');
INSERT INTO trabajadores (nombre) VALUES ('Santiago Davio');
INSERT INTO trabajadores (nombre) VALUES ('Gino Marcone');
INSERT INTO trabajadores (nombre) VALUES ('César Monasterio');
INSERT INTO trabajadores (nombre) VALUES ('Lucas Fuenzalida');
INSERT INTO trabajadores (nombre) VALUES ('Andrés Montenegro');
INSERT INTO trabajadores (nombre) VALUES ('Kevin Hoyos');
INSERT INTO trabajadores (nombre) VALUES ('Andrés Iglesias');
INSERT INTO trabajadores (nombre) VALUES ('Lisandro Quimares');
INSERT INTO trabajadores (nombre) VALUES ('José María Bianco');
INSERT INTO trabajadores (nombre) VALUES ('Renato Lienes');
INSERT INTO trabajadores (nombre) VALUES ('Iván Delfino');
INSERT INTO trabajadores (nombre) VALUES ('Alfonso Malcorra');

INSERT INTO puestos (puesto, sueldo) VALUES ('Director Tecnico', 2000000);
INSERT INTO puestos (puesto, sueldo) VALUES ('Preparador Fisico', 500000);

INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (1, 1);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (1, 2);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (1, 3);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (1, 4);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (2, 5);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (2, 6);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (2, 7);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (2, 8);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (3, 9);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (3, 10);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (3, 11);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (3, 12);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (4, 13);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (4, 14);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (4, 15);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (4, 16);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (5, 17);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (5, 18);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (5, 19);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (5, 20);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (6, 21);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (6, 22);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (6, 23);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (6, 24);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (7, 25);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (7, 26);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (7, 27);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (7, 28);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (8, 29);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (8, 30);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (8, 31);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (8, 32);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (9, 33);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (9, 34);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (9, 35);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (9, 36);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (10, 37);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (10, 38);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (10, 39);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (10, 40);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (11, 41);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (11, 42);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (11, 43);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (11, 44);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (12, 45);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (12, 46);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (12, 47);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (12, 48);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (13, 49);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (13, 50);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (13, 51);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (13, 52);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (14, 53);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (14, 54);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (14, 55);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (14, 56);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (15, 57);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (15, 58);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (15, 59);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (15, 60);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (16, 61);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (16, 62);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (16, 63);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (16, 64);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (17, 65);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (17, 66);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (17, 67);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (17, 68);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (18, 69);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (18, 70);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (18, 71);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (18, 72);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (19, 73);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (19, 74);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (19, 75);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (19, 76);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (20, 77);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (20, 78);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (20, 79);
INSERT INTO equipos_jugador (idEquipo, idJugador) VALUES (20, 80);

INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (1, 1);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (1, 2);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (2, 3);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (2, 4);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (3, 5);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (3, 6);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (4, 7);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (4, 8);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (5, 9);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (5, 10);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (6, 11);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (6, 12);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (7, 13);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (7, 14);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (8, 15);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (8, 16);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (9, 17);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (9, 18);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (10, 19);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (10, 20);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (11, 21);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (11, 22);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (12, 23);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (12, 24);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (13, 25);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (13, 26);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (14, 27);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (14, 28);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (15, 29);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (15, 30);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (16, 31);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (16, 32);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (17, 33);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (17, 34);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (18, 35);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (18, 36);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (19, 37);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (19, 38);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (20, 39);
INSERT INTO equipos_trabajador (idEquipo, idTrabajador) VALUES (20, 40);

INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 1);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 2);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 3);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 4);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 5);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 6);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 7);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 8);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 9);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 10);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 11);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 12);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 13);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 14);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 15);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 16);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 17);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 18);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 19);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 20);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 21);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 22);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 23);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 24);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 25);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 26);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 27);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 28);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 29);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 30);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 31);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 32);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 33);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 34);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 35);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 36);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 37);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 38);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (1, 39);
INSERT INTO puestos_trabajador (idPuesto, idTrabajador) VALUES (2, 40);

