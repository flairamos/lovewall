CREATE TABLE IF NOT EXISTS bookmark (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(255) NOT NULL,
    collect_id INTEGER NOT NULL,
    type SMALLINT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE bookmark IS '收藏表';

COMMENT ON COLUMN bookmark.id IS '主键';
COMMENT ON COLUMN bookmark.phone IS '账户';
COMMENT ON COLUMN bookmark.collect_id IS '收藏id';
COMMENT ON COLUMN bookmark.type IS '类型';
COMMENT ON COLUMN bookmark.created_at IS '创建时间';
COMMENT ON COLUMN bookmark.updated_at IS '更新时间';
COMMENT ON COLUMN bookmark.deleted_at IS '删除时间';

