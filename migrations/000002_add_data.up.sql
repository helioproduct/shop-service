INSERT INTO users (username, hashed_password, balance)
VALUES 
  ('alice', 'hashedpassword1', 1000),
  ('bob', 'hashedpassword2', 1500),
  ('carol', 'hashedpassword3', 2000),
  ('dave', 'hashedpassword4', 500),
  ('eve', 'hashedpassword5', 3000)
ON CONFLICT (username) DO NOTHING;

INSERT INTO products (name, price)
VALUES 
  ('t-shirt', 80),
  ('cup', 20),
  ('book', 50),
  ('pen', 10),
  ('powerbank', 200),
  ('hoody', 300),
  ('umbrella', 200),
  ('socks', 10),
  ('wallet', 50),
  ('pink-hoody', 500)
ON CONFLICT (name) DO NOTHING;
