# Install

Data migrations are run as simple sql files stored under the migrate directory. Initial migration is called `install.sql` but each subsequent migration should lead with the date in `yyyy_mm_dd` format

## Set Up

Create your database (MySQL)

```sh
mysql -u<username> -p<password> -e "CREATE DATABASE <dbname>"
```

Copy .example.env to .env and populate dummy settings with your actual settings

## Build

just run make

```sh
make
```
