var baseURL ="http://127.0.0.1:8080"
document.getElementById('main').onload = function()
{
	console.log('Hello')
  LoadTasks()
}

async function deleteTask(id) 
{
  request = {
    method: "DELETE",
  }
  let response = await fetch(baseURL + "/task/"+id, request);
    if (!response.ok) { 
        alert("Ошибка удаления задачи: " + response.status);
        return
    }
    LoadTasks()
}


function renderTask(element)
{
  let container = document.getElementById('container')
  let taskElement = document.createElement('li')
  taskElement.className = 'list-group-item'
  taskElement.id = element.id
  container.appendChild(taskElement)
  let divcont = document.createElement('div')
  divcont.className = 'item-content'
  taskElement.appendChild(divcont)
  let div1 = document.createElement('div')
  div1.className = "task-name"
  divcont.appendChild(div1)
  let checkbox = document.createElement('input')    
  checkbox.type = 'checkbox'
  checkbox.className = 'form-check-input me-1'


  checkbox.addEventListener('change',function(e) {
    changecheckbox(element.id,e);
  });


  if (element.is_done == true) 
  {
    checkbox.checked = true
  }
  div1.appendChild(checkbox)
  let p = document.createElement('p')
  p.id='taskname' + element.id
  p.innerHTML = element.name
  if (element.is_done == true) 
    {
      p.style.textDecoration = 'line-through'
    }
  div1.appendChild(p)

  let input = document.createElement('input')
  input.id = 'taskinput'+element.id
  input.value = element.name
  input.style.display = 'none'
  input.className = 'form-control'
  div1.appendChild(input)

  let div2 = document.createElement('div')
  div2.style.minWidth = '100px'
  divcont.appendChild(div2)


  
  let buttsave = document.createElement('button')
  buttsave.id = 'buttsave'+element.id
  buttsave.className = 'btn btn-outline-primary btn-sm'
  buttsave.style.display = 'none'
  buttsave.addEventListener('click',function() {
    updatetask(element.id);
  });
  
  buttsave.textContent = 'Save'
  div2.appendChild(buttsave)




  let buttedit = document.createElement('button')
  buttedit.id = 'buttedit'+element.id
  buttedit.className = 'btn btn-outline-primary btn-sm'

  buttedit.addEventListener('click',function() {
    edittask(element.id);
  });
  
  buttedit.textContent = 'Edit'
  div2.appendChild(buttedit)


  buttdel = document.createElement('button')
  buttdel.className = 'btn btn-outline-danger btn-sm'

  buttdel.addEventListener('click',function() {
    deleteTask(element.id);
  });

  buttdel.textContent = 'Delete'
  div2.appendChild(buttdel)
}
async function LoadTasks()
{
  document.getElementById('container').innerHTML = ''
    let response = await fetch(baseURL + "/task");

    if (!response.ok) { // если HTTP-статус в диапазоне 200-299
        alert("Ошибка HTTP: " + response.status);
        return
    }
    // получаем тело ответа (см. про этот метод ниже)
    let json = await response.json();
    // console.log(json)
    json.forEach(element => {
        console.log(element)
        renderTask(element)
        // render item
    });
}

async function createtask() 
{
  json_body = {
    name: document.getElementById('newtask').value
  }
  request = {
    method: "POST",
    headers: {
      'Content-type': 'application/json'
    },
    body: JSON.stringify(json_body)
}
  let response = await fetch(baseURL + "/task", request);
    if (response.status != 201 ) { 
        alert("Ошибка добавления задачи : " + response.status);
        return
    }
    document.getElementById('newtask').value = ''
    LoadTasks()
}

async function changecheckbox(id,e)
{
  if (e.target.checked == true)
  {
    request = {
      method: "PUT"
    }
    let response = await fetch(baseURL + "/task/"+id+"/set-done", request);
      if (response.status != 200 ) { 
          alert("Ошибка изменения статуса задачи : " + response.status);
          return
      }
  }
  else
  {
    request = {
      method: "PUT"
    }
    let response = await fetch(baseURL + "/task/"+id+"/unset-done", request);
      if (response.status != 200 ) { 
          alert("Ошибка изменения статуса задачи : " + response.status);
          return
      }
  }
  LoadTasks()
}

function edittask(id)
{
  document.getElementById('taskinput'+id).style.display = 'inline-block'
  document.getElementById('buttedit'+id).style.display = 'none'
  document.getElementById('taskname'+id).style.display = 'none'
  document.getElementById('buttsave'+id).style.display = 'inline-block'
}
async function updatetask(id)
{
  if (document.getElementById('taskname'+id).innerHTML == document.getElementById('taskinput'+id).value)
  {
    document.getElementById('taskinput'+id).style.display = 'none'
    document.getElementById('buttedit'+id).style.display = 'inline-block'
    document.getElementById('taskname'+id).style.display = 'inline'
    document.getElementById('buttsave'+id).style.display = 'none'
  }
  else
  {
    json_body = {
      name: document.getElementById('taskinput'+id).value
    }
    request = {
      method: "PUT",
      headers: {
        'Content-type': 'application/json'
      },
      body: JSON.stringify(json_body)
    }
    let response = await fetch(baseURL + "/task/"+id, request);
      if (response.status != 200 ) { 
          alert("Ошибка обновления задачи : " + response.status);
          return
      }
      LoadTasks()
  }
}