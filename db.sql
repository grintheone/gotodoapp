DROP TABLE IF EXISTS todos;

CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status BOOL NOT NULL,
    created DATETIME NOT NULL,
    due DATETIME NOT NULL
);

INSERT INTO todos (title, status, created, due) VALUES
    ("Buy milk", false, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 10 DAY)),
    ("Buy bread", true, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 2 DAY)),
    ("Go surf", false, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 3 DAY)),
    ("Finish worktask", true, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 5 DAY)),
    ("Enjoy some sun", true, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL 1 DAY));
