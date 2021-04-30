CREATE TABLE IF NOT EXISTS payment_codes(
   id uuid PRIMARY KEY,
   payment_code TEXT,
   name TEXT,
   status TEXT,
   expiration_date TEXT,
   created_at TIMESTAMP,
   updated_at TIMESTAMP
);
