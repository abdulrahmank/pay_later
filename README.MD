**Dependencies**

1. Postgres

**Installation**

Set following env variables:
 
`DATABASE_HOST`, `DATABASE_PORT`, `DATABASE_USER`, `DATABASE_DB`, `DATABASE_PSSWD`

Or Provide psql url:

`DATABASE_URL`

**SQL Migrations**

    create table users(
        name text, 
        mailid text, 
        creditlimit decimal, 
        dues decimal);

    create table merchants(
        name text, 
        discount decimal);
    
    create table txn(
        id serial primary key, 
        user_name text, 
        merchant_name text, 
        amt decimal);
       
**Running**

Ensure postgres server is running, run following command in terminal

`./pay_later_service {new|report|update|payback} {args}`

All commands have help doc.