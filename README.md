# Go / PGX Compiled Size

This is test of the size of compiled Go programs and how much pgx and other database libraries cost in size.

Tests run with Go 1.15.3 on MacOS with pgx v4.9.2.

| Libary | Size |
| ------ | ---- |
| fmt / Hello world | 2.1M |
| pgconn | 6.5M |
| pgx | 8.4M |
| pgxpool | 8.5M |
| stdlib | 8.7M |
