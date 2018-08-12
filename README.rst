About
=====

Tool which scrapes all Katas from codewars.com and puts them into PogstreSQL.

Currently the following parameters are scraped::

    - Kata ID,
    - Title,
    - URL,
    - kyu (kata difficulty level),
    - languages that kata is implemented with,
    - Kata keyword tags.

Dependencies
============

- go1.10.3
- psql (PostgreSQL) 10.4

Usage
=====

To compile and run the scraper::

    $ go run main.go

To compile the binary::

    $ go build main.go
