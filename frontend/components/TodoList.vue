<template>
  <div class="mt-8 bg-white overflow-hidden shadow sm:rounded-lg p-6">
    <h2 class="text-2xl leading-7 font-semibold" >
      <p v-if=!isNameEdit @dblclick="isNameEdit=!isNameEdit" @click="hideDoneTask=!hideDoneTask">{{name}}'s Todo List</p>
      <input v-else type="text" @keyup.enter="isNameEdit=!isNameEdit" v-model="name" class="w-3/6">
    </h2>
    <p class="mt-3 text-gray-600">

      <div v-if="!activeEdit" class="m-2">
        <label for="newtodo">Create Todo:</label>
        <input v-model="newContent" @keyup.enter="createData" class="outline-none" id="newtodo" type="text" placeholder="Your todo here!">
        <hr>
      </div>

      <div v-else class="m-2">
        <label for="newtodo">Edit Todo:</label>
        <input v-model="activeEditData.content" @keyup.enter="unsetEditData(activeEditData)" class="outline-none" id="newtodo" type="text" placeholder="Your todo here!">
        <hr>
      </div>
      

      <ol class="p-2">
          <li v-for="(task, index) in todolist">
            <div v-if="!task.is_done" class="p-1">
              <i @dblclick="setEditData(task)">{{task.content}} |</i>
              <b @click="setDone(task.id, task)">Check</b>
            </div>
            <div v-if="task.is_done && !hideDoneTask">
              <strike  @click="unsetDone(task.id, task)">{{task.content}}</strike>
            </div>
          </li>
      </ol>
    </p>
  </div>
</template>

<script>
export default {
    name: "MyTodoList",
    data(){ return {
      name: "Accalina",
      newContent: "",
      hideDoneTask: false,
      isNameEdit: false,
      activeEditData: {},
      activeEdit: false,
      todolist: []
    }},
    methods: {
      fetchData : function(){
        this.$axios.$get('http://localhost:8080/todo/')
          .then(response => this.todolist = response.data)
      },
      createData: function(){
        this.$axios.$post('http://localhost:8080/todo/', {
          content: this.newContent
        })
          .then(
            response => this.todolist.push(response.data),
            this.newContent = ""
          )
      },
      updateData: function(id, data){
        this.$axios.$put('http://localhost:8080/todo/' + id, {
          content: data.content,
          is_done: data.is_done
        })
          .then(response => console.log(response.data))
      },
      setDone: function(id, task){
        task.is_done = true
        this.updateData(id, task)
      },
      unsetDone: function(id, task){
        task.is_done = false
        this.updateData(id, task)
      },
      setEditData: function(task){
        this.activeEdit = true
        this.activeEditData = task
      },
      unsetEditData: function(task){
        this.updateData(task.id, task)
        this.activeEdit = false
        this.activeEditData = {}
      }
      
    },
    mounted() {
      this.fetchData();
    },
    
}
</script>