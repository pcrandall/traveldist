import{p as e,w as t,d as a,N as s,a as i,b as r}from"./index-a3afd6e1.js";(()=>{e.t=t.__cssshim;const r=Array.from(a.querySelectorAll("script")).find((e=>new RegExp(`/${s}(\\.esm)?\\.js($|\\?|#)`).test(e.src)||e.getAttribute("data-stencil-namespace")===s)),n={};return n.resourcesUrl=new URL(".",new URL(r.getAttribute("data-resources-url")||r.src,t.location.href)).href,((e,i)=>{const r=`__sc_import_${s.replace(/\s|-/g,"_")}`;try{t[r]=new Function("w",`return import(w);//${Math.random()}`)}catch(s){const n=new Map;t[r]=s=>{const o=new URL(s,e).href;let c=n.get(o);if(!c){const e=a.createElement("script");e.type="module",e.crossOrigin=i.crossOrigin,e.src=URL.createObjectURL(new Blob([`import * as m from '${o}'; window.${r}.m = m;`],{type:"application/javascript"})),c=new Promise((a=>{e.onload=()=>{a(t[r].m),e.remove()}})),n.set(o,c),a.head.appendChild(e)}return c}}})(n.resourcesUrl,r),t.customElements?i(n):__sc_import_duet("./dom-fb6a473e.js").then((()=>n))})().then((e=>r([["duet-date-picker",[[0,"duet-date-picker",{name:[1],identifier:[1],disabled:[516],role:[1],direction:[1],required:[4],value:[513],min:[1],max:[1],firstDayOfWeek:[2,"first-day-of-week"],localization:[16],dateAdapter:[16],activeFocus:[32],focusedDay:[32],open:[32],setFocus:[64],show:[64],hide:[64]},[[6,"click","handleDocumentClick"]]]]]],e)));
