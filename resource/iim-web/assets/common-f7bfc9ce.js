var a=Object.defineProperty;var c=(e,t,i)=>t in e?a(e,t,{enumerable:!0,configurable:!0,writable:!0,value:i}):e[t]=i;var r=(e,t,i)=>(c(e,typeof t!="symbol"?t+"":t,i),i);import{p as o}from"./index-9dec9867.js";class u{constructor(t,i=60,s=m=>{}){r(this,"timer",null);r(this,"lockTime",60);r(this,"lockName","");r(this,"callBack",()=>{});this.lockTime=i,this.lockName=`SMSLOCK_${t}`,this.callBack=s,this.compute()}start(){localStorage.setItem(this.lockName,this.getCurrentTime()+this.lockTime),this.compute()}compute(){this.clear();const t=this.getExpireTime();if(t===null)return;if(t<=this.getCurrentTime()){this.callBack(0),localStorage.removeItem(this.lockName);return}const i=t-this.getCurrentTime();this.callBack(i),this.timer=setTimeout(()=>this.compute(),1e3)}getCurrentTime(){return Math.floor(new Date().getTime()/1e3)}getExpireTime(){return localStorage.getItem(this.lockName)}clear(){clearTimeout(this.timer)}}const h=e=>/^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((.[a-zA-Z0-9_-]{2,3}){1,2})$/.test(e),k=e=>/^1[0-9]{10}$/.test(e),p=e=>o("/api/v1/common/sms-code",e),T=e=>o("/api/v1/common/email-code",e);export{u as S,p as a,T as b,h as c,k as i};
