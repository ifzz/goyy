// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesCoreCommPage = `<script type="text/javascript">
	function {%.param%}Page(options){
		var option={};
		if(arguments.length>0&&typeof(arguments[0])!="object"){
			var os=["pageNo","pageSize"];
			for(var i=0;i<os.length;i++){
				if(arguments.length>i&&$.isNotBlank(arguments[i])){
					option[os[i]]=arguments[i];
				}else{
					break;
				}
			}
		}else{
			option=options;
		}
		var obj=this.object;
		if(typeof(obj)=="object"){
			obj.args=$.extend(obj.args,option);
		}else{
			this.object=new Object();
			obj=this.object;
			//默认参数
			obj.args=$.extend({
				pageNoName		:"sPageNoTR", 					//分页当前页 input
				pageSizeName	:"sPageSizeTR", 				//分页条数 input
				formId			:"{%.param%}ListSform",		//ajaxSubmit 提交的表单 form
				alertId			:"{%.param%}ListAlert",		//提示信息id div
				dataId			:"{%.param%}ListData", 		//放模版内容的id div
				templateId		:"{%.param%}ListTemplate",	//模版id script
				pageId			:"{%.param%}ListPagination",	//分页标签的id ul
			}, option);
		}
		var form=$("#"+obj.args.formId);
		if($.isNotBlank(obj.args.pageNo)){
			form.find("input[name=\""+obj.args.pageNoName+"\"]").val(obj.args.pageNo);
		}
		if($.isNotBlank(obj.args.pageSize)){
			form.find("input[name=\""+obj.args.pageSizeName+"\"]").val(obj.args.pageSize);
		}
		$("#"+obj.args.alertId).hideAlert();
		if(typeof({%.param%}PrePage)=="function"){
			try{
				{%.param%}PrePage;
			}catch(e){
				console.log("{%.param%}PrePage err:",e);
			}
		}
		form.ajaxSubmit({
			dataType : "json",
			success : function(result) {
				if (result.success) {
					if(typeof({%.param%}Breadcrumb)=="function"){//如果存在导行函数
						try{
							{%.param%}Breadcrumb(result.tag)//将父节点数据传入导行生成函数
						}catch(e){
							console.log("{%.param%}Breadcrumb err:",e);
						}
					}
					var data=result.data;
					if($.isBlank(data)){
						console.log("result data is null",data);
						return;
					}
					var winUrl = $.url();
					for(var key in winUrl.param()){
						data[key]=winUrl.param()[key];
					}
					$("#"+obj.args.dataId).handlebars(obj.args.templateId,result.data);
					$("#"+obj.args.pageId).pagination({
						pageNo: result.data.pageNo,
						pageSize: result.data.pageSize,
						pageFunc: "{%.param%}Page",
						totalPages: result.data.totalPages,
						totalElements: result.data.totalElements
					});
					$("#"+obj.args.dataId).dict();//查询字典
					if(typeof({%.param%}PostPage)=="function"){
						try{
							{%.param%}PostPage(result);
						}catch(e){
							console.log("{%.param%}PostPage err:",e);
						}
					}
				} else {
					alert(result.message);
				}
			}
		});
	}
	$(function(){
		$("#{%.param%}ListSform").permission();
		if("{%.loadable%}"!="false"){
			{%.param%}Page(1);
		}
		$("#{%.param%}ListSbtn").click(function(){
			{%.param%}Page(1);
		});
	});
</script>
`
