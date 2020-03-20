package services

const feedBacksDDL = `create table if not exists feedbacks (
    id bigserial primary key,
    feedBack varchar(30) not null,
	userId_Who integer not null,
	userId_Whom integer not null,
	feedBackTime date not null default CURRENT_DATE,
	remove boolean not null default false
);`