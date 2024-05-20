CREATE TABLE IF NOT EXISTS feature_flags(
  feature_name varchar(130) NOT NULL PRIMARY KEY,
  enabled boolean NOT NULL DEFAULT false,
  settings jsonb NULL DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP NULL DEFAULT NULL
);