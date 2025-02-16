DELETE FROM transfers
WHERE from_user_id IN (SELECT id FROM users WHERE username IN ('alice', 'bob', 'carol', 'dave', 'eve'))
   OR to_user_id IN (SELECT id FROM users WHERE username IN ('alice', 'bob', 'carol', 'dave', 'eve'));

DELETE FROM purchases
WHERE user_id IN (SELECT id FROM users WHERE username IN ('alice', 'bob', 'carol', 'dave', 'eve'));

DELETE FROM products
WHERE name IN ('t-shirt', 'cup', 'book', 'pen', 'powerbank', 'hoody', 'umbrella', 'socks', 'wallet', 'pink-hoody');

DELETE FROM users
WHERE username IN ('alice', 'bob', 'carol', 'dave', 'eve');
