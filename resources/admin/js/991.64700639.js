"use strict";(self["webpackChunkhalo_admin"]=self["webpackChunkhalo_admin"]||[]).push([[991],{90991:function(t,e,n){n.r(e),n.d(e,{default:function(){return x}});var s=function(){var t=this,e=t._self._c;return e("exception-page",{attrs:{type:"404"}})},a=[],o=function(){var t=this,e=t._self._c;return e("div",{staticClass:"exception"},[e("a-result",{attrs:{status:t.type,subTitle:t.config[t.type].desc,title:t.type},scopedSlots:t._u([{key:"extra",fn:function(){return[e("a-button",{attrs:{type:"primary"},on:{click:t.handleToHome}},[t._v("返回仪表盘")])]},proxy:!0}])})],1)},r=[];n(70560);const u={404:{desc:"抱歉，你访问的页面不存在"},500:{desc:"抱歉，服务器出错了"}};var l=u,c={name:"Exception",props:{type:{type:String,default:"404"}},data(){return{config:l}},methods:{handleToHome(){this.$router.push({name:"Dashboard"})}}},i=c,p=n(1001),d=(0,p.Z)(i,o,r,!1,null,null,null),f=d.exports,h={components:{ExceptionPage:f}},y=h,m=(0,p.Z)(y,s,a,!1,null,null,null),x=m.exports}}]);