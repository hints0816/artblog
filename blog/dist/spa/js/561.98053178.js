"use strict";(self["webpackChunkblog"]=self["webpackChunkblog"]||[]).push([[561],{3561:(e,t,a)=>{a.r(t),a.d(t,{default:()=>B});var s=a(83673),l=a(62323);const o=e=>((0,s.dD)("data-v-54f1d72f"),e=e(),(0,s.Cn)(),e),i=(0,s.Uk)("ArtBlog"),r={class:"row justify-center"},d=(0,s.Uk)(" OOPSIES! OUR BAD "),n=(0,s.Uk)(" This Set-up Link is Bad "),u=o((()=>(0,s._)("div",{align:"center",class:"text-subtitle2"}," Hmm, this account set up link isn't valid. ",-1)));function c(e,t,a,o,c,p){const m=(0,s.up)("router-link"),w=(0,s.up)("q-toolbar-title"),f=(0,s.up)("q-toolbar"),b=(0,s.up)("q-header"),g=(0,s.up)("q-icon"),k=(0,s.up)("q-input"),y=(0,s.up)("q-btn"),_=(0,s.up)("q-form"),q=(0,s.up)("q-card-section"),v=(0,s.up)("q-card"),h=(0,s.up)("q-item-label"),W=(0,s.up)("q-item-section"),Z=(0,s.up)("q-item"),Q=(0,s.up)("q-card-actions"),P=(0,s.up)("q-page"),C=(0,s.up)("q-page-container"),S=(0,s.up)("q-layout");return(0,s.wg)(),(0,s.j4)(S,null,{default:(0,s.w5)((()=>[(0,s.Wm)(b,null,{default:(0,s.w5)((()=>[(0,s.Wm)(f,{class:(0,l.C_)(e.$q.dark.isActive?"bg-dark":"bg-primary")},{default:(0,s.w5)((()=>[(0,s.Wm)(w,{shrink:"",class:"text-weight-bold"},{default:(0,s.w5)((()=>[(0,s.Wm)(m,{to:"/",style:{color:"#fff"}},{default:(0,s.w5)((()=>[i])),_:1})])),_:1})])),_:1},8,["class"])])),_:1}),(0,s.Wm)(C,{class:(0,l.C_)(e.$q.dark.isActive?"bg-dark":"bg-primary")},{default:(0,s.w5)((()=>[(0,s._)("div",r,[(0,s.Wm)(P,{padding:""},{default:(0,s.w5)((()=>[e.valid?((0,s.wg)(),(0,s.j4)(v,{key:0,flat:"",style:{width:"380px"}},{default:(0,s.w5)((()=>[(0,s.Wm)(q,null,{default:(0,s.w5)((()=>[(0,s.Wm)(_,{onSubmit:e.onSubmit},{default:(0,s.w5)((()=>[(0,s.Wm)(k,{outlined:"",dense:"","bg-color":e.$q.dark.isActive?"dark":"grey-2",modelValue:e.formdata.password,"onUpdate:modelValue":t[1]||(t[1]=t=>e.formdata.password=t),type:e.isPwd?"password":"text",label:"password","lazy-rules":"",rules:[e=>!!e||"password is required"]},{append:(0,s.w5)((()=>[(0,s.Wm)(g,{name:e.isPwd?"visibility_off":"visibility",class:"cursor-pointer",onClick:t[0]||(t[0]=t=>e.isPwd=!e.isPwd)},null,8,["name"])])),_:1},8,["bg-color","modelValue","type","rules"]),(0,s.Wm)(k,{outlined:"",dense:"","bg-color":e.$q.dark.isActive?"dark":"grey-2",modelValue:e.formdata.comfirm_password,"onUpdate:modelValue":t[3]||(t[3]=t=>e.formdata.comfirm_password=t),type:e.isPwd?"password":"text",label:"comfirm password","lazy-rules":"",rules:[e=>!!e||"comfirm password is required",t=>t==e.formdata.password||"Make sure the password twice"]},{append:(0,s.w5)((()=>[(0,s.Wm)(g,{name:e.isPwd?"visibility_off":"visibility",class:"cursor-pointer",onClick:t[2]||(t[2]=t=>e.isPwd=!e.isPwd)},null,8,["name"])])),_:1},8,["bg-color","modelValue","type","rules"]),(0,s.Wm)(y,{class:"full-width",style:{"padding-top":"0"},label:"sign up",type:"submit",color:"primary"})])),_:1},8,["onSubmit"])])),_:1})])),_:1})):((0,s.wg)(),(0,s.j4)(v,{key:1,class:"q-pa-lg",style:{width:"380px"}},{default:(0,s.w5)((()=>[(0,s.Wm)(Z,{align:"center"},{default:(0,s.w5)((()=>[(0,s.Wm)(W,null,{default:(0,s.w5)((()=>[(0,s.Wm)(h,{class:"text-weight-bold text-red-6"},{default:(0,s.w5)((()=>[d])),_:1}),(0,s.Wm)(h,{class:"text-h6"},{default:(0,s.w5)((()=>[n])),_:1})])),_:1})])),_:1}),(0,s.Wm)(Q,{align:"center"},{default:(0,s.w5)((()=>[u])),_:1})])),_:1}))])),_:1})])])),_:1},8,["class"])])),_:1})}var p=a(61959),m=a(79582),w=a(10011),f=a(64434);const b={name:"Login",components:{},setup(){const e=(0,m.tv)(),t=(0,m.yj)();let a=(0,p.qj)({isPwd:!0,valid:!0,formdata:{password:""}});const{ctx:l}=(0,s.FN)(),o={async onSubmit(){const s={token:t.query.token,password:a.formdata.password};let l=await(0,w.IU)(s);200==l.status&&(f.Z.create({message:"Sign Up Successful",color:"positive",icon:"report_problem",position:"top",timeout:3e3}),e.push({path:"/login",query:{login_redirect:"claim"}}))},async checkToken(){const e={token:t.query.token};let s=await(0,w.a_)(e);200==s.status&&(a.valid=!0)}};return(0,s.wF)((async()=>{await o.checkToken()})),{...(0,p.BK)(a),...o}}};var g=a(74260),k=a(19214),y=a(53812),_=a(39570),q=a(13747),v=a(72652),h=a(24379),W=a(10151),Z=a(25589),Q=a(68689),P=a(64689),C=a(24554),S=a(2165),x=a(83414),I=a(52035),U=a(2350),A=a(99367),j=a(7518),T=a.n(j);const V=(0,g.Z)(b,[["render",c],["__scopeId","data-v-54f1d72f"]]),B=V;T()(b,"components",{QLayout:k.Z,QHeader:y.Z,QToolbar:_.Z,QToolbarTitle:q.Z,QPageContainer:v.Z,QPage:h.Z,QCard:W.Z,QCardSection:Z.Z,QForm:Q.Z,QInput:P.Z,QIcon:C.Z,QBtn:S.Z,QItem:x.Z,QItemSection:I.Z,QItemLabel:U.Z,QCardActions:A.Z})}}]);