

-- 消息表
CREATE TABLE IF NOT EXISTS message (
    id SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL,
    receiver_id INTEGER,
    group_id INTEGER,
    message_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP 
    -- 约束：要么是私信（receiver_id 不为空），要么是群聊（group_id 不为空）
    CONSTRAINT check_message_type CHECK ((receiver_id IS NOT NULL AND group_id IS NULL) OR (receiver_id IS NULL AND group_id IS NOT NULL))
);


COMMENT ON TABLE message IS '消息';

COMMENT ON COLUMN message.id IS '主键';
COMMENT ON COLUMN message.sender_id IS '发送人主键';
COMMENT ON COLUMN message.receiver_id IS '接收人主键';
COMMENT ON COLUMN message.group_id IS '类型';
COMMENT ON COLUMN message.message_text IS '消息内容';
COMMENT ON COLUMN message.created_at IS '创建时间';
COMMENT ON COLUMN message.updated_at IS '更新时间';
COMMENT ON COLUMN message.deleted_at IS '删除时间';



-- 群表
CREATE TABLE IF NOT EXISTS message_group (
    id SERIAL PRIMARY KEY,
    create_user INTEGER NOT NULL,
    create_phone VARCHAR(100) NOT NULL,
    group_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP 
);

COMMENT ON TABLE message_group IS '消息群组';

COMMENT ON COLUMN message_group.id IS '主键';
COMMENT ON COLUMN message_group.create_user IS '创建人主键';
COMMENT ON COLUMN message_group.create_phone IS '创建人账户';
COMMENT ON COLUMN message_group.group_name IS '群组名称';
COMMENT ON COLUMN message_group.created_at IS '创建时间';
COMMENT ON COLUMN message_group.updated_at IS '更新时间';
COMMENT ON COLUMN message_group.deleted_at IS '删除时间';




-- 用户与群的关联表
CREATE TABLE IF NOT EXISTS user_group_relation (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP 
);


COMMENT ON TABLE user_group_relation IS '消息群组关联表';

COMMENT ON COLUMN user_group_relation.id IS '主键';
COMMENT ON COLUMN user_group_relation.user_id IS '用户主键';
COMMENT ON COLUMN user_group_relation.group_id IS '群组主键';
COMMENT ON COLUMN user_group_relation.created_at IS '创建时间';
COMMENT ON COLUMN user_group_relation.updated_at IS '更新时间';
COMMENT ON COLUMN user_group_relation.deleted_at IS '删除时间';