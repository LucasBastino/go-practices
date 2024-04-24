-- Active: 1713822963476@@w0x.h.filess.io@3307@goProjects_cowboydish
USE goProjects_cowboydish;

CREATE TABLE team(
	idTeam INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    points INT
);

CREATE TABLE player(
	idPlayer INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    age INT,
    idTeam INT,
    FOREIGN KEY (idTeam) REFERENCES team(idTeam)
);

SELECT * FROM team;
SELECT * FROM player;

SELECT player.name, player.age, team.name FROM player
 INNER JOIN team 
 ON player.idTeam = team.idTeam
 WHERE player.idTeam = 2

