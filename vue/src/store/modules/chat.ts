import {
  Module,
  VuexModule,
  Mutation,
  Action,
  getModule
} from "vuex-module-decorators";
import store from "../index";

export interface ChatI {
  message: string;
}

const websocket = new WebSocket("ws://localhost:8888/chat");

@Module({ dynamic: true, store: store, name: "chat", namespaced: true })
class Chat extends VuexModule {
  messages: string[] = ["asdf", "asdfasf"];

  @Mutation
  addMessage(mes: string) {
    this.messages.push(mes);
  }

  @Action
  async sendMessage(mes: string) {
    const sendData: ChatI = {
      message: mes
    };
    websocket.send(JSON.stringify(sendData));
  }
}

export const chatModule = getModule(Chat);

websocket.onmessage = function(e: any) {
  const res = JSON.parse(e.data);
  chatModule.addMessage(res.message);
};
