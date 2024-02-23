-- +goose Up
-- +goose StatementBegin


CREATE TYPE status AS ENUM ('new','pending','found courier','in shop','go client','in home','success','refused');


CREATE TABLE public.orders
(
    order_id     uuid                                                         NOT NULL,
    company_id   integer                                                      NOT NULL,
    courier_id   character varying(555)      DEFAULT ' '::character varying,
    status       status                      DEFAULT 'new'::character varying NOT NULL,
    map_url      character varying(555)      DEFAULT ' '::character varying,
    put_url      character varying(555)      DEFAULT ' '::character varying,
    created_at   timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at   timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    price        integer                     DEFAULT 0,
    transport_id integer                                                      NOT NULL,
    code         varchar(10)                 default ''
);
CREATE TABLE public.addresses
(
    address_id integer                                              NOT NULL,
    city_id    integer                                              NOT NULL,
    apartment  character varying(100) DEFAULT ''::character varying NOT NULL,
    floor      smallint               DEFAULT 0,
    entrance   character varying(75)  DEFAULT ''::character varying NOT NULL,
    address    character varying(250)                               NOT NULL
);

CREATE TABLE public.items
(
    item_id   integer NOT NULL,
    order_id  uuid    NOT NULL,
    image_url text DEFAULT ' '::text,
    title     text    NOT NULL,
    CONSTRAINT order_fk FOREIGN KEY (order_id) REFERENCES orders (order_id)
);

CREATE TABLE public.receivers
(
    receiver_id integer                                             NOT NULL,
    order_id    uuid                                                NOT NULL,
    person_name character varying(75) DEFAULT ''::character varying NOT NULL,
    address_id  integer                                             NOT NULL,
    phone       character varying(50)                               NOT NULL,
    CONSTRAINT order_fk FOREIGN KEY (order_id) REFERENCES orders (order_id),
    CONSTRAINT address_fk FOREIGN KEY (address_id) REFERENCES addresses (address_id)

);
CREATE TABLE public.senders
(
    sender_id    integer                                             NOT NULL,
    order_id     uuid                                                NOT NULL,
    person_name  character varying(75) DEFAULT ''::character varying NOT NULL,
    address_id   integer                                             NOT NULL,
    phone        character varying(50)                               NOT NULL,
    company_name character varying(50)                               NOT NULL,
    CONSTRAINT order_fk FOREIGN KEY (order_id) REFERENCES orders (order_id),
    CONSTRAINT address_fk FOREIGN KEY (address_id) REFERENCES addresses (address_id)
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS senders;
DROP TABLE IF EXISTS receivers;
DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS orders;
DROP TYPE IF EXISTS status;
-- +goose StatementEnd
