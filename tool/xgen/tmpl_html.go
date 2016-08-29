// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplHtmlMain = `<%range $i, $e := .Entities%><!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="description" content="">
<meta name="author" content="">
<!--#settings project="<%$e.Project%>" module="<%$e.Name%>" title="<%$e.Comment%>"--><!--#endsettings-->
<!--#with prefix="e" --><!--#endwith-->
<title go:title="/title.html">{%title%}</title>
<!--#include file="/core/include/head.html"--><!--#endinclude-->

<script type="text/javascript" go:src="{%statics%}/js/{%project%}/{%module%}/{%module%}.js"></script>
</head>
<body>
<!--#include file="/core/include/header.html" param="{%project%}"--><!--#endinclude-->
<div class="container-fluid">
<div class="content" class="row-fluid">
	<!--#include file="/core/comm/page.html" param="{%.prefix%}"--><!--#endinclude-->
	<!--#include file="/core/comm/action.html" param="{%.prefix%}"--><!--#endinclude-->
	<!--#include file="/core/comm/disable.html" param="{%.prefix%}"--><!--#endinclude-->
	<div class="tabbable">
		<!--#include file="/core/comm/navtabs.list.html" param="{%.prefix%}"--><!--#endinclude-->
		<div class="tab-content">
			<div id="{%.prefix%}ListContent" class="tab-pane active">
				<br/>
				<div class="well well-sm">
					<form id="{%.prefix%}ListSform" class="form-inline" method="post" go:action="{%apis%}/{%project%}/{%module%}/index">
						<input type="hidden" id="sDeletionU0" name="sDeletionU0" value="0" />
						<input type="hidden" id="sPageNoTR" name="sPageNoTR" value="1" />
						<input type="hidden" id="sPageSizeTR" name="sPageSizeTR" value="" />
						<div class="form-group">
							<label class="control-label">序号</label>
							<input class="form-control" id="sIdEQ" name="sIdEQ">
						</div>
						<button type="button" id="{%.prefix%}ListSbtn" class="btn btn-primary">查询</button>
					</form>
				</div>
				<!--#include file="/core/comm/alert.html" param="{%.prefix%}List"--><!--#endinclude-->
				<table class="table table-bordered table-condensed">
					<thead>
						<tr>
							<th class="f-w50">序号</th><%range $f := $e.Fields%>
							<th><%$f.Comment%></th><%end%>
							<th class="f-w150">创建时间</th>
							<th>操作</th>
						</tr>
					</thead>
					<tbody id="{%.prefix%}ListData"></tbody>
					<script id="{%.prefix%}ListTemplate" type="text/x-handlebars-template">
					{{#each slice}}
						<tr{{#if (divisible @index 2)}} class="active"{{/if}}>
							<td title="{{id}}">{{abbr id 5}}</td><%range $f := $e.Fields%>
							<td>{{<%$f.Name%>}}</td><%end%>
							<td>{{uyymdhms created}}</td>
							<td>
								<a href="#disable" go:data-state="disable" go:data-url="{%apis%}/{%project%}/{%module%}/disable" onclick="disable(this,'{{id}}',{%.prefix%}Page)" go:data-permissions="{%project%}:{%module%}:disable">删除</a>
							</td>
						</tr>
					{{/each}}
					</script>
				</table>
                <ul id="{%.prefix%}ListPagination" class="pagination"></ul>
			</div>
		</div>
	</div>
</div>
<!--#include file="/core/include/footer.html"--><!--#endinclude-->
</div>
</body>
</html><%end%>
`
