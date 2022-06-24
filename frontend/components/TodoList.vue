<template>
  <div class="mt-8 bg-white overflow-hidden shadow sm:rounded-lg p-6">
    <h2 class="text-2xl leading-7 font-semibold" >
      <p v-if=!isNameEdit @dblclick="isNameEdit=!isNameEdit">{{name}}'s Todo List</p>
      <input v-else type="text" @keyup.enter="isNameEdit=!isNameEdit" v-model="name" class="w-3/6">
    </h2>
    <p class="mt-3 text-gray-600">
      <ol v-for="(task, index) in todolist">
          <li>
            <i v-if="!task.is_done">{{index + 1}}. {{task.content}} |</i>
            <strike v-else>{{index + 1}}. {{task.content}}</strike>

            <b v-if="!task.is_done" @click="setDone(task.id, task)">Check</b>
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
      isNameEdit: false,
      todolist: []
    }},
    methods: {
      fetchData : function(){
        this.$axios.$get('http://localhost:8080/todo/')
          .then(response => this.todolist = response.data)
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
      }
      
    },
    mounted() {
      this.fetchData();
    },
    
}
</script>