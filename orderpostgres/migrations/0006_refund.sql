-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE refund
(
    "id" serial4,
    "uid" varchar(32) NOT NULL,
    "oid" int4 NOT NULL,
    "pid" int4 NOT NULL,
    "oprice" int4 NOT NULL,
    "price" int4 NOT NULL,
    "way" int2 NOT NULL,
    "number" varchar(100) NOT NULL,
    "serial" varchar(100) NOT NULL,
    "status" int2 NOT NULL,
    "del" int2 NOT NULL,
    "createtime" int4 NOT NULL,
    "updatetime" int4 NOT NULL,
    "details" varchar(255) NOT NULL,
    PRIMARY KEY ("id")
);
CREATE INDEX refund_uid_index ON refund (uid);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE refund;