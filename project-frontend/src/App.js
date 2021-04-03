import React from "react"

class App extends React.Component {
  constructor() {
    super()
    this.state = {
      error: null,
      isLoaded: false,
      todos: []
    }
  }

  componentDidMount() {
    fetch("/todos")
      .then(res => res.json())
      .then(
        (result) => {
          var todoList = JSON.parse(result)
          this.setState({
            isLoaded: true,
            todos: todoList
          })
        },
        (error) => {
          this.setState({
            isLoaded: true,
            error
          })
        }
      )
    this.setState({
      isLoaded: true,
      todos: [{id: 3, text: "test", done: true}, {text: "test2", done: false}]
    })
  }

  handleTodoStatusChanged(event) {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const id = target.name;

    const namespace = process.env.REACT_APP_NAMESPACE

    fetch(`/todos/${id}`, {
      method: "PUT",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        id: id,
        done: value
      })
    }).then(res => {
      if(res.status == 200) {
        target.checked = value
      }
    })
  }

  render() {
    const { error, isLoaded, todos} = this.state
    if (error) {
      return <div>Error: {error.message} </div>
    } else if (!isLoaded) { 
      return <div>Loading...</div>
    }

    return (

      <div className="App">
        
        <img src="/images"/>
    
        <form method="POST" action="/todos" >
          <input type="text" id="todo-input" name="todo-name" />
          <input type="submit" value="Create TODO" />
        </form>
  
        <ul>
          { todos.map(todo => (
            <li>
              <form>
                <label>{todo.text} </label>
                <input name={todo.id} type="checkbox" checked={todo.done} onChange={this.handleTodoStatusChanged}/>
              </form>
            </li>
          ))}
        </ul>
  
      </div>
    );
  }
}

export default App;
