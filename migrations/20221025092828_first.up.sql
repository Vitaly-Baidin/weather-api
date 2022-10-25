CREATE TABLE IF NOT EXISTS city
(
    id        serial PRIMARY KEY,
    name      varchar(200)     NOT NULL,
    fmt_name  varchar(200)     NOT NULL,
    country   varchar(150)     NOT NULL,
    latitude  double precision NOT NULL,
    longitude double precision NOT NULL,
    constraint coordinates unique (latitude, longitude)
);

CREATE TABLE IF NOT EXISTS temperature
(
    timestamp bigserial,
    temp      float NOT NULL,
    city_id   int REFERENCES city (id),
    data      jsonb,
    PRIMARY KEY (timestamp, city_id)
);