import{r as f,d as C,f as $,h as P,w as l,i as n,c as e,C as _,j as t,ah as q,aX as y,aD as S,a$ as x,s as U,b1 as B,e as E,q as M,v as N,t as V,b2 as D,b3 as R,b as F,aU as A,k as L,F as T,x as G,y as H}from"./index-9f644312.js";import{S as I,i as j,a as O,b as X}from"./common-78499294.js";const z={style:{width:"100%","text-align":"right"}},J={__name:"EditorPassword",props:["modelValue"],emits:["update:modelValue","close"],setup(g,{emit:h}){const c=f(),s=C({oldPassword:"",newPassword:"",newPassword2:""}),v={oldPassword:{required:!0,trigger:["blur","input"],message:"登录密码不能为空"},newPassword:{required:!0,trigger:["blur","input"],message:"新密码不能为空"},newPassword2:{required:!0,trigger:["blur","change"],validator(o,d){if(d){if(s.newPassword!=s.newPassword2)return new Error("两次密码填写不一致")}else return new Error("确认密码不能为空");return!0}}},u=f(!1),b=()=>{u.value=!0;let o=B({old_password:s.oldPassword,new_password:s.newPassword});o.then(d=>{d.code==200?(window.$message.success("密码修改成功"),h("update:modelValue",!1)):window.$message.warning(d.message)}),o.finally(()=>{u.value=!1})},w=o=>{o.preventDefault(),c.value.validate(d=>{!d&&b()})};return(o,d)=>($(),P(t(U),{show:g.modelValue,preset:"card",title:"修改密码",class:"modal-radius",style:{"max-width":"400px"},"on-update:show":r=>{o.$emit("update:modelValue",r)}},{footer:l(()=>[n("div",z,[e(t(q),{type:"tertiary",onClick:d[3]||(d[3]=r=>o.$emit("update:modelValue",!1))},{default:l(()=>[_(" 取消 ")]),_:1}),e(t(q),{type:"primary",class:"mt-l15",loading:u.value,onClick:w},{default:l(()=>[_(" 保存修改 ")]),_:1},8,["loading"])])]),default:l(()=>[e(t(x),{ref_key:"formRef",ref:c,model:s,rules:v},{default:l(()=>[e(t(y),{label:"登录密码",path:"oldPassword"},{default:l(()=>[e(t(S),{placeholder:"请填写登录密码",type:"password",value:s.oldPassword,"onUpdate:value":d[0]||(d[0]=r=>s.oldPassword=r)},null,8,["value"])]),_:1}),e(t(y),{label:"设置新密码",path:"newPassword"},{default:l(()=>[e(t(S),{placeholder:"请填写新密码",type:"password",value:s.newPassword,"onUpdate:value":d[1]||(d[1]=r=>s.newPassword=r)},null,8,["value"])]),_:1}),e(t(y),{label:"确认新密码",path:"newPassword2"},{default:l(()=>[e(t(S),{placeholder:"请再次填写新密码",type:"password",value:s.newPassword2,"onUpdate:value":d[2]||(d[2]=r=>s.newPassword2=r)},null,8,["value"])]),_:1})]),_:1},8,["model"])]),_:1},8,["show","on-update:show"]))}},K={style:{width:"100%","text-align":"right"}},Q={__name:"EditorMobile",props:["modelValue"],emits:["update:modelValue","success"],setup(g,{emit:h}){const c=f(),s=C({password:"",mobile:"",code:""}),v={password:{required:!0,trigger:["input"],message:"账号密码不能为空"},mobile:{required:!0,trigger:["input"],message:"手机号不能为空"},code:{required:!0,trigger:["change"],message:"验证码不能为空"}},u=f(!1),b=f(0),w=new I("CHANGE_MOBILE_SMS",60,i=>b.value=i),o=()=>{if(!j(s.mobile)){window.$message.warning("请正确填写手机号");return}O({mobile:s.mobile,channel:"change_mobile"}).then(({code:a,data:p,message:m})=>{a==200?(w.start(),p.is_debug?(s.code=p.code,window.$message.success("已开启验证码自动填充")):window.$message.success("验证码发送成功")):window.$message.warning(m)})},d=()=>{u.value=!0;let i=D(s);i.then(({code:a,message:p})=>{a==200?(window.$message.success("手机号修改成功"),h("success",s.mobile)):window.$message.warning(p)}),i.finally(()=>{u.value=!1})},r=i=>{i.preventDefault(),c.value.validate(a=>{!a&&d()})};return(i,a)=>{const p=E("n-button");return $(),P(t(U),{show:g.modelValue,preset:"card",title:"换绑手机号",class:"modal-radius",style:{"max-width":"400px"},"on-update:show":m=>{i.$emit("update:modelValue",m)}},{footer:l(()=>[n("div",K,[e(p,{type:"tertiary",onClick:a[3]||(a[3]=m=>i.$emit("update:modelValue",!1))},{default:l(()=>[_(" 取消 ")]),_:1}),e(p,{type:"primary",class:"mt-l15",loading:u.value,onClick:r},{default:l(()=>[_(" 保存修改 ")]),_:1},8,["loading"])])]),default:l(()=>[e(t(x),{ref_key:"formRef",ref:c,model:s,rules:v},{default:l(()=>[e(t(y),{label:"登录密码",path:"password"},{default:l(()=>[e(t(S),{placeholder:"请填写登录密码",type:"password",value:s.password,"onUpdate:value":a[0]||(a[0]=m=>s.password=m)},null,8,["value"])]),_:1}),e(t(y),{label:"新手机号",path:"mobile"},{default:l(()=>[e(t(S),{placeholder:"请填写新手机号",type:"text",value:s.mobile,"onUpdate:value":a[1]||(a[1]=m=>s.mobile=m)},null,8,["value"])]),_:1}),e(t(y),{label:"短信验证码",path:"code"},{default:l(()=>[e(t(S),{placeholder:"请填写验证码",type:"text",value:s.code,"onUpdate:value":a[2]||(a[2]=m=>s.code=m)},null,8,["value"]),e(p,{tertiary:"",class:"mt-l5",onClick:o,disabled:b.value>0},{default:l(()=>[_(" 获取验证码 "),M(n("span",null,"("+V(b.value)+"s)",513),[[N,b.value>0]])]),_:1},8,["disabled"])]),_:1})]),_:1},8,["model"])]),_:1},8,["show","on-update:show"])}}},W={style:{width:"100%","text-align":"right"}},Y={__name:"EditorEmail",props:["modelValue"],emits:["update:modelValue","success"],setup(g,{emit:h}){const c=f(0),s=new I("CHANGE_EMAIL_SMS",60,i=>c.value=i),v=f(),u=C({password:"",email:"",code:""}),b={password:{required:!0,trigger:["input"],message:"账号密码不能为空"},email:{required:!0,trigger:["input"],message:"邮箱不能为空"},code:{required:!0,trigger:["change"],message:"验证码不能为空"}},w=f(!1),o=()=>{X({email:u.email,channel:"change_email"}).then(({code:a,message:p})=>{a==200?(s.start(),window.$message.success("验证码发送成功")):window.$message.warning(p)})},d=()=>{w.value=!0;let i=R(u);i.then(({code:a,message:p})=>{a==200?(window.$message.success("邮箱修改成功"),h("success",u.email)):window.$message.warning(p)}),i.finally(()=>{w.value=!1})},r=i=>{i.preventDefault(),v.value.validate(a=>{!a&&d()})};return(i,a)=>{const p=E("n-button");return $(),P(t(U),{show:g.modelValue,preset:"card",title:"换绑邮箱",class:"modal-radius",style:{"max-width":"400px"},"on-update:show":m=>{i.$emit("update:modelValue",m)}},{footer:l(()=>[n("div",W,[e(p,{type:"tertiary",onClick:a[3]||(a[3]=m=>i.$emit("update:modelValue",!1))},{default:l(()=>[_(" 取消 ")]),_:1}),e(p,{type:"primary",class:"mt-l15",loading:w.value,onClick:r},{default:l(()=>[_(" 保存修改 ")]),_:1},8,["loading"])])]),default:l(()=>[e(t(x),{ref_key:"formRef",ref:v,model:u,rules:b},{default:l(()=>[e(t(y),{label:"登录密码",path:"password"},{default:l(()=>[e(t(S),{placeholder:"请填写登录密码",type:"password",value:u.password,"onUpdate:value":a[0]||(a[0]=m=>u.password=m)},null,8,["value"])]),_:1}),e(t(y),{label:"新邮箱",path:"email"},{default:l(()=>[e(t(S),{placeholder:"请填写新邮箱",type:"text",value:u.email,"onUpdate:value":a[1]||(a[1]=m=>u.email=m)},null,8,["value"])]),_:1}),e(t(y),{label:"邮箱验证码",path:"code"},{default:l(()=>[e(t(S),{placeholder:"请填写验证码",type:"text",value:u.code,"onUpdate:value":a[2]||(a[2]=m=>u.code=m)},null,8,["value"]),e(p,{tertiary:"",class:"mt-l5",onClick:o,disabled:c.value>0},{default:l(()=>[_(" 获取验证码 "),M(n("span",null,"("+V(c.value)+"s)",513),[[N,c.value>0]])]),_:1},8,["disabled"])]),_:1})]),_:1},8,["model"])]),_:1},8,["show","on-update:show"])}}};const k=g=>(G("data-v-aeb89c11"),g=g(),H(),g),Z={class:"container"},ee=k(()=>n("h3",{class:"title"},"安全设置",-1)),se={class:"view-box"},ae={class:"view-list"},le=k(()=>n("div",{class:"content"},[n("div",{class:"name"},"账户密码"),n("div",{class:"desc"},"当前密码强度 ：中")],-1)),te={class:"tools"},oe={class:"view-list"},de={class:"content"},ne=k(()=>n("div",{class:"name"},"绑定邮箱",-1)),re={class:"desc"},ie={class:"tools"},ue={class:"view-list"},me={class:"content"},pe=k(()=>n("div",{class:"name"},"绑定手机号",-1)),ce={class:"desc"},we={class:"tools"},ve={__name:"security",setup(g){const h=A(),c=f(!1),s=f(!1),v=f(!1),u=w=>{s.value=!1,h.mobile=w},b=w=>{v.value=!1,h.email=w};return(w,o)=>{const d=E("n-button");return $(),L(T,null,[n("section",Z,[ee,n("div",se,[n("div",ae,[le,n("div",te,[e(d,{type:"primary",text:"",onClick:o[0]||(o[0]=r=>c.value=!0)},{default:l(()=>[_(" 修改 ")]),_:1})])]),n("div",oe,[n("div",de,[ne,n("div",re,"已绑定邮箱 ："+V(t(h).email||"未绑定"),1)]),n("div",ie,[e(d,{type:"primary",text:"",onClick:o[1]||(o[1]=r=>v.value=!0)},{default:l(()=>[_(" 修改 ")]),_:1})])]),n("div",ue,[n("div",me,[pe,n("div",ce,"已绑定手机号 ："+V(t(h).mobile||"未绑定"),1)]),n("div",we,[e(d,{type:"primary",text:"",onClick:o[2]||(o[2]=r=>s.value=!0)},{default:l(()=>[_(" 修改 ")]),_:1})])])])]),e(J,{modelValue:c.value,"onUpdate:modelValue":o[3]||(o[3]=r=>c.value=r)},null,8,["modelValue"]),e(Q,{modelValue:s.value,"onUpdate:modelValue":o[4]||(o[4]=r=>s.value=r),onSuccess:u},null,8,["modelValue"]),e(Y,{modelValue:v.value,"onUpdate:modelValue":o[5]||(o[5]=r=>v.value=r),onSuccess:b},null,8,["modelValue"])],64)}}},ge=F(ve,[["__scopeId","data-v-aeb89c11"]]);export{ge as default};
