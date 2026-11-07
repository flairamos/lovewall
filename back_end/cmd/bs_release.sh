#!/bin/bash

# 提示用户输入文件名
echo "请输入Go文件的名称（不包括扩展名，例如：main）："
read fileName

# 如果用户未输入文件名，默认使用test
if [ -z "$fileName" ]; then
    fileName="test"
fi

# 将文件名的首字母转为大写
capitalizedFileName="${fileName^}"

handler="../handler"
model="../model"
router="../router"
service="../service"

if [ ! -d "$handler" ]; then
    mkdir -p "$handler"
fi

if [ ! -d "$model" ]; then
    mkdir -p "$model"
fi

if [ ! -d "$router" ]; then
    mkdir -p "$router"
fi

if [ ! -d "$service" ]; then
    mkdir -p "$service"
fi



# handler 开始
handlerContent="package $(basename $handler)

import (
	\"github.com/gofiber/fiber/v2\"
	\"shoollove/model\"
	\"shoollove/service\"
	\"shoollove/pkg/respond\"
)

// ${capitalizedFileName}Add 添加
func ${capitalizedFileName}Add(c *fiber.Ctx) error {

	var req model.${capitalizedFileName}

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, \"参数错误\", err)
	}

    resp , err := service.${capitalizedFileName}Add(req)
    if err != nil {
        return respond.Fail(c, err.Error())
    }
	return respond.Ok(c, resp)
}


// ${capitalizedFileName}Edit 修改
func ${capitalizedFileName}Edit(c *fiber.Ctx) error {

	var req model.${capitalizedFileName}

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, \"参数错误\", err)
	}

    resp , err := service.${capitalizedFileName}Edit(req)
    if err != nil {
        return respond.Fail(c, err.Error())
    }
	return respond.Ok(c, resp)
}


// ${capitalizedFileName}Delete 删除
func ${capitalizedFileName}Delete(c *fiber.Ctx) error {

	var req model.${capitalizedFileName}

	if err := c.BodyParser(&req); err != nil {
		return respond.Fail(c, \"参数错误\", err)
	}

    resp , err := service.${capitalizedFileName}Delete(req)
    if err != nil {
        return respond.Fail(c, err.Error())
    }
	return respond.Ok(c, resp)
}

// ${capitalizedFileName}Get 查询
func ${capitalizedFileName}Get(c *fiber.Ctx) error {

	var req model.${capitalizedFileName}

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, \"参数错误\", err)
	}

    resp , err := service.${capitalizedFileName}Get(req)
    if err != nil {
        return respond.Fail(c, err.Error())
    }
	return respond.Ok(c, resp)
}


// ${capitalizedFileName}Form 精准查询
func ${capitalizedFileName}Form(c *fiber.Ctx) error {

	var req model.${capitalizedFileName}

	if err := c.QueryParser(&req); err != nil {
		return respond.Fail(c, \"参数错误\", err)
	}

    resp , err := service.${capitalizedFileName}Form(req)
    if err != nil {
        return respond.Fail(c, err.Error())
    }
	return respond.Ok(c, resp)
}
"
# handler 结束

# model 开始
modelContent="package $(basename $model)

import (
    \"gorm.io/gorm\"
    \"shoollove/pkg/utils\"
)

// $capitalizedFileName 模型
type $capitalizedFileName struct {
}

// ${capitalizedFileName}Req 添加请求参数
type ${capitalizedFileName}Req struct {
    Id   string \`json:\"id,omitempty\"\`
    Page int \`json:\"page,omitempty\"\`
    Limit int \`json:\"limit,omitempty\"\`
}
// ${capitalizedFileName}Resp 添加返回参数
type ${capitalizedFileName}Resp struct {
}
"

# model 结束

# router 开始
routerContent="package $(basename $router)

import(
	\"github.com/gofiber/fiber/v2\"
    \"shoollove/handler\"
)

// ${fileName} 路由，记得把 ${fileName}(app) 添加把 router.go 中
func ${fileName^}(app *fiber.App) {

	${fileName}Router := app.Group(\"/${fileName}\")
	{
		${fileName}Router.Post(\"/add\", handler.${capitalizedFileName}Add)
		${fileName}Router.Post(\"/edit\", handler.${capitalizedFileName}Edit)
		${fileName}Router.Post(\"/delete\", handler.${capitalizedFileName}Delete)
		${fileName}Router.Get(\"/get\", handler.${capitalizedFileName}Get)
		${fileName}Router.Get(\"/form\", handler.${capitalizedFileName}Form)
	}
}"
# router 结束

# service 开始
serviceContent="package $(basename $service)

import (
    \"context\"
    \"shoollove/model\"
    \"shoollove/pkg/db\"
    \"shoollove/pkg/utils\"
    \"shoollove/pkg/serr\"
)

// ${capitalizedFileName}Add 添加
func ${capitalizedFileName}Add(req model.${capitalizedFileName}) (resp any, err error) {

    dbBegin := db.Gorm.Begin()

	param := model.${capitalizedFileName}{}

	// var ${fileName} ${fileName}
	// 请求参数处理后添加到数据库
	// 赋值操作


    // 确保事务在函数结束时被正确处理
    defer func() {
        if r := recover(); r != nil {
            // 如果有 panic，回滚事务
            dbBegin.Rollback()
            // 处理错误（模拟 respond.Error 的处理）
            log.ZapSqlLog().Warn(context.Background(),err.Error()+\"添加失败，以回溯\")
            return
        } else {
            // 没有 panic，提交事务
            if err := dbBegin.Commit().Error; err != nil {
                // 提交失败的情况下处理错误
                log.ZapSqlLog().Warn(context.Background(), err.Error())
                return
            }
        }
    }()

	r := dbBegin.Create(&param)

	if r.Error != nil {
		panic(\"保存失败\")
	}
	return resp, nil
}

// ${capitalizedFileName}Edit 编辑
func ${capitalizedFileName}Edit(req model.${capitalizedFileName}) (resp any, err error) {
    session := db.Gorm

	param := model.${capitalizedFileName}{}

	// var ${fileName} ${fileName}
	// 请求参数处理后添加到数据库
	// 赋值操作
    err = session.Updates(&param).Error
    if err != nil {
        log.ZapSqlLog().Error(context.Background(), \"UserUpdate\", err.Error())
        return resp, serr.Serr{Code: serr.Fail, Msg: err.Error()}
    }
    return resp, nil
}

// ${capitalizedFileName}Delete 删除
func ${capitalizedFileName}Delete(req model.${capitalizedFileName}) (resp any, err error) {
    session := db.Gorm

	param := model.${capitalizedFileName}{}

	// var ${fileName} ${fileName}
	// 请求参数处理后添加到数据库
	// 赋值操作
	err = session.Delete(&param).Error
	if err != nil {
		log.ZapSqlLog().Error(context.Background(), \"UserDelete\", err.Error())
		return resp, serr.Serr{Code: serr.Fail, Msg: err.Error()}
	}
    return resp, nil
}


// ${capitalizedFileName}Get 查询
func ${capitalizedFileName}Get(req model.${capitalizedFileName}) (resp any, err error) {
    session := db.Gorm

	param := model.${capitalizedFileName}{}

	// var ${fileName} ${fileName}
	// 请求参数处理后添加到数据库
	// 赋值操作
	result, err := utils.IPage(req.CurrentPage, req.PageSize, session, param)
	if err != nil {
		log.ZapSqlLog().Error(context.Background(), \"UserDelete\", err.Error())
		return resp, serr.Serr{Code: serr.Fail, Msg: err.Error()}
	}
	resp.Result = result
    return resp, nil
}


// ${capitalizedFileName}Form 查询
func ${capitalizedFileName}Form(req model.${capitalizedFileName}) (resp any, err error) {
    session := db.Gorm

	param := model.${capitalizedFileName}{}

	// var ${fileName} ${fileName}
	// 请求参数处理后添加到数据库
	// 赋值操作
	err = session.Take(&param, req.Id).Error
	if err != nil {
		log.ZapSqlLog().Error(context.Background(), \"UserDelete\", err.Error())
		return resp, serr.Serr{Code: serr.Fail, Msg: err.Error()}
	}
    return resp, nil
}
"
# service 结束

# 在目录中创建一个main.go文件并写入内容
echo "$handlerContent" > "$handler/$fileName.go"
echo "$modelContent"   > "$model/$fileName.go"
echo "$routerContent"  > "$router/$fileName.go"
echo "$serviceContent" > "$service/$fileName.go"