CREATE TABLE "users" (
    ID serial PRIMARY KEY NOT NULL,
    USERNAME VARCHAR(500) UNIQUE NOT NULL,
    PASSWORD VARCHAR(500) NOT NULL,
    EMAIL VARCHAR(500) UNIQUE NOT NULL,
    IMAGE VARCHAR(500) NOT NULL
);
