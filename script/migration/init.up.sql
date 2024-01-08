CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL,
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
                       deleted_at TIMESTAMP WITH TIME ZONE,
                       name VARCHAR(255),
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password BYTEA
);

CREATE TABLE rooms (
                        id SERIAL PRIMARY KEY,
                        created_at TIMESTAMP WITH TIME ZONE NOT NULL,
                        updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
                        deleted_at TIMESTAMP WITH TIME ZONE,
                        name VARCHAR(255)
);

CREATE TABLE room_members (
                              id SERIAL PRIMARY KEY,
                              created_at TIMESTAMP WITH TIME ZONE NOT NULL,
                              updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
                              deleted_at TIMESTAMP WITH TIME ZONE,
                              user_id INTEGER,
                              room_id INTEGER
);

CREATE TABLE messages (
                            id SERIAL PRIMARY KEY,
                            created_at TIMESTAMP WITH TIME ZONE NOT NULL,
                            updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
                            deleted_at TIMESTAMP WITH TIME ZONE,
                            sender_id INTEGER,
                            room_id INTEGER,
                            metadata TEXT NULL DEFAULT NULL,
                            content TEXT
);

CREATE TABLE friendships (
                              id SERIAL PRIMARY KEY,
                              created_at TIMESTAMP WITH TIME ZONE NOT NULL,
                              updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
                              deleted_at TIMESTAMP WITH TIME ZONE,
                              user_id INTEGER,
                              friend_id INTEGER
);

CREATE TABLE languages (
                            id SERIAL PRIMARY KEY,
                            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            deleted_at TIMESTAMP WITH TIME ZONE,
                            name VARCHAR(255),
                            code VARCHAR(255)
);

CREATE TABLE message_translations (
                          id SERIAL PRIMARY KEY,
                          message_id INTEGER,
                          language_id INTEGER,
                          supported BOOLEAN NOT NULL DEFAULT FALSE,
                          translated_content TEXT,
                          created_at TIMESTAMP WITH TIME ZONE NOT NULL,
                          updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
                          deleted_at TIMESTAMP WITH TIME ZONE
);