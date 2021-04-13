CREATE TYPE weapon_type AS ENUM ('spear', 'bow', 'axe', 'dagger', 'staff', 'melee', 'sword', 'katana', 'harp', 'gun');

CREATE TYPE weapon_rarity AS ENUM ('n', 'r', 'sr', 'ssr');

CREATE TYPE element AS ENUM ('water', 'fire', 'earth', 'wind', 'light', 'dark');

CREATE TABLE Weapons (
id SERIAL PRIMARY KEY,
name varchar(80) NOT NULL,
wep_type weapon_type NOT NULL,
wep_ele element NULL,
---rarity weapon_rarity NOT NULL,
ougi_desc varchar(512) NOT NULL,
skill1 FOREIGN KEY REFERENCES WeaponSkills(id) NULL,
skill2 FOREIGN KEY REFERENCES WeaponSkills(id) NULL,
skill3 FOREIGN KEY REFERENCES WeaponSkills(id) NULL,

);

CREATE TABLE WeaponSkills (
id SERIAL INT,
name varchar(80),
description varchar(256)
);
