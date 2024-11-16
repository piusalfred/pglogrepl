[![godoc](https://godoc.org/github.com/piusalfred/pglogrepl?status.svg)](https://godoc.org/github.com/jackc/pglogrepl)
[![ci pipeline](https://github.com/jackc/pglogrepl/actions/workflows/ci.yml/badge.svg)](https://github.com/piusalfred/pglogrepl/actions/workflows/tasks.yml)

# pglogrepl

Golang package for PostgreSQL logical replication. Makes the use of [github.com/jackc/pgx/v5/pgconn](https://pkg.go.dev/github.com/jackc/pgx/v5/pgconn) to connect to PostgreSQL.

> [!NOTE]
> This is a fork of [github.com/jackc/pglogrepl](https://github.com/jackc/pglogrepl) with some modifications I made while learning about PostgreSQL logical replication.


### examples
You can find examples in [example](example) directory.

### links
- [PostgreSQL Logical Decoding Concepts](https://www.postgresql.org/docs/current/logicaldecoding-explanation.html#LOGICALDECODING-EXPLANATION)
- [Streaming Replication protocol Interface](https://www.postgresql.org/docs/current/logicaldecoding-walsender.html)
- [PostgreSQL Streaming Replication Protocol](https://www.postgresql.org/docs/current/protocol-replication.html)
- [PostgreSQL Logical Streaming Replication Protocol](https://www.postgresql.org/docs/current/protocol-logical-replication.html)
- [Replication Protocol Message Types](https://www.postgresql.org/docs/current/protocol-message-types.html)
- [Replication Protocol Message Formats](https://www.postgresql.org/docs/current/protocol-message-formats.html)
- [Replication Protocol Error Fields](https://www.postgresql.org/docs/current/protocol-error-fields.html)
- [Logical Replication Message Formats](https://www.postgresql.org/docs/current/protocol-logicalrep-message-formats.html)
- [Replication Protocol Changes](https://www.postgresql.org/docs/current/protocol-changes.html)
