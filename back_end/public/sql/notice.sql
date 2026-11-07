CREATE TABLE IF NOT EXISTS notice (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(11) NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    attachments TEXT NOT NULL,
    like_number BIGINT NOT NULL,
    view_number  BIGINT NOT NULL,
    comment_number BIGINT NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE notice IS '公告'

COMMENT ON COLUMN notice.id IS '主键';
COMMENT ON COLUMN notice.phone IS '手机号';
COMMENT ON COLUMN notice.title IS '标题';
COMMENT ON COLUMN notice.content IS '内容';
COMMENT ON COLUMN notice.attachments IS '附件';
COMMENT ON COLUMN notice.like_number IS '点赞数';
COMMENT ON COLUMN notice.view_number IS '浏览数';
COMMENT ON COLUMN notice.comment_number IS '评论数';
COMMENT ON COLUMN notice.status IS '状态';
COMMENT ON COLUMN notice.created_at IS '创建时间';
COMMENT ON COLUMN notice.updated_at IS '更新时间';
COMMENT ON COLUMN notice.deleted_at IS '删除时间';
