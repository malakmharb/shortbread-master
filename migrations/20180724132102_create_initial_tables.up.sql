create table current_parking_occupancy (
  floor_label char(1) not null, /* 1 or 2 */
  space_label text not null, /* usually a number, but supports char data */
  status text not null,  /* available, occupied, expired */
  started_at timestamp,
  expires_at timestamp,
  ended_at timestamp,
  occupant_id uuid not null,
  unique (floor_label, space_label)
);

create table parking_occupant (
  id uuid primary key,
  resident_name text not null,
  resident_phone text not null,
  other_phones text[] not null,
  resident_room_number text not null
);

create table delivered_sms_messages (
  id uuid primary key,
  parking_occupant_id uuid not null,
  content text not null,
  floor_label char(1) not null,
  space_label text not null,
  sent_at timestamp not null,
  sent_to text[] not null
);

create table resident_information (
  id uuid primary key,
  first_name text,
  last_name text,
  phone_number text,
  room_number text
);
