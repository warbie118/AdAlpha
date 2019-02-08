CREATE TABLE assets (
   isin           VARCHAR PRIMARY KEY,
   asset_name     VARCHAR(1000) NOT NULL
);

CREATE TABLE investors (
  investor_id INT PRIMARY KEY,
  title       VARCHAR NOT NULL,
  surname     VARCHAR NOT NULL
);

CREATE TABLE portfolio (
 investor_id  INT NOT NULL REFERENCES investors(investor_id),
 isin         VARCHAR NOT NULL REFERENCES assets(isin),
 units        DECIMAL NOT NULL,
 constraint units_nonnegative check (units >= 0)
);

CREATE TABLE instructions (
 instruction_id SERIAL PRIMARY KEY,
 investor_id    INT NOT NULL REFERENCES investors(investor_id),
 isin           VARCHAR NOT NULL REFERENCES assets(isin),
 asset_price    NUMERIC(10,2) NOT NULL,
 instruction    VARCHAR NOT NULL,
 currency_code  VARCHAR(3),
 amount         NUMERIC(10,2),
 units          DECIMAL
);

INSERT INTO assets VALUES ('IE00B52L4369', 'BlackRock Institutional Cash Series Sterling Liquidity Agency Inc');
INSERT INTO assets VALUES ('GB00BQ1YHQ70', 'Threadneedle UK Property Authorised Investment Net GBP 1 Acc');
INSERT INTO assets VALUES ('GB00B3X7QG63', 'Vanguard FTSE U.K. All Share Index Unit Trust Accumulation');
INSERT INTO assets VALUES ('GB00BG0QP828', 'Legal & General Japan Index Trust C Class Accumulation');
INSERT INTO assets VALUES ('GB00BPN5P238', 'Vanguard US Equity Index Institutional Plus GBP Accumulation');
INSERT INTO assets VALUES ('IE00B1S74Q32', 'Vanguard U.K. Investment Grade Bond Index Fund GBP Accumulation');

INSERT INTO investors VALUES (1, 'Mr', 'Investor');

INSERT INTO portfolio VALUES (1, 'IE00B52L4369', 44000);
INSERT INTO portfolio VALUES (1, 'GB00BQ1YHQ70', 37931.03448275862069);
INSERT INTO portfolio VALUES (1, 'GB00B3X7QG63', 117.377154137544683);
INSERT INTO portfolio VALUES (1, 'GB00BG0QP828', 179.416082205186756);
INSERT INTO portfolio VALUES (1, 'GB00BPN5P238', 187.862916998747581);
INSERT INTO portfolio VALUES (1, 'IE00B1S74Q32', 695.982284087314141);