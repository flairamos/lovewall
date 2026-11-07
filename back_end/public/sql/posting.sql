CREATE TABLE IF NOT EXISTS posting (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(11) NOT NULL,
    name VARCHAR(20) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    authentication BOOLEAN NOT NULL,
    authentication_name VARCHAR(20) NOT NULL,
    type SMALLINT NOT NULL,
    like_people text,
    content VARCHAR(1000) NOT NULL,
    photos TEXT NOT NULL,
    comment_people text,
    ref_comment integer,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE notice IS '帖子'

COMMENT ON COLUMN posting.id IS '主键';
COMMENT ON COLUMN posting.phone IS '手机号';
COMMENT ON COLUMN posting.name IS '昵称';
COMMENT ON COLUMN posting.avatar IS '头像';
COMMENT ON COLUMN posting.authentication IS '是否认证';
COMMENT ON COLUMN posting.authentication_name IS '认证信息';
COMMENT ON COLUMN posting.type IS '帖子类型';
COMMENT ON COLUMN posting.content IS '帖子内容';
COMMENT ON COLUMN posting.photos IS '帖子图片';
COMMENT ON COLUMN posting.created_at IS '创建时间';
COMMENT ON COLUMN posting.updated_at IS '更新时间';
COMMENT ON COLUMN posting.deleted_at IS '删除时间';