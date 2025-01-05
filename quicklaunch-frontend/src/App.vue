<template>
    <div id="app">
      <command-input></command-input>
    </div>
  </template>
  
  <script>
  import CommandInput from './components/CommandInput.vue';
  
  export default {
    name: 'App',
    components: {
      CommandInput
    }
  };
  </script>