const _0x4b9c2d=_0x3bc8;(function(_0x2c8ce6,_0x29e3e5){const _0x2a3615=_0x3bc8,_0x30762e=_0x2c8ce6();while(!![]){try{const _0xf01a95=parseInt(_0x2a3615(0x105))/0x1*(-parseInt(_0x2a3615(0x108))/0x2)+-parseInt(_0x2a3615(0xf6))/0x3+parseInt(_0x2a3615(0xfd))/0x4*(parseInt(_0x2a3615(0x113))/0x5)+-parseInt(_0x2a3615(0x112))/0x6+parseInt(_0x2a3615(0x100))/0x7*(-parseInt(_0x2a3615(0x109))/0x8)+-parseInt(_0x2a3615(0x10d))/0x9*(-parseInt(_0x2a3615(0x10a))/0xa)+parseInt(_0x2a3615(0xf7))/0xb*(parseInt(_0x2a3615(0xf4))/0xc);if(_0xf01a95===_0x29e3e5)break;else _0x30762e['push'](_0x30762e['shift']());}catch(_0x376b0c){_0x30762e['push'](_0x30762e['shift']());}}}(_0x6842,0xd48b5));let currentPage=0x1;const params=new Proxy(new URLSearchParams(window[_0x4b9c2d(0x115)][_0x4b9c2d(0x104)]),{'get':(_0x1be019,_0x1c22a5)=>_0x1be019[_0x4b9c2d(0xf5)](_0x1c22a5)});function _0x6842(){const _0x3b639a=['offset','19650ETcadB','1088UjcMKF','10dmaJsM','style','posts','8559369LgjqUU','\x0a\x20\x20<div\x20class=\x22card\x20col-sm-5\x20mt-4\x20mr-4\x22>\x0a\x20\x20<img\x20class=\x22card-img-top\x22\x20src=\x22','row','json','createElement','4663494TYozVD','15LeWLdj','getElementById','location','onload','addEventListener','innerHTML','click','490068ArfPSV','get','345363SUZNTT','407tzCrjM','Next','&offset=','display','log','\x22\x20class=\x22btn\x20btn-primary\x22>Go\x20to\x20the\x20article</a>\x0a\x20\x20</div>\x0a\x20\x20</div>\x0a\x20\x20<br>','949424ISLSLJ','div','appendChild','19537UVQXxw','</h5>\x0a\x20\x20\x20\x20<p\x20class=\x22card-text\x22>','length','className','search','105bUcPug','block'];_0x6842=function(){return _0x3b639a;};return _0x6842();}let cidValue=params['cid'],offset=parseInt(params[_0x4b9c2d(0x107)]),apiRequestURL='/api/v1/postByCat?cid='+cidValue+_0x4b9c2d(0xf9)+offset,apiRequestURLNext='/api/v1/postByCat?cid='+cidValue+_0x4b9c2d(0xf9)+(offset+0x4);console[_0x4b9c2d(0xfb)](apiRequestURL);async function getArticlesInCat(){const _0x1b1907=_0x4b9c2d;let _0x3a3a34=await fetch(apiRequestURL),_0x9ca94d=await _0x3a3a34[_0x1b1907(0x110)](),_0xb6ce5=await fetch(apiRequestURLNext),_0xb77af2=await _0xb6ce5[_0x1b1907(0x110)]();nextButton=document['getElementById'](_0x1b1907(0xf8));_0xb77af2[_0x1b1907(0x102)]!=0x0&&(nextButton[_0x1b1907(0x10b)][_0x1b1907(0xfa)]=_0x1b1907(0x106));console[_0x1b1907(0xfb)](_0x9ca94d),html='';let _0x3f17c1='',_0x444ed2='',_0x2970ad='',_0x327df4='',_0x194a06=0x0;for(_0x194a06=0x0;_0x194a06<_0x9ca94d[_0x1b1907(0x102)];_0x194a06++){_0x3f17c1=_0x9ca94d[_0x194a06][0x0],_0x444ed2=_0x9ca94d[_0x194a06][0x1],_0x2970ad=_0x9ca94d[_0x194a06][0x2],_0x327df4=_0x9ca94d[_0x194a06][0x4],html+=createPost(_0x3f17c1,_0x444ed2,_0x327df4,_0x2970ad),(_0x194a06+0x1)%0x2==0x0?(g=document[_0x1b1907(0x111)](_0x1b1907(0xfe)),g['id']=_0x3f17c1,g['className']=_0x1b1907(0x10f),postTag=document[_0x1b1907(0x114)](_0x1b1907(0x10c)),postTag[_0x1b1907(0xff)](g),document[_0x1b1907(0x114)](g['id'])[_0x1b1907(0x118)]=html):html=createPost(_0x3f17c1,_0x444ed2,_0x327df4,_0x2970ad);}_0x194a06+0x1%0x2!=0x0&&(g=document['createElement'](_0x1b1907(0xfe)),g['id']=_0x3f17c1,g[_0x1b1907(0x103)]=_0x1b1907(0x10f),postTag=document[_0x1b1907(0x114)](_0x1b1907(0x10c)),postTag[_0x1b1907(0xff)](g),document[_0x1b1907(0x114)](g['id'])[_0x1b1907(0x118)]=html);}function _0x3bc8(_0xee0b8a,_0xf992f9){const _0x68425=_0x6842();return _0x3bc8=function(_0x3bc8a6,_0x355468){_0x3bc8a6=_0x3bc8a6-0xf4;let _0x1b4435=_0x68425[_0x3bc8a6];return _0x1b4435;},_0x3bc8(_0xee0b8a,_0xf992f9);}function createPost(_0x4b2f42,_0x24361a,_0x958dc1,_0x55381d){const _0x3c2678=_0x4b9c2d;return html=_0x3c2678(0x10e)+_0x55381d+'\x22\x20alt=\x22Card\x20image\x20cap\x22>\x0a\x20\x20<div\x20class=\x22card-body\x22>\x0a\x20\x20\x20\x20<h5\x20class=\x22card-title\x22>'+_0x4b2f42+_0x3c2678(0x101)+_0x24361a+'</p>\x0a\x20\x20\x20\x20<a\x20href=\x22/posts/'+_0x958dc1+'/'+_0x4b2f42+_0x3c2678(0xfc),html;}getArticlesInCat(),window[_0x4b9c2d(0x116)]=function(){const _0x5451b2=_0x4b9c2d;nextButton=document['getElementById'](_0x5451b2(0xf8)),nextButton[_0x5451b2(0x117)](_0x5451b2(0x119),function(){const _0x129651=_0x5451b2;let _0x2e776c=params[_0x129651(0x107)];console[_0x129651(0xfb)]('click'),_0x2e776c=parseInt(_0x2e776c),_0x2e776c+=0x4,window[_0x129651(0x115)]['href']='/articles?cid='+cidValue+_0x129651(0xf9)+_0x2e776c;});};