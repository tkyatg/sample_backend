CREATE TABLE IF NOT EXISTS users (
  user_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
  display_name VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  telephone_number VARCHAR NOT NULL,
  gender INT NOT NULL,
  image_url VARCHAR,
  free_time VARCHAR,
  self_introduction VARCHAR,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at  TIMESTAMP,
  CONSTRAINT users_pkey PRIMARY KEY (user_uuid)
);

COMMENT ON TABLE users IS 'ユーザーを保存するためのテーブルです';
COMMENT ON COLUMN users.user_uuid IS 'レコードを一意に識別するIDです';
COMMENT ON COLUMN users.display_name IS '表示名です';
COMMENT ON COLUMN users.email IS 'メールアドレスです';
COMMENT ON COLUMN users.password IS 'パスワードです';
COMMENT ON COLUMN users.telephone_number IS '電話番号です';
COMMENT ON COLUMN users.gender IS '性別です';
COMMENT ON COLUMN users.image_url IS 'イメージ画像です';
COMMENT ON COLUMN users.free_time IS '時間を取りやすい日時です';
COMMENT ON COLUMN users.self_introduction IS '自己紹介です';
COMMENT ON COLUMN users.created_at IS '作成日時です';
COMMENT ON COLUMN users.updated_at IS '更新日時です';
COMMENT ON COLUMN users.deleted_at IS '削除日時です';
COMMENT ON CONSTRAINT users_pkey ON users IS 'PK制約です';
