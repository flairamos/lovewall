CREATE TABLE IF NOT EXISTS order (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    nonce_str VARCHAR(32) NOT NULL,
    sign VARCHAR(32) NOT NULL,
    sign_type VARCHAR(32) NOT NULL,
    body VARCHAR(127) NOT NULL,
    out_trade_no VARCHAR(32) NOT NULL,
    total_fee INTEGER NOT NULL,
    spbill_create_ip VARCHAR(64) NOT NULL,
    notify_url VARCHAR(256) NOT NULL,
    trade_type VARCHAR(32) NOT NULL,
    product_id VARCHAR(32) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);


COMMENT ON TABLE order IS '订单表';

COMMENT ON COLUMN order.id IS '主键';
COMMENT ON COLUMN order.user_id IS '用户id';
COMMENT ON COLUMN order.nonce_str IS '随机字符串';
COMMENT ON COLUMN order.sign IS '签名';
COMMENT ON COLUMN order.sign_type IS '签名类型';
COMMENT ON COLUMN order.body IS '商品描述';
COMMENT ON COLUMN order.out_trade_no IS '商户订单号';
COMMENT ON COLUMN order.total_fee IS '标价金额';
COMMENT ON COLUMN order.spbill_create_ip IS '终端IP';
COMMENT ON COLUMN order.notify_url IS '通知地址';
COMMENT ON COLUMN order.trade_type IS '交易类型';
COMMENT ON COLUMN order.product_id IS '商品ID';
COMMENT ON COLUMN order.created_at IS '创建时间';
COMMENT ON COLUMN order.updated_at IS '更新时间';
COMMENT ON COLUMN order.deleted_at IS '删除时间';