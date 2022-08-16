
CREATE TABLE public.users (
	id SERIAL,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(120) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

---

CREATE TABLE public.asserts (
	id SERIAL,
    name VARCHAR(50) NOT NULL,
    amount INT NULL,
    price NUMERIC(6, 2) NULL,
    average_price NUMERIC(6, 2) NULL,
    user_id INT NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT asserts_pkey PRIMARY KEY (id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id)
);

CREATE INDEX asserts_identifier_idx ON public.asserts (IDENTIFIER);

CREATE INDEX asserts_name_idx ON public.asserts (name);

---

CREATE TABLE public.launchs (
	id SERIAL,
    operation VARCHAR(4) NOT NULL,
    amount INT NOT NULL,
    price NUMERIC(6, 2) NOT NULL,
    date_operation TIMESTAMP NOT NULL, 
    broker VARCHAR(30) NOT NULL,
    assert_id INT NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT launchs_pkey PRIMARY KEY (id),
    CONSTRAINT fk_assert_id FOREIGN KEY (assert_id) REFERENCES public.asserts(id)
);

CREATE INDEX launchs_identifier_idx ON public.launchs (IDENTIFIER);

CREATE INDEX launchs_operation_idx ON public.launchs (operation);

CREATE INDEX launchs_date_operation_idx ON public.launchs (date_operation);

CREATE INDEX launchs_broker_idx ON public.launchs (broker);

---
