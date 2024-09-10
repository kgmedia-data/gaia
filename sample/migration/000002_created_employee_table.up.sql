create table if not exists employee (
    id bigserial primary key,
    employee_number varchar(255) not null unique,
    first_name varchar(255) not null,
    last_name varchar(255),
    birth_date date,
    department_id bigint not null,
    is_deleted boolean not null default false,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    foreign key (department_id) references department (id) on delete cascade
);