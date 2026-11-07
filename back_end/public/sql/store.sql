CREATE TABLE IF NOT EXISTS store (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(11) NOT NULL,
    name VARCHAR(20) NOT NULL,
    authentication BOOLEAN NOT NULL,
    authentication_name VARCHAR(20) NOT NULL,
    background_img TEXT NOT NULL,
    content VARCHAR(1000) NOT NULL,
    photos TEXT NOT NULL,
    price NUMERIC(5,2),
    status BOOLEAN NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE store IS '帖子'

COMMENT ON COLUMN store.id IS '主键';
COMMENT ON COLUMN store.phone IS '开店人手机号';
COMMENT ON COLUMN store.name IS '店铺名称';
COMMENT ON COLUMN store.authentication IS '是否认证';
COMMENT ON COLUMN store.authentication_name IS '认证信息';
COMMENT ON COLUMN store.background_img IS '店铺背景图片';
COMMENT ON COLUMN store.content IS '店铺介绍';
COMMENT ON COLUMN store.photos IS '店铺介绍图片';
COMMENT ON COLUMN store.content IS '帖子内容';
COMMENT ON COLUMN store.price IS '店铺消费水平';
COMMENT ON COLUMN store.status IS '店铺审核状态';
COMMENT ON COLUMN store.created_at IS '创建时间';
COMMENT ON COLUMN store.updated_at IS '更新时间';
COMMENT ON COLUMN store.deleted_at IS '删除时间';