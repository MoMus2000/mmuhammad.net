const _0xc33883=_0x2e15;(function(_0x2351d6,_0x4237e0){const _0x2b3935=_0x2e15,_0x54e48f=_0x2351d6();while(!![]){try{const _0x2addee=-parseInt(_0x2b3935(0xe6))/0x1*(-parseInt(_0x2b3935(0xdf))/0x2)+parseInt(_0x2b3935(0xe3))/0x3+-parseInt(_0x2b3935(0xe9))/0x4+parseInt(_0x2b3935(0xc7))/0x5+-parseInt(_0x2b3935(0xd5))/0x6+parseInt(_0x2b3935(0xda))/0x7*(parseInt(_0x2b3935(0xc8))/0x8)+-parseInt(_0x2b3935(0xcd))/0x9;if(_0x2addee===_0x4237e0)break;else _0x54e48f['push'](_0x54e48f['shift']());}catch(_0x218087){_0x54e48f['push'](_0x54e48f['shift']());}}}(_0xdb92,0xe589e));function _0xdb92(){const _0x46d329=['\x22\x20alt=\x22Card\x20image\x20cap\x22>\x0a\x20\x20<div\x20class=\x22card-body\x22>\x0a\x20\x20\x20\x20<h5\x20class=\x22card-title\x22>','</h5>\x0a\x20\x20\x20\x20<p\x20class=\x22card-text\x22>','offset','createElement','2626962RCISdC','click','Next','/api/v1/postByCat?cid=','\x22\x20class=\x22btn\x20btn-primary\x22>Go\x20to\x20the\x20article</a>\x0a\x20\x20</div>\x0a\x20\x20</div>\x0a\x20\x20<br>','7gsFpDI','&offset=','/articles?cid=','block','\x0a\x20\x20<div\x20class=\x22card\x20col-sm-6\x22>\x0a\x20\x20<img\x20class=\x22card-img-top\x22\x20src=\x22','49506xiDrZm','display','appendChild','cid','400701XZbZaG','get','getElementById','19lkGlJq','log','div','5546616HYUAvh','posts','location','href','innerHTML','6277250wYbsxb','10330504ZPdNwz','</p>\x0a\x20\x20\x20\x20<a\x20href=\x22/posts/','className','json','row','3473694DqPNVp','addEventListener','search','onload'];_0xdb92=function(){return _0x46d329;};return _0xdb92();}let currentPage=0x1;const params=new Proxy(new URLSearchParams(window[_0xc33883(0xc4)][_0xc33883(0xcf)]),{'get':(_0x2f1ab6,_0x5a7732)=>_0x2f1ab6[_0xc33883(0xe4)](_0x5a7732)});let cidValue=params[_0xc33883(0xe2)],offset=parseInt(params[_0xc33883(0xd3)]),apiRequestURL=_0xc33883(0xd8)+cidValue+_0xc33883(0xdb)+offset,apiRequestURLNext='/api/v1/postByCat?cid='+cidValue+_0xc33883(0xdb)+(offset+0x4);console[_0xc33883(0xe7)](apiRequestURL);async function getArticlesInCat(){const _0x3e47d8=_0xc33883;let _0x1fa794=await fetch(apiRequestURL),_0x54c6e5=await _0x1fa794[_0x3e47d8(0xcb)](),_0x4b6209=await fetch(apiRequestURLNext),_0x151c82=await _0x4b6209['json']();nextButton=document[_0x3e47d8(0xe5)](_0x3e47d8(0xd7));_0x151c82['length']!=0x0&&(nextButton['style'][_0x3e47d8(0xe0)]=_0x3e47d8(0xdd));console[_0x3e47d8(0xe7)](_0x54c6e5),html='';let _0x13ecb0='',_0x1b0c2a='',_0x30ad0e='',_0x505745='',_0x3fa7f1=0x0;for(_0x3fa7f1=0x0;_0x3fa7f1<_0x54c6e5['length'];_0x3fa7f1++){_0x13ecb0=_0x54c6e5[_0x3fa7f1][0x0],_0x1b0c2a=_0x54c6e5[_0x3fa7f1][0x1],_0x30ad0e=_0x54c6e5[_0x3fa7f1][0x2],_0x505745=_0x54c6e5[_0x3fa7f1][0x4],html+=createPost(_0x13ecb0,_0x1b0c2a,_0x505745,_0x30ad0e),(_0x3fa7f1+0x1)%0x2==0x0?(g=document[_0x3e47d8(0xd4)](_0x3e47d8(0xe8)),g['id']=_0x13ecb0,g[_0x3e47d8(0xca)]=_0x3e47d8(0xcc),postTag=document[_0x3e47d8(0xe5)](_0x3e47d8(0xc3)),postTag[_0x3e47d8(0xe1)](g),document[_0x3e47d8(0xe5)](g['id'])[_0x3e47d8(0xc6)]=html):html=createPost(_0x13ecb0,_0x1b0c2a,_0x505745,_0x30ad0e);}_0x3fa7f1+0x1%0x2!=0x0&&(g=document['createElement'](_0x3e47d8(0xe8)),g['id']=_0x13ecb0,g[_0x3e47d8(0xca)]=_0x3e47d8(0xcc),postTag=document[_0x3e47d8(0xe5)]('posts'),postTag['appendChild'](g),document['getElementById'](g['id'])[_0x3e47d8(0xc6)]=html);}function _0x2e15(_0x216281,_0x43f1a9){const _0xdb9204=_0xdb92();return _0x2e15=function(_0x2e1570,_0x1808d7){_0x2e1570=_0x2e1570-0xc3;let _0x5709c4=_0xdb9204[_0x2e1570];return _0x5709c4;},_0x2e15(_0x216281,_0x43f1a9);}function createPost(_0x3bd959,_0x285366,_0x59ac8f,_0x31dd3b){const _0x4d8c1c=_0xc33883;return html=_0x4d8c1c(0xde)+_0x31dd3b+_0x4d8c1c(0xd1)+_0x3bd959+_0x4d8c1c(0xd2)+_0x285366+_0x4d8c1c(0xc9)+_0x59ac8f+'/'+_0x3bd959+_0x4d8c1c(0xd9),html;}getArticlesInCat(),window[_0xc33883(0xd0)]=function(){const _0x3ec60f=_0xc33883;nextButton=document[_0x3ec60f(0xe5)](_0x3ec60f(0xd7)),nextButton[_0x3ec60f(0xce)](_0x3ec60f(0xd6),function(){const _0x1c80b5=_0x3ec60f;let _0xe7f505=params['offset'];console[_0x1c80b5(0xe7)](_0x1c80b5(0xd6)),_0xe7f505=parseInt(_0xe7f505),_0xe7f505+=0x4,window['location'][_0x1c80b5(0xc5)]=_0x1c80b5(0xdc)+cidValue+'&offset='+_0xe7f505;});};