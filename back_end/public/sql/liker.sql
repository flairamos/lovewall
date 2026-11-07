CREATE TABLE IF NOT EXISTS liker (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(255) NOT NULL,
    collect_id INTEGER NOT NULL,
    type SMALLINT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE liker IS '点赞表';

COMMENT ON COLUMN liker.id IS '主键';
COMMENT ON COLUMN liker.phone IS '账户';
COMMENT ON COLUMN liker.collect_id IS '点赞帖子id';
COMMENT ON COLUMN liker.type IS '类型';
COMMENT ON COLUMN liker.created_at IS '创建时间';
COMMENT ON COLUMN liker.updated_at IS '更新时间';
COMMENT ON COLUMN liker.deleted_at IS '删除时间';