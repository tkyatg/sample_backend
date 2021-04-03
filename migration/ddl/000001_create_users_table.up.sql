CREATE TYPE user_gender AS ENUM ('男', '女');
COMMENT ON TYPE user_gender IS '性別です';
CREATE TYPE user_residence AS ENUM ('北海道','青森','岩手','宮城','秋田','山形','福島','茨城','栃木','群馬','埼玉','千葉','東京','神奈川','新潟','富山','石川','福井','山梨','長野','岐阜','静岡','愛知','三重','滋賀','京都','大阪','兵庫','奈良','和歌山','鳥取','島根','岡山','広島','山口','徳島','香川','愛媛','高知','福岡','佐賀','長崎','熊本','大分','宮崎','鹿児島','沖縄');
COMMENT ON TYPE user_residence IS '居住地です';
CREATE TYPE user_body_shape AS ENUM ('普通', 'スリム', 'グラマー', 'ぽっちゃり');
COMMENT ON TYPE user_body_shape IS '体型です';

CREATE TABLE IF NOT EXISTS users (
  user_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
  display_name VARCHAR DEFAULT NULL,
  display BOOLEAN NOT NULL,
  email VARCHAR NOT NULL,
  birthday DATE DEFAULT NULL,
  residence user_residence,
  password VARCHAR NOT NULL,
  telephone_number VARCHAR DEFAULT NULL,
  gender user_gender NOT NULL,
  image_url VARCHAR DEFAULT NULL,
  free_time VARCHAR DEFAULT NULL,
  height INTEGER DEFAULT NULL,
  body_shape user_body_shape DEFAULT NULL,
  self_introduction VARCHAR DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at  TIMESTAMP,
  CONSTRAINT users_pkey PRIMARY KEY (user_uuid)
);

COMMENT ON TABLE users IS 'ユーザーを保存するためのテーブルです';
COMMENT ON COLUMN users.user_uuid IS 'レコードを一意に識別するIDです';
COMMENT ON COLUMN users.display_name IS '表示名です';
COMMENT ON COLUMN users.email IS 'メールアドレスです';
COMMENT ON COLUMN users.birthday IS 'メールアドレスです';
COMMENT ON COLUMN users.password IS 'パスワードです';
COMMENT ON COLUMN users.telephone_number IS '電話番号です';
COMMENT ON COLUMN users.image_url IS 'イメージ画像です';
COMMENT ON COLUMN users.free_time IS '時間を取りやすい日時です';
COMMENT ON COLUMN users.self_introduction IS '自己紹介です';
COMMENT ON COLUMN users.created_at IS '作成日時です';
COMMENT ON COLUMN users.updated_at IS '更新日時です';
COMMENT ON COLUMN users.deleted_at IS '削除日時です';
COMMENT ON CONSTRAINT users_pkey ON users IS 'PK制約です';
