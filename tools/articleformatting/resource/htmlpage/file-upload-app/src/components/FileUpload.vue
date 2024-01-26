<template>
  <div>
    <input type="file" @change="onFileChange">
    <button @click="upload">Upload</button>
    <div v-if="response">{{ response }}</div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      selectedFile: null,
      response: null
    };
  },
  methods: {
    onFileChange(e) {
      this.selectedFile = e.target.files[0];
    },
    upload() {
      const formData = new FormData();
      formData.append('file', this.selectedFile);

      axios.post('http://localhost:18080/format-article-file', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }).then(response => {
        this.response = response.data;
        console.log("response", response)
      }).catch(error => {
        // handle error
        alert('发生错误: ' + error.message);
      });
    }
  }
};
</script>
