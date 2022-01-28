CREATE TABLE beer (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255),
    beer_type VARCHAR(255)
);

CREATE TABLE coffee (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO beer VALUES (1, 'Brahma', 'NORMAL');
INSERT INTO beer VALUES (2, 'Petra', 'NORMAL');
INSERT INTO beer VALUES (3, 'Imperio', 'NORMAL');
INSERT INTO beer VALUES (4, 'Brahma Duplo Malte', 'MALTE');


INSERT INTO coffee VALUES (1, 'Café gelado');
INSERT INTO coffee VALUES (2, 'Café Munddo novo coado');
INSERT INTO coffee VALUES (3, 'Café Catuai prensa francesa');