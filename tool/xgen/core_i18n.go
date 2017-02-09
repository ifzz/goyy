// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

var i18N = i18n.NewByEnv(locales)

var locales = map[string]map[string]string{
	i18n.LocaleZhCN: map[string]string{
		"col.comment.id":                 "标识",
		"col.comment.memo":               "备注",
		"col.comment.creates":            "创建机构",
		"col.comment.creater":            "创建人员",
		"col.comment.created":            "创建时间",
		"col.comment.modifier":           "更新人员",
		"col.comment.modified":           "更新时间",
		"col.comment.version":            "乐观锁",
		"col.comment.deletion":           "删除标志",
		"col.comment.artifical":          "人造数据",
		"col.comment.history":            "历史数据",
		"col.comment.code":               "编号",
		"col.comment.name":               "名称",
		"col.comment.fullname":           "全称",
		"col.comment.genre":              "类型",
		"col.comment.leaf":               "是否叶子",
		"col.comment.grade":              "节点等级",
		"col.comment.ordinal":            "排序",
		"col.comment.parent_id":          "父表主键",
		"col.comment.parent_ids":         "父ID集",
		"col.comment.parent_codes":       "父编号集",
		"col.comment.parent_names":       "父名称集",
		"html.list.sid":                  "序号",
		"html.list.sbtn":                 "查询",
		"html.list.created":              "创建时间",
		"html.list.opr":                  "操作",
		"html.list.edit":                 "修改",
		"html.list.del":                  "删除",
		"html.form.save":                 "保存",
		"tmpl.form.combo.placeholder":    "请选择",
		"tmpl.form.submit.alert":         "保存成功",
		"tmpl.form.repwd.alert":          "修改密码成功！",
		"tmpl.form.area.province":        "请选择省",
		"tmpl.form.area.city":            "请选择市",
		"tmpl.form.area.district":        "请选择区",
		"tmpl.form.page.current":         "当前",
		"tmpl.form.page.row":             "条",
		"tmpl.form.page.total":           "共",
		"tmpl.form.amount.limit":         "超出最大处理数字",
		"tmpl.form.interval.time":        "秒后重新发送",
		"tmpl.note.form.valid":           "表单校验",
		"tmpl.note.post.init":            "初始化岗位树结构数据",
		"tmpl.note.role.init":            "初始化角色树结构数据",
		"tmpl.note.user.get":             "获取当前用户登录信息",
		"tmpl.note.captcha.reset":        "重置图片验证码",
		"tmpl.valid.login.name":          "请填写用户名",
		"tmpl.valid.login.passwd":        "请填写密码",
		"tmpl.valid.login.captcha":       "请填写验证码",
		"tmpl.home.title":                "用户",
		"tmpl.home.tab.info":             "个人信息",
		"tmpl.home.tab.passwd":           "修改密码",
		"tmpl.home.form.login":           "登录名：",
		"tmpl.home.form.name":            "名称：",
		"tmpl.home.form.time":            "登录时间：",
		"tmpl.home.form.oldpwd":          "*旧密码：",
		"tmpl.home.form.newpwd":          "*新密码：",
		"tmpl.home.form.okpwd":           "*确认新密码：",
		"tmpl.home.form.save":            "保存",
		"tmpl.home.form.goback":          "返回",
		"tmpl.login.title":               "登录",
		"tmpl.login.form.name":           "登录名：",
		"tmpl.login.form.pwd":            "密码：",
		"tmpl.login.form.captcha":        "验证码：",
		"tmpl.login.form.look":           "看不清",
		"tmpl.login.form.change":         "换一张",
		"tmpl.login.form.submit":         "登 录",
		"tmpl.core.comm.action.add":      "添加",
		"tmpl.core.comm.action.edit":     "修改",
		"tmpl.core.comm.action.show":     "查看",
		"tmpl.core.comm.dialog.cancel":   "取消",
		"tmpl.core.comm.dialog.ok":       "确定",
		"tmpl.core.comm.dialog.tips":     "系统提示",
		"tmpl.core.comm.disable.cancel":  "取消",
		"tmpl.core.comm.disable.ok":      "确定",
		"tmpl.core.comm.disable.tips":    "系统提示",
		"tmpl.core.comm.disable.ask":     "确认要删除该记录吗？",
		"tmpl.core.comm.formtree.cancel": "取消",
		"tmpl.core.comm.formtree.ok":     "确定",
		"tmpl.core.comm.formtree.select": "选择",
		"tmpl.core.comm.formtree.title":  "请选择",
		"tmpl.core.comm.navtabs.list":    "列表",
		"tmpl.core.comm.navtabs.add":     "添加",
		"tmpl.err.goback":                "返回上一页",
		"tmpl.err.401.title":             "无权访问页面",
		"tmpl.err.401.h1":                "您没有权限访问该页面！",
		"tmpl.err.403.title":             "访问被拒绝",
		"tmpl.err.403.h1":                "抱歉，访问被拒绝！",
		"tmpl.err.404.title":             "出错了",
		"tmpl.err.404.h1":                "抱歉，您访问的页面不存在！",
		"tmpl.err.500.title":             "出错了",
		"tmpl.err.500.h1":                "抱歉，您访问的页面出现了问题！",
		"tmpl.sys.col.id":                "序号",
		"tmpl.sys.col.name":              "名称",
		"tmpl.sys.col.name2":             "姓名",
		"tmpl.sys.col.passwd":            "密码",
		"tmpl.sys.col.email":             "邮箱",
		"tmpl.sys.col.mobile":            "手机",
		"tmpl.sys.col.tel":               "电话",
		"tmpl.sys.col.login.name":        "登录名称",
		"tmpl.sys.col.login.ip":          "登录IP",
		"tmpl.sys.col.login.time":        "登录时间",
		"tmpl.sys.col.reg.time":          "注册时间",
		"tmpl.sys.col.code":              "编号",
		"tmpl.sys.col.classify":          "业务分类",
		"tmpl.sys.col.genre":             "类型",
		"tmpl.sys.col.genre.sys":         "系统类型",
		"tmpl.sys.col.genre.like":        "类型（包含）",
		"tmpl.sys.col.genre.label":       "类型（命名规范：表名.字段名；比如：sys_user.genre）",
		"tmpl.sys.col.descr":             "描述",
		"tmpl.sys.col.mkey":              "键",
		"tmpl.sys.col.mval":              "值",
		"tmpl.sys.col.filters":           "过滤",
		"tmpl.sys.col.usable":            "是否可用",
		"tmpl.sys.col.href":              "链接",
		"tmpl.sys.col.target":            "目标",
		"tmpl.sys.col.icon":              "图标",
		"tmpl.sys.col.perm":              "权限",
		"tmpl.sys.col.hidden":            "隐藏",
		"tmpl.sys.col.content":           "内容",
		"tmpl.sys.col.fullname":          "全称",
		"tmpl.sys.col.parent":            "上级",
		"tmpl.sys.col.ordinal":           "排序",
		"tmpl.sys.col.memo":              "备注",
		"tmpl.sys.col.created":           "创建时间",
		"tmpl.sys.col.opt":               "操作",
		"tmpl.sys.col.area":              "所属区域",
		"tmpl.sys.col.org":               "所属机构",
		"tmpl.sys.col.menu":              "菜单",
		"tmpl.sys.col.role":              "角色",
		"tmpl.sys.btn.query":             "查询",
		"tmpl.sys.btn.edit":              "修改",
		"tmpl.sys.btn.disable":           "删除",
		"tmpl.sys.btn.show":              "查看",
		"tmpl.sys.btn.save":              "保存",
		"tmpl.sys.list.title":            "列表",
		"tmpl.sys.add.title":             "添加",
		"tmpl.sys.area.title":            "区域",
		"tmpl.sys.blacklist.title":       "黑名单",
		"tmpl.sys.cache.title":           "缓存",
		"tmpl.sys.conf.title":            "配置",
		"tmpl.sys.dict.title":            "字典",
		"tmpl.sys.menu.title":            "菜单",
		"tmpl.sys.org.title":             "机构",
		"tmpl.sys.post.title":            "岗位",
		"tmpl.sys.role.title":            "角色",
		"tmpl.sys.user.title":            "用户",
	},
	i18n.LocaleEnUS: map[string]string{
		"col.comment.id":                  "ID",
		"col.comment.memo":                "MEMO",
		"col.comment.creates":             "CREATES",
		"col.comment.creater":             "CREATER",
		"col.comment.created":             "CREATED",
		"col.comment.modifier":            "MODIFIER",
		"col.comment.modified":            "MODIFIED",
		"col.comment.version":             "VERSION",
		"col.comment.deletion":            "DELETION",
		"col.comment.artifical":           "ARTIFICAL",
		"col.comment.history":             "HISTORY",
		"col.comment.code":                "CODE",
		"col.comment.name":                "NAME",
		"col.comment.fullname":            "FULLNAME",
		"col.comment.genre":               "GENRE",
		"col.comment.leaf":                "LEAF",
		"col.comment.grade":               "GRADE",
		"col.comment.ordinal":             "ORDINAL",
		"col.comment.parent_id":           "PARENT_ID",
		"col.comment.parent_ids":          "PARENT_IDS",
		"col.comment.parent_codes":        "PARENT_CODES",
		"col.comment.parent_names":        "PARENT_NAMES",
		"html.list.sid":                   "id",
		"html.list.sbtn":                  "search",
		"html.list.created":               "created",
		"html.list.opr":                   "operate",
		"html.list.edit":                  "edit",
		"html.list.del":                   "delete",
		"html.form.save":                  "save",
		"tmpl.form.combo.placeholder":     "select",
		"tmpl.form.submit.alert":          "save success",
		"tmpl.form.repwd.alert":           "change password successfully!",
		"tmpl.form.area.province":         "Please select province",
		"tmpl.form.area.city":             "Please select city",
		"tmpl.form.area.district":         "Please select district",
		"tmpl.form.page.current":          "current",
		"tmpl.form.page.row":              "row",
		"tmpl.form.page.total":            "total",
		"tmpl.form.amount.limit":          "Maximum number exceeded",
		"tmpl.form.interval.time":         "seconds to resend",
		"tmpl.note.form.valid":            "form validation",
		"tmpl.note.post.init":             "initialize the tree structure data on the post",
		"tmpl.note.role.init":             "initialize the tree structure data on the role",
		"tmpl.note.user.get":              "get the information of the logged in user",
		"tmpl.note.captcha.reset":         "reset captcha",
		"tmpl.valid.login.name":           "Please fill in the user name",
		"tmpl.valid.login.passwd":         "Please fill in the password",
		"tmpl.valid.login.captcha":        "Please fill in the captcha",
		"tmpl.home.title":                 "USER",
		"tmpl.home.tab.info":              "User information",
		"tmpl.home.tab.passwd":            "Modify password",
		"tmpl.home.form.login":            "login name:",
		"tmpl.home.form.name":             "name:",
		"tmpl.home.form.time":             "login time:",
		"tmpl.home.form.oldpwd":           "*old password:",
		"tmpl.home.form.newpwd":           "*new password:",
		"tmpl.home.form.okpwd":            "*confirm new password:",
		"tmpl.home.form.save":             "save",
		"tmpl.home.form.goback":           "go back",
		"tmpl.login.title":                "LOGIN",
		"tmpl.login.form.name":            "login name:",
		"tmpl.login.form.pwd":             "password:",
		"tmpl.login.form.captcha":         "captcha:",
		"tmpl.login.form.look":            "invisibility ",
		"tmpl.login.form.change":          "change",
		"tmpl.login.form.submit":          "login",
		"tmpl.core.comm.action.add":       "add",
		"tmpl.core.comm.action.edit":      "edit",
		"tmpl.core.comm.action.show":      "show",
		"tmpl.core.comm.dialog.cancel":    "cancel",
		"tmpl.core.comm.dialog.ok":        "ok",
		"tmpl.core.comm.dialog.tips":      "System prompt",
		"tmpl.core.comm.disable.cancel":   "cancel",
		"tmpl.core.comm.disable.ok":       "ok",
		"tmpl.core.comm.disable.tips":     "System prompt",
		"tmpl.core.comm.disable.ask":      "Are you sure you want to delete this record?",
		"tmpl.core.comm.formtree.cancel":  "cancel",
		"tmpl.core.comm.formtree.ok":      "ok",
		"tmpl.core.comm.formtree.select":  "select",
		"tmpl.core.comm.formtree.pselect": "please select",
		"tmpl.core.comm.navtabs.list":     " list",
		"tmpl.core.comm.navtabs.add":      " add",
		"tmpl.err.goback":                 "return to previous page",
		"tmpl.err.401.title":              "Unauthorized access page",
		"tmpl.err.401.h1":                 "You do not have permission to access this page!",
		"tmpl.err.403.title":              "Access was denied",
		"tmpl.err.403.h1":                 "Sorry, access denied!",
		"tmpl.err.404.title":              "Error",
		"tmpl.err.404.h1":                 "Sorry, the page you visit does not exist!",
		"tmpl.err.500.title":              "Error",
		"tmpl.err.500.h1":                 "Sorry, there is a problem with the page you visited!",
		"tmpl.sys.col.id":                 "id",
		"tmpl.sys.col.name":               "name",
		"tmpl.sys.col.name2":              "name",
		"tmpl.sys.col.passwd":             "password",
		"tmpl.sys.col.email":              "email",
		"tmpl.sys.col.mobile":             "mobile",
		"tmpl.sys.col.tel":                "telephone",
		"tmpl.sys.col.login.name":         "login name",
		"tmpl.sys.col.login.ip":           "login IP",
		"tmpl.sys.col.login.time":         "login time",
		"tmpl.sys.col.reg.time":           "register time",
		"tmpl.sys.col.code":               "code",
		"tmpl.sys.col.classify":           "classify",
		"tmpl.sys.col.genre":              "genre",
		"tmpl.sys.col.genre.sys":          "system genre",
		"tmpl.sys.col.genre.like":         "genre(like)",
		"tmpl.sys.col.genre.label":        "genre(Naming specification: tableName.fieldName, eg: sys_user.genre)",
		"tmpl.sys.col.descr":              "describe",
		"tmpl.sys.col.mkey":               "map key",
		"tmpl.sys.col.mval":               "map value",
		"tmpl.sys.col.filters":            "filters",
		"tmpl.sys.col.usable":             "usable",
		"tmpl.sys.col.href":               "href",
		"tmpl.sys.col.target":             "target",
		"tmpl.sys.col.icon":               "icon",
		"tmpl.sys.col.perm":               "permission",
		"tmpl.sys.col.hidden":             "hidden",
		"tmpl.sys.col.content":            "content",
		"tmpl.sys.col.fullname":           "fullname",
		"tmpl.sys.col.parent":             "parent",
		"tmpl.sys.col.ordinal":            "ordinal",
		"tmpl.sys.col.memo":               "memo",
		"tmpl.sys.col.created":            "create time",
		"tmpl.sys.col.opt":                "operation",
		"tmpl.sys.col.area":               "area",
		"tmpl.sys.col.org":                "org",
		"tmpl.sys.col.menu":               "menu",
		"tmpl.sys.col.role":               "role",
		"tmpl.sys.btn.query":              "query",
		"tmpl.sys.btn.edit":               "edit",
		"tmpl.sys.btn.disable":            "delete",
		"tmpl.sys.btn.show":               "show",
		"tmpl.sys.btn.save":               "save",
		"tmpl.sys.list.title":             "list",
		"tmpl.sys.add.title":              "add",
		"tmpl.sys.area.title":             "area",
		"tmpl.sys.blacklist.title":        "blacklist",
		"tmpl.sys.cache.title":            "cache",
		"tmpl.sys.conf.title":             "conf",
		"tmpl.sys.dict.title":             "dictionary",
		"tmpl.sys.menu.title":             "menu",
		"tmpl.sys.org.title":              "org",
		"tmpl.sys.post.title":             "post",
		"tmpl.sys.role.title":             "role",
		"tmpl.sys.user.title":             "user",
	},
}
