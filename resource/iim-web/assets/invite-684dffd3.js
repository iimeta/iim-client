import{b as d,a2 as v,d as u,r as p,e as m,f as o,k as l,i as e,t as s,c as h,w as b,F as f,l as x,cg as y,C as g,x as k,y as w}from"./index-8b922204.js";import{c as C}from"./vip-9348872c.js";const I=i=>(k("data-v-02fb2a46"),i=i(),w(),i),S={class:"el-container is-vertical height100"},F={class:"invite"},B={class:"inviteUrl"},N={class:"tools"},U={class:"inviteUrl mt10"},V={class:"table el-main scroller me-scrollbar me-scrollbar-thumb"},L=I(()=>e("div",{class:"theader"},[e("div",null,"昵称"),e("div",null,"邮箱"),e("div",null,"注册时间")],-1)),R={__name:"invite",setup(i){v();const t=u({invite_url:"",items:p([])});C().then(a=>{if(a.code==200){t.invite_url=a.data.invite_url,t.items=a.data.items||[];let c="https://iim.ai";t.invite_url=c+t.invite_url}});const r=()=>y(t.invite_url,()=>window.$message.success("已复制"));return(a,c)=>{const _=m("n-button");return o(),l("section",S,[e("div",F,[e("span",B,"邀请链接: "+s(t.invite_url),1),e("span",N,[h(_,{type:"primary",text:"",onClick:r},{default:b(()=>[g(" 复制 ")]),_:1})]),e("p",U,"邀请人数: "+s(t.items.length),1)]),e("div",V,[L,(o(!0),l(f,null,x(t.items,n=>(o(),l("div",{class:"row",key:n.email},[e("div",null,s(n.nickname),1),e("div",null,s(n.email),1),e("div",null,s(n.created_at),1)]))),128))])])}}},T=d(R,[["__scopeId","data-v-02fb2a46"]]);export{T as default};