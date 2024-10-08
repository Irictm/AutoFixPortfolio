-- Table: public.vehicles

CREATE SEQUENCE vehicles_id_seq AS bigint;
CREATE TABLE IF NOT EXISTS public.vehicles
(
    id bigint NOT NULL DEFAULT nextval('vehicles_id_seq'::regclass),
    patent text,
    brand text,
    model text,
    type text,
    fabrication_date timestamp with time zone,
    motor_type text,
    seats smallint,
    mileage integer,
    CONSTRAINT vehicles_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.vehicles
    OWNER to postgres;
ALTER SEQUENCE vehicles_id_seq OWNED BY public.vehicles.id;


-- Table: public.bonuses

CREATE SEQUENCE bonuses_id_seq AS bigint;
CREATE TABLE IF NOT EXISTS public.bonuses
(
    id bigint NOT NULL DEFAULT nextval('bonuses_id_seq'::regclass),
    brand text,
    remaining smallint,
    amount integer,
    CONSTRAINT bonuses_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.bonuses
    OWNER to postgres;
ALTER SEQUENCE bonuses_id_seq OWNED BY public.bonuses.id;


-- Table: public.receipts

CREATE SEQUENCE receipts_id_seq AS bigint;
CREATE TABLE IF NOT EXISTS public.receipts
(
    id bigint NOT NULL DEFAULT nextval('receipts_id_seq'::regclass),
    operations_amount integer,
    recharge_amount integer,
    discount_amount integer,
    iva_amount integer,
    total_amount integer,
    bonus_consumed boolean,
    CONSTRAINT receipts_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.receipts
    OWNER to postgres;
ALTER SEQUENCE receipts_id_seq OWNED BY public.receipts.id;


-- Table: public.repairs

CREATE SEQUENCE repairs_id_seq AS bigint;
CREATE TABLE IF NOT EXISTS public.repairs
(
    id bigint NOT NULL DEFAULT nextval('repairs_id_seq'::regclass),
    date_of_admission timestamp with time zone,
    date_of_release timestamp with time zone,
    date_of_pick_up timestamp with time zone,
    id_receipt bigint REFERENCES receipts ON DELETE CASCADE ON UPDATE CASCADE,
    id_vehicle bigint REFERENCES vehicles ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT repairs_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.repairs
    OWNER to postgres;
ALTER SEQUENCE repairs_id_seq OWNED BY public.repairs.id;


-- Table: public.operations

CREATE SEQUENCE operations_id_seq AS bigint;
CREATE TABLE IF NOT EXISTS public.operations
(
    id bigint NOT NULL DEFAULT nextval('operations_id_seq'::regclass),
    patent text,
    id_operation_type bigint,
    date timestamp with time zone,
    cost integer,
    id_repair bigint REFERENCES repairs ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT operations_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.operations
    OWNER to postgres;
ALTER SEQUENCE operations_id_seq OWNED BY public.operations.id;


----- TABLES FOR TARIFFS -----

-- Table: public.tariff_operations

CREATE SEQUENCE tariff_operations_id_seq AS bigint;
CREATE TABLE IF NOT EXISTS public.tariff_operations
(
    id bigint NOT NULL DEFAULT nextval('tariff_operations_id_seq'::regclass),
    motor_type text,
    id_operation_type bigint,
    value double precision,
    CONSTRAINT tariff_operations_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tariff_operations
    OWNER to postgres;
ALTER SEQUENCE tariff_operations_id_seq OWNED BY public.tariff_operations.id;


-- Table: public.tariff_repair_number

CREATE SEQUENCE tariff_repair_number_id_seq AS bigint;
CREATE TABLE IF NOT EXISTS public.tariff_repair_number
(
    id bigint NOT NULL DEFAULT nextval('tariff_repair_number_id_seq'::regclass),
    motor_type text,
    bottom double precision,
    top double precision,
    value double precision,
    CONSTRAINT tariff_repair_number_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tariff_repair_number
    OWNER to postgres;
ALTER SEQUENCE tariff_repair_number_id_seq OWNED BY public.tariff_repair_number.id;


-- Table: public.tariff_mileage

CREATE SEQUENCE tariff_mileage_id_seq AS bigint;
CREATE TABLE IF NOT EXISTS public.tariff_mileage
(
    id bigint NOT NULL DEFAULT nextval('tariff_mileage_id_seq'::regclass),
    vehicle_type text,
    bottom double precision,
    top double precision,
    value double precision,
    CONSTRAINT tariff_mileage_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tariff_mileage
    OWNER to postgres;
ALTER SEQUENCE tariff_mileage_id_seq OWNED BY public.tariff_mileage.id;


-- Table: public.tariff_antiquity

CREATE SEQUENCE tariff_antiquity_id_seq AS bigint;
CREATE TABLE IF NOT EXISTS public.tariff_antiquity
(
    id bigint NOT NULL DEFAULT nextval('tariff_antiquity_id_seq'::regclass),
    vehicle_type text,
    bottom double precision,
    top double precision,
    value double precision,
    CONSTRAINT tariff_antiquity_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tariff_antiquity
    OWNER to postgres;
ALTER SEQUENCE tariff_antiquity_id_seq OWNED BY public.tariff_antiquity.id;