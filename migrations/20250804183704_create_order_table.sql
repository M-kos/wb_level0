-- +goose Up
-- +goose StatementBegin
CREATE USER customer WITH PASSWORD 'customer';

CREATE TABLE IF NOT EXISTS "order"
(
	id                  BIGSERIAL,
	order_uid           VARCHAR(255) NOT NULL
		UNIQUE,
	track_number        VARCHAR(255) NOT NULL
		UNIQUE,
	entry               VARCHAR(255) NOT NULL
		UNIQUE,
	delivery_id         BIGINT       NOT NULL,
	payment_id          BIGINT       NOT NULL,
	locale_id           BIGINT       NOT NULL,
	internal_signature  VARCHAR(255),
	customer_id         VARCHAR(255) NOT NULL,
	delivery_service_id BIGINT       NOT NULL,
	shardkey            VARCHAR(255) NOT NULL,
	sm_id               INT          NOT NULL,
	date_created        TIMESTAMP    NOT NULL DEFAULT now( ),
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

CREATE TABLE IF NOT EXISTS delivery
(
	id BIGSERIAL,
	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS payment
(
	id BIGSERIAL,
	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS locale
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS delivery_service
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS item
(
	id   BIGSERIAL,
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY ( id )
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
DROP TYPE IF EXISTS user_status;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd


-- insert into users (uuid, first_name, last_name, user_type, email, "password") values ('04a2b9eb-2c3f-4946-a5f3-2db6d51435fe', '111', '111', '1', 'qwe@qwe.rr', '123456');
-- insert into users (uuid, first_name, last_name, user_type, email, "password") values ('1232582a-5bf3-486d-9ccb-9b399c0eddc2', '222', '222', '2', 'asd@qwe.rr', '123456');
