-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE orderinfo
(
    "id" serial4,
    "uid" varchar(32) NOT NULL,
    "cid" int4 NOT NULL,
    "pid" int4 NOT NULL,
    "cprice" int4 NOT NULL,
    "priceremark" varchar(255),
    "createtime" int4 NOT NULL,
    "updatetime" int4 NOT NULL,
    "status" int2 NOT NULL,
    "del" int2 NOT NULL,
    "details" varchar(255),
    PRIMARY KEY ("id")
);
CREATE INDEX orderinfo_uid_index ON orderinfo (uid);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE orderinfo;