-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS customer
(
	id         BIGSERIAL,
	first_name VARCHAR(255) NOT NULL,
	last_name  VARCHAR(255) NOT NULL,
	phone      VARCHAR(255) NOT NULL
		UNIQUE,
	email      VARCHAR(255),

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS region
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL
		UNIQUE,

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS city
(
	id        BIGSERIAL,
	name      VARCHAR(255) NOT NULL,
	region_id BIGINT       NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( region_id )
		REFERENCES region ( id )
);

CREATE TABLE IF NOT EXISTS address
(
	id      BIGSERIAL,
	zip     VARCHAR(255) NOT NULL,
	address VARCHAR(255) NOT NULL,
	city_id BIGINT       NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( city_id )
		REFERENCES city ( id )
);

CREATE TABLE IF NOT EXISTS customer_address
(
	customer_id BIGINT NOT NULL,
	address_id  BIGINT NOT NULL,

	PRIMARY KEY ( customer_id, address_id ),
	FOREIGN KEY ( customer_id )
		REFERENCES customer ( id )
		ON DELETE CASCADE,
	FOREIGN KEY ( address_id )
		REFERENCES address ( id )
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS delivery
(
	id          BIGSERIAL,
	customer_id BIGINT NOT NULL,
	address_id  BIGINT NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( address_id )
		REFERENCES address ( id ),
	FOREIGN KEY ( customer_id )
		REFERENCES customer ( id )
);

CREATE TABLE IF NOT EXISTS currency
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL
		UNIQUE,

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS provider
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL
		UNIQUE,

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS bank
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL
		UNIQUE,

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS payment
(
	id            BIGSERIAL,
	transaction   VARCHAR(255) NOT NULL,
	request_id    VARCHAR(255),
	currency_id   BIGINT       NOT NULL,
	provider_id   BIGINT       NOT NULL,
	amount        INT          NOT NULL,
	payment_dt    TIMESTAMP    NOT NULL,
	bank_id       BIGINT       NOT NULL,
	delivery_cost INT          NOT NULL,
	goods_total   INT          NOT NULL,
	custom_fee    INT          NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( currency_id )
		REFERENCES currency ( id ),
	FOREIGN KEY ( provider_id )
		REFERENCES provider ( id ),
	FOREIGN KEY ( bank_id )
		REFERENCES bank ( id )
);


CREATE TABLE IF NOT EXISTS locale
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL
		UNIQUE,

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS delivery_service
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL
		UNIQUE,

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS brand
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL
		UNIQUE,

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS item_status
(
	id    BIGSERIAL,
	value INT NOT NULL
		UNIQUE,

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS item
(
	id           BIGSERIAL,
	chrt_id      BIGINT       NOT NULL,
	track_number VARCHAR(255) NOT NULL,
	price        INT          NOT NULL,
	rid          VARCHAR(255) NOT NULL,
	name         VARCHAR(255) NOT NULL,
	sale         INT          NOT NULL,
	size         VARCHAR(255),
	total_price  INT          NOT NULL,
	nm_id        BIGINT       NOT NULL,
	brand_id     BIGINT       NOT NULL,
	status_id    BIGINT       NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( brand_id )
		REFERENCES brand ( id ),
	FOREIGN KEY ( status_id )
		REFERENCES item_status ( id )
);

CREATE TABLE IF NOT EXISTS "order"
(
	id                  BIGSERIAL,
	order_uid           VARCHAR(255) NOT NULL
		UNIQUE,
	track_number        VARCHAR(255) NOT NULL
		UNIQUE,
	entry               VARCHAR(255) NOT NULL,
	delivery_id         BIGINT       NOT NULL,
	payment_id          BIGINT       NOT NULL,
	locale_id           BIGINT       NOT NULL,
	internal_signature  VARCHAR(255),
	customer_id         VARCHAR(255) NOT NULL,
	delivery_service_id BIGINT       NOT NULL,
	shardkey            VARCHAR(255) NOT NULL,
	sm_id               INT          NOT NULL,
	date_created        TIMESTAMP    NOT NULL,
	oof_shard           VARCHAR(255) NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( delivery_id )
		REFERENCES delivery ( id ),
	FOREIGN KEY ( payment_id )
		REFERENCES payment ( id ),
	FOREIGN KEY ( locale_id )
		REFERENCES locale ( id ),
	FOREIGN KEY ( delivery_service_id )
		REFERENCES delivery_service ( id )
);


CREATE TABLE IF NOT EXISTS order_item
(
	order_id BIGINT NOT NULL,
	item_id  BIGINT NOT NULL,

	PRIMARY KEY ( order_id, item_id ),
	FOREIGN KEY ( order_id )
		REFERENCES "order" ( id )
		ON DELETE CASCADE,
	FOREIGN KEY ( item_id )
		REFERENCES item ( id )
		ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_item;
DROP TABLE IF EXISTS customer_address;
DROP TABLE IF EXISTS "order";
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS item_status;
DROP TABLE IF EXISTS brand;
DROP TABLE IF EXISTS delivery_service;
DROP TABLE IF EXISTS locale;
DROP TABLE IF EXISTS payment;
DROP TABLE IF EXISTS bank;
DROP TABLE IF EXISTS provider;
DROP TABLE IF EXISTS currency;
DROP TABLE IF EXISTS delivery;
DROP TABLE IF EXISTS address;
DROP TABLE IF EXISTS city;
DROP TABLE IF EXISTS region;
DROP TABLE IF EXISTS customer;

-- +goose StatementEnd


-- insert into users (uuid, first_name, last_name, user_type, email, "password") values ('04a2b9eb-2c3f-4946-a5f3-2db6d51435fe', '111', '111', '1', 'qwe@qwe.rr', '123456');
-- insert into users (uuid, first_name, last_name, user_type, email, "password") values ('1232582a-5bf3-486d-9ccb-9b399c0eddc2', '222', '222', '2', 'asd@qwe.rr', '123456');
