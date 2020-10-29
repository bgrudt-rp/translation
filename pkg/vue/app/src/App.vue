<template>
  <v-app>
    <v-main>
      <HelloWorld/>
    </v-main>
  </v-app>
</template>

<script>
import CodeTypeTest from './components/CodeTypeTest';

export default {
  name: 'App',

  components: {
    CodeTypeTest,
  },

  data: () => ({
    //
  }),
};
</script>
