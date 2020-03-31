<template>
  <div class="home">
    <div class="container">
      <div class="input-group my-3">
        <input
          type="text"
          class="form-control"
          v-model="mes"
          @keyup.enter="send"
        />
        <div class="input-group-append">
          <button class="btn-secondary" @click="send">submit</button>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <ul>
            <li v-for="(message, idx) in messages" :key="message + idx">
              {{ message }} - {{ idx }}
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { chatModule } from "@/store/modules/chat";

@Component({ components: {} })
export default class Home extends Vue {
  mes = "";

  send() {
    if (this.mes != "") {
      chatModule.sendMessage(this.mes);
      this.mes = "";
    }
  }

  get messages(): string[] {
    return chatModule.messages;
  }
}
</script>
