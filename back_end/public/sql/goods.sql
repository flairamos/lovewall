CREATE TABLE IF NOT EXISTS goods (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    imageUrl text NOT NULL,
    category VARCHAR(255) NOT NULL,
    price    VARCHAR(255) NOT NULL,
    original_price VARCHAR(255) NOT NULL,
    publish_time VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    user_id integer,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE goods IS '二手商品表';

COMMENT ON COLUMN goods.id IS '主键';
COMMENT ON COLUMN goods.title IS '标题';
COMMENT ON COLUMN goods.description IS '描述';
COMMENT ON COLUMN goods.imageUrl IS '图片地址';
COMMENT ON COLUMN goods.category IS '商品类型';
COMMENT ON COLUMN goods.price IS '价格';
COMMENT ON COLUMN goods.original_price IS '商品原价';
COMMENT ON COLUMN goods.publish_time IS '发布时间';
COMMENT ON COLUMN goods.status IS '发布时间';
COMMENT ON COLUMN goods.user_id IS '用户主键';
COMMENT ON COLUMN goods.created_at IS '创建时间';
COMMENT ON COLUMN goods.updated_at IS '更新时间';
COMMENT ON COLUMN goods.deleted_at IS '删除时间';
