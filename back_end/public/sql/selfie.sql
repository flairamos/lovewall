CREATE TABLE IF NOT EXISTS selfie (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(11) NOT NULL,
    user_id INTEGER NOT NULL,
    like_num INTEGER NOT NULL,
    photos TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE selfie IS '晒美图';

COMMENT ON COLUMN selfie.id IS '主键';
COMMENT ON COLUMN selfie.phone IS '手机号';
COMMENT ON COLUMN selfie.user_id IS '用户id';
COMMENT ON COLUMN selfie.like_num IS '点赞数';
COMMENT ON COLUMN selfie.photos IS '图片';
COMMENT ON COLUMN selfie.created_at IS '创建时间';
COMMENT ON COLUMN selfie.updated_at IS '更新时间';
COMMENT ON COLUMN selfie.deleted_at IS '删除时间';