"use strict";(self["webpackChunkblog"]=self["webpackChunkblog"]||[]).push([[495],{10011:(a,t,e)=>{e.d(t,{x4:()=>s,a_:()=>o,P1:()=>l,IU:()=>n,Ai:()=>p,cc:()=>r,jC:()=>c,pk:()=>d,fq:()=>u,qM:()=>m,nQ:()=>w,Ir:()=>g,sh:()=>y,uc:()=>b,bb:()=>f,Ix:()=>_,Yq:()=>k});var i=e(11768);const s=async a=>await i.api.post("/api/login/users",a),o=async a=>await i.api.get("/api/login/checktoken",{params:a}),l=async a=>await i.api.post("/api/login/sendvalidator",a),n=async a=>await i.api.post("/api/blog/signup",a),p=async a=>await i.api.get(`/api/blog/profile/${a}`),r=async()=>await i.api.get("/api/blog/getme"),c=async a=>await i.api.get("/api/blog/articlelist",{params:a}),d=async a=>await i.api.get("/api/blog/category",{params:a}),u=async a=>await i.api.get(`/api/blog/article/${a}`),m=async a=>await i.api.post("/api/blog/article",a),w=async(a,t)=>await i.api.get(`/api/blog/comment/list/${a}`,{params:t}),g=async a=>await i.api.post("/api/blog/comment",a),y=async a=>await i.api.post("/api/blog/comment/digg",a),b=async a=>await i.api.post("/api/blog/comment/undigg",a),f=async a=>await i.api["delete"](`/api/blog/article/${a}`),_=async a=>{const t={headers:{"Content-Type":"multipart/form-data"}};return await i.api.post("/api/blog/upload",a,t)},k=async a=>await i.api.post("/api/blog/emoji",a)},57495:(a,t,e)=>{e.r(t),e.d(t,{default:()=>B});var i=e(83673);const s=a=>((0,i.dD)("data-v-685851f3"),a=a(),(0,i.Cn)(),a),o=(0,i.Uk)("blogName"),l={class:"row justify-center"},n=(0,i.Uk)(" OOPSIES! OUR BAD "),p=(0,i.Uk)(" This Set-up Link is Bad "),r=s((()=>(0,i._)("div",{align:"center",class:"text-subtitle2"}," Hmm, this account set up link isn't valid. ",-1)));function c(a,t,e,s,c,d){const u=(0,i.up)("router-link"),m=(0,i.up)("q-toolbar-title"),w=(0,i.up)("q-toolbar"),g=(0,i.up)("q-header"),y=(0,i.up)("q-icon"),b=(0,i.up)("q-input"),f=(0,i.up)("q-btn"),_=(0,i.up)("q-form"),k=(0,i.up)("q-card-section"),q=(0,i.up)("q-card"),h=(0,i.up)("q-item-label"),v=(0,i.up)("q-item-section"),W=(0,i.up)("q-item"),Z=(0,i.up)("q-card-actions"),Q=(0,i.up)("q-page"),P=(0,i.up)("q-page-container"),x=(0,i.up)("q-layout");return(0,i.wg)(),(0,i.j4)(x,null,{default:(0,i.w5)((()=>[(0,i.Wm)(g,null,{default:(0,i.w5)((()=>[(0,i.Wm)(w,null,{default:(0,i.w5)((()=>[(0,i.Wm)(m,{shrink:"",class:"text-weight-bold"},{default:(0,i.w5)((()=>[(0,i.Wm)(u,{to:"/",style:{color:"#fff"}},{default:(0,i.w5)((()=>[o])),_:1})])),_:1})])),_:1})])),_:1}),(0,i.Wm)(P,{class:"bg-primary"},{default:(0,i.w5)((()=>[(0,i._)("div",l,[(0,i.Wm)(Q,{padding:""},{default:(0,i.w5)((()=>[a.valid?((0,i.wg)(),(0,i.j4)(q,{key:0,style:{width:"380px"}},{default:(0,i.w5)((()=>[(0,i.Wm)(k,null,{default:(0,i.w5)((()=>[(0,i.Wm)(_,{onSubmit:a.onSubmit},{default:(0,i.w5)((()=>[(0,i.Wm)(b,{outlined:"",dense:"","bg-color":"grey-2",modelValue:a.formdata.password,"onUpdate:modelValue":t[1]||(t[1]=t=>a.formdata.password=t),type:a.isPwd?"password":"text",label:"password","lazy-rules":"",rules:[a=>!!a||"password is required"]},{append:(0,i.w5)((()=>[(0,i.Wm)(y,{name:a.isPwd?"visibility_off":"visibility",class:"cursor-pointer",onClick:t[0]||(t[0]=t=>a.isPwd=!a.isPwd)},null,8,["name"])])),_:1},8,["modelValue","type","rules"]),(0,i.Wm)(b,{outlined:"",dense:"","bg-color":"grey-2",modelValue:a.formdata.comfirm_password,"onUpdate:modelValue":t[3]||(t[3]=t=>a.formdata.comfirm_password=t),type:a.isPwd?"password":"text",label:"comfirm password","lazy-rules":"",rules:[a=>!!a||"comfirm password is required",t=>t==a.formdata.password||"Make sure the password twice"]},{append:(0,i.w5)((()=>[(0,i.Wm)(y,{name:a.isPwd?"visibility_off":"visibility",class:"cursor-pointer",onClick:t[2]||(t[2]=t=>a.isPwd=!a.isPwd)},null,8,["name"])])),_:1},8,["modelValue","type","rules"]),(0,i.Wm)(f,{class:"full-width",style:{"padding-top":"0"},label:"sign up",type:"submit",color:"primary"})])),_:1},8,["onSubmit"])])),_:1})])),_:1})):((0,i.wg)(),(0,i.j4)(q,{key:1,class:"q-pa-lg",style:{width:"380px"}},{default:(0,i.w5)((()=>[(0,i.Wm)(W,{align:"center"},{default:(0,i.w5)((()=>[(0,i.Wm)(v,null,{default:(0,i.w5)((()=>[(0,i.Wm)(h,{class:"text-weight-bold text-red-6"},{default:(0,i.w5)((()=>[n])),_:1}),(0,i.Wm)(h,{class:"text-h6"},{default:(0,i.w5)((()=>[p])),_:1})])),_:1})])),_:1}),(0,i.Wm)(Z,{align:"center"},{default:(0,i.w5)((()=>[r])),_:1})])),_:1}))])),_:1})])])),_:1})])),_:1})}var d=e(61959),u=e(79582),m=e(10011),w=e(64434);const g={name:"Login",components:{},setup(){const a=(0,u.tv)(),t=(0,u.yj)();let e=(0,d.qj)({isPwd:!0,valid:!0,formdata:{password:""}});const{ctx:s}=(0,i.FN)();console.log(s);const o={async onSubmit(){const i={token:t.query.token,password:e.formdata.password};let s=await(0,m.IU)(i);200==s.status&&(w.Z.create({message:"Sign Up Successful",color:"positive",icon:"report_problem",position:"top",timeout:3e3}),a.push({path:"/login",query:{login_redirect:"claim"}}))},async checkToken(){const a={token:t.query.token};let i=await(0,m.a_)(a);200==i.status&&(e.valid=!0)}};return(0,i.wF)((async()=>{await o.checkToken()})),{...(0,d.BK)(e),...o}}};var y=e(74260),b=e(19214),f=e(53812),_=e(39570),k=e(13747),q=e(72652),h=e(24379),v=e(10151),W=e(25589),Z=e(68689),Q=e(64689),P=e(24554),x=e(2165),C=e(83414),I=e(52035),S=e(2350),U=e(99367),j=e(7518),T=e.n(j);const V=(0,y.Z)(g,[["render",c],["__scopeId","data-v-685851f3"]]),B=V;T()(g,"components",{QLayout:b.Z,QHeader:f.Z,QToolbar:_.Z,QToolbarTitle:k.Z,QPageContainer:q.Z,QPage:h.Z,QCard:v.Z,QCardSection:W.Z,QForm:Z.Z,QInput:Q.Z,QIcon:P.Z,QBtn:x.Z,QItem:C.Z,QItemSection:I.Z,QItemLabel:S.Z,QCardActions:U.Z})}}]);