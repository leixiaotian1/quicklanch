<template>
    <div>
      <input v-model="command" placeholder="输入软件名称" @keyup.enter="sendCommand"/>
      <button @click="sendCommand">启动软件</button>
      <p v-if="feedback">{{ feedback }}</p>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        command: '',
        feedback: ''
      };
    },
    methods: {
      sendCommand() {
        this.$axios.post('/api/launch', { command: this.command })
          .then(response => {
            this.feedback = `软件启动成功: ${response.data}`;
          })
          .catch(error => {
            this.feedback = `启动失败: ${error}`;
          });
      }
    }
  };
  </script>