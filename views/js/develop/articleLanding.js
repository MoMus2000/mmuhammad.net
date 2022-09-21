let currentPage = 1
const params = new Proxy(new URLSearchParams(window.location.search), {
  get: (searchParams, prop) => searchParams.get(prop),
});
let cidValue = params.cid; // "some_value"
let offset = parseInt(params.offset) // "offset value"
let apiRequestURL = `/api/v1/postByCat?cid=${cidValue}&offset=${offset}`
let apiRequestURLNext = `/api/v1/postByCat?cid=${cidValue}&offset=${offset+4}`
console.log(apiRequestURL)
async function getArticlesInCat(){
  let response  = await fetch(apiRequestURL);
  let data = await response.json();

  let nextresponse = await fetch(apiRequestURLNext);
  let nextdata = await nextresponse.json();

  nextButton = document.getElementById("Next")
  if(nextdata.length != 0){
    nextButton.style.display='block';
  }
  console.log(data)
  html = ""
  let topic = ""
  let summary = ""
  let imgur = ""
  let id = ""
  let i = 0
  for (i=0; i<data.length; i++){
    topic = data[i][0]
    summary = data[i][1]
    imgur = data[i][2]
    id = data[i][4]
    html += createPost(topic, summary, id, imgur)
    if((i+1)%2 == 0){
      g= document.createElement('div');
      g.id = topic;
      g.className = "row"
      postTag = document.getElementById("posts")
      postTag.appendChild(g);
      document.getElementById(g.id).innerHTML=html;
    }
    else{
      html = createPost(topic, summary, id, imgur)
    }
  }
  if(i+1 % 2 != 0){
    g= document.createElement('div');
    g.id = topic;
    g.className = "row"
    postTag = document.getElementById("posts")
    postTag.appendChild(g);
    document.getElementById(g.id).innerHTML=html;
  }
}

function createPost(topic, summary, id, imgur){
  html = `
  <div class="card col-sm-6 mt-2">
  <img class="card-img-top" src="${imgur}" alt="Card image cap">
  <div class="card-body">
    <h5 class="card-title">${topic}</h5>
    <p class="card-text">${summary}</p>
    <a href="/posts/${id}/${topic}" class="btn btn-primary">Go to the article</a>
  </div>
  </div>
  <br>`
  return html
}

getArticlesInCat()

window.onload = function(){
  nextButton = document.getElementById("Next")
  nextButton.addEventListener('click', function(){
    let offset = params.offset
    console.log('click')
    offset = parseInt(offset)
    offset += 4
    window.location.href = `/articles?cid=${cidValue}&offset=${offset}`
  })
}