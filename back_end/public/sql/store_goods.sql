CREATE TABLE IF NOT EXISTS store_goods (
    id SERIAL PRIMARY KEY,
    store_id INT NOT NULL,
    goods_name VARCHAR(20) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    photos TEXT NOT NULL,
    price NUMERIC(5,2),
    status BOOLEAN NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE store_goods IS '商铺商品'

COMMENT ON COLUMN store_goods.id IS '主键';
COMMENT ON COLUMN store_goods.store_id IS '商铺id';
COMMENT ON COLUMN store_goods.goods_name IS '商品名称';
COMMENT ON COLUMN store_goods.content IS '商品介绍';
COMMENT ON COLUMN store_goods.photos IS '商品图片';
COMMENT ON COLUMN store_goods.price IS '商品价格';
COMMENT ON COLUMN store_goods.status IS '上下架状态';
COMMENT ON COLUMN store_goods.created_at IS '创建时间';
COMMENT ON COLUMN store_goods.updated_at IS '更新时间';
COMMENT ON COLUMN store_goods.deleted_at IS '删除时间';