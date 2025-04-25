var url ="127.0.0.1:8080"
document.getElementById('body1').onload = function()
{
	alert('Hello')	
    CreateTask()
}
function GetTasks()
{

}
function CreateTask(id, name)
{
    Task = document.createElement('div')
    Task.className = 'card'
    document.getElementById('container').appendChild(Task)
    TaskBody = document.createElement('div')
    TaskBody.className = 'card-body'
    Task.appendChild(TaskBody)
    CheckBox = document.createElement('input')
    CheckBox.type = 'checkbox'
    CheckBox.className = 'form-check-input me-1'
    TaskBody.appendChild(CheckBox)
    Label = document.createElement('<label')
    Label.className = 'form-check-label'
    Label.innerHTML = 'Task1'
    TaskBody.appendChild(Label)
    ButtEdit = document.createElement('button')
    ButtEdit.className = 'btn btn-light'
    ButtEdit.textContent = 'Edit'
    ButtEdit.style.float = 'right'
    TaskBody.appendChild(ButtEdit)
    ButtDel = document.createElement('button')
    ButtDel.className = 'btn btn-light'
    ButtDel.textContent = 'Del'
    ButtDel.style.float = 'right'
    TaskBody.appendChild(ButtDel)
}