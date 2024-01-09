"use strict";(self["webpackChunkhalo_admin"]=self["webpackChunkhalo_admin"]||[]).push([[574],{70297:function(e,t,a){a.d(t,{Z:function(){return d}});var r=function(){var e=this,t=e._self._c;return t("a-tree-select",{attrs:{allowClear:!0,treeData:e.categoryTreeData,treeDataSimpleMode:!0,placeholder:"请选择上级目录，默认为顶级目录",treeDefaultExpandAll:"","dropdown-style":{overflow:"auto"}},on:{change:e.handleChange},model:{value:e.categoryIdString,callback:function(t){e.categoryIdString=t},expression:"categoryIdString"}})},o=[],l=(a(70560),{name:"CategorySelectTree",props:{categoryId:{type:Number,required:!0,default:0},categories:{type:Array,required:!1,default:()=>[]},root:{type:Object,required:!1,default:()=>({id:0,title:"根目录",value:"0",pId:-1})}},computed:{categoryTreeData(){return[this.root,...this.convertDataToTree(this.categories)]},categoryIdString:{get(){return this.categoryId.toString()},set(e){this.$emit("update:categoryId",e?parseInt(e):0)}}},methods:{handleChange(){this.$emit("change")},convertDataToTree(e){const t={},a=[];e.forEach((e=>t[e.id]={...e,title:e.name,value:e.id.toString(),pId:e.parentId,children:[]})),e.forEach((e=>{const r=t[e.id];e.parentId?t[e.parentId].children.push(r):a.push(r)}));const r=(e,t=!1)=>{e.forEach((e=>{e.hasPassword=!!e.password||t,e.hasPassword&&(e.title=`${e.title}（加密）`),Object.hasOwn(e,"postCount")&&(e.title=`${e.title} - ${e.postCount} 篇`),e.children&&e.children.length&&r(e.children,e.hasPassword)}))};return r(a),a}}}),s=l,i=a(1001),n=(0,i.Z)(s,r,o,!1,null,null,null),d=n.exports},20574:function(e,t,a){a.d(t,{Z:function(){return N}});var r=function(){var e=this,t=e._self._c;return t("a-modal",{attrs:{afterClose:e.onClosed,bodyStyle:{padding:0},maskClosable:!1,width:680,destroyOnClose:""},scopedSlots:e._u([{key:"title",fn:function(){return[e._v(" "+e._s(e.modalTitle)+" "),e.loading?t("a-icon",{attrs:{type:"loading"}}):e._e()]},proxy:!0},{key:"footer",fn:function(){return[e._t("extraFooter"),e.draftSaveVisible?t("ReactiveButton",{attrs:{errored:e.form.draftSaveErrored,loading:e.form.draftSaving,text:(e.hasId?"转为":"保存")+"草稿",erroredText:"保存失败",loadedText:"保存成功",type:"danger"},on:{callback:function(t){return e.handleSavedCallback()},click:function(t){return e.handleSaveDraft()}}}):e._e(),e.publishVisible?t("ReactiveButton",{attrs:{errored:e.form.publishErrored,loading:e.form.publishing,erroredText:"发布失败",loadedText:"发布成功",text:"转为发布"},on:{callback:function(t){return e.handleSavedCallback()},click:function(t){return e.handlePublish()}}}):e._e(),t("ReactiveButton",{attrs:{errored:e.form.saveErrored,erroredText:(e.hasId?"保存":"发布")+"失败",loadedText:(e.hasId?"保存":"发布")+"成功",loading:e.form.saving,text:""+(e.hasId?"保存":"发布")},on:{callback:function(t){return e.handleSavedCallback()},click:function(t){return e.handleSave()}}}),t("a-button",{attrs:{disabled:e.loading},on:{click:function(t){e.modalVisible=!1}}},[e._v("关闭")])]},proxy:!0}],null,!0),model:{value:e.modalVisible,callback:function(t){e.modalVisible=t},expression:"modalVisible"}},[t("div",{staticClass:"card-container"},[t("a-tabs",{attrs:{type:"card"}},[t("a-tab-pane",{key:"normal",attrs:{tab:"常规"}},[t("a-form",{attrs:{"label-col":e.form.labelCol,"wrapper-col":e.form.wrapperCol,labelAlign:"left"}},[t("a-form-item",{attrs:{label:"文章标题"}},[t("a-input",{model:{value:e.form.model.title,callback:function(t){e.$set(e.form.model,"title",t)},expression:"form.model.title"}})],1),t("a-form-item",{attrs:{help:e.fullPath,label:"文章别名"}},[t("a-input",{scopedSlots:e._u([{key:"addonAfter",fn:function(){return[t("a-popconfirm",{attrs:{"cancel-text":"取消","ok-text":"确定",placement:"left",title:"是否确定根据标题重新生成别名？"},on:{confirm:e.handleGenerateSlug}},[t("a-icon",{staticClass:"cursor-pointer",attrs:{type:"sync"}})],1)]},proxy:!0}]),model:{value:e.form.model.slug,callback:function(t){e.$set(e.form.model,"slug",t)},expression:"form.model.slug"}})],1),t("a-form-item",{attrs:{label:"分类目录"}},[t("a-space",{attrs:{direction:"vertical"}},[t("category-tree",{ref:"categoryTree",model:{value:e.form.model.categoryIds,callback:function(t){e.$set(e.form.model,"categoryIds",t)},expression:"form.model.categoryIds"}}),t("a-button",{attrs:{type:"dashed"},on:{click:function(t){e.categoryCreateModalVisible=!0}}},[e._v("新增")])],1)],1),t("a-form-item",{attrs:{label:"标签"}},[t("TagSelect",{model:{value:e.form.model.tagIds,callback:function(t){e.$set(e.form.model,"tagIds",t)},expression:"form.model.tagIds"}})],1),t("a-form-item",{attrs:{label:"摘要"}},[t("a-input",{attrs:{autoSize:{minRows:5},placeholder:"如不填写，会从文章中自动截取",type:"textarea"},model:{value:e.form.model.summary,callback:function(t){e.$set(e.form.model,"summary",t)},expression:"form.model.summary"}})],1)],1)],1),t("a-tab-pane",{key:"advanced",attrs:{tab:"高级"}},[t("a-form",{attrs:{"label-col":e.form.labelCol,"wrapper-col":e.form.wrapperCol,labelAlign:"left"}},[t("a-form-item",{attrs:{label:"禁止评论"}},[t("a-switch",{model:{value:e.form.model.disallowComment,callback:function(t){e.$set(e.form.model,"disallowComment",t)},expression:"form.model.disallowComment"}})],1),t("a-form-item",{attrs:{label:"是否置顶"}},[t("a-switch",{model:{value:e.topPriority,callback:function(t){e.topPriority=t},expression:"topPriority"}})],1),t("a-form-item",{attrs:{label:"发表时间："}},[t("a-date-picker",{attrs:{defaultValue:e.createTimeDefaultValue,format:"YYYY-MM-DD HH:mm:ss",placeholder:"选择文章发表时间",showTime:""},on:{change:e.onCreateTimeSelect,ok:e.onCreateTimeSelect}})],1),t("a-form-item",{attrs:{label:"自定义模板："}},[t("a-select",{model:{value:e.form.model.template,callback:function(t){e.$set(e.form.model,"template",t)},expression:"form.model.template"}},[t("a-select-option",{key:"",attrs:{value:""}},[e._v("无")]),e._l(e.templates,(function(a){return t("a-select-option",{key:a,attrs:{value:a}},[e._v(" "+e._s(a)+" ")])}))],2)],1),t("a-form-item",{attrs:{label:"访问密码："}},[t("a-input-password",{attrs:{autocomplete:"new-password"},model:{value:e.form.model.password,callback:function(t){e.$set(e.form.model,"password",t)},expression:"form.model.password"}})],1),t("a-form-item",{attrs:{label:"封面图："}},[t("a-space",{attrs:{direction:"vertical"}},[t("img",{staticClass:"w-1/2 cursor-pointer",staticStyle:{"border-radius":"4px"},attrs:{src:e.form.model.thumbnail||"/images/placeholder.jpg",alt:"Post cover thumbnail"},on:{click:function(t){e.attachmentSelectVisible=!0}}}),t("a-input",{attrs:{"allow-clear":"",placeholder:"点击封面图选择图片，或者输入外部链接"},model:{value:e.form.model.thumbnail,callback:function(t){e.$set(e.form.model,"thumbnail",t)},expression:"form.model.thumbnail"}})],1)],1)],1)],1),t("a-tab-pane",{key:"seo",attrs:{tab:"SEO"}},[t("a-form",{attrs:{"label-col":e.form.labelCol,"wrapper-col":e.form.wrapperCol,labelAlign:"left"}},[t("a-form-item",{attrs:{label:"自定义关键词"}},[t("a-input",{attrs:{autoSize:{minRows:5},placeholder:"多个关键词以英文逗号隔开，如不填写，将自动使用标签作为关键词",type:"textarea"},model:{value:e.form.model.metaKeywords,callback:function(t){e.$set(e.form.model,"metaKeywords",t)},expression:"form.model.metaKeywords"}})],1),t("a-form-item",{attrs:{label:"自定义描述"}},[t("a-input",{attrs:{autoSize:{minRows:5},placeholder:"如不填写，会从文章中自动截取",type:"textarea"},model:{value:e.form.model.metaDescription,callback:function(t){e.$set(e.form.model,"metaDescription",t)},expression:"form.model.metaDescription"}})],1)],1)],1),t("a-tab-pane",{key:"meta",attrs:{tab:"元数据"}},[t("MetaEditor",{attrs:{metas:e.form.model.metas,targetId:e.form.model.id,target:"post"},on:{"update:metas":function(t){return e.$set(e.form.model,"metas",t)}}})],1)],1)],1),t("AttachmentSelectModal",{attrs:{multiSelect:!1,visible:e.attachmentSelectVisible},on:{"update:visible":function(t){e.attachmentSelectVisible=t},confirm:e.handleSelectPostThumbnail}}),t("CategoryCreateModal",{attrs:{visible:e.categoryCreateModalVisible},on:{"update:visible":function(t){e.categoryCreateModalVisible=t},close:e.onCategoryCreateModalClose}})],1)},o=[],l=function(){var e=this,t=e._self._c;return e.categories.loading?e._e():t("a-tree",{attrs:{checkedKeys:e.categoryIds,treeData:e.categories.data,checkStrictly:"",checkable:"",defaultExpandAll:"",showLine:""},on:{"update:treeData":function(t){return e.$set(e.categories,"data",t)},"update:tree-data":function(t){return e.$set(e.categories,"data",t)},check:e.onCheck}})},s=[],i=(a(70560),a(80068)),n={name:"CategoryTree",model:{prop:"categoryIds",event:"check"},props:{categoryIds:{type:Array,required:!1,default:()=>[]}},data(){return{categories:{data:[],loading:!1}}},created(){this.handleListCategories()},methods:{async handleListCategories(){try{this.categories.loading=!0;const{data:e}=await i.Z.category.list({sort:[],more:!1});this.categories.data=this.convertDataToTree(e)}catch(e){this.$log.error(e)}finally{this.categories.loading=!1}},convertDataToTree(e){const t={},a=[];e.forEach((e=>t[e.id]={...e,key:e.id,title:e.name,children:[]})),e.forEach((e=>{const r=t[e.id];e.parentId?t[e.parentId].children.push(r):a.push(r)}));const r=(e,t=!1)=>{e.forEach((e=>{e.hasPassword=!!e.password||t,e.hasPassword&&(e.title=`${e.title}（加密）`),e.children&&e.children.length&&r(e.children,e.hasPassword)}))};return r(a),a},onCheck(e,t){this.$log.debug("Chekced keys",e),this.$log.debug("e",t),this.$emit("check",e.checked)}}},d=n,m=a(1001),c=(0,m.Z)(d,l,s,!1,null,null,null),h=c.exports,u=function(){var e=this,t=e._self._c;return t("a-select",{staticClass:"w-full",attrs:{"token-separators":[",","|"],allowClear:"",mode:"tags",placeholder:"选择或输入标签"},on:{change:e.handleChange},model:{value:e.selectedTagNames,callback:function(t){e.selectedTagNames=t},expression:"selectedTagNames"}},e._l(e.tags,(function(a){return t("a-select-option",{key:a.id,attrs:{value:a.name}},[e._v(e._s(a.name))])})),1)},p=[],f=a(9669),g=a.n(f),b={name:"TagSelect",model:{prop:"tagIds",event:"change"},props:{tagIds:{type:Array,required:!1,default:()=>[]}},data(){return{tags:[],selectedTagNames:[]}},created(){this.handleListTags()},watch:{tags(e){e&&(this.selectedTagNames=this.tagIds.map((e=>this.tagIdMap[e].name)))},tagIds:{handler(e){this.tags.length&&(this.selectedTagNames=e.map((e=>this.tagIdMap[e].name)))},deep:!0}},computed:{tagIdMap(){const e={};return this.tags.forEach((t=>{e[t.id]=t})),e},tagNameMap(){const e={};return this.tags.forEach((t=>{e[t.name]=t})),e}},methods:{handleListTags(e){i.Z.tag.list({sort:"name,asc",more:!0}).then((t=>{this.tags=t.data,e&&e()}))},handleChange(){const e=this.selectedTagNames.filter((e=>!this.tagNameMap[e]));if(this.$log.debug("Tag names to create",e),!e.length){const e=this.selectedTagNames.map((e=>this.tagNameMap[e].id));return void this.$emit("change",e)}const t=e.map((e=>i.Z.tag.create({name:e})));g().all(t).then(g().spread((()=>{this.handleListTags((()=>{this.$log.debug("Tag name map",this.tagNameMap);const e=this.selectedTagNames.map((e=>this.tagNameMap[e].id));this.$emit("change",e)}))})))}}},y=b,v=(0,m.Z)(y,u,p,!1,null,null,null),C=v.exports,w=a(8641),k=function(){var e=this,t=e._self._c;return t("a-modal",{attrs:{afterClose:e.onClose,width:512,destroyOnClose:"",title:"新建分类"},scopedSlots:e._u([{key:"footer",fn:function(){return[t("ReactiveButton",{attrs:{errored:e.form.errored,loading:e.form.saving,erroredText:"保存失败",loadedText:"保存成功",text:"保存",type:"primary"},on:{callback:e.handleSavedCallback,click:e.handleCreate}}),t("a-button",{on:{click:function(t){e.modalVisible=!1}}},[e._v("关闭")])]},proxy:!0}]),model:{value:e.modalVisible,callback:function(t){e.modalVisible=t},expression:"modalVisible"}},[t("a-form-model",{ref:"categoryForm",attrs:{"label-col":e.form.labelCol,model:e.form.model,rules:e.form.rules,"wrapper-col":e.form.wrapperCol,labelAlign:"left"}},[t("a-form-model-item",{attrs:{help:"* 页面上所显示的名称",label:"名称：",prop:"name"}},[t("a-input",{ref:"nameInput",model:{value:e.form.model.name,callback:function(t){e.$set(e.form.model,"name",t)},expression:"form.model.name"}})],1),t("a-form-model-item",{attrs:{help:"* 一般为单个分类页面的标识，最好为英文",label:"别名：",prop:"slug"}},[t("a-input",{model:{value:e.form.model.slug,callback:function(t){e.$set(e.form.model,"slug",t)},expression:"form.model.slug"}})],1),t("a-form-model-item",{attrs:{label:"上级目录：",prop:"parentId"}},[t("category-select-tree",{attrs:{categories:e.list.data,"category-id":e.form.model.parentId},on:{"update:categoryId":function(t){return e.$set(e.form.model,"parentId",t)},"update:category-id":function(t){return e.$set(e.form.model,"parentId",t)}}})],1),t("a-form-model-item",{attrs:{help:"* 在分类页面可展示，需要主题支持",label:"封面图：",prop:"thumbnail"}},[t("AttachmentInput",{attrs:{title:"选择封面图"},model:{value:e.form.model.thumbnail,callback:function(t){e.$set(e.form.model,"thumbnail",t)},expression:"form.model.thumbnail"}})],1),t("a-form-model-item",{attrs:{help:"* 分类密码",label:"密码：",prop:"password"}},[t("a-input-password",{attrs:{autocomplete:"new-password"},model:{value:e.form.model.password,callback:function(t){e.$set(e.form.model,"password",t)},expression:"form.model.password"}})],1),t("a-form-model-item",{attrs:{help:"* 分类描述，需要主题支持",label:"描述：",prop:"description"}},[t("a-input",{attrs:{autoSize:{minRows:3},type:"textarea"},model:{value:e.form.model.description,callback:function(t){e.$set(e.form.model,"description",t)},expression:"form.model.description"}})],1)],1)],1)},$=[],S=a(70297),T={name:"CategoryCreateModal",components:{CategorySelectTree:S.Z},props:{visible:{type:Boolean,default:!1}},data(){return{list:{data:[],loading:!1},form:{model:{},saving:!1,errored:!1,rules:{name:[{required:!0,message:"* 分类名称不能为空",trigger:["change"]},{max:255,message:"* 分类名称的字符长度不能超过 255",trigger:["change"]}],slug:[{max:255,message:"* 分类别名的字符长度不能超过 255",trigger:["change"]}],thumbnail:[{max:1023,message:"* 封面图链接的字符长度不能超过 1023",trigger:["change"]}],description:[{max:100,message:"* 分类描述的字符长度不能超过 100",trigger:["change"]}]},labelCol:{sm:{span:4},xs:{span:24}},wrapperCol:{sm:{span:20},xs:{span:24}}}}},computed:{modalVisible:{get(){return this.visible},set(e){this.$emit("update:visible",e)}}},created(){this.handleListCategories()},watch:{modalVisible(e){e&&this.$nextTick((()=>{this.$refs.nameInput.focus()}))}},methods:{async handleListCategories(){try{this.list.loading=!0;const{data:e}=await i.Z.category.list({});this.list.data=e}catch(e){this.$log.error("Failed to get categories",e)}finally{this.list.loading=!1}},handleCreate(){this.$refs.categoryForm.validate((async e=>{if(e)try{this.form.saving=!0,await i.Z.category.create(this.form.model)}catch(t){this.form.errored=!0,this.$log.error("Failed to create category",t)}finally{setTimeout((()=>{this.form.saving=!1}),400)}}))},handleSavedCallback(){this.form.errored?this.form.errored=!1:(this.form.model={},this.handleListCategories())},onClose(){this.$emit("close")}}},x=T,I=(0,m.Z)(x,k,$,!1,null,null,null),_=I.exports,D=a(83428),E=a(64203),M=a(80981),V=a.n(M),Y=a(20629),P=a(32422),A={name:"PostSettingModal",mixins:[D.jB,D.KT],components:{CategoryTree:h,TagSelect:C,MetaEditor:w.Z,CategoryCreateModal:_},props:{visible:{type:Boolean,default:!1},loading:{type:Boolean,default:!1},post:{type:Object,default:()=>({})},savedCallback:{type:Function,default:null}},data(){return{postStatuses:P.JQ,form:{model:{},saving:!1,saveErrored:!1,draftSaving:!1,draftSaveErrored:!1,publishing:!1,publishErrored:!1,labelCol:{sm:{span:4},xs:{span:24}},wrapperCol:{sm:{span:20},xs:{span:24}}},templates:[],attachmentSelectVisible:!1,categoryCreateModalVisible:!1}},computed:{...(0,Y.Se)(["options"]),modalVisible:{get(){return this.visible},set(e){this.$emit("update:visible",e)}},modalTitle(){return this.form.model.id?"文章设置":"文章发布"},createTimeDefaultValue(){if(this.form.model.createTime){const e=new Date(this.form.model.createTime);return(0,E._)(e,"YYYY-MM-DD HH:mm:ss")}return(0,E._)(new Date,"YYYY-MM-DD HH:mm:ss")},topPriority:{get(){return void 0!==this.form.model.topPriority&&0!==this.form.model.topPriority},set(e){this.$set(this.form.model,"topPriority",e?1:0)}},fullPath(){const{post_permalink_type:e,archives_prefix:t,blog_url:a,path_suffix:r=""}=this.options,{slug:o="{slug}",createTime:l=new Date,id:s="{id}"}=this.form.model;switch(e){case"DEFAULT":return`${a}/${t}/${o}${r}`;case"YEAR":return`${a}${(0,E._)(l,"/YYYY/")}${o}${r}`;case"DATE":return`${a}${(0,E._)(l,"/YYYY/MM/")}${o}${r}`;case"DAY":return`${a}${(0,E._)(l,"/YYYY/MM/DD/")}${o}${r}`;case"ID":return`${a}/?p=${s}`;case"ID_SLUG":return`${a}/${t}/${s}${r}`;default:return""}},hasId(){return!!this.form.model.id},draftSaveVisible(){const{draftSaving:e,publishing:t}=this.form;return(this.form.model.status!==P.JQ.DRAFT.value||!this.hasId||e)&&!t},publishVisible(){const{draftSaving:e,publishing:t}=this.form;return(this.form.model.status===P.JQ.DRAFT.value&&this.hasId||t)&&!e}},watch:{modalVisible(e){e&&(this.form.model=Object.assign({},this.post),this.form.model.slug||this.form.model.id||this.handleGenerateSlug())},post:{deep:!0,handler(e){this.form.model=Object.assign({},e)}}},created(){this.handleListCustomTemplates()},methods:{async handleCreateOrUpdate(){if(!this.form.model.title)throw this.$notification["error"]({message:"提示",description:"文章标题不能为空！"}),new Error("文章标题不能为空！");this.form.model.keepRaw=!0;try{this.hasId?await i.Z.post.update(this.form.model.id,this.form.model):await i.Z.post.create(this.form.model)}catch(e){throw this.$log.error(e),new Error(e)}},async handleSave(){try{this.form.saving=!0;const{status:e}=this.form.model;e||(this.form.model.status=this.postStatuses.PUBLISHED.value),await this.handleCreateOrUpdate()}catch(e){this.form.saveErrored=!0,this.$log.error("Failed to save post",e)}finally{setTimeout((()=>{this.form.saving=!1}),400)}},async handlePublish(){try{this.form.publishing=!0,this.form.model.status=this.postStatuses.PUBLISHED.value,await this.handleCreateOrUpdate()}catch(e){this.form.publishErrored=!0,this.$log.error("Failed to publish post",e)}finally{setTimeout((()=>{this.form.publishing=!1}),400)}},async handleSaveDraft(){try{this.form.draftSaving=!0,this.form.model.status=this.postStatuses.DRAFT.value,await this.handleCreateOrUpdate()}catch(e){this.form.draftSaveErrored=!0,this.$log.error("Failed to save draft post",e)}finally{setTimeout((()=>{this.form.draftSaving=!1}),400)}},handleSavedCallback(){this.form.saveErrored||this.form.draftSaveErrored||this.form.publishErrored?(this.form.saveErrored=!1,this.form.draftSaveErrored=!1,this.form.publishErrored=!1):this.savedCallback&&this.savedCallback()},async handleListCustomTemplates(){try{const e=await i.Z.theme.listCustomPostTemplates();this.templates=e.data}catch(e){this.$log.error(e)}},onCreateTimeSelect(e){this.form.model.createTime=e.valueOf()},handleGenerateSlug(){if(this.form.model.title&&V().isSupported()){let e="";const t=V().parse(this.form.model.title.replace(/\s+/g,"").toLowerCase());let a;t.forEach((t=>{if(2===t.type){const r=t.target?t.target.toLowerCase():"";e+=e&&!/\n|\s/.test(a.target)?"-"+r:r}else e+=(a&&2===a.type?"-":"")+t.target;a=t})),this.$set(this.form.model,"slug",e)}},handleSelectPostThumbnail({raw:e}){e.length&&(this.form.model.thumbnail=e[0].path),this.attachmentSelectVisible=!1},onClosed(){this.$emit("onClose"),this.$emit("onUpdate",this.form.model)},onCategoryCreateModalClose(){this.$refs.categoryTree.handleListCategories()}}},L=A,Z=(0,m.Z)(L,r,o,!1,null,null,null),N=Z.exports}}]);