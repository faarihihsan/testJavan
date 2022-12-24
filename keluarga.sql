-- language : postgreSQL
CREATE TABLE keluarga
(
    id     SERIAL,
    nama   VARCHAR(255),
    parent INT ,

    PRIMARY KEY (id),
    FOREIGN KEY (parent) REFERENCES keluarga (id) ON DELETE SET NULL
);

CREATE TABLE aset
(
    id   SERIAL,
    nama VARCHAR(255),
    price int,

    PRIMARY KEY (id)
);

CREATE TABLE keluarga_aset
(
    id          SERIAL,
    id_keluarga INT,
    id_aset     INT,

    PRIMARY KEY (id),
    FOREIGN KEY (id_keluarga) REFERENCES keluarga (id) ON DELETE CASCADE,
    FOREIGN KEY (id_aset) REFERENCES aset (id) ON DELETE CASCADE
);

INSERT INTO keluarga(id, nama, parent)
VALUES (1, 'Bani', null),
       (2, 'Budi', 1),
       (3, 'Nida', 1),
       (4, 'Andi', 1),
       (5, 'Sigit', 1),
       (6, 'Hari', 2),
       (7, 'Siti', 2),
       (8, 'Bila', 3),
       (9, 'Lesti', 3),
       (10, 'Diki', 4),
       (11, 'Doni', 5),
       (12, 'Toni', 5);


INSERT INTO aset(id, nama)
VALUES (1, 'Samsung Universe 9'),
       (2, 'Samsung Galaxy Book'),
       (3, 'iPhone 9'),
       (4, 'iPhone X'),
       (5, 'Huawei P30');


INSERT INTO keluarga_aset(id, id_keluarga, id_aset)
VALUES (1, 2, 1),
       (2, 2, 2),
       (3, 6, 3),
       (4, 7, 4),
       (5, 3, 5),
       (6, 8, 1),
       (7, 9, 5),
       (8, 9, 4),
       (9, 4, 1),
       (10, 10, 2),
       (11, 5, 5),
       (12, 11, 4);