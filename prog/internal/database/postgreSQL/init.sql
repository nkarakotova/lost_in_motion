drop table if exists coaches cascade;
create table coaches
(
	coach_id serial primary key,
    name text,
	description text
);
alter table coaches add constraint unique_name_coach unique (name);

drop table if exists directions cascade;
create table directions
(
	direction_id serial primary key,
    name text,
	description text,
    acceptable_gender int
);
alter table directions add constraint unique_name_direction unique (name);

drop table if exists halls cascade;
create table halls
(
	hall_id serial primary key,
    number int,
	capacity int
);
alter table halls add constraint unique_number_hall unique (number);

drop table if exists subscriptions cascade;
create table subscriptions
(
	subscription_id serial primary key,
    trainings_num int,
	remaining_trainings_num int,
	cost int,
    start_date date,
	end_date date
);

drop table if exists clients cascade;
create table clients
(
	client_id serial primary key,
    subscription_id int references subscriptions (subscription_id) default null,
    name text,
	telephone text,
	mail text,
    password text,
	age int,
	gender int
);
alter table clients add constraint unique_telephone_client unique (telephone);

drop table if exists trainings cascade;
create table trainings 
(
	training_id serial  primary key,
	coach_id int references coaches (coach_id),
	hall_id int references halls (hall_id),
	direction_id int references directions (direction_id),
    name text,
	date_time timestamp,
	places_num int,
    available_places_num int,
    acceptable_age int
);

drop table if exists clients_trainings cascade;
create table clients_trainings
(
	client_id int references clients(client_id) on delete cascade, 
	training_id int references trainings(training_id) on delete cascade
);

drop table if exists coaches_directions cascade;
create table coaches_directions
(
	coach_id int references coaches(coach_id) on delete cascade, 
	direction_id int references directions(direction_id) on delete cascade
);
