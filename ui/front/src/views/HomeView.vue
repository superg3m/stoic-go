<script setup>
import {onMounted, ref} from 'vue';
import Draggable from 'vuedraggable';
import { UserStore } from "@/UserStore.ts";

const newTodoText = ref('');
const todos = ref([]);
const inProgressTodos = ref([]);
const doneTodos = ref([]);

// status: 0, 1, 2
// 0: todo
// 1: in-progress
// 2: done

const addTodo = async () => {
  if (!newTodoText.value.trim()) {
    return
  }

  let newID = -1;

  try {
    const todoRequest = {
      "OwnerID": UserStore.User.ID,
      "Message": newTodoText.value.trim(),
      "Status": 0
    };

    const response = await fetch("http://localhost:8080/TodoItem", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(todoRequest)
    });

    const jsonRes = await response.json();
    newID = jsonRes.ID
  } catch (e) {
    console.warn(e)
  }

  todos.value.push({
    ID: newID,
    Message: newTodoText.value.trim(),
    OwnerID: UserStore.User.ID,
    Status: 0
  });
  newTodoText.value = '';
};

async function onChangeDone(event){
  if (event.added) {
    const todo = event.added.element
    await tryUpdate(todo, 2)
  }
}

async function onChangeinProgress(event){
  if (event.added) {
    const todo = event.added.element
    await tryUpdate(todo, 1)
  }
}

async function onChangeTodo(event){
  if (event.added) {
    const todo = event.added.element
    await tryUpdate(todo, 0)
  }
}

async function tryUpdate(todo, status) {
  try {
    todo.Status = status;
    await fetch("http://localhost:8080/TodoItem", {
      method: "PATCH",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(todo)
    })

  } catch(e){
    console.error(e)
  }
}

const removeTodo = async (todo, list) => {
  const index = list.findIndex(t => t.ID === todo.ID);
  if (index === -1) {
    return
  }

  try {
    const todoRequest = {
      "ID": todo.ID,
    };

    await fetch("http://localhost:8080/TodoItem", {
      method: "DELETE",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(todoRequest)
    });
  } catch (e) {
    console.warn(e)
  }

  list.splice(index, 1);
};

onMounted(async () => {
  try {
    const response = await fetch(`http://localhost:8080/TodoItem?OwnerID=${UserStore.User.ID}`, {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const todosResponse = await response.json()
    for (let i = 0; i < todosResponse.length; i++) {
      const todo = todosResponse[i];
      if (todo.Status === 0) {
        todos.value.push(todo)
      } else if (todo.Status === 1) {
        inProgressTodos.value.push(todo)
      } else {
        doneTodos.value.push(todo)
      }
    }

  } catch (e) {
    console.warn(e)
  }
})

</script>

<template>
  <div class="todo-app">
    <div class="add-todo-section">
      <input
        v-model="newTodoText"
        @keyup.enter="addTodo"
        placeholder="Enter a new todo"
        class="todo-input"
      />
      <Button @click="addTodo">Add Todo</Button>
    </div>

    <div class="lists-container">
      <div class="todo-list">
        <h2>Todo</h2>
        <Draggable
          v-model="todos"
          group="todos"
          item-key="id"
          class="drag-area"
          ghostClass="ghost"
          @change="onChangeTodo"
        >
          <template #item="{ element }">
            <div class="todo-card">
              <span>{{ element.Message }}</span>
              <button
                @click="removeTodo(element, todos)"
                class="delete-btn"
              >
                ✕
              </button>
            </div>
          </template>
        </Draggable>
      </div>

      <div class="todo-list">
        <h2>In Progress</h2>
        <Draggable
          v-model="inProgressTodos"
          group="todos"
          item-key="id"
          class="drag-area"
          ghostClass="ghost"
          @change="onChangeinProgress"
        >
          <template #item="{ element }">
            <div class="todo-card">
              <span>{{ element.Message }}</span>
              <button
                @click="removeTodo(element, inProgressTodos)"
                class="delete-btn"
              >
                ✕
              </button>
            </div>
          </template>
        </Draggable>
      </div>
    </div>

    <div class="done-list">
      <h2>Done</h2>
      <Draggable
        v-model="doneTodos"
        group="todos"
        item-key="id"
        class="done-drag-area"
        ghostClass="ghost"
        @change="onChangeDone"
      >
        <template #item="{ element }">
          <div class="todo-card">
            <span>{{ element.Message }}</span>
            <button
              @click="removeTodo(element, doneTodos)"
              class="delete-btn"
            >
              ✕
            </button>
          </div>
        </template>
      </Draggable>
    </div>
  </div>
</template>



<style scoped>
span {
  color: #333;
  font-size: 16px;
}

h2 {
  color: #222;
  font-size: 18px;
  text-align: center;
}

.todo-app {
  max-width: 50vw;
  height: fit-content;
  max-height: 90vh;
  margin: 40px auto;
  padding: 20px;
  font-family: Arial, sans-serif;
  background: lightgray;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.add-todo-section {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

.todo-input {
  flex-grow: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
}

.lists-container {
  display: flex;
  justify-content: center;
  gap: 32px;
}

.done-list {
  width: 100%;
  background: white;
  padding: 16px;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  height: fit-content;
}

.todo-list {
  width: 100%;
  background: white;
  padding: 16px;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.done-drag-area {
  height: fit-content;
  padding: 10px;
  border: 2px dashed #ccc;
  border-radius: 6px;
  max-height: 244px;
  overflow-y: scroll;
}

.drag-area {
  min-height: 200px;
  max-height: 328px;
  height: fit-content;
  overflow-y: scroll;
  padding: 10px;
  border: 2px dashed #ccc;
  border-radius: 6px;
}

.todo-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #e6e6e6;
  padding: 12px;
  margin-bottom: 8px;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  cursor: grab;
  transition: transform 0.1s, background 0.2s;
}

.todo-card:hover {
  background: #cccccc;
  transform: scale(1.02);
}

.delete-btn {
  background: #ff4d4d;
  color: white;
  border: none;
  border-radius: 50%;
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 14px;
  transition: background 0.2s;
}

.delete-btn:hover {
  background: #e60000;
}

.ghost {
  background-color: rgba(76, 175, 80, 0.35) !important;
}
</style>
