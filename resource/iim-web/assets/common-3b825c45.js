var a=Object.defineProperty;var c=(e,t,i)=>t in e?a(e,t,{enumerable:!0,configurable:!0,writable:!0,value:i}):e[t]=i;var o=(e,t,i)=>(c(e,typeof t!="symbol"?t+"":t,i),i);import{p as r}from"./index-7540fc5d.js";class h{constructor(t,i=60,s=m=>{}){o(this,"timer",null);o(this,"lockTime",60);o(this,"lockName","");o(this,"callBack",()=>{});this.lockTime=i,this.lockName=`SMSLOCK_${t}`,this.callBack=s,this.compute()}start(){localStorage.setItem(this.lockName,this.getCurrentTime()+this.lockTime),this.compute()}end(){this.callBack(0),localStorage.removeItem(this.lockName)}compute(){this.clear();const t=this.getExpireTime();if(t===null)return;if(t<=this.getCurrentTime()){this.callBack(0),localStorage.removeItem(this.lockName);return}const i=t-this.getCurrentTime();this.callBack(i),this.timer=setTimeout(()=>this.compute(),1e3)}getCurrentTime(){return Math.floor(new Date().getTime()/1e3)}getExpireTime(){return localStorage.getItem(this.lockName)}clear(){clearTimeout(this.timer)}}const u=e=>/^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((.[a-zA-Z0-9_-]{2,3}){1,2})$/.test(e),k=e=>/^1[0-9]{10}$/.test(e),g=e=>r("/api/v1/common/sms-code",e),p=e=>r("/api/v1/common/email-code",e);export{h as S,g as a,p as b,u as c,k as i};
