let apiRequestURL = `/api/v1/categories`
async function getCats(){
  let response  = await fetch(apiRequestURL);
  let data = await response.json();
  console.log(data)
  html = ""
  let categoryName = ""
  let cetegorySummary = ""
  let imgur = ""
  let id = ""
  let i = 0
  for (i=0; i<data.length; i++){
    categoryName = data[i][0]
    categorySummary = data[i][1]
    imgur = data[i][2]
    id = data[i][4]
    html += createPost(categoryName, categorySummary, id, imgur)
    if((i+1)%2 == 0){
      g= document.createElement('div');
      g.id = categoryName;
      g.className = "row"
      postTag = document.getElementById("posts")
      postTag.appendChild(g);
      document.getElementById(g.id).innerHTML=html;
    }
    else{
      html = createPost(categoryName, categorySummary, id, imgur)
    }
  }
  if(i+1 % 2 != 0){
    g= document.createElement('div');
    g.id = categoryName;
    g.className = "row"
    postTag = document.getElementById("posts")
    postTag.appendChild(g);
    document.getElementById(g.id).innerHTML=html;
  }
}

function createPost(topic, summary, id, imgur){
  html = `
  <div class="card col-sm-6 mt-2" style="padding: 10px;">
  <img class="card-img-top" src="${imgur}" alt="Card image cap">
  <div class="card-body">
    <h5 class="card-title">${topic}</h5>
    <p class="card-text">${summary}</p>
    <a href="/articles?cid=${id}&offset=0" class="btn btn-primary align-self-end">Read articles !</a>
  </div>
  </div>
  <br>`
  return html
}

getCats()
