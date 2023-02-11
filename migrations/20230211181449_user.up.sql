CREATE TABLE IF NOT EXISTS "users" (
                                       "id" uuid not null,
                                       "name" varchar(255) not null,
                                       "email" varchar(320) not null,
                                       "words_per_day" smallint,
                                       PRIMARY KEY ("id")
);