-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE commodity
(
  "id" serial4,
  "name" varchar(20) NOT NULL,
  "details" varchar(255) NOT NULL,
  "oldprice" int4 NOT NULL,
  "newprice" int4 NOT NULL,
  "priceremark" varchar(255) NOT NULL,
  "updatetime" int4 NOT NULL,
  "status" int2 NOT NULL,
  PRIMARY KEY ("id")
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE commodity;