import React from "react"

import { Route } from 'react-router-dom'

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

        <Route path="/">
        
          <img src="/images"/>
      
          <form method="POST" action="/todos" >
            <input type="text" id="todo-input" name="todo-name" />
            <input type="submit" value="Create TODO" />
          </form>
    
          <ul>
            { todos.map(todo => (
              <li>
                {todo.text}
              </li>
            ))}
          </ul>
          
        </Route>

        <Route path="/healthz">
          <p>Healthy</p>
        </Route>
  
      </div>
    );
  }
}

export default App;
