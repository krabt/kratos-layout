CREATE TABLE "user" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT, -- 组织主键
  "name" VARCHAR(64) NOT NULL, -- 组织名称
  "code" VARCHAR(64) NOT NULL, -- 组织编码，系统内唯一
  "status" BOOLEAN NOT NULL DEFAULT TRUE, -- 启用状态
  "created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
  "updated_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 更新时间
  "deleted_at" DATETIME, -- 软删除时间
);